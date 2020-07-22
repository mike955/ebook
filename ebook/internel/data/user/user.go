package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(ctx *gin.Context)  {
	var body SignInRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	return
	fmt.Println("SignUp")
}