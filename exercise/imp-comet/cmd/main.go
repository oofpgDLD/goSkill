package main

import (
	"flag"
	"github.com/oofpgDLD/goSkill/exercise/imp-comet/conf"
	"github.com/oofpgDLD/goSkill/exercise/imp-comet/server"
	"github.com/oofpgDLD/goSkill/exercise/imp-comet/server/grpc"
	"github.com/oofpgDLD/goSkill/exercise/imp-comet/service"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	log "github.com/golang/glog"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	println(conf.Conf.Debug)
	log.Infof("goim-comet [version: %s env: %+v] start", ver, conf.Conf.Env)


	// new comet server
	srv := service.New(conf.Conf)

	if err := server.InitTCP(srv); err != nil {
		panic(err)
	}

	go grpc.New(conf.Conf, srv)
	// signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("goim-comet get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			//rpcSrv.GracefulStop()
			srv.Close()
			log.Infof("goim-comet [version: %s] exit", ver)
			log.Flush()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}