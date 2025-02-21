package core

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/JsonLee12138/jsonix/embed"
	"github.com/JsonLee12138/jsonix/pkg/utils"
)

const (
	entitiesDestPackageName = "auto_migrate"
)

func AggregateEntities(root, dest string) error {
	var entities []string
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fileName := info.Name()
		if info.IsDir() && (fileName == "entity" || fileName == "entities") {
			entities = append(entities, CopyEntities(path, dest)...)
		}
		return nil
	})
	if len(entities) == 0 {
		return nil
	}
	entitiesStr := strings.Join(entities, ",")
	tmpl := string(utils.Raise(embed.TemplatesPath.ReadFile("templates/auto_migrate.tmpl")))
	return GenerateFileExistsHandler(fmt.Sprintf("%s/entry.go", dest), tmpl, map[string]string{
		"Entities": entitiesStr,
	}, true)
}

func genAutoMigrate(entitiesStr string) *ast.FuncDecl {
	params := &ast.FieldList{
		List: []*ast.Field{
			{
				Names: []*ast.Ident{
					ast.NewIdent("db"),
				},
				Type: &ast.StarExpr{
					X: &ast.SelectorExpr{
						X:   ast.NewIdent("gorm"),
						Sel: ast.NewIdent("DB"),
					},
				},
			},
		},
	}
	content := &ast.BlockStmt{
		List: []ast.Stmt{
			&ast.ReturnStmt{
				Results: []ast.Expr{
					&ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent("db"),
							Sel: ast.NewIdent("AutoMigrate"),
						},
						Args: []ast.Expr{
							&ast.BasicLit{
								Kind:  token.STRING,
								Value: entitiesStr,
							},
						},
					},
				},
			},
		},
	}

	return &ast.FuncDecl{
		Name: ast.NewIdent("AutoMigrate"),
		Type: &ast.FuncType{
			Params: params,
			Results: &ast.FieldList{
				List: []*ast.Field{
					{Type: ast.NewIdent("error")},
				},
			},
		},
		Body: content,
	}
}

func ReadEntities(path string) (entities []string, node *ast.File, err error) {
	fset := token.NewFileSet()
	n, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, nil, err
	}
	for _, decl := range n.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			if genDecl.Doc != nil {
				for _, comment := range genDecl.Doc.List {
					if strings.Contains(comment.Text, "@AutoMigrate") {
						entities = append(entities, fmt.Sprintf("&%s{}", typeSpec.Name.Name))
						break
					}
				}
			}
		}
	}
	return entities, n, nil
}

func CopyEntities(path, destPath string) (entities []string) {
	exists, isDir, err := utils.Exists(destPath)
	if err != nil {
		return nil
	}
	if !exists || !isDir {
		utils.CreateDir(destPath)
	}
	filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fileName := info.Name()
		if strings.HasSuffix(fileName, ".go") && !strings.Contains(fileName, "_easyjson") {
			entitiesItems, node, err := ReadEntities(filePath)
			entities = append(entities, entitiesItems...)
			if err == nil && len(entities) > 0 {
				newNode := &ast.File{
					Name:  ast.NewIdent(entitiesDestPackageName),
					Decls: node.Decls,
				}
				err = utils.WriteToFile(newNode, fmt.Sprintf("%s/%s", destPath, info.Name()))
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	return entities
}
