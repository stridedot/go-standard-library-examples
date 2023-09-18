package unsafe_test

import (
	"testing"
	"unsafe"
)

func TestUnsafePointer(t *testing.T) {
	var a int
	var s string
	var f float64
	t.Logf("a = %q, &a = %p", a, unsafe.Pointer(&a))
	t.Logf("s = %q, &s = %p", s, unsafe.Pointer(&s))
	t.Logf("f = %v, &f = %p", f, unsafe.Pointer(&f))
}

func TestUnsafePointer2(t *testing.T) {
	var arr = []int{177, 123, 3, 221, 5, 1211}

	pointer := unsafe.Pointer(&arr[1])
	t.Logf("pointer = %p\n", pointer)

	uPointer := uintptr(pointer)
	t.Logf("uPointer = %v\n", uPointer)

	uPointer += 8
	t.Logf("uPointer + 8 = %v\n", uPointer)

	pointer = unsafe.Pointer(uPointer)
	t.Logf("uPointer => pointer = %p\n", pointer)

	intPointer := (*int)(pointer)
	t.Logf("intPointer = %v\n", *intPointer)
}

func TestUnsafePointer3(t *testing.T) {
	t.Log(unsafe.Sizeof("komeiji satori"))
	t.Log(unsafe.Sizeof("satori"))

	// len 返回的底层数组的长度，而不是字符串的长度
	name := "琪露诺"
	t.Logf("len = %d\n", len(name))
	t.Logf("[]byte(name) = %v\n", []byte(name))

	// 转为 []rune 切片后，len 返回的是字符串的长度
	t.Logf("len([]rune(name)) = %d\n", len([]rune(name)))

	s := "憨pi"
	t.Log([]byte(s))
	t.Log([]rune(s))

	s1 := []byte{230, 134, 168, 112, 105}
	t.Logf("s1 = %s\n", s1)

	s2 := []rune{25000, 112, 105}
	t.Logf("s2 = %v\n", string(s2))
}

func TestUnsafePointer4(t *testing.T) {
	s := "abc"
	s1 := []byte(s)
	t.Logf("p1 = %p, p1 = %p\n", &s, unsafe.Pointer(&s))
	t.Logf("p2 = %p\n", &s1)
}

func TestUnsafePointer5(t *testing.T) {
	s1 := []int8{1, 2, 3, 4}
	t.Logf("pointer of s1 = %p\n", &s1)

	s2 := *(*[]int16)(unsafe.Pointer(&s1))
	t.Logf("bytes of s2 = %p\n", s2)
}

// 字符串和切片共享数组
func TestUnsafePointer6(t *testing.T) {
	s := "abc"
	slice := (*[]byte)(unsafe.Pointer(&s))
	t.Logf("s = %p, slice = %p\n", &s, slice)
	t.Logf("s = %v, slice = %v\n", s, *slice)

	// 结果：cap of slice = 12433427
	// 原因是字符串没有容量，转成切片时，容量丢失
	t.Logf("cap of slice = %v\n", cap(*slice))

	// 字符串转为切片
	slice2 := stringToBytes(s)
	t.Logf("stringToBytes = %v\n", slice2)
	// 切片转为字符串
	t.Logf("bytesToString = %v\n", bytesToString(*slice))

	// 底层数组是否相同
	slice1 := []byte{97, 98, 99}
	s2 := bytesToString(slice1)
	t.Logf("s2 = %v\n", s2)
	slice1[0] = 'A'
	t.Logf("s2 = %v\n", s2)
}

// 字符串转为切片
func stringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&struct {
		string
		Cap int
	}{s, len(s)}))
}

// 切片转为字符串
func bytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 新版
func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// 新版
func BytesToString(b []byte) string {
	return unsafe.String(&b[0], unsafe.IntegerType(len(b)))
}
