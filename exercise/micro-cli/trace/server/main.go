package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	registry2 "github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/uber/jaeger-client-go"
	"time"

	proto "github.com/oofpgDLD/goSkill/exercise/micro-cli"

	"github.com/uber/jaeger-client-go/config"
)

/*func NewClientWrapper(ot opentracing.Tracer) client.Wrapper {
	return func(c client.Client) client.Client {
		return &otWrapper{ot, c}
	}
}
*/

func main() {
	cfg := config.Configuration{
		ServiceName: "enterprise",//自定义服务名称
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  "172.16.103.31:5775",//jaeger agent
		},
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer closer.Close()

	//etcdv3插件
	registry := etcdv3.NewRegistry(
		registry2.Addrs("172.16.103.31:2379"), //etch 服务器地址
		etcdv3.Auth("root", "admin"),
	)

	service := micro.NewService(
		// Set service registry
		micro.Registry(registry),
		micro.Name("enterprise"),
		micro.Version("1.0.0"),
		micro.WrapHandler(opentracing.NewHandlerWrapper(tracer)),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	err = proto.RegisterEnterpriseHandler(service.Server(), new(server))
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