package encoding_test

import (
	"bytes"
	"encoding/csv"
	"io"
	"strings"
	"testing"
)

// 测试 ReadAll csv文件数据
func TestCsvNewReader(t *testing.T) {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))
	r.Comma = ','
	records, err := r.ReadAll()
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	t.Logf("ReadAll() result = %v\n", records)
}

// 测试 Read csv 数据
func TestCsvRead(t *testing.T) {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Read err = %v", err)
		}
		t.Logf("Record = %q\n", record)
	}
}

// FieldPos返回对应于最近由Read返回的片断中具有给定索引的字段开始的行和列。
// 行和列的编号从1开始;列的计数单位是字节,而不是符码。
func TestCsvFieldPos(t *testing.T) {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Read err = %v", err)
		}
		t.Logf("Record = %q\n", record)

		for i := range record {
			line, column := r.FieldPos(i)
			t.Logf("FieldPos: line = %d, column = %d\n", line, column)
		}
	}
}

// 返回当前读取器位置的输入流字节偏移量。
// 偏移量给出了最近读取行的末尾位置和下一行的开始位置
func TestCsvInputOffset(t *testing.T) {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))
	inputOffset := r.InputOffset()
	t.Logf("InputOffset = %d\n", inputOffset)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Read err = %v", err)
		}
		t.Logf("Record = %q\n", record)

		inputOffset = r.InputOffset()
		t.Logf("InputOffset = %d\n", inputOffset)
	}
}

// 测试 csv.NewWriter 写入数据
func TestCsvWriter(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	w := csv.NewWriter(buf)

	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	for _, record := range records {
		err := w.Write(record)
		if err != nil {
			t.Fatalf("Write err = %v", err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		t.Fatalf("Error = %v\n", err)
	}

	t.Logf("buf.String = %s\n", buf.String())
}

// 测试 csv.WriteAll 写入所有数据
func TestCsvWriteAll(t *testing.T) {
	var benchmarkWriteData = [][]string{
		{"abc", "def", "12356", "1234567890987654311234432141542132"},
		{"abc", "def", "12356", "1234567890987654311234432141542132"},
		{"abc", "def", "12356", "1234567890987654311234432141542132"},
	}

	buf := &bytes.Buffer{}
	w := csv.NewWriter(buf)
	err := w.WriteAll(benchmarkWriteData)
	if err != nil {
		t.Fatalf("WriteAll err := %v", err)
	}
	w.Flush()
	t.Logf("buf = %s\n", buf.String())
}
