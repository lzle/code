package agentTranslateHandle

import (
	"github.com/gin-gonic/gin"
	"go-commweb/amqp"
	"go-commweb/common"
	"go-commweb/global"
	log "go-commweb/log"
)

func ExecuteAgentTranslate(c *gin.Context) {
	action := c.PostForm("action")

	log.LOGGER.Info("agent translate get param is action [%s]", action)
	if action == "set" {
		setAgentTranslate(c)
	}else if action == "cancel" {
		cancelAgentTranslate(c)
	}else{
		log.LOGGER.Error("action[%s] can not match", action)
		common.ResponseJson(c, global.RequestSuccess, global.ParamMissingError, "参数不全", "")
	}
}

func setAgentTranslate(c *gin.Context)  {
	actionId := c.PostForm("actionId")
	agentId := c.PostForm("agentId")
	compId := c.PostForm("compId")
	log.LOGGER.Info("agent set translate get param is agentId [%s] compId [%s]", agentId, compId)
	if !verifyArguments(actionId, agentId, compId) {
		common.ResponseJson(c, global.RequestSuccess, global.ParamMissingError, "参数不全", "")
		return
	}
	ctiServerId := getCtiServerId(compId, agentId)
	if ctiServerId == "" {
		log.LOGGER.Error("agent [%s] compId [%s] can not get cti serverid", agentId, compId)
		common.ResponseJson(c, global.RequestSuccess, global.GetServerIdError, "serverId获取失败", "")
		return
	}
	amqp.SendSetAgentTranslate(ctiServerId, agentId, compId)
	common.ResponseJson(c, global.RequestSuccess, global.RequestSuccess, "ok", "")
}

func cancelAgentTranslate(c *gin.Context)  {
	actionId := c.PostForm("actionId")
	agentId := c.PostForm("agentId")
	compId := c.PostForm("compId")

	log.LOGGER.Info("agent cancel translate get param is agentId [%s] compId [%s]", agentId, compId)
	if !verifyArguments(actionId, agentId, compId) {
		common.ResponseJson(c, global.RequestSuccess, global.ParamMissingError, "参数不全", "")
		return
	}
	ctiServerId := getCtiServerId(compId, agentId)
	if ctiServerId == "" {
		log.LOGGER.Error("agent [%s] compId [%s] can not get cti serverid", agentId, compId)
		common.ResponseJson(c, global.RequestSuccess, global.GetServerIdError, "serverId获取失败", "")
		return
	}
	amqp.SendCancelAgentTranslate(ctiServerId, agentId, compId)
	common.ResponseJson(c, global.RequestSuccess, global.RequestSuccess, "ok", "")
}

func getCtiServerId(compId, agentId string) (serverId string) {
	serverId = getAgentServerId(compId, agentId)
	if serverId == "" {
		serverId = getServerId(compId, "3", agentId)
	}
	return serverId
}

func verifyArguments(actionId, agentId, compId string) bool {
	if actionId == "" || agentId == "" || compId == "" {
		return false
	}
	return true
}
