package data

import (
	pb "ebook/api/user"
	"ebook/internal/dao"
	"ebook/pkg/utils"
	"fmt"
	"log"
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
	response = &pb.SignUpResponse{
		Error: 0,
		Errmsg: "",
	}
	response.Error = 0
	response.Errmsg = ""
	account, err := data.UserDao.FindByAccountName(params.AccountName)
	fmt.Println(account)
	fmt.Println(err)
	if err != nil {
		log.Println("Sign up error")
		response.Errmsg = err.Error()
		return
	}
	if account != nil {
		response.Errmsg = "account_name has exist"
		return
	}
	insertData := make(map[string]interface{})
	
	salt := utils.GenerateAccountId()
	fmt.Println(salt)
	insertData["accountId"] = utils.GenerateRandom(16)
	insertData["accountName"] = params.AccountName
	insertData["accountEmail"] = params.AccountEmail
	insertData["salt"] = utils.GenerateRandom(32)
	insertData["accountPassword"] = utils.Sha512(fmt.Sprintf("%s%s", insertData["salt"], params.AccountPassword))
	insertData["accountRole"] = uint(0)
	insertData["status"] = uint(0)
	
	if err = data.UserDao.Add(insertData); err != nil {
		response.Error = 111
		response.Data = false
		response.Errmsg = "sign up error"
	} else {
		response.Data = true
	}
	return
}