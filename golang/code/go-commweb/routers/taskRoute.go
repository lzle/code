package routers

import (
	"github.com/gin-gonic/gin"
	"go-commweb/module/taskHandle"
)

func taskRoute (c *gin.Engine)  {
	taskApi := c.Group("/robot")
	{
		taskApi.POST("start", taskHandle.TaskStart)
		taskApi.POST("stop", taskHandle.TaskStop)
		taskApi.POST("listen", taskHandle.TaskListen)
		taskApi.POST("intercept", taskHandle.TaskIntercept)
	}
}


func taskInterFace(c *gin.Engine) {
	taskRoute(c)
	taskHandle.TaskCls = new(taskHandle.Task)
	taskHandle.TaskCls.TaskQueue = make(chan *taskHandle.TaskParam, 100)
	taskHandle.TaskCls.CloseTask = make(chan interface{}, 1)
	go taskHandle.TaskCls.Execute()
}