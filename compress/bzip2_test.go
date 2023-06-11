package compress_test

import (
	"compress/bzip2"
	"io"
	"os"
	"testing"
)

func TestBzip2(t *testing.T) {
	f, err := os.Open("testdata/e.txt.bz2")
	if err != nil {
		t.Fatalf("open file failed, err=%v", err)
	}
	r := bzip2.NewReader(f)
	buf, err := io.ReadAll(r)
	t.Log(string(buf), err)
}
