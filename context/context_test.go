package context_test

import (
	"context"
	"testing"
	"time"
)

// 测试 context.Background
func TestBackground(t *testing.T) {
	c := context.Background()
	if c == nil {
		t.Fatal("Want content.Background, got nil")
	}
	select {
	case x := <-c.Done():
		t.Fatalf("%v <-c.Done() should block", x)
	default:
	}
}

// 测试 context.TODO
func TestTODO(t *testing.T) {
	c := context.TODO()
	if c == nil {
		t.Fatal("Want content.TODO, got nil")
	}
}

// 测试 context.WithCancel，取消父协程所有关联的子协程
// 被取消时调用 context.Done
func TestWithCancel(t *testing.T) {
	// gen 函数在协程中将n存入channel，调用 gen 方法输出n的时候
	// 需要取消 context
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				// 父 context 的 Done 通道关闭时，所有子协程的通道都将被关闭
				case <-ctx.Done():
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 循环防止协程过早结束
	for n := range gen(ctx) {
		t.Logf("n = %d\n", n)
		if n == 5 {
			break
		}
	}
}

// 测试 context.WithDeadline
// WithDeadline 运行到截止时间，返回的取消函数被调用，或者父 Context 的 Done 被关闭时，
// 超时时调用 context.Done
func TestWithDeadline(t *testing.T) {
	d := time.Now().Add(1 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		t.Log("over slept")
	case <-ctx.Done():
		t.Logf("err = %v\n", ctx.Err())
	}
}

// 测试 超时关闭 Done 通道
func TestWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Millisecond):
		t.Log("over slept")
	case <-ctx.Done():
		t.Log("err = ", ctx.Err())
	}
}

// 将值传递到 context，并在 context 中检索
func TestWithValue(t *testing.T) {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v == nil {
			t.Fatalf("Key %q not found", k)
		}
		t.Logf("key %q found", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	k = favContextKey("color")
	f(ctx, k)
}
