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

func InjectController(controllerName, output string) error {
	return utils.TryCatchVoid(func() {
		entryPath := fmt.Sprintf("%s/entry.go", output)
		fmt.Println("entryPath:", entryPath)
		isDir, exists, err := utils.Exists(entryPath)
		if err != nil {
			utils.Throw(err)
		}
		if !exists || isDir {
			//core.RaiseVoid(core.OnlyCreateFile(entryPath))
			utils.RaiseVoid(GenerateEntry(output))
		}
		file, _, err := utils.ParseFile(entryPath)
		if err != nil {
			utils.Throw(err)
		}
		pathArr := strings.Split(output, "/")
		moduleName := pathArr[len(pathArr)-1]
		upperModuleName := utils.UpperCamelCase(moduleName)
		upperControllerName := utils.UpperCamelCase(controllerName)
		funs := utils.Raise(utils.FindFunctions(file, ".*ModuleSetup$", utils.RegexMatch))
		content := &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent("container"),
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
		routerContent := &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ExprStmt{
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
				},
			},
		}
		fmt.Println("routerContent:", routerContent)
		if len(funs) == 0 {
			file.Imports = append(file.Imports, &ast.ImportSpec{
				Name: &ast.Ident{
					Name: "dig",
				},
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf("\"%s/service\"", output),
				},
			})
			params := &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							ast.NewIdent("container"),
						},
						Type: &ast.StarExpr{
							X: &ast.SelectorExpr{
								X: ast.NewIdent("dig"),
								Sel: &ast.Ident{
									Name: "Container",
								},
							},
						},
					},
				},
			}
			utils.CreateFunc(file, fmt.Sprintf("%sModuleSetup", upperModuleName), params, content)
		} else {
			fn := funs[0].Node.(*ast.FuncDecl)
			//fn.Body.List = append(fn.Body.List, content.List...)
			for _, stmt := range fn.Body.List {
				if retStmt, ok := stmt.(*ast.ReturnStmt); ok {
					fmt.Println("returnStmt:", retStmt)
					if len(retStmt.Results) == 1 {
						if callExpr, ok := retStmt.Results[0].(*ast.CallExpr); ok {
							if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
								if selExpr.Sel.Name == "TryCatchVoid" {
									if ident, ok := selExpr.X.(*ast.Ident); ok && ident.Name == "core" {
										//return true
										fmt.Println("trycatch:", ident)

									}
								}
							}
						}
					}
					//if callExpr, ok := returnStmt.X.(*ast.CallExpr); ok {
					//	fmt.Println("callExpr:", callExpr)
					//	if isRaiseVoidWithInvoke(callExpr) {
					//		fmt.Println("i:", i)
					//		fmt.Println("callExpr:", callExpr)
					//		fn.Body.List = append(fn.Body.List[:i], routerContent.List[i:]...)
					//		// Step 1: Insert new `scope.Provide` before this statement
					//		//injection := createInjectionStmt()
					//		//body.List = append(body.List[:i], append([]ast.Stmt{injection}, body.List[i:]...)...)
					//		//
					//		//// Step 2: Modify `scope.Invoke` body
					//		//modifyInvokeBody(callExpr)
					//		break
					//	}
					//}
				}
			}
		}
		utils.RaiseVoid(utils.WriteToFile(file, entryPath))
		fmt.Printf("✅ '%s' controller has been successfully injected!\n", controllerName)
	}, utils.DefaultErrorHandler)
}

func isRaiseVoidWithInvoke(callExpr *ast.CallExpr) bool {
	if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok && selExpr.Sel.Name == "RaiseVoid" {
		if len(callExpr.Args) == 1 {
			if innerCall, ok := callExpr.Args[0].(*ast.CallExpr); ok {
				if innerSel, ok := innerCall.Fun.(*ast.SelectorExpr); ok {
					return innerSel.Sel.Name == "Invoke"
				}
			}
		}
	}
	return false
}

func GenerateInjectController(controllerName, output string, override bool, moduleName string) error {
	return utils.TryCatchVoid(func() {
		utils.RaiseVoid(GenerateController(controllerName, output, override, moduleName))
		utils.RaiseVoid(InjectController(controllerName, output))
	}, utils.DefaultErrorHandler)
}
