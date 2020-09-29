package main

import (
	"context"
	"github.com/inconshreveable/log15"
	"github.com/oofpgDLD/goSkill/exercise/http/server/api"
	"net/http"
	"os"
	"time"
	"os/signal"
	"syscall"
)

var log = log15.New()

func main() {
	address := "localhost:1111"
	router := api.New()
	srv := http.Server{
		Addr: address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil{
			log.Error("after http server", "err", err.Error())
		}
	}()

	log.Info("server started")
	c := make(chan os.Signal)
	signal.Notify(c)
	for {
		s := <-c
		log.Info("get signal", "signal", s)
		switch s {
		case syscall.SIGHUP:
			log.Info("do restart http server")
			go func() {
				if err := srv.ListenAndServe(); err != nil{
					log.Error("after http server", "err", err.Error())
				}
			}()
		case syscall.SIGQUIT:
			log.Info("do shutdown http server")
			ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
			if err := srv.Shutdown(ctx); err != nil {
				log.Error("Server Shutdown", "err", err)
			}
		default:
			goto exit
		}
	}
	exit:
	log.Info("server stop")
}


