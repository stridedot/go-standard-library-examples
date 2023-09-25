package testing

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func Greet(w io.Writer, name string) {
	_, _ = fmt.Fprintf(w, "Hello, %s", name)
}

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
