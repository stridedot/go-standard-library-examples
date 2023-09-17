package testing_test

import (
	"math/rand"
	"testing"
	"time"
)

// 写法：
// go test testing/benchmark_test.go
// -bench 支持正则表达式，只有匹配到的测试用例才会执行，使用 . 则运行所有测试用例
// -benchtime 测试速度
//		1.增加时长，默认1s，例：-benchtime=5s;
//		2.增加运行次数，例：-benchtime=30x;
// -count 运行次数，默认1次，例：-count=5;
// -cpu 指定运行cpu核数，默认1核，例：-cpu=2,4,6,8;
// -benchmem 显示内存分配情况，例：-benchmem;

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

// go test testing/benchmark_test.go -bench=BenchmarkFib
func BenchmarkFib(b *testing.B) {
	// 模拟耗时准备任务
	time.Sleep(3 * time.Second)
	// 重置定时器
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fib(30)
	}
}

// 分配切片容量
func generateWithCap(n int) []int {
	rand.NewSource(time.Now().UnixNano())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

// 不分配容量
func generateWithZeroCap(n int) []int {
	rand.NewSource(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func BenchmarkWithCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateWithCap(10000000)
	}
}

func BenchmarkWithZeroCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateWithZeroCap(10000000)
	}
}

// 循环 1000w 次
func loop(n int) {
	for i := 0; i < n; i++ {
		//
	}
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		loop(10000000)
		b.StartTimer()
		generateWithCap(1000000)
	}
}
