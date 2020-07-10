package main

import (
	"context"
	"fmt"
	pb "github.com/oofpgDLD/goSkill/exercise/grpc/api"
	"github.com/oofpgDLD/goSkill/exercise/grpc/server/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8091")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		os.Exit(1)
	}

	srv := grpc.NewServer(

		)
	pb.RegisterUserSrvServer(srv, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(srv)

	log.Println("server start")
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalf("grpc serve shutdown: %v", err)
		os.Exit(1)
	}
}

type server struct {
}

func (s *server) UserInfo(ctx context.Context, in *pb.UserReq) (*pb.User, error){
	//get method
	md,b := metadata.FromIncomingContext(ctx)
	fmt.Println(b," metadata is:", md)
	name,_ := grpc.Method(ctx)
	log.Println("method is:" + name)
	return mock.GetUser(), nil
}