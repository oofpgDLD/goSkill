package main

import (
	"encoding/json"
	"fmt"
	"github.com/oofpgDLD/goSkill/config-center/etcd"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var	e etcd.EtcdM

	data := &struct {
		Name string
	}{
		Name: "dongge",
	}

	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	e.Add("test", b)
}

func serve() {


	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("up-service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			time.Sleep(time.Second * 2)
			log.Info("user server exit")
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}