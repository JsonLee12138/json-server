package generate

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"testing"
)

func TestGenerateProvideControllerHandler_Success(t *testing.T) {
	controllerHandler := GenerateProvideControllerHandler("test")
	fset := token.NewFileSet()
	var buf bytes.Buffer

	if err := format.Node(&buf, fset, controllerHandler); err != nil {
		fmt.Println(err)
	}
	const want = "utils.RaiseVoid(container.Provide(controller.NewTestController))"
	if buf.String() != want {
		t.Errorf("GenerateProvideControllerHandler() failed = %q, want %q", buf.String(), want)
		return
	}
	t.Logf("GenerateProvideControllerHandler() success = %q", buf.String())
}

func TestGenerateRouterHandler_Success(t *testing.T) {
	routerHandler := GenerateRouterHandler("test")
	fset := token.NewFileSet()
	var buf bytes.Buffer

	if err := format.Node(&buf, fset, routerHandler); err != nil {
		fmt.Println(err)
	}
	const want = "{\n\tgroup.Get(\"test\", testController.HelloWorld)\n}"
	if buf.String() != want {
		t.Errorf("GenerateRouterHandler() = %q, want %q", buf.String(), want)
		return
	}
	t.Logf("GenerateRouterHandler() = %q, want %q", buf.String(), want)
}

func TestGenerateProvideControllerFn_Success(t *testing.T) {
	imports, fn := GenerateProvideControllerFn("test", "test")

	fset := token.NewFileSet()

	fmt.Println(imports)
	fmt.Println(fn)
	// for _, spec := range imports {
	// 	var buf bytes.Buffer
	// 	err := format.Node(&buf, fset, spec)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	t.Logf("imports: %q", buf.String())
	// }

	var buf bytes.Buffer

	if err := format.Node(&buf, fset, fn); err != nil {
		t.Errorf("GenerateProvideControllerFn() failed = %q", buf.String())
	}

	t.Logf("GenerateProvideControllerFn() success = %q", buf.String())
}

func TestIsRaiseVoidWithInvoke_Success(t *testing.T) {
	_, fn := GenerateProvideControllerFn("test", "./test")

	for _, stmt := range fn.Body.List {
		if returnStmt, ok := stmt.(*ast.ReturnStmt); ok {
			for _, result := range returnStmt.Results {
				if callExpr, ok := result.(*ast.CallExpr); ok {
					if inTryCatchVoid(callExpr) {
						t.Logf("isRaiseVoidWithInvoke() success = %q", callExpr.Args[0])
						if f, ok := callExpr.Args[0].(*ast.FuncLit); ok {
							fmt.Println(1, f.Body.List)
							fn.Body.List = append(fn.Body.List, GenerateProvideControllerHandler("aaa"))
						}
						return
					}
				}
			}
		}
	}
	t.Errorf("isRaiseVoidWithInvoke() failed")
}
