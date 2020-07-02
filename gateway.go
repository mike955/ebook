package main

//var (
//	grpcServerEndpoint = &conf.PORT
//)

//func run() error {
//ctx := context.Background()
//ctx, cancel := context.WithCancel(ctx)
//defer cancel()
//
//// Register gRPC server endpoint
//// Note: Make sure the gRPC server is running properly and accessible
//mux := runtime.NewServeMux()
//opts := []grpc.DialOption{grpc.WithInsecure()}
//err := pb.RegisterUserHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
//if err != nil {
//return err
//}
//
//// Start HTTP server (and proxy calls to gRPC server endpoint)
//return http.ListenAndServe(":50802", mux)
//}

func main()  {
	//if err := run(); err != nil {
	//	glog.Fatal(err)
	//}
}