package zip_test

import (
	"archive/zip"
	"bytes"
	"log"
	"os"
	"testing"
)

// TestReadCloser 测试打开一个 zip 文件
func TestReadCloser(t *testing.T) {
	rc, err := zip.OpenReader("testdata/readme.zip")
	if err != nil {
		t.Fatalf("zip.OpenReader err: %v\n", err)
	}
	defer rc.Close()

	for _, file := range rc.File {
		osRc, err := file.Open()
		if err != nil {
			t.Fatalf("file.Open err: %v\n", err)
		}

		buf := make([]byte, 4096)
		n, _ := osRc.Read(buf)
		t.Logf("osRc.Read: %s\n", buf[:n])
		osRc.Close()
	}
}

// TestReader 测试打开 zip 压缩包中的某个文件
func TestReader(t *testing.T) {
	file, _ := os.Open("testdata/test.zip")
	fileInfo, _ := file.Stat()
	reader, err := zip.NewReader(file, fileInfo.Size())

	if err == zip.ErrFormat {
		t.Logf("zip.NewReader err: %v\n", err)
	}

	_, err = reader.Open("test.txt")
	if err != nil {
		t.Logf("reader.Open err: %v\n", err)
	}
}

func TestReaderOpen(t *testing.T) {
	data := []byte{
		0x50, 0x4b, 0x03, 0x04, 0x14, 0x00, 0x08, 0x00,
		0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x2e, 0x2e,
		0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x74, 0x78,
		0x74, 0x0a, 0xc9, 0xc8, 0x2c, 0x56, 0xc8, 0x2c,
		0x56, 0x48, 0x54, 0x28, 0x49, 0x2d, 0x2e, 0x51,
		0x28, 0x49, 0xad, 0x28, 0x51, 0x48, 0xcb, 0xcc,
		0x49, 0xd5, 0xe3, 0x02, 0x04, 0x00, 0x00, 0xff,
		0xff, 0x50, 0x4b, 0x07, 0x08, 0xc0, 0xd7, 0xed,
		0xc3, 0x20, 0x00, 0x00, 0x00, 0x1a, 0x00, 0x00,
		0x00, 0x50, 0x4b, 0x01, 0x02, 0x14, 0x00, 0x14,
		0x00, 0x08, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00,
		0x00, 0xc0, 0xd7, 0xed, 0xc3, 0x20, 0x00, 0x00,
		0x00, 0x1a, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2e,
		0x2e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x74,
		0x78, 0x74, 0x50, 0x4b, 0x05, 0x06, 0x00, 0x00,
		0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x39, 0x00,
		0x00, 0x00, 0x59, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	r, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	_, err = r.Open("test.txt")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	if len(r.File) != 1 {
		t.Fatalf("No entries in the file list")
	}
	if r.File[0].Name != "../test.txt" {
		t.Errorf("Unexpected entry name: %s", r.File[0].Name)
	}
	if _, err := r.File[0].Open(); err != nil {
		t.Errorf("Error opening file: %v", err)
	}
}

// TestWriterCreate 创建一个 zip 压缩包
func TestWriterCreate(t *testing.T) {
	outFile, err := os.Create("testdata/test0.zip")
	if err != nil {
		log.Fatalf("os.Create %v\n", err)
	}

	// Create a new zip archive.
	writer := zip.NewWriter(outFile)
	defer writer.Close()

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}

	for _, file := range files {
		f, err := writer.Create(file.Name)
		if err != nil {
			t.Fatalf("writer.Create err :%v\n", err)
		}

		_, err = f.Write([]byte(file.Body))
		if err != nil {
			t.Fatalf("f.Write err : %v\n", err)
		}
	}
}

// TestWriterCreateHeader 测试通过 FileHeader 创建一个 zip 压缩包
func TestWriterCreateHeader(t *testing.T) {
	zipFile, err := os.Create("testdata/test1.zip")
	if err != nil {
		log.Fatalf("os.Create %v\n", err)
	}

	// Create a new zip archive.
	writer := zip.NewWriter(zipFile)
	defer writer.Close()

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}

	for _, file := range files {
		f, err := os.Open("testdata/" + file.Name)
		if err != nil {
			t.Fatalf("os.Open err:%v\n", err)
		}
		fileInfo, err := f.Stat()
		if err != nil {
			t.Fatalf("f.Stat err:%v\n", err)
		}
		fileHeader, err := zip.FileInfoHeader(fileInfo)
		if err != nil {
			t.Fatalf("zip.FileInfoHeader err:%v\n", err)
		}
		w, err := writer.CreateHeader(fileHeader)
		if err != nil {
			t.Fatalf("writer.CreateHeader err:%v\n", err)
		}

		_, err = w.Write([]byte(file.Body))
		if err != nil {
			t.Fatalf("w.Write err : %v\n", err)
		}

		f.Close()
	}
}

// TestWriterCopy 测试 zip copy 到新的 zip 压缩包
func TestWriterCopy(t *testing.T) {
	zipFile, err := os.Create("testdata/test2.zip")
	if err != nil {
		log.Fatalf("os.Create %v\n", err)
	}

	// Create a new zip archive.
	writer := zip.NewWriter(zipFile)
	defer writer.Close()

	f, _ := os.Open("testdata/test.zip")
	fi, _ := f.Stat()
	reader, _ := zip.NewReader(f, fi.Size())

	for _, file := range reader.File {
		writer.Copy(file)
	}
}
