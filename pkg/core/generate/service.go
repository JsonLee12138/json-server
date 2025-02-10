package generate

import (
	"fmt"
	"github.com/JsonLee12138/json-server/embed"
	"github.com/JsonLee12138/json-server/pkg/utils"
	"go/ast"
	"go/token"
	"strings"
)

func GenerateService(serviceName, output string, override bool, moduleName string) error {
	return utils.TryCatchVoid(func() {
		tmplFile := string(utils.Raise(embed.TemplatesPath.ReadFile("templates/service.tmpl")))
		outputPath := fmt.Sprintf("%s/service/%s_service.go", output, serviceName)
		upperName := utils.UpperCamelCase(serviceName)
		pkgPath := utils.Raise(utils.GetModuleFullPath(moduleName))
		params := map[string]string{
			"Name":    upperName,
			"PkgPath": pkgPath,
		}
		utils.RaiseVoid(GenerateFileExistsHandler(outputPath, tmplFile, params, override))
		fmt.Printf("✅ '%s' service has been successfully generated!\n", serviceName)
	}, utils.DefaultErrorHandler)
}

func InjectService(serviceName, output string) error {
	return utils.TryCatchVoid(func() {
		entryPath := fmt.Sprintf("%s/entry.go", output)
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
		upperServiceName := utils.UpperCamelCase(serviceName)
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
								X:   ast.NewIdent("service"),
								Sel: ast.NewIdent(fmt.Sprintf("New%sService", upperServiceName)),
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
			utils.CreateFunc(file, fmt.Sprintf("%sModuleSetup", upperModuleName), params, content)
		} else {
			fn := funs[0].Node.(*ast.FuncDecl)
			fn.Body.List = append(fn.Body.List, content.List...)
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
