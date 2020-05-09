package main

import (
	"context"
	pb "github.com/oofpgDLD/goSkill/exercise/grpc/client/api"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8091", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserSrvClient(conn)

	r, err := c.UserInfo(context.Background(), &pb.UserReq{Id: "1"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r)
}