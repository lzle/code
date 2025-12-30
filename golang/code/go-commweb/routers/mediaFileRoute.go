package routers

import (
	"github.com/gin-gonic/gin"
	"go-commweb/module/mediaFileHandle"
)


func mediaFileRoute(c *gin.Engine) {
	c.POST("/media", mediaFileHandle.MediaFileAdd)
	c.DELETE("/media", mediaFileHandle.MediaFileDelete)
}


func mediaFileInterface(c *gin.Engine) {
	mediaFileRoute(c)
}
