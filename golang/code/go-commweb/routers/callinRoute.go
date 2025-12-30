package routers

import (
	"github.com/gin-gonic/gin"
	"go-commweb/module/callInHandle"
)


func callInRoute(r *gin.Engine)  {
	callInApi := r.Group("/callin")
	{
		// 申请
		callInApi.POST("apply", callInHandle.CallInApply)

		// 释放
		callInApi.POST("release", callInHandle.CallRelease)
	}
}

func callInInterface(r *gin.Engine) {
	callInRoute(r)
	go callInHandle.ClearCalls()
	go callInHandle.UpdateCompany()
	go callInHandle.Monitor()
}