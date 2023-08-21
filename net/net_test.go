package net_test

import (
	"bytes"
	"net"
	"testing"
	"time"
)

// 测试net.JoinHostPort
func TestNetJoinHostPort(t *testing.T) {
	addr1 := net.JoinHostPort("localhost", "8080")
	addr2 := net.JoinHostPort("localhost:80", "8080")
	t.Logf("addr1: %s", addr1)
	t.Logf("addr2: %s", addr2)
}

func TestNetLookupAddr(t *testing.T) {
	names, err := net.LookupAddr("127.0.0.1")
	if err != nil {
		t.Fatal("LookupAddr error:", err)
	}

	t.Logf("names: %v", names)
}

// 测试net.LookupCNAME
func TestNetLookupCNAME(t *testing.T) {
	cname, err := net.LookupCNAME("any.weizhe.net")
	if err != nil {
		t.Fatal("LookupCNAME error:", err)
	}

	t.Logf("cname: %s", cname)
}

func TestNetLookupHost(t *testing.T) {
	addrs, err := net.LookupHost("github.com")
	if err != nil {
		t.Fatal("LookupHost error:", err)
	}

	t.Logf("addrs: %v", addrs)
}

// 测试net.LookupPort， 返回端口号
func TestNetLookupPort(t *testing.T) {
	port, err := net.LookupPort("tcp", "https")
	if err != nil {
		t.Fatal("LookupPort error:", err)
	}

	t.Logf("port: %d", port)
}

func TestNetLookupTXT(t *testing.T) {
	txts, err := net.LookupTXT("github.com")
	if err != nil {
		t.Fatal("LookupTXT error:", err)
	}

	t.Logf("txts: %v", txts)
}

func TestNetParseCIDR(t *testing.T) {
	ip, ipnet, err := net.ParseCIDR("192.0.2.1/24")
	if err != nil {
		t.Fatal("ParseCIDR error:", err)
	}
	t.Logf("ip: %v, ipnet: %v", ip, ipnet)

	ip, ipnet, err = net.ParseCIDR("2001:db8::1/32")
	if err != nil {
		t.Fatal("ParseCIDR error:", err)
	}
	t.Logf("ip: %v, ipnet: %v", ip, ipnet)
}

// 测试net.Pipe， 返回两个net.Conn
// io下的Pipe(), io下的Pipe()是单向的，一端的读取匹配另一端的写入，直接在这两端之间复制数据
// net下的Pipe()是双向的，也就是说返回的Conn既可以读也可以写，通常用于不同区域的代码之间相互传递数据
func TestNetPipe(t *testing.T) {
	r, w := net.Pipe()
	defer r.Close()

	go func() {
		w.Write([]byte("hello world"))
		w.Close()
	}()

	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		t.Fatal("Read error:", err)
	}

	t.Logf("read: %s", buf[:n])
}

func TestNetSplitHostPort(t *testing.T) {
	host, port, err := net.SplitHostPort("localhost:8080")
	if err != nil {
		t.Fatal("SplitHostPort error:", err)
	}

	t.Logf("host: %s, port: %s", host, port)
}

func TestNetInterfaceAddrs(t *testing.T) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		t.Fatal("InterfaceAddrs error:", err)
	}

	t.Logf("addrs: %v", addrs)
}

func TestNetBuffers(t *testing.T) {
	b := net.Buffers{}
	w := bytes.NewBuffer([]byte("hello world"))

	_, err := b.WriteTo(w)
	if err != nil {
		t.Fatal("WriteTo error:", err)
	}

	buf := make([]byte, 1024)
	n, err := w.Read(buf)
	if err != nil {
		t.Fatal("Read error:", err)
	}

	t.Logf("read: %s", buf[:n])
}

// 测试net.Dial
func TestNetDialer(t *testing.T) {
	conn, err := net.Dial("tcp", "www.163.com:80")
	if err != nil {
		t.Fatal("Dial error:", err)
	}
	defer conn.Close()

	t.Logf("conn: %#v", conn)
}

// 测试net.DialTimeout
func TestNetDialerTimeout(t *testing.T) {
	timeout := time.Duration(10) * time.Second
	conn, err := net.DialTimeout("tcp", "www.google.com:80", timeout)
	if err != nil {
		t.Fatal("DialTimeout error:", err)
	}
	defer conn.Close()

	t.Logf("conn: %#v", conn)
}

// todo 测试net.Dialer
func TestNetDialerFileConn(t *testing.T) {

}

// 测试net.IPv4
func TestNetIPIPv4(t *testing.T) {
	ip := net.IPv4(127, 0, 0, 1)
	t.Logf("ip: %v", ip)
}

// 测试net.IPv4
// 使用本地解析器解析域名，返回IP地址
func TestNetIPLookupIP(t *testing.T) {
	ips, err := net.LookupIP("www.baidu.com")
	if err != nil {
		t.Fatal("LookupIP error:", err)
	}

	t.Logf("ips: %v", ips)
}

func TestNetIPParseIP(t *testing.T) {
	ip := net.ParseIP("127.0.0.1")
	t.Logf("ip: %v", ip)
}
