package routers

import (
	"github.com/gin-gonic/gin"
	"go-commweb/module/agentTranslateHandle"
)

func agentTranslateRoute(c *gin.Engine) {
	c.POST("/agent/translate", agentTranslateHandle.ExecuteAgentTranslate)
}

func agentTranslateInterface(c *gin.Engine)  {
	agentTranslateRoute(c)
}
