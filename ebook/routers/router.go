package routers

import (
	"ebook/ebook/internel/service/ebook"
	"ebook/ebook/internel/service/user"
	"ebook/ebook/internel/utils/jwt"
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
	
	userApi := engine.Group("/ebook/user")
	userApi.POST("/signIn", user.SignIn)   // 登录
	userApi.Use(jwt.JWT())
	{
		userApi.POST("/signUp", user.SignUp)   // 注册
		userApi.POST("/signOut", user.SignOut) // 登出
	}
	
	ebookApi := engine.Group("/ebook/ebook")
	//userApi.Use(jwt.JWT())
	{
		ebookApi.POST("/add", ebook.Add) // 上传电子书
		ebookApi.POST("/get", ebook.Get)   // 查看电子图详情
		ebookApi.POST("/getList", ebook.GetList) // 获取电子书列表
		ebookApi.POST("/download", ebook.Download) // 下载电子书
		ebookApi.POST("/view", ebook.View) // 预览
	}
	
	
	//categoryApi := engine.Group("/ebook/category")
	////categoryApi.MaxMultipartMemory = 8 << 20  // 8 MiB		设置上传文件大小
	//categoryApi.Use(jwt.JWT())
	//{
	//	categoryApi.POST("/getList", user.SignUp)   // 获取类别列表
	//	categoryApi.POST("/upload", user.SignOut) // 上传电子书
	//	categoryApi.POST("/view", user.SignOut) // 预览电子书或图片
	//}
	//return r
}