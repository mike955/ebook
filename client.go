package main

import (
	"context"
	pb "ebook/api/user"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address     = "localhost:50801"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserClient(conn)
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SignUp(ctx, &pb.SignUpRequest{AccountName: "mike"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("data: %s", r.Data)
}