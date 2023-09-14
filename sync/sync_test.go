package sync_test

import (
	"net/http"
	"sync"
	"testing"
	"time"
)

// sync.NewCond
func TestSyncNewCond(t *testing.T) {
	locker := new(sync.Mutex)
	cond := sync.NewCond(locker)

	var done = false
	read := func(c *sync.Cond, name string) {
		c.L.Lock()

		for !done {
			// 调用 Wait 会自动释放锁 c.L，并挂起调用者所在的 goroutine，
			// 因此当前协程会阻塞在 Wait 方法调用的地方
			c.Wait()
		}
		t.Logf("%s start reading", name)
		c.L.Unlock()
	}
	write := func(c *sync.Cond, name string) {
		t.Logf("%s start writing", name)

		c.L.Lock()
		done = true
		c.L.Unlock()

		t.Logf("%s wakes all", name)
		// 唤醒所有等待的goroutine
		c.Broadcast()
	}

	go read(cond, "reader1")
	go read(cond, "reader2")
	go read(cond, "reader3")
	write(cond, "writer1")
	time.Sleep(time.Second * 3)
}

// sync.Map 读写并发安全的map
func TestSyncMap(t *testing.T) {
	var m sync.Map
	m.Store("name", "zhangsan")
	m.Store("age", 18)
	m.LoadOrStore("sex", "男")

	name, _ := m.Load("name")
	age, _ := m.Load("age")
	sex, _ := m.Load("sex")
	t.Logf("name: %s, age: %d, sex: %s\n", name, age, sex)

	m.LoadAndDelete("sex")
	previous, loaded := m.Swap("name", "lisi")
	t.Logf("previous: %s, loaded: %t\n", previous, loaded)

	m.Range(func(key, value interface{}) bool {
		t.Logf("key: %s, value: %s", key, value)
		return true
	})
}

// sync.Mutex 互斥锁
func TestSyncMutex(t *testing.T) {
	var m sync.Mutex
	m.Lock()
	try := m.TryLock()
	t.Logf("try: %t", try)
	defer m.Unlock()
	t.Log("do something")
}

// sync.Once 保证只执行一次
func TestSyncOnce(t *testing.T) {
	var once sync.Once
	ch := make(chan int, 3)

	for i := 0; i < 3; i++ {
		go func(i int) {
			once.Do(func() {
				t.Log("do something")
			})
			ch <- i
		}(i)
	}

	for i := 0; i < 3; i++ {
		t.Log(<-ch)
	}
}

// sync.Pool 对象池
func TestSyncPool(t *testing.T) {
	var pool sync.Pool
	pool.Put("hello")
	pool.Put("world")
	pool.Put("golang")

	t.Log(pool.Get())
	t.Log(pool.Get())
	t.Log(pool.Get())
	t.Log(pool.Get())
}

// sync.RWMutex 读写锁
func TestSyncRWMutex(t *testing.T) {
	data := make(map[string]string)
	rw := new(sync.RWMutex)

	read := func(key string) {
		rw.RLock()
		defer rw.RUnlock()
		t.Logf("%s = %s", key, data[key])
	}

	write := func(key, value string) {
		rw.Lock()
		defer rw.Unlock()
		data[key] = value
	}

	go read("baz")
	go read("baz")

	write("baz", "hello")
	go read("baz")
	time.Sleep(time.Second * 3)
}

func TestSyncRWMutex2(t *testing.T) {
	var count int
	m := new(sync.RWMutex)

	getCount := func() {
		m.RLock()
		defer m.RUnlock()
		t.Logf("count: %d\n", count)
	}

	setCount := func(c int) {
		m.Lock()
		count = c
		m.Unlock()
	}

	setCount(1)

	go getCount()
	go getCount()
}

func TestSyncRWMutex3(t *testing.T) {
	rwGuard := new(sync.RWMutex)
	wg := new(sync.WaitGroup)
	var a = 10

	read := func() {
		defer wg.Done()
		rwGuard.RLock()
		defer rwGuard.RUnlock()
		t.Logf("read starting..., a = %d\n", a)
	}

	write := func() {
		defer wg.Done()
		rwGuard.Lock()
		defer rwGuard.Unlock()
		t.Log("write starting...")
		a++
		t.Log("write end")
	}

	wg.Add(3)
	go write()
	go read()
	go write()
	wg.Wait()

	t.Log("main end")
}

// sync.WaitGroup 等待一组goroutine执行结束
func TestSyncWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		t.Log("goroutine 1")
	}()

	go func() {
		defer wg.Done()
		t.Log("goroutine 2")
	}()

	go func() {
		defer wg.Done()
		t.Log("goroutine 3")
	}()

	wg.Wait()
	t.Log("main end")
}

func TestSyncWaitGroup2(t *testing.T) {
	// 声明一个等待组
	var wg sync.WaitGroup
	// 准备一系列的网站地址
	var urls = []string{
		"http://www.github.com/",
		"https://www.qiniu.com/",
		"https://www.golangtc.com/",
	}
	// 遍历这些地址
	for _, url := range urls {
		// 每一个任务开始时, 将等待组增加1
		wg.Add(1)
		// 开启一个并发
		go func(url string) {
			// 使用defer, 表示函数完成时将等待组值减1
			defer wg.Done()
			// 使用http访问提供的地址
			_, err := http.Get(url)
			// 访问完成后, 打印地址和可能发生的错误
			t.Logf("url = %s, err = %v\n", url, err)
			// 通过参数传递url地址
		}(url)
	}
	// 等待所有的任务完成
	wg.Wait()
	t.Log("over......")
}
