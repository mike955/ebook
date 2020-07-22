package response

import "github.com/gin-gonic/gin"

func OK(ctx *gin.Context, data interface{})  {
	ctx.JSON(200, map[string]interface{}{
		"errno": 0,
		"errmsg": "",
		"data": data,
	})
	return
}

func Error(ctx *gin.Context, infos ...string) {
	ErrorCode := infos[0]
	Errno, Errmsg := Code(ErrorCode)
	if len(infos) > 1 {
		Errmsg = infos[1]
	}
	ctx.JSON(200, map[string]interface{}{
		"errno": Errno,
		"errmsg": Errmsg,
		"data": "",
	})
	return
}