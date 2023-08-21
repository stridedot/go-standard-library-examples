package math_test

import (
	"math"
	"testing"
)

// 测试 math.Abs，math.Abs 返回一个浮点数的绝对值
func TestMathAbs(t *testing.T) {
	x := 3.14
	y := -3.14

	if math.Abs(x) != 3.14 {
		t.Errorf("Abs(%f) should be 3.14", x)
	}

	if math.Abs(y) != 3.14 {
		t.Errorf("Abs(%f) should be 3.14", y)
	}
}

// 测试 math.Acos 反余弦
func TestMathAcos(t *testing.T) {
	t.Logf("math.Acos = %v", math.Acos(0.5))
	t.Logf("math.Acos = %v", math.Acos(5))
}

func TestMathAcosh(t *testing.T) {
	t.Logf("math.Acosh = %v", math.Acosh(1))
}
