package taskHandle

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-commweb/common"
	"go-commweb/global"
	log "go-commweb/log"
	"io/ioutil"
)

var TaskCls *Task

type TaskParam struct {
	ActionId  string  `json:"actionId" binding:"required"`
	CompId    string  `json:"compId" binding:"required"`
	TaskId    string  `json:"taskId" binding:"required"`
	Auth      string  `json:"auth"`
	Mode      string  `json:"mode"`
	CallId    string  `json:"callId"`
	From      string  `json:"from"`
	FromType  string  `json:"fromType"`
	To        string  `json:"to"`
	ToType    string  `json:"toType"`
}


func TaskStart (c *gin.Context) {
	taskParam := new(TaskParam)
	data, err := ioutil.ReadAll(c.Request.Body)
	if err == nil {
		err = json.Unmarshal(data, &taskParam)
		if err != nil {
			log.LOGGER.Error("start task param type is error [%v]", err)
			common.ResponseJson(c, global.RequestSuccess, global.JsonParseError, "json解析失败", "")
			return
		}
	}else{
		log.LOGGER.Error("start task param type is error [%v]", err)
		common.ResponseJson(c, global.RequestSuccess, global.ParamReadError, "读取参数失败", "")
		return
	}
	res := paramDetection(taskParam)
	if !res {
		log.LOGGER.Error("start task param loss")
		common.ResponseJson(c, global.RequestSuccess, global.ParamMissingError, "参数不全", "")
		return
	}
	taskParam.Mode = startTask
	TaskCls.TaskQueue <- taskParam
	common.ResponseJson(c, global.RequestSuccess, global.RequestSuccess, "ok", "")
}

func TaskStop (c *gin.Context) {
	taskParam := new(TaskParam)
	data, err := ioutil.ReadAll(c.Request.Body)
	if err == nil {
		err = json.Unmarshal(data, &taskParam)
		if err != nil {
			log.LOGGER.Error("stop task param type is error [%v]", err)
			common.ResponseJson(c, global.RequestSuccess, global.JsonParseError, "json解析失败", "")
			return
		}
	}else{
		log.LOGGER.Error("stop task param type is error [%v]", err)
		common.ResponseJson(c, global.RequestSuccess, global.ParamReadError, "读取参数失败", "")
		return
	}
	res := paramDetection(taskParam)
	if !res {
		log.LOGGER.Error("stop task param loss")
		common.ResponseJson(c, global.RequestSuccess, global.ParamMissingError, "参数不全", "")
		return
	}
	taskParam.Mode = stopTask
	TaskCls.TaskQueue <- taskParam
	common.ResponseJson(c, global.RequestSuccess, global.RequestSuccess, "ok", "")
}

func TaskListen (c *gin.Context) {
	taskParam := new(TaskParam)
	data, err := ioutil.ReadAll(c.Request.Body)
	if err == nil {
		err = json.Unmarshal(data, &taskParam)
		if err != nil {
			log.LOGGER.Error("listen task param type is error [%v]", err)
			common.ResponseJson(c, global.RequestSuccess, global.JsonParseError, "json解析失败", "")
			return
		}
	}else {
		log.LOGGER.Error("listen task param type is error [%v]", err)
		common.ResponseJson(c, global.RequestSuccess, global.ParamReadError, "读取参数失败", "")
		return
	}
	res := dealParamDetection(taskParam)
	if !res {
		log.LOGGER.Error("listen task param loss")
		common.ResponseJson(c, global.RequestSuccess, global.ParamMissingError, "参数不全", "")
		return
	}
	taskParam.Mode = listenTask
	TaskCls.TaskQueue <- taskParam
	common.ResponseJson(c, global.RequestSuccess, global.RequestSuccess, "ok", "")
}

func TaskIntercept (c *gin.Context) {
	taskParam := new(TaskParam)
	data, err := ioutil.ReadAll(c.Request.Body)
	if err == nil {
		err = json.Unmarshal(data, &taskParam)
		if err != nil {
			log.LOGGER.Error("Intercept task param type is error [%v]", err)
			common.ResponseJson(c, global.RequestSuccess, global.JsonParseError, "json解析失败", "")
			return
		}
	}else {
		log.LOGGER.Error("Intercept task param type is error [%v]", err)
		common.ResponseJson(c, global.RequestSuccess, global.ParamReadError, "读取参数失败", "")
		return
	}
	res := dealParamDetection(taskParam)
	if !res {
		log.LOGGER.Error("Intercept task param loss")
		common.ResponseJson(c, global.RequestSuccess, global.ParamMissingError, "参数不全", "")
		return
	}
	taskParam.Mode = interceptTask
	TaskCls.TaskQueue <- taskParam
	common.ResponseJson(c, global.RequestSuccess, global.RequestSuccess, "ok", "")
}

func CallTaskStart(c *gin.Context) {
	actionId := c.PostForm("actionId")
	company := c.PostForm("company")
	taskid := c.PostForm("taskid")
	auth := c.PostForm("auth")

	taskParam := new(TaskParam)
	taskParam.CompId = company
	taskParam.TaskId = taskid
	taskParam.ActionId = actionId
	taskParam.Auth = auth

	//data, err := ioutil.ReadAll(c.Request.Body)
	//if err == nil {
	//	err = json.Unmarshal(data, &taskParam)
	//	if err != nil {
	//		log.LOGGER.Error("start call task param type is error [%v]", err)
	//		common.ResponseJson(c, global.RequestSuccess, global.JsonParseError, "json解析失败", "")
	//		return
	//	}
	//}else{
	//	log.LOGGER.Error("start call task param type is error [%v]", err)
	//	common.ResponseJson(c, global.RequestSuccess, global.ParamReadError, "读取参数失败", "")
	//	return
	//}
	res := callTaskParamDetection(taskParam)
	if !res {
		log.LOGGER.Error("start call task param loss")
		common.ResponseJson(c, global.RequestSuccess, global.ParamMissingError, "参数不全", "")
		return
	}

	//  md5 加密 验证
	md5Encryption := common.CreateMD5(taskParam.CompId, taskParam.TaskId)
	log.LOGGER.Info("create md5 compid [%s] taskid [%s] md5 [%s]",taskParam.CompId, taskParam.TaskId, md5Encryption)
	if md5Encryption != taskParam.Auth {
		common.ResponseJson(c, global.RequestSuccess, global.AuthEncryptError, "秘钥错误", "")
		return
	}

	taskParam.Mode = startCallTask
	TaskCls.TaskQueue <- taskParam
	common.ResponseJson(c, global.RequestSuccess, global.RequestSuccess, "ok", "")
}

func CallTaskStop(c *gin.Context) {
	actionId := c.PostForm("actionId")
	company := c.PostForm("company")
	taskid := c.PostForm("taskid")
	auth := c.PostForm("auth")

	taskParam := new(TaskParam)
	taskParam.CompId = company
	taskParam.TaskId = taskid
	taskParam.ActionId = actionId
	taskParam.Auth = auth

	//data, err := ioutil.ReadAll(c.Request.Body)
	//if err == nil {
	//	err = json.Unmarshal(data, &taskParam)
	//	if err != nil {
	//		log.LOGGER.Error("stop call task param type is error [%v]", err)
	//		common.ResponseJson(c, global.RequestSuccess, global.JsonParseError, "json解析失败", "")
	//		return
	//	}
	//}else{
	//	log.LOGGER.Error("stop call task param type is error [%v]", err)
	//	common.ResponseJson(c, global.RequestSuccess, global.ParamReadError, "读取参数失败", "")
	//	return
	//}
	res := callTaskParamDetection(taskParam)
	if !res {
		log.LOGGER.Error("start call task param loss")
		common.ResponseJson(c, global.RequestSuccess, global.ParamMissingError, "参数不全", "")
		return
	}

	//  md5 加密 验证
	md5Encryption := common.CreateMD5(taskParam.CompId, taskParam.TaskId)
	log.LOGGER.Info("create md5 compid [%s] taskid [%s] md5 [%s]",taskParam.CompId, taskParam.TaskId, md5Encryption)
	if md5Encryption != taskParam.Auth {
		common.ResponseJson(c, global.RequestSuccess, global.AuthEncryptError, "秘钥错误", "")
		return
	}

	taskParam.Mode = stopCallTask
	TaskCls.TaskQueue <- taskParam
	common.ResponseJson(c, global.RequestSuccess, global.RequestSuccess, "ok", "")
}

func paramDetection (task *TaskParam) (res bool) {
	log.LOGGER.Info("recv task message actionId [%s] compId [%s] taskId [%s]", task.ActionId, task.CompId, task.TaskId)
	if task.ActionId != "" && task.CompId != "" && task.TaskId != ""{
		return true
	}
	return false
}

func callTaskParamDetection (callTask *TaskParam) (res bool) {
	log.LOGGER.Info("recv task message actionId [%s] compId [%s] taskId [%s] auth [%s]", callTask.ActionId, callTask.CompId, callTask.TaskId, callTask.Auth)
	if callTask.ActionId != "" && callTask.CompId != "" && callTask.TaskId != "" && callTask.Auth != ""{
		return true
	}
	return false
}

func dealParamDetection(taskDeal *TaskParam) (res bool)  {
	log.LOGGER.Info("recv task message actionId [%s] compId [%s] taskId [%s] callid [%s] from [%s, %s] to [%s, %s]", taskDeal.ActionId, taskDeal.CompId, taskDeal.TaskId, taskDeal.CallId, taskDeal.From, taskDeal.FromType, taskDeal.To, taskDeal.ToType)
	if taskDeal.ActionId != "" && taskDeal.CompId != "" && taskDeal.TaskId != "" && taskDeal.CallId != "" && taskDeal.From != "" && taskDeal.FromType != "" && taskDeal.To != "" && taskDeal.ToType != ""{
		return true
	}
	return false
}

