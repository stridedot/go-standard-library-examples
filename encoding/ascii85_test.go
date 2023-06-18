package encoding_test

import (
	"bytes"
	"encoding/ascii85"
	"io"
	"testing"
)

// 测试 ascii85 编码解码
func TestAscii85New(t *testing.T) {
	buf := new(bytes.Buffer)
	encoder := ascii85.NewEncoder(buf)

	n, err := encoder.Write([]byte("hello world"))
	_ = encoder.Close()
	if err != nil {
		t.Fatalf("encoder.Writer err: %v", err)
	}
	t.Logf("buf = %s, n = %d\n", buf.String(), n)

	decoder := ascii85.NewDecoder(buf)
	decoded, err := io.ReadAll(decoder)
	if err != nil {
		t.Fatalf("io.ReadAll err: %v", err)
	}
	t.Logf("decoded = %s\n", decoded)
}

func TestAscii85(t *testing.T) {
	src := []byte("hello world")
	dst := make([]byte, ascii85.MaxEncodedLen(len(src)))
	ascii85.Encode(dst, src)
	t.Logf("dst = %s\n", dst)

	p := make([]byte, 12)
	_, _, err := ascii85.Decode(p, dst, false)
	if err != nil {
		t.Fatalf("ascii85.Decode err: %v", err)
	}
	t.Logf("p = %s\n", p)
}
