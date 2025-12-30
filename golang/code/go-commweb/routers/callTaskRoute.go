package routers

import (
	"github.com/gin-gonic/gin"
	"go-commweb/module/taskHandle"
)

func callTaskRoute(c *gin.Engine) {
	taskApi := c.Group("/task")
	{
		taskApi.POST("start", taskHandle.CallTaskStart)
		taskApi.POST("stop", taskHandle.CallTaskStop)
	}
}

func callTaskInterface(c *gin.Engine) {
	callTaskRoute(c)
	taskHandle.TaskCls = new(taskHandle.Task)
	taskHandle.TaskCls.TaskQueue = make(chan *taskHandle.TaskParam, 100)
	taskHandle.TaskCls.CloseTask = make(chan interface{}, 1)
	go taskHandle.TaskCls.Execute()
}