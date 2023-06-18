package encoding_test

import (
	"bytes"
	"encoding/binary"
	"math"
	"testing"
)

// 测试 binary.Write
func TestBinaryWrite(t *testing.T) {
	buf := new(bytes.Buffer)
	pi := math.Pi
	err := binary.Write(buf, binary.LittleEndian, pi)
	if err != nil {
		t.Fatalf("Write err: %v", err)
	}
	t.Logf("buf = %v\n", buf.Bytes())
}

// 测试 binary 写入多个
func TestBinaryWriteMulti(t *testing.T) {
	buf := new(bytes.Buffer)
	data := []any{
		uint16(61374),
		int8(-54),
		uint8(254),
	}
	for _, v := range data {
		_ = binary.Write(buf, binary.LittleEndian, v)
	}
	t.Logf("buf = %v\n", buf.String())
}

// 测试 binary PutVarint
// 将可变长值写入内存并返回写入的字节数，有符号
func TestBinaryPutVarint(t *testing.T) {
	buf := make([]byte, binary.MaxVarintLen64)
	for _, x := range []int64{-65, -64, -2, -1, 0, 1, 2, 63, 64} {
		n := binary.PutVarint(buf, x)
		t.Logf("%x\n", buf[:n])
	}
}

// 测试 binary PutUvarint
// 将可变长值写入内存，无符号
func TestBinaryPutUvarint(t *testing.T) {
	buf := make([]byte, binary.MaxVarintLen64)

	for _, x := range []uint64{1, 2, 127, 128, 255, 256} {
		n := binary.PutUvarint(buf, x)
		t.Logf("%x\n", buf[:n])
	}
}

// 测试 binary.Varint
// 将字节码转为十进制
func TestBinaryVarint(t *testing.T) {
	inputs := [][]byte{
		{0x81, 0x01},
		{0x7f},
		{0x03},
		{0x01},
		{0x00},
		{0x02},
		{0x04},
		{0x7e},
		{0x80, 0x01},
	}
	for _, b := range inputs {
		x, n := binary.Varint(b)
		if n != len(b) {
			t.Fatal("Varint did not consume all of in")
		}
		t.Logf("x = %v\n", x)
	}
}

func TestBinaryUvarint(t *testing.T) {
	inputs := [][]byte{
		{0x01},
		{0x02},
		{0x7f},
		{0x80, 0x01},
		{0xff, 0x01},
		{0x80, 0x02},
	}
	for _, b := range inputs {
		x, n := binary.Uvarint(b)
		if n != len(b) {
			t.Fatal("Varint did not consume all of in")
		}
		t.Logf("x = %v\n", x)
	}
}
