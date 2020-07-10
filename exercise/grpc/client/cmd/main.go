package main

import (
	"context"
	"flag"
	pb "github.com/oofpgDLD/goSkill/exercise/grpc/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
	"log"
	"os"
)

var(
	tp int
)

func init(){
	flag.IntVar(&tp,"type", 0, "the type of grpc client: 0 pb client, 1 grpc client")
}

func main() {
	flag.Parse()

	if len(os.Args) > 0 {
		flag.Set("type", os.Args[1])
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8091", grpc.WithInsecure(),
		)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()


	encoding.RegisterCodec()

	switch tp {
	case 0:
		c := pb.NewUserSrvClient(conn)

		r, err := c.UserInfo(context.Background(), &pb.UserReq{Id: "1"})
		if err != nil {
			log.Fatalf("could not call: %v", err)
		}
		log.Printf("user: %s", r)
	case 1:
		GetUser(conn)
	case 2:
		GetUserCallOption(conn)
	}
}

func GetUser(c *grpc.ClientConn) {
	reply := &pb.User{}
	err := c.Invoke(context.Background(), "/grpc.UserSrv/UserInfo", &pb.UserReq{Id: "1"}, reply)
	if err != nil {
		log.Fatalf("could not call: %v", err)
	}
	log.Printf("invoke user: %s", reply)
	log.Printf("type of reply is:%T",reply)
}

func GetUserCallOption(c *grpc.ClientConn) {
	opt := grpc.CallContentSubtype("json")
	reply := &pb.User{}
	err := c.Invoke(context.Background(), "/grpc.UserSrv/UserInfo", &pb.UserReq{Id: "1"}, reply, opt)
	if err != nil {
		log.Fatalf("could not call: %v", err)
	}
	log.Printf("invoke user: %s", reply)
	log.Printf("type of reply is:%T",reply)
}