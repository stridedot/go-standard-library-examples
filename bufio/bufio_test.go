package bufio_test

import (
	"bufio"
	"bytes"
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

// TestReaderDiscard 测试 Reader 丢弃的字节数
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
