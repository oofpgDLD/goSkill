package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
)

func t1() {
	go func() {
		buf := make([]byte, 100)
		for {
			//buf := make([]byte, 100)
			n, err := os.Stdin.Read(buf)
			if err != nil {
				if err == io.EOF {
					return
				}
				fmt.Println("err:", err)
				return
			}
			fmt.Println("rev:",string(buf[:n]))
		}
	}()
}

func t2() {
	go func() {
		w, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			fmt.Println("err", err)
			return
		}
		fmt.Println("copy", w)
	}()
}

func t3() {
	go func() {
		rd := bufio.NewReader(os.Stdin)
		for {
			//buf := make([]byte, 100)
			_, err := rd.Peek(2)
			if err != nil {
				if err == io.EOF {
					return
				}
				fmt.Println("err:", err)
				return
			}
			b,err :=ioutil.ReadAll(rd)
			if err != nil {
				fmt.Println("err:", err)
				return
			}
			fmt.Println("rev:",string(b))
		}
	}()
}

func main() {
	//t2()
	//t1()
	t3()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Println("user server exit")
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}