package net_test

import (
	"bytes"
	"context"
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
	cname, err := net.LookupCNAME("www.baidu.com")
	if err != nil {
		t.Fatal("LookupCNAME error:", err)
	}

	t.Logf("cname: %s", cname)
}

func TestNetLookupHost(t *testing.T) {
	addrs, err := net.LookupHost("www.baidu.com")
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

// 测试返回IP的默认掩码
func TestNetIPDefaultMask(t *testing.T) {
	ip := net.ParseIP("127.0.0.1")
	t.Logf("ip: %v", ip.DefaultMask())
}

// 测试 net.IP.Equal
func TestNetIPEqual(t *testing.T) {
	ip := net.IPv4(127, 0, 0, 1)
	t.Logf("ip: %v", ip.Equal(net.ParseIP("127.0.0.1")))
}

// 测试net.IP.IsGlobalUnicast
func TestNetIsGlobalUnicast(t *testing.T) {
	ipv4Private := net.ParseIP("10.255.0.0")
	t.Logf("ipv4Private: %v", ipv4Private.IsGlobalUnicast())
}

// 测试返回端点的地址
func TestResolveIPAddr(t *testing.T) {
	addr, err := net.ResolveIPAddr("ip", "www.baidu.com")
	if err != nil {
		t.Fatal("ResolveIPAddr error:", err)
	}

	t.Logf("addr: %v", addr)
	t.Logf("addr.Network(): %v", addr.Network())
	t.Logf("addr.String(): %v", addr.String())
}

// 测试net.DialIP
// laddr 是本地地址，raddr 是远程地址
// 如果 laddr 为 nil，系统会自动选择本地地址
// 如果 raddr 为 nil，系统会假定是地址
func TestNetDialIP(t *testing.T) {
	addr, err := net.ResolveIPAddr("ip", "www.baidu.com")
	if err != nil {
		t.Fatal("ResolveIPAddr error:", err)
	}
	ipconn, err := net.DialIP("ip4:icmp", nil, addr)
	if err != nil {
		t.Logf("DialIP error: %v", err)
	}

	t.Logf("ipconn: %#v", ipconn)
}

// 参数列表:
//
// - ones 网络字符串
// - bits 服务名
//
// 返回列表:
//
// - IPMask 掩码结构
//
// [百度百科 - CIDR](http://baike.baidu.com/view/4217886.htm)
//
// 返回CIDR规范的掩码
// ones 必须 大于0 小于 bits
// bits 必须为 4(IPv4的字节数)或16(IPv6的字节数)的8倍
func TestNetCIDRMask(t *testing.T) {
	t.Logf("mask: %v", net.CIDRMask(24, 32))
}

// 返回 IPv4 掩码
func TestNetIPMask(t *testing.T) {
	ipmask := net.IPv4Mask(255, 255, 255, 0)
	t.Logf("ipmask: %v", ipmask)
	t.Logf("ipmask.String(): %v", ipmask.String())
	ones, bits := ipmask.Size()
	t.Logf("ipmask.Size(): %v, %v", ones, bits)
}

// 测试net.IPNet
// IPNet 表示一个IP网络
// ipnet.Contains(ip) 判断ip是否在ipnet中
// ipnet.Network() 返回此IP网络名称  "ip+net"
// ipnet.String() 返回IP网络的字符串形式
func TestIPNet(t *testing.T) {
	ip := net.IPv4(127, 0, 0, 1)
	ipnet := net.IPNet{
		IP:   ip,
		Mask: ip.DefaultMask(),
	}
	t.Logf("ipnet.Contains(ip): %v", ipnet.Contains(ip))
	t.Logf("ipnet.Network(): %v", ipnet.Network())
	t.Logf("ipnet.String(): %v", ipnet.String())
}

// 测试net.ListenConfig
func TestListenConfig(t *testing.T) {
	lc := net.ListenConfig{KeepAlive: 1000 * time.Second}
	l, err := lc.Listen(context.Background(), "tcp", ":8080")
	if err != nil {
		t.Fatal("Listen error:", err)
	}
	defer l.Close()

	t.Logf("l: %#v", l)
}

// 测试 net.LookupMX	返回域名的MX记录
func TestLookupMX(t *testing.T) {
	mx, err := net.LookupMX("baidu.com")
	if err != nil {
		t.Fatal("LookupMX error:", err)
	}

	t.Logf("mx: %v", mx)
}

// net.DialTCP 的作用同 DialIP，但是网络协议是 TCP
func TestNetDialTcp(t *testing.T) {
	ip := net.IPv4(127, 0, 0, 1)
	tcpAddr, err := net.ResolveTCPAddr("tcp", ip.String()+":8080")
	if err != nil {
		t.Fatal("ResolveTCPAddr error:", err)
	}

	conn, err := net.DialTCP("tcp", tcpAddr, tcpAddr)
	if err != nil {
		t.Fatal("DialTCP error:", err)
	}
	defer conn.Close()

	t.Logf("conn: %#v", conn)
}
