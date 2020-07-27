package main

import (
	pb "ebook/ebook-ebook/api/ebook"
	"ebook/ebook-ebook/conf"
	"ebook/ebook-ebook/internel/dao"
	"ebook/ebook-ebook/internel/service"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

// protoc --gofast_out=plugins=grpc:ebook ebook.proto

func init()  {
	dao.SetUp()
}

func main()  {
	listen, err := net.Listen("tcp", conf.GRPC_EBOOK_PORT)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterEbookServer(server, service.EbookService)
	fmt.Println("starting ebook-user gRPC server: ", conf.GRPC_EBOOK_PORT)
	if err := server.Serve(listen); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}