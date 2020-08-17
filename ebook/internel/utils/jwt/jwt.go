package jwt

import (
	"ebook/ebook/api/user"
	"ebook/ebook/conf"
	"ebook/ebook/internel/utils/response"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

//type tokenClaims jwt.Claims

type tokenClaims struct {
	userId string
	expire time.Time
}

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			ctx.Abort()
			response.Error(ctx, "TOKEN_ERROR")
			return
		} else {
			data, err := Parse(token)
			if err != nil {
				ctx.Abort()
				response.Error(ctx, "TOKEN_ERROR")
				return
			}
			ctx.Set("User", data)
		}
		ctx.Next()
	}
}

func Sign(user *user.UserResponseInfo) (string, error) {
	expire := time.Now().Add(conf.JWT_EXPIRE_TIME).Unix()
	secret := conf.JWT_SECRET
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.UserId,
		"username": user.Username,
		"expire": expire,
	})
	return token.SignedString([]byte(secret))
}

func Parse(tokenString string) (data map[string]interface{}, err error)  {
	secret := conf.JWT_SECRET
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		data = map[string]interface{}{
			"userId": claims["userId"],
			"username": claims["username"],
			"expire": claims["expire"],
		}
	}
	return
}