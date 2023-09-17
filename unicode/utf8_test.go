package unicode_test

import (
	"testing"
	"unicode/utf8"
)

// utf8.AppendRune 将 rune 追加到 dst 的尾部，返回新的 slice
func TestUtf8AppendRune(t *testing.T) {
	p := []byte("hello")
	r := '中'
	p = utf8.AppendRune(p, r)
	t.Logf("p: %s\n", p)
}

// utf8.EncodeRune 将 rune 编码为 UTF-8，返回编码后的长度
func TestUtf8EncodeRune(t *testing.T) {
	p := make([]byte, 3)
	r := '中'
	n := utf8.EncodeRune(p, r)
	t.Logf("p: %s, n: %d\n", p, n)
}

// utf8.DecodeRune 将 UTF-8 编码解码为 rune，返回 rune 和编码长度
func TestUtf8DecodeRune(t *testing.T) {
	p := []byte("hello")
	r, n := utf8.DecodeRune(p)
	t.Logf("r: %c, n: %d\n", r, n)
}

// utf8.DecodeRuneInString 将 UTF-8 编码解码为 rune，返回 rune 和编码长度
func TestUtf8DecodeRuneInString(t *testing.T) {
	p := "hello"
	r, n := utf8.DecodeRuneInString(p)
	t.Logf("r: %c, n: %d\n", r, n)
}

// utf8.DecodeLastRune 将 UTF-8 编码解码为 rune，返回最后一个 rune 和编码长度
func TestUtf8DecodeLastRune(t *testing.T) {
	p := []byte("hello")
	r, n := utf8.DecodeLastRune(p)
	t.Logf("r: %c, n: %d\n", r, n)
}

// utf8.DecodeLastRuneInString 将 UTF-8 编码解码为 rune，返回最后一个 rune 和编码长度
func TestUtf8DecodeLastRuneInString(t *testing.T) {
	p := "hello"
	r, n := utf8.DecodeLastRuneInString(p)
	t.Logf("r: %c, n: %d\n", r, n)
}

// utf8.FullRune 判断 p 是否是一个完整的 UTF-8 编码
func TestUtf8FullRune(t *testing.T) {
	p := []byte("hello")
	b := utf8.FullRune(p)
	t.Logf("b: %v\n", b)
}

// utf8.FullRuneInString 判断 p 是否是一个完整的 UTF-8 编码
func TestUtf8FullRuneInString(t *testing.T) {
	p := "世界"
	b := utf8.FullRuneInString(p)
	t.Logf("b: %v\n", b)

	p1 := []byte(p)[:2]
	b1 := utf8.FullRuneInString(string(p1))
	t.Logf("b: %v\n", b1)
}

// utf8.RuneCount 将 p 解码为 rune，返回 rune 的个数
func TestUtf8RuneCount(t *testing.T) {
	p := []byte("hello, 世界")
	t.Logf("bytes = %d\n", len(p))
	n := utf8.RuneCount(p)
	t.Logf("n: %d\n", n)
}

// utf8.RuneCountInString 将 p 解码为 rune，返回 rune 的个数
func TestUtf8RuneCountInString(t *testing.T) {
	p := "hello, 世界"
	t.Logf("bytes = %d\n", len(p))
	n := utf8.RuneCountInString(p)
	t.Logf("n: %d\n", n)
}

// utf8.RuneLen 返回 rune 的编码长度
func TestUtf8RuneLen(t *testing.T) {
	r := '中'
	n := utf8.RuneLen(r)
	t.Logf("n: %d\n", n)
}

// utf8.RuneStart 判断 b 是否是一个 UTF-8 编码的开始
func TestUtf8RuneStart(t *testing.T) {
	p := []byte("a界")
	b := utf8.RuneStart(p[0])
	t.Logf("b: %v\n", b)
	b = utf8.RuneStart(p[1])
	t.Logf("b: %v\n", b)
	b = utf8.RuneStart(p[2])
	t.Logf("b: %v\n", b)
}

// utf8.Valid 判断 p 是否是一个合法的 UTF-8 编码
func TestUtf8Valid(t *testing.T) {
	valid := []byte("Hello, 世界")
	invalid := []byte{0xff, 0xfe, 0xfd}
	t.Logf("valid: %v\n", utf8.Valid(valid))
	t.Logf("invalid: %v\n", utf8.Valid(invalid))
}

// utf8.ValidRune 判断 r 是否是一个合法的 UTF-8 编码
func TestUtf8ValidRune(t *testing.T) {
	valid := 'a'
	invalid := rune(0xfffffff)
	t.Logf("valid: %v\n", utf8.ValidRune(valid))
	t.Logf("invalid: %v\n", utf8.ValidRune(invalid))
}

// utf8.ValidString 判断 s 是否是一个合法的 UTF-8 编码
func TestUtf8ValidString(t *testing.T) {
	valid := "Hello, 世界"
	invalid := string([]byte{0xff, 0xfe, 0xfd})
	t.Logf("valid: %v\n", utf8.ValidString(valid))
	t.Logf("invalid: %v\n", utf8.ValidString(invalid))
}
