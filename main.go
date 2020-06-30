package main

import (
	"ebook/api/user"
	pb "ebook/api/user"
	"ebook/conf"
	"ebook/internal/service"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	user.UnimplementedUserServer
}

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService Hello服务
var HelloService = helloService{}

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