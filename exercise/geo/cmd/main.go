package main

import (
	"fmt"
	"github.com/oofpgDLD/goSkill/exercise/geo"
	"os"
)

func main() {
	cfg := &geo.NodeConfig{
		Url:         "redis://172.16.103.31:6380",
		Password:    "",
		MaxIdle:     30,
		MaxActive:   50,
		IdleTimeout: 240,
	}
	ca := geo.NewLocRedis(geo.NewPool(cfg))

	//err := ca.GeoAdd("1001", "start", "120.12333333333333","30.273888888888887")
	err := ca.GeoAdd("1001", "start", "120.123333","30.273888")
	if err != nil {
		fmt.Println("GeoAdd", err)
		os.Exit(1)
	}
	//err = ca.GeoAdd("1001", "end", "120.14305555555555","30.279166666666665")
	err = ca.GeoAdd("1001", "end", "120.143055","30.279166")
	if err != nil {
		fmt.Println("GeoAdd",err)
		os.Exit(1)
	}

	l ,err := ca.GeoDIST("1001", "start", "end")
	if err != nil {
		fmt.Println("GeoDIST",err)
		os.Exit(1)
	}

	fmt.Println("success:", l)
}