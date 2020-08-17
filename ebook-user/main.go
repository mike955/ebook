package main

import (
	pb "ebook/ebook-user/api/user"
	"ebook/ebook-user/conf"
	"ebook/ebook-user/internel/dao"
	"ebook/ebook-user/internel/service"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func init()  {
	dao.SetUp()
}

func main()  {
	
	port := conf.GRPC_USER_PORT
	if os.Getenv("port") != "" {
		port = fmt.Sprintf(":%s", os.Getenv("port"))
	}
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterUserServer(server, service.UserService)
	fmt.Println("starting ebook-user gRPC server: ", port)
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