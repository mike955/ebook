package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Account struct {
	CommonModel
	ID        uint64	`json:"id" gorm"index"`
	AccountId string	`json:"account_id"`
	AccountName string	`json:"account_name"`
	AccountEmail string	`json:"account_email"`
	AccountPassword string	`json:"account_password"`
	Salt string	`json:"salt"`
	AccountRole uint	`json:"account_role"`
	Status uint	`gorm:"default:0" json:"status"`
}

type UserDao struct {
}

func (dao UserDao) Add(data map[string]interface{}) (err error) {
	account := Account{
		AccountId:       data["accountId"].(string),
		AccountName:     data["accountName"].(string),
		AccountEmail:    data["accountEmail"].(string),
		AccountPassword: data["accountPassword"].(string),
		Salt:            data["salt"].(string),
		AccountRole:     data["accountRole"].(uint),
		Status:          data["status"].(uint),
	}
	if err := DB.Create(&account).Error; err !=nil {
		return err
	}
	return nil
}

func (dao UserDao) FindByID(id uint64) (*Account, error) {
	var account = new(Account)
	err := DB.Where(&Account{ID: id}).First(account).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return account, nil
}

func (dao UserDao) FindByAccountId(accountId string) (*Account, error) {
	var account = new(Account)
	err := DB.Where(&Account{AccountId: accountId}).First(account).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return account, nil
}

func (dao UserDao) FindByAccountName(accountName string) (*Account, error) {
	var account = new(Account)
	fmt.Println("============ accountName: ", accountName)
	err := DB.Where(&Account{AccountName: accountName}).Find(account).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return account, nil
}

func  (dao UserDao) FindByFields (fields map[string]interface{}) (*Account, error)  {
	var account = new(Account)
	err := DB.Where(fields).First(account).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return account, nil
}
