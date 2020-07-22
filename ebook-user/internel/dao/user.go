package dao

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	CommonModel
	ID        uint64	`json:"id" gorm"index"`
	UserId string	`json:"user_id"`
	Username string	`json:"username"`
	Email string	`json:"email"`
	Password string	`json:"password"`
	Salt string	`json:"salt"`
	RoleId uint64	`json:"role"`
	Status uint64	`gorm:"default:0" json:"status"`
	IsDelete  uint64 `gorm:"default:0" json:"is_delete"`
}

type UserDao struct {
}

func (dao UserDao) Add(data map[string]interface{}) (err error) {
	account := User{
		UserId:       data["userId"].(string),
		Username:     data["username"].(string),
		Email:    data["email"].(string),
		Password: data["password"].(string),
		Salt:            data["salt"].(string),
		RoleId:     data["roleId"].(uint64),
		Status:          data["status"].(uint64),
	}
	if err := DB.Create(&account).Error; err !=nil {
		return err
	}
	return nil
}

func (dao UserDao) FindByID(id uint64) (*User, error) {
	var account = new(User)
	err := DB.Where(&User{ID: id}).First(account).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return account, nil
}

func (dao UserDao) FindByUserId(userId string) (*User, error) {
	var user = new(User)
	err := DB.Where(&User{UserId: userId}).First(user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return user, nil
}

func  (dao UserDao) FindByFields (fields map[string]interface{}) ([]*User, error)  {
	var users []*User
	err := DB.Where(fields).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, nil
}

func  (dao UserDao) DeleteByUserId (userId string) (err error)  {
	if err := DB.Update("is_delete", 1).Error; err !=nil {
		return err
	}
	return nil
}

func (dao UserDao) UpdateFields(where map[string]interface{}, updateFileds map[string]interface{}) (err error) {
	err = DB.Where(&where).Update(updateFileds).Error
	return
}