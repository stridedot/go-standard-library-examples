package strings_test

import (
	"strings"
	"testing"
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
