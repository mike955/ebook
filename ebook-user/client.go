package main

import (
	"context"
	pb "ebook/ebook-user/api/user"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	userAddress     = "127.0.0.1:50801"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(userAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserClient(conn)
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//r, err := c.Get(ctx, &pb.GetRequest{UserId: "1234567" , Username: "lll"})
	r, err := c.Add(ctx, &pb.AddRequest{
		Username:             "mike",
		Email:                "mike955@163.com",
		Password:             "d56381bffae08ab3ee6297aedf3474e9",
		RoleId:               uint64(2),
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("data: %s", r)
}