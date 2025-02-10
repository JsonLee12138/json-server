package generate

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"github.com/JsonLee12138/json-server/embed"
	"github.com/JsonLee12138/json-server/pkg/utils"
)

func GenerateRepository(repositoryName, output string, override bool) error {
	return utils.TryCatchVoid(func() {
		tmplFile := string(utils.Raise[[]byte](embed.TemplatesPath.ReadFile("templates/repository.tmpl")))
		upperName := utils.UpperCamelCase(repositoryName)
		outputPath := fmt.Sprintf("%s/repository/%s_repository.go", output, repositoryName)
		params := map[string]string{
			"Name": upperName,
		}
		utils.RaiseVoid(GenerateFileExistsHandler(outputPath, tmplFile, params, override))
		fmt.Printf("✅ '%s' repository has been successfully generated!\n", repositoryName)
	}, utils.DefaultErrorHandler)
}

func InjectRepository(repositoryName, output string) error {
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
		upperRepositoryName := utils.UpperCamelCase(repositoryName)
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
			utils.CreateFunc(file, fmt.Sprintf("%sModuleSetup", upperModuleName), params, content)
		} else {
			fn := funs[0].Node.(*ast.FuncDecl)
			fn.Body.List = append(fn.Body.List, content.List...)
		}
		utils.RaiseVoid(utils.WriteToFile(file, entryPath))
		fmt.Printf("✅ '%s' repository has been successfully injected!\n", repositoryName)
	}, utils.DefaultErrorHandler)
}

func GenerateInjectRepository(repositoryName, output string, override bool) error {
	return utils.TryCatchVoid(func() {
		utils.RaiseVoid(GenerateRepository(repositoryName, output, override))
		utils.RaiseVoid(InjectRepository(repositoryName, output))
	}, utils.DefaultErrorHandler)
}
