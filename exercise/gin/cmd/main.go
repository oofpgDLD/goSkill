package main

import (
	"github.com/oofpgDLD/goSkill/exercise/gin/api"
	"net/http"
	"time"

	"github.com/inconshreveable/log15"
)

var log = log15.New()

func main() {
	address := "localhost:1112"

	router := api.New()
	srv := http.Server{
		Addr: address,
		Handler: router,
		ReadTimeout: 10 *time.Second,
		WriteTimeout: 10 *time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("after http server", "err", err.Error())
	}
}
