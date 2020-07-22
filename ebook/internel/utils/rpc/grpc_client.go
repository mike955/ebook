package rpc

import (
	"context"
	privilege_pb "ebook/ebook/api/privilege"
	user_pb "ebook/ebook/api/user"
	"ebook/ebook/conf"
	"google.golang.org/grpc"
	"log"
)

var _gRPCClientMap = map[string]interface{}{}

type Rpc struct {
	UserRpc *grpc.ClientConn
	PrivilegeRpc *privilege_pb.PrivilegeClient
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

