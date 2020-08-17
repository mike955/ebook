package service

import (
	"context"
	pb "ebook/ebook-privilege/api/privilege"
	"ebook/ebook-privilege/internel/dao"
	"ebook/ebook-privilege/pkg/err_code"
	"fmt"
)

type privilegeService struct {
	roleDao          dao.RoleDao
	privilegeDao     dao.PrivilegeDao
	userRoleDao      dao.UserRoleDao
	rolePrivilegeDao dao.RolePrivilegeDao
}

var PrivilegeService = &privilegeService{}

func (service *privilegeService) AddUser(ctx context.Context, req *pb.AddUserRequest) (response *pb.AddUserResponse, err error) {
	response = new(pb.AddUserResponse)
	userId := req.GetUserId()
	roleId := req.GetRoleId()
	// 验证 grpc 是否会校验参数类型

	// 判断角色 id 是否存在
	roles, err := service.roleDao.FindByFields(map[string]interface{}{"id": roleId, "is_delete": 0})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_ROLE_ERROR")
		return
	}
	if len(roles) == 0 {
		response.Errno, response.Errmsg = err_code.Code("ROLE_IS_NOT_EXIST_ERROR")
		return
	}

	// 判断用户是否存在
	users, err := service.userRoleDao.FindByFields(map[string]interface{}{"user_id": userId, "role_id": roleId})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_USER_ROLE_ERROR")
		return
	}
	if len(users) > 1 {
		response.Errno, response.Errmsg = err_code.Code("USER_ROLE_IS_EXIST_ERROR")
		return
	}
	err = service.userRoleDao.Add(map[string]interface{}{"user_id": userId, "role_id": roleId})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("ADD_USER_ROLE_ERROR")
		return
	}
	response.Data = true
	return
}

func (service *privilegeService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (response *pb.DeleteUserResponse, err error) {
	response = new(pb.DeleteUserResponse)
	userId := req.GetUserId()
	// 判断角色 id 是否存在
	users, err := service.userRoleDao.FindByFields(map[string]interface{}{"user_id": userId, "is_delete": 0})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_USER_ROLE_ERROR")
		return
	}
	if len(users) == 0 {
		response.Errno, response.Errmsg = err_code.Code("USER_ROLE_IS_NOT_EXIST_ERROR")
		return
	}
	if err = service.userRoleDao.UpdateFields(map[string]interface{}{"user_id": userId}, map[string]interface{}{"is_delete": 1}); err != nil {
		response.Errno, response.Errmsg = err_code.Code("DELETE_USER_ROLE_ERROR")
		return
	}
	response.Data = true
	return
}

func (service *privilegeService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (response *pb.UpdateUserResponse, err error) {
	response = new(pb.UpdateUserResponse)
	userId := req.UserId
	roleId := req.RoleId
	// 判断用户 id 是否存在
	users, err := service.userRoleDao.FindByFields(map[string]interface{}{"user_id": userId, "is_delete": 0})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_USER_ROLE_ERROR")
		return
	}
	if len(users) == 0 {
		response.Errno, response.Errmsg = err_code.Code("USER_IS_NOT_EXIST_ERROR")
		return
	}
	// 判断角色 id 是否存在
	roles, err := service.roleDao.FindByFields(map[string]interface{}{"id": roleId, "is_delete": 0})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_ROLE_ERROR")
		return
	}
	if len(roles) == 0 {
		response.Errno, response.Errmsg = err_code.Code("ROLE_IS_NOT_EXIST_ERROR")
		return
	}

	if err = service.userRoleDao.UpdateFields(map[string]interface{}{"user_id": userId}, map[string]interface{}{"role_id": roleId}); err != nil {
		response.Errno, response.Errmsg = err_code.Code("UPDATE_USER_ROLE_ERROR")
		return
	}
	response.Data = true
	return
}

func (service *privilegeService) GetUser(ctx context.Context, req *pb.GetUserRequest) (response *pb.GetUserResponse, err error) {
	// 判断用户是否存在
	userId := req.UserId
	users, err := service.userRoleDao.FindByFields(map[string]interface{}{"user_id": userId, "is_delete": 0})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_USER_ROLE_ERROR")
		return
	}
	if len(users) != 0 {
		response.Errno, response.Errmsg = err_code.Code("USER_ROLE_IS_NOT_EXIST_ERROR")
		return
	}
	roleId := users[0].RoleId
	roles, err := service.roleDao.FindByFields(map[string]interface{}{"id": roleId, "is_delete": 0})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_ROLE_ERROR")
		return
	}
	if len(roles) == 0 {
		response.Errno, response.Errmsg = err_code.Code("ROLE_IS_NOT_EXIST_ERROR")
		return
	}
	rolePrivileges, err := service.rolePrivilegeDao.FindByFields(map[string]interface{}{"role_id": roleId})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_ROLE_PRIVILEGE_ERROR")
		return
	}
	if len(rolePrivileges) == 0 {
		response.Errno, response.Errmsg = err_code.Code("ROLE_IS_NOT_EXIST_ERROR")
		return
	}
	privilegeIds := make([]uint64, len(rolePrivileges))
	for _, rolePrivilege := range rolePrivileges {
		privilegeIds = append(privilegeIds, rolePrivilege.PrivilegeId)
	}
	privileges, err := service.privilegeDao.FindByFields(map[string]interface{}{"id": privilegeIds})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_PRIVILEGE_ERROR")
		return
	}
	if len(rolePrivileges) == 0 {
		response.Errno, response.Errmsg = err_code.Code("ROLE_PRIVILEGE_IS_NOT_EXIST_ERROR")
		return
	}

	privilegesInfo := make([]*pb.PrivilegeInfo, len(privileges))
	for _, privilege := range privileges {
		privilegeInfo := &pb.PrivilegeInfo{
			Id:            privilege.ID,
			PrivilegeName: privilege.PrivilegeName,
			Uri:           privilege.Uri,
			Sn:            privilege.Sn,
			PrivilegeDesc: privilege.PrivilegeDesc,
			IsDelete:      privilege.IsDelete,
			CreateTime:    privilege.CreateTime,
			UpdateTime:    privilege.UpdateTime,
		}
		privilegesInfo = append(privilegesInfo, privilegeInfo)
	}
	role := roles[0]
	response.Data = (*pb.GetUserResponse_Data)(&pb.GetRoleResponse_Data{
		RoleInfo: &pb.RoleInfo{
			Id:         role.ID,
			RoleName:   role.RoleName,
			RoleDesc:   role.RoleDesc,
			IsDelete:   role.IsDelete,
			CreateTime: role.CreateTime,
			UpdateTime: role.UpdateTime,
		},
		Privileges: privilegesInfo,
	})
	return
}

func (service *privilegeService) AddRole(ctx context.Context, req *pb.AddRoleRequest) (response *pb.AddRoleResponse, err error) {
	response = new(pb.AddRoleResponse)
	roleName := req.RoleName
	roleDesc := req.RoleDesc
	// 判断角色 id 是否存在
	roles, err := service.roleDao.FindByFields(map[string]interface{}{"role_name": roleName, "is_delete": 0})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_ROLE_ERROR")
		return
	}
	if len(roles) > 0 {
		response.Errno, response.Errmsg = err_code.Code("ROLE_IS_EXIST_ERROR")
		return
	}
	if err = service.roleDao.Add(map[string]interface{}{"role_name": roleName, "role_desc": roleDesc}); err != nil {
		response.Errno, response.Errmsg = err_code.Code("ADD_ROLE_ERROR")
		return
	}
	response.Data = true
	return
}

func (service *privilegeService) DeleteRole(ctx context.Context, req *pb.DeleteRoleRequest) (response *pb.DeleteRoleResponse, err error) {
	response = new(pb.DeleteRoleResponse)
	roleId := req.RoleId
	// 判断角色 id 是否存在
	roles, err := service.roleDao.FindByFields(map[string]interface{}{"id": roleId, "is_delete": 0})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_ROLE_ERROR")
		return
	}
	if len(roles) == 0 {
		response.Errno, response.Errmsg = err_code.Code("ROLE_IS_NOT_EXIST_ERROR")
		return
	}
	if err = service.roleDao.UpdateFields(map[string]interface{}{"id": roleId}, map[string]interface{}{"is_delete": 1}); err != nil {
		response.Errno, response.Errmsg = err_code.Code("ADD_ROLE_ERROR")
		return
	}
	response.Data = true
	return
}

func (service *privilegeService) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (response *pb.UpdateRoleResponse, err error) {
	response = new(pb.UpdateRoleResponse)
	roleId := req.Id
	// 判断角色 id 是否存在
	roles, err := service.roleDao.FindByFields(map[string]interface{}{"id": roleId, "is_delete": 0})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_ROLE_ERROR")
		return
	}
	if len(roles) == 0 {
		response.Errno, response.Errmsg = err_code.Code("ROLE_IS_NOT_EXIST_ERROR")
		return
	}
	if err = service.roleDao.UpdateFields(map[string]interface{}{"id": roleId}, map[string]interface{}{"is_delete": 1}); err != nil {
		response.Errno, response.Errmsg = err_code.Code("ADD_ROLE_ERROR")
		return
	}
	response.Data = true
	return
}

func (service *privilegeService) GetRole(ctx context.Context, req *pb.GetRoleRequest) (response *pb.GetRoleResponse, err error) {
	fmt.Println(req)
	response = new(pb.GetRoleResponse)
	conditions := map[string]interface{}{"is_delete": 0}
	if roleId := req.Id; roleId != 0 {
		conditions["id"] = roleId
	}
	if roleName := req.RoleName; roleName != "" {
		conditions["role_name"] = roleName // todo like
	}
	if roleDesc := req.RoleDesc; roleDesc != "" {
		conditions["role_desc"] = roleDesc // todo like
	}
	roles, err := service.roleDao.FindByFields(conditions)
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_ROLE_ERROR")
		return
	}
	if len(roles) == 0 {
		response.Errno, response.Errmsg = err_code.Code("ROLE_IS_NOT_EXIST_ERROR")
		return
	}
	roleId := roles[0].ID
	// get role privilege
	rolePrivileges, err := service.rolePrivilegeDao.FindByFields(map[string]interface{}{"role_id": roleId, "is_delete": 0})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_ROLE_PRIVILEGE_ERROR")
		return
	}
	privilegesInfo := make([]*pb.PrivilegeInfo, 0)
	response.Data = &pb.GetRoleResponse_Data{
		RoleInfo: &pb.RoleInfo{
			Id:         roles[0].ID,
			RoleName:   roles[0].RoleName,
			RoleDesc:   roles[0].RoleDesc,
			IsDelete:   roles[0].IsDelete,
			CreateTime: roles[0].CreateTime,
			UpdateTime: roles[0].UpdateTime,
		},
		Privileges: privilegesInfo,
	}
	fmt.Println(rolePrivileges)
	if len(rolePrivileges) != 0 {
		privilegeIds := make([]uint64, len(rolePrivileges))
		for _, rolePrivilege := range rolePrivileges {
			privilegeIds = append(privilegeIds, rolePrivilege.PrivilegeId)
		}
		privileges, err := service.privilegeDao.FindByFields(map[string]interface{}{"id": privilegeIds})
		if err != nil {
			response.Errno, response.Errmsg = err_code.Code("GET_PRIVILEGE_ERROR")
			return response, nil
		}
		if len(privileges) > 0 {
			for _, privilege := range privileges {
				privilegeInfo := &pb.PrivilegeInfo{
					Id:            privilege.ID,
					PrivilegeName: privilege.PrivilegeName,
					Uri:           privilege.Uri,
					Sn:            privilege.Sn,
					PrivilegeDesc: privilege.PrivilegeDesc,
					IsDelete:      privilege.IsDelete,
					CreateTime:    privilege.CreateTime,
					UpdateTime:    privilege.UpdateTime,
				}
				privilegesInfo = append(privilegesInfo, privilegeInfo)
			}
			response.Data.Privileges = privilegesInfo
		}
	}
	return
}

func (service *privilegeService) AddPrivilege(ctx context.Context, req *pb.AddPrivilegeRequest) (response *pb.AddPrivilegeResponse, err error) {
	response = new(pb.AddPrivilegeResponse)
	privileges, err := service.privilegeDao.FindByFields(map[string]interface{}{"uri": req.Uri})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_PRIVILEGE_ERROR")
		return
	}
	if len(privileges) != 0 {
		response.Errno, response.Errmsg = err_code.Code("PRIVILEGE_IS_EXIST_ERROR")
		return
	}
	newPrivilege := map[string]interface{}{
		"privilege_name": req.PrivilegeName,
		"uri":            req.Uri,
		"sn":             req.Sn,
		"privilege_desc": req.PrivilegeDesc,
	}
	err = service.privilegeDao.Add(newPrivilege)
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("ADD_PRIVILEGE_ERROR")
		return
	}
	return
}

func (service *privilegeService) AddPrivileges(ctx context.Context, req *pb.AddPrivilegesRequest) (response *pb.AddPrivilegesResponse, err error) {
	return
}

func (service *privilegeService) DeletePrivilege(ctx context.Context, req *pb.DeletePrivilegeRequest) (response *pb.DeletePrivilegeResponse, err error) {
	response = new(pb.DeletePrivilegeResponse)
	privilegeId := req.PrivilegeId
	privileges, err := service.privilegeDao.FindByFields(map[string]interface{}{"id": privilegeId, "is_delete": 0})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_PRIVILEGE_ERROR")
		return
	}
	if len(privileges) == 0 {
		response.Errno, response.Errmsg = err_code.Code("PRIVILEGE_IS_NOT_EXIST_ERROR")
		return
	}
	err = service.privilegeDao.UpdateFields(map[string]interface{}{"id": privilegeId}, map[string]interface{}{"is_delete": 1})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("DELETE_PRIVILEGE_ERROR")
		return
	}
	return
}

func (service *privilegeService) DeletePrivileges(ctx context.Context, req *pb.DeletePrivilegesRequest) (response *pb.DeletePrivilegesResponse, err error) {
	response = new(pb.DeletePrivilegesResponse)
	privilegeIds := req.PrivilegeIds
	privileges, err := service.privilegeDao.FindByFields(map[string]interface{}{"id": privilegeIds, "is_delete": 0})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_PRIVILEGE_ERROR")
		return
	}
	if len(privileges) != len(privilegeIds) {
		response.Errno, response.Errmsg = err_code.Code("PRIVILEGE_IS_NOT_EXIST_ERROR")
		return
	}
	err = service.privilegeDao.UpdateFields(map[string]interface{}{"id": privilegeIds}, map[string]interface{}{"is_delete": 1})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("DELETE_PRIVILEGE_ERROR")
		return
	}
	return
}

func (service *privilegeService) UpdatePrivilege(ctx context.Context, req *pb.UpdatePrivilegeRequest) (response *pb.UpdatePrivilegeResponse, err error) {
	response = new(pb.UpdatePrivilegeResponse)
	privilegeId := req.PrivilegeId
	privileges, err := service.privilegeDao.FindByFields(map[string]interface{}{"id": privilegeId})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_PRIVILEGE_ERROR")
		return
	}
	if len(privileges) == 0 {
		response.Errno, response.Errmsg = err_code.Code("PRIVILEGE_IS_NOT_EXIST_ERROR")
		return
	}
	updateFields := map[string]interface{}{}
	if privilegeName := req.PrivilegeName; privilegeName != "" {
		updateFields["privilege_name"] = privilegeName
	}
	if uri := req.Uri; uri != "" {
		updateFields["uri"] = uri
	}
	if sn := req.Sn; sn != "" {
		updateFields["sn"] = sn
	}
	if privilegeDesc := req.PrivilegeDesc; privilegeDesc != "" {
		updateFields["privilege_desc"] = privilegeDesc
	}
	err = service.privilegeDao.UpdateFields(map[string]interface{}{"id": privilegeId}, updateFields)
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("DELETE_PRIVILEGE_ERROR")
		return
	}
	return
}

func (service *privilegeService) GetPrivilege(ctx context.Context, req *pb.GetPrivilegeRequest) (response *pb.GetPrivilegeResponse, err error) {
	response = new(pb.GetPrivilegeResponse)
	privilegeId := req.PrivilegeId
	privileges, err := service.privilegeDao.FindByFields(map[string]interface{}{"id": privilegeId, "is_delete": 0})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_PRIVILEGE_ERROR")
		return
	}
	if len(privileges) == 0 {
		response.Errno, response.Errmsg = err_code.Code("PRIVILEGE_IS_NOT_EXIST_ERROR")
		return
	}
	privilege := privileges[0]
	response.Data = &pb.PrivilegeInfo{
		Id:            privilege.ID,
		PrivilegeName: privilege.PrivilegeName,
		Uri:           privilege.Uri,
		Sn:            privilege.Sn,
		PrivilegeDesc: privilege.PrivilegeDesc,
		IsDelete:      privilege.IsDelete,
		CreateTime:    privilege.CreateTime,
		UpdateTime:    privilege.UpdateTime,
	}
	return
}
