package io_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// 测试 io.Copy
func TestIOCopy(t *testing.T) {
	reader := bytes.NewBuffer([]byte("hello world"))
	writer := bytes.NewBuffer(nil)
	_, err := io.Copy(writer, reader)
	if err != nil {
		t.Logf("io.Copy error: %v", err)
	}
}

// 测试 io.CopyBuffer，与 copy 相同，只是多了一个 buffer 参数提供缓冲区
func TestIOCopyBuffer(t *testing.T) {
	reader := bytes.NewReader([]byte("hello world"))
	writer := bytes.NewBuffer(nil)
	_, err := io.CopyBuffer(writer, reader, make([]byte, 3))
	if err != nil {
		t.Logf("io.CopyBuffer error: %v", err)
	}
}

// 测试 io.CopyN，从 reader 复制 n 个字节到 writer
func TestIOCopyN(t *testing.T) {
	reader := strings.NewReader("hello world")
	writer := new(strings.Builder)
	_, err := io.CopyN(writer, reader, 5)
	if err != nil {
		t.Logf("io.CopyN error: %v", err)
	}
	t.Logf("writer: %s", writer.String())
}

// 测试 io.Pipe
// io.Pipe 创建了一个同步内存管道，可用于将需要 io.Reader 的代码与需要 io.Writer 的代码连接起来。
// 调用时，io.Pipe() 返回一个 PipeReader 和一个 PipeWriter。它们是连接的（管道），
// 因此写入 PipeWriter 的所有内容都可以从 PipeReader 读取
func TestPipe(t *testing.T) {
	r, w := io.Pipe()

	go func() {
		w.Write([]byte("hello world"))
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		t.Logf("io.Copy error: %v", err)
	}
}

// 测试 io.ReadAll
func TestIOReadAll(t *testing.T) {
	reader := bytes.NewReader([]byte("hello world"))
	b, err := io.ReadAll(reader)
	if err != nil {
		t.Logf("io.ReadAll error: %v", err)
	}
	t.Logf("b: %s", b)
}

// 测试 io.ReadAtLeast, 从 reader 读取至少 min 个字节到 buf
// 注意，min 是最少读取（贪婪读取）的字节数，具体读取长度由 buf 的长度决定
// 如果 min > buf 的长度，会报错
func TestReadAtLeast(t *testing.T) {
	buf := make([]byte, 10)
	reader := bytes.NewReader(nil)
	n, err := io.ReadAtLeast(reader, buf, 10)
	if err != io.EOF {
		t.Logf("io.ReadAtLeast error: %v", err)
	}
	t.Logf("n: %d", n)

	buf1 := make([]byte, 5)
	reader = bytes.NewReader([]byte("hello world"))
	n, err = io.ReadAtLeast(reader, buf1, 4)
	if err != nil {
		t.Logf("io.ReadAtLeast error: %v", err)
	}
	t.Logf("n: %d, buf1:%s", n, buf1)

	buf2 := make([]byte, 5)
	reader = bytes.NewReader([]byte("hello world"))
	n, err = io.ReadAtLeast(reader, buf2, 10)
	if err != nil {
		t.Logf("io.ReadAtLeast error: %v", err)
	}
	t.Logf("n: %d, buf2:%s", n, buf2)
}

// 测试 io.ReadFull, 从 reader 读取至少 len(buf) 个字节到 buf
func TestIOReadFull(t *testing.T) {
	reader := bytes.NewReader(nil)
	buf := make([]byte, 5)
	n, err := io.ReadFull(reader, buf)
	if err != io.EOF {
		t.Logf("io.ReadFull error: %v", err)
	}
	t.Logf("n: %d, buf: %s", n, buf)

	buf1 := make([]byte, 20)
	reader = bytes.NewReader([]byte("hello world"))
	n, err = io.ReadFull(reader, buf1)
	if err != nil {
		t.Logf("io.ReadFull error: %v", err)
	}
	t.Logf("n: %d, buf1: %s", n, buf1)
}

// 测试 io.WriteString
func TestIOWriteString(t *testing.T) {
	file, err := os.Create("test.txt")
	if err != nil {
		t.Logf("os.Create error: %v", err)
	}
	defer file.Close()
	_, err = io.WriteString(file, "hello world")
	if err != nil {
		t.Logf("io.WriteString error: %v", err)
	}
}

// 测试 io.LimitReader
func TestIOLimitReader(t *testing.T) {
	reader1 := io.LimitReader(nil, 0)
	t.Logf("reader1: %v", reader1)

	file, err := os.Open("test.txt")
	if err != nil {
		t.Logf("os.Open error: %v", err)
	}
	reader2 := io.LimitReader(file, 5)

	buf := make([]byte, 10)
	n, _ := reader2.Read(buf)
	if n != 5 {
		t.Fatalf("n: %d", n)
	}
}

// 测试 io.OffsetReader
func TestIOOffsetWriter(t *testing.T) {
	file, err := os.Create("test1.txt")
	if err != nil {
		t.Logf("os.Create error: %v", err)
	}
	offWriter := io.NewOffsetWriter(file, 0)

	b := []byte("hello kitty")

	// 将 b 写入到 offWriter 中
	_, err = offWriter.Write(b)
	if err != nil {
		t.Fatalf("n: %v", err)
	}
}

// 测试 io.OffsetWriter writeAt
// WriteAt 将 b 的内容写入到 offWriter 中，写入的起始位置为 off
func TestOffsetWriterWriteAt(t *testing.T) {
	file, err := os.Create("test1.txt")
	if err != nil {
		t.Fatalf("os.Create error: %v", err)
	}
	offWriter := io.NewOffsetWriter(file, 3)

	offWriter.WriteAt([]byte("hello WORLD"), 2)
}
