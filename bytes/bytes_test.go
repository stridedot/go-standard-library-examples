package bytes_test

import (
	"bytes"
	"testing"
	"unicode"
)

// TestClone 测试返回一个新的切片备份
// 返回的新的切片
func TestClone(t *testing.T) {
	var b1 []byte
	b2 := []byte("abc")
	b3 := bytes.Clone(b1)
	if b3 != nil {
		t.Fatalf("Want nil, got %v", b3)
	}
	t.Logf("b1 = %v, b2 = %v, b3 = %v\n", b1, b2, b3)

	b4 := bytes.Clone(b2)
	t.Logf("b2 capacity = %d, b4 capacity = %d\n", cap(b2), cap(b4))
}

// TestCompare 测试比较 2 个切片
func TestCompare(t *testing.T) {
	b1 := []byte("abc")
	b2 := []byte("bcd")
	if bytes.Compare(b1, b2) > -1 {
		t.Fatalf("Want b1 is less than b2")
	}

	b1 = []byte("你好")
	b2 = []byte("世界")
	t.Logf("res = %d\n", bytes.Compare(b1, b2))
}

// TestContains 测试一个切片是否包含另一个切片
func TestContains(t *testing.T) {
	b1 := []byte("123456")
	b2 := []byte("345")
	if bytes.Contains(b1, b2) == false {
		t.Fatalf("Want b1 contains b2, got false")
	}

	b3 := []byte("567")
	if bytes.Contains(b1, b3) == true {
		t.Fatalf("Want b1 not contains b3, got true")
	}
}

// TestContainsAny 测试切片中是否包含字符串中的任意字符
func TestContainsAny(t *testing.T) {
	b := []byte("abcd.efg ")
	s1 := "b1234"
	s2 := "1234."
	s3 := ""
	t.Logf("b >>> s1: %t\n", bytes.ContainsAny(b, s1))
	t.Logf("b >>> s2: %t\n", bytes.ContainsAny(b, s2))
	t.Logf("b >>> s3: %t\n", bytes.ContainsAny(b, s3))
}

// TestContainsRune 测试切片中是否包含 rune 字符
func TestContainsRune(t *testing.T) {
	b := []byte("I like this book.")
	r1 := 'b'
	r2 := 'ö'
	r3 := '.'
	t.Logf("b >>> r1: %t\n", bytes.ContainsRune(b, r1))
	t.Logf("b >>> r2: %t\n", bytes.ContainsRune(b, r2))
	t.Logf("b >>> r3: %t\n", bytes.ContainsRune(b, r3))
}

// TestCount 测试一个切片在另一个切片中的个数
func TestCount(t *testing.T) {
	b := []byte("I like this city!")
	s1 := []byte("i")
	s2 := []byte("")
	t.Logf("b >>> s1: %d\n", bytes.Count(b, s1))
	t.Logf("b >>> s2: %d\n", bytes.Count(b, s2))
}

// TestCut 测试使用 sep 切割 b，返回 b 中在 sep 之前和之后的字节
// 注意，此函数是区分大小写的
func TestCut(t *testing.T) {
	b := []byte("PHPer")
	sep := []byte("r")
	before, after, found := bytes.Cut(b, sep)
	t.Logf("before = %q, after = %q, found = %t\n", before, after, found)
}

// TestCutPrefix 测试从字节开头进行分割
func TestCutPrefix(t *testing.T) {
	b := []byte("I like the food")
	prefix := []byte("like")
	after, found := bytes.CutPrefix(b, prefix)
	if found == true {
		t.Fatalf("Want false, got true, after = %v", after)
	}

	prefix = []byte("I ")
	after, found = bytes.CutPrefix(b, prefix)
	if found == false {
		t.Fatalf("Want true, got false, after = %v", after)
	}
}

// TestCutSuffix 测试从字节末尾进行分割
func TestCutSuffix(t *testing.T) {
	b := []byte("I like the food")
	suffix := []byte("the")
	before, found := bytes.CutSuffix(b, suffix)
	if found == true {
		t.Fatalf("Want false, got true, before = %v", before)
	}

	suffix = []byte("food")
	before, found = bytes.CutSuffix(b, suffix)
	if found == false {
		t.Fatalf("Want true, got false, before = %v", string(before))
	}
}

// TestEqual 测试2个切片是否相等（长度和内容）
func TestEqual(t *testing.T) {
	var b1, b2 []byte
	t.Logf("b1 = b2: %t\n", bytes.Equal(b1, b2))
}

// TestEqualFold 测试 2 个字节在 utf-8 编码下是否相等，不区分大小写
func TestEqualFold(t *testing.T) {
	b1 := []byte("Go")
	b2 := []byte("go")
	if bytes.EqualFold(b1, b2) == false {
		t.Fatalf("Want true, got false")
	}
}

// TestFields 将字节按照 utf-8 编码使用空格分割
func TestFields(t *testing.T) {
	b := []byte("  foo bar  baz   ")
	bs := bytes.Fields(b)
	for _, v := range bs {
		t.Logf("v = %v\n", string(v))
	}
}

// TestFieldsFunc 将字节按照 uft-8 编码使用自定义函数进行分割
func TestFieldsFunc(t *testing.T) {
	s := []byte("  foo1;bar2,baz3...")
	f := func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}
	ss := bytes.FieldsFunc(s, f)
	for _, v := range ss {
		t.Logf("v = %v\n", string(v))
	}
}

// TestHasPrefix 测试 s 切片中的前缀是否是 b
func TestHasPrefix(t *testing.T) {
	s := []byte("good good study, day day up")
	b := []byte("go")
	t.Logf("bytes.HasPrefix: %t\n", bytes.HasPrefix(s, b))
	t.Logf("bytes.HasPrefix: %t\n", bytes.HasPrefix(s, []byte("")))
	t.Logf("bytes.HasPrefix: %t\n", bytes.HasPrefix(s, []byte("day")))
}

// TestHasSuffix 测试 s 切片中是否包含指定后缀
func TestHasSuffix(t *testing.T) {
	s := []byte("good good study, day day up")
	t.Logf("btes.HasSuffix: %t\n", bytes.HasSuffix(s, []byte("up")))
}

// TestIndex 返回切片 sep 在切片 s 中第一次出现的位置
func TestIndex(t *testing.T) {
	s := []byte("football")
	sep := []byte("oo")
	if r := bytes.Index(s, sep); r == -1 {
		t.Fatalf("Want index %d, got -1", r)
	}
	t.Logf("index %d\n", bytes.Index(s, sep))
}

// TestIndexAny 返回字符串 chars 任一字符在切片 s 中的位置
func TestIndexAny(t *testing.T) {
	s := []byte("football is very famous")
	chars := "ty"
	if r := bytes.IndexAny(s, chars); r == -1 {
		t.Fatalf("Want index %d, got -d", r)
	}
	t.Logf("index %d\n", bytes.IndexAny(s, chars))
}

func TestIndexByte(t *testing.T) {

}
