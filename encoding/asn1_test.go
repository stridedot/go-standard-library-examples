package encoding_test

import (
	"encoding/asn1"
	"testing"
)

// Marshal 和 Unmarshal 数据
func TestAsn1Marshal(t *testing.T) {
	type testSetSET []string
	testSet := testSetSET{"a", "aa", "b", "bb", "c", "cc"}

	// Expected ordering of the SET should be:
	// a, b, c, aa, bb, cc

	output, err := asn1.Marshal(testSet)
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("output = %s\n", output)

	// expectedOrder := testSetSET{"a", "b", "c", "aa", "bb", "cc"}
	var resultSet testSetSET
	// var raw asn1.RawValue
	_, err = asn1.Unmarshal(output, &resultSet)
	t.Logf("rest = %v\n", resultSet)
}

// Marshal 为顶级元素指定字段参数
func TestAsn1MarshalWithParam(t *testing.T) {
	type testSetSET []string
	testSet := testSetSET{"a", "aa", "b", "bb", "c", "cc"}
	output, err := asn1.MarshalWithParams(testSet, "private")
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	t.Logf("output = %v\n", output)
}
