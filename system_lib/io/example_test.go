package io

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_ExampleLimitReader(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := io.LimitReader(r, 100)

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		t.Fatal(err)
	}

	// Output:
	// some
}

func Test_ReadAll(t *testing.T) {
	srd := strings.NewReader("hello world")
	data, err := ioutil.ReadAll(srd)
	if err != nil{
		t.Error(err)
		return
	}
	t.Log(string(data))
}

func Test_NewBuffer(t *testing.B) {
}