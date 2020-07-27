package service

import (
	"context"
	pb "ebook/ebook-user/api/user"
	"ebook/ebook-user/internel/dao"
	"ebook/ebook-user/pkg/err_code"
	"ebook/ebook-user/pkg/utils"
	"log"
)

type userService struct {
	userDao dao.UserDao
}

var UserService = &userService{
}


func (service *userService) Add(ctx context.Context, req *pb.AddRequest) (response *pb.AddResponse, err error){
	response = new(pb.AddResponse)
	// 判断用户是否存在
	checkUsernameRes, err := service.userDao.FindByFields(map[string]interface{}{"username": req.Username})
	if err != nil {
		log.Println("")
		response.Errno, response.Errmsg = err_code.Code("ADD_USER_ERROR")
		return
	}
	if len(checkUsernameRes) != 0 {
		response.Errno, response.Errmsg = err_code.Code("USERNAME_IS_EXIST_ERROR")
		return
	}
	
	// 判断邮箱是否存在
	checkEmailRes, err := service.userDao.FindByFields(map[string]interface{}{"email": req.Email})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("ADD_USER_ERROR")
		return
	}
	if len(checkEmailRes) != 0 {
		response.Errno, response.Errmsg = err_code.Code("EMAIL_IS_EXIST_ERROR")
		return
	}
	
	// add salt and hash(salt+password)
	condition := make(map[string]interface{})
	condition["userId"] = utils.GenerateRandomString(16)
	condition["username"] = req.Username
	condition["email"] = req.Email
	condition["roleId"] = req.RoleId
	condition["status"] = uint64(0)
	if status := req.Status; err != nil {
		condition["status"] = status
	}
	if isDelete := req.IsDelete; err != nil {
		condition["is_delete"] = isDelete
	}
	
	salt := utils.GenerateRandomHex(64)
	condition["salt"] = salt
	condition["password"] = utils.GeneratePassword(req.Password, salt)
	err = service.userDao.Add(condition)
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("ADD_USER_ERROR")
		return
	}
	user, err := service.userDao.FindByUserId(condition["userId"].(string))
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("ADD_USER_ERROR")
		return
	}
	response.Data = &pb.UserResponseInfo{
		Id:                   user.ID,
		UserId:               user.UserId,
		Username:             user.Username,
		Email:                user.Email,
		RoleId:               user.RoleId,
		Status:               user.Status,
		IsDelete:             user.IsDelete,
		CreateTime:           user.CreateTime,
		UpdateTime:           user.UpdateTime,
	}
	return
}

func (service *userService) Delete(ctx context.Context, req *pb.DeleteRequest) (response *pb.DeleteResponse, err error){
	response = new(pb.DeleteResponse)
	users, err := service.userDao.FindByFields(map[string]interface{}{"user_id": req.UserId})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_USER_ERROR")
		return
	}
	if len(users) == 0 {
		response.Errno, response.Errmsg = err_code.Code("USER_IS_NOT_EXIST_ERROR")
		return
	}
	if err = service.userDao.DeleteByUserId(req.UserId); err != nil {
		response.Errno, response.Errmsg = err_code.Code("DELETE_USER_ERROR")
		return
	}
	response.Data = true
	return
}

func (service *userService) Update(ctx context.Context, req *pb.UpdateRequest) (response *pb.UpdateResponse, err error){
	users, err := service.userDao.FindByFields(map[string]interface{}{"user_id": req.UserId})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_USER_ERROR")
		return
	}
	if len(users) == 0 {
		response.Errno, response.Errmsg = err_code.Code("USER_IS_NOT_EXIST_ERROR")
		return
	}
	where := map[string]interface{}{"user_id": req.UserId}
	updateFields := map[string]interface{}{}
	if username := req.Username; err != nil {
		updateFields["username"] = username
	}
	if email := req.Email; err != nil {
		updateFields["email"] = email
	}
	if roleId := req.RoleId; err != nil {
		updateFields["role_id"] = roleId
	}
	if status := req.Status; err != nil {
		updateFields["status"] = status
	}
	if isDelete := req.IsDelete; err != nil {
		updateFields["is_delete"] = isDelete
	}
	err = service.userDao.UpdateFields(where, updateFields)
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("UPDATE_ERROR")
		return
	}
	response.Data = true
	return
}

func (service *userService) Get(ctx context.Context, req *pb.GetRequest) (response *pb.GetResponse, err error){
	response = new(pb.GetResponse)
	condition := make(map[string]interface{})
	if useId := req.UserId; useId != ""  {
		condition["user_id"] = useId
	}
	if username := req.Username; username != "" {		// todo convert like
		condition["username"] = "%" + username + "%"
	}
	if email := req.Email; email != "" {
		condition["email"] = "%" + email + "%"		// todo convert like
	}
	if roleId := req.RoleId; roleId != 0 {
		condition["role_id"] = roleId
	}
	if status := req.Status; status != 0 {
		condition["status"] = status
	}
	if isDelete := req.IsDelete; isDelete != 0 {
		condition["is_delete"] = isDelete
	}
	if createTime := req.CreateTime; createTime != "" {
		condition["create_time"] = createTime
	}
	if updateTime := req.UpdateTime; updateTime != "" {
		condition["update_time"] = updateTime
	}
	users, err := service.userDao.FindByFields(condition)
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_USER_ERROR")
		response.Errmsg = err.Error()
		return
	}
	for _, user := range users {
		userInfo := &pb.UserResponseInfo{
			Id:                   user.ID,
			UserId:               user.UserId,
			Username:             user.Username,
			Email:                user.Email,
			RoleId:               user.RoleId,
			Status:               user.Status,
			IsDelete:             user.IsDelete,
			CreateTime:           user.CreateTime,
			UpdateTime:           user.UpdateTime,
		}
		response.Data = append(response.Data, userInfo)
	}
	return
}

func (service *userService) Gets(ctx context.Context, req *pb.GetsRequest) (response *pb.GetsResponse, err error){
	response = new(pb.GetsResponse)
	condition := make(map[string]interface{})
	if useIds := req.UserIds; len(useIds) > 0  {
		condition["user_ids"] = useIds
	}
	if username := req.Username; username != "" {		// todo convert like
		condition["username"] = "%" + username + "%"
	}
	if email := req.Email; email != "" {
		condition["email"] = "%" + email + "%"		// todo convert like
	}
	if roleId := req.RoleId; roleId != 0 {
		condition["role_id"] = roleId
	}
	if status := req.Status; status != 0 {
		condition["status"] = status
	}
	if isDelete := req.IsDelete; isDelete != 0 {
		condition["is_delete"] = isDelete
	}
	if createTime := req.CreateTime; createTime != "" {
		condition["create_time"] = createTime
	}
	if updateTime := req.UpdateTime; updateTime != "" {
		condition["update_time"] = updateTime
	}
	users, err := service.userDao.FindByFields(condition)
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_USER_ERROR")
		return
	}
	for _, user := range users {
		userInfo := &pb.UserResponseInfo{
			Id:                   user.ID,
			UserId:               user.UserId,
			Username:             user.Username,
			Email:                user.Email,
			RoleId:               user.RoleId,
			Status:               user.Status,
			IsDelete:             user.IsDelete,
			CreateTime:           user.CreateTime,
			UpdateTime:           user.UpdateTime,
		}
		response.Data = append(response.Data, userInfo)
	}
	return
}

func (service *userService) VerifyPassword(ctx context.Context, req *pb.VerifyPasswordRequest) (response *pb.VerifyPasswordResponse, err error){
	response = new(pb.VerifyPasswordResponse)
	users, err := service.userDao.FindByFields(map[string]interface{}{"username": req.Username, "is_delete": 0, "status": 0})
	if err != nil {
		response.Errno, response.Errmsg = err_code.Code("GET_USER_ERROR")
		return
	}
	if len(users) == 0 {
		response.Errno, response.Errmsg = err_code.Code("USER_IS_NOT_EXIST_ERROR")
		return
	}
	user := users[0]
	salt := user.Salt
	inputPassword := req.Password
	if checkPassword := utils.GeneratePassword(inputPassword, salt); checkPassword != user.Password {
		response.Errno, response.Errmsg = err_code.Code("PASSWORD_ERROR")
		return
	}
	response.Data = &pb.UserResponseInfo{
		Id:                   user.ID,
		UserId:               user.UserId,
		Username:             user.Username,
		Email:                user.Email,
		RoleId:               user.RoleId,
		Status:               user.Status,
		IsDelete:             user.IsDelete,
		CreateTime:           user.CreateTime,
		UpdateTime:           user.UpdateTime,
	}
	return
}
