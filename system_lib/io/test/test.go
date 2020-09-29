package test

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	myio "github.com/oofpgDLD/goSkill/system_lib/io"
	"github.com/oofpgDLD/goSkill/system_lib/io/task"
)

var TaskStore = map[string]*task.Task{
	"wd":           task.NewTask(testWallDoor),
	"std-read":     task.NewTask(testStdReader),
	"std-copy":     task.NewTask(testStdCopy),
	"std-read-all": task.NewTask(testStdReadAll),
	"string-eof":   task.NewTask(testStringsReadEOF),
	"read-file":    task.NewTask(testReadFile),
}

//测试自定义信号发射装置
func testWallDoor(closer chan struct{}) {
	wd := myio.NewWallDoor()
	stream := myio.NewWDStream(wd)
	go func() {
		time.Sleep(time.Second * 50)
		wd.Close()
	}()

	buf := make([]byte, 4)
	for {
		select {
		case <-closer:
			fmt.Println("testWallDoor closing")
			return
		default:
			n, err := stream.Read(buf)
			if err != nil {
				if err == io.EOF {
					fmt.Println("err:", "wd stream closed")
					return
				}
				fmt.Println("err:", err)
				return
			}
			fmt.Println("rev:", string(buf[:n]))
		}
	}
}

func testStdReader(closer chan struct{}) {
	go func() {
		time.Sleep(time.Second * 50)
		os.Stdin.Close()
	}()

	buf := make([]byte, 4)
	for {
		select {
		case <-closer:
			fmt.Println("testStdReader closing")
			return
		default:
			n, err := os.Stdin.Read(buf)
			if err != nil {
				if err == io.EOF {
					fmt.Println("err:", "stdin stream closed")
					return
				}
				fmt.Println("err:", err)
				return
			}
			fmt.Println("rev:", string(buf[:n]))
		}
	}
}

func testStdCopy(closer chan struct{}) {
	go func() {
		time.Sleep(time.Second * 50)
		os.Stdout.Close()
	}()

	for {
		select {
		case <-closer:
			fmt.Println("testCopy closing")
			return
		default:
			w, err := io.Copy(os.Stdout, os.Stdin)
			if err != nil {
				fmt.Println("err", err)
				return
			}
			fmt.Println("copy", w)
		}
	}
}

func testStdReadAll(closer chan struct{}) {
	rd := bufio.NewReader(os.Stdin)
	for {
		select {
		case <-closer:
			fmt.Println("testReadAll closing")
			return
		default:
			b, err := ioutil.ReadAll(rd)
			if err != nil {
				fmt.Println("err:", err)
				return
			}
			fmt.Println("rev:", string(b))
		}
	}
}

func testStringsReadEOF(closer chan struct{}) {
	srd := strings.NewReader("hello world")
	buf := make([]byte, 4)
	for {
		n, err := srd.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("err:", "strings read EOF")
				break
			}
			fmt.Println("err:", err)
			return
		}
		fmt.Println("rev:", string(buf[:n]))
	}
}

func testReadFile(closer chan struct{}) {
	filename := "testfile1.txt"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	fmt.Println("read:", data)
}
