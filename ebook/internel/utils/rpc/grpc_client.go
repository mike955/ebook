package rpc

import (
	"context"
	privilege_pb "ebook/ebook/api/privilege"
	"ebook/ebook/conf"
	"google.golang.org/grpc"
	"log"
	"time"
)

var _map map[]string]interface{}

func createPrivilegeGRPCClient() privilege_pb.PrivilegeClient  {
	conn, err := grpc.Dial(conf.GRPC_ADDR_MAP["key"], grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := privilege_pb.NewPrivilegeClient(conn)
	return client
}

func createUserGRPCClient(key string) privilege_pb.PrivilegeClient  {
	conn, err := grpc.Dial(conf.GRPC_ADDR_MAP["key"], grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//if (key == "privilege") {
	client := privilege_pb.NewPrivilegeClient(conn)
	//}
	return client
}

func createGRPCClient(key string) privilege_pb.PrivilegeClient  {
	conn, err := grpc.Dial(conf.GRPC_ADDR_MAP["key"], grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//if (key == "privilege") {
	client := privilege_pb.NewPrivilegeClient(conn)
	//}
	return client
}

func Rpc(key string, fname string, params interface{})  {
	if(_map[key] == nil) {
		client := createGRPCClient(key)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := client.AddUser()
	}
}

func PrivilegeRpc() privilege_pb.PrivilegeClient {
	var client privilege_pb.PrivilegeClient
	if _map["privilege"] == nil {
		client = createPrivilegeGRPCClient()
		_map["privilege"] = client
	}
	return client
}

func UserRpc(){

}