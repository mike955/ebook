package user

import (
	"context"
	"ebook/ebook/api/privilege"
	"ebook/ebook/api/user"
	"ebook/ebook/internel/utils/jwt"
	"ebook/ebook/internel/utils/response"
	"log"
	
	"ebook/ebook/internel/utils/rpc"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(ctx *gin.Context)  {
	var requestBody SignUpRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		log.Println("request params error: ", err.Error())
		response.Error(ctx, "PARAMS_ERROR")
		return
	}
	password := requestBody.Password
	roleId := requestBody.RoleId
	if len(password) != 32 {
		log.Println("request params error: password")
		response.Error(ctx, "PARAMS_ERROR", "password error")
		return
	}
	// 判断角色是否存在
	getRoleParams := &privilege.GetRoleRequest{
		Id:                   roleId,
	}
	roleInfo, err := rpc.PrivilegeRpc().GetRole(context.Background(), getRoleParams)
	if roleInfo.Errno != 0 {
		log.Println("request params error: ", roleInfo.Errmsg)
		response.Error(ctx, roleInfo.Errmsg)
		return
	}
	signUpParams := &user.AddRequest{
		Username: requestBody.Username,
		Email:    requestBody.Email,
		Password: requestBody.Password,
		RoleId:   requestBody.RoleId,
	}
	signUpRes, err := rpc.UserRpc().Add(context.Background(), signUpParams)
	if err != nil {
		log.Println("request params error: ", err.Error())
		response.Error(ctx, signUpRes.Errmsg)
		return
	}
	if signUpRes.Errno != 0 {
		log.Println("request params error: ", signUpRes.Errmsg)
		response.Error(ctx, signUpRes.Errmsg)
		return
	}
	response.OK(ctx, signUpRes.Data)
	return
}

func SignIn(ctx *gin.Context)  {
	var requestBody SignInRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		response.Error(ctx, "PARAMS_ERROR")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(requestBody.Password) != 32 {
		response.Error(ctx, "PARAMS_ERROR", "password error")
		return
	}
	signInParams := &user.VerifyPasswordRequest{
		Username: requestBody.Username,
		Password: requestBody.Password,
	}
	signInRes, err := rpc.UserRpc().VerifyPassword(context.Background(), signInParams)
	if err != nil {
		response.Error(ctx, signInRes.Errmsg)
		return
	}
	if signInRes.Errno != 0 {
		response.Error(ctx, signInRes.Errmsg)
		return
	}
	user := signInRes.Data
	token, err := jwt.Sign(user)
	if err != nil {
		response.Error(ctx, "SIGNIN_ERROR")
		return
	}
	ctx.Header("token", token)
	response.OK(ctx, signInRes.Data)
	return
}

func SignOut(ctx *gin.Context)  {
	fmt.Println("SignOut")
}