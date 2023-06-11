package compress_test

import (
	"bytes"
	"compress/flate"
	"io"
	"testing"
)

// TestFlateNewReader 创建一个 ReaderCloser
// flate.NewReader 参数列表：
// 参数：r deflate压缩文件的文件标识符
// 返回值：解压后的ReadCloser数据
// 功能：从r读取deflate压缩数据，返回一个解压过的io.ReadCloser，使用后需要调用关闭该io.ReadCloser
func TestFlateNewReader(t *testing.T) {
	b := new(bytes.Buffer)
	w, err := flate.NewWriter(b, flate.BestSpeed)
	if err != nil {
		t.Fatalf("NewWriter err: %v", err)
	}
	w.Write([]byte("你好世界"))
	w.Flush()
	w.Close()

	r := flate.NewReader(b)
	data, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("ReadAll err: %v", err)
	}
	t.Logf("data = %s\n", data)
}

// TestFlateNewReaderDict
// flate.NewReaderDict 参数列表：
// r deflate压缩的数据
// dict 解压数据时预设的字典，和NewWriteDict函数里得dict相同
// 返回值：解压后ReadCloser数据
// 功能说明：从r读取deflate压缩数据，使用预设得dict字典压缩数据，返回一个压缩过得io.ReadCloser，使用后需要调用者关闭该io.ReadCloser。主要用来读取NewWriteDict压缩的数据
func TestFlateNewReaderDict(t *testing.T) {
	const (
		dict = "hello world"
		text = "hello again world"
	)
	var b bytes.Buffer
	w, err := flate.NewWriter(&b, flate.BestSpeed)
	if err != nil {
		t.Fatalf("NewWriter: %v", err)
	}
	w.Write([]byte(dict))
	w.Flush()
	w.Close()

	r := flate.NewReaderDict(&b, []byte(dict))
	data, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "hello again world" {
		t.Fatalf("read returned %q want %q", string(data), text)
	}
}

// TestFlateNewWriter 参数列表：
// 1）w 表示输出数据的Write
// 2）level 表示压缩级别
// 返回列表：
//
//	1）*Write 基于压缩级别新生成的压缩数据的Writer
//	2）error 表示该函数的错误信息
//
// 功能说明：
//
//	该函数返回一个压缩级别为level的新的压缩用的Writer，
//	压缩级别的范围时1（BestSpeed）to 9（BestCompression）。
//	压缩效果越好的意味着压缩速度越慢。0（NoCompression）表示不做任何压缩；
//	仅仅只需要添加必要的deflate信息，
//	-1（DefaultCompression）表示用默认的压缩级别。
//	如果压缩级别在-1~9的范围内，error返回nil，否则将返回非nil的错误信息。
func TestFlateNewWriter(t *testing.T) {
	b := new(bytes.Buffer)
	w, err := flate.NewWriter(b, flate.BestCompression)
	if err != nil {
		t.Fatalf("NewWriter err: %v", err)
	}
	w.Write([]byte("hi Emma"))
	w.Flush()
	defer w.Close()

	t.Logf("b.String = %s\n", b.String())
}

// TestFlateNewWriterDict 参数列表：
//
//	1）w 代表输出数据的Writer
//	2）level 代表压缩级别
//	3）dict 代表压缩预设字典
//
// 返回列表：
//
//	1）*Writer 基于压缩级别和预设字典新生成的压缩数据的Writer
//	2）error 该函数的错误信息
//
// 功能说明：
// 该函数和NewWriter差不多，只不过使用了预设字典进行初始化Writer。使用该Writer压缩的数据只能被使用相同字典初始化的Reader解压。可以实现基于密码的解压缩。
func TestFlateNewWriterDict(t *testing.T) {
	b := new(bytes.Buffer)
	w, err := flate.NewWriterDict(b, flate.BestCompression, []byte("key"))
	if err != nil {
		t.Fatalf("NewWriterDict err: %v", err)
	}
	defer w.Close()
	w.Write([]byte("hi Emma"))
	w.Flush()

	t.Logf("b.String = %s\n", b.String())
}
