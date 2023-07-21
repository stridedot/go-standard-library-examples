package encoding_test

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"
)

// 参数列表:
//
// - dst 表示字符缓冲区指针
// - src 表示JSON格式的字符串切片
//
// 返回值:
//
// - 返回error错误信息
// 这个函数主要是用于将JSON格式的src追加到dst中，正确则返回nil，如果发生错误则返回error信息
func TestJsonCompact(t *testing.T) {
	dst := new(bytes.Buffer)
	src := []byte(`{
		"Name":"tony.shao",
		"Age":25,
		"Job":"Programmer"
	}`)
	err := json.Compact(dst, src)
	if err != nil {
		t.Fatalf("Compact err := %v", err)
	}
	t.Logf("dst = %v\n", dst.String())
}

// 测试 json html 转义
func TestJsonHtmlEscape(t *testing.T) {
	b := new(bytes.Buffer)
	json.HTMLEscape(b, []byte(`{"Name": "<b>HTML content</b>"}`))
	t.Logf("b = %v\n", b)
}

// 测试 json 缩进
func TestJsonIdent(t *testing.T) {
	type Road struct {
		Name   string
		Number int
	}
	roads := []Road{
		{"Diamond Fork", 29},
		{"Sheep Creek", 51},
	}
	b, err := json.Marshal(roads)
	if err != nil {
		t.Fatalf("Marshal err: %v", err)
	}
	out := new(bytes.Buffer)
	_ = json.Indent(out, b, "=", "\t")
	_, _ = out.WriteTo(os.Stdout)
}

// encode 数据
func TestJsonMarshal(t *testing.T) {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"red", "yellow", "grey"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		t.Fatalf("Marshal err: %v", err)
	}
	_, _ = os.Stdout.Write(b)
}

// decode 数据
func TestJsonUnmarshal(t *testing.T) {
	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`)
	type Animal struct {
		Name  string
		Order string
	}

	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		t.Fatalf("Unmarshal err: %v", err)
	}
	t.Logf("animals = %v\n", animals)
}

// 带有缩进的 encoded 数据
func TestJsonMarshalIndent(t *testing.T) {
	data := map[string]int{
		"a": 1,
		"b": 2,
	}
	b, err := json.MarshalIndent(data, "---", "\t")
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	t.Logf("b = %s\n", b)
}

// 判断 json 是否有效
func TestJsonValid(t *testing.T) {
	json1 := `{"name": "Job"}`
	json2 := `{"name": "Ben"}}`
	if json.Valid([]byte(json1)) == false {
		t.Fatal("Expected true, return false")
	}
	if json.Valid([]byte(json2)) == true {
		t.Fatalf("Expected false, return true")
	}
}

// 测试 encode
// 测试转义 HTML 字符
func TestJsonEncoder(t *testing.T) {
	type Dummy struct {
		Name string
		Age  int
	}
	dummy := Dummy{
		Name: "<Duo",
		Age:  25,
	}
	b := new(bytes.Buffer)
	enc := json.NewEncoder(b)
	enc.SetEscapeHTML(true)
	err := enc.Encode(dummy)
	if err != nil {
		t.Fatalf("Encode err:%v", err)
	}
	t.Logf("encoded: %v\n", b.String())
}

// 测试 json 缩进
func TestJsonSetIndent(t *testing.T) {
	var streamTest = []any{
		0.1,
		"hello",
		nil,
		true,
		false,
		[]any{"a", "b", "c"},
		map[string]any{"K": "Kelvin", "ß": "long s"},
		3.14, // another value to make sure something can follow map
	}
	var buf strings.Builder
	enc := json.NewEncoder(&buf)
	enc.SetIndent(">", ".")
	for _, v := range streamTest {
		enc.Encode(v)
	}
	t.Log(buf.String())
}

// 测试 decode 数据
func TestJsonDecoder(t *testing.T) {
	const jsonStream = `{"Name": "Ed", "Text": "Knock knock."}`
	type Stream struct {
		Name string
		Text string
	}
	b := bytes.NewReader([]byte(jsonStream))
	dec := json.NewDecoder(b)

	var stream Stream
	err := dec.Decode(&stream)
	if err != nil {
		t.Fatalf("Decode err: %v", err)
	}
	t.Logf("dec = %v\n", stream)
}

// 测试 json decoder buffered
func TestJsonDecoderBuffered(t *testing.T) {
	const jsonStream = `
			{"Name": "Ed", "Text": "Knock knock."}
			{"Name": "Sam", "Text": "Who's there?"}
			{"Name": "Ed", "Text": "Go fmt."}
			{"Name": "Sam", "Text": "Go fmt who?"}
			{"Name": "Ed", "Text": "Go fmt yourself!"}
		`
	type Stream struct {
		Name string
		Text string
	}
	b := bytes.NewReader([]byte(jsonStream))
	dec := json.NewDecoder(b)

	for {
		var stream Stream
		err := dec.Decode(&stream)
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Decode err: %v", err)
		}
		t.Logf("Decode = %s\n", dec.Buffered())
	}
}

// 测试 struct 中含有 json 中没有的字段时的情况
func TestDecoderDisallow(t *testing.T) {
	const jsonStream = `{"Name": "Ed", "Text": "Knock knock.", "Age": 20}`

	type Stream struct {
		Name string
		Text string
	}
	b := bytes.NewReader([]byte(jsonStream))
	dec := json.NewDecoder(b)
	dec.DisallowUnknownFields()

	var stream Stream
	err := dec.Decode(&stream)
	if err == nil {
		t.Fatalf("Expect err: %v, nil got", err)
	}
}

// 测试 json decoder token
func TestDecoderToken(t *testing.T) {
	const jsonStream = `
	{"Message": "Hello", "Array": [1, 2, 3], "Null": null, "Number": 1.234}
`
	b := bytes.NewReader([]byte(jsonStream))
	dec := json.NewDecoder(b)

	for {
		token, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Token err: %v", err)
		}
		t.Logf("Token: %v", token)
		if dec.More() {
			t.Log("(...)")
		}
	}
}

// 测试使用 useNumber，将数字类型转为 int64, float64, string
func TestDecoderUseNumber(t *testing.T) {
	const jsonStream = `{"name": "ethancai", "fansCount": 9223372036854775807}`
	type User struct {
		Name      string
		FansCount int64
	}
	var user User
	dec := json.NewDecoder(bytes.NewReader([]byte(jsonStream)))
	err := dec.Decode(&user)
	if err != nil {
		t.Fatalf("Decode err: %v", err)
	}
	t.Log(user)

	var user1 interface{}
	dec = json.NewDecoder(bytes.NewReader([]byte(jsonStream)))
	dec.UseNumber()
	err = dec.Decode(&user1)
	if err != nil {
		t.Fatalf("Decode err: %v", err)
	}
	t.Log(user1)

	u := user1.(map[string]interface{})
	// 转为 json.Number 类型
	i, err := u["fansCount"].(json.Number).Int64()
	t.Logf("fansCount = %v, err = %v", i, err)
}
