package generate

import (
	"fmt"
	"github.com/JsonLee12138/json-server/core"
	"github.com/JsonLee12138/json-server/embed"
	"go/ast"
	"go/token"
	"strings"
)

func GenerateController(controllerName, output string, override bool, moduleName string) error {
	return core.TryCatchVoid(func() {
		tmplFile := string(core.Raise(embed.TemplatesPath.ReadFile("templates/controller.tmpl")))
		outputPath := fmt.Sprintf("%s/controller/%s_controller.go", output, controllerName)
		upperName := core.UpperCamelCase(controllerName)
		pkgPath := core.Raise(core.GetModuleFullPath(moduleName))
		params := map[string]string{
			"Name":    upperName,
			"PkgPath": pkgPath,
		}
		core.RaiseVoid(GenerateFileExistsHandler(outputPath, tmplFile, params, override))
		fmt.Printf("✅ '%s' controller has been successfully generated!\n", controllerName)
	}, core.DefaultErrorHandler)
}

func InjectController(controllerName, output string) error {
	return core.TryCatchVoid(func() {
		entryPath := fmt.Sprintf("%s/entry.go", output)
		fmt.Println("entryPath:", entryPath)
		isDir, exists, err := core.Exists(entryPath)
		if err != nil {
			core.Throw(err)
		}
		if !exists || isDir {
			//core.RaiseVoid(core.OnlyCreateFile(entryPath))
			core.RaiseVoid(GenerateEntry(output))
		}
		file, _, err := core.ParseFile(entryPath)
		if err != nil {
			core.Throw(err)
		}
		pathArr := strings.Split(output, "/")
		moduleName := pathArr[len(pathArr)-1]
		upperModuleName := core.UpperCamelCase(moduleName)
		upperControllerName := core.UpperCamelCase(controllerName)
		funs := core.Raise(core.FindFunctions(file, ".*ModuleSetup$", core.RegexMatch))
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
			core.CreateFunc(file, fmt.Sprintf("%sModuleSetup", upperModuleName), params, content)
		} else {
			fn := funs[0].Node.(*ast.FuncDecl)
			fn.Body.List = append(fn.Body.List, content.List...)
		}
		core.RaiseVoid(core.WriteToFile(file, entryPath))
		fmt.Printf("✅ '%s' controller has been successfully injected!\n", controllerName)
	}, core.DefaultErrorHandler)
}

func GenerateInjectController(controllerName, output string, override bool, moduleName string) error {
	return core.TryCatchVoid(func() {
		core.RaiseVoid(GenerateController(controllerName, output, override, moduleName))
		core.RaiseVoid(InjectController(controllerName, output))
	}, core.DefaultErrorHandler)
}
