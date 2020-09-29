package main

import (
	"context"
	"fmt"
	"github.com/oofpgDLD/goSkill/system_lib/io/test"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func readerTest(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("failed:%v", err)))
		return
	}
	names := req.Form["name"]
	methods := req.Form["method"]
	if len(names) < 1 {
		w.Write([]byte(fmt.Sprintf("failed:%v", "not find [name] param")))
		return
	}
	if len(methods) < 1 {
		w.Write([]byte(fmt.Sprintf("failed:%v", "not find [method] param")))
		return
	}
	name := names[0]
	method := methods[0]
	if t, ok := test.TaskStore[name]; !ok || t == nil {
		w.Write([]byte(fmt.Sprintf("failed:[%v]%v", name, "task not find")))
		return
	} else {
		log.Printf("%v %v", name, method)
		switch method {
		case "run":
			if err := t.Run(); err != nil {
				w.Write([]byte(fmt.Sprintf("failed:[%v]%v", name, err)))
				return
			}
		case "stop":
			if err := t.Stop(); err != nil {
				w.Write([]byte(fmt.Sprintf("failed:[%v]%v", name, err)))
				return
			}
		default:
			w.Write([]byte(fmt.Sprintf("failed:[%v]%v", method, "method not find")))
			return
		}
	}
	w.Write([]byte("success"))
}

func main() {
	l, err := net.Listen("tcp", ":28000")
	if err != nil {
		log.Fatalf("listen: %s\n", err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/reader-test", readerTest)

	srv := http.Server{
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
	}

	go func() {
		log.Printf("http server listen %v", l.Addr())
		if err := srv.Serve(l); err != nil && err != http.ErrServerClosed {
			log.Fatalf("serve: %s\n", err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Println("user server exit")
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("Server Shutdown:", err)
			}
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
