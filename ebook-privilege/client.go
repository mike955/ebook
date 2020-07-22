package main

import (
	"context"
	pb "ebook/ebook-privilege/api/privilege"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address     = "127.0.0.1:50802"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPrivilegeClient(conn)
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddUser(ctx, &pb.AddUserRequest{UserId: "mike" , RoleId: uint32(122)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("data: %s", r.Data)
}