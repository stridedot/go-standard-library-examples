package os_test

import (
	"context"
	"io"
	"os/exec"
	"testing"
	"time"
)

//

// exec.LookPath 用来查询可执行二进制文件的路径
func TestExecLookPath(t *testing.T) {
	path, err := exec.LookPath("echo")
	if err != nil {
		t.Fatalf("Look path failed, err = %v", err)
	}

	t.Logf("path = %v", path)
}

// 初始化一个Cmd指针，Path和Args按参数初始化，其他字段执行默认初始化
func TestExecCommand(t *testing.T) {
	arg := []string{"hello", "world"}
	cmd := exec.Command("echo", arg...)
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("cmd output failed, err = %v", err)
	}

	t.Logf("output = %s", output)
}

// 返回带有 context 上下文的 cmd 指针
func TestExecCommandContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Microsecond)
	defer cancel()

	// 将会在 100 微秒后报错，5 秒延迟会打断
	cmd := exec.CommandContext(ctx, "sleep", "5")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("error, err = %v", err)
	}
}

// CombinedOutput 运行该命令并返回其组合的标准输出和标准错误
func TestExecCombinedOutput(t *testing.T) {
	cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Set Combined output failed, err = %v", err)
	}
	t.Logf("%s\n", stdoutStderr)
}

// 返回当前配置的命令运行环境的副本
func TestExecEnviron(t *testing.T) {
	cmd := exec.Command("pwd")
	t.Logf("environ = %v", cmd.Environ())

	cmd.Dir = ".."
	cmd.Env = append(cmd.Environ(), "POSIXLY_CORRECT=1")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("Output failed, err = %v", out)
	}
	t.Logf("out = %s", out)
}

// 运行命令并返回其标准输出
func TestExecOutput(t *testing.T) {
	out, err := exec.Command("date").Output()
	if err != nil {
		t.Fatalf("output failed, err = %v", err)
	}

	t.Logf("out = %s", out)
}

// Run() 方法相当于 Start() + Wait()
func TestExecRun(t *testing.T) {
	cmd := exec.Command("sleep", "1")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Run failed, err = %v", err)
	}
}

// 用于连接到命令启动时错误标准输出的管道，命令结束时，管道会自动关闭
func TestExecStderrPipe(t *testing.T) {
	cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		t.Fatalf("StderrPipe failed, err = %v", err)
	}
	t.Logf("stderr = %#v", stderr)

	if err = cmd.Start(); err != nil {
		t.Fatalf("Start command, err = %v", err)
	}

	slurp, err := io.ReadAll(stderr)
	if err != nil {
		t.Fatalf("ReadAll failed, err = %v", err)
	}

	if err = cmd.Wait(); err != nil {
		t.Fatal(err)
	}

	t.Logf("Read all %s\n", slurp)
}

// 用于连接到命令启动时标准输入的管道
func TestExecStdinPipe(t *testing.T) {
	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatalf("Call StdinPipe failed, err = %v", err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Combined output failed, err = %v", err)
	}
	t.Logf("out = %s", out)
}

// 连接到命令启动时标准输出的管道，命令结束时，管道会自动关闭
func TestExecStdoutPipe(t *testing.T) {
	cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatalf("Call StdoutPipe failed, err = %v", err)
	}

	if err = cmd.Start(); err != nil {
		t.Fatalf("Start failed, err = %v", err)
	}

	b, err := io.ReadAll(stdout)
	if err != nil {
		t.Fatalf("ReadAll failed, err = %v", err)
	}
	t.Logf("stdout = %s", b)

	if err = cmd.Wait(); err != nil {
		t.Fatalf("Wait failed, err = %v", err)
	}
}
