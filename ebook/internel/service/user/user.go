package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context)  {
	
	fmt.Println("SignUp")
}

func SignIn(ctx *gin.Context)  {
	fmt.Println("SignIn")
}

func SignOut(ctx *gin.Context)  {
	fmt.Println("SignOut")
}