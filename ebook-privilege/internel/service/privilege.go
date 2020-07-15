package service

import (
	"context"
	pb "ebook/ebook-privilege/api/privilege"
	"ebook/ebook-privilege/internel/data"
)

type privilegeService struct {
	privilegeData data.PrivilegeData
}

var PrivilegeService = &privilegeService{}

func (service *privilegeService) AddUser(ctx context.Context, req *pb.AddUserRequest) (response *pb.AddUserResponse, err error){
	userId := req.GetUserId()
	roleId := req.GetRoleId()
	res, err := service.privilegeData.AddUser(userId, roleId)
	if err != nil {
	
	}
	response.Data = res
	return
}

func (service *privilegeService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (response *pb.DeleteUserResponse, err error){
	return
}

func (service *privilegeService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (response *pb.UpdateUserResponse, err error){
	return
}

func (service *privilegeService) GetUser(ctx context.Context, req *pb.GetUserRequest) (response *pb.GetUserResponse, err error){
	return
}

func (service *privilegeService) AddRole(ctx context.Context, req *pb.AddRoleRequest) (response *pb.AddRoleResponse, err error){
	return
}

func (service *privilegeService) DeleteRole(ctx context.Context, req *pb.DeleteRoleRequest) (response *pb.DeleteRoleResponse, err error){
	return
}

func (service *privilegeService) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (response *pb.UpdateRoleResponse, err error){
	return
}

func (service *privilegeService) GetRole(ctx context.Context, req *pb.GetRoleRequest) (response *pb.GetRoleResponse, err error){
	return
}

func (service *privilegeService) AddPrivilege(ctx context.Context, req *pb.AddPrivilegeRequest) (response *pb.AddPrivilegeResponse, err error){
	return
}

func (service *privilegeService) AddPrivileges(ctx context.Context, req *pb.AddPrivilegesRequest) (response *pb.AddPrivilegesResponse, err error){
	return
}

func (service *privilegeService) DeletePrivilege(ctx context.Context, req *pb.DeletePrivilegeRequest) (response *pb.DeletePrivilegeResponse, err error){
	return
}

func (service *privilegeService) DeletePrivileges(ctx context.Context, req *pb.DeletePrivilegesRequest) (response *pb.DeletePrivilegesResponse, err error){
	return
}

func (service *privilegeService) UpdatePrivilege(ctx context.Context, req *pb.UpdatePrivilegeRequest) (response *pb.UpdatePrivilegeResponse, err error){
	return
}

func (service *privilegeService) GetPrivilege(ctx context.Context, req *pb.GetPrivilegeRequest) (response *pb.GetPrivilegeResponse, err error){
	return
}

