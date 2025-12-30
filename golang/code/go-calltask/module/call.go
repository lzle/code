package module

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"go-calltask/amqp"
	"go-calltask/cdr"
	"go-calltask/config"
	log "go-calltask/log"
	"go-calltask/tools"
	"sync"
	"time"
)

const (
	CALL_INTERCEPE_MSGID = "304"
	CALL_HANGUP_MSGID    = "305"
	CALL_INSERT_MSGID    = "306"
	CALL_BREAK_MSGID     = "307"
	CALL_LISTEN_MSGID    = "308"
	CALL_RESET_MSGID     = "312"
	CHANSWER             = 1 // 接通
	CHUNANSWER           = 0 // 未接
	CHHANGUP             = 2 // 接通挂断
)

var (
	callIdToCall = make(map[string]*Call)
	mutexCall            sync.RWMutex
)

type Call struct {
	// 企业id
	compId string
	// 通话id
	callId string
	// 任务id
	taskId string
	// 消息队列
	MsgQueue chan *simplejson.Json
	// first成员
	firstMem *Member
	// second成员
	secondMem *Member
	// 特殊成员
	thirdMem *Member
	// memId mem绑定
	memIdToMem map[string]*Member
	// bridgeid
	bridgeId string
	// 底层serverId
	eslServerId string
	// 短信serverId
	smsServerId string
	// 销毁
	closeCall chan interface{}
	// 录音id
	recordId string
	// 主话单
	cdrMaster *cdr.CdrMaster
	// 当前通话发起呼叫使用的接入号
	showNumber string
}

// 接收消息处理
func (call *Call) execute() {
	defer func() {
		if r := recover(); r != nil {
			log.LOGGER.Error("%v", r)
		}
	}()
	for {
		select {
		case message := <-call.MsgQueue:
			{
				log.LOGGER.Info("callId[%s] taskId[%s] call recv message[%v]", call.callId, call.taskId, *message)
				actionType, _ := message.Get("action_type").String()
				if actionType == "callout" {
					call.onRouteEvent(message)
				} else if actionType == "playmedia" {
					call.playWaitMedia(message)
				} else if actionType == "hangup" {
					call.hangup()
				} else if actionType == "cti" {
					call.onCtiEvent(message)
				} else if actionType == "ari" {
					call.onAriEvent(message)
				} else {
					log.LOGGER.Error("callId[%s] taskId[%s] match call application failed", call.callId, call.taskId)
				}
			}
		case <-call.closeCall:
			{
				goto CLOSE
			}
		}
	}
CLOSE:
	log.LOGGER.Info("callId[%s] close", call.callId)
	return
}

// 处理cti消息
func (call *Call) onCtiEvent(message *simplejson.Json) {
	msgId, _ := message.Get("msgId").String()
	// 发起挂断、强复位、拦截、强插、强拆、监听操作的坐席
	agentId, _ := message.Get("msgFrom").String()
	compId, _ := message.Get("compId").String()
	ctiServerid, _ := message.Get("ctiServerid").String()

	if msgId == CALL_HANGUP_MSGID {
		hangType, _ := message.Get("msgInfo").Get("hangType").String()
		if hangType == "1" {
			if call.firstMem != nil && call.firstMem.DeviceId == agentId {
				call.firstMem.Hangup()
			} else if call.secondMem != nil && call.secondMem.DeviceId == agentId {
				call.secondMem.Hangup()
			} else if call.thirdMem != nil && call.thirdMem.DeviceId == agentId {
				call.thirdMem.Hangup()
			}
		} else {
			for _, mem := range call.memIdToMem {
				if mem != call.firstMem && mem != call.secondMem && agentId == mem.DeviceId {
					mem.Hangup()
				}
			}
		}
	} else if msgId == CALL_RESET_MSGID {
		deviceId, _ := message.Get("msgFrom").String()
		deviceType, _ := message.Get("msgFromType").String()
		for _, mem := range call.memIdToMem {
			if mem.DeviceId == deviceId && mem.DeviceType == deviceType {
				mem.Hangup()
			}
		}
	} else if msgId == CALL_INTERCEPE_MSGID {
		deviceType, _ := message.Get("msgFromType").String()
		// 被拦截
		peerId, _ := message.Get("msgInfo").Get("deviceid").String()
		ret := call.intercept(compId, agentId, deviceType, peerId)
		if ret == true {
			amqp.CtiResponse(ctiServerid, "404", agentId, deviceType, compId, "", call.callId, "yes", "ok")
		} else {
			amqp.CtiResponse(ctiServerid, "404", agentId, deviceType, compId, "", call.callId, "no", "unknown")
		}
	} else if msgId == CALL_INSERT_MSGID {
		deviceType, _ := message.Get("msgFromType").String()
		// 被强插
		peerId, _ := message.Get("msgInfo").Get("deviceid").String()
		ret := call.insert(compId, agentId, deviceType, peerId)
		if ret == true {
			amqp.CtiResponse(ctiServerid, "406", agentId, deviceType, compId, "", call.callId, "yes", "ok")
		} else {
			amqp.CtiResponse(ctiServerid, "406", agentId, deviceType, compId, "", call.callId, "no", "unknown")
		}
	} else if msgId == CALL_BREAK_MSGID {
		deviceType, _ := message.Get("msgFromType").String()
		breakMem := []*Member{}
		for _, mem := range call.memIdToMem {
			// 先挂断坐席
			if mem.DeviceType == "1" {
				breakMem = append([]*Member{mem}, breakMem...)
			} else {
				breakMem = append(breakMem, mem)
			}
		}
		for index := range breakMem {
			breakMem[index].Hangup()
		}
		amqp.CtiResponse(ctiServerid, "407", agentId, deviceType, compId, "", call.callId, "yes", "ok")
	} else if msgId == CALL_LISTEN_MSGID {
		deviceType, _ := message.Get("msgFromType").String()
		// 被监听
		peerId, _ := message.Get("msgInfo").Get("deviceid").String()
		ret := call.listen(compId, agentId, deviceType, peerId)
		if ret == true {
			amqp.CtiResponse(ctiServerid, "408", agentId, deviceType, compId, "", call.callId, "yes", "ok")
		} else {
			amqp.CtiResponse(ctiServerid, "408", agentId, deviceType, compId, "", call.callId, "no", "unknown")
		}
	}
}

// 拦截
func (call *Call) intercept(compId string, deviceId string, deviceType string, peerId string) (result bool) {
	_mem := new(Member)
	if call.firstMem != nil && call.firstMem.DeviceId == peerId {
		_mem = call.firstMem
	} else if call.secondMem != nil && call.secondMem.DeviceId == peerId {
		_mem = call.secondMem
	} else {
		log.LOGGER.Error("callId[%s] firstMem[%s] secondMem[%s] not match intercept mem",
			call.callId, call.firstMem.DeviceId, call.secondMem.DeviceId)
		return false
	}
	pickMem := new(Member)
	memId, _ := tools.NewUUid()
	pickMem.MemId = memId
	pickMem.CompId = compId
	pickMem.DeviceId = deviceId
	pickMem.DeviceType = deviceType
	pickMem.PeerId = _mem.PeerId
	pickMem.PeerType = _mem.PeerType
	pickMem.CallId = call.callId
	pickMem.MemType = "pick"
	pickMem.EslServerId = call.eslServerId
	pickMem.TaskId = call.taskId

	var onCallback func(mem *Member, result int)
	onCallback = func(mem *Member, result int) {
		log.LOGGER.Info("call result[%v] callId[%s] compId[%s] deviceId[%s]", result, call.callId, call.compId, mem.DeviceId)
		if result == CHUNANSWER {
			mem.CallState = STATE_HANGUP
			mem.StateCallback = nil
			call.thirdMem = nil
			if call.firstMem == nil && call.secondMem == nil {
				call.tryDestroyCall()
			}
		} else if result == CHANSWER {
			mem.CallState = STATE_ANSWER
			_mem.StateCallback = onCallback
			call.noticeAgentState(mem, "3")
		} else if result == CHHANGUP {
			mem.CallState = STATE_HANGUP
			call.thirdMem = nil
			if call.firstMem == nil && call.secondMem == nil {
				call.tryDestroyCall()
			} else {
				if mem == call.firstMem {
					call.firstMem = pickMem
				} else if mem == call.secondMem {
					call.secondMem = pickMem
				}
				pickMem.StateCallback = call.onCallResult
			}
		}
	}

	call.memIdToMem[memId] = pickMem
	pickMem.StateCallback = onCallback
	call.thirdMem = pickMem
	pickMem.Intercept(_mem.MemId, call.bridgeId, "60")
	//call.noticeCallInfo(pickMem.DeviceId)
	return true
}

// 强插
func (call *Call) insert(compId string, deviceId string, deviceType string, peerId string) (result bool) {
	_mem := new(Member)
	if call.firstMem != nil && call.firstMem.DeviceId == peerId {
		_mem = call.firstMem
	} else if call.secondMem != nil && call.secondMem.DeviceId == peerId {
		_mem = call.secondMem
	} else {
		log.LOGGER.Error("callId[%s] firstMem[%s] secondMem[%s] not match insert mem",
			call.callId, call.firstMem.DeviceId, call.secondMem.DeviceId)
		return false
	}
	insertMem := new(Member)
	memId, _ := tools.NewUUid()
	insertMem.MemId = memId
	insertMem.CompId = compId
	insertMem.DeviceId = deviceId
	insertMem.DeviceType = deviceType
	insertMem.PeerId = _mem.PeerId
	insertMem.PeerType = _mem.PeerType
	insertMem.CallId = call.callId
	insertMem.MemType = "insert"
	insertMem.EslServerId = call.eslServerId
	insertMem.TaskId = call.taskId

	onCallback := func(mem *Member, result int) {
		log.LOGGER.Info("call result[%v] callId[%s] compId[%s] deviceId[%s]", result, call.callId, call.compId, mem.DeviceId)
		if result == CHUNANSWER {
			call.thirdMem = nil
			mem.CallState = STATE_HANGUP
			if call.firstMem == nil && call.secondMem == nil {
				call.tryDestroyCall()
			}
		} else if result == CHANSWER {
			mem.CallState = STATE_ANSWER
			call.noticeAgentState(mem, "3")
		} else if result == CHHANGUP {
			mem.CallState = STATE_HANGUP
			if mem == insertMem {
				call.thirdMem = nil
				call.noticeAgentState(mem, "4")
				delete(call.memIdToMem, memId)
				if call.firstMem == nil && call.secondMem == nil {
					call.tryDestroyCall()
				}
			} else {
				insertMem.Hangup()
				call.onCallResult(mem, 2)
			}
		}
	}

	call.memIdToMem[memId] = insertMem
	call.firstMem.StateCallback = onCallback
	call.secondMem.StateCallback = onCallback
	insertMem.StateCallback = onCallback
	call.thirdMem = insertMem
	insertMem.Insert(_mem.MemId)
	//call.noticeCallInfo(insertMem.DeviceId)
	return true
}

// 监听
func (call *Call) listen(compId string, deviceId string, deviceType string, peerId string) (result bool) {
	_mem := new(Member)
	if call.firstMem != nil && call.firstMem.DeviceId == peerId {
		_mem = call.firstMem
	} else if call.secondMem != nil && call.secondMem.DeviceId == peerId {
		_mem = call.secondMem
	} else {
		log.LOGGER.Error("callId[%s] firstMem[%s] secondMem[%s] not match listen mem",
			call.callId, call.firstMem.DeviceId, call.secondMem.DeviceId)
		return false
	}
	listenMem := new(Member)
	memId, _ := tools.NewUUid()
	listenMem.MemId = memId
	listenMem.CompId = compId
	listenMem.DeviceId = deviceId
	listenMem.DeviceType = deviceType
	listenMem.PeerId = _mem.PeerId
	listenMem.PeerType = _mem.PeerType
	listenMem.CallId = call.callId
	listenMem.MemType = "listen"
	listenMem.EslServerId = call.eslServerId
	listenMem.TaskId = call.taskId

	onCallback := func(mem *Member, result int) {
		log.LOGGER.Info("call result[%v] callId[%s] compId[%s] deviceId[%s]", result, call.callId, call.compId, mem.DeviceId)
		if result == CHUNANSWER {
			call.thirdMem = nil
			mem.CallState = STATE_HANGUP
			if call.firstMem == nil && call.secondMem == nil {
				call.tryDestroyCall()
			}
		} else if result == CHANSWER {
			mem.CallState = STATE_ANSWER
			call.noticeAgentState(mem, "3")
		} else if result == CHHANGUP {
			mem.CallState = STATE_HANGUP
			if mem == listenMem {
				call.thirdMem = nil
				call.noticeAgentState(mem, "4")
				delete(call.memIdToMem, memId)
				if mem.DestoryCall == true {
					call.tryDestroyCall()
				}
			} else {
				if mem.DeviceType == "1" {
					listenMem.DestoryCall = true
				}
				call.onCallResult(mem, 2)
				listenMem.Hangup()
			}
		}
	}

	call.memIdToMem[memId] = listenMem
	call.firstMem.StateCallback = onCallback
	call.secondMem.StateCallback = onCallback
	listenMem.StateCallback = onCallback
	call.thirdMem = listenMem
	listenMem.Listen(_mem.MemId)
	//call.noticeCallInfo(listenMem.DeviceId)
	return true
}

// 处理来自ari底层的消息命令
func (call *Call) onAriEvent(message *simplejson.Json) {
	action, _ := message.Get("state").String()
	memId, _ := message.Get("memId").String()
	date, _ := message.Get("date").String()

	mem, ok := call.memIdToMem[memId]
	if !ok {
		log.LOGGER.Error("callId[%s] memId[%s] does not have mem application", call.callId, memId)
		return
	}
	log.LOGGER.Info("callId[%s] compId[%s] memId[%s] deviceId[%s,%s] mem_type[%s] action[%s]", call.callId, call.compId, memId, mem.DeviceId,
		mem.DeviceType, mem.MemType, action)

	onCallResult := call.onCallResult
	if mem.StateCallback != nil {
		onCallResult = mem.StateCallback
	}

	if action == "CHANNEL_CREATE" {
		call.cdrMaster.DetailCount += 1
		call.statePush(action, mem, date, "")
	} else if action == "CHANNEL_ANSWER" {
		onCallResult(mem, CHANSWER)
		if mem == call.secondMem {
			if mem.DeviceType == "1" {
				call.noticeAgentState(mem, "3")
			} else {
				call.noticeAgentState(call.firstMem, "3")
			}
		}
		call.statePush(action, mem, date, "")
	} else if action == "CHANNEL_RINGING" {
		call.noticeAgentState(mem, "0")
		call.statePush(action, mem, date, "")
	} else if action == "CHANNEL_PROGRESS" {
		if mem == call.secondMem {
			call.noticeCustomerProgress(mem)
		}
		call.statePush(action, mem, date, "")
	} else if action == "CHANNEL_HANGUP" {
		causeTxt, _ := message.Get("resp").String()
		log.LOGGER.Info("[INFO] callId[%s] deviceId[%s,%s] deviceType[%s] causeTxt[%s]", call.callId, mem.DeviceId,
			mem.DeviceType, mem.DeviceType, causeTxt)
		// 坐席
		if mem.DeviceType == "1" {
			call.noticeAgentState(mem, "4")
		}
		if causeTxt == "588" {
			if mem == call.firstMem {
				call.cdrMaster.CallResult = 13
			} else if mem == call.secondMem {
				call.cdrMaster.CallResult = 23
			}
		}
		call.statePush(action, mem, date, causeTxt)

		taskdt := getTaskDt(call.taskId)
		if mem.DeviceType == "3" && taskdt != nil {

			smsConditon := func() {
				if (taskdt.dailModel == 2 && mem == call.secondMem) || (taskdt.dailModel == 3 && mem == call.firstMem) {
					call.sendHangupSms(mem)
				}
			}
			// 电话呼叫了就要发送短信（不管是否接通）
			if taskdt.smsMode == 1 {
				smsConditon()
		    // 双方建立通话后才发送短信
			} else if taskdt.smsMode == 2 {
				if call.recordId != "" && call.cdrMaster.Atime != "" {
					// 计算时间间隔
					timeLayout := "2006-01-02 15:04:05"
					ats, err := time.ParseInLocation(timeLayout, call.cdrMaster.Atime, time.Local)
					if err == nil {
						_ats := ats.Unix()
						nts := time.Now().Unix()
						log.LOGGER.Info("call callid[%s] atime[%v] ntime[%v] smsduration[%v]", call.callId, _ats,
							nts, taskdt.smsDuration)

						if nts-_ats >= int64(taskdt.smsDuration) {
							call.sendHangupSms(mem)
						}
					}
				}
			}
		}

		if mem.CallState == STATE_INIT {
			onCallResult(mem, CHUNANSWER)
		} else {
			onCallResult(mem, CHHANGUP)
		}
	} else if action == "CHANNEL_CALLOUT_ERROR" {
		log.LOGGER.Info("callId[%s] deviceId[%s,%s] deviceType[%s]", call.callId, mem.DeviceId,
			mem.DeviceType, mem.DeviceType, )
		// 坐席
		if mem.DeviceType == "1" {
			call.noticeAgentState(mem, "4")
		}
		if mem == call.firstMem {
			call.cdrMaster.CallResult = 12
		} else if mem == call.secondMem {
			call.cdrMaster.CallResult = 22
		}
		call.statePush(action, mem, date, "")
		onCallResult(mem, CHUNANSWER)
	} else if action == "CHANNEL_LINK" {
		// todo
	} else if action == "playFinished" {
		playId, _ := message.Get("playId").String()
		playCallBAck, ok := mem.PlayCallback[playId]
		if ok {
			delete(mem.PlayCallback, playId)
			playCallBAck()
		}
	} else {
		log.LOGGER.Error("action[%s] does not match", action)
	}
}

// 呼叫
func (call *Call) onRouteEvent(message *simplejson.Json) {
	var (
		dailMode int
		ret      bool
	)
	cusId, _ := message.Get("cusid").String()
	deviceId, _ := message.Get("deviceid").String()
	deviceType, _ := message.Get("device_type").String()
	transParam, _ := message.Get("transparam").String()
	showNum, _ := message.Get("shownum").String()
	memId, _ := tools.NewUUid()

	mem := new(Member)

	if call.firstMem == nil {
		mem.DeviceId = deviceId
		mem.DeviceType = deviceType
		mem.MemType = "first"
		call.firstMem = mem
		call.cdrMaster.Caller = deviceId
	} else {
		mem.DeviceId = deviceId
		mem.DeviceType = deviceType
		mem.PeerId = call.firstMem.DeviceId
		mem.PeerType = call.firstMem.DeviceType
		mem.MemType = "second"
		call.secondMem = mem
		call.firstMem.PeerId = deviceId
		call.firstMem.PeerType = deviceType
		call.cdrMaster.Callee = deviceId
	}
	mem.CallId = call.callId
	mem.CompId = call.compId
	mem.MemId = memId
	mem.CusId = cusId
	mem.TaskId = call.taskId
	mem.TransParam = transParam
	mem.EslServerId = call.eslServerId
	mem.PlayCallback = make(map[string]func())
	call.memIdToMem[memId] = mem
	call.cdrMaster.RequestId = transParam

	if deviceType == "1" {
		call.cdrMaster.AgentId = deviceId
		//call.noticeCallInfo(deviceId)
	}
	if call.bridgeId == "" {
		call.bridgeId, _ = tools.NewUUid()
	}
	taskdt := getTaskDt(call.taskId)
	if taskdt != nil {
		dailMode = taskdt.dailModel
	}

	if res := call.numberPrefixMatch(mem); res == false {
		//showNum = call.applyAccessNumber(mem, showNum)
		if showNum != "" {
			if dailMode == 2 {
				ret = mem.MakeCall(showNum, call.bridgeId)
			} else {
				if mem == call.firstMem {
					ret = mem.MakeCall(showNum, call.bridgeId)
				} else {
					ret = mem.MakeCall(showNum, "")
				}
			}
			if deviceType == "3" {
				call.cdrMaster.Userphone = deviceId
				call.cdrMaster.Dh = mem.ShowNum
			}
			if ret == false {
				if dailMode == 3 {
					if deviceType == "3" {
						call.cdrMaster.CallResult = 16
					} else {
						call.cdrMaster.CallResult = 21
					}
				} else if dailMode == 2 {
					if deviceType == "3" {
						call.cdrMaster.CallResult = 26
					} else {
						call.cdrMaster.CallResult = 11
					}
				}
				call.onCallResult(mem, CHUNANSWER)
			}
		} else {
			log.LOGGER.Error("callId[%s] compId[%s] apply access number error shownum[%s]",
				call.callId, call.compId, showNum)
			call.cdrMaster.CallResult = 53
			date := time.Now().Format("2006-01-02 15:04:05")
			call.statePush("CHANNEL_CREATE", mem, date, "")
			call.statePush("CHANNEL_HANGUP", mem, date, "601")
			call.onCallResult(mem, CHUNANSWER)
		}
	} else {
		if dailMode == 3 {
			call.cdrMaster.CallResult = 10
		} else if dailMode == 2 {
			call.cdrMaster.CallResult = 20
		}
		date := time.Now().Format("2006-01-02 15:04:05")
		call.statePush("CHANNEL_CREATE", mem, date, "")
		call.statePush("CHANNEL_HANGUP", mem, date, "562")
		call.onCallResult(mem, CHUNANSWER)
	}

}

// 过滤客户号码前缀
func (call *Call) numberPrefixMatch(mem *Member) (ret bool) {
	if mem.DeviceType == "3" {
		memPrefix := mem.DeviceId[:3]
		for _, prefix := range config.NumberPrefixConfig.NumberPrefix {
			if memPrefix == prefix {
				log.LOGGER.Error("callId[%s] compId[%s] match prefix[%s] number[%s]",
					call.callId, call.compId, prefix, mem.DeviceId)
				return true
			}
		}
		return false
	} else {
		return false
	}

}

// 排队等待音
func (call *Call) playWaitMedia(message *simplejson.Json) {
	var (
		playCallBack func()
	)
	deviceId, _ := message.Get("deviceid").String()
	media, _ := message.Get("media").String()
	playCallBack = func() {
		if call.firstMem != nil {
			log.LOGGER.Info("call[%s] prepare deviceId[%s] play waiting media[%s]", call.callId, deviceId, media)
			call.firstMem.playMedia(media, playCallBack)
		}
	}
	playCallBack()
}

// 通知cti坐席与callid的绑定
func (call *Call) noticeCallInfo(agentId string) {
	toServerid := getAgentCtiServerid(call.compId, agentId)
	if toServerid != "" {
		amqp.SendCallInfo(toServerid, call.compId, agentId, call.callId)
	}
}

// 推送坐席状态到cti
func (call *Call) noticeAgentState(mem *Member, status string) {
	var (
		caller   string
		taskType int
	)
	// status "0"响铃 "3"接通 "4"挂断
	toServerid := getAgentCtiServerid(mem.CompId, mem.DeviceId)
	if toServerid != "" {
		if status == "0" {
			if mem == call.firstMem {
				caller = mem.ShowNum
			} else {
				caller = mem.PeerId
			}
			callType := "2"
			callee := mem.DeviceId
			taskdt := getTaskDt(call.taskId)
			if taskdt != nil {
				taskType = taskdt.dailModel
			}
			amqp.AgentStateChange(toServerid, mem.DeviceId, mem.DeviceType, call.compId, status, caller, callee, callType, call.taskId, taskType, mem.TransParam)
		} else {
			amqp.AgentStateChange(toServerid, mem.DeviceId, mem.DeviceType, call.compId, status, "", "", "", "", 0, "")
		}
	}
}

// 通知客户响铃
func (call *Call) noticeCustomerProgress(mem *Member) {
	var (
		taskType int
	)
	toServerId := getAgentCtiServerid(call.compId, mem.PeerId)
	if toServerId != "" {
		callType := "2"
		taskdt := getTaskDt(call.taskId)
		if taskdt != nil {
			taskType = taskdt.dailModel
		}
		amqp.SendCustomerProgressInfo(toServerId, mem.PeerId, mem.PeerType, call.compId, mem.ShowNum, mem.DeviceId, callType, call.taskId, taskType, call.firstMem.TransParam)
	}
}

// 通知录音地址
func (call *Call) noticeRecordInfo(mem *Member, fileName string) {
	toServerId := getAgentCtiServerid(call.compId, mem.DeviceId)
	if toServerId != "" {
		callType := "2"
		caller := call.cdrMaster.Caller
		callee := call.cdrMaster.Callee
		amqp.SendRecordInfo(toServerId, mem.DeviceId, mem.DeviceType, call.compId, fileName, caller, callee, callType, call.taskId, "", call.firstMem.TransParam)
	}
}

// 发送挂机短信
func (call *Call) sendHangupSms(mem *Member) {
	agentId := mem.PeerId
	deviceId := mem.DeviceId
	param := mem.TransParam
	log.LOGGER.Info(call.smsServerId)
	if call.smsServerId != "" {
		amqp.SendHangupSms(call.smsServerId, call.compId, agentId, deviceId, param, call.taskId)
	}
}

// 推送呼叫状态
func (call *Call) statePush(action string, mem *Member, date string, causeTxt string) {
	data := make(map[string]string)
	data["action"] = "notifycall"
	data["actionId"], _ = tools.NewUUid()
	data["state"] = action
	data["compId"] = mem.CompId
	data["calltype"] = "2"
	data["seq"] = mem.MemType
	data["direction"] = "2"
	data["date"] = date
	data["caller"] = mem.ShowNum
	data["callee"] = mem.DeviceId
	data["callId"] = mem.CallId
	data["taskId"] = call.taskId
	data["cusId"] = mem.CusId
	data["param"] = mem.TransParam
	if action == "CHANNEL_HANGUP" {
		data["respcode"] = causeTxt
	}
	if mem.DeviceType == "1" {
		data["agentId"] = mem.DeviceId
	}
	taskdt := getTaskDt(call.taskId)
	if taskdt != nil {
		HttpRequestPost(taskdt.callNoticeUrl, data)
	}
}


func (call *Call) onCallResult(mem *Member, result int) {
	log.LOGGER.Info("call result[%v] callId[%s] compId[%s] deviceId[%s]", result, call.callId, call.compId, mem.DeviceId)
	task := getTask(call.taskId)
	if result == CHUNANSWER {
		mem.CallState = STATE_HANGUP
		if mem == call.firstMem {
			call.firstMem = nil
			if call.cdrMaster.CallResult == 2 {
				call.cdrMaster.CallResult = 14
			}
			if call.cdrMaster.Hanguper == 0 {
				call.cdrMaster.Hanguper = 1
			}
		} else if mem == call.secondMem {
			call.secondMem = nil
			if call.cdrMaster.CallResult == 2 {
				call.cdrMaster.CallResult = 40
			}
			if call.cdrMaster.Hanguper == 0 {
				call.cdrMaster.Hanguper = 2
			}
		}
		if task != nil {
			task.callResult("0", mem.DeviceType, mem.DeviceId, call.callId)
		}
		call.hangup()
	} else if result == CHANSWER {
		mem.CallState = STATE_ANSWER
		if mem == call.secondMem {
			call.cdrMaster.CallResult = 1
			if call.firstMem != nil && call.firstMem.PlayId != "" {
				call.firstMem.StopPlay()
				delete(mem.PlayCallback, mem.PlayId)
			}
			taskdt := getTaskDt(call.taskId)
			if taskdt != nil && taskdt.dailModel != 2 {
				//call.bridgeAdd(call.firstMem.MemId)
				call.bridgeAdd(call.secondMem.MemId)
			}
			call.cdrMaster.Atime = tools.DateTime()
			call.bridgeRecord()
		} else if mem == call.firstMem {
			call.cdrMaster.CallResult = 40
		}
		if task != nil {
			task.callResult("1", mem.DeviceType, mem.DeviceId, call.callId)
		}
	} else if result == CHHANGUP {
		mem.CallState = STATE_HANGUP
		if mem == call.firstMem {
			call.firstMem = nil
			if call.cdrMaster.Hanguper == 0 {
				call.cdrMaster.Hanguper = 1
			}
		} else if mem == call.secondMem {
			call.secondMem = nil
			if call.cdrMaster.Hanguper == 0 {
				call.cdrMaster.Hanguper = 2
			}
		}
		if task != nil {
			task.callDestroy(call.callId, mem.DeviceId)
		}
		call.hangup()
	}
}

func (call *Call) bridgeAdd(memId string) {
	if call.bridgeId == "" {
		call.bridgeId, _ = tools.NewUUid()
	}
	call.SendCommandBridgeAdd(call.bridgeId, memId)
}

// 录音
func (call *Call) bridgeRecord() {
	var (
		extension string
		format    string
	)
	if call.recordId == "" {
		call.recordId, _ = tools.NewUUid()
		t := time.Now()
		datetime := fmt.Sprintf("%4d-%02d-%02d", t.Year(), t.Month(), t.Day())
		stime := fmt.Sprintf("%4d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
		if call.cdrMaster.AgentId != "" {
			extension = call.compId + call.cdrMaster.AgentId
		} else {
			extension = call.secondMem.DeviceId
		}

		dh := call.cdrMaster.Dh
		userphone := call.cdrMaster.Userphone
		// calltask_new5/830070/2019-10-15/8300701110_201909243_13649893400_20191015103325_00163e0cd7cc8b6911e9eef42a003e14.mp3
		recordPath := fmt.Sprintf("%s/%s/%s/%s_%s_%s_%s_%s", config.RecordConfig.RecordDirName, call.compId, datetime, extension,
			dh, userphone, stime, call.callId)

		compdt := getCompDt(call.compId)
		if compdt != nil {
			format = compdt.format
		}
		recordFile := fmt.Sprintf("%s%s", recordPath, format)
		call.cdrMaster.RecPath = recordFile
		call.SendCommandBridgeRecord(call.bridgeId, call.recordId, recordFile)
		call.firstMem.SendCommandSetRecordPath(recordFile)
		call.secondMem.SendCommandSetRecordPath(recordFile)
		if call.firstMem.DeviceType == "1" {
			call.noticeRecordInfo(call.firstMem, recordFile)
		} else if call.secondMem.DeviceType == "1" {
			call.noticeRecordInfo(call.secondMem, recordFile)
		}
	}
}

func (call *Call) hangup() {
	call.tryDestroyCall()
	if call.firstMem != nil {
		call.firstMem.Hangup()
		//call.firstMem = nil
	}
	if call.secondMem != nil {
		call.secondMem.Hangup()
		//call.secondMem = nil
	}
	if call.thirdMem != nil {
		call.thirdMem.Hangup()
	}
}

func (call *Call) tryDestroyCall() {
	if call.firstMem == nil && call.secondMem == nil && call.thirdMem == nil {
		call.memIdToMem = map[string]*Member{}
		log.LOGGER.Info("call[%s] destroyed", call.callId)
		//call.releaseAccessNumber()
		call.SendCommandDestroy()
		call.createMasterCdr()
		close(call.closeCall)
		delCall(call.callId)
	}
}

func (call *Call) createMasterCdr() {
	call.cdrMaster.TotalTime = time.Now().Unix() - call.cdrMaster.StimeStamp
	call.cdrMaster.StimeStamp = 0
	call.cdrMaster.Etime = tools.DateTime()
	call.cdrMaster.Flags = 1
	call.cdrMaster.ServerId = config.AmqpConfig.ServerId
	call.cdrMaster.CallType = 2
	call.cdrMaster.Direction = 2
	call.cdrMaster.CallId = call.callId
	call.cdrMaster.CompId = call.compId
	call.cdrMaster.TaskId = call.taskId
	cdr := cdr.GetCdrInstance()
	cdr.CdrQueue <- call.cdrMaster
	call.cdrMaster = nil
}

func (call *Call) SendCommandBridgeAdd(bridgeId string, memId string) {
	messageBridgeAdd := new(amqp.MessageBridgeAdd)
	messageBridgeAdd.Action = "bridgeAdd"
	messageBridgeAdd.CallId = call.callId
	messageBridgeAdd.MemId = memId
	messageBridgeAdd.BridgeId = bridgeId
	messageBridgeAdd.ServerId = config.AmqpConfig.ServerId
	messageBridgeAdd.ActionId, _ = tools.NewUUid()
	amqp.SendMsgToAmqp(call.eslServerId, messageBridgeAdd)
}

func (call *Call) SendCommandBridgeRecord(bridgeId string, recordId string, file string) {
	messageBridgeRecord := new(amqp.MessageBridgeRecord)
	messageBridgeRecord.Action = "bridgeRecord"
	messageBridgeRecord.CallId = call.callId
	messageBridgeRecord.RecordId = recordId
	messageBridgeRecord.File = fmt.Sprintf("%s%s", config.RecordConfig.AbsRecordPath, file)
	messageBridgeRecord.BridgeId = bridgeId
	messageBridgeRecord.ServerId = config.AmqpConfig.ServerId
	messageBridgeRecord.ActionId, _ = tools.NewUUid()
	amqp.SendMsgToAmqp(call.eslServerId, messageBridgeRecord)
}

func (call *Call) SendCommandDestroy() {
	messageDestroy := new(amqp.MessageDestroy)
	messageDestroy.Action = "destroy"
	messageDestroy.CallId = call.callId
	messageDestroy.ServerId = config.AmqpConfig.ServerId
	messageDestroy.ActionId, _ = tools.NewUUid()
	amqp.SendMsgToAmqp(call.eslServerId, messageDestroy)
}

// 获取call实例
func getCall(callId string) (call *Call) {
	mutexCall.RLock()
	defer mutexCall.RUnlock()
	call, _ = callIdToCall[callId]
	return call
}

// 绑定call实例
func setCall(callId string, call *Call) {
	mutexCall.Lock()
	defer mutexCall.Unlock()
	if call != nil {
		callIdToCall[callId] = call
	}
}

// 删除call实例
func delCall(callId string) {
	mutexCall.Lock()
	defer mutexCall.Unlock()
	delete(callIdToCall, callId)
}

// 生成call实例
func newCall(callId string, compId string, taskId string) (call *Call) {
	call = new(Call)
	call.taskId = taskId
	call.compId = compId
	call.callId = callId
	call.memIdToMem = make(map[string]*Member)
	call.MsgQueue = make(chan *simplejson.Json, 100)
	call.closeCall = make(chan interface{}, 1)
	call.cdrMaster = new(cdr.CdrMaster)
	call.cdrMaster.Stime = tools.DateTime()
	call.cdrMaster.StimeStamp = time.Now().Unix()
	call.cdrMaster.CallResult = 2
	compDt := getCompDt(compId)
	if compDt != nil {
		call.eslServerId = compDt.getEslServerId()
		call.smsServerId = compDt.getSmsServerId()
	}
	return call
}

// 初始化call实例
func initCall(callId string, compId string, taskId string) (call *Call) {
	call = newCall(callId, compId, taskId)
	setCall(callId, call)
	go call.execute()
	return call
}
