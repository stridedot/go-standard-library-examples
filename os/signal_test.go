package os_test

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

// 忽略提供的信号
func TestSignalIgnore(t *testing.T) {
	c := make(chan syscall.Signal, 1)
	signal.Ignore(syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	s := <-c
	t.Log("Got signal:", s)
	time.Sleep(5 * time.Second)
}

// 判断当前信号是否是忽略信号
func TestSignalIgnored(t *testing.T) {
	signal.Ignore(syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	t.Log("Ignored:", signal.Ignored(syscall.SIGHUP))
	time.Sleep(5 * time.Second)
}

// 将信号转到 channel c
func TestSignalNotify(t *testing.T) {
	c := make(chan os.Signal, 1)

	// 如果是 signal.Notify(c) 表示所有的信号将会被传递到 channel c
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	// 阻塞状态，直到 notify 中的 c 有信号进入
	s := <-c
	t.Logf("s = %v", s)
}

// 将信号转到 ctx
func TestSignalNotifyContent(t *testing.T) {
	// 监听中断信号（ctrl + c）
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)

	// 重置 os.Interrupt 的默认行为，类似 signal.Reset
	defer stop()

	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatalf("FindProcess failed, err = %v", err)
	}

	if err = p.Signal(os.Interrupt); err != nil {
		t.Fatalf("err = %v", err)
	}

	select {
	case <-time.After(time.Second):
		t.Log("missed signal")
	case <-ctx.Done():
		ctx.Err()
		// 立即停止接收信号
		stop()
	}
}

// 撤销之前对所有提供信号调用 Notify 的影响
// 放 reset 返回时，c 不会收到任何信号
func TestSignalReset(t *testing.T) {
	c := make(chan os.Signal, 1)
	signal.Notify(c)
	signal.Reset()
	select {
	case s := <-c:
		t.Fatalf("unexpected signal %v", s)
	case <-time.After(100 * time.Millisecond):
		t.Log("reset signal") // nothing to read - good
	}
}

// 撤销之前使用 c 调用 Notify 的所有效果
// 当 stop 返回时，c 不会接收到任何信号
func TestSignalStop(t *testing.T) {
	c := make(chan os.Signal, 1)
	signal.Notify(c)
	signal.Stop(c)
	select {
	case s := <-c:
		t.Fatalf("unexpected signal %v", s)
	case <-time.After(100 * time.Millisecond):
		t.Log("stop signal") // nothing to read - good
	}
}
