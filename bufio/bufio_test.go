package bufio_test

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"
)

// TestReadWriter 测试创建一个 rw 对象
func TestReadWriter(t *testing.T) {
	r := bytes.NewBuffer([]byte("123456780"))
	w := bytes.NewBuffer(nil)
	rb := bufio.NewReader(r)
	wb := bufio.NewWriter(w)
	rw := bufio.NewReadWriter(rb, wb)

	b := make([]byte, 128)
	n, _ := rw.Read(b)
	t.Logf("rw.Read %s\n", b[:n])

	n, _ = rw.Write([]byte("this is a new string"))
	rw.Flush()
	t.Logf("w.String %s\n", w.String())
}

// TestReaderSize
func TestReaderSize(t *testing.T) {
	buf := bytes.NewBuffer([]byte("abcdefg"))
	reader := bufio.NewReader(buf)

	b := make([]byte, 128)
	n, err := reader.Read(b)
	if err != nil {
		t.Fatalf("reader.Read err: %v\n", err)
	}
	t.Logf("reader.Read: %s\n", string(b[:n]))

	buf = bytes.NewBuffer([]byte("123456789"))
	reader = bufio.NewReaderSize(buf, 3)
	n, err = reader.Read(b)
	if err != nil {
		t.Fatalf("reader Size Read err: %v\n", err)
	}
	t.Logf("reader Size Read: %s\n", string(b[:n]))
}

// TestReaderBuffered 测试 Reader 可以从当前缓冲区读取的字节数（剩余未读取的）
func TestReaderBuffered(t *testing.T) {
	buf := bytes.NewBuffer([]byte("123456789"))
	reader := bufio.NewReader(buf)

	b := make([]byte, 3)
	n, err := reader.Read(b)
	if err != nil {
		t.Fatalf("reader Read err: %v\n", err)
	}
	t.Logf("reader read: %s\n", b[:n])

	if m := reader.Buffered(); m != 6 {
		t.Fatalf("Got %d bytes buffered in bufio.Reader; want 6 bytes", m)
	}
}

// bufio.Reader.Discard 测试 Reader 丢弃的字节数
// bufio.Reader.Peek 测试返回指定 n 个字节，Peek 读取后不会影响后续的读取
func TestReaderDiscardAndPeek(t *testing.T) {
	buf := bytes.NewBuffer([]byte("1234567890"))
	reader := bufio.NewReader(buf)

	discarded, err := reader.Discard(3)
	if err != nil {
		t.Fatalf("reader.Discard err: %v\n", err)
	}
	t.Logf("reader.Discard: %d\n", discarded)

	if s, err := reader.Peek(1); string(s) != "4" {
		t.Fatalf("Want %q got %q, err=%v", "4", string(s), err)
	}

	if s, err := reader.Peek(1); string(s) != "4" {
		t.Fatalf("Want %q got %q, err=%v", "4", string(s), err)
	}

	p := make([]byte, 6)
	if n, err := reader.Read(p); n != 6 {
		t.Fatalf("Want %d bytes in reader.Reader, got %d bytes, err=%v", 6, n, err)
	}
}

// TestReaderRead 测试从 Reader 中读取数据
func TestReaderRead(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("foo foo"))
	buf := make([]byte, 3)
	_, _ = r.Read(buf)
	if string(buf) != "foo" {
		t.Errorf("Want foo, got %q", string(buf))
	}
}

// TestReaderReadByte 测试读取并返回单个字节
func TestReaderReadByte(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("hello world"))
	b, _ := r.ReadByte()
	if string(b) != "h" {
		t.Fatalf("Want h, got %q", string(b))
	}

	b, _ = r.ReadByte()
	if string(b) != "e" {
		t.Fatalf("Want e, got %q", string(b))
	}
}

// TestReaderReadBytes 测试读取直到第一次出现输入中的定界符
// 并返回一个包含数据（包括分隔符）的切片
func TestReaderReadBytes(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("bar,foo"))
	b, _ := r.ReadBytes(',')
	if string(b) != "bar," {
		t.Fatalf("Want bar, got %q\n", string(b))
	}

	buf := make([]byte, 3)
	r.Read(buf)
	if string(buf) != "foo" {
		t.Fatalf("Want foo, got %q", string(buf))
	}
}

// TestReaderReadLine 测试读取一行数据
func TestReaderReadLine(t *testing.T) {
	var testInput = []byte("012\n345\n678\n9ab\ncde\nfgh\nijk\nlmn\nopq\nrst\nuvw\nxy")
	reader := bufio.NewReader(bytes.NewReader(testInput))

	for {
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		t.Logf("Line = %q\n", string(line))
		t.Logf("isPrefix = %v\n", isPrefix)
		t.Logf("err = %v\n", err)
	}
}

// TestReaderReadRune 测试读取一个 utf-8 编码的Unicode 字符
func TestReaderReadRune(t *testing.T) {
	reader := bufio.NewReader(bytes.NewBuffer([]byte("你好，世界")))
	c, size, err := reader.ReadRune()
	if string(c) != "你" || size != 3 {
		t.Fatalf("Want c='你' got %q, Want size=3 got %d", string(c), size)
	}

	c, size, err = reader.ReadRune()
	if string(c) != "好" || size != 3 {
		t.Fatalf("Want c='好' got %q, Want size=3 got %d", string(c), size)
	}

	c, size, err = reader.ReadRune()
	if string(c) != "，" || size != 3 {
		t.Fatalf("Want c='，' got %q, Want size=3 got %d", string(c), size)
	}

	c, size, err = reader.ReadRune()
	if string(c) != "世" || size != 3 {
		t.Fatalf("Want c='世' got %q, Want size=3 got %d", string(c), size)
	}

	c, size, err = reader.ReadRune()
	if string(c) != "界" || size != 3 {
		t.Fatalf("Want c='界' got %q, Want size=3 got %d", string(c), size)
	}

	c, size, err = reader.ReadRune()
	if size != 0 || err != io.EOF {
		t.Fatalf("Want size=0 got %d, Want err=EOL got %q", size, err)
	}
}

// TestReaderReadSlice 测试读取切片
func TestReaderReadSlice(t *testing.T) {
	reader := bufio.NewReader(bytes.NewBuffer([]byte("日本語日本語日本語")))
	_, err := reader.ReadSlice(',')
	if err != io.EOF {
		t.Fatalf("Want EOL, got %v", err)
	}
}

// TestReaderReadString 测试读取数据输入中的定界符第一次出现
func TestReaderReadString(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("你好$世界$"))
	s, _ := reader.ReadString('$')
	if s != "你好$" {
		t.Fatalf("Want `你好$`, got %q\n", s)
	}
}

// TestReaderReset 测试重置
func TestReaderReset(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("foo foo"))
	b := make([]byte, 3)
	_, _ = reader.Read(b)
	if string(b) != "foo" {
		t.Fatalf("Watn foo, got %q\n", string(b))
	}

	t.Logf("size = %d\n", reader.Size())

	reader.Reset(strings.NewReader("bar bar"))
	all, _ := io.ReadAll(reader)
	if string(all) != "bar bar" {
		t.Fatalf("Want `bar bar`, got %q\n", string(all))
	}
}

// TestReaderUnreadByte 测试反读最后一个读取的字节
func TestReaderUnreadByte(t *testing.T) {
	reader := bufio.NewReader(bytes.NewBuffer([]byte("abc123")))
	b, _ := reader.ReadByte()
	if string(b) != "a" {
		t.Fatalf("Want `a`, got %q\n", string(b))
	}

	err := reader.UnreadByte()
	if err != nil {
		t.Fatalf("reader.UnreadByte err = %v\n", err)
	}

	buf := make([]byte, 6)
	reader.Read(buf)
	if string(buf) != "abc123" {
		t.Fatalf("Want `abc123`, got %q\n", string(buf))
	}
}

// TestReaderUnreadRune 测试反读最后一个读取的 rune
func TestReaderUnreadRune(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("你好，世界"))
	r, _, _ := reader.ReadRune()
	if r != '你' {
		t.Fatalf("Want `你`, got %q\n", r)
	}

	buf := make([]byte, 128)
	reader.UnreadRune()
	n, _ := reader.Read(buf)
	if string(buf[:n]) != "你好，世界" {
		t.Fatalf("Want `你好，世界`, got %q\n", string(buf))
	}
}

// TestReaderWriteTo 测试将 reader 数据写入到 writer 中
func TestReaderWriteTo(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("你好，世界"))
	w := new(bytes.Buffer)
	n, _ := reader.WriteTo(w)
	if n != 15 {
		t.Fatalf("Want n=15, got %d\n", n)
	}
	if n != int64(w.Len()) {
		t.Fatalf("Want n=%d, got n!=%d\n", w.Len(), w.Len())
	}
}
