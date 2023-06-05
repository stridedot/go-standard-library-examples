package builtin_test

import (
	"testing"
)

func TestAppend(t *testing.T) {
	s := []string{"hello "}
	s = append(s, "world ")
	t.Logf("s = %s\n", s)

	s = append(s, "I ", "am ", "here")
	t.Logf("s = %s\n", s)
}

func TestCap(t *testing.T) {
	arr := [3]int{1, 2, 3}
	t.Logf("arr capacity = %d\n", cap(arr))
	t.Logf("&arr capacity = %d\n", cap(&arr))
	t.Logf("slice arr capacity = %d\n", cap(arr[:]))

	var v []int
	t.Logf("v = %v\n", v)
	t.Logf("nil capacity = %d\n", cap(v))
}

// close函数是一个内建函数，用来关闭channel，这个channel要么是双向的，要么是只写(只发送)的（chan<- Type）。
// 这个方法应该只由发送者调用， 而不是接收者。
// 当最后一个发送的值都被接收者从关闭的channel(下简称为c)中接收时,
// 接下来所有接收的值都会非阻塞直接成功，返回channel元素的零值。
// 如下的代码：
// 如果c已经关闭（c中所有值都被接收）， x, ok := <- c， 读取ok将会得到false。
func TestClose(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan bool)

	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}(ch1)

	go func(ch1 chan int, ch2 chan bool) {
		for {
			v, ok := <-ch1
			if ok {
				t.Logf("readed int is %d\n", v)
			} else {
				// 断言判断，如果 ok = false，则写入 channel
				ch2 <- true
			}
		}
	}(ch1, ch2)

	<-ch2
}

// 复数运算法则
func TestComplex(t *testing.T) {
	x := complex(1, 2)
	y := complex(3, 4)
	t.Logf("x + y = %v\n", x+y)
	t.Logf("x * y = %v\n", x*y)
}

// 测试复制到新的切片
func TestCopy(t *testing.T) {
	src := []byte("abcdefg")
	dst := make([]byte, 7)
	n := copy(dst, src)
	if n != len(src) {
		t.Fatalf("Want length %d, got length %d", len(src), n)
	}
}

func TestDelete(t *testing.T) {
	m := map[int]string{0: "A", 1: "B", 2: "C"}
	delete(m, 1)
	t.Logf("m = %v\n", m)
}

// 测试获取复数的实数和虚数部分
func TestImag(t *testing.T) {
	x := complex(1, 2)
	y := complex(3, 4)
	s := x + y
	t.Logf("x + y = %v\n", s)
	t.Logf("real of s is %v, imag of s is %v\n", real(s), imag(s))
}

// 数组：v 中元素的数量。
// 数组指针：*v 中元素的数量（即使 v 为 nil）。
// 切片或映射：v 中元素的数量；若 v 为 nil，len(v) 即为零。
// 字符串：v 中字节的数量。
// channel：信道缓存中队列（未读取）元素的数量；若 v 为 nil，len(v) 即为零。
func TestLen(t *testing.T) {
	arr := [3]int{1, 2, 3}
	t.Logf("arr len = %d\n", len(arr))
	t.Logf("&arr len = %d\n", len(&arr))
	t.Logf("slice arr len = %d\n", len(arr[:]))

	var s []int
	t.Logf("nil slice len = %d\n", len(s))
}

// slice：
//
//	size 指定了切片长度。该切片的容量等于其长度。
//	第二个整数实参可用来指定不同的容量；
//	它必须不小于其长度，因此 make([]int, 0, 10)
//	会分配一个长度为0，	容量为10的切片。
//
// map：
//
//	初始分配取决于 size，但产生的 map 长度为0。
//	size 可以省略，这种情况下就会分配一个最小的起始容量。
//
// chan：
//
//	信道的缓存根据指定的缓存容量初始化。
//	若 size 为零或被省略，该信道即为无缓存的。
func TestMake(t *testing.T) {
	s := make([]int, 0, 5)
	t.Logf("len = %d, cap = %d\n", len(s), cap(s))
}

// 测试 panic 调用
func TestPanic(t *testing.T) {
	defer func() {
		t.Logf("here will print: %s\n", "before panic")
	}()

	panic("手动出发 panic 调用~~~~")

	defer func() {
		t.Logf("here will not print: %s\n", "after panic")
	}()
}

// 使用 recover 捕获异常
func TestRecover(t *testing.T) {
	t.Log("Hello")
	try(func() {
		panic("World")
	}, func(e interface{}) {
		t.Logf("catch: %v\n", e)
	})
}

func try(fun func(), catch func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catch(err)
		}
	}()
	fun()
}
