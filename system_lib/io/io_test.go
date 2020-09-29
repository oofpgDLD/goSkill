package io

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func Test_IOWrite(t *testing.T) {
	ps := []string{
		"hello",
		"world",
	}

	writer := bytes.Buffer{}
	for _,p := range ps{
		n, err := writer.Write([]byte(p))
		if err != nil {
			t.Error(err)
			return
		}
		if n != len(p) {
			t.Error(err)
			return
		}
	}
	t.Log(writer.String())
}

func Test_BufioReadLine(t *testing.T) {
	srd := strings.NewReader("hhhhhhhh")

	bufr := bufio.NewReader(srd)
	t.Log(bufr.ReadLine())
}