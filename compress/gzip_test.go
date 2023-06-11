package compress_test

import (
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"testing"
	"time"
)

// gzip 解压缩文件，zlib 解压缩字符串

// TestGzipNewReader
func TestGzipNewReader(t *testing.T) {
	r, err := gzip.NewReader(bytes.NewReader(nil))
	if err != io.EOF {
		t.Fatalf("Want EOF, got err: %v", err)
	}

	defer r.Close()
}

// TestGzipRead 读取解压后的字节到 buf
func TestGzipRead(t *testing.T) {
	f, err := os.Open("testdata/e.txt.gz")
	if err != nil {
		t.Fatalf("Open file failed, err=%v", err)
	}
	defer f.Close()

	r, err := gzip.NewReader(f)
	if err != nil {
		t.Fatalf("NewReader failed err = %v\n", err)
	}
	defer r.Close()

	buf := make([]byte, 4096)
	n, err := r.Read(buf)
	if err != nil {
		t.Fatalf("Read err: %v", err)
	}

	t.Logf("buf = %s\n", buf[:n])
}

// TestGzipReset
// 使用Reset方法可以重置已经创建的Reader对象，
// 并将下一个读取目标设置为io.Reader类型的参数r，
// 从而达到复用的目的。在需要解压多个文件的时候避免内存分配，提高效率。
func TestGzipReset(t *testing.T) {
	f, err := os.Open("testdata/e.txt.gz")
	if err != nil {
		t.Fatalf("Open e.txt.gz failed, err=%v", err)
	}
	defer f.Close()

	r, err := gzip.NewReader(f)
	if err != nil {
		t.Fatalf("NewReader failed err = %v\n", err)
	}
	defer r.Close()

	f, err = os.Open("testdata/gzip.txt.gz")
	if err != nil {
		t.Fatalf("Open gzip.txt.gz failed, err=%v", err)
	}
	defer f.Close()

	err = r.Reset(f)
	t.Logf("err = %v\n", err)
}

// TestGzipNewWriter gzip.NewWriter 测试
func TestGzipNewWriter(t *testing.T) {
	gf, err := os.Create("testdata/demo.gz")
	if err != nil {
		t.Fatalf("Create file testdata/demo.gz failed, err = %v", err)
	}
	defer gf.Close()

	gw := gzip.NewWriter(gf)
	defer gw.Close()

	ff, err := os.Open("testdata/gzip.txt")
	if err != nil {
		t.Fatalf("Open testdata/gzip.txt failed, err=%v", err)
	}
	defer ff.Close()
	fi, err := ff.Stat()
	if err != nil {
		t.Fatalf("Stat testdata/gzip.txt failed, err=%v", err)
	}

	gw.Header.Name = fi.Name()
	b := make([]byte, fi.Size())
	n, _ := ff.Read(b)
	_, _ = gw.Write(b[:n])
	gw.Flush()
}

// TestGzipNewWriterWrite NewWriter 写入
func TestGzipNewWriterWrite(t *testing.T) {
	buf := new(bytes.Buffer)
	zw := gzip.NewWriter(buf)

	// Setting the Header fields is optional.
	zw.Name = "a-new-hope.txt"
	zw.Comment = "an epic space opera by George Lucas"
	zw.ModTime = time.Date(1977, time.May, 25, 0, 0, 0, 0, time.UTC)

	_, err := zw.Write([]byte("A long time ago in a galaxy far, far away..."))
	if err != nil {
		t.Fatalf("Write failed, err = %v", err)
	}

	// 按照 zw.Close() 的说法，close 将为写入的数据写入到底层的 io.Writer，
	// 并编写 GZip footer，但是不会关闭底层的 io.Writer
	// 使用 defer zw.Close()，会缺少 gzip footer
	// 因此需要在 return 之前关闭
	if err = zw.Close(); err != nil {
		t.Fatalf("Close Writer failed, err = %v", err)
	}

	zr, err := gzip.NewReader(buf)
	if err != nil {
		t.Fatalf("NewReader failed, err = %v", err)
	}

	if _, err = io.Copy(os.Stdout, zr); err != nil {
		t.Fatalf("Copy failed, err = %v", err)
	}

	if err = zr.Close(); err != nil {
		t.Fatalf("Close Reader failed, err = %v", err)
	}
}

// TestGzipNewWriterReset 重置 gzip.NewWriter
func TestGzipNewWriterReset(t *testing.T) {
	b1 := new(bytes.Buffer)
	b2 := new(bytes.Buffer)
	buf := []byte("hello world")

	w := gzip.NewWriter(b1)
	w.Write(buf)
	w.Close()

	w.Reset(b2)
	w.Write(buf)
	w.Close()

	if b1.String() != b2.String() {
		t.Fatal("Want b1 = b2, got err")
	}
}
