package generate

import (
	"fmt"
	"github.com/JsonLee12138/json-server/core"
	"github.com/JsonLee12138/json-server/embed"
	"go/ast"
	"go/token"
	"strings"
)

func GenerateRepository(repositoryName, output string, override bool) error {
	return core.TryCatchVoid(func() {
		tmplFile := string(core.Raise[[]byte](embed.TemplatesPath.ReadFile("templates/repository.tmpl")))
		upperName := core.UpperCamelCase(repositoryName)
		outputPath := fmt.Sprintf("%s/repository/%s_repository.go", output, repositoryName)
		params := map[string]string{
			"Name": upperName,
		}
		core.RaiseVoid(GenerateFileExistsHandler(outputPath, tmplFile, params, override))
		fmt.Printf("✅ '%s' repository has been successfully generated!\n", repositoryName)
	}, core.DefaultErrorHandler)
}

func InjectRepository(repositoryName, output string) error {
	return core.TryCatchVoid(func() {
		entryPath := fmt.Sprintf("%s/entry.go", output)
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
		upperRepositoryName := core.UpperCamelCase(repositoryName)
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
								X:   ast.NewIdent("repository"),
								Sel: ast.NewIdent(fmt.Sprintf("New%sRepository", upperRepositoryName)),
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
					Value: fmt.Sprintf("\"%s/repository\"", output),
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
		fmt.Printf("✅ '%s' repository has been successfully injected!\n", repositoryName)
	}, core.DefaultErrorHandler)
}

func GenerateInjectRepository(repositoryName, output string, override bool) error {
	return core.TryCatchVoid(func() {
		core.RaiseVoid(GenerateRepository(repositoryName, output, override))
		core.RaiseVoid(InjectRepository(repositoryName, output))
	}, core.DefaultErrorHandler)
}
