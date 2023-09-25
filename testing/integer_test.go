package testing_test

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Expected '%d', but got '%d'", expected, sum)
	}
}

func Add(m, n int) int {
	return m + n
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

const repeatCount = 5

func Repeat(str string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += str
	}
	return repeated
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a"))
	// Output: aaaaa
}
