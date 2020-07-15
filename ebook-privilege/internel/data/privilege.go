package data

import (
	"ebook/ebook-privilege/internel/dao"
	"fmt"
)

type PrivilegeData struct {
	roleDao dao.RoleDao
	privilegeDao dao.PrivilegeDao
	userRoleDao dao.UserRoleDao
	rolePrivilegeDao dao.RolePrivilegeMap
}

func (data *PrivilegeData) AddUser(userId string, roleId uint32) ( bool, error) {
	roles, err := data.roleDao.GetById(roleId)
	if err != nil {
	
	}
	if len(roles) > 0 {
	
	}
	findUserRolesParams := map[string]interface{}{
		"user_id": userId,
		"is_delete": 0,
	}
	fmt.Println(findUserRolesParams)
	userRoles, err := data.userRoleDao.FindByFields(findUserRolesParams)
	if err != nil {
		return false, err
	}
	if len(userRoles) > 0 {
		// 用户已存在
	}
	addUserRoleParams := map[string]interface{}{
		"user_id": userId,
		"role_id": roleId,
	}
	if err := data.userRoleDao.Add(addUserRoleParams); err != nil {
		return false, err
	}
	fmt.Println("==========addUserRoleRes: \n", addUserRoleParams)
	return true, nil
}

func (data *PrivilegeData) GetRole(roleId uint32) bool {
	role, err := data.roleDao.GetById(roleId)
	if err != nil {
	
	}
	fmt.Println("==========role: \n", role)
	return false
}