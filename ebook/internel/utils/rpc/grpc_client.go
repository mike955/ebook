package rpc

import (
	"context"
	privilege_pb "ebook/ebook/api/privilege"
	user_pb "ebook/ebook/api/user"
	ebook_pb "ebook/ebook/api/ebook"
	"ebook/ebook/conf"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

var _gRPCClientMap = map[string]interface{}{}

type Rpc struct {
	UserRpc *user_pb.UserClient
	PrivilegeRpc *privilege_pb.PrivilegeClient
	EbookRpc *ebook_pb.EbookClient
}

type userRpc struct {

}

type privilegeRpc struct {

}

func UserRpc() user_pb.UserClient  {
	if _gRPCClientMap["user"] == nil{
		var ctx = context.Background()
		conn, err := grpc.DialContext(ctx, conf.GRPC_ADDR_MAP["user"], grpc.WithInsecure()); if err != nil {
			log.Fatalf("user grpc client did not connect: %v", err)
		}
		client := user_pb.NewUserClient(conn)
		_gRPCClientMap["user"] = client
	}
	return _gRPCClientMap["user"].(user_pb.UserClient)
}

func PrivilegeRpc() privilege_pb.PrivilegeClient {
	if _gRPCClientMap["privilege"] == nil{
		var ctx = context.Background()
		conn, err := grpc.DialContext(ctx, conf.GRPC_ADDR_MAP["privilege"], grpc.WithInsecure());	if err != nil {
			log.Fatalf("privilege grpc client did not connect: %v", err)
		}
		client := privilege_pb.NewPrivilegeClient(conn)
		_gRPCClientMap["privilege"] = client
	}
	return _gRPCClientMap["privilege"].(privilege_pb.PrivilegeClient)
}

func EbookRpc() ebook_pb.EbookClient {
	if _gRPCClientMap["ebook"] == nil{
		var ctx = context.Background()
		conn, err := grpc.DialContext(ctx, conf.GRPC_ADDR_MAP["ebook"], grpc.WithInsecure());	if err != nil {
			fmt.Println(err)
			log.Fatalf("ebook grpc client did not connect: %v", err)
		}
		client := ebook_pb.NewEbookClient(conn)
		_gRPCClientMap["ebook"] = client
	}
	return _gRPCClientMap["ebook"].(ebook_pb.EbookClient)
}
