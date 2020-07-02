package main

import (
	"context"
	pb "ebook/api/user"
	"ebook/conf"
	"ebook/internal/dao"
	"ebook/internal/service"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func init()  {
	dao.SetUp()
}

var (
	grpcServerEndpoint = &conf.GRPC_PORT
)

func main()  {
	listen, err := net.Listen("tcp", conf.GRPC_PORT)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServer(grpcServer, service.UserService)
	go func() {
		// start gRPC server
		log.Println("starting gRPC server: ", conf.GRPC_PORT)
		grpcServer.Serve(listen)
	}()
	run()
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterUserHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Println("starting http server: ", conf.HTTP_PORT)
	return http.ListenAndServe(conf.HTTP_PORT, mux)
}