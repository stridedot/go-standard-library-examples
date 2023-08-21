package io_test

import (
	"bytes"
	"fmt"
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

// 测试 io.Pipe 拆分成了 PipeReader 和 PipeWriter 暴露出去。
// PipeReader.Read() 方法写入数据，PipeWriter.Write() 读取数据。由于其实现了接口 io.ReadCloser 和 io.WriteCloser 。所以可以使用更高层的工具函数来操作它们。
// 对于 PipeReader ，可以使用 bytes.Buffer.ReadFrom 或者 bufio.NewScanner 。
// 对于 PipeWriter 可以使用 fmt.Fprintf() 或者 bufio.NewWriter
func TestIOPipe(t *testing.T) {
	r, w := io.Pipe()
	defer r.Close()

	go func() {
		_, err := w.Write([]byte("hello world"))
		defer w.Close()
		if err != nil {
			t.Fatalf("w.Write error: %v", err)
		}
	}()

	b := make([]byte, 10)
	n, err := r.Read(b)
	if err != nil {
		t.Logf("r.Read error: %v", err)
	}
	t.Logf("b = %s", b[:n])
}

// 测试 io.ReadCloser
func TestIOReadCloser(t *testing.T) {
	r := strings.NewReader("hello world")
	rc := io.NopCloser(r)
	defer rc.Close()
}

// 测试 io.LimitedReader
func TestIOLimitedReader(t *testing.T) {
	r := bytes.NewReader([]byte("test io reader"))
	lr := io.LimitReader(r, 20)

	var total int
	b := make([]byte, 5)

	for {
		n, err := lr.Read(b)
		if err == io.EOF {
			t.Log("Read EOF ", total)
			break
		}
		if err != nil {
			t.Fatalf("lr.Read error: %v", err)
		}

		t.Logf("b = %s", b[:n])
		total += n
	}
}

// 测试 io.MultiReader.Read
func TestIOMultiReader(t *testing.T) {
	r1 := bytes.NewReader([]byte("test io r1"))
	r2 := bytes.NewReader([]byte("test io r2"))
	mr := io.MultiReader(r1, r2)

	var n, total int
	var err error
	b := make([]byte, 5)

	for {
		n, err = mr.Read(b)
		if err == io.EOF {
			t.Log("Read EOF ", total)
			break
		}
		if err != nil {
			t.Fatalf("mr.Read error: %v", err)
		}
		t.Logf("b = %s", b[:n])
		total += n
	}
}

// 测试 io.MultiReader.WriteTo
// 需要先将 io.MultiReader 转换成 io.WriterTo 接口
func TestIOMultiReaderWriteTo(t *testing.T) {
	r1 := bytes.NewReader([]byte("test io r1"))
	r2 := bytes.NewReader([]byte("test io r2"))
	mr := io.MultiReader(r1, r2)

	mr.(io.WriterTo).WriteTo(os.Stdout)
}

// 测试 io.TeeReader
func TestIOTeeReader(t *testing.T) {
	r := bytes.NewReader([]byte("test io reader"))
	tr := io.TeeReader(r, os.Stdout)

	b := make([]byte, 5)
	_, err := tr.Read(b)
	if err != nil {
		t.Fatalf("tr.Read error: %v", err)
	}

	p, err := io.ReadAll(tr)
	if err != nil {
		t.Fatalf("io.ReadAll error: %v", err)
	}
	t.Logf("b = %s", p)
}

func testIOTeeReaderRead(t *testing.T) {
	reader, _ := os.Open("test.txt")
	writer, _ := os.Create("dst.txt")
	teeReader := io.TeeReader(reader, writer)
	var n, total int
	var err error
	p := make([]byte, 20)
	for {
		n, err = teeReader.Read(p)
		total = total + n
		if err == nil {
			fmt.Println("Read and write value", string(p[0:n]))
			fmt.Println("Read and write count", total)
		}
		if err == io.EOF {
			fmt.Println("Read and write end total", total)
			break
		}
	}
}

// 测试 io.PipeReader.Read
func TestIOSectionReaderRead(t *testing.T) {
	r, _ := os.Open("test.txt")
	sr := io.NewSectionReader(r, 6, 20)

	t.Logf("sr.Size() = %d", sr.Size())

	var n, total int
	var err error
	b := make([]byte, 10)

	for {
		n, err = sr.Read(b)
		if err == io.EOF {
			t.Logf("Read EOF, total size = %d", total)
			break
		}
		if err != nil {
			t.Fatalf("sr.Read error: %v", err)
		}
		total += n

		t.Logf("b = %q", b[:n])
	}
}

// 测试 io.PipeReader.ReadAt
// ReadAt 从指定位置开始读取数据，而不是从当前位置开始读取数据
func TestIOSectionReaderReadAt(t *testing.T) {
	r, _ := os.Open("test.txt")
	sr := io.NewSectionReader(r, 6, 20)

	b := make([]byte, 10)
	n, err := sr.ReadAt(b, 5)
	if err != nil {
		t.Fatalf("sr.ReadAt error: %v", err)
	}
	t.Logf("b = %q", b[:n])
}

// 测试 io.PipeReader.Seek
// 参数：
//   - offset 偏移量
//   - whence 设定选项
//     0:读取起始点，
//     1:当前读取点，
//     2:结束点(不好用)，
//     其他：将抛出Seek: invalid whence异常
//
// 返回值：
//   - 当前读取点相对读取起始点的偏移量
func TestIOSectionReaderSeek(t *testing.T) {
	r, _ := os.Open("test.txt")
	sr := io.NewSectionReader(r, 6, 20)

	off, err := sr.Seek(0, io.SeekStart)
	if err != nil {
		t.Fatalf("sr.Seek error: %v", err)
	}
	t.Logf("off = %d", off)

	b1 := make([]byte, 10)
	n, err := sr.Read(b1)
	if err != nil {
		t.Fatalf("sr.Read error: %v", err)
	}
	t.Logf("b1 = %q", b1[:n])

	// 从当前位置向后移动 5 个字节
	off, err = sr.Seek(5, io.SeekCurrent)
	if err != nil {
		t.Fatalf("sr.Seek error: %v", err)
	}
	t.Logf("off = %d", off)

	b := make([]byte, 10)
	n, err = sr.Read(b)
	if err != nil {
		t.Fatalf("sr.Read error: %v", err)
	}
	t.Logf("b = %q", b[:n])
}
