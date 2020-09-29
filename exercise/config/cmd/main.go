package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/inconshreveable/log15"
	"github.com/oofpgDLD/goSkill/exercise/config"
)

var l = log15.New()

var (
	conf = config.Config{}
)

func main() {
	cfgPath, err := findConfiger()
	if err != nil {
		l.Error("config path not find:", "err", err.Error())
		os.Exit(1)
	}
	_, err = toml.DecodeFile(cfgPath, &conf)
	if err != nil {
		fmt.Println("read config file failed:", "err", err.Error())
		os.Exit(1)
	}

	l.Info("config is", "config", conf)
}

func findConfiger() (string, error) {
	var configPath = ""
	l.Info("runtime:", "os", runtime.GOOS)
	if runtime.GOOS == `windows` {
		configPath = "etc/config.toml"
	} else {
		err := os.Chdir(pwd())
		if err != nil {
			l.Info("get project pwd err", "err", err)
			return configPath, err
		}
		d, _ := os.Getwd()
		l.Info("project info:", "dir", d)
		configPath = d + "/etc/config.toml"
	}
	return configPath, nil
}

/*
	---workdir/
		| -- bin/
		|     |-- chat(I am here)
		|
		| -- etc/
			  |-- config.toml
			  |-- config.json
*/
func pwd() string {
	dir, err := filepath.Abs(filepath.Dir(filepath.Dir(os.Args[0])))
	if err != nil {
		panic(err)
	}
	return dir
}
