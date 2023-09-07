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

// strconv.FormatComplex 将复数 complex128 转为字符串
func TestStrconvFormatComplex(t *testing.T) {
	s := strconv.FormatComplex(2.718281828+3.1415926535i, 'g', 10, 128)
	t.Logf("c = %s, type = %T", s, s)
}

// strconv.FormatFloat 将浮点数 f 转为字符串
func TestStrconvFormatFloat(t *testing.T) {
	s := strconv.FormatFloat(3.1415926535, 'b', -1, 64)
	t.Logf("f = %s, type = %T", s, s)

	s = strconv.FormatFloat(3.1415926535, 'e', -1, 64)
	t.Logf("f = %s, type = %T", s, s)

	s = strconv.FormatFloat(3.1415926535, 'E', -1, 64)
	t.Logf("f = %s, type = %T", s, s)

	s = strconv.FormatFloat(3.1415926535, 'f', -1, 64)
	t.Logf("f = %s, type = %T", s, s)

	s = strconv.FormatFloat(3.1415926000, 'g', -1, 64)
	t.Logf("f = %s, type = %T", s, s)

	s = strconv.FormatFloat(3.1415926535, 'G', -1, 64)
	t.Logf("f = %s, type = %T", s, s)
}

// strconv.FormatInt 将整数 i 转为字符串，当进制大于10时，大于10的值将使用小写a-z表示
func TestStrconvFormatInt(t *testing.T) {
	s := strconv.FormatInt(-42, 10)
	t.Logf("s = %s, type = %T", s, s)
}

// strconv.FormatUint 将无符号整数 i 转为字符串，当进制大于10时，大于10的值将使用小写a-z表示
func TestStrconvFormatUint(t *testing.T) {
	s := strconv.FormatUint(42, 10)
	t.Logf("s = %s, type = %T", s, s)
}

// strconv.IsGraphic 判断字符 c 是否为可显示字符，空格除外
func TestStrconvIsGraphic(t *testing.T) {
	t.Logf("c = %v", strconv.IsGraphic('☘'))
}

// strconv.Itoa 将整数 i 转为字符串
func TestStrconvItoa(t *testing.T) {
	s := strconv.Itoa(123)
	t.Logf("type is %T, s = %v", s, s)
}

// strconv.ParseBool 将字符串转为布尔值
func TestStrconvParseBool(t *testing.T) {
	b, err := strconv.ParseBool("t")
	if err != nil {
		t.Fatalf("ParseBool failed, err = %v", err)
	}
	t.Logf("b = %v, type = %T", b, b)
}

// strconv.ParseComplex 将字符串转为复数
func TestStrconvParseComplex(t *testing.T) {
	c, err := strconv.ParseComplex("2.718281828+3.1415926535i", 128)
	if err != nil {
		t.Fatalf("ParseComplex failed, err = %v", err)
	}
	t.Logf("c = %v, type = %T", c, c)
}

// strconv.ParseFloat 将字符串转为浮点数
func TestStrconvParseFloat(t *testing.T) {
	f, err := strconv.ParseFloat("3.1415926535", 64)
	if err != nil {
		t.Fatalf("ParseFloat failed, err = %v", err)
	}
	t.Logf("f = %v, type = %T", f, f)
}

// strconv.ParseInt 将字符串转为整数
func TestStrconvParseInt(t *testing.T) {
	i, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		t.Fatalf("ParseInt failed, err = %v", err)
	}
	t.Logf("i = %v, type = %T", i, i)
}

// strconv.ParseUint 将字符串转为无符号整数
func TestStrconvParseUint(t *testing.T) {
	i, err := strconv.ParseUint("123", 10, 64)
	if err != nil {
		t.Fatalf("ParseUint failed, err = %v", err)
	}
	t.Logf("i = %v, type = %T", i, i)
}

// strconv.Quote 将字符串 s 转为带双引号的字符串
func TestStrconvQuote(t *testing.T) {
	s := strconv.Quote(`"Fran & Freddie's Diner	☺"`)
	t.Logf("s = %v, type = %T", s, s)
}

// strconv.QuoteRune 将字符 r 转为带单引号的字符串
func TestStrconvQuoteRune(t *testing.T) {
	s := strconv.QuoteRune('☺')
	t.Logf("s = %v, type = %T", s, s)
}

// strconv.QuoteRuneToASCII 将字符 r 转为带单引号的 ASCII 字符串
func TestStrconvQuoteRuneToASCII(t *testing.T) {
	s := strconv.QuoteRuneToASCII('☺')
	t.Logf("s = %v, type = %T", s, s)
}

// strconv.QuoteRuneToGraphic 将字符 r 转为带单引号的可显示字符串
func TestStrconvQuoteRuneToGraphic(t *testing.T) {
	s := strconv.QuoteRuneToGraphic('☺')
	t.Logf("s = %v, type = %T", s, s)

	s = strconv.QuoteRuneToGraphic('\u263a')
	t.Logf("s = %v, type = %T", s, s)

	s = strconv.QuoteRuneToGraphic('\u000a')
	t.Logf("s = %v, type = %T", s, s)

	s = strconv.QuoteRuneToGraphic('	') // tab character
	t.Logf("s = %v, type = %T", s, s)
}

// strconv.QuoteToASCII 将字符串 s 转为带双引号的 ASCII 字符串
func TestStrconvQuoteToASCII(t *testing.T) {
	s := strconv.QuoteToASCII(`"Fran & Freddie's Diner	☺"`)
	t.Logf("s = %v, type = %T", s, s)
}

// strconv.QuoteToGraphic 将字符串 s 转为带双引号的可显示字符串
func TestStrconvQuoteToGraphic(t *testing.T) {
	s := strconv.QuoteToGraphic("☺")
	t.Logf("s = %v, type = %T", s, s)

	// This string literal contains a tab character.
	s = strconv.QuoteToGraphic("This is a \u263a	\u000a")
	t.Logf("s = %v, type = %T", s, s)

	s = strconv.QuoteToGraphic(`" This is a ☺ \n "`)
	t.Logf("s = %v, type = %T", s, s)
}

// strconv.QuotedPrefix QuotedPrefix函数返回 s 的前缀处的带引号字符串(如 Unquote函数理解的那样)。
// 如果 s 不以有效的带引号字符串开头，则 QuotedPrefix函数返回一个错误。
func TestStrconvQuotedPrefix(t *testing.T) {
	prefix, err := strconv.QuotedPrefix("Hello, world! How are you?")
	t.Logf("prefix = %v, ok = %v", prefix, err)

	prefix, err = strconv.QuotedPrefix("\"Hello, world!\" How are you?")
	t.Logf("prefix = %v, ok = %v", prefix, err)

	prefix, err = strconv.QuotedPrefix("'☺'Hello, world! How are you?")
	t.Logf("prefix = %v, ok = %v", prefix, err)

	prefix, err = strconv.QuotedPrefix("`Hello, world!` How are you?")
	t.Logf("prefix = %v, ok = %v", prefix, err)
}

// strconv.Unquote 将字符串 s 转为不带引号的字符串
func TestStrconvUnquote(t *testing.T) {
	s, err := strconv.Unquote("You can't unquote a string without quotes")
	t.Logf("%q, %v\n", s, err)
	s, err = strconv.Unquote("\"The string must be either double-quoted\"")
	t.Logf("%q, %v\n", s, err)
	s, err = strconv.Unquote("`or backquoted.`")
	t.Logf("%q, %v\n", s, err)
	s, err = strconv.Unquote("'\u263a'") // single character only allowed in single quotes
	t.Logf("%q, %v\n", s, err)
	s, err = strconv.Unquote("'\u2639\u2639'")
	t.Logf("%q, %v\n", s, err)
}

// strconv.UnquoteChar 将字符串 s 转为字符
func TestStrconvUnquoteChar(t *testing.T) {
	sr := `\"大\\家\\好！\"`
	var c rune
	var mb bool
	var err error
	for ; len(sr) > 0; c, mb, sr, err = strconv.UnquoteChar(sr, '"') {
		t.Log(c, mb, sr, err)
	}
}
