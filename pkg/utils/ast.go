package utils

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"regexp"
	"strings"
)

type MatchType int

const (
	ExactMatch    MatchType = iota // 精确匹配
	RegexMatch                     // 正则匹配
	ContainsMatch                  // 包含匹配
)

type SearchResult struct {
	Name string
	Node ast.Node
}

func matchString(pattern string, target string, matchType MatchType) bool {
	switch matchType {
	case ExactMatch:
		return pattern == target
	case ContainsMatch:
		return strings.Contains(target, pattern)
	case RegexMatch:
		matched, _ := regexp.MatchString(pattern, target)
		return matched
	default:
		return false
	}
}

// parseFile 解析 Go 文件为 AST
func ParseFile(filePath string) (node *ast.File, fset *token.FileSet, err error) {
	err = TryCatchVoid(func() {
		src := Raise(os.ReadFile(filePath))
		fset = token.NewFileSet()
		node = Raise(parser.ParseFile(fset, filePath, src, parser.ParseComments))
	}, DefaultErrorHandler)
	return node, fset, err
}

// FindFunctions 查找函数定义
func FindFunctions(file *ast.File, pattern string, matchType MatchType) ([]SearchResult, error) {
	var results []SearchResult
	ast.Inspect(file, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok {
			if matchString(pattern, fn.Name.Name, matchType) {
				results = append(results, SearchResult{Name: fn.Name.Name, Node: n})
			}
		}
		return true
	})
	return results, nil
}

// FindStructs 查找结构体定义
func FindStructs(filePath string, pattern string, matchType MatchType) ([]SearchResult, error) {
	node, _, err := ParseFile(filePath)
	if err != nil {
		return nil, err
	}

	var results []SearchResult
	ast.Inspect(node, func(n ast.Node) bool {
		if ts, ok := n.(*ast.TypeSpec); ok {
			if _, isStruct := ts.Type.(*ast.StructType); isStruct {
				if matchString(pattern, ts.Name.Name, matchType) {
					results = append(results, SearchResult{Name: ts.Name.Name, Node: n})
				}
			}
		}
		return true
	})
	return results, nil
}

// FindInterfaces 查找接口定义
func FindInterfaces(filePath string, pattern string, matchType MatchType) ([]SearchResult, error) {
	node, _, err := ParseFile(filePath)
	if err != nil {
		return nil, err
	}

	var results []SearchResult
	ast.Inspect(node, func(n ast.Node) bool {
		if ts, ok := n.(*ast.TypeSpec); ok {
			if _, isInterface := ts.Type.(*ast.InterfaceType); isInterface {
				if matchString(pattern, ts.Name.Name, matchType) {
					results = append(results, SearchResult{Name: ts.Name.Name, Node: n})
				}
			}
		}
		return true
	})
	return results, nil
}

func CreateFunc(file *ast.File, name string, params *ast.FieldList, content *ast.BlockStmt) {
	funcName := ast.NewIdent(name)
	funcDecl := &ast.FuncDecl{
		Name: funcName,
		Type: &ast.FuncType{
			Params: params,
		},
		Body: content,
	}
	file.Decls = append(file.Decls, funcDecl)
}

func WriteToFile(file *ast.File, filePath string) error {
	fset := token.NewFileSet()
	return WriteToFileByFset(fset, file, filePath)
}

func WriteToFileByFset(fset *token.FileSet, file *ast.File, filePath string) error {
	return TryCatchVoid(func() {
		var buf bytes.Buffer
		if err := printer.Fprint(&buf, fset, file); err != nil {
			Throw(err)
		}
		RaiseVoid(os.WriteFile(filePath, buf.Bytes(), 0644))
	}, DefaultErrorHandler)
}

func CreateFunction(file *ast.File, name string, content *ast.BlockStmt) {
	funcName := ast.NewIdent(name)
	params := &ast.FieldList{
		List: []*ast.Field{
			{
				Names: []*ast.Ident{
					ast.NewIdent("container"),
				},
				Type: &ast.StarExpr{
					X: &ast.SelectorExpr{
						X:   ast.NewIdent("dig"),
						Sel: &ast.Ident{Name: "Container"},
					},
				},
			},
		},
	}

	funcDecl := &ast.FuncDecl{
		Name: funcName,
		Type: &ast.FuncType{
			Params: params,
		},
		Body: content,
	}
	// 将函数声明添加到文件节点中
	file.Decls = append(file.Decls, funcDecl)

	// 打印生成的代码
	fset := token.NewFileSet()
	err := printer.Fprint(os.Stdout, fset, file)
	if err != nil {
		fmt.Printf("Error printing code: %v\n", err)
	}
}

func FindTryCatchVoid(content *ast.SelectorExpr) {
	if content.Sel.Name == "TryCatchVoid" {
		//if ident, ok := content.X.(*ast.Ident); ok && ident.Name == "core" {
		//
		//}
		//fmt.Println(content.)
	}
}

func UniqueImports(imports []*ast.ImportSpec) []*ast.ImportSpec {
	uniqueImports := make(map[string]*ast.ImportSpec)
	for _, importSpec := range imports {
		uniqueImports[importSpec.Path.Value] = importSpec
	}
	uniqueImportList := make([]*ast.ImportSpec, 0, len(uniqueImports))
	for _, importSpec := range uniqueImports {
		uniqueImportList = append(uniqueImportList, importSpec)
	}
	return uniqueImportList
}
