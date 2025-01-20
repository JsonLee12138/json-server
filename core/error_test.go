package core

import (
	"errors"
	"testing"
)

// ✅ 测试 TryCatch 正常执行
func TestTryCatch_Success(t *testing.T) {
	res, err := TryCatch(func() int {
		return 42
	}, DefaultErrorHandler)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if res != 42 {
		t.Errorf("Expected result to be 42, but got %d", res)
	}
}

// ✅ 测试 TryCatch 触发 panic
func TestTryCatch_Panic(t *testing.T) {
	_, err := TryCatch(func() int {
		panic("simulated panic")
	}, func(err error) error {
		t.Logf("Caught panic: %v", err)
		return err
	})

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

// ✅ 测试 TryCatchVoid 正常执行
func TestTryCatchVoid_Success(t *testing.T) {
	err := TryCatchVoid(func() {
		println("Running task...")
	}, DefaultErrorHandler)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

// ✅ 测试 TryCatchVoid 触发 panic
func TestTryCatchVoid_Panic(t *testing.T) {
	err := TryCatchVoid(func() {
		panic("panic in TryCatchVoid")
	}, func(err error) error {
		t.Logf("Caught panic: %v", err)
		return err
	})

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

// ✅ 测试 Raise 正常执行
func TestRaise_Success(t *testing.T) {
	res := Raise(100, nil)
	if res != 100 {
		t.Errorf("Expected result to be 100, but got %d", res)
	}
}

// ✅ 测试 Raise 触发 panic
func TestRaise_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		}
	}()
	Raise(0, errors.New("test error"))
}

// ✅ 测试 RaiseVoid 正常执行
func TestRaiseVoid_Success(t *testing.T) {
	RaiseVoid(nil) // 无 error，不应触发 panic
}

// ✅ 测试 RaiseVoid 触发 panic
func TestRaiseVoid_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		}
	}()
	RaiseVoid(errors.New("test error in RaiseVoid"))
}

// ✅ 测试 handlePanic 处理 panic
func TestHandlePanic(t *testing.T) {
	err := handlePanic("panic message", func(err error) error {
		t.Logf("Caught panic: %v", err)
		return err
	})

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

// ✅ 测试 DefaultErrorHandler
func TestDefaultErrorHandler(t *testing.T) {
	DefaultErrorHandler(errors.New("test error message")) // 仅用于观察输出，无需断言
}

// ✅ 测试 Throw
func TestPanicErrorHandler(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		}
	}()
	Throw(errors.New("test panic"))
}
