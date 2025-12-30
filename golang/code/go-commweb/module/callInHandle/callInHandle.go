package callInHandle

import (
	"github.com/gin-gonic/gin"
	"go-commweb/common"
	"go-commweb/global"
	log "go-commweb/log"
)

func CallInApply (c *gin.Context) {
	actionId := c.PostForm("actionId")
	compId := c.PostForm("compId")
	callId := c.PostForm("callId")
	caller := c.PostForm("caller")
	callee := c.PostForm("callee")

	res := paramDetection(actionId, compId, callId, caller, callee)
	if !res {
		common.ResponseJson(c, global.RequestSuccess, global.ParamMissingError , "参数不全", "")
		return
	}

	var company *company
	company = getCompany(compId)
	if company == nil {
		company = newCompany(compId)
	}
	if company.canCall() {
		company.callContain = append(company.callContain, callId)
		newCall(callId, compId, caller, callee)
		log.LOGGER.Info("compid[%s] callid[%s] can call limit[%d] currentcalls[%d]", compId, callId, company.maxLimit, len(company.callContain))
		common.ResponseJson(c, global.RequestSuccess, global.RequestSuccess , "OK", 1)
	}else{
		log.LOGGER.Info("callid[%s] compid[%s] over limit[%s] currentcalls[%s]", callId, compId, company.maxLimit, company.callContain)
		common.ResponseJson(c, global.RequestSuccess, global.RequestSuccess , "当前并发限制已满", 0)
	}
}

func CallRelease (c *gin.Context) {
	actionId := c.PostForm("actionId")
	compId := c.PostForm("compId")
	callId := c.PostForm("callId")
	caller := c.PostForm("caller")
	callee := c.PostForm("callee")

	res := paramDetection(actionId, compId, callId, caller, callee)

	if !res {
		common.ResponseJson(c, global.RequestSuccess, global.ParamMissingError , "缺少参数", "")
		return
	}

	call := getCall(callId)
	if call == nil {
		log.LOGGER.Error("call [%s] not exist", callId)
		common.ResponseJson(c, global.RequestSuccess, global.CallNotExistError , "释放的call没有申请", "")
		return
	}
	if caller == call.caller {
		delCall(callId)
		company := getCompany(compId)
		company.callContain = common.SliceDelete(company.callContain, callId)
		log.LOGGER.Info("compId [%s] callid [%s] release caller [%s]", compId, callId, caller)
		common.ResponseJson(c, global.RequestSuccess, global.RequestSuccess , "ok", "")
	} else {
		log.LOGGER.Error(" callId [%s] caller [%s] not match", callId, caller)
		common.ResponseJson(c, global.RequestSuccess, global.CallerNotMatchError , "主叫不匹配", "")
	}
}

func paramDetection (actionId, compId, callId, caller, callee string) (res bool) {
	log.LOGGER.Info("recv message actionId [%s], compId [%s], callId [%s], caller [%s], callee [%s]", actionId, compId, callId, caller, callee)
	if actionId != "" && compId != "" && callId != "" && caller != "" && callee != "" {
		return true
	}
	return false
}
