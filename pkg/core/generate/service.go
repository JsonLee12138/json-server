package generate

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/JsonLee12138/jsonix/embed"
	"github.com/JsonLee12138/jsonix/pkg/core"
	"github.com/JsonLee12138/jsonix/pkg/utils"
)

func GenerateService(serviceName, output string, override bool, moduleName string) error {
	return utils.TryCatchVoid(func() {
		tmplFile := string(utils.Raise(embed.TemplatesPath.ReadFile("templates/service.tmpl")))
		outputPath := fmt.Sprintf("%s/service/%s.go", output, serviceName)
		upperName := utils.UpperCamelCase(serviceName)
		pkgPath := utils.Raise(utils.GetModuleFullPath(moduleName))
		params := map[string]string{
			"Name":    upperName,
			"PkgPath": pkgPath,
		}
		utils.RaiseVoid(core.GenerateFileExistsHandler(outputPath, tmplFile, params, override))
		fmt.Printf("✅ '%s' service has been successfully generated!\n", serviceName)
	}, utils.DefaultErrorHandler)
}

func GenerateProvideServiceHandler(serviceName string) *ast.ExprStmt {
	upperServiceName := utils.UpperCamelCase(serviceName)
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
							X:   ast.NewIdent("service"),
							Sel: ast.NewIdent(fmt.Sprintf("New%sService", upperServiceName)),
						},
					},
				},
			},
		},
	}
}

func GenerateProvideServiceFn(serviceName string, outPath string) ([]*ast.ImportSpec, *ast.FuncDecl) {
	imports := []*ast.ImportSpec{
		{
			Name: &ast.Ident{
				Name: "dig",
			},
		},
		{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: "github.com/JsonLee12138/jsonix/pkg/utils",
			},
		},
		{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("\"%s/service\"", outPath),
			},
		},
	}
	provideServiceHandler := GenerateProvideServiceHandler(serviceName)
	fn := &ast.FuncDecl{
		Name: ast.NewIdent("ProvideService"),
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
											provideServiceHandler,
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

func InjectService(serviceName, output string) error {
	return utils.TryCatchVoid(func() {
		entryPath := fmt.Sprintf("%s/entry.go", output)
		isDir, exists, err := utils.Exists(entryPath)
		if err != nil {
			utils.Throw(err)
		}
		if !exists || isDir {
			utils.RaiseVoid(GenerateEntry(output))
		}
		file, _, err := utils.ParseFile(entryPath)
		utils.RaiseVoid(err)

		provideServiceFns := utils.Raise(utils.FindFunctions(file, "ProvideService", utils.ExactMatch))
		if len(provideServiceFns) == 0 {
			provideServiceImports, provideServiceFn := GenerateProvideServiceFn(serviceName, output)
			file.Imports = utils.UniqueImports(append(file.Imports, provideServiceImports...))
			file.Decls = append(file.Decls, provideServiceFn)
		} else {
			serviceContent := GenerateProvideServiceHandler(serviceName)
			provideServiceFn := provideServiceFns[0].Node.(*ast.FuncDecl)
			for _, stmt := range provideServiceFn.Body.List {
				if returnStmt, ok := stmt.(*ast.ReturnStmt); ok {
					for _, result := range returnStmt.Results {
						if callExpr, ok := result.(*ast.CallExpr); ok {
							if inTryCatchVoid(callExpr) {
								if fn, ok := callExpr.Args[0].(*ast.FuncLit); ok {
									fn.Body.List = append(fn.Body.List, serviceContent)
								}
							}
						}
					}
				}
			}
		}

		utils.RaiseVoid(utils.WriteToFile(file, entryPath))
		fmt.Printf("✅ '%s' service has been successfully injected!\n", serviceName)
	}, utils.DefaultErrorHandler)
}

func GenerateInjectService(serviceName, output string, override bool, moduleName string) error {
	return utils.TryCatchVoid(func() {
		utils.RaiseVoid(GenerateService(serviceName, output, override, moduleName))
		utils.RaiseVoid(InjectService(serviceName, output))
	}, utils.DefaultErrorHandler)
}
