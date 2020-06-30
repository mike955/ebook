package data

import (
	pb "ebook/api/user"
	"ebook/internal/dao"
	"ebook/pkg/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type UserData struct {
	UserDao dao.UserDao
}

func (data UserData) getByID(id int64) (response *pb.UserInfo, err error){
	
	fmt.Println("hello ")
	//response = new(pb.SignUpResponse)
	//response.Error = 0
	//response.Errmsg = ""
	//response.Data = &pb.UserInfo{
	//}
	return
}

func (data UserData) SignUp(params *pb.SignUpRequest) (response *pb.SignUpResponse, err error){
	//data := make(map[string]interface{})
	_, err = data.UserDao.FindByAccountName(params.AccountName)
	if err != gorm.ErrRecordNotFound {
	
	}
	insertData := make(map[string]interface{})
	insertData["accountId"] = utils.GenerateAccountId()
	insertData["accountName"] = params.AccountName
	insertData["accountEmail"] = params.AccountEmail
	
	//fmt.Println("hello ", params.AccountName)
	//response = new(pb.SignUpResponse)
	//response.Error = 0
	//response.Errmsg = ""
	//response.Data = &pb.UserInfo{
	//}
	return
}