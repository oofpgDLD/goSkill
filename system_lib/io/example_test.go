package io

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_ExampleCopy(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read\n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		t.Fatal(err)
	}

	// Output:
	// some io.Reader stream to be read
}

func Test_ExampleCopyBuffer(t *testing.T) {
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 8)

	// buf is used here...
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		t.Fatal(err)
	}

	// ... reused here also. No need to allocate an extra buffer.
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		t.Fatal(err)
	}

	// Output:
	// first reader
	// second reader
}


func Test_ExampleLimitReader(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := io.LimitReader(r, 100)

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		t.Fatal(err)
	}

	// Output:
	// some
}

func Test_ExampleCopyN(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read")

	if _, err := io.CopyN(os.Stdout, r, 100); err != nil {
		t.Fatal(err)
	}

	// Output:
	// some
}