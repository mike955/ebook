package main

import (
	pb "ebook/api/user"
	"ebook/conf"
	"ebook/internal/dao"
	"ebook/internal/service"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init()  {
	dao.SetUp()
}

func main()  {
	port := conf.PORT
	lis, err := net.Listen("tcp", port)
	
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, service.UserService)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
	log.Fatalf("server start %s", port)
}