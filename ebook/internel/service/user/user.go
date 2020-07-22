package user

import (
	"context"
	"ebook/ebook/api/privilege"
	"ebook/ebook/api/user"
	"ebook/ebook/internel/utils/response"
	
	"ebook/ebook/internel/utils/rpc"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(ctx *gin.Context)  {
	var requestBody SignUpRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	password := requestBody.Password
	roleId := requestBody.RoleId
	if len(password) != 32 {
	
	}
	// 判断角色是否存在
	getRoleParams := &privilege.GetRoleRequest{
		Id:                   roleId,
	}
	roleInfo, err := rpc.PrivilegeRpc().GetRole(context.Background(), getRoleParams)
	if roleInfo.Errno != 0 {
		response.Error(ctx, roleInfo.Errmsg)
	}
	signUpParams := &user.AddRequest{
		Username: requestBody.Username,
		Email:    requestBody.Email,
		Password: requestBody.Password,
		RoleId:   requestBody.RoleId,
	}
	signUpRes, err := rpc.UserRpc().Add(context.Background(), signUpParams)
	if err != nil {
		response.Error(ctx, signUpRes.Errmsg)
	}
	
	// 生成 token and set into redis
	response.OK(ctx, signUpRes.Data)
	return
}

func SignIn(ctx *gin.Context)  {
	fmt.Println("SignIn")
}

func SignOut(ctx *gin.Context)  {
	fmt.Println("SignOut")
}