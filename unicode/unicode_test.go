package unicode_test

import (
	"testing"
	"unicode"
)

// unicode.In 判断 rune 是否属于某个范围的成员
func TestUnicodeIn(t *testing.T) {
	// 判断 rune 是否属于某个范围的成员
	if !unicode.In('a', unicode.Latin) {
		t.Fatalf("expected %q to be in Latin", 'a')
	}
	if unicode.In('a', unicode.Armenian) {
		t.Fatalf("expected %q to not be in Armenian", 'a')
	}
}

// unicode.Is 判断 rune 是否属于某个类别
func TestUnicodeIs(t *testing.T) {
	// 判断 rune 是否属于某个类别
	if !unicode.Is(unicode.Letter, 'a') {
		t.Fatalf("expected %q to be a letter", 'a')
	}
	if unicode.Is(unicode.Number, 'a') {
		t.Fatalf("expected %q to not be a number", 'a')
	}
}

// unicode.IsControl 判断 rune 是否是控制字符
func TestUnicodeIsControl(t *testing.T) {
	// 判断 rune 是否是控制字符
	if !unicode.IsControl('\t') {
		t.Fatalf("expected %q to be a control character", '\t')
	}
	if unicode.IsControl('a') {
		t.Fatalf("expected %q to not be a control character", 'a')
	}
}

// unicode.IsDigit 判断 rune 是否是十进制数字
func TestUnicodeIsDigit(t *testing.T) {
	// 判断 rune 是否是十进制数字
	if !unicode.IsDigit('1') {
		t.Fatalf("expected %q to be a digit", '1')
	}
	if unicode.IsDigit('a') {
		t.Fatalf("expected %q to not be a digit", 'a')
	}
}

// unicode.IsGraphic 判断 rune 是否是图形字符
func TestUnicodeIsGraphic(t *testing.T) {
	// 判断 rune 是否是图形字符
	if !unicode.IsGraphic('a') {
		t.Fatalf("expected %q to be a graphic character", 'a')
	}
	if unicode.IsGraphic('\t') {
		t.Fatalf("expected %q to not be a graphic character", '\t')
	}
}

// unicode.IsLetter 判断 rune 是否是字母
func TestUnicodeIsLetter(t *testing.T) {
	// 判断 rune 是否是字母
	if !unicode.IsLetter('a') {
		t.Fatalf("expected %q to be a letter", 'a')
	}
	if unicode.IsLetter('1') {
		t.Fatalf("expected %q to not be a letter", '1')
	}
}

// unicode.IsMark 判断 rune 是否是标记字符
func TestUnicodeIsMark(t *testing.T) {
	// 判断 rune 是否是标记字符
	if !unicode.IsMark('Ὂ') {
		t.Fatalf("expected %q to be a mark character", 'Ὂ')
	}
	if unicode.IsMark('a') {
		t.Fatalf("expected %q to not be a mark character", 'a')
	}
}

// unicode.IsNumber 判断 rune 是否是数字
func TestUnicodeIsNumber(t *testing.T) {
	// 判断 rune 是否是数字
	if !unicode.IsNumber('1') {
		t.Fatalf("expected %q to be a number", '1')
	}
	if !unicode.IsNumber('Ⅷ') {
		t.Fatalf("expected %q to not be a number", 'a')
	}
}

// unicode.IsOneOf 判断 rune 是否属于某个类别
func TestUnicodeIsOneOf(t *testing.T) {
	var letterDigit = []*unicode.RangeTable{
		unicode.Letter,
		unicode.Digit,
		{R16: []unicode.Range16{{'_', '_', 1}}},
	}

	// 判断 rune 是否属于某个类别
	if !unicode.IsOneOf(letterDigit, 'a') {
		t.Fatalf("expected %q to be a letter", 'a')
	}
	if !unicode.IsOneOf(letterDigit, '1') {
		t.Fatalf("expected %q to not be a number", '1')
	}
}

// unicode.IsPrint 判断 rune 是否是可打印字符
func TestUnicodeIsPrint(t *testing.T) {
	// 判断 rune 是否是可打印字符
	if !unicode.IsPrint('a') {
		t.Fatalf("expected %q to be a printable character", 'a')
	}
	if unicode.IsPrint('\t') {
		t.Fatalf("expected %q to not be a printable character", '\t')
	}
}

// unicode.IsPunct 判断 rune 是否是标点符号
func TestUnicodeIsPunct(t *testing.T) {
	// 判断 rune 是否是标点符号
	if !unicode.IsPunct('.') {
		t.Fatalf("expected %q to be a punctuation character", '.')
	}
	if unicode.IsPunct('a') {
		t.Fatalf("expected %q to not be a punctuation character", 'a')
	}
}

// unicode.IsSpace 判断 rune 是否是空白字符
func TestUnicodeIsSpace(t *testing.T) {
	// 判断 rune 是否是空白字符
	if !unicode.IsSpace(' ') {
		t.Fatalf("expected %q to be a space character", ' ')
	}
	if unicode.IsSpace('a') {
		t.Fatalf("expected %q to not be a space character", 'a')
	}
}

// unicode.IsSymbol 判断 rune 是否是符号
func TestUnicodeIsSymbol(t *testing.T) {
	// 判断 rune 是否是符号
	if !unicode.IsSymbol('©') {
		t.Fatalf("expected %q to be a symbol", '©')
	}
	if unicode.IsSymbol('a') {
		t.Fatalf("expected %q to not be a symbol", 'a')
	}
}

// unicode.IsTitle 判断 rune 是否是标题大小写字母
func TestUnicodeIsTitle(t *testing.T) {
	if !unicode.IsTitle('ǅ') {
		t.Fatalf("expected %q to be a title case letter", 'ǅ')
	}
	if unicode.IsTitle('a') {
		t.Fatalf("expected %q to not be a title case letter", 'a')
	}
	if unicode.IsTitle('1') {
		t.Fatalf("expected %q to not be a title case letter", '1')
	}
}

// unicode.IsUpper 判断 rune 是否是大写字母
func TestUnicodeIsUpper(t *testing.T) {
	// 判断 rune 是否是大写字母
	if !unicode.IsUpper('A') {
		t.Fatalf("expected %q to be an upper case letter", 'A')
	}
	if unicode.IsUpper('a') {
		t.Fatalf("expected %q to not be an upper case letter", 'a')
	}
	if unicode.IsUpper('1') {
		t.Fatalf("expected %q to not be an upper case letter", '1')
	}
}

// unicode.SimpleFold 返回与 rune 等效的最简单的大小写折叠形式
func TestUnicodeSimpleFold(t *testing.T) {
	// 返回与 rune 等效的最简单的大小写折叠形式
	if unicode.SimpleFold('A') != 'a' {
		t.Fatalf("expected %q to be the simple fold of %q", 'a', 'A')
	}
	if unicode.SimpleFold('a') != 'A' {
		t.Fatalf("expected %q to be the simple fold of %q", 'A', 'a')
	}
	if unicode.SimpleFold('1') != '1' {
		t.Fatalf("expected %q to be the simple fold of %q", '1', '1')
	}
}

// unicode.To 函数返回与 rune 等效的大小写转换形式
func TestUnicodeTo(t *testing.T) {
	// 返回与 rune 等效的大小写转换形式
	if unicode.To(unicode.UpperCase, 'a') != 'A' {
		t.Fatalf("expected %q to be the upper case of %q", 'A', 'a')
	}
	if unicode.To(unicode.LowerCase, 'A') != 'a' {
		t.Fatalf("expected %q to be the lower case of %q", 'a', 'A')
	}
	if unicode.To(unicode.TitleCase, 'a') != 'A' {
		t.Fatalf("expected %q to be the title case of %q", 'A', 'a')
	}
	if unicode.To(unicode.TitleCase, '1') != '1' {
		t.Fatalf("expected %q to be the title case of %q", '1', '1')
	}
}

// unicode.ToLower 返回与 rune 等效的小写形式
func TestUnicodeToLower(t *testing.T) {
	// 返回与 rune 等效的小写形式
	if unicode.ToLower('A') != 'a' {
		t.Fatalf("expected %q to be the lower case of %q", 'a', 'A')
	}
	if unicode.ToLower('a') != 'a' {
		t.Fatalf("expected %q to be the lower case of %q", 'a', 'a')
	}
	if unicode.ToLower('1') != '1' {
		t.Fatalf("expected %q to be the lower case of %q", '1', '1')
	}
}

// unicode.ToTitle 返回与 rune 等效的标题大写形式
func TestUnicodeToTitle(t *testing.T) {
	// 返回与 rune 等效的标题大写形式
	if unicode.ToTitle('a') != 'A' {
		t.Fatalf("expected %q to be the title case of %q", 'A', 'a')
	}
	if unicode.ToTitle('A') != 'A' {
		t.Fatalf("expected %q to be the title case of %q", 'A', 'A')
	}
	if unicode.ToTitle('1') != '1' {
		t.Fatalf("expected %q to be the title case of %q", '1', '1')
	}
}

// unicode.ToUpper 返回与 rune 等效的大写形式
func TestUnicodeToUpper(t *testing.T) {
	// 返回与 rune 等效的大写形式
	if unicode.ToUpper('a') != 'A' {
		t.Fatalf("expected %q to be the upper case of %q", 'A', 'a')
	}
	if unicode.ToUpper('A') != 'A' {
		t.Fatalf("expected %q to be the upper case of %q", 'A', 'A')
	}
	if unicode.ToUpper('1') != '1' {
		t.Fatalf("expected %q to be the upper case of %q", '1', '1')
	}
}
