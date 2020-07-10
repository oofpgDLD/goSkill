package server

import (
	"google.golang.org/grpc"
	"net"
)

type ServerConfig struct {
	// Network is grpc listen network,default value is tcp
	Network string `dsn:"network"`
	// Addr is grpc listen addr,default value is 0.0.0.0:9000
	Addr string `dsn:"address"`
}


func NewServer() *Server{
	//TODO
	return &Server{}
}

type Server struct {
	conf *ServerConfig
	server *grpc.Server
}

func (s *Server) Start() error{
	lis, err := net.Listen(s.conf.Network, s.conf.Addr)
	if err != nil {
		return err
	}

	if err := s.server.Serve(lis); err != nil {

	}
	return nil
}

