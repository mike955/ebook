package main

import (
	"context"
	pb_user "ebook/api/user"
	pb_ebook "ebook/api/ebook"
	"ebook/conf"
	"ebook/internal/dao"
	"ebook/internal/service"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
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
	pb_user.RegisterUserServer(grpcServer, service.UserService)
	pb_ebook.RegisterEbookServer(grpcServer, service.EbookService)
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
	
	// // handle upload file
	// mx := http.NewServeMux()
	// mx.HandleFunc("/ebook/add", func(w http.ResponseWriter, req *http.Request) {
	// 	return service.EbookService.Add(req)
	// })
	
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb_user.RegisterUserHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	err = pb_ebook.RegisterEbookHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Println("starting http server: ", conf.HTTP_PORT)
	return http.ListenAndServe(conf.HTTP_PORT, mux)
}