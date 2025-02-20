package generate

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"github.com/JsonLee12138/json-server/embed"
	"github.com/JsonLee12138/json-server/pkg/core"
	"github.com/JsonLee12138/json-server/pkg/utils"
)

func GenerateController(controllerName, output string, override bool, moduleName string) error {
	return utils.TryCatchVoid(func() {
		tmplFile := string(utils.Raise(embed.TemplatesPath.ReadFile("templates/controller.tmpl")))
		outputPath := fmt.Sprintf("%s/controller/%s_controller.go", output, controllerName)
		upperName := utils.UpperCamelCase(controllerName)
		pkgPath := utils.Raise(utils.GetModuleFullPath(moduleName))
		params := map[string]string{
			"Name":    upperName,
			"PkgPath": pkgPath,
		}
		utils.RaiseVoid(core.GenerateFileExistsHandler(outputPath, tmplFile, params, override))
		fmt.Printf("✅ '%s' controller has been successfully generated!\n", controllerName)
	}, utils.DefaultErrorHandler)
}

func GenerateProvideControllerHandler(controllerName string) *ast.ExprStmt {
	upperControllerName := utils.UpperCamelCase(controllerName)
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent("utils"),
				Sel: ast.NewIdent("RaiseVoid"),
			},
			Args: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   ast.NewIdent("scope"),
						Sel: ast.NewIdent("Provide"),
					},
					Args: []ast.Expr{
						&ast.SelectorExpr{
							X:   ast.NewIdent("controller"),
							Sel: ast.NewIdent(fmt.Sprintf("New%sController", upperControllerName)),
						},
					},
				},
			},
		},
	}
}

func GenerateRouterHandler(controllerName string) *ast.ExprStmt {
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent("group"),
				Sel: ast.NewIdent("Get"),
			},
			Args: []ast.Expr{
				&ast.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%s"`, controllerName),
				},
				&ast.SelectorExpr{
					X:   ast.NewIdent(fmt.Sprintf("%sController", controllerName)),
					Sel: ast.NewIdent("HelloWorld"),
				},
			},
		},
	}
}

func GenerateProvideControllerFn(controllerName string, outPath string) ([]*ast.ImportSpec, *ast.FuncDecl) {
	imports := []*ast.ImportSpec{
		{
			Name: &ast.Ident{
				Name: "dig",
			},
		},
		{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: "github.com/JsonLee12138/json-server/pkg/utils",
			},
		},
		{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("\"%s/controller\"", outPath),
			},
		},
	}
	provideControllerHandler := GenerateProvideControllerHandler(controllerName)
	fn := &ast.FuncDecl{
		Name: ast.NewIdent("ProvideController"),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							ast.NewIdent("scope"),
						},
						Type: &ast.StarExpr{
							X: &ast.SelectorExpr{
								X:   ast.NewIdent("dig"),
								Sel: ast.NewIdent("Scope"),
							},
						},
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: ast.NewIdent("error"),
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   ast.NewIdent("utils"),
								Sel: ast.NewIdent("RaiseVoid"),
							},
							Args: []ast.Expr{
								&ast.FuncLit{
									Type: &ast.FuncType{
										Params: &ast.FieldList{
											List: []*ast.Field{},
										},
									},
									Body: &ast.BlockStmt{
										List: []ast.Stmt{
											provideControllerHandler,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	return imports, fn
}

func GenerateRouterFn(controllerName string, outPath string) ([]*ast.ImportSpec, *ast.FuncDecl) {
	imports := []*ast.ImportSpec{
		{
			Name: ast.NewIdent("fiber"),
		},
		{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("\"%s/controller\"", outPath),
			},
		},
	}
	upperControllerName := utils.UpperCamelCase(controllerName)
	routerHandler := GenerateRouterHandler(controllerName)
	pathArr := strings.Split(outPath, "/")
	moduleName := pathArr[len(pathArr)-1]
	fn := &ast.FuncDecl{
		Name: ast.NewIdent("RouterSetup"),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							ast.NewIdent("app"),
						},
						Type: &ast.StarExpr{
							X: &ast.SelectorExpr{
								X:   ast.NewIdent("fiber"),
								Sel: ast.NewIdent("App"),
							},
						},
					},
					{
						Names: []*ast.Ident{
							ast.NewIdent(fmt.Sprintf("%sController", controllerName)),
						},
						Type: &ast.StarExpr{
							X: &ast.SelectorExpr{
								X:   ast.NewIdent("controller"),
								Sel: ast.NewIdent(fmt.Sprintf("%sController", upperControllerName)),
							},
						},
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						ast.NewIdent("group"),
					},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   ast.NewIdent("app"),
								Sel: ast.NewIdent("Group"),
							},
							Args: []ast.Expr{
								&ast.BasicLit{
									Kind:  token.STRING,
									Value: fmt.Sprintf(`"%s"`, moduleName),
								},
							},
						},
					},
				},
				routerHandler,
			},
		},
	}
	return imports, fn
}

func InjectController(controllerName, output string) error {
	return utils.TryCatchVoid(func() {
		entryPath := fmt.Sprintf("%s/entry.go", output)
		fmt.Println("entryPath:", entryPath)
		isDir, exists, err := utils.Exists(entryPath)
		if err != nil {
			utils.Throw(err)
		}
		if !exists || isDir {
			utils.RaiseVoid(GenerateEntry(output))
		}
		file, _, err := utils.ParseFile(entryPath)
		utils.RaiseVoid(err)

		provideControllerFns := utils.Raise(utils.FindFunctions(file, "ProvideController", utils.ExactMatch))
		if len(provideControllerFns) == 0 {
			provideControllerImports, provideControllerFn := GenerateProvideControllerFn(controllerName, output)
			file.Imports = utils.UniqueImports(append(file.Imports, provideControllerImports...))
			file.Decls = append(file.Decls, provideControllerFn)
		} else {
			controllerContent := GenerateProvideControllerHandler(controllerName)
			provideControllerFn := provideControllerFns[0].Node.(*ast.FuncDecl)
			for _, stmt := range provideControllerFn.Body.List {
				if returnStmt, ok := stmt.(*ast.ReturnStmt); ok {
					for _, result := range returnStmt.Results {
						if callExpr, ok := result.(*ast.CallExpr); ok {
							if inTryCatchVoid(callExpr) {
								if fn, ok := callExpr.Args[0].(*ast.FuncLit); ok {
									fn.Body.List = append(fn.Body.List, controllerContent)
								}
							}
						}
					}
				}
			}
		}

		routerFns := utils.Raise(utils.FindFunctions(file, "RouterSetup", utils.ExactMatch))
		if len(routerFns) == 0 {
			routerImports, routerFn := GenerateRouterFn(controllerName, output)
			file.Imports = utils.UniqueImports(append(file.Imports, routerImports...))
			file.Decls = append(file.Decls, routerFn)
		} else {
			routerContent := GenerateRouterHandler(controllerName)
			routerFn := routerFns[0].Node.(*ast.FuncDecl)
			upperControllerName := utils.UpperCamelCase(controllerName)
			newParams := []*ast.Field{
				{
					Names: []*ast.Ident{
						ast.NewIdent(fmt.Sprintf("%sController", controllerName)),
					},
					Type: &ast.StarExpr{
						X: &ast.SelectorExpr{
							X:   ast.NewIdent("controller"),
							Sel: ast.NewIdent(fmt.Sprintf("%sController", upperControllerName)),
						},
					},
				},
			}
			routerFn.Type.Params.List = append(routerFn.Type.Params.List, newParams...)
			routerFn.Body.List = append(routerFn.Body.List, routerContent)
		}

		utils.RaiseVoid(utils.WriteToFile(file, entryPath))
		fmt.Printf("✅ '%s' controller has been successfully injected!\n", controllerName)
	}, utils.DefaultErrorHandler)
}

func inTryCatchVoid(callExpr *ast.CallExpr) bool {
	if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok && selExpr.Sel.Name == "TryCatchVoid" {
		return true
	}
	return false
}

func GenerateInjectController(controllerName, output string, override bool, moduleName string) error {
	return utils.TryCatchVoid(func() {
		utils.RaiseVoid(GenerateController(controllerName, output, override, moduleName))
		utils.RaiseVoid(InjectController(controllerName, output))
	}, utils.DefaultErrorHandler)
}
