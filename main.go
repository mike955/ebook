package main

import (
	"ebook/api/user"
	"ebook/conf"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	user.UnimplementedUserServer
}

func main()  {
	port := conf.PORT
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	user.RegisterUserServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}