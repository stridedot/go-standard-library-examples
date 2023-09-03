package path_test

import (
	"path"
	"testing"
)

// path 包可以用于处理 URL 相关的字符串

// Base 返回路径的最后一个元素
func TestPathBase(t *testing.T) {
	t.Log(path.Base("/a/b/c"))
	t.Log(path.Base("/a/b/c/"))
	t.Log(path.Base("/a"))
	t.Log(path.Base("/"))
	t.Log(path.Base(""))
	t.Log(path.Base("\\a\\b"))
}

// 通过纯词法处理返回与 path 等效的最短路径名
func TestPathClean(t *testing.T) {
	paths := []string{
		"a/c",
		"a//c",
		"a/c/.",
		"a/c/b/..",
		"a/c/b/d/..",
		"/../a/c",
		"/../a/b/../././/c",
		"",
		"\\a\\b\\..",
	}

	for _, p := range paths {
		t.Logf("Clean(%q) = %q\n", p, path.Clean(p))
	}
}

// 回路径中除最后一个元素之外的所有元素，通常是路径的目录
func TestPathDir(t *testing.T) {
	dirs := []string{
		"/a/b/c",
		"a/b/c",
		"/a/",
		"a/",
		"/",
		"",
		"\\a\\b",
	}

	for _, d := range dirs {
		t.Logf("Dir(%q) = %q\n", d, path.Dir(d))
	}
}

// 返回路径使用的文件扩展名
func TestPathExt(t *testing.T) {
	paths := []string{
		"/a/b/c.css",
		"/a/b",
		"/",
	}

	for _, d := range paths {
		t.Logf("Ext(%q) = %q\n", d, path.Ext(d))
	}
}

// 是否是绝对路径
func TestPathIsAbs(t *testing.T) {
	paths := []string{
		"/dev/null",
		"..",
		"../../",
	}

	for _, d := range paths {
		t.Logf("IsAbs(%q) = %t\n", d, path.IsAbs(d))
	}
}

// 将多个元素链接到单个路径中
func TestPathJoin(t *testing.T) {
	t.Log(path.Join("a", "b", "c"))
	t.Log(path.Join("a", "b/c"))
	t.Log(path.Join("a/b", "c"))
	t.Log(path.Join("a/b", "../../../xyz"))

	t.Log(path.Join("", ""))
	t.Log(path.Join("a", ""))
	t.Log(path.Join("", "a"))
}

// 使用指定的模式匹配 name
func TestPathMatch(t *testing.T) {
	t.Log(path.Match("abc", "abc"))
	t.Log(path.Match("a*", "abc"))
	t.Log(path.Match("a*/b", "a/c/b"))
}

func TestPathSplit(t *testing.T) {
	split := func(s string) {
		dir, file := path.Split(s)
		t.Logf("path.Split(%q) = dir: %q, file: %q", s, dir, file)
	}

	split("static/c.css")
	split("c.css")
	split("/a/b/c/")
}
