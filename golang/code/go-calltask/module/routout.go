package module

import (
	"fmt"
	"go-calltask/mysql"
	"go-calltask/redis"
	log "go-calltask/log"
	"strings"
)

type ShowNumber struct {
	Number string `json:"number"`
}

type ExtensionShowNumber struct {
	ExtensionNumber string `json:"callerid"`
}

type Extension struct {
	LoginExt string `json:"loginext"`
}

type OutlineTrunkId struct {
	Prefix string `json:"prefix"`
	TrunkId string `json:"serverid"`
}

type BusinessServerId struct {
	ServerId string `json:"serverid"`
}

func getAgentCtiServerid(compid string, agentid string) (serverid string) {
	name := fmt.Sprintf("agent:%s:%s", compid, agentid)
	redisdb := redis.GetRedisDb()
	serverid = redisdb.HashGet(name, "serverid")
	return serverid
}

// 获取外显号码
func getShowNum(deviceId string, deviceType string, peerId string, peerType string, compId string) (showNum string) {
	if deviceType == "1" || deviceType == "2"{
		if peerType == "1"{
			showNum = getLoginExtension(peerId, compId)
		}else{
			showNum = peerId
		}
	}else{
		if peerType == "1"{
			showNum = getAgentShowNum(peerId, compId)
		}else if peerType == "2"{
			showNum = getExtensionShowMum(peerId, compId)
		}else{
			showNum = getCompanyShowNum(compId)
		}

	}
	return showNum
}

// 获取企业外显
func getCompanyShowNum(compId string) (showNum string) {
	sql := "select number from tx_access_num where enable=1 and isdefault=1 and compid=?"
	DB := mysql.GetMysqlDb().DB
	rows, err := DB.Query(sql, compId)
	if err != nil {
		return
	}
	defer func() {
		if rows != nil {
			rows.Close() //可以关闭掉未scan连接一直占用
		}
	}()
	for rows.Next() {
		_showNum := new(ShowNumber)
		err = rows.Scan(&_showNum.Number) //不scan会导致连接不释放
		if err != nil {
			log.LOGGER.Error("scan failed %s", err.Error())
			return ""
		}
		if _showNum.Number != ""{
			return _showNum.Number
		}
	}
	return ""
}

// 根据坐席id获取外呼外显号码
func getAgentShowNum(agentId string, compId string)  (showNum string){
	loginExt := getLoginExtension(agentId, compId)
	if loginExt != ""{
		showNum = getExtensionShowMum(loginExt, compId)
		return
	}
	return ""
}

// 根据分机id获取外呼外显号码
func getExtensionShowMum(extension string, compId string) (showNum string){
	sql := "select callerid from tx_extension where name=?"
	DB := mysql.GetMysqlDb().DB
	rows, err := DB.Query(sql, extension)
	if err != nil {
		return
	}
	defer func() {
		if rows != nil {
			rows.Close() //可以关闭掉未scan连接一直占用
		}
	}()
	for rows.Next() {
		_showNum := new(ExtensionShowNumber)
		err = rows.Scan(&_showNum.ExtensionNumber) //不scan会导致连接不释放
		if err != nil {
			log.LOGGER.Error("scan failed %s", err.Error())
			return ""
		}
		if _showNum.ExtensionNumber != ""{
			return _showNum.ExtensionNumber
		}else{
			showNum = getCompanyShowNum(compId)
			return
		}
	}
	return ""
}

func getTrunkId(deviceId string, deviceType string, compId string) (trunkId string, webrtc string, _deviceId string) {
	if deviceType == "1" {
		extension := getLoginExtension(deviceId, compId)
		if extension != "" {
			trunkId,webrtc = getExtensionTrunkid(extension,compId)
			return  trunkId,webrtc,extension
		}
	} else if deviceType == "2" {
		trunkId,webrtc = getExtensionTrunkid(deviceId,compId)
		return  trunkId,webrtc,deviceId
	} else {
		trunkId = getOutlineTrunkId(deviceId,compId)
		return  trunkId,webrtc,deviceId
	}
	return
}

func getLoginExtension(deviceId string, compId string) (loginext string) {
	sql := "select loginext from tx_agent where agentid=? and compid=?"
	DB := mysql.GetMysqlDb().DB
	row := DB.QueryRow(sql, deviceId, compId)
	extension := new(Extension)
	if err := row.Scan(&extension.LoginExt); err != nil {
		log.LOGGER.Error("scan failed %s", err.Error())
		return
	}
	loginext = extension.LoginExt
	return
}

func getExtensionTrunkid(extension string, compId string) (serverid string, webrtc string) {
	name := fmt.Sprintf("extension_reg:%s", extension)
	redisdb := redis.GetRedisDb()
	serverid = redisdb.HashGet(name, "serverid")
	webrtc = redisdb.HashGet(name, "webrtc")
	if serverid == ""{
		log.LOGGER.Error("get extension_reg error,name[%s] not register", extension)
	}
	return
}

func getOutlineTrunkId(deviceId string, compId string) (serverid string){
	sql := "select prefix,serverid from pbx.tx_route_out order by priority"
	DB := mysql.GetMysqlDb().DB
	rows, err := DB.Query(sql)
	if err != nil {
		return
	}
	defer func() {
		if rows != nil {
			rows.Close() //可以关闭掉未scan连接一直占用
		}
	}()
	for rows.Next() {
		outlineTrunkId := new(OutlineTrunkId)
		err = rows.Scan(&outlineTrunkId.Prefix,&outlineTrunkId.TrunkId) //不scan会导致连接不释放
		if err != nil {
			log.LOGGER.Error("scan failed %s", err.Error())
			return
		}
		if outlineTrunkId.Prefix != ""{
			ok :=strings.HasPrefix(deviceId,outlineTrunkId.Prefix)
			if ok {
				return outlineTrunkId.TrunkId
			}
		}
	}
	return
}

func getEslServerIds(compId string)(eslServerIds []string){
	log.LOGGER.Info("get esl serviceId from tx_route_business bustype[17] compid[%s]", compId)
	sql := "SELECT serverid FROM pbx.tx_route_business where bustype = 17  and compid = ? and enable = 1"
	DB := mysql.GetMysqlDb().DB
	rows, err := DB.Query(sql,compId)
	if err != nil {
		return
	}
	defer func() {
		if rows != nil {
			rows.Close() //可以关闭掉未scan连接一直占用
		}
	}()
	for rows.Next() {
		businessServerId := new(BusinessServerId)
		err = rows.Scan(&businessServerId.ServerId) //不scan会导致连接不释放
		if err != nil {
			log.LOGGER.Error("scan failed %s", err.Error())
			return
		}
		if businessServerId.ServerId != ""{
			eslServerIds = append(eslServerIds,businessServerId.ServerId)
		}
	}
	if len(eslServerIds) > 0 {
		return
	} else {
		sql := "SELECT serverid FROM pbx.tx_route_business where bustype = 17 and compid ='' and enable = 1"
		_rows, err := DB.Query(sql)
		if err != nil {
			return
		}
		defer func() {
			if _rows != nil {
				_rows.Close() //可以关闭掉未scan连接一直占用
			}
		}()
		for _rows.Next() {
			businessServerId := new(BusinessServerId)
			err = _rows.Scan(&businessServerId.ServerId) //不scan会导致连接不释放
			if err != nil {
				log.LOGGER.Error("scan failed %s", err.Error())
				return
			}
			if businessServerId.ServerId != ""{
				eslServerIds = append(eslServerIds,businessServerId.ServerId)
			}
		}
	}
	return
}

func getSmsServerIds(compId string)(smsServerIds []string){
	log.LOGGER.Info("get sms serviceId from tx_route_business bustype[24] compid[%s]", compId)
	sql := "SELECT serverid FROM pbx.tx_route_business where bustype = 24  and compid = ? and enable = 1"
	DB := mysql.GetMysqlDb().DB
	rows, err := DB.Query(sql,compId)
	if err != nil {
		return
	}
	defer func() {
		if rows != nil {
			rows.Close() //可以关闭掉未scan连接一直占用
		}
	}()
	for rows.Next() {
		businessServerId := new(BusinessServerId)
		err = rows.Scan(&businessServerId.ServerId) //不scan会导致连接不释放
		if err != nil {
			log.LOGGER.Error("scan failed %s", err.Error())
			return
		}
		if businessServerId.ServerId != ""{
			smsServerIds = append(smsServerIds,businessServerId.ServerId)
		}
	}
	if len(smsServerIds) > 0 {
		return
	} else {
		sql := "SELECT serverid FROM pbx.tx_route_business where bustype = 24 and compid ='' and enable = 1"
		_rows, err := DB.Query(sql)
		if err != nil {
			return
		}
		defer func() {
			if _rows != nil {
				_rows.Close() //可以关闭掉未scan连接一直占用
			}
		}()
		for _rows.Next() {
			businessServerId := new(BusinessServerId)
			err = _rows.Scan(&businessServerId.ServerId) //不scan会导致连接不释放
			if err != nil {
				log.LOGGER.Error("scan failed %s", err.Error())
				return
			}
			if businessServerId.ServerId != ""{
				smsServerIds = append(smsServerIds,businessServerId.ServerId)
			}
		}
	}
	return
}

// 坐席加入任务写redis
func setAgentCti(compId string, agentId string, taskId string, ctiServerid string){
	redisdb := redis.GetLocalRedisDb()
	key := fmt.Sprintf("%s:%s", compId, agentId)
	content := make(map[string]string)
	content["taskid"] = taskId
	content["serverid"] = ctiServerid
	ret := redisdb.HashMSet(key, content)
	log.LOGGER.Info("agent %s join task %s compid %s set redis result[%v]", agentId, taskId, compId, ret)
}

// 坐席退出任务删除redis数据
func delAgentCti(compId string, agentId string, taskId string){
	redisdb := redis.GetLocalRedisDb()
	key := fmt.Sprintf("%s:%s", compId, agentId)
	ret := redisdb.Del(key)
	log.LOGGER.Info("agent %s exit task %s compid %s delete redis result[%v]", agentId, taskId, compId, ret)
}









