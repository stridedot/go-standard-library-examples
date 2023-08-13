package errors_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrorNew(t *testing.T) {
	if errors.New("abc") == errors.New("abc") {
		t.Log(`New("abc") == New("abc")`)
	} else {
		t.Log(`New("abc") != New("abc")`)
	}
	if errors.New("abc") != errors.New("xyz") {
		t.Log(`New("abc") != New("xyz")`)
	}

	err := errors.New("abc")
	if err == err {
		t.Log(`err == err`)
	}

	if err.Error() != "abc" {
		t.Errorf(`New("abc").Error() = %q, want %q`, err.Error(), "abc")
	}
}

// 测试被包装过的error是否包含指定错误
func TestErrorIs(t *testing.T) {
	var BaseErr = errors.New("base error")

	err1 := fmt.Errorf("wrap base: %w", BaseErr)
	err2 := fmt.Errorf("wrap err1: %w", err1)

	t.Logf("err1 == BaseErr: %v", err1 == BaseErr)

	if errors.Is(err2, BaseErr) {
		t.Log("err2 is BaseErr")
		return
	}

	t.Log("err2 is not BaseErr")
}

type TypicalErr struct {
	e string
}

func (te TypicalErr) Error() string {
	return te.e
}

// 测试被包装过的error是否为指定类型
// 具体说明：提取指定类型的错误，判断包装的 error 链中，
// 某一个 error 的类型是否与 target 相同，并提取第一个符合目标类型的错误的值，
// 将其赋值给 target
func TestErrorAs(t *testing.T) {
	err := TypicalErr{"typical error"}
	err1 := fmt.Errorf("wrap err: %w", err)
	err2 := fmt.Errorf("wrap err1: %w", err1)

	var e TypicalErr
	t.Logf("err == e: %v", err == e)

	if errors.As(err2, &e) {
		t.Log("TypicalErr is on the chain of err2")
		return
	}

	t.Log("TypicalErr is not on the chain of err2")
}

// 测试集合多个错误
func TestErrorJoin(t *testing.T) {
	err1 := errors.New("err1")
	err2 := errors.New("err2")
	err := errors.Join(err1, err2)
	t.Logf("err: %v", err)
	t.Logf("err is err1: %v", errors.Is(err, err1))
	t.Logf("err is err2: %v", errors.Is(err, err2))
}

// 暴露内部的Err，而配合使用的还有Is和As两个方法
func TestErrorUnwrap(t *testing.T) {
	err1 := errors.New("error1")
	err2 := fmt.Errorf("error2: [%w]", err1)
	err3 := fmt.Errorf("error3: [%w]", err2)
	t.Logf("err3: %v", err3)
	t.Logf("errors.Unwrap(err3): %v", errors.Unwrap(err3))
}
