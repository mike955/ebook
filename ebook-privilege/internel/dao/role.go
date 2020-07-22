package dao

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	CommonModel
	ID        uint64	`gorm:"primary_key" json:"id"`
	RoleName string	`json:"role_name"`
	RoleDesc string	`json:"role_desc"`
	IsDelete  uint64 `gorm:"default:0" json:"is_delete"`
}

type RoleDao struct {
}

func (dao RoleDao) Add(data map[string]interface{}) (err error) {
	role := Role{
		RoleName:       data["roleName"].(string),
		RoleDesc:     data["roleDesc"].(string),
	}
	if err := DB.Create(&role).Error; err !=nil {
		return err
	}
	return nil
}

func (dao RoleDao) delete(id uint64) (err error) {
	if err := DB.Update("is_delete", 1).Error; err !=nil {
		return err
	}
	return nil
}

func (dao RoleDao) GetById(id uint64) ([]*Role, error) {
	var roles  []*Role
	err := DB.Where(&Role{ID:id}).Find(&roles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return roles, nil
}


func  (dao RoleDao) FindByFields (fields map[string]interface{}) ([]*Role, error)  {
	var roles []*Role
	err := DB.Where(fields).Find(&roles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return roles, nil
}

func (dao RoleDao) UpdateFields(where map[string]interface{}, updateFields map[string]interface{}) (err error) {
	err = DB.Where(&where).Update(updateFields).Error
	return
}