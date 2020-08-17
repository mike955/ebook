package dao

import "github.com/jinzhu/gorm"

type UserRoleMap struct {
	CommonModel
	ID       uint64 `gorm:"primary_key" json:"id"`
	UserId   string `json:"user_id"`
	RoleId   uint64 `json:"role_id"`
	IsDelete uint   `gorm:"default:0" json:"is_delete"`
}

type UserRoleDao struct {
}

func (dao UserRoleDao) Add(data map[string]interface{}) (err error) {
	rolePrivilegeMap := UserRoleMap{
		UserId: data["roleDesc"].(string),
		RoleId: data["roleName"].(uint64),
	}
	if err := DB.Create(&rolePrivilegeMap).Error; err != nil {
		return err
	}
	return nil
}

func (dao UserRoleDao) Delete(id uint64) (err error) {
	if err := DB.Update("is_delete", 1).Error; err != nil {
		return err
	}
	return nil
}

func (dao UserRoleDao) FindByFields(fields map[string]interface{}) ([]*UserRoleMap, error) {
	var userRoles []*UserRoleMap
	err := DB.Where(fields).Find(&userRoles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return userRoles, nil
}

func (dao UserRoleDao) UpdateFields(where map[string]interface{}, updateFileds map[string]interface{}) (err error) {
	err = DB.Where(&where).Update(updateFileds).Error
	return
}
