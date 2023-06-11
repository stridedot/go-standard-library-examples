package compress_test

import (
	"bytes"
	"compress/zlib"
	"io"
	"os"
	"testing"
)

// gzip 解压缩文件，zlib 解压缩字符串

// TestZlibNewReader 创建 zlib reader
func TestZlibNewReader(t *testing.T) {
	buff := []byte{120, 156, 202, 72, 205, 201, 201, 215, 81, 40, 207,
		47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 33, 231, 4, 147}
	r, err := zlib.NewReader(bytes.NewReader(buff))
	if err != nil {
		t.Fatalf("NewReader err : = %v", err)
	}
	io.Copy(os.Stdout, r)
	r.Close()
}

// TestZlibNewWriter 创建一个 zlib writer
func TestZlibNewWriter(t *testing.T) {
	b := new(bytes.Buffer)
	w := zlib.NewWriter(b)
	w.Write([]byte("Hello World\n"))
	w.Close()
	t.Logf("b = %v\n", b.Bytes())
}
