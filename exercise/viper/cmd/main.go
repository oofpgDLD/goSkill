package main

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"log"
	"os"
)

type Config struct {
	Server       *HttpServer
	Chat33Server *Chat33Server
}

type Chat33Server struct {
	Host string
}

type HttpServer struct {
	Addr string
}

func main() {
	address := "http://172.16.103.31:2379"
	store := "etcd"
	path := "config"
	viper.AddRemoteProvider(store, address, path)
	//viper.SetConfigFile("config2")
	viper.SetConfigType("toml")
	//viper.SetConfigType("toml")
	//viper.AddConfigPath("etc/")               // 还可以在工作目录中查找配置

	//viper.WriteConfig()

	var cfg Config
	//从远程配置中心读取
	if err := viper.ReadRemoteConfig(); err != nil{
		//
		log.Println("read from remote")
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			log.Fatal("未找到配置文件", err.Error())
			os.Exit(1)
		} else {
			// 配置文件被找到，但产生了另外的错误
			log.Fatal("配置文件格式错误", err.Error())
			os.Exit(1)
		}

		//从本地读取
		if err := viper.ReadInConfig(); err != nil {
			log.Println("read from local")
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				// 配置文件未找到错误；如果需要可以忽略
				log.Fatal("未找到配置文件", err.Error())
				os.Exit(1)
			} else {
				// 配置文件被找到，但产生了另外的错误
				log.Fatal("配置文件格式错误", err.Error())
				os.Exit(1)
			}
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			log.Fatal("未找到配置文件")
			os.Exit(1)
		} else {
			// 配置文件被找到，但产生了另外的错误
			log.Fatal("配置文件格式错误")
			os.Exit(1)
		}
	}

	log.Println("成功", cfg.Server, cfg.Chat33Server)
	os.Exit(0)
}
