package bytes_test

import (
	"bytes"
	"io"
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

// TestIndexByte 返回字节在 s 切片中第一次出现的位置索引
func TestIndexByte(t *testing.T) {
	s := []byte("Good morning")
	t.Logf("First instance of c in b is %d\n", bytes.IndexByte(s, 'o'))
}

// TestIndexFunc 返回 s 中第一个满足 f(c) 的 Unicode 代码点的字节索引，如果没有就返回 -1
func TestIndexFunc(t *testing.T) {
	s := []byte("Go programming 编程")
	f := func(r rune) bool {
		return unicode.Is(unicode.Han, r)
	}
	t.Logf("index = %d\n", bytes.IndexFunc(s, f))
}

// TestIndexRune 返回 rune 字符在 s 中第一次出现的索引
func TestIndexRune(t *testing.T) {
	s := []byte("从前有座山，山上有座庙，庙里有个老和尚")
	t.Logf("index = %d\n", bytes.IndexRune(s, '有'))
}

// TestJoin 使用 sep 拼接 s 切片
func TestJoin(t *testing.T) {
	s := [][]byte{[]byte("bar"), []byte("foo"), []byte("baz")}
	t.Logf("join: %s\n", bytes.Join(s, []byte(", ")))
}

// TestLastIndex 返回 s 中 sep 最后一次出现的索引
func TestLastIndex(t *testing.T) {
	s := []byte("go gopher")
	t.Logf("LastIndex: %d\n", bytes.LastIndex(s, []byte("go")))
}

// TestLastIndexAny 返回 chars 中任一个字符在 s 中最后一次出现的字节索引
func TestLastIndexAny(t *testing.T) {
	s := []byte("go 地鼠")
	t.Logf("LastIndexAny: %d\n", bytes.LastIndexAny(s, "地球"))
}

// TestLastIndexByte 返回字节 c 在 s 中最后一次出现的字节索引
func TestLastIndexByte(t *testing.T) {
	s := []byte("go gopher")
	t.Logf("LastIndexByte: %d\n", bytes.LastIndexByte(s, 'o'))
}

// TestLastIndexFunc 返回满足 f(c) 的最后一个字符在 s 中的索引
func TestLastIndexFunc(t *testing.T) {
	s := []byte("go gopher！")
	t.Logf("LastIndexFunc: %d\n", bytes.LastIndexFunc(s, unicode.IsLetter))
	t.Logf("LastIndexFunc: %d\n", bytes.LastIndexFunc(s, unicode.IsPunct))
	t.Logf("LastIndexFunc: %d\n", bytes.LastIndexFunc(s, unicode.IsNumber))
}

// TestMap 对所有字符根据函数进行修改
func TestMap(t *testing.T) {
	s := []byte("There are three children in the park")
	f := func(r rune) rune {
		if r == 'e' {
			return 'd'
		}
		return r
	}
	b := bytes.Map(f, s)
	t.Logf("bytes.Map: %s\n", b)
}

// TestRepeat 返回有 b 的 count 个副本组成的新切片
func TestRepeat(t *testing.T) {
	t.Logf("repeat: %s\n", bytes.Repeat([]byte("# "), 5))
}

// TestReplace 将 s 中的前 n 个 k 替换为 ky
func TestReplace(t *testing.T) {
	s := []byte("oink oink oink")
	t.Logf("replace: %s\n", bytes.Replace(s, []byte("k"), []byte("ky"), -1))
}

// TestReplaceAll 将 s 中的所有 k 替换为 ky
func TestReplaceAll(t *testing.T) {
	s := []byte("oink oink oink")
	t.Logf("replaceAll: %s\n", bytes.ReplaceAll(s, []byte("k"), []byte("ky")))
}

// TestRunes 将 byte 字符转为 rune 字符
func TestRunes(t *testing.T) {
	s := []byte("go gopher")
	t.Logf("bytes.Runes: %v\n", string(bytes.Runes(s)))
}

// TestSplit 将 s 按照 sep 分隔，返回子切片
func TestSplit(t *testing.T) {
	s := []byte("a,b,c")
	t.Logf("bytes.Split: %q\n", bytes.Split(s, []byte(",")))
	t.Logf("bytes.Split: %q\n", bytes.Split([]byte(" abc "), []byte("")))
	t.Logf("bytes.Split: %q\n", bytes.Split([]byte(""), []byte("")))
}

// TestSplitAfter 将 s 按照 sep 分割，返回带有 sep 的子切片
func TestSplitAfter(t *testing.T) {
	t.Logf("bytes.SplitAfter: %q\n", bytes.SplitAfter([]byte("a,b,c"), []byte(",")))
}

// TestSplitAfterN 将 s 按照 sep 分割，返回 n 个带有 sep 的子切片
func TestSplitAfterN(t *testing.T) {
	// output: ["a," "b,c"]
	t.Logf("bytes.SplitAfterN: %q\n", bytes.SplitAfterN([]byte("a,b,c"), []byte(","), 2))
}

// TestSplitN 将 s 按照 sep 分隔，返回 n 个子切片
func TestSplitN(t *testing.T) {
	t.Logf("bytes.SplitN: %q\n", bytes.SplitN([]byte("a,b,c"), []byte(","), 2))
}

// TestToLower 将 s 转为小写
func TestToLower(t *testing.T) {
	t.Logf("bytes.ToLower: %s\n", bytes.ToLower([]byte("ABC")))
}

// TesToLowerSpecial 将 s 转为小写，优先考虑指定小写规则
func TestToLowerSpecial(t *testing.T) {
	t.Logf("bytes.ToLowerSpecial: %s\n", bytes.ToLowerSpecial(unicode.AzeriCase, []byte("ahoj vývojári golang")))
}

// TestToTitle 转为大写，注意 ToTitle 和 ToUpper 的区别
func TestToTitle(t *testing.T) {
	t.Logf("bytes.ToTitle: %s\n", bytes.ToTitle([]byte("ǳ")))
	t.Logf("bytes.ToUpper: %s\n", bytes.ToUpper([]byte("ǳ")))
}

// TestToTitleSpecial 转为大写，优先考虑大写规则
func TestToTitleSpecial(t *testing.T) {
	s := []byte("ahoj vývojári golang")
	t.Logf("bytes.ToTitleSpecial:%s\n", bytes.ToTitleSpecial(unicode.AzeriCase, s))
}

func TestToUpper(t *testing.T) {
	t.Logf("bytes.ToUpper:%s\n", bytes.ToUpper([]byte("Gopher")))
}

func TestToUpperSpecial(t *testing.T) {
	t.Logf("bytes.ToUpperSpecial:%s\n", bytes.ToUpperSpecial(unicode.AzeriCase, []byte("ahoj vývojári golang")))
}

// TesToValidUTF8 将无效 utf-8 的字节替换为指定字节
func TestToValidUTF8(t *testing.T) {
	s := bytes.ToValidUTF8([]byte("typical"), []byte("\uFFFD"))
	t.Logf("s = %s\n", s)

	s = bytes.ToValidUTF8([]byte("foo\xffbar"), []byte("\uFFFD"))
	t.Logf("s = %s\n", s)

	s = bytes.ToValidUTF8([]byte("日本語\xff日本語"), []byte("\uFFFD"))
	t.Logf("s = %s\n", s)
}

// TestTrim 将 s 两端指定的字符移除
func TestTrim(t *testing.T) {
	t.Logf("%s\n", bytes.Trim([]byte("!!! Achtung! Achtung! !!!"), "! "))
}

// TestTrimFun 将 s 按指定函数移除两端字符
func TestTrimFunc(t *testing.T) {
	t.Logf("IsLetter: %s\n", bytes.TrimFunc([]byte("go-gopher!"), unicode.IsLetter))
	t.Logf("IsLetter: %s\n", bytes.TrimFunc([]byte("\"go-gopher!\""), unicode.IsLetter))
	t.Logf("IsPunct: %s\n", bytes.TrimFunc([]byte("go-gopher!"), unicode.IsPunct))
	t.Logf("IsNumber: %s\n", bytes.TrimFunc([]byte("1234go-gopher!567"), unicode.IsNumber))
}

// TestTrimLeft 移除 s 左边字符
func TestTrimLeft(t *testing.T) {
	t.Logf("%s\n", bytes.TrimLeft([]byte("123go-gopher!"), "012134"))
}

// TestTrimLeftFunc 将 s 按指定函数移除左端字符
func TestTrimLeftFunc(t *testing.T) {
	t.Logf("IsNumber: %s\n", bytes.TrimLeftFunc([]byte("1234go-gopher!567"), unicode.IsNumber))
}

// TestTrimPrefix 返回 s 中没有指定前缀的部分
func TestTrimPrefix(t *testing.T) {
	t.Logf("trimPrefix:%s\n", bytes.TrimPrefix([]byte("Goodbye, teacher"), []byte("Good")))
}

// TestTrimSpace 移除 s 中两端的空白
func TestTrimSpace(t *testing.T) {
	t.Logf("bytes.TrimSpace: %s\n", bytes.TrimSpace([]byte(" \t\n a lone gopher \n\t\r\n")))
}

// TestTrimSuffix 返回 s 中没有指定后缀的部分
func TestTrimSuffix(t *testing.T) {
	t.Logf("bytes.TrimSuffix:%s\n", bytes.TrimSuffix([]byte("Goodbye, teacher"), []byte("er")))
}

// TestNewBuffer 初始化一个 buffer
func TestNewBuffer(t *testing.T) {
	buf := bytes.NewBuffer([]byte("123456"))
	var data [10]byte
	n, err := buf.Read(data[:])
	if n != 6 {
		t.Fatalf("Read error: %v\n", err)
	}
}

// TestNewBufferString 初始化一个以字符串为参数的新 buffer
func TestNewBufferString(t *testing.T) {
	buf := bytes.NewBufferString("abbbaaaba")
	var data [10]byte
	n, err := buf.Read(data[:])
	if n != 9 {
		t.Fatalf("Read error: %v\n", err)
	}
}

// TestBufferBytes 返回一个切片
func TestBufferBytes(t *testing.T) {
	buf := bytes.Buffer{}
	buf.Write([]byte{'h', 'e', 'l', 'l', 'o'})
	t.Logf("buf.Bytes: %s\n", buf.Bytes())
}

// TestBufferCap 返回切片的容量
func TestBufferCap(t *testing.T) {
	buf1 := bytes.NewBuffer(make([]byte, 10))
	buf2 := bytes.NewBuffer(make([]byte, 0, 10))
	t.Logf("buf1 cap = %d\n", buf1.Cap())
	t.Logf("buf2 cap = %d\n", buf2.Cap())
}

// TestBufferGroup 增加缓冲区的容量
func TestBufferGrow(t *testing.T) {
	var buf bytes.Buffer
	t.Logf("buf = %q\n", buf.Bytes())

	buf.Grow(64)
	buf.Write([]byte("more or fewer"))
	t.Logf("buf = %q\n", buf.Bytes())
}

// TestBufferLen 返回缓冲区的字节数
func TestBufferLen(t *testing.T) {
	buf := bytes.NewBuffer([]byte("abc"))
	if buf.Len() != len(buf.Bytes()) {
		t.Fatalf("Want %d, got %d", buf.Len(), len(buf.Bytes()))
	}
}

// TestBufferNext 返回缓冲区中 n 个字节数的切片
func TestBufferNext(t *testing.T) {
	buf := bytes.NewBuffer([]byte("123456"))
	t.Logf("next = %s\n", buf.Next(2))
	t.Logf("next = %s\n", buf.Next(2))
	t.Logf("bytes.Bytes = %s, len = %d\n", buf.Bytes(), buf.Len())
}

// TestBufferRead 返回读取的字节
func TestBufferRead(t *testing.T) {
	var b bytes.Buffer
	b.Grow(15)
	b.Write([]byte("123456"))

	buf := make([]byte, 7)
	n, err := b.Read(buf)
	if err != nil {
		t.Fatalf("Got err: %v", err)
	}

	buf1 := make([]byte, 7)
	n, err = b.Read(buf1)
	if err != io.EOF {
		t.Fatalf("Want empty, got %s\n", buf1[:n])
	}
}

// TestBufferReadByte 测试返回字节
func TestBufferReadByte(t *testing.T) {
	s := make([]byte, 10)
	buf := bytes.NewBuffer([]byte("123456"))

	for {
		b, err := buf.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("buf.ReadByte err: %v\n", err)
		}
		s = append(s, b)
	}

	t.Logf("s = %s\n", s)
}

// TestBufferReadBytes 测试返回字节数组
func TestBufferReadBytes(t *testing.T) {
	buf := bytes.NewBuffer([]byte("bar,foo,baz"))
	line, err := buf.ReadBytes(',')
	if err != nil {
		t.Fatalf("buf.ReadBytes err: %v\n", err)
	}
	// output "bar,"
	t.Logf("buf.ReadBytes = %s\n", line)
}

// TestBufferReadFrom 从 reader 中读取数据
func TestBufferReadFrom(t *testing.T) {
	b := bytes.Buffer{}
	buf := bytes.NewBuffer([]byte("123456"))
	_, err := b.ReadFrom(buf)
	if err != nil {
		t.Fatalf("b.ReadFrom err: %v\n", err)
	}
	t.Logf("b = %s\n", b.String())
}

// TestBufferReadRune 返回一个 utf-8 编码的字符
func TestBufferReadRune(t *testing.T) {
	b := bytes.NewBuffer([]byte("你好，世界"))
	c, size, err := b.ReadRune()
	if err != nil {
		t.Fatalf("b.ReadRune err: %v\n", err)
	}
	t.Logf("c = %s, size = %d\n", string(c), size)
}

// TestBufferReadString 返回包含 delim 第一次出现的字符数据
func TestBufferReadString(t *testing.T) {
	s := "gopher$phper$rust"
	b := bytes.NewBufferString(s)

	line, err := b.ReadString('$')
	t.Logf("err = %v\n", err)
	t.Logf("line = %v\n", line)
}

// TestBufferReset 将缓冲区重置
func TestBufferReset(t *testing.T) {
	var b bytes.Buffer
	b.Reset()
	b.WriteString("go语言学习")
	t.Logf("%s\n", b.String())
}

// TestBufferTruncate 丢弃缓冲区中的所有字节，但保留前 n 个字节
// 如果 n < 0 或 n > b.Len()，将会抛出异常
func TestBufferTruncate(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fatal("Truncate(-1) should have panicked")
		}
	}()
	var b bytes.Buffer
	b.Truncate(-1)
}

// TestBufferUnreadByte 取消读取最近读取的字节
func TestBufferUnreadByte(t *testing.T) {
	var b bytes.Buffer
	b.Write([]byte("let's go"))
	n, err := b.ReadByte()
	t.Logf("n = %d, err = %v\n", n, err)
	t.Logf("b.String = %s\n", b.String())

	err = b.UnreadByte()
	if err != nil {
		t.Fatalf("UnreadByte err: %v\n", err)
	}
	t.Logf("b.String = %s\n", b.String())

	line, err := b.ReadString('\'')
	t.Logf("line = %s, err = %v\n", line, err)
	t.Logf("b.String = %s\n", b.String())

	err = b.UnreadByte()
	if err != nil {
		t.Fatalf("UnreadByte err: %v\n", err)
	}
	t.Logf("b.String = %s\n", b.String())
}

// TestBufferUnreadRune 取消读取最后一个 utf-8 字符
func TestBufferUnreadRune(t *testing.T) {
	b := bytes.NewBuffer([]byte("一二三四五六七八九十"))
	_, _, err := b.ReadRune()
	if err != nil {
		t.Fatalf("ReadRune err: %v", err)
	}
	err = b.UnreadRune()
	if err != nil {
		t.Fatalf("UnreadRune err: %v", err)
	}
	t.Logf("String: %s\n", b.String())
}

// TestBufferWrite 写入
func TestBufferWrite(t *testing.T) {
	var b bytes.Buffer
	_, err := b.Write([]byte("Go is simple secure scalable language"))
	if err != nil {
		t.Fatalf("Write err: %v\n", err)
	}
}

// TestBufferWriteByte 写入字节
func TestBufferWriteByte(t *testing.T) {
	var b bytes.Buffer
	b.WriteByte('a')
	b.WriteByte('b')
	b.WriteByte('c')
	b.WriteByte('d')
	t.Logf("WriteByte: %s\n", b.String())
}

// TestBufferWriteRune 写入 rune 字符
func TestBufferWriteRune(t *testing.T) {
	var b bytes.Buffer
	b.WriteRune('你')
	b.WriteRune('好')
	b.WriteRune('世')
	b.WriteRune('界')
	t.Logf("WriteRune: %s\n", b.String())
}

// TestBufferWriteString 写入字符串
func TestBufferWriteString(t *testing.T) {
	b := new(bytes.Buffer)
	b.WriteString("别来无恙")
	t.Logf("WriteString: %s\n", b.String())
}

// TestBufferWriteTo 将数据写入 w
func TestBufferWriteTo(t *testing.T) {
	b1 := bytes.NewBuffer([]byte("Sleep fragrance"))
	b2 := new(bytes.Buffer)
	_, err := b1.WriteTo(b2)
	if err != nil {
		t.Fatalf("WriteTo err: %v\n", err)
	}
}

// TestNewReader 返回一个新 reader
func TestNewReader(t *testing.T) {
	reader := bytes.NewReader([]byte("123456"))
	if reader.Len() != 6 {
		t.Fatalf("Want 6, got %d\n", reader.Len())
	}
}

// TestReaderRead 读取
func TestReaderRead(t *testing.T) {
	reader := bytes.NewReader([]byte("123456"))
	buf := make([]byte, 3)
	_, err := reader.Read(buf)
	if err != nil {
		t.Fatalf("Read err: %v\n", err)
	}
}

// TestReaderReadAt 返回从指定偏移量读取的数据
func TestReaderReadAt(t *testing.T) {
	reader := bytes.NewReader([]byte("Go语言学习"))
	p := make([]byte, 6)
	n, err := reader.ReadAt(p, 2)
	if err != nil {
		t.Fatalf("ReadAt err: %v", err)
	}
	t.Logf("ReadAt: %s\n", p[:n])
}

// TestReadReadRune 读取 utf-8 编码的字符
func TestReaderReadRune(t *testing.T) {
	reader := bytes.NewReader([]byte("123456你好"))
	for {
		c, size, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("ReadRune err: %v\n", err)
		}
		t.Logf("c = %s, size = %d\n", string(c), size)
	}
}

// TestReaderReset 将 Reader 重置为从 b 读取
func TestReaderReset(t *testing.T) {
	var b bytes.Reader
	b.Reset([]byte("123456"))
	buf, err := io.ReadAll(&b)
	if err != nil {
		t.Fatalf("Got err: %v\n", err)
	}
	t.Logf("buf = %s\n", buf)
}

// TestReaderSeek 设置下次读取的位置
func TestReaderSeek(t *testing.T) {
	b := bytes.NewReader([]byte("123456"))
	n, err := b.Seek(2, io.SeekStart)
	t.Logf("n = %d, err = %v\n", n, err)

	c, err := b.ReadByte()
	t.Logf("byte = %s, err = %v\n", string(c), err)

	b1 := bytes.NewReader([]byte("abcdefg"))
	b1.ReadByte()
	b1.ReadByte()
	n, err = b1.Seek(1, 1)
	c, _ = b1.ReadByte()
	if string(c) != "d" {
		t.Fatalf("Want d, got %s\n", string(c))
	}

	b2 := bytes.NewReader([]byte("abcdefg"))
	n, _ = b2.Seek(-2, 2)
	c, _ = b2.ReadByte()
	if string(c) != "f" {
		t.Fatalf("Want f, got %s\n", string(c))
	}
}
