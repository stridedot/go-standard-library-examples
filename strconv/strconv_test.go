package strconv_test

import (
	"strconv"
	"testing"
)

// strconv.AppendBool 将布尔值 b 转为字符串"true"或"false"，然后将其追加到 dst 尾部
func TestStrconvAppendBool(t *testing.T) {
	t.Logf("b = %s", strconv.AppendBool([]byte("abc"), true))
}

// strconv.AppendFloat 将浮点数 f 转为字符串，然后将其追加到 dst 尾部
func TestStrconvAppendFloat(t *testing.T) {
	b32 := []byte("float32:")
	b32 = strconv.AppendFloat(b32, 3.1415926535, 'E', -1, 32)
	t.Logf("b32 = %s", b32)

	b64 := []byte("float64:")
	b64 = strconv.AppendFloat(b64, 3.1415926535, 'g', -1, 64)
	t.Logf("b64 = %s", b64)
}

// strconv.AppendInt 将整数 i 转为字符串，然后将其追加到 dst 尾部
func TestStrconvAppendInt(t *testing.T) {
	t.Logf("i = %s", strconv.AppendInt([]byte("abc"), 123, 10))
}

// strconv.AppendQuote 将字符串 s 转为带双引号的字符串，然后将其追加到 dst 尾部
func TestStrconvAppendQuote(t *testing.T) {
	t.Logf("s = %s", strconv.AppendQuote([]byte("abc"), `"def"`))
}

// strconv.AppendQuoteRune 将字符 r 转为带单引号的字符串，然后将其追加到 dst 尾部
func TestStrconvAppendQuoteRune(t *testing.T) {
	t.Logf("r = %s", strconv.AppendQuoteRune([]byte("abc"), '☺'))
}

// strconv.AppendQuoteRuneToASCII 将字符 r 转为带单引号的 ASCII 字符串，然后将其追加到 dst 尾部
func TestStrconvAppendQuoteRuneToASCII(t *testing.T) {
	t.Logf("r = %s", strconv.AppendQuoteRuneToASCII([]byte("rune (ascii):"), '☺'))
}

// strconv.AppendQuoteToASCII 将字符串 s 转为带双引号的 ASCII 字符串，然后将其追加到 dst 尾部
func TestStrconvAppendQuoteToASCII(t *testing.T) {
	t.Logf("s = %s", strconv.AppendQuoteToASCII([]byte("quote (ascii):"), `"Fran & Freddie's Diner"`))
}

// strconv.AppendUint 将无符号整数 i 转为字符串，然后将其追加到 dst 尾部
func TestStrconvAppendUint(t *testing.T) {
	t.Logf("i = %s", strconv.AppendUint([]byte("abc"), 123, 10))
}

// strconv.Atoi 将字符串转为 int 类型
func TestStrconvAtoi(t *testing.T) {
	i, err := strconv.Atoi("123")
	if err != nil {
		t.Fatalf("Atoi failed, err = %v", err)
	}
	t.Logf("i = %d, type = %T", i, i)
}

// strconv.CanBackquote 判断字符串 s 是否可以不经转义地表示为一个单行的反引号字符串
func TestStrconvCanBackquote(t *testing.T) {
	t.Logf("s = %v", strconv.CanBackquote("Fran & Freddie's Diner ☺"))
}

// strconv.FormatBool 将布尔值 b 转为字符串"true"或"false"
func TestStrconvFormatBool(t *testing.T) {
	s := strconv.FormatBool(true)
	t.Logf("b = %s, type = %T", s, s)
}
