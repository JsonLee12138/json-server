package generate

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/JsonLee12138/json-server/embed"
	"github.com/JsonLee12138/json-server/pkg/core"
	"github.com/JsonLee12138/json-server/pkg/utils"
)

func GenerateRepository(repositoryName, output string, override bool) error {
	return utils.TryCatchVoid(func() {
		tmplFile := string(utils.Raise(embed.TemplatesPath.ReadFile("templates/repository.tmpl")))
		upperName := utils.UpperCamelCase(repositoryName)
		outputPath := fmt.Sprintf("%s/repository/%s_repository.go", output, repositoryName)
		params := map[string]string{
			"Name": upperName,
		}
		utils.RaiseVoid(core.GenerateFileExistsHandler(outputPath, tmplFile, params, override))
		fmt.Printf("✅ '%s' repository has been successfully generated!\n", repositoryName)
	}, utils.DefaultErrorHandler)
}

func GenerateProvideRepositoryHandler(repositoryName string) *ast.ExprStmt {
	upperRepositoryName := utils.UpperCamelCase(repositoryName)
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
							X:   ast.NewIdent("repository"),
							Sel: ast.NewIdent(fmt.Sprintf("New%sRepository", upperRepositoryName)),
						},
					},
				},
			},
		},
	}
}

func GenerateProvideRepositoryFn(repositoryName string, outPath string) ([]*ast.ImportSpec, *ast.FuncDecl) {
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
				Value: fmt.Sprintf("\"%s/repository\"", outPath),
			},
		},
	}
	provideRepositoryHandler := GenerateProvideRepositoryHandler(repositoryName)
	fn := &ast.FuncDecl{
		Name: ast.NewIdent("ProvideRepository"),
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
											provideRepositoryHandler,
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

func InjectRepository(repositoryName, output string) error {
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

		provideRepositoryFns := utils.Raise(utils.FindFunctions(file, "ProvideRepository", utils.ExactMatch))
		if len(provideRepositoryFns) == 0 {
			provideRepositoryImports, provideRepositoryFn := GenerateProvideRepositoryFn(repositoryName, output)
			file.Imports = utils.UniqueImports(append(file.Imports, provideRepositoryImports...))
			file.Decls = append(file.Decls, provideRepositoryFn)
		} else {
			repositoryContent := GenerateProvideRepositoryHandler(repositoryName)
			provideRepositoryFn := provideRepositoryFns[0].Node.(*ast.FuncDecl)
			for _, stmt := range provideRepositoryFn.Body.List {
				if returnStmt, ok := stmt.(*ast.ReturnStmt); ok {
					for _, result := range returnStmt.Results {
						if callExpr, ok := result.(*ast.CallExpr); ok {
							if inTryCatchVoid(callExpr) {
								if fn, ok := callExpr.Args[0].(*ast.FuncLit); ok {
									fn.Body.List = append(fn.Body.List, repositoryContent)
								}
							}
						}
					}
				}
			}
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
