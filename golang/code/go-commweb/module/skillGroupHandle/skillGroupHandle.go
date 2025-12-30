package skillGroupHandle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-commweb/common"
	"go-commweb/global"
	log "go-commweb/log"
)

func SkillGroupPreview (c *gin.Context)  {
	actionId := c.PostForm("actionId")
	company := c.PostForm("company")
	groupid := c.PostFormArray("groupid")

	log.LOGGER.Info("skill group preview recv param actionId [%s] compId [%s] groupId [%s]", actionId, company, groupid)

	res := paramDetection(actionId, company, groupid)
	if !res {
		common.ResponseJson(c, global.RequestSuccess, global.ParamMissingError, "参数不全", "")
		return
	}
	var sgArray []map[string]string
	for _, id := range groupid {
		sgInfo := make(map[string]string)
		keys := fmt.Sprintf("%s:%s", company, id)
		log.LOGGER.Info("group search key is %s", keys)
		numberOfWaiting := redisHget(keys, "number_of_waiting")
		maximumWaitingDatetime := redisHget(keys, "maximum_waiting_datetime")
		sgInfo["number_of_waiting"] = numberOfWaiting
		sgInfo["maximum_waiting_datetime"] = maximumWaitingDatetime
		sgInfo["groupid"] = id
		sgArray = append(sgArray, sgInfo)
	}
	common.ResponseJson(c, global.RequestSuccess, global.RequestSuccess, "ok", sgArray)
}

func SkillGroupDetail (c *gin.Context) {
	actionId := c.PostForm("actionId")
	company := c.PostForm("company")
	groupid := c.PostFormArray("groupid")

	log.LOGGER.Info("skill group detail recv param actionId [%s] compId [%s] groupId [%s]", actionId, company, groupid)

	res := paramDetection(actionId, company, groupid)
	if !res {
		common.ResponseJson(c, global.RequestSuccess, global.ParamMissingError, "参数不全", "")
		return
	}
	sgDetail := make(map[string][]map[string]string)

	for _, id := range groupid {
		keyBuild := fmt.Sprintf("%s:%s:*", company, id)
		getKeys := redisKeys(keyBuild)
		perGroupDetail := []map[string]string{}
		for _, keys := range getKeys.([]interface{}) {
			perCustomerDetail := make(map[string]string)
			keysConvered := string(keys.([]uint8))
			if !common.StringContain(keysConvered, "24h") {
				deviceId :=  redisHget(keysConvered, "deviceid")
				callIdDatetime :=  redisHget(keysConvered, "callin_datetime")
				waitingDatetime :=  redisHget(keysConvered, "waiting_datetime")
				waitingPosition :=  redisHget(keysConvered, "waiting_position")

				perCustomerDetail["deviceid"] = deviceId
				perCustomerDetail["callin_datetime"] = callIdDatetime
				perCustomerDetail["waiting_datetime"] = waitingDatetime
				perCustomerDetail["waiting_position"] = waitingPosition

				perGroupDetail = append(perGroupDetail, perCustomerDetail)
			}
		}
		sgDetail[id] = perGroupDetail
	}
	common.ResponseJson(c, global.RequestSuccess, global.RequestSuccess, "ok", sgDetail)
}

func paramDetection(actionId, compId string, groupId []string) bool {
	if actionId != "" && compId != "" && len(groupId) > 0 {
		return true
	}
	return false
}