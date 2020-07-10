package mock

import pb "github.com/oofpgDLD/goSkill/exercise/grpc/api"

func GetUser() *pb.User{
	return &pb.User{
		Id: 1,
		Name: "mock",
	}
}