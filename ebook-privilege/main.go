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
	port := conf.GRPC_PRIVILEGE_PORT
	if os.Getenv("port") != "" {
		port = fmt.Sprintf(":%s", os.Getenv("port"))
	}
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterPrivilegeServer(server, service.PrivilegeService)
	fmt.Println("starting ebook-privilege gRPC server: ", port)
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