package strings_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
	"unicode"
)

// strings.Clone 返回 s 的副本，
// 如果 s 为空则返回空字符串。
func TestStringsClone(t *testing.T) {
	s := "abc"
	clone := strings.Clone(s)
	t.Logf("s: %p, clone: %p", &s, &clone)
}

// strings.Compare 按字典顺序比较两个字符串，
func TestStringsCompare(t *testing.T) {
	t.Logf("Compare: %d", strings.Compare("a", "b"))
	t.Logf("Compare: %d", strings.Compare("a", "a"))
	t.Logf("Compare: %d", strings.Compare("b", "a"))
}

// strings.Contains 判断字符串 s 是否包含子串 substr。
func TestStringsContains(t *testing.T) {
	t.Logf("Contains: %t", strings.Contains("seafood", "foo"))
	t.Logf("Contains: %t", strings.Contains("seafood", "bar"))
	t.Logf("Contains: %t", strings.Contains("seafood", ""))
	t.Logf("Contains: %t", strings.Contains("", ""))
}

// strings.ContainsAny 判断字符串 s 中是否包含 chars 中的任一字符。
func TestStringsContainsAny(t *testing.T) {
	t.Logf("ContainsAny: %t", strings.ContainsAny("team", "i"))
	t.Logf("ContainsAny: %t", strings.ContainsAny("failure", "u & i"))
	t.Logf("ContainsAny: %t", strings.ContainsAny("foo", ""))
	t.Logf("ContainsAny: %t", strings.ContainsAny("", ""))
}

// strings.ContainsFunc 判断字符串 s 中是否包含满足函数 f 的任一字符。
func TestStringsContainsFunc(t *testing.T) {
	//f := func(c rune) bool {
	//	return c == 'a' || c == '1'
	//}
	//t.Logf("ContainsFunc: %t", strings.ContainsFunc("team", f('a')))
	//t.Logf("ContainsFunc: %t", strings.ContainsFunc("failure", f('u')))
	//t.Logf("ContainsFunc: %t", strings.ContainsFunc("foo", f('o')))
	//t.Logf("ContainsFunc: %t", strings.ContainsFunc("", f('1')))
}

// strings.ContainsRune 判断字符串 s 中是否包含字符 r。
func TestStringsContainsRune(t *testing.T) {
	t.Logf("ContainsRune: %t", strings.ContainsRune("team", 'a'))
	t.Logf("ContainsRune: %t", strings.ContainsRune("failure", 'u'))
	t.Logf("ContainsRune: %t", strings.ContainsRune("foo", 'o'))
	t.Logf("ContainsRune: %t", strings.ContainsRune("", '1'))
}

// strings.Count 计算字符串 s 在字符串 str 中出现的非重叠次数。
func TestStringsCount(t *testing.T) {
	t.Logf("Count: %d", strings.Count("cheese", "e"))
	t.Logf("Count: %d", strings.Count("five", "")) // before & after each rune
}

// strings.Cut 返回将 s 按 sep 分割的所有子串，结果中不包含 sep 自身。
func TestStringsCut(t *testing.T) {
	show := func(s, sep string) {
		before, after, found := strings.Cut(s, sep)
		t.Logf("Cut: %q, %q, %v", before, after, found)
	}
	show("Gopher", "Go")
	show("Gopher", "ph")
	show("Gopher", "er")
	show("Gopher", "Badger")
}

// strings.CutPrefix 返回将 s 前缀 prefix 去除的结果。
func TestStringsCutPrefix(t *testing.T) {
	show := func(s, prefix string) {
		after, found := strings.CutPrefix(s, prefix)
		t.Logf("CutPrefix: %q, %v", after, found)
	}
	show("Gopher", "Go")
	show("Gopher", "ph")
}

// strings.CutSuffix 返回将 s 后缀 suffix 去除的结果。
func TestStringsCutSuffix(t *testing.T) {
	show := func(s, suffix string) {
		before, found := strings.CutSuffix(s, suffix)
		t.Logf("CutSuffix: %q, %v", before, found)
	}
	show("Gopher", "er")
	show("Gopher", "ph")
}

// strings.EqualFold 判断 s 和 t 是否相等，忽略大小写。
func TestStringsEqualFold(t *testing.T) {
	t.Logf("EqualFold: %t", strings.EqualFold("Go", "go"))
	t.Logf("EqualFold: %t", strings.EqualFold("Go", "Go"))
}

// strings.Fields 将字符串按空白分割，返回一个字符串的切片。
func TestStringsFields(t *testing.T) {
	t.Logf("Fields: %q", strings.Fields("  foo bar  baz   "))
}

// strings.FieldsFunc 将字符串按满足函数 f 的字符分割，返回一个字符串的切片。
func TestStringsFieldsFunc(t *testing.T) {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	t.Logf("FieldsFunc: %q", strings.FieldsFunc("  foo1;bar2,baz3...", f))
}

// strings.HasPrefix 判断字符串 s 是否以 prefix 开头。
func TestStringsHasPrefix(t *testing.T) {
	t.Logf("HasPrefix: %t", strings.HasPrefix("Gopher", "Go"))
	t.Logf("HasPrefix: %t", strings.HasPrefix("Gopher", "C"))
	t.Logf("HasPrefix: %t", strings.HasPrefix("Gopher", ""))
}

// strings.HasSuffix 判断字符串 s 是否以 suffix 结尾。
func TestStringsHasSuffix(t *testing.T) {
	t.Logf("HasSuffix: %t", strings.HasSuffix("Amigo", "go"))
	t.Logf("HasSuffix: %t", strings.HasSuffix("Amigo", "O"))
	t.Logf("HasSuffix: %t", strings.HasSuffix("Amigo", ""))
}

// strings.Index 返回字符串 str 在字符串 s 中的第一个索引，-1 表示字符串 s 不包含字符串 str。
func TestStringsIndex(t *testing.T) {
	t.Logf("Index: %d", strings.Index("chicken", "ken"))
	t.Logf("Index: %d", strings.Index("chicken", "dmr"))
}

// strings.IndexAny 返回字符串 chars 中的任一字符在字符串 s 中第一次出现的位置，-1 表示字符串 s 不包含字符串 chars。
func TestStringsIndexAny(t *testing.T) {
	t.Logf("IndexAny: %d", strings.IndexAny("chicken", "aeiouy"))
	t.Logf("IndexAny: %d", strings.IndexAny("crwth", "aeiouy"))
}

// strings.IndexByte 返回字符 c 在字符串 s 中第一次出现的位置，-1 表示字符串 s 不包含字符 c。
func TestStringsIndexByte(t *testing.T) {
	t.Logf("IndexByte: %d", strings.IndexByte("golang", 'g'))
	t.Logf("IndexByte: %d", strings.IndexByte("gophers", 'h'))
	t.Logf("IndexByte: %d", strings.IndexByte("golang", 'x'))
}

// strings.IndexFunc 返回满足函数 f 的第一个字符的索引，-1 表示字符串 s 中没有满足函数 f 的字符。
func TestStringsIndexFunc(t *testing.T) {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	t.Logf("IndexFunc: %d", strings.IndexFunc("Hello, 世界", f))
	t.Logf("IndexFunc: %d", strings.IndexFunc("Hello, world", f))
}

// strings.IndexRune 返回字符 r 在字符串 s 中第一次出现的位置，-1 表示字符串 s 不包含字符 r。
func TestStringsIndexRune(t *testing.T) {
	t.Logf("IndexRune: %d", strings.IndexRune("chicken", 'k'))
	t.Logf("IndexRune: %d", strings.IndexRune("chicken", 'd'))
}

// strings.Join 将字符串切片 a 用字符串 sep 连接，返回生成的字符串。
func TestStringsJoin(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	t.Logf("Join: %s", strings.Join(s, ", "))
}

// strings.LastIndex 返回字符串 str 在字符串 s 中最后一次出现的位置，-1 表示字符串 s 不包含字符串 str。
func TestStringsLastIndex(t *testing.T) {
	t.Logf("LastIndex: %d", strings.LastIndex("go gopher", "go"))
	t.Logf("LastIndex: %d", strings.LastIndex("go gopher", "rodent"))
}

// strings.LastIndexAny 返回字符串 chars 中的任一字符在字符串 s 中最后一次出现的位置，-1 表示字符串 s 不包含字符串 chars。
func TestStringsLastIndexAny(t *testing.T) {
	t.Logf("LastIndexAny: %d", strings.LastIndexAny("go gopher", "go"))
	t.Logf("LastIndexAny: %d", strings.LastIndexAny("go gopher", "rodent"))
	t.Logf("LastIndexAny: %d", strings.LastIndexAny("go gopher", "fail"))
}

// strings.LastIndexByte 返回字符 c 在字符串 s 中最后一次出现的位置，-1 表示字符串 s 不包含字符 c。
func TestStringsLastIndexByte(t *testing.T) {
	t.Logf("LastIndexByte: %d", strings.LastIndexByte("Hello, world", 'o'))
	t.Logf("LastIndexByte: %d", strings.LastIndexByte("Hello, world", 'x'))
}

// strings.LastIndexFunc 返回满足函数 f 的最后一个字符的索引，-1 表示字符串 s 中没有满足函数 f 的字符。
func TestStringsLastIndexFunc(t *testing.T) {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	t.Logf("LastIndexFunc: %d", strings.LastIndexFunc("Hello, 世界", f))
	t.Logf("LastIndexFunc: %d", strings.LastIndexFunc("Hello, world", f))
}

// strings.Map 将字符串 s 中满足函数 mapping 的字符替换为 mapping(c)，返回替换后的字符串。
func TestStringsMap(t *testing.T) {
	t.Logf("Map: %s", strings.Map(func(r rune) rune {
		return r + 1
	}, "HAL-9000"))
}

// strings.Repeat 返回将字符串 s 重复 count 次的新字符串。
func TestStringsRepeat(t *testing.T) {
	t.Logf("Repeat: %s", strings.Repeat("na", 2))
	t.Logf("Repeat: %s", strings.Repeat("na", -1))
}

// strings.Replace 返回将字符串 s 中前 n 个不重叠的 old 子串替换为 new 子串的新字符串，如果 n < 0 则替换所有的 old 子串。
func TestStringsReplace(t *testing.T) {
	t.Logf("Replace: %s", strings.Replace("oink oink oink", "k", "ky", 2))
	t.Logf("Replace: %s", strings.Replace("oink oink oink", "oink", "moo", -1))
}

// strings.ReplaceAll 返回将字符串 s 中所有的 old 子串替换为 new 子串的新字符串。
func TestStringsReplaceAll(t *testing.T) {
	t.Logf("ReplaceAll: %s", strings.ReplaceAll("oink oink oink", "oink", "moo"))
}

// strings.Split 将字符串 s 按照 sep 分割，返回一个字符串的切片。
func TestStringsSplit(t *testing.T) {
	t.Logf("Split: %q", strings.Split("a,b,c", ","))
	t.Logf("split. %q", strings.Split("a man a plan a canal panama", "a "))
	t.Logf("split. %q", strings.Split(" xyz ", ""))
}

// strings.SplitAfter 将字符串 s 按照 sep 分割，返回一个字符串的切片，包含 sep。
func TestStringsSplitAfter(t *testing.T) {
	t.Logf("SplitAfter: %q", strings.SplitAfter("a,b,c", ","))
	t.Logf("SplitAfter: %q", strings.SplitAfter("a,b,c", ""))
	t.Logf("SplitAfter: %q", strings.SplitAfter("a,b,c", "z"))
}

// strings.SplitAfterN 将字符串 s 按照 sep 分割，返回一个字符串的切片，包含 sep，最多分割 n 个子串。
func TestStringsSplitAfterN(t *testing.T) {
	t.Logf("SplitAfterN: %q", strings.SplitAfterN("a,b,c", ",", 2))
	t.Logf("SplitAfterN: %q", strings.SplitAfterN("a,b,c", ",", -1))
	t.Logf("SplitAfterN: %q", strings.SplitAfterN("a,b,c", ",", 0))
}

// strings.SplitN 将字符串 s 按照 sep 分割，返回一个字符串的切片，不包含 sep，最多分割 n 个子串。
func TestStringsSplitN(t *testing.T) {
	t.Logf("SplitN: %q", strings.SplitN("a,b,c", ",", 2))
	t.Logf("SplitN: %q", strings.SplitN("a,b,c", ",", -1))
	t.Logf("SplitN: %q", strings.SplitN("a,b,c", ",", 0))
}

// strings.ToLower 将字符串 s 转换为小写。
func TestStringsToLower(t *testing.T) {
	t.Logf("ToLower: %s", strings.ToLower("Gopher"))
	t.Logf("ToLower: %s", strings.ToLower("Gopher"))
}

// strings.ToLowerSpecial 将字符串 s 转换为小写，使用特定的规则。
func TestStringsToLowerSpecial(t *testing.T) {
	t.Logf("ToLowerSpecial: %s", strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))
}

// strings.ToTitle 将字符串 s 转换为标题格式。
func TestStringsToTitle(t *testing.T) {
	t.Logf("ToTitle: %s", strings.ToTitle("loud noises"))
	t.Logf("ToTitle: %s", strings.ToTitle("хлеб"))
}

// strings.ToTitleSpecial 将字符串 s 转换为标题格式，使用特定的规则。
func TestStringsToTitleSpecial(t *testing.T) {
	t.Logf("ToTitleSpecial: %s", strings.ToTitleSpecial(unicode.TurkishCase, "örnek iş"))
}

// strings.ToUpper 将字符串 s 转换为大写。
func TestStringsToUpper(t *testing.T) {
	t.Logf("ToUpper: %s", strings.ToUpper("Gopher"))
	t.Logf("ToUpper: %s", strings.ToUpper("Gopher"))
}

// strings.ToUpperSpecial 将字符串 s 转换为大写，使用特定的规则。
func TestStringsToUpperSpecial(t *testing.T) {
	t.Logf("ToUpperSpecial: %s", strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))
}

// strings.ToValidUTF8 将字符串 s 中的无效 UTF-8 字符替换为 replacement。
func TestStringsToValidUTF8(t *testing.T) {
	t.Logf("ToValidUTF8: %s", strings.ToValidUTF8("hello\x80world", ""))
}

// strings.Trim 返回将字符串 s 前后端所有 cutset 包含的字符去除的字符串。
func TestStringsTrim(t *testing.T) {
	t.Logf("Trim: %s", strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))
}

// strings.TrimFunc 返回将字符串 s 前后端满足函数 f 的字符去除的字符串。
func TestStringsTrimFunc(t *testing.T) {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	t.Logf("TrimFunc: %s", strings.TrimFunc("¡¡¡Hello, Gophers!!!", f))
}

// strings.TrimLeft 返回将字符串 s 前端所有 cutset 包含的字符去除的字符串。
func TestStringsTrimLeft(t *testing.T) {
	t.Logf("TrimLeft: %s", strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡"))
}

// strings.TrimLeftFunc 返回将字符串 s 前端满足函数 f 的字符去除的字符串。
func TestStringsTrimLeftFunc(t *testing.T) {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	t.Logf("TrimLeftFunc: %s", strings.TrimLeftFunc("¡¡¡Hello, Gophers!!!", f))
}

// strings.TrimPrefix 返回将字符串 s 前缀 prefix 去除的结果。
func TestStringsTrimPrefix(t *testing.T) {
	t.Logf("TrimPrefix: %s", strings.TrimPrefix("¡¡¡Hello, Gophers!!!", "¡¡¡Hello, "))
}

// strings.TrimRight 返回将字符串 s 后端所有 cutset 包含的字符去除的字符串。
func TestStringsTrimRight(t *testing.T) {
	t.Logf("TrimRight: %s", strings.TrimRight("¡¡¡Hello, Gophers!!!", "!¡"))
}

// strings.TrimRightFunc 返回将字符串 s 后端满足函数 f 的字符去除的字符串。
func TestStringsTrimRightFunc(t *testing.T) {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	t.Logf("TrimRightFunc: %s", strings.TrimRightFunc("¡¡¡Hello, Gophers!!!", f))
}

// strings.TrimSpace 返回将字符串 s 前后端所有空白字符去除的字符串。
func TestStringsTrimSpace(t *testing.T) {
	t.Logf("TrimSpace: %s", strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))
}

// strings.TrimSuffix 返回去除字符串后缀的 s
func TestStringsTrimSuffix(t *testing.T) {
	var s = "¡¡¡Hello, Gophers!!!"
	t.Logf("TrimSuffix: %s", strings.TrimSuffix(s, ", Gophers!!!"))
}

// strings.Builder 使用 Write 方法有效地构建字符串
func TestStringsBuilder(t *testing.T) {
	var b strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("ignition")
	t.Logf("b = %s", b.String())
}

// strings.Builder.Cap() 返回底层切片地容量
func TestStringsBuilderCap(t *testing.T) {
	var b strings.Builder
	t.Logf("b.Cap = %v\n", b.Cap())

	b.WriteString("hello world")
	t.Logf("b.Cap = %v\n", b.Cap())
	t.Logf("b.Len = %v\n", b.Len())

	b.Reset()
	t.Logf("b.Cap = %v\n", b.Cap())
}

// strings.Builder.Grow() 增加 b 的容量
func TestStringsBuilderGrow(t *testing.T) {
	var b strings.Builder
	b.Grow(10)
	t.Logf("b.Cap = %v\n", b.Cap())
	t.Logf("b.Len = %v\n", b.Len())

	b.Reset()
	t.Logf("b.Cap = %v\n", b.Cap())
}

// strings.Builder.Write() 将 p 中的内容写入 b
func TestStringsBuilderWrite(t *testing.T) {
	var b strings.Builder
	b.Write([]byte("hello world"))
	t.Logf("b = %s\n", b.String())

	b.WriteString("你好，世界")
	t.Logf("b = %s\n", b.String())

	b.WriteRune('a')
	t.Logf("b = %s\n", b.String())
}

// strings.Reader
func TestStringsReader(t *testing.T) {
	reader := strings.NewReader("sf[,r2wre")
	t.Logf("Len = %v\n", reader.Len())

	b := make([]byte, 10)
	n, err := reader.Read(b)
	if err != nil {
		t.Fatalf("Read failed, err = %v", err)
	}
	t.Logf("Len = %s\n", b[:n])

	reader.Reset("very nice!!!")

	b = make([]byte, 10)
	n, err = reader.ReadAt(b, 2)
	if err != nil {
		t.Fatalf("ReatAt failed, err = %v", err)
	}
	t.Logf("Len = %s\n", b[:n])

	reader.Reset("wuhan")
	b1, err := reader.ReadByte()
	if err != nil {
		t.Fatalf("ReadByte failed, err = %v", err)
	}
	t.Logf("byte = %c\n", b1)

	reader.Reset("中国")
	ch, size, err := reader.ReadRune()
	if err != nil {
		t.Fatalf("ReadRune failed, err = %v", err)
	}
	t.Logf("ch = %c, size = %d\n", ch, size)
}

// strings.Reader.Seek
func TestStringsReaderSeek(t *testing.T) {
	reader := strings.NewReader(" returns the original length")
	off, err := reader.Seek(0, io.SeekStart)
	if err != nil {
		t.Fatalf("Seek failed, %v", off)
	}
	t.Logf("off = %v\n", off)
}

// strings.NewReplacer从旧的、新的字符串对列表中返回一个新的 Replacer
func TestStringsNewReplacer(t *testing.T) {
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	t.Logf("replace = %v", r.Replace("This is <b>HTML</b>!"))
}

// strings.WriteString 将 s 写入 w 并执行所有替换
func TestStringsReplacerWriteString(t *testing.T) {
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	writer := bytes.NewBuffer(nil)
	_, _ = r.WriteString(writer, "replacer")
	t.Logf("S = %v", writer.String())
}
