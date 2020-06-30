package dao

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	CommonModel
	ID        uint64	`json:"id" gorm"index"`
	AccountId string	`json:"account_id"`
	AccountName string	`json:"account_name"`
	AccountEmail string	`json:"account_email"`
	AccountPassword string	`json:"account_password"`
	Salt string	`json:"salt"`
	AccountRole uint	`json:"account_role"`
	Status uint	`json:"status"`
	IsDelete uint	`json:"is_delete"`
}

type UserDao struct {
}

func (dao UserDao) Add(data map[string]interface{}) (err error) {
	user := User{
		AccountId:       data["accountId"].(string),
		AccountName:     data["accountName"].(string),
		AccountEmail:    data["accountEmail"].(string),
		AccountPassword: data["accountPassword"].(string),
		Salt:            data["salt"].(string),
		AccountRole:     data["accountRole"].(uint),
		Status:          data["status"].(uint),
		IsDelete:        data["isDelete"].(uint),
	}
	if err := DB.Create(&user).Error; err !=nil {
		return err
	}
	return nil
}

func (dao UserDao) FindByID(id uint64) (user User, err error) {
	err = DB.Where("id=?", id).Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return user, nil
}

func (dao UserDao) FindByAccountId(accountId string) (user User, err error) {
	err = DB.Where("account_id=?", accountId).Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}
	return user, nil
}

func (dao UserDao) FindByAccountName(accountName string) (user User, err error) {
	err = DB.Where("account_name=?", accountName).Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return user, nil
}

