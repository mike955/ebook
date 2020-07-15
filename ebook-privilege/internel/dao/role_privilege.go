package dao

import "github.com/jinzhu/gorm"

type RolePrivilegeMap struct {
	CommonModel
	ID        uint64	`json:"id" gorm"index"`
	RoleId uint64	`json:"role_id"`
	PrivilegeId uint64	`json:"privilege_id"`
}

type RolePrivilegeDao struct {
}

func (dao RolePrivilegeDao) Add(data map[string]interface{}) (err error) {
	rolePrivilegeMap := RolePrivilegeMap{
		RoleId:       data["roleName"].(uint64),
		PrivilegeId:     data["roleDesc"].(uint64),
	}
	if err := DB.Create(&rolePrivilegeMap).Error; err !=nil {
		return err
	}
	return nil
}

func (dao RolePrivilegeDao) delete(id uint64) (err error) {
	if err := DB.Update("is_delete", 1).Error; err !=nil {
		return err
	}
	return nil
}


func  (dao RolePrivilegeDao) FindByFields (fields map[string]interface{}) (*RolePrivilegeMap, error)  {
	var rolePrivilege = new(RolePrivilegeMap)
	err := DB.Where(fields).First(rolePrivilege).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return rolePrivilege, nil
}