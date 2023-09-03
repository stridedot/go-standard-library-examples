package path_test

import (
	"io/fs"
	"path/filepath"
	"testing"
)

// filepath 包用于操作文件名路径，可处理正斜杠和反斜杠的文件路径

// 返回给定 path 的绝对路径
func TestFilepathAbs(t *testing.T) {
	path := "path_test.go"
	absPath, err := filepath.Abs(path)
	if err != nil {
		t.Fatalf("Abs failed, err = %v", err)
	}
	t.Logf("filepath.Abs(%q) = %q\n", path, absPath)
}

// 返回给定 path 的最后一个元素
func TestFilepathBase(t *testing.T) {
	paths := []string{
		"/foo/bar/baz.js",
		"/foo/bar/baz",
		"/foo/bar/baz/",
		"dev.txt",
		"../todo.txt",
		"..",
		".",
		"/",
		"",
	}

	for _, path := range paths {
		t.Logf("filepath.Base(%q) = %q\n", path, filepath.Base(path))
	}
}

// 见 path.Clean()
func TestFilepathClean(t *testing.T) {

}

// 见 path.Dir()
func TestFilepathDir(t *testing.T) {

}

// todo
func TestFilepathEvalSymlinks(t *testing.T) {

}

// 见 path.Ext()
func TestFilepathExt(t *testing.T) {

}

// 返回将路径中的每个斜杠（"/"）字符替换为分隔符的结果
func TestFilepathFromSlash(t *testing.T) {
	s := "/a/b/c.css"
	t.Logf("filepath.FromSlash(%q) = %q\n", s, filepath.FromSlash(s))
}

// 返回匹配模式的所有文件
func TestFilepathGlob(t *testing.T) {
	files, err := filepath.Glob("./*.go")
	if err != nil {
		t.Fatalf("Glob failed, err = %v", err)
	}

	for _, file := range files {
		t.Logf("file = %v\n", file)
	}
}

// 见 path.IsAbs
func TestFilepathIsAbs(t *testing.T) {

}

func TestFilepathIsLocal(t *testing.T) {
	type IsLocalTest struct {
		path    string
		isLocal bool
	}
	localTests := []IsLocalTest{
		{"NUL", false},
		{"nul", false},
		{"nul.", false},
		{"com1", false},
		{"./nul", false},
		{`\`, false},
		{`\a`, false},
		{`C:`, false},
		{`C:\a`, false},
		{`..\a`, false},
		{`a/../c:`, false},
		{`CONIN$`, false},
		{`conin$`, false},
		{`CONOUT$`, false},
		{`conout$`, false},
		{`dollar$`, true}, // not a special file name
	}
	for _, test := range localTests {
		t.Logf("filepath.IsLocal(%q) = `%t`\n", test.path, filepath.IsLocal(test.path))
	}
}

// 见 path.Join
func TestFilepathJoin(t *testing.T) {

}

// 见 path.Match
func TestFilepathMath(t *testing.T) {

}

// 返回以 basepath 为基准的相对路径
func TestFilepathRel(t *testing.T) {
	paths := []string{
		"/a/b/c",
		"/b/c",
		"./b/c",
	}
	base := "/a"

	// on unix
	for _, p := range paths {
		rel, err := filepath.Rel(base, p)
		t.Logf("%q: %q %v\n", p, rel, err)
	}
}

// 见 path.Split
func TestFilepathSplit(t *testing.T) {

}

// 将路径字符串使用路径列表分隔符分开。路径列表分隔符见os.PathListSeparator,
// linux的路径列表分隔符是":", windows的路径列表分隔符是";"，
// 主要用于PATH或是GOPATH等环境变量。
func TestFilepathSplitList(t *testing.T) {
	t.Log("on unix", filepath.SplitList("/a/b/c:/usr/bin"))
	t.Log("on windows", filepath.SplitList("/a/b/c;/usr/bin"))
}

// 返回用斜杠（"/"）字符替换路径中每个分隔符的结果
// windows 文件路径分隔符："\\"，见 os.PathSeparator
func TestFilepathToSlash(t *testing.T) {
	t.Log("ToSlash:", filepath.ToSlash("\\a\\b\\c"))
}

// 返回前导卷名
func TestFilepathVolumeName(t *testing.T) {
	t.Log(filepath.VolumeName("C:\\foo\\bar"))
}

// 使用 walkFun 遍历 path 目录下的所有文件
func TestFilepathWalk(t *testing.T) {
	walk := func(path string, info fs.FileInfo, err error) error {
		t.Logf("path = %v\n", path)
		return nil
	}

	err := filepath.Walk("/home", walk)
	if err != nil {
		t.Fatalf("error = %v", err)
	}
}

// 使用 walkDirFun 遍历 path 目录下的所有文件
func TestFilepathWalkDir(t *testing.T) {
	fn := func(path string, d fs.DirEntry, err error) error {
		t.Logf("path = %v\n", path)
		return nil
	}
	filepath.WalkDir(".", fn)
}
