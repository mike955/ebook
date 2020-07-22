package main

import (
	pb "ebook/ebook-privilege/api/privilege"
	"ebook/ebook-privilege/conf"
	"ebook/ebook-privilege/internel/dao"
	"ebook/ebook-privilege/internel/service"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func init()  {
	dao.SetUp()
}

func main()  {
	listen, err := net.Listen("tcp", conf.GRPC_PRIVILEGE_PORT)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterPrivilegeServer(server, service.PrivilegeService)
	fmt.Println("starting ebook-privilege gRPC server: ", conf.GRPC_PRIVILEGE_PORT)
	if err := server.Serve(listen); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}