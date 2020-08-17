package dao

import "github.com/jinzhu/gorm"

type Privilege struct {
	CommonModel
	ID            uint64 `gorm:"primary_key" json:"id"`
	PrivilegeName string `json:"privilege_name"`
	Uri           string `json:"uri"`
	Sn            string `json:"sn"`
	PrivilegeDesc string `json:"privilege_desc"`
	IsDelete      uint64 `gorm:"default:0" json:"is_delete"`
}

type PrivilegeDao struct {
}

func (dao PrivilegeDao) Add(data map[string]interface{}) (err error) {
	privilege := Privilege{
		PrivilegeName: data["accountId"].(string),
		Uri:           data["accountName"].(string),
		Sn:            data["accountEmail"].(string),

		PrivilegeDesc: data["accountPassword"].(string),
	}
	if err := DB.Create(&privilege).Error; err != nil {
		return err
	}
	return nil
}

func (dao PrivilegeDao) delete(id uint64) (err error) {
	if err := DB.Update("is_delete", 1).Error; err != nil {
		return err
	}
	return nil
}

func (dao PrivilegeDao) FindByFields(fields map[string]interface{}) ([]*Privilege, error) {
	var privileges []*Privilege
	err := DB.Where(fields).Find(&privileges).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return privileges, nil
}

func (dao PrivilegeDao) UpdateFields(where map[string]interface{}, updateFields map[string]interface{}) (err error) {
	err = DB.Where(&where).Update(updateFields).Error
	return
}
