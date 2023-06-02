package tar_test

import (
	"archive/tar"
	"bytes"
	"io"
	"os"
	"testing"
)

// TestHeaderFileInfoHeader 测试 tar 格式文件的头部信息
func TestHeaderFileInfoHeader(t *testing.T) {
	fi, err := os.Stat("testdata/small.txt")
	if err != nil {
		t.Fatalf("Stat: %v", err)
	}
	header, err := tar.FileInfoHeader(fi, "")
	if err != nil {
		t.Fatalf("FileInfoHeader: %v", err)
	}
	if g, e := header.Name, "small.txt"; g != e {
		t.Fatalf("Expected %q, got %q", e, g)
	}
	t.Logf("Size: %d\n", header.Size)
	t.Logf("Mode: %d\n", header.Mode)
	t.Logf("ModTime: %v\n", header.ModTime)

	hfi := header.FileInfo()
	t.Logf("Size: %v\n", hfi.Size())
	t.Logf("Sys: %v\n", hfi.Sys())
}

// TestReader 测试 Reader
func TestReader(t *testing.T) {
	file, err := os.Open("testdata/hdr-only.tar")
	if err != nil {
		t.Fatalf("Open: %v\n", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var headers []*tar.Header
	reader := tar.NewReader(file)

	for {
		header, err := reader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("reader.Next: %v\n", err)
		}

		name := header.Name
		switch header.Typeflag {
		case tar.TypeDir:
			break
		case tar.TypeReg:
			t.Logf("header.Name: %v\n", name)
			// 复制输出到控制台
			// io.Copy(os.Stdout, reader)

			// 读取内容
			b := make([]byte, 1024)
			n, _ := reader.Read(b[:])
			t.Logf("reader.Read: %s\n", b[:n])
		default:
			t.Logf("header.Typeflag: %c\n", header.Typeflag)
			headers = append(headers, header)
		}
	}
}

// TestWriter 测试 Writer
func TestWriter(t *testing.T) {
	var buf bytes.Buffer
	writer := tar.NewWriter(&buf)
	defer writer.Close()

	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	}

	for _, file := range files {
		// 先写头部信息，头部信息包含字节长度等限制
		hdr := &tar.Header{
			Name:     file.Name,
			Mode:     0600,
			Size:     int64(len(file.Body)),
			Typeflag: tar.TypeReg,
		}
		err := writer.WriteHeader(hdr)
		if err != nil {
			t.Fatalf("writer.WriteHeader: %v\n", err)
		}

		// 将文件写入到 tar
		_, err = writer.Write([]byte(file.Body))
		if err != nil {
			t.Fatalf("writer.Write: %v\n", err)
		}
	}

	reader := tar.NewReader(&buf)
	for {
		header, err := reader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("reader.Next: %v\n", err)
		}
		t.Logf("header.Name: %v\n", header.Name)

		var b [128]byte
		n, _ := reader.Read(b[:])
		t.Logf("reader.Read: %s\n", b[:n])
	}
}

// TestWriteToTar 测试打包到指定目录和文件
func TestWriteToTar(t *testing.T) {
	var buf bytes.Buffer
	writer := tar.NewWriter(&buf)
	defer writer.Close()

	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	}

	for _, file := range files {
		// 先写头部信息，头部信息包含字节长度等限制
		hdr := &tar.Header{
			Name:     file.Name,
			Mode:     0600,
			Size:     int64(len(file.Body)),
			Typeflag: tar.TypeReg,
		}
		err := writer.WriteHeader(hdr)
		if err != nil {
			t.Fatalf("writer.WriteHeader: %v\n", err)
		}

		// 将文件写入到 tar
		_, err = writer.Write([]byte(file.Body))
		if err != nil {
			t.Fatalf("writer.Write: %v\n", err)
		}
	}

	// 写入到指定文件
	tarFile := "testdata/test.tar"
	err := os.WriteFile(tarFile, buf.Bytes(), os.ModePerm)
	if err != nil {
		t.Fatalf("os.WriteFile err: %v\n", err)
	}
}
