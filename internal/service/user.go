package service

import (
	"context"
	pb "ebook/api/user"
	"ebook/internal/data"
	"fmt"
)

type userService struct {
	UserData data.UserData
	//dao dao.User
}

var UserService = &userService{}

func (service *userService) SignUp(ctx context.Context, req *pb.SignUpRequest) (response *pb.SignUpResponse, err error){
	
	// check request
	fmt.Println("hello ", req.AccountName)
	return service.UserData.SignUp(req)
}
func (service *userService) SignIn(ctx context.Context, req *pb.SignInRequest) (response *pb.SignInResponse,err error){
	return
}
func (service *userService) SignOut(ctx context.Context, req *pb.SignOutRequest) (response *pb.SignOutResponse,err error){
	return
}
func (service *userService) GetUser(ctx context.Context, req *pb.GetUserRequest) (response *pb.GetUserResponse,err error){
	return
}
func (service *userService) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (response *pb.GetUsersResponse,err error){
	return
}

//func (s *userService) SignUp(ctx context.Context, req *pb.SignUpRequest) (response *pb.SignUpResponse, err error) {
//	//reply = new(empty.Empty)
//	//fmt.Printf("hello %s", req.AccountName)
//	//accountName := req.AccountName
//	//accountInfo :=
//	// 判断用户名是否存在
//
//	// 判断邮箱是否存在
//
//	// 判断用户角色
//
//	// 生成盐、id
//
//	// 加密密码
//
//	// 存储数据
//
//	// 返回结果
//	return
//}
//
//
//func (s *userService) SignIn(ctx context.Context, req *pb.SignInRequest) (response *pb.SignInResponse, err error) {
//	//reply = new(empty.Empty)
//	//fmt.Printf("hello %s", req.AccountName)
//	//accountName := req.AccountName
//	//accountInfo :=
//	// 判断用户名是否存在
//
//	// 判断邮箱是否存在
//
//	// 判断用户角色
//
//	// 生成盐、id
//
//	// 加密密码
//
//	// 存储数据
//
//	// 返回结果
//	return
//}
//
//
//func (s *userService) SignOut(ctx context.Context, req *pb.SignOutRequest) (response *pb.SignOutRequest, err error) {
//	//reply = new(empty.Empty)
//	//fmt.Printf("hello %s", req.AccountName)
//	//accountName := req.AccountName
//	//accountInfo :=
//	// 判断用户名是否存在
//
//	// 判断邮箱是否存在
//
//	// 判断用户角色
//
//	// 生成盐、id
//
//	// 加密密码
//
//	// 存储数据
//
//	// 返回结果
//	return
//}
//
//
//func (s *userService) GetUser(ctx context.Context, req *pb.GetUserRequest) (response *pb.GetUserResponse, err error) {
//	//reply = new(empty.Empty)
//	//fmt.Printf("hello %s", req.AccountName)
//	//accountName := req.AccountName
//	//accountInfo :=
//	// 判断用户名是否存在
//
//	// 判断邮箱是否存在
//
//	// 判断用户角色
//
//	// 生成盐、id
//
//	// 加密密码
//
//	// 存储数据
//
//	// 返回结果
//	return
//}
//
//
//func (s *userService) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (response *pb.GetUsersResponse, err error) {
//	//reply = new(empty.Empty)
//	//fmt.Printf("hello %s", req.AccountName)
//	//accountName := req.AccountName
//	//accountInfo :=
//	// 判断用户名是否存在
//
//	// 判断邮箱是否存在
//
//	// 判断用户角色
//
//	// 生成盐、id
//
//	// 加密密码
//
//	// 存储数据
//
//	// 返回结果
//	return
//}