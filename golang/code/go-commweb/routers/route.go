package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	//router.Use(Cors.Next()) //允许跨域，如果nginx已经开启跨域，请注释该行

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "connect success")
	})

	callInInterface(router)

	taskInterFace(router)

	skillGroupInterface(router)

	callTaskInterface(router)

	agentTranslateInterface(router)

	mediaFileInterface(router)

	return router
}
