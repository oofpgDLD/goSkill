package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/grpc"
	registry2 "github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/uber/jaeger-client-go"
	"time"

	"github.com/uber/jaeger-client-go/config"

	proto "github.com/oofpgDLD/goSkill/exercise/micro-cli"
)

func main() {
	Client2()
}
func Client2() {
	cfg := config.Configuration{
		ServiceName: "enterprise.client",//自定义服务名称
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

	registry := etcdv3.NewRegistry(
		registry2.Addrs("172.16.103.31:2379"),
		etcdv3.Auth("root", "admin"),
	) //a default to using env vars for master API

	c := grpc.NewClient(
		client.Registry(registry),
		client.Wrap(opentracing.NewClientWrapper(tracer)),
	)

	// Create new greeter client
	src := proto.NewEnterpriseService("enterprise", c)

	ret, err := src.Auth(context.TODO(), &proto.AuthRequest{AppId:"1001",Token:"session-login=MTU4OTc5MzA2NXxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRUJfNElBQmdaemRISnBibWNNQ1FBSFpHVjJkSGx3WlFaemRISnBibWNNQ1FBSFFXNWtjbTlwWkFaemRISnBibWNNQmdBRWRYVnBaQVp6ZEhKcGJtY01JZ0FnTmtSR05USXpOak5HTmtORlEwSkRNRVEwTlRjNU5EUXlPRE15TlVFM016RUdjM1J5YVc1bkRBY0FCV0Z3Y0Vsa0JuTjBjbWx1Wnd3R0FBUXhNREF4Qm5OMGNtbHVad3dHQUFSMGFXMWxCV2x1ZERZMEJBZ0EtZ0xrVGhvWFJnWnpkSEpwYm1jTUNRQUhkWE5sY2w5cFpBWnpkSEpwYm1jTUF3QUJOQVp6ZEhKcGJtY01Cd0FGZEc5clpXNEdjM1J5YVc1bkRDb0FLR1EyTVdZek5EVmhPVFJpWWprME1HSTBORGMwTlRVd01qY3hPVEExWlRVd1pHTTFNR1psTWpVPXwcqpFI8g8aQmOJzn0QjMJips-2pD37C4rWTJ0DMpJSoQ=="})
	//ret, err := src.Auth(context.TODO(), &proto.AuthRequest{AppId:"1001",Token:"session-login=MTU4OTc5MzA2NXxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRUJfNElBQmdaemRISnBibWNNQ1FBSFpHVjJkSGx3WlFaemRISnBibWNNQ1FBSFFXNWtjbTlwWkFaemRISnBibWNNQmdBRWRYVnBaQVp6ZEhKcGJtY01JZ0FnTmtSR05USXpOak5HTmtORlEwSkRNRVEwTlRjNU5EUXlPRE15TlVFM016RUdjM1J5YVc1bkRBY0FCV0Z3Y0Vsa0JuTjBjbWx1Wnd3R0FBUXhNREF4Qm5OMGNtbHVad3dHQUFSMGFXMWxCV2x1ZERZMEJBZ0EtZ0xrVGhvWFJnWnpkSEpwYm1jTUNRQUhkWE5sY2w5cFpBWnpkSEpwYm1jTUF3QUJOQVp6ZEhKcGJtY01Cd0FGZEc5clpXNEdjM1J5YVc1bkRDb0FLR1EyTVdZek5EVmhPVFJpWWprME1HSTBORGMwTlRVd0"})
	if err != nil {
		fmt.Println("call auth failed", err)
		return
	}

	if ret != nil {
		fmt.Println(ret.AppId)
		fmt.Println(ret.UserId)
		fmt.Println(ret.Token)
		fmt.Println(ret.Uid)
		fmt.Println(ret.EndTime)
	}
	fmt.Println("call auth success")
}

func Client1() {
	cfg := config.Configuration{
		ServiceName: "enterprise.client",//自定义服务名称
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

	registry := etcdv3.NewRegistry(
		registry2.Addrs("172.16.103.31:2379"),
		etcdv3.Auth("root", "admin"),
	) //a default to using env vars for master API

	service := micro.NewService(
		// Set service name
		micro.Name("enterprise.client"),
		// Set service registry
		micro.Registry(registry),
		//
		micro.WrapClient(opentracing.NewClientWrapper(tracer)),
	)

	service.Init()

	// Create new greeter client
	src := proto.NewEnterpriseService("enterprise", service.Client())

	//ret, err := src.Auth(context.TODO(), &proto.AuthRequest{AppId:"1001",Token:"session-login=MTU4OTc5MzA2NXxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRUJfNElBQmdaemRISnBibWNNQ1FBSFpHVjJkSGx3WlFaemRISnBibWNNQ1FBSFFXNWtjbTlwWkFaemRISnBibWNNQmdBRWRYVnBaQVp6ZEhKcGJtY01JZ0FnTmtSR05USXpOak5HTmtORlEwSkRNRVEwTlRjNU5EUXlPRE15TlVFM016RUdjM1J5YVc1bkRBY0FCV0Z3Y0Vsa0JuTjBjbWx1Wnd3R0FBUXhNREF4Qm5OMGNtbHVad3dHQUFSMGFXMWxCV2x1ZERZMEJBZ0EtZ0xrVGhvWFJnWnpkSEpwYm1jTUNRQUhkWE5sY2w5cFpBWnpkSEpwYm1jTUF3QUJOQVp6ZEhKcGJtY01Cd0FGZEc5clpXNEdjM1J5YVc1bkRDb0FLR1EyTVdZek5EVmhPVFJpWWprME1HSTBORGMwTlRVd01qY3hPVEExWlRVd1pHTTFNR1psTWpVPXwcqpFI8g8aQmOJzn0QjMJips-2pD37C4rWTJ0DMpJSoQ=="})
	ret, err := src.Auth(context.TODO(), &proto.AuthRequest{AppId:"1001",Token:"session-login=MTU4OTc5MzA2NXxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRUJfNElBQmdaemRISnBibWNNQ1FBSFpHVjJkSGx3WlFaemRISnBibWNNQ1FBSFFXNWtjbTlwWkFaemRISnBibWNNQmdBRWRYVnBaQVp6ZEhKcGJtY01JZ0FnTmtSR05USXpOak5HTmtORlEwSkRNRVEwTlRjNU5EUXlPRE15TlVFM016RUdjM1J5YVc1bkRBY0FCV0Z3Y0Vsa0JuTjBjbWx1Wnd3R0FBUXhNREF4Qm5OMGNtbHVad3dHQUFSMGFXMWxCV2x1ZERZMEJBZ0EtZ0xrVGhvWFJnWnpkSEpwYm1jTUNRQUhkWE5sY2w5cFpBWnpkSEpwYm1jTUF3QUJOQVp6ZEhKcGJtY01Cd0FGZEc5clpXNEdjM1J5YVc1bkRDb0FLR1EyTVdZek5EVmhPVFJpWWprME1HSTBORGMwTlRVd0"})
	if err != nil {
		fmt.Println("call auth failed", err)
		return
	}

	if ret != nil {
		fmt.Println(ret.AppId)
		fmt.Println(ret.UserId)
		fmt.Println(ret.Token)
		fmt.Println(ret.Uid)
		fmt.Println(ret.EndTime)
	}
	fmt.Println("call auth success")
}