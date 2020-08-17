package main

import (
	pb "ebook/ebook-ebook/api/ebook"
	"ebook/ebook-ebook/conf"
	"ebook/ebook-ebook/internel/dao"
	"ebook/ebook-ebook/internel/service"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	dao.SetUp()
}

func main() {
	port := conf.GRPC_USER_PORT
	if os.Getenv("port") != "" {
		port = fmt.Sprintf(":%s", os.Getenv("port"))
	}
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterEbookServer(server, service.EbookService)
	fmt.Println("starting ebook-ebook gRPC server: ", port)
	go func() {
		if err := server.Serve(listen); err != nil {
			fmt.Printf("failed to serve: %v", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGKILL)
	<-sig
	fmt.Println("Get Signal:", sig)
	fmt.Println("Shutdown Server ...")
	log.Println("Server exiting")
	server.GracefulStop()
}
