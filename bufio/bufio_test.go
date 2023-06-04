package bufio_test

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
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

// TestScannerScan 测试扫描数据
func TestScannerScan(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader("gopher"))
	for scanner.Scan() {
		if len(scanner.Bytes()) != 6 {
			t.Fatalf("Scan wrong length")
		}
	}
}

// TestScannerWithoutBuffer 测试默认最大缓冲
// s.Err() return token too long
func TestScannerWithoutBuffer(t *testing.T) {
	text := strings.Repeat("x", 2*bufio.MaxScanTokenSize)
	s := bufio.NewScanner(strings.NewReader(text + "\n"))
	for s.Scan() {
		if text != s.Text() {
			t.Fatalf("Scan got incorrect token of length %d", len(text))
		}
	}
	if s.Err() != nil {
		t.Fatalf("after scan: %v", s.Err())
	}
}

// Scanner.Buffer 方法可以重置缓冲区的最大长度
// 默认最大值是 bufio.MaxScanTokenSize
func TestScannerBuffer(t *testing.T) {
	text := strings.Repeat("x", 2*bufio.MaxScanTokenSize)
	s := bufio.NewScanner(strings.NewReader(text + "\n"))
	s.Buffer(make([]byte, 100), 3*bufio.MaxScanTokenSize)
	for s.Scan() {
		if text != s.Text() {
			t.Fatalf("Scan got incorrect token of length %d", len(text))
		}
	}
	if s.Err() != nil {
		t.Fatalf("after scan: %v", s.Err())
	}
}

// TestScannerSplit 测试按单词分割
func TestScannerSplit(t *testing.T) {
	var input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		t.Logf("scan = %v\n", scanner.Text())
	}
	if scanner.Err() != nil {
		t.Fatalf("reading error: %v\n", scanner.Err())
	}
}

// TestCustomSplit 测试自定义分割函数
func TestCustomSplit(t *testing.T) {
	input := "1234 5678 12345 67901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			// 转为64位，会输出 67901234567890
			// 转为32位，不会输出 67901234567890
			_, err = strconv.ParseInt(string(token), 10, 64)
		}
		return
	}
	scanner.Split(split)
	for scanner.Scan() {
		t.Logf(scanner.Text())
	}

	if scanner.Err() != nil {
		t.Fatalf("Invalid input: %v", scanner.Err())
	}
}

func TestNewWriter(t *testing.T) {
	w := bufio.NewWriter(os.Stdout)
	w.WriteString("hello,")
	w.WriteString("world!")
	w.Flush()
}

// TestNewWriterSize 测试返回指定长度的 writer
func TestNewWriterSize(t *testing.T) {
	w := bufio.NewWriterSize(bytes.NewBuffer([]byte("你好世界")), 6)
	if w.Size() != 6 {
		t.Fatalf("Want 6, got %d", w.Size())
	}
}

// TestWriterAvailable 测试未使用的 buffer 长度
func TestWriterAvailable(t *testing.T) {
	w := bufio.NewWriter(os.Stdout)
	if w.Available() != 4096 {
		t.Fatalf("wrong available lenth %d", w.Available())
	}

	_, _ = w.WriteString("你好世界")

	if w.Available() != 4084 {
		t.Fatalf("Got error length %d", w.Available())
	}

	_ = w.Flush()
}

// TestWriterFlush 测试输出到终端
// 如果需要输出，NewWriter 必不能是 nil
func TestWriterFlush(t *testing.T) {
	str := strings.Repeat("x", 1<<10)
	bw := bufio.NewWriter(io.Discard)
	bw.WriteString(str)
	bw.Flush()
}

// TestWriterAvailableBuffer 测试 writer 向外暴露一个空的 []byte
func TestWriterAvailableBuffer(t *testing.T) {
	w := bufio.NewWriter(os.Stdout)
	for _, i := range []int64{1, 2, 3, 4} {
		b := w.AvailableBuffer()
		b = strconv.AppendInt(b, i, 10)
		b = append(b, ' ')
		_, _ = w.Write(b)
	}
	_ = w.Flush()
}

// TestWriterBuffered 测试已写入的字节数
func TestWriterBuffered(t *testing.T) {
	w := bufio.NewWriter(os.Stdout)
	n, _ := w.WriteString("你好")
	if w.Buffered() != n {
		t.Fatalf("Want 6, got %d", w.Buffered())
	}
	if w.Available() != 4090 {
		t.Fatalf("Want 4090, got %d", w.Available())
	}
}

// TestWriterReadFrom 测试从 reader 中读取数据到 writer
func TestWriterReadFrom(t *testing.T) {
	w := bufio.NewWriter(os.Stdout)
	w.ReadFrom(strings.NewReader("abcdef"))
	w.Flush()
}

// TestWriterReset
// 重置放弃任何未刷新的缓冲数据，清除任何错误，并重置 w 以将其输出写入 buf2。
// 注意：重置时是将 buf1 重置为 buf2
func TestWriterReset(t *testing.T) {
	var buf1, buf2 strings.Builder

	w := bufio.NewWriter(&buf1)
	_, _ = w.WriteString("foo")

	w.Reset(&buf2) // and not flushed
	_, _ = w.WriteString("bar")
	_ = w.Flush()

	if buf1.String() != "" {
		t.Errorf("buf1 = %q; want empty", buf1.String())
	}
	if buf2.String() != "bar" {
		t.Errorf("buf2 = %q; want bar", buf2.String())
	}
}

// TestWriterWrite 测试写入
func TestWriterWrite(t *testing.T) {
	buf := []byte("hello world")
	writer := bufio.NewWriter(os.Stdout)
	n, err := writer.Write(buf)
	if n < len(buf) {
		t.Fatalf("Want %d, got %d", len(buf), n)
	}
	if err != nil {
		t.Fatalf("Got err: %v", err)
	}
}

// TestWriterWriteByte 写入单个字节
func TestWriterWriteByte(t *testing.T) {
	writer := bufio.NewWriter(os.Stdout)
	err := writer.WriteByte('a')
	if err != nil {
		t.Fatalf("Got err: %v\n", err)
	}
	_ = writer.Flush()
}

// TestWriterWriteRune 写入单个 utf-8 编码的字符
func TestWriterWriteRune(t *testing.T) {
	writer := bufio.NewWriter(os.Stdout)
	_, err := writer.WriteRune('你')
	if err != nil {
		t.Fatalf("Got err: %v\n", err)
	}
	_ = writer.Flush()
}

// TestWriterWriteString 测试写入字符串
func TestWriterWriteString(t *testing.T) {
	writer := bufio.NewWriter(os.Stdout)
	_, err := writer.WriteString("你好世界")
	if err != nil {
		t.Fatalf("Got err: %v\n", err)
	}
	_ = writer.Flush()
}
