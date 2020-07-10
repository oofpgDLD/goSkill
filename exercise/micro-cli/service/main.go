package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	registry2 "github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"

	proto "github.com/oofpgDLD/goSkill/exercise/micro-cli"
)

func main() {
	//etcdv3插件
	registry := etcdv3.NewRegistry(
		registry2.Addrs("172.16.103.31:2379"), //etch 服务器地址
		etcdv3.Auth("root", "admin"),
	)

	service := micro.NewService(
		// Set service registry
		micro.Registry(registry),
		micro.Name("enterprise"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	err := proto.RegisterEnterpriseHandler(service.Server(), new(server))
	if err != nil {
		fmt.Println(err)
	}
	// Run the server
	if err = service.Run(); err != nil {
		fmt.Println(err)
	}
}

type server struct {
	//srv *service.Service
}

func (s *server) Auth(ctx context.Context, req *proto.AuthRequest, rsp *proto.AuthResponse) error {
	fmt.Println("auth call", req.Token, req.AppId)
	rsp = &proto.AuthResponse{

	}
	return nil
}