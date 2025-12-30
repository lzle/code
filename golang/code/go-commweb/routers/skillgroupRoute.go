package routers

import (
	"github.com/gin-gonic/gin"
	"go-commweb/module/skillGroupHandle"
	)


func skillGroupRoute(c *gin.Engine) {
	sgGroup := c.Group("/group")
	{
		// 技能组预览接口
		sgGroup.POST("preview", skillGroupHandle.SkillGroupPreview)

		// 技能组详情接口
		sgGroup.POST("detail", skillGroupHandle.SkillGroupDetail)
	}
}

func skillGroupInterface (c *gin.Engine) {
	skillGroupRoute(c)
}

