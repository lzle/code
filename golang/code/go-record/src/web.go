package src

import (
	"github.com/gin-gonic/gin"
	"go-record/core"
	"go-record/utils"
	"strings"
)

type Web struct {
}

func (web *Web) download(c *gin.Context) {
	filePath := c.Param("param")
	core.LOGGER.Info("request url %s", filePath)
	if len(filePath) >= 20 {
		for _, dir := range core.CONFIG.RecordDirs() {
			fullPath := dir + filePath
			core.LOGGER.Info(fullPath)
			if utils.PathExists(fullPath) {
				if strings.HasSuffix(fullPath, ".wav") || strings.HasSuffix(fullPath, ".mp3") {
					c.Writer.Header().Add("Content-Type", "audio/mpeg")
				}
				core.LOGGER.Info("download file %s", fullPath)
				c.File(fullPath)
				return
			}
		}
	}
	c.String(404, "404: Not Found")
}

func (web *Web) Run() {
	router := gin.Default()
	router.GET("/*param", web.download)
	router.POST("/*param", web.download)
	if core.CONFIG.EnTls() {
		router.RunTLS(core.CONFIG.HttpAddr(), core.CONFIG.Crt(), core.CONFIG.Key())
	} else {
		router.Run(core.CONFIG.HttpAddr())
	}
}
