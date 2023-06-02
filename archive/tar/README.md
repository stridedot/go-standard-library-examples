## Overview
`import "archive/tar"`    
tar 包实现了对 tar 格式文件的存取。这个包的目标的是覆盖大多数 tar 格式的变体，包括 GUN 和 BSD 生成的文件。

## Constants

```
const (
	TypeReg           = '0'    // 正常的普通文件
	TypeRegA          = '\x00' // 普通文件（已废弃）
	TypeLink          = '1'    // 硬链接
	TypeSymlink       = '2'    // 符号链接
	TypeChar          = '3'    // 字符设备节点
	TypeBlock         = '4'    // 块设备节点
	TypeDir           = '5'    // 目录
	TypeFifo          = '6'    // 先进先出（FIFO）队列节点
	TypeCont          = '7'    // 保留数字=
	TypeXHeader       = 'x'    // 扩展数据头
	TypeXGlobalHeader = 'g'    // 全局扩展数据头
	TypeGNUSparse     = 'S'    // 表示GNU格式的稀疏文件

	// TypeGNULongName TypeGNULongLink 类型“L”和“K”被GNU格式用于元文件，
	// 用于存储下一个文件的路径或链接名。
	// 这个软件包可以透明地处理这些类型。
	TypeGNULongName = 'L'
	TypeGNULongLink = 'K'
)
```
此常量用于下面 header 结构体中的 Typeflag

## Variables

```
var (
	ErrHeader          = errors.New("archive/tar: invalid tar header")      // 无效的 tar 头部信息
	ErrWriteTooLong    = errors.New("archive/tar: write too long")          // 数据过长
	ErrFieldTooLong    = errors.New("archive/tar: header field too long")   // 头部字段太长
	ErrWriteAfterClose = errors.New("archive/tar: write after close")       // 关闭后写入
	ErrInsecurePath    = errors.New("archive/tar: insecure file path")      // 不安全的文件路径
)
```

## Types

### type Format
```
type Format init
```
`Format` 表示 tar 的存档格式

- func (Format) String() string

### type Header
```
type Header struct {
	Typeflag byte               // 记录头的类型
	Name     string             // 文件头域的名称
	Linkname string             // 链接的目标名
	Size  int64                 // 字节数
	Mode  int64                 // 权限和模式位
	Uid   int                   // 所有者的用户 ID
	Gid   int                   // 所有者的组 ID
	Uname string                // 所有者的用户名
	Gname string                // 所有者的组名
	ModTime    time.Time        // 修改时间
	AccessTime time.Time        // 访问时间
	ChangeTime time.Time        // 状态改变时间

	Devmajor int64 // Major device number (valid for TypeChar or TypeBlock)
	Devminor int64 // Minor device number (valid for TypeChar or TypeBlock)

	// Xattrs stores extended attributes as PAX records under the
	// "SCHILY.xattr." namespace.
	//
	// The following are semantically equivalent:
	//  h.Xattrs[key] = value
	//  h.PAXRecords["SCHILY.xattr."+key] = value
	//
	// When Writer.WriteHeader is called, the contents of Xattrs will take
	// precedence over those in PAXRecords.
	//
	// Deprecated: Use PAXRecords instead.
	Xattrs map[string]string

	// PAXRecords is a map of PAX extended header records.
	//
	// User-defined records should have keys of the following form:
	//	VENDOR.keyword
	// Where VENDOR is some namespace in all uppercase, and keyword may
	// not contain the '=' character (e.g., "GOLANG.pkg.version").
	// The key and value should be non-empty UTF-8 strings.
	//
	// When Writer.WriteHeader is called, PAX records derived from the
	// other fields in Header take precedence over PAXRecords.
	PAXRecords map[string]string

	// Format specifies the format of the tar header.
	//
	// This is set by Reader.Next as a best-effort guess at the format.
	// Since the Reader liberally reads some non-compliant files,
	// it is possible for this to be FormatUnknown.
	//
	// If the format is unspecified when Writer.WriteHeader is called,
	// then it uses the first format (in the order of USTAR, PAX, GNU)
	// capable of encoding this Header (see Format).
	Format Format
}
```