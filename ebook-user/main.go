package main

import (
	pb "ebook/ebook-user/api/user"
	"ebook/ebook-user/conf"
	"ebook/ebook-user/internel/dao"
	"ebook/ebook-user/internel/service"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func init()  {
	dao.SetUp()
}

func main()  {
	listen, err := net.Listen("tcp", conf.GRPC_USER_PORT)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterUserServer(server, service.UserService)
	fmt.Println("starting gRPC server: ", conf.GRPC_USER_PORT)
	if err := server.Serve(listen); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}