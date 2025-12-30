package module

import (
	"fmt"
	"go-calltask/amqp"
	"go-calltask/config"
	log "go-calltask/log"
	"go-calltask/tools"
)

const (
	STATE_INIT = 0    // 初始化
	STATE_HOLD = 1    // 保持
	STATE_ANSWER = 2  // 应答
	STATE_HANGUP = 3  // 挂断
)

type Member struct {
	// 主叫
	DeviceId   string
	DeviceType string
	// 被叫
	PeerId   string
	PeerType string
	// 成员标示
	MemId string
	// 通话标示
	CallId string
	// 企业id
	CompId string
	// 任务id
	TaskId string
	// 播放id
	PlayId string
	// first second
	MemType string
	// device标示
	CusId string
	// 随路数据
	TransParam string
	// 外显
	ShowNum string
	// 底层serverId
	EslServerId string
	// member 通话状态
	CallState int
	// 监听销毁通话标识
	DestoryCall bool
	// 状态回调
	StateCallback func(member *Member, result int)
	// 放音回调
	PlayCallback map[string]func()
}

func (mem *Member) createVariables()(variables *amqp.CdrVariables){
	variables = new(amqp.CdrVariables)
	variables.CallerId = mem.ShowNum
	variables.RequestId = mem.TransParam
	variables.CallType = "2"
	variables.CallId = mem.CallId
	variables.CompId = mem.CompId
	variables.TaskId = mem.TaskId
	if mem.DeviceType == "3" {
		variables.GroupType = "2"
		variables.Callee = mem.DeviceId
		variables.Dh = mem.ShowNum
	}else if mem.DeviceType == "1" {
		variables.GroupType = "1"
		variables.AgentId = mem.DeviceId
		variables.Callee = fmt.Sprintf("%s%s",mem.CompId,mem.DeviceId)
	}
	if mem.PeerType == "1" {
		variables.AgentId = mem.PeerId
	}
	variables.App = mem.MemType
	variables.Direction = "2"
	variables.Cdrtype = "2"
	variables.Caller = mem.ShowNum
	variables.Serverid = config.AmqpConfig.ServerId
	return variables
}

func (mem *Member) SetRecordPath(recordFile string){
	mem.SendCommandSetRecordPath(recordFile)
}

// 开始呼叫
func (mem *Member) MakeCall(showNum string, bridgeId string) (ret bool) {
	var (
		deviceId string
		caller string
		prefix string
	)
	taskdt := getTaskDt(mem.TaskId)
	compdt := getCompDt(mem.CompId)
	if mem.DeviceType == "3" {
		data := make(map[string]string)
		data["actionId"],_ = tools.NewUUid()
		data["compId"] = mem.CompId
		data["deviceId"] = mem.DeviceId
		if bodyJson := HttpPostForm(config.RecognitionConf.Url,data); bodyJson != nil {
			prefix, _ = bodyJson.Get("prefix").String();
		}
	}
	deviceId = fmt.Sprintf("%s%s", prefix, mem.DeviceId)
	if showNum == "" {
		showNum = getCompanyShowNum(mem.CompId)
	}
	if showNum == ""{
		log.LOGGER.Error("callId[%s] compId[%s] deviceId[%s] showNum is none", mem.CallId, mem.CompId, deviceId)
		return false
	}
	mem.ShowNum = showNum

	trunkid, webrtc, deviceId := getTrunkId(deviceId, mem.DeviceType, mem.CompId)
	log.LOGGER.Info("callId[%s] trunkid[%s] webrtc[%s] deviceId[%s]", mem.CallId, trunkid, webrtc, deviceId)
	if trunkid == "" {
		return false
	}
	if webrtc == "1" {
		webrtc = "true"
	}else {
		webrtc = "false"
	}
	if tools.IsExistInArray(mem.DeviceType,[]string{"1","2"}) && compdt.isAnonymous{
		caller = "anonymous"
	}else {
		caller = showNum
	}
	mem.SendCommandCallout(caller, deviceId, bridgeId, webrtc, trunkid, taskdt.ringingDuration)
	return true
}

func (mem *Member) playMedia(media string, playCallBack func()){
	mem.PlayId,_ = tools.NewUUid()
	mem.SendCommandPlay(mem.PlayId,media)
	mem.PlayCallback[mem.PlayId] = playCallBack
}

func (mem *Member) StopPlay(){
	if mem.PlayId != "" {
		delete(mem.PlayCallback, mem.PlayId)
		mem.SendCommandStopPlay(mem.PlayId)
	}
}

func (mem *Member) Answer() {
	if mem.CallState != STATE_ANSWER {
		mem.SendCommandAnswer()
	}
}

func (mem *Member) Hangup() {
	if mem.CallState != STATE_HANGUP{
		mem.SendCommandHangup()
		mem.CallState = STATE_HANGUP
	}
}

func (mem *Member) Intercept(interceptMemId string, bridgeId string, timeout string) {
	showNum := getShowNum(mem.DeviceId,mem.DeviceType,mem.PeerId,mem.PeerType,mem.CompId)
	if showNum == "" || mem.DeviceId == ""{
		log.LOGGER.Error("compId[%s] callId[%s] showNum[%s] is None", mem.CompId,mem.CallId,showNum)
	}else {
		trunkId, webrtc, deviceId := getTrunkId(mem.DeviceId, mem.DeviceType, mem.CompId)
		if webrtc == "1" {
			webrtc = "true"
		}else {
			webrtc = "false"
		}
		mem.ShowNum = showNum
		mem.SendCommandPick(showNum,deviceId,bridgeId,interceptMemId,webrtc,trunkId,timeout)
	}
}

func (mem *Member) Insert(peerMemId string){
	showNum := getShowNum(mem.DeviceId,mem.DeviceType,mem.PeerId,mem.PeerType,mem.CompId)
	if showNum == "" || mem.DeviceId == ""{
		log.LOGGER.Error("compId[%s] callId[%s] showNum[%s] is None", mem.CompId,mem.CallId,showNum)
	}else {
		trunkId, webrtc, deviceId := getTrunkId(mem.DeviceId, mem.DeviceType, mem.CompId)
		if webrtc == "1" {
			webrtc = "true"
		}else {
			webrtc = "false"
		}
		mem.ShowNum = showNum
		mem.SendCommandInsert(showNum,deviceId,peerMemId,webrtc,trunkId,"60")
	}
}

func (mem *Member) Listen(peerMemId string){
	showNum := getShowNum(mem.DeviceId,mem.DeviceType,mem.PeerId,mem.PeerType,mem.CompId)
	if showNum == "" || mem.DeviceId == ""{
		log.LOGGER.Error("compId[%s] callId[%s] showNum[%s] is None", mem.CompId,mem.CallId,showNum)
	}else {
		trunkId, webrtc, deviceId := getTrunkId(mem.DeviceId, mem.DeviceType, mem.CompId)
		if webrtc == "1" {
			webrtc = "true"
		}else {
			webrtc = "false"
		}
		mem.ShowNum = showNum
		mem.SendCommandListen(showNum,deviceId,peerMemId,webrtc,trunkId,"60")
	}
}

func (mem *Member) SendCommandCallout(caller string, deviceId string, bridgeId string, webrtc string, trunkid string, timeout string) {
	variables := mem.createVariables()
	variables.RingTimeout = timeout
	messageCallout :=new(amqp.MessageCallout)
	messageCallout.Action = "callout"
	messageCallout.Caller = caller
	messageCallout.CallId = mem.CallId
	messageCallout.Callee = deviceId
	messageCallout.Timeout = timeout
	messageCallout.MemId = mem.MemId
	messageCallout.BridgeId = bridgeId
	messageCallout.Webrtc = webrtc
	messageCallout.TrunkId = trunkid
	messageCallout.Media = "1"
	messageCallout.AutoAnswer = "false"
	messageCallout.ActionId,_ = tools.NewUUid()
	messageCallout.ServerId = config.AmqpConfig.ServerId
	messageCallout.Variables = variables
	amqp.SendMsgToAmqp(mem.EslServerId,messageCallout)
}

func (mem *Member) SendCommandPlay(playId string,media string){
	messagePlay :=new(amqp.MessagePlay)
	messagePlay.Action = "play"
	messagePlay.CallId = mem.CallId
	messagePlay.MemId = mem.MemId
	messagePlay.PlayId = playId
	messagePlay.Data = media
	messagePlay.DataType = "2"
	messagePlay.ServerId = config.AmqpConfig.ServerId
	messagePlay.ActionId,_ = tools.NewUUid()
	amqp.SendMsgToAmqp(mem.EslServerId,messagePlay)
}

func (mem *Member) SendCommandStopPlay(playId string){
	messageStopPlay :=new(amqp.MessageStopPlay)
	messageStopPlay.Action = "stopplay"
	messageStopPlay.CallId = mem.CallId
	messageStopPlay.MemId = mem.MemId
	messageStopPlay.PlayId = playId
	messageStopPlay.ServerId = config.AmqpConfig.ServerId
	messageStopPlay.ActionId,_ = tools.NewUUid()
	amqp.SendMsgToAmqp(mem.EslServerId,messageStopPlay)
}

func (mem *Member) SendCommandAnswer(){
	messageAnswer :=new(amqp.MessageAnswer)
	messageAnswer.Action = "answer"
	messageAnswer.CallId = mem.CallId
	messageAnswer.MemId = mem.MemId
	messageAnswer.ServerId = config.AmqpConfig.ServerId
	messageAnswer.ActionId,_ = tools.NewUUid()
	amqp.SendMsgToAmqp(mem.EslServerId,messageAnswer)
}

func (mem *Member) SendCommandHangup(){
	messageHangup :=new(amqp.MessageHangup)
	messageHangup.Action = "hangup"
	messageHangup.CallId = mem.CallId
	messageHangup.MemId = mem.MemId
	messageHangup.ServerId = config.AmqpConfig.ServerId
	messageHangup.ActionId,_ = tools.NewUUid()
	amqp.SendMsgToAmqp(mem.EslServerId,messageHangup)
}

func (mem *Member) SendCommandSetRecordPath(recordFile string){
	variables := new(amqp.VariablesRecordPath)
	variables.RecPath = recordFile
	messageSet :=new(amqp.MessageSetRecordPath)
	messageSet.Action = "set"
	messageSet.CallId = mem.CallId
	messageSet.MemId = mem.MemId
	messageSet.ServerId = config.AmqpConfig.ServerId
	messageSet.ActionId,_ = tools.NewUUid()
	messageSet.Variables = variables
	amqp.SendMsgToAmqp(mem.EslServerId,messageSet)
}

func (mem *Member) SendCommandPick(caller string, deviceId string, bridgeId string, interceptMemId string, webrtc string, trunkId string, timeout string){
	variables := mem.createVariables()
	messagePick :=new(amqp.MessagePick)
	messagePick.Action = "pick"
	messagePick.CallId = mem.CallId
	messagePick.BridgeId = bridgeId
	messagePick.MemId = interceptMemId
	messagePick.PeerMemId = ""
	messagePick.PickMemId = mem.MemId
	messagePick.Caller = caller
	messagePick.Callee = deviceId
	messagePick.Webrtc = webrtc
	messagePick.TrunkId = trunkId
	messagePick.Timeout = timeout
	messagePick.ActionId,_ = tools.NewUUid()
	messagePick.ServerId = config.AmqpConfig.ServerId
	messagePick.Variables = variables
	amqp.SendMsgToAmqp(mem.EslServerId,messagePick)
}

func (mem *Member) SendCommandInsert(caller string, deviceId string, peerMemId string, webrtc string, trunkId string, timeout string){
	variables := mem.createVariables()
	messageInsert :=new(amqp.MessageInsert)
	messageInsert.Action = "ins"
	messageInsert.CallId = mem.CallId
	messageInsert.MemId = mem.MemId
	messageInsert.PeerMemId = peerMemId
	messageInsert.Caller = caller
	messageInsert.Callee = deviceId
	messageInsert.Webrtc = webrtc
	messageInsert.TrunkId = trunkId
	messageInsert.Timeout = timeout
	messageInsert.Media = "1"
	messageInsert.AutoAnswer = "false"
	messageInsert.ActionId,_ = tools.NewUUid()
	messageInsert.ServerId = config.AmqpConfig.ServerId
	messageInsert.Variables = variables
	amqp.SendMsgToAmqp(mem.EslServerId,messageInsert)
}

func (mem *Member) SendCommandListen(caller string, deviceId string, peerMemId string, webrtc string, trunkId string, timeout string){
	variables := mem.createVariables()
	messageListen :=new(amqp.MessageListen)
	messageListen.Action = "listen"
	messageListen.CallId = mem.CallId
	messageListen.MemId = mem.MemId
	messageListen.PeerMemId = peerMemId
	messageListen.Caller = caller
	messageListen.Callee = deviceId
	messageListen.Webrtc = webrtc
	messageListen.TrunkId = trunkId
	messageListen.Timeout = timeout
	messageListen.Media = "1"
	messageListen.ActionId,_ = tools.NewUUid()
	messageListen.ServerId = config.AmqpConfig.ServerId
	messageListen.Variables = variables
	amqp.SendMsgToAmqp(mem.EslServerId,messageListen)
}