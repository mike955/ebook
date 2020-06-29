package service

import (
	"context"
	pb "ebook/api/user"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
)

type UserService struct {
	data data.User
	dao dao.User
}


func (s *UserService) SignUp(ctx context.Context, req *pb.SignUpRequest) (reply *empty.Empty, err error) {
	reply = new(empty.Empty)
	fmt.Printf("hello %s", req.AccountName)
	//accountName := req.AccountName
	//accountInfo :=
	// 判断用户名是否存在

	// 判断邮箱是否存在

	// 判断用户角色

	// 生成盐、id

	// 加密密码

	// 存储数据

	// 返回结果
	return
}