package encoding_test

import (
	"bytes"
	"encoding/base64"
	"strings"
	"testing"
)

// 测试 base64 的编码和解码
func TestBase64Encoder(t *testing.T) {
	msg := []byte("你好，世界")
	encoded := base64.StdEncoding.EncodeToString(msg)
	decoder, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		t.Fatalf("err = %v", decoder)
	}
	t.Logf("decoder = %s\n", decoder)
}

// 测试 Encoding 的 encode 和 decode
// 不适合大型数据流
func TestEncoding(t *testing.T) {
	msg := []byte("hello, world")
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(msg)))
	base64.StdEncoding.Encode(dst, msg)

	b := make([]byte, base64.StdEncoding.DecodedLen(len(dst)))
	n, err := base64.StdEncoding.Decode(b, dst)
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	t.Logf("b = %s\n", b[:n])
}

// 测试 base64 NewEncoder, 适合大型数据流
func TestBase64NewEncoder(t *testing.T) {
	buf := new(bytes.Buffer)
	encoder := base64.NewEncoder(base64.StdEncoding, buf)
	_, err := encoder.Write([]byte("Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit..."))
	if err != nil {
		t.Fatalf("Write err: %v", err)
	}
	_ = encoder.Close()
	t.Logf("buf = %s\n", buf.String())
}

// 测试 base64 NewDecoder
func TestBase64NewDecoder(t *testing.T) {
	encoded := "TmVxdWUgcG9ycm8gcXVpc3F1YW0gZXN0IHF1aSBkb2xvcmVtIGlwc3VtIHF1aWEgZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyLCBhZGlwaXNjaSB2ZWxpdC4uLg=="
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(encoded))
	n, err := decoder.Read(buf)
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	t.Logf("buf = %s\n", buf[:n])
}
