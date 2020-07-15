package dao

import "github.com/jinzhu/gorm"

type Privilege struct {
	CommonModel
	ID        uint64	`json:"id" gorm"index"`
	PrivilegeName string	`json:"privilege_name"`
	Uri string	`json:"uri"`
	Sn string	`json:"sn"`
	PrivilegeDesc string	`json:"privilege_desc"`
}

type PrivilegeDao struct {
}

func (dao PrivilegeDao) Add(data map[string]interface{}) (err error) {
	privilege := Privilege{
		PrivilegeName:       data["accountId"].(string),
		Uri:     data["accountName"].(string),
		Sn:    data["accountEmail"].(string),
		
		PrivilegeDesc: data["accountPassword"].(string),
	}
	if err := DB.Create(&privilege).Error; err !=nil {
		return err
	}
	return nil
}

func (dao PrivilegeDao) delete(id uint64) (err error) {
	if err := DB.Update("is_delete", 1).Error; err !=nil {
		return err
	}
	return nil
}


func  (dao PrivilegeDao) FindByFields (fields map[string]interface{}) (*Privilege, error)  {
	var account = new(Privilege)
	err := DB.Where(fields).First(account).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return account, nil
}