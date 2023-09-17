package text_test

import (
	"strings"
	"testing"
	"text/scanner"
)

// scanner.TokenString 将 tok 转为字符串
func TestScannerTokenString(t *testing.T) {
	tokens := []rune("abcdefg")
	for _, token := range tokens {
		t.Logf("token = %s\n", scanner.TokenString(token))
	}
}

// scanner.Position 代表一个源文件中的位置
func TestScannerPosition(t *testing.T) {
	const src = `
// This is scanned code.
if a > 10 {
	someParsable = text
}`
	var s scanner.Scanner
	reader := strings.NewReader(src)
	s.Init(reader)
	s.Filename = "example"

	s.Next()
	pos := s.Pos()
	if pos.IsValid() == false {
		t.Fatal("Want pos true, got false")
	}
	t.Logf("pos = %+v\n", pos.String())
}

// scanner.Peek 返回下一个 token，但不会移动扫描器，多次 peek，返回相同的 token
func TestScannerPeek(t *testing.T) {
	const src = "int num = 1;"
	var s scanner.Scanner
	reader := strings.NewReader(src)
	s.Init(reader)
	t.Logf("peek1 = %s\n", scanner.TokenString(s.Peek()))
	t.Logf("pos1 = %+v\n", s.Pos())
	t.Logf("peek2 = %s\n", scanner.TokenString(s.Peek()))
	t.Logf("pos2 = %+v\n", s.Pos())
}

// scanner.Next 返回下一个 token，并将扫描器前进到下一个 token
func TestScannerNext(t *testing.T) {
	const src = "int num = 1;"
	var s scanner.Scanner
	reader := strings.NewReader(src)
	s.Init(reader)
	t.Logf("next1 = %s\n", scanner.TokenString(s.Next()))
	t.Logf("pos1 = %+v\n", s.Pos())
	t.Logf("next2 = %s\n", scanner.TokenString(s.Next()))
	t.Logf("pos2 = %+v\n", s.Pos())
	t.Logf("peek3 = %s\n", scanner.TokenString(s.Peek()))
	t.Logf("pos3 = %+v\n", s.Pos())
}

// scanner.Scan 返回下一个 token 的类型
func TestScannerScan(t *testing.T) {
	const src = "int num = 1;"
	var s scanner.Scanner
	reader := strings.NewReader(src)
	s.Init(reader)
	s.Scan()
	t.Logf("scan1 = %s\n", s.TokenText())
	t.Logf("pos1 = %+v\n", s.Pos())
	t.Logf("scan2 = %v\n", s.Scan())
	t.Logf("pos2 = %+v\n", s.Pos())
	t.Logf("peek3 = %s\n", scanner.TokenString(s.Peek()))
	t.Logf("pos3 = %+v\n", s.Pos())
}

// “Scan”方法返回一个“rune”，它代表令牌类型。 “text/scanner”包定义了各种令牌类型的常量，
// 例如 "scanner.Ident"、"scanner.Int"、"scanner.Float"、
// "scanner.Char"、"scanner.String"和"scanner.RawString"
func TestScannerScan2(t *testing.T) {
	input := "x = 42 + 3.14"
	var s scanner.Scanner
	s.Init(strings.NewReader(input))

	for token := s.Scan(); token != scanner.EOF; token = s.Scan() {
		// s.TokenText() 返回 scan 后的 token 的文本
		text := s.TokenText()

		switch token {
		case scanner.Ident:
			t.Logf("Identifier: %s\n", text)
		case scanner.Int:
			t.Logf("Integer: %s\n", text)
		case scanner.Float:
			t.Logf("Float: %s\n", text)
		default:
			t.Logf("Other: %s\n", text)
		}
	}
}
