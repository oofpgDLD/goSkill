package main

import (
	"github.com/inconshreveable/log15"
	"os"
	"os/signal"
	"syscall"
)

var log = log15.New()

func main() {
	log.Info("server started")
	c := make(chan os.Signal)
	signal.Notify(c )
	for {
		s := <-c
		log.Info("get signal", "signal", s)
		switch s {
		case syscall.SIGQUIT:
			log.Info("do quit")
		case syscall.SIGHUP:
			log.Info("do hup")
		case syscall.SIGTERM:
			log.Info("do term")
		case syscall.SIGINT:
			log.Info("do int")
		case syscall.SIGTSTP:
			log.Info("do stop")
		case syscall.SIGKILL:
			log.Info("do kill")
		case syscall.SIGSTOP:
			log.Info("do stop2")
		default:
			return
		}
	}
	log.Info("server stop")
}
