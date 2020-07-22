package routers

import (
	"ebook/ebook/internel/service/user"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter(engine *gin.Engine) {
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	
	//r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	//r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	//r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))
	
	//r.POST("/auth", api.GetAuth)
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.POST("/upload", api.UploadImage)
	
	//userApi := r.Group("/ebook/user")
	//userApi.Use(jwt.JWT())
	userApi := engine.Group("/ebook/user")
	{
		userApi.POST("/signUp", user.SignUp)   // 注册
		userApi.POST("/signIn", user.SignIn)   // 登录
		userApi.POST("/signOut", user.SignOut) // 登出
	}
	
	//ebookApi := r.Group("/ebook/ebook")
	//return r
}