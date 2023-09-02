//go:build windows

package os_test

import (
	"io"
	"os"
	"syscall"
	"testing"
	"time"
)

// 改变当前工作目录
func TestOSChdir(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal("Getwd failed")
	}
	t.Logf("Current dir: %s", dir)

	err = os.Chdir("D:\\projects\\go")
	if err != nil {
		t.Fatal("Chdir failed")
	}
	dir, err = os.Getwd()
	if err != nil {
		t.Fatal("Getwd failed")
	}
	t.Logf("Current dir: %s", dir)
}

// 改变文件权限
func TestOSChmod(t *testing.T) {
	err := os.Chmod("D:\\projects\\go\\go_code\\standard-library-examples\\io\\test.txt", 0777)
	if err != nil {
		t.Fatalf("Chmod failed, err: %v", err)
	}
}

// 改变文件所有者
// 该方法只能在 Unix 系统上使用
func TestOSChown(t *testing.T) {
	file, err := os.Create("test.txt")
	if err != nil {
		t.Fatalf("Create failed, err: %v", err)
	}
	defer file.Close()

	err = os.Chown("test.txt", -1, -1)
	if err != nil {
		t.Fatalf("Chown failed, err: %v", err)
	}
}

// 改变文件的最后写入时间和最后访问时间
func TestOSChtimes(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Open failed, err: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		t.Fatalf("Stat failed, err: %v", err)
	}

	sys := fileInfo.Sys().(*syscall.Win32FileAttributeData)
	t.Logf("%+v", sys)

	atime := time.Now()
	mtime := time.Now()
	os.Chtimes("test.txt", atime, mtime)

	fileInfo, err = file.Stat()
	if err != nil {
		t.Fatalf("Stat failed, err: %v", err)
	}
	sys = fileInfo.Sys().(*syscall.Win32FileAttributeData)
	t.Logf("%+v", sys)
}

// 清除环境变量
func TestOSClearEnv(t *testing.T) {
	os.Clearenv()
}

// DirFS为目录dir下的文件树返回一个文件系统(fs.FS)
func TestOSDirFS(t *testing.T) {
	fs := os.DirFS(".")
	file, err := fs.Open("test.txt")
	if err != nil {
		t.Fatalf("Open failed, err: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		t.Fatalf("Stat failed, err: %v", err)
	}

	t.Logf("%+v", fileInfo)
}

// 返回以 "key=value "的形式返回代表环境的字符串副本
func TestOSEnviron(t *testing.T) {
	t.Logf("%+v", os.Environ())
}

// 返回当前可执行的路径
func TestOSExecutable(t *testing.T) {
	pathname, err := os.Executable()
	if err != nil {
		t.Fatalf("Executable failed, err: %v", err)
	}
	t.Logf("Executable path: %s", pathname)
}

// Expand 会根据映射函数替换字符串中的 ${var} 或 $var
func TestOSExpand(t *testing.T) {
	mapper := func(placeholder string) string {
		switch placeholder {
		case "DAY_PART":
			return "morning"
		case "NAME":
			return "Gopher"
		}
		return ""
	}
	t.Logf("%s", os.Expand("Good ${DAY_PART}, ${NAME}!", mapper))
}

// os.Getegid 返回调用者的有效组 ID
func TestOSGetegid(t *testing.T) {
	t.Logf("egid: %d", os.Getegid())
}

// os.Getenv 获取环境变量的值
func TestOSGetenv(t *testing.T) {
	t.Logf("GOPATH: %s", os.Getenv("GOPATH"))
}

// os.Geteuid 返回调用者的有效用户 ID
func TestOSGeteuid(t *testing.T) {
	t.Logf("euid: %d", os.Geteuid())
}

// os.Getgid 返回调用者的组 ID
func TestOSGetgid(t *testing.T) {
	t.Logf("gid: %d", os.Getgid())
}

// os.Getgroups 返回调用者所属的所有用户组的组 ID
// 该方法只能在 Unix 系统上使用
func TestOSGetgroups(t *testing.T) {

}

// os.Getpagesize 返回底层的系统内存页的尺寸
func TestOSGetpagesize(t *testing.T) {
	t.Logf("pagesize: %d", os.Getpagesize())
}

func TestOSGetppid(t *testing.T) {
	t.Logf("ppid: %d", os.Getppid())
}

func TestOSGetuid(t *testing.T) {
	t.Logf("uid: %d", os.Getuid())
}

// os.Getwd 返回当前工作目录的根路径
func TestOSwd(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd failed, err: %v", err)
	}
	t.Logf("wd: %s", dir)
}

// os.Hostname 返回内核提供的主机名
func TestOSHostname(t *testing.T) {
	hostname, err := os.Hostname()
	if err != nil {
		t.Fatalf("Hostname failed, err: %v", err)
	}
	t.Logf("hostname: %s", hostname)
}

// os.IsExist 判断指定的错误是否是文件或目录已经存在的错误
func TestOSIsExist(t *testing.T) {
	t.Logf("IsExist: %v", os.IsExist(os.ErrExist))
}

// os.IsNotExist 判断指定的错误是否是文件或目录不存在的错误
func TestOSIsNotExist(t *testing.T) {
	t.Logf("IsNotExist: %v", os.IsNotExist(os.ErrNotExist))
}

// os.IsPathSeparator 判断指定的字符是否是路径分隔符
func TestOSIsPathSeparator(t *testing.T) {
	t.Logf("IsPathSeparator: %v", os.IsPathSeparator('\\'))
}

// os.IsPermission 判断指定的错误是否是权限不足的错误
func TestOSIsPermission(t *testing.T) {
	t.Logf("IsPermission: %v", os.IsPermission(os.ErrPermission))
}

// os.IsTimeout 判断指定的错误是否是超时错误
func TestOSIsTimeout(t *testing.T) {
	t.Logf("IsTimeout: %v", os.IsTimeout(os.ErrDeadlineExceeded))
}

// os.Lchown 修改指定路径文件的用户 ID 和组 ID
// 该方法只能在 Unix 系统上使用
func TestOSLchown(t *testing.T) {
	err := os.Lchown("test.txt", 0, 0)
	if err != nil {
		t.Fatalf("Lchown failed, err: %v", err)
	}
}

// os.Link 创建硬链接
func TestOSLink(t *testing.T) {
	err := os.Link("test.txt", "test_link.txt")
	if err != nil {
		t.Fatalf("Link failed, err: %v", err)
	}
}

// os.LookupEnv 获取环境变量的值
func TestOSLookupEnv(t *testing.T) {
	value, ok := os.LookupEnv("GOPATH")
	if !ok {
		t.Fatalf("LookupEnv failed")
	}
	t.Logf("GOPATH: %s", value)
}

// os.Mkdir 创建指定名称的目录
func TestOSMkdir(t *testing.T) {
	err := os.Mkdir("test_dir", 0755)
	if err != nil {
		t.Fatalf("Mkdir failed, err: %v", err)
	}
}

// os.MkdirAll 递归创建指定名称的目录
func TestOSMkdirAll(t *testing.T) {
	err := os.MkdirAll("test_dir/test_dir", 0755)
	if err != nil {
		t.Fatalf("MkdirAll failed, err: %v", err)
	}
}

// os.MkdirTemp 在指定的目录中创建一个新的临时目录
func TestOSMkdirTemp(t *testing.T) {
	dir, err := os.MkdirTemp("test_dir", "test")
	if err != nil {
		t.Fatalf("MkdirTemp failed, err: %v", err)
	}
	t.Logf("MkdirTemp: %s", dir)

	dir, err = os.MkdirTemp("test_dir", "*-logs")
	if err != nil {
		t.Fatalf("MkdirTemp failed, err: %v", err)
	}
	t.Logf("MkdirTemp: %s", dir)
}

// os.Pipe 创建一个管道，返回两个文件对象
func TestOSPipe(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Pipe failed, err: %v", err)
	}

	_, err = w.WriteString("hello world")
	if err != nil {
		t.Fatalf("WriteString failed, err: %v", err)
	}

	var buf [128]byte
	n, err := r.Read(buf[:])
	if err != nil {
		t.Fatalf("Read failed, err: %v", err)
	}

	t.Logf("read: %s", buf[:n])
}

// os.ReadFile 读取指定文件的所有内容
func TestOSReadFile(t *testing.T) {
	data, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatalf("ReadFile failed, err: %v", err)
	}
	t.Logf("ReadFile: %s", data)
}

// os.Readlink 读取指定符号链接文件的内容
// 该方法只能在 Unix 系统上使用
func TestOSReadlink(t *testing.T) {
	link, err := os.Readlink("test_link.txt")
	if err != nil {
		t.Fatalf("Readlink failed, err: %v", err)
	}
	t.Logf("Readlink: %s", link)
}

// os.Remove 删除指定的文件或目录
func TestOSRemove(t *testing.T) {
	err := os.Remove("test.txt")
	if err != nil {
		t.Fatalf("Remove failed, err: %v", err)
	}
}

// os.RemoveAll 递归删除指定的文件或目录
func TestOSRemoveAll(t *testing.T) {
	err := os.RemoveAll("test_dir")
	if err != nil {
		t.Fatalf("RemoveAll failed, err: %v", err)
	}
}

// os.Rename 重命名文件或目录
func TestOSRename(t *testing.T) {
	err := os.Rename("test_link.txt", "test_rename.txt")
	if err != nil {
		t.Fatalf("Rename failed, err: %v", err)
	}
}

// os.SameFile 判断两个文件是否是同一个文件
func TestOSSameFile(t *testing.T) {
	file1, err := os.Create("test.txt")
	if err != nil {
		t.Fatalf("Create failed, err: %v", err)
	}
	defer file1.Close()
	fi1, _ := os.Stat("test.txt")

	file2, err := os.Open("test_rename.txt")
	if err != nil {
		t.Fatalf("Open failed, err: %v", err)
	}
	defer file2.Close()
	fi2, _ := os.Stat("test_rename.txt")

	same := os.SameFile(fi1, fi2)
	t.Logf("SameFile: %v", same)
	t.Logf("SameFile: %v", os.SameFile(fi1, fi1))
}

// os.Setenv 设置环境变量的值
func TestOSSetenv(t *testing.T) {
	t.Logf("GOBIN: %s", os.Getenv("GOBIN"))
	err := os.Setenv("GOBIN", "test")
	if err != nil {
		t.Fatalf("Setenv failed, err: %v", err)
	}
	t.Logf("GOBIN: %s", os.Getenv("GOBIN"))
}

// os.Symlink 创建指定名称的符号链接文件
// 该方法只能在 Unix 系统上使用
func TestOSSymlink(t *testing.T) {
	// A required privilege is not held by the client.
	err := os.Symlink("test.txt", "test_symlink.txt")
	if err != nil {
		t.Fatalf("Symlink failed, err: %v", err)
	}
}

// os.TempDir 返回用于临时文件的目录
func TestOSTempDir(t *testing.T) {
	t.Logf("TempDir: %s", os.TempDir())
}

// os.Truncate 修改指定文件的大小
func TestOSTruncate(t *testing.T) {
	err := os.Truncate("test_rename.txt", 10)
	if err != nil {
		t.Fatalf("Truncate failed, err: %v", err)
	}
}

// os.Unsetenv 删除指定的环境变量
func TestOSUnsetenv(t *testing.T) {
	t.Logf("TEMPDIR: %s", os.Getenv("TMPDIR"))
	err := os.Unsetenv("TMPDIR")
	if err != nil {
		t.Fatalf("Unsetenv failed, err: %v", err)
	}
	defer os.Unsetenv("TMPDIR")

	t.Logf("TEMPDIR: %s", os.Getenv("TMPDIR"))
}

// os.UserCacheDir 返回用户的缓存目录
func TestOSUserCacheDir(t *testing.T) {
	dir, err := os.UserCacheDir()
	if err != nil {
		t.Fatalf("UserCacheDir failed, err: %v", err)
	}
	t.Logf("UserCacheDir: %s", dir)
}

// os.UserConfigDir 返回用户的配置目录
func TestOSUserConfigDir(t *testing.T) {
	dir, err := os.UserConfigDir()
	if err != nil {
		t.Fatalf("UserConfigDir failed, err: %v", err)
	}
	t.Logf("UserConfigDir: %s", dir)
}

// os.UserHomeDir 返回用户的主目录
func TestOSUserHomeDir(t *testing.T) {
	dir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("UserHomeDir failed, err: %v", err)
	}
	t.Logf("UserHomeDir: %s", dir)
}

// os.WriteFile 向指定文件写入内容
func TestOSWriteFile(t *testing.T) {
	err := os.WriteFile("test.txt", []byte("Golang study"), 0666)
	if err != nil {
		t.Fatalf("WriteFile failed, err: %v", err)
	}
}

// os.Create 创建指定名称的文件，如果文件已存在则清空
func TestOSCreate(t *testing.T) {
	file, err := os.Create("test.txt")
	if err != nil {
		t.Fatalf("Create failed, err: %v", err)
	}
	defer file.Close()
}

// os.CreateTemp 在指定目录下创建临时文件
// test_dir 目录必须存在
func TestOSCreateTemp(t *testing.T) {
	_ = os.Mkdir("test_dir", 0666)
	file, err := os.CreateTemp("test_dir", "test_*.txt")
	if err != nil {
		t.Fatalf("CreateTemp failed, err: %v", err)
	}
	defer file.Close()

	t.Logf("CreateTemp: %s", file.Name())
}

// os.NewFile 使用指定的文件描述符创建一个文件
func TestOSNewFile(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Open failed, err: %v", err)
	}
	defer file.Close()

	newFile := os.NewFile(file.Fd(), "test1.txt")
	t.Logf("NewFile: %s", newFile.Name())
}

// os.Open 打开指定名称的文件
func TestOSOpen(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Open failed, err: %v", err)
	}
	defer file.Close()
}

// os.OpenFile 打开指定名称的文件
func TestOSOpenFile(t *testing.T) {
	file, err := os.OpenFile("test.txt", os.O_RDWR, 0666)
	if err != nil {
		t.Fatalf("OpenFile failed, err: %v", err)
	}
	defer file.Close()
}

// os.Chdir 修改当前工作目录
// 只能用于 Unix 系统
func TestOSFileChdir(t *testing.T) {
	dir, _ := os.Getwd()
	t.Logf("Getwd: %s", dir)

	f, err := os.Open("test_dir")
	if err != nil {
		t.Fatalf("Open failed, err: %v", err)
	}

	err = f.Chdir()
	if err != nil {
		t.Fatalf("Chdir failed, err: %v", err)
	}

	t.Logf("Getwd: %s", dir)
}

// os.Chmod 修改指定文件的权限
func TestOSFileChmod(t *testing.T) {
	fi, _ := os.Stat("test.txt")
	t.Logf("Chmod1: %v", fi.Mode())

	f, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Open failed, err: %v", err)
	}

	err = f.Chmod(0644)
	if err != nil {
		t.Fatalf("Chmod failed, err: %v", err)
	}

	f.Close()

	fi, _ = os.Stat("test.txt")
	t.Logf("Chmod2: %v", fi.Mode())
}

// os.Chown 修改指定文件的所有者
// 只能用于 Unix 系统
func TestOSFileChown(t *testing.T) {
	fi, err := os.Stat("test.txt")
	if err != nil {
		t.Fatalf("Stat failed, err: %v", err)
	}
	t.Logf("Chown1: %v", fi.Sys())

	f, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Open failed, err: %v", err)
	}

	err = f.Chown(501, 20)
	if err != nil {
		t.Fatalf("Chown failed, err: %v", err)
	}

	f.Close()
}

// os.Fd 返回文件的描述符
func TestOSFileFd(t *testing.T) {
	f, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Open failed, err: %v", err)
	}
	f.Close()

	fd := f.Fd()
	t.Logf("Fd: %v", fd)
}

// os.Name 返回文件的名称
func TestOSFileName(t *testing.T) {
	f, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Open failed, err: %v", err)
	}
	defer f.Close()

	name := f.Name()
	t.Logf("Name: %s", name)
}

// os.Read 从指定文件读取内容
func TestOSFileRead(t *testing.T) {
	f, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Open failed, err: %v", err)
	}
	defer f.Close()

	data := make([]byte, 128)
	n, err := f.Read(data)
	if err != nil {
		t.Fatalf("Read failed, err: %v", err)
	}
	t.Logf("Read: %s", string(data[:n]))
}

// os.ReadAt 从指定位置读取内容
func TestOSReadAt(t *testing.T) {
	f, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Open failed, err: %v", err)
	}
	defer f.Close()

	data := make([]byte, 1000)
	n, err := f.ReadAt(data, 10)
	if err == io.EOF || err == nil {
		t.Logf("ReadAt: %s", string(data[:n]))
		return
	}
	t.Fatalf("ReadAt failed, err: %v", err)
}

// os.ReadDir 读取指定目录下的所有文件和子目录
func TestOSReadDir(t *testing.T) {
	f, err := os.Open(".")
	if err != nil {
		t.Fatalf("Open failed, err: %v", err)
	}
	defer f.Close()

	dir, err := f.ReadDir(10)
	if err != nil {
		t.Fatalf("ReadDir failed, err: %v", err)
	}

	for _, d := range dir {
		t.Logf("ReadDir: %s", d.Name())
	}
}

// os.Readdir	读取指定目录下的所有文件和子目录
func TestOSReaddir(t *testing.T) {
	f, err := os.Open(".")
	if err != nil {
		t.Fatalf("Open failed, err: %v", err)
	}
	defer f.Close()

	dir, err := f.Readdir(10)
	if err != nil {
		t.Fatalf("Readdir failed, err: %v", err)
	}

	for _, d := range dir {
		t.Logf("Readdir: %s", d.Name())
	}
}

// os.Readdirnames 读取指定目录下的所有文件和子目录的名称
func TestOSReaddirnames(t *testing.T) {
	f, err := os.Open(".")
	if err != nil {
		t.Fatalf("Open failed, err: %v", err)
	}
	defer f.Close()

	names, err := f.Readdirnames(10)
	if err != nil {
		t.Fatalf("Readdirnames failed, err: %v", err)
	}

	for _, name := range names {
		t.Logf("Readdirnames: %s", name)
	}
}

// os.Seek 设置指定文件的读写位置
func TestOSSeek(t *testing.T) {
	f, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Open failed, err: %v", err)
	}
	defer f.Close()

	// 设置读写位置为 5
	ret, err := f.Seek(5, io.SeekStart)
	if err != nil {
		t.Fatalf("Seek failed, err: %v", err)
	}

	b := make([]byte, 10)
	n, _ := f.Read(b)
	t.Logf("Seek: %d, %s", ret, string(b[:n]))

	ret, err = f.Seek(5, io.SeekCurrent)
	if err != nil {
		t.Fatalf("Seek failed, err: %v", err)
	}

	b = make([]byte, 10)
	n, _ = f.Read(b)
	t.Logf("Seek: %d, %s", ret, string(b[:n]))

	ret, err = f.Seek(5, io.SeekEnd)
	if err != nil {
		t.Fatalf("Seek failed, err: %v", err)
	}

	b = make([]byte, 10)
	n, _ = f.Read(b)
	t.Logf("Seek: %d, %s", ret, string(b[:n]))
}

// os.SetReadDeadline 设置文件的读取截至时间
func TestOSSetReadline(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Open file failed")
	}
	defer file.Close()

	err = file.SetReadDeadline(time.Now().Add(3 * time.Second))
	if err != nil {
		t.Fatalf("Set read deadline error, err := %v", err)
	}

	time.Sleep(5)
	b := make([]byte, 63350)
	n, err := file.Read(b)
	if err != nil {
		t.Fatalf("Read file failed, err = %v", err)
	}
	t.Logf("read content: %s", b[:n])
}

// os.SetWriteDeadline 设置文件的写入截至时间
func TestOSSetWriteDeadline(t *testing.T) {
	file, err := os.Create("test1.txt")
	if err != nil {
		t.Fatalf("Create file failed, err = %v", err)
	}
	defer file.Close()

	err = file.SetWriteDeadline(time.Now().Add(3 * time.Second))
	if err != nil {
		t.Fatalf("Set write deadline failed, err = %v", err)
	}
}

// os.SetDeadLine 设置文件的读写截至时间，
// 相当于同时调用 SetReadDeadline 和 SetWriteDeadline
func TestOSSetDeadline(t *testing.T) {
	file, err := os.Create("test2.txt")
	if err != nil {
		t.Fatalf("Create file failed, err = %v", err)
	}
	defer file.Close()

	err = file.SetDeadline(time.Now().Add(3 * time.Second))
	if err != nil {
		t.Fatalf("Set deadline err, err = %v", err)
	}
}

// os.Stat 返回描述文件的 FileInfo 结构
func TestOSStat(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Open file failed, err = %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		t.Fatalf("Return fileinfo err, err = %v", err)
	}

	t.Logf("fileInfo = %#v", fileInfo)
}

// 将 file 内存里的内容写到磁盘
// 一般不用调用这个方法
func TestOSSync(t *testing.T) {
	file, err := os.Create("test1.txt")
	if err != nil {
		t.Fatalf("Create file failed, err = %v", err)
	}
	defer file.Close()

	b := []byte("台风来了，五运全停，明天不上班")
	_, err = file.Write(b)
	if err != nil {
		t.Fatalf("Write to file error, err = %v", err)
	}
	file.Sync()
	file.SyscallConn()
}

// f.SyscallConn 返回原始文件
func TestOSSyscallConn(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Open file failed, err = %v", err)
	}

	rconn, err := file.SyscallConn()
	if err != nil {
		t.Fatalf("error, err = %v", err)
	}
	t.Logf("rconn = %v", rconn)
}

// f.Truncate 将文件截取到指定大小
func TestFileTruncate(t *testing.T) {
	file, err := os.Create("test1.txt")
	if err != nil {
		t.Fatalf("Create file failed, err = %v", err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		t.Fatalf("Stat file failed, err = %v", err)
	}
	t.Logf("file size = %v", fi.Size())

	err = file.Truncate(6)
	if err != nil {
		t.Fatalf("Truncate file failed, err = %v", err)
	}

	fi, err = file.Stat()
	if err != nil {
		t.Fatalf("Stat file failed, err = %v", err)
	}
	t.Logf("Truncated file size = %v", fi.Size())
}

// f.Write 将 b 写入到文件
// f.WriteAt 往文件的指定位置写入 b
func TestFileWrite(t *testing.T) {
	file, err := os.Create("test1.txt")
	if err != nil {
		t.Fatalf("Create file failed, err = %v", err)
	}
	defer file.Close()

	_, err = file.Write([]byte("给你了三天，作业都没做完。"))
	if err != nil {
		t.Fatalf("Write to file failed, err = %v", err)
	}

	file.WriteString("明天再写不完，罚你站着听课3天。")
	file.WriteAt([]byte("你听见没？"), 50)

	file.Sync()
}

// os.FindProcess 通过进程 id 获取一个进程对象
// Release 释放与 Process p 关联的所有资源，使其在将来无法使用。仅当未调用 Wait 时才需要调用 Release。
// Signal 向 Process 发送信号。 Windows 上未实现发送中断。
func TestOSProcess(t *testing.T) {
	pid := os.Getpid()
	t.Logf("pid = %v", pid)
	p, err := os.FindProcess(pid)
	if err != nil {
		t.Fatalf("Find process failed, err = %v", err)
	}
	t.Logf("p = %#v", p)

	// p.Kill()
	// p.Release()
	// p.Signal(os.Kill)
}

func TestOSStartProcessWait(t *testing.T) {

}
