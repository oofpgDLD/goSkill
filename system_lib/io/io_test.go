package io

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"testing"
	"time"
)

func Test_IOReadEOF(t *testing.T) {
	srd := strings.NewReader("hello world")
	buf := make([]byte, 4)
	for {
		n, err := srd.Read(buf)
		if err != nil {
			if err == io.EOF {
				t.Log("EOF:",n)
				break
			}
			t.Error(err)
			return
		}
		t.Log(string(buf[:n]))
	}
}

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

func Test_IOReadStd(t *testing.T) {
	in := os.Stdin
	defer in.Close()
	go func() {
		for {
			buf := make([]byte, 100)
			n, err := in.Read(buf)
			if err != nil {
				if err == io.EOF {
					return
				}
			}
			fmt.Println(string(buf[:n]))
			time.Sleep(1 * time.Second)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			t.Log("user server exit")
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}