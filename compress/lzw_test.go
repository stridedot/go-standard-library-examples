package compress_test

import (
	"bytes"
	"compress/lzw"
	"io"
	"testing"
)

// TestNewReader
// NewReader创建一个新的io.ReadCloser。
// 从返回的io.ReadCloser中读取并解压r中的数据。
// 如果r没有同时实现io.ByteReader，解压器可能会从r中读取更多的数据，
// 当完成读取时，调用者有责任对ReadCloser调用Close。
// 用于字面编码的比特数，litWidth，必须在[2,8]的范围内，通常是8
func TestNewReader(t *testing.T) {
	in := make([]byte, 5406)
	in = append(in, 0x80, 0xff, 0x0f, 0x08)
	rc := lzw.NewReader(bytes.NewReader(in), lzw.LSB, 8)

	_, err := io.Copy(io.Discard, rc)
	if err != nil {
		t.Fatalf("Copy err: %v", err)
	}
}

// TestNewReaderRead 读取未压缩的字节
func TestNewReaderRead(t *testing.T) {
	b := new(bytes.Buffer)
	w := lzw.NewWriter(b, lzw.LSB, 8)
	_, err := w.Write([]byte("how to study Go"))
	if err != nil {
		t.Fatalf("Write failed, err = %v", err)
	}
	w.Close()

	r := lzw.NewReader(b, lzw.LSB, 8)
	defer r.Close()
	p := make([]byte, 20)
	n, err := r.Read(p)
	if err != nil {
		t.Fatalf("Read failed, err = %v", err)
	}
	t.Logf("p = %s\n", p[:n])
}

// TestNewWriter 创建一个 io.Writer
func TestNewWriter(t *testing.T) {
	buf := new(bytes.Buffer)
	w := lzw.NewWriter(buf, lzw.LSB, 8)
	_, err := w.Write([]byte("Hello Kitty"))
	if err != nil {
		t.Fatalf("Write failed, err = %v", err)
	}
	_ = w.Close()

	r := lzw.NewReader(buf, lzw.LSB, 8)
	p := make([]byte, 20)
	n, err := r.Read(p)
	if err != nil {
		t.Fatalf("Read failed, err = %v", err)
	}
	t.Logf("Read p = %s\n", p[:n])
	_ = r.Close()
}
