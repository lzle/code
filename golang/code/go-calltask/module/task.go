package module

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"go-calltask/amqp"
	"go-calltask/config"
	log "go-calltask/log"
	"go-calltask/tools"
	"math"
	"strconv"
	"sync"
	"time"
)

const (
	// 任务开启/关闭
	TASK_START_MSGID = "313"
	TASK_STOP_MSGID  = "314"
	// 坐席加入/退出
	AGENT_JOIN_MSGID  = "320"
	AGENT_JOIN2_MSGID = "322"
	AGENT_EXIT_MSGID  = "323"
	// 标记/订阅
	MARK_AGENT_MSGID      = "358"
	AGENT_SUBSCRIBE_MSGID = "356"
	// 坐席状态改变
	AGENT_STATE_CHANGE_MSGID = "352"
	//响应状态码
	TASK_START_RESP  = "413"
	TASK_STOP_RESP   = "414"
	AGENT_JOIN_RESP  = "420"
	AGENT_JOIN2_RESP = "422"
	AGENT_EXIT_RESP  = "423"
)

var (
	// taskid与task绑定
	taskIdToTask = make(map[string]*Task)
	mutexTask            sync.RWMutex
)

type Task struct {
	// 企业id
	compId string
	// 任务id
	taskId string
	// 任务状态
	taskState bool
	// 任务详情
	taskDt *taskDt
	// 企业详情
	compDt *compDt
	// agents
	allAgents [] string
	// 空闲agents {"1001":timestamp}
	avaAgents *AvaAgent
	// agent与serverid
	agentToCti *AgentToCti
	// 坐席标记回调
	agentMarkCallback *AgentMarkCallback
	// 邀请信息
	invites *Invites
	// 通话信息
	calls *Calls
	// 客户接通等待空闲坐席
	deviceWaitAvaAgent [] *Device
	// 未呼叫客户
	deviceQueue chan *Device
	// 号码获取接收
	deviceFinish bool
	// 发起呼叫时间 timestamp
	calloutTime int64
	// 消息队列
	MsgQueue chan *simplejson.Json
	// 销毁
	closeTask chan interface{}
	// 呼叫量统计
	calloutNumberMark []int
}

// 接收cti消息
func (ts *Task) onCtiEvent() {
	defer func() {
		if r := recover(); r != nil {
			log.LOGGER.Error("%v", r)
		}
	}()
	for {
		select {

		case message := <-ts.MsgQueue:
			{
				msgId, _ := message.Get("msgId").String()
				agentId, _ := message.Get("msgFrom").String()
				fromType, _ := message.Get("msgFromType").String()
				ctiServerid, _ := message.Get("ctiServerid").String()
				if msgId == TASK_START_MSGID {
					ts.start()
					amqp.CtiResponse(ctiServerid, TASK_START_RESP, agentId, fromType, ts.compId, ts.taskId, "", "yes", "ok")
				} else if msgId == TASK_STOP_MSGID {
					ts.stop()
					amqp.CtiResponse(ctiServerid, TASK_STOP_RESP, agentId, fromType, ts.compId, ts.taskId, "", "yes", "ok")
					time.Sleep(time.Second * 3)
					// 坐席加入
				} else if tools.IsExistInArray(msgId, []string{AGENT_JOIN_MSGID, AGENT_JOIN2_MSGID,}) {
					var (
						RESP_CODE string
					)
					if msgId == AGENT_JOIN_MSGID {
						RESP_CODE = AGENT_JOIN_RESP
						if !ts.taskState {
							log.LOGGER.Error("compId[%s] taskId[%s] Task does not start ", ts.compId, ts.taskId)
							amqp.CtiResponse(ctiServerid, RESP_CODE, agentId, fromType, ts.compId, ts.taskId, "", "yes", "ok")
						} else {
							ts.agentJoin(agentId, ctiServerid)
							amqp.CtiResponse(ctiServerid, RESP_CODE, agentId, fromType, ts.compId, ts.taskId, "", "yes", "ok")
						}
					} else {
						RESP_CODE = AGENT_JOIN2_RESP
						if !ts.taskState {
							log.LOGGER.Error("compId[%s] taskId[%s] Task does not start ", ts.compId, ts.taskId)
							amqp.CtiResponse(ctiServerid, RESP_CODE, agentId, fromType, ts.compId, ts.taskId, "", "no", "task has stopped")
						} else {
							ts.agentJoin(agentId, ctiServerid)
							amqp.CtiResponse(ctiServerid, RESP_CODE, agentId, fromType, ts.compId, ts.taskId, "", "yes", "ok")
						}
					}
					// 坐席退出
				} else if msgId == AGENT_EXIT_MSGID {
					log.LOGGER.Info("compId[%s] taskId[%s] agent[%s] exit allAgents[%s]", ts.compId, ts.taskId, agentId, ts.allAgents)
					if tools.IsExistInArray(agentId, ts.allAgents) {
						ts.allAgents = tools.SliceDelete(ts.allAgents, agentId)
						ts.unsubscribe([]string{agentId})
						delAgentCti(ts.compId, agentId, ts.taskId)
					}
					if ts.avaAgents.has(agentId) {
						ts.avaAgents.delete(agentId)
					}
					ts.agentStatusPush("exittask", agentId)
					amqp.CtiResponse(ctiServerid, AGENT_EXIT_RESP, agentId, fromType, ts.compId, ts.taskId, "", "yes", "ok")
					// 标记坐席
				} else if msgId == MARK_AGENT_MSGID {
					if msgInfo, _ := message.Get("msgInfo").Map(); msgInfo != nil {
						if value, ok := msgInfo["requestId"].(string); ok {
							requestId := value
							callback, ok := ts.agentMarkCallback.get(requestId)
							if ok {
								callback(message)
							}
							ts.agentMarkCallback.delete(requestId)
						}
					}
				} else if msgId == AGENT_STATE_CHANGE_MSGID {
					ts.agentStateChange(msgId, message)
				} else if msgId == AGENT_SUBSCRIBE_MSGID {
					ts.agentStateChange(msgId, message)
				}
			}
		case <-ts.closeTask:
			{
				goto CLOSE
			}
		}
	}
CLOSE:
	log.LOGGER.Info("taskId[%s] close", ts.taskId)
	return
}

func (ts *Task) start() {
	log.LOGGER.Info("start task compId[%s] taskId[%s]", ts.compId, ts.taskId)
	if ts.taskState == true {
		log.LOGGER.Error("compId[%s] taskId[%s] task already started, can not restart", ts.compId, ts.taskId)
		return
	}
	if ts.taskDt == nil {
		log.LOGGER.Error("compId[%s] taskId[%s] task detail is not found", ts.compId, ts.taskId)
		return
	}
	if ts.compDt == nil {
		log.LOGGER.Error("compId[%s] taskId[%s] dompany detail is not found", ts.compId, ts.taskId)
		return
	}
	if ts.taskDt.dailModel == 4 {
		log.LOGGER.Error("compId[%s] taskId[%s] dailmodel[%s] unsupported", ts.compId, ts.taskId, ts.taskDt.dailModel)
		return
	}
	ts.postTaskIdBindServerId()
	ts.taskState = true
	ts.reloadAgents()
	ts.calloutTime = time.Now().Unix()
	// todo 重复开启关闭
	go ts.reservePauseTask()
	go ts.executeDetection()
}

func (ts *Task) stop() {
	ts.taskState = false
	log.LOGGER.Info("compId[%s] taskId[%s] stop ", ts.compId, ts.taskId)
}

func (ts *Task) agentJoin(agentId string, serverId string) {
	if tools.IsExistInArray(agentId, ts.allAgents) {
		log.LOGGER.Error("compId[%s] taskId[%s] agentId[%s] already in task", ts.compId, ts.taskId, agentId)
		return
	}
	ts.agentStatusPush("jointask", agentId)
	ts.subscribe([]string{agentId}, serverId)
	setAgentCti(ts.compId, agentId, ts.taskId, serverId)
	ts.allAgents = append(ts.allAgents, agentId)
}

//加载坐席，订阅
func (ts *Task) reloadAgents() {
	agents := ts.getAgentsJoinedTask()
	ts.allAgents = agents
	ts.subscribe(agents, "")
	log.LOGGER.Info("compId[%s] taskId[%s] reload agents[%s]", ts.compId, ts.taskId, agents)
}

// 订阅坐席
func (ts *Task) subscribe(subsagents []string, serverId string) () {
	for _, agentId := range subsagents {
		if serverId == "" {
			serverId = getAgentCtiServerid(ts.compId, agentId)
		}
		if serverId != "" {
			ts.agentToCti.set(agentId, serverId)
			log.LOGGER.Info("compId[%s] taskId[%s] subscribe serverId[%s] agent[%s]", ts.compId, ts.taskId, serverId, agentId)
			amqp.SubscribeAgentState(serverId, ts.compId, ts.taskId, []string{agentId})
		}
	}
}

// 取消订阅
func (ts *Task) unsubscribe(unsubsagents []string) () {
	for _, agentId := range unsubsagents {
		serverId, ok := ts.agentToCti.get(agentId)
		if !ok {
			serverId = getAgentCtiServerid(ts.compId, agentId)
		}
		if serverId != "" {
			log.LOGGER.Info("compId[%s] taskId[%s] unsubscribe serverId[%s] agent[%s]", ts.compId, ts.taskId, serverId, agentId)
			amqp.UnSubscribeAgentState(serverId, ts.compId, ts.taskId, []string{agentId})
		}
	}
}

// 通知坐席任务完成
func (ts *Task) toAgentTaskFinish(agents []string){
	for _, agentId := range agents {
		serverId, ok := ts.agentToCti.get(agentId)
		if !ok {
			serverId = getAgentCtiServerid(ts.compId, agentId)
		}
		if serverId != "" {
			log.LOGGER.Info("compId[%s] taskId[%s] notice agent task finish serverId[%s] agent[%s]", ts.compId, ts.taskId, serverId, agentId)
			amqp.SendTaskFinishToAgent(serverId,ts.taskId,agentId,"1",ts.compId)
		}
	}
}


// 推送任务绑定
func (ts *Task) postTaskIdBindServerId() {
	data := make(map[string]string)
	data["action"] = "task_bind"
	data["actionId"], _ = tools.NewUUid()
	data["compId"] = ts.compId
	data["taskId"] = ts.taskId
	data["serverId"] = config.AmqpConfig.ServerId
	HttpRequestPost(ts.taskDt.taskNoticeUrl, data)
}

// 推送任务完成
func (ts *Task) postTaskFinish() {
	data := make(map[string]string)
	data["action"] = "task_finish"
	data["actionId"], _ = tools.NewUUid()
	data["compId"] = ts.compId
	data["taskId"] = ts.taskId
	HttpRequestPost(ts.taskDt.taskNoticeUrl, data)
}

// 推送任务暂停
func (ts *Task) postTaskPause() {
	data := make(map[string]string)
	data["action"] = "task_pause"
	data["actionId"], _ = tools.NewUUid()
	data["compId"] = ts.compId
	data["taskId"] = ts.taskId
	HttpRequestPost(ts.taskDt.taskNoticeUrl, data)
}

func (ts *Task) getAgentsJoinedTask() (agents []string) {
	params := make(map[string]string)
	params["taskId"] = ts.taskId
	params["compId"] = ts.compId
	url := config.TaskConfig.GetTaskAgentUrl
	bodyJson := HttpRequestGet(url, params)
	agents = []string{}
	if bodyJson != nil {
		array, _ := bodyJson.Get("data").Array()
		for _, item := range array {
			if item_map, ok := item.(map[string]interface{}); ok {
				agents = append(agents, item_map["agentid"].(string))
			}
		}
	}
	return
}

// 执行
func (ts *Task) executeDetection() {
	defer func (){
		if r := recover(); r != nil{
			log.LOGGER.Error("%v", r)
			go ts.executeDetection()
		}
	}()

	for {
		// 关闭
		if !ts.taskState {
			log.LOGGER.Error("compId[%s] taskId[%s] state[%v] is stop ", ts.compId, ts.taskId, ts.taskState)
			return
		}
		// 预定任务
		if ts.taskDt.isReserve {
			log.LOGGER.Warn("compId[%s] taskId[%s] task reserve", ts.compId, ts.taskId)
			time.Sleep(time.Second * 1)
			continue
		}
		// 任务过期
		if ts.taskDt.isExpiry {
			log.LOGGER.Error("compId[%s] taskId[%s] task expiry", ts.compId, ts.taskId)
			ts.tryCloseTask()
			return
		}
		// 工作时间
		if !ts.taskDt.isWorkTime {
			log.LOGGER.Error("compId[%s] taskId[%s] not in work time", ts.compId, ts.taskId)
			time.Sleep(time.Second * 1)
			continue
		}
		// 企业过期
		if ts.compDt.isCompExpriy {
			log.LOGGER.Error("compId[%s] taskId[%s] company expiry", ts.compId, ts.taskId)
			time.Sleep(time.Second * 1)
			continue
		}
		// 企业欠费
		if ts.compDt.isCompForbidden {
			log.LOGGER.Error("compId[%s] taskId[%s] company forbidden", ts.compId, ts.taskId)
			time.Sleep(time.Second * 1)
			continue
		}
		// 企业禁用
		if !ts.compDt.isCompEnable {
			log.LOGGER.Error("compId[%s] taskId[%s] company it not enable", ts.compId, ts.taskId)
			time.Sleep(time.Second * 1)
			continue
		}
		if ts.taskDt.dailModel == 2 {
			ts.makePreviewCall()
		} else if ts.taskDt.dailModel == 3 {
			ts.makePredictionCall()
		} else {
			log.LOGGER.Error("compId[%s] taskId[%s] dailmodel[%s] unsupported", ts.compId, ts.taskId, ts.taskDt.dailModel)
			ret := ts.tryCloseTask()
			if ret {
				return
			}
		}
		time.Sleep(time.Second * 1)
	}
}

// 坐席状态更改
func (ts *Task) agentStateChange(msgId string, message *simplejson.Json) {
	if msgId == AGENT_STATE_CHANGE_MSGID {
		if msgInfo, _ := message.Get("msgInfo").Map(); msgInfo != nil {
			agentId := msgInfo["agentId"].(string)
			state := msgInfo["state"].(string)
			if tools.IsExistInArray(agentId, ts.allAgents) {
				if state == "0" && !ts.avaAgents.has(agentId) {
					log.LOGGER.Info("compId[%s] taskId[%s] agent[%s] state changes available, append ava_agents", ts.compId, ts.taskId, agentId)
					ts.avaAgents.set(agentId, time.Now().Unix())
					ts.agentStatusPush("available", agentId)
				} else if state == "1" && ts.avaAgents.has(agentId) {
					log.LOGGER.Info("compId[%s] taskId[%s] agent[%s] state changes unavailable, remove ava_agents", ts.compId, ts.taskId, agentId)
					ts.avaAgents.delete(agentId)
					ts.agentStatusPush("unavailable", agentId)
				}
			}
		}
	} else if msgId == AGENT_SUBSCRIBE_MSGID {
		if agentState, err := message.Get("msgInfo").Get("agentState").Array(); err == nil {
			for _, item := range agentState {
				if each_map, ok := item.(map[string]interface{}); ok {
					state, _ := each_map["state"].(string)
					if agentId, ok := each_map["agentId"].(string); ok {
						if state == "0" && !ts.avaAgents.has(agentId) {
							ts.avaAgents.set(agentId, time.Now().Unix())
							ts.agentStatusPush("available", agentId)
						} else if state == "1" && ts.avaAgents.has(agentId) {
							ts.avaAgents.delete(agentId)
							ts.agentStatusPush("unavailable", agentId)
						}
					}
				}
			}
		}
	}
}

// 精确预览 先坐席后客户
func (ts *Task) makePreviewCall() {
	log.LOGGER.Info("compId[%s] taskId[%s] make preview call dailmodel[2] avaAgentNum[%v]",
		ts.compId, ts.taskId, ts.avaAgents.len())

	defer func() {
		if r := recover(); r != nil {
			log.LOGGER.Error("%v", r)
		}
	}()

	pOutNum := ts.avaAgents.len()
	if pOutNum <= 0 {
		return
	}
	hit := 0
	for i := 0; i < pOutNum; i++ {
		if len(ts.deviceQueue) > 0 {
			device := <-ts.deviceQueue
			ts.tryMarkPreview(device)
			hit++
		} else {
			break
		}
	}
	pOutNum = pOutNum - hit
	if pOutNum <= 0 {
		return
	}
	for _, device := range ts.getCustomerGroup(pOutNum) {
		ts.tryMarkPreview(device)
		ts.calloutTime = time.Now().Unix()
	}
}

// 渐进拨号 *先客户后坐席
func (ts *Task) makePredictionCall() {
	defer func() {
		if r := recover(); r != nil {
			log.LOGGER.Error("%v", r)
		}
	}()

	invitingNum := ts.invites.len()
	avaAgentsNum := ts.avaAgents.len()
	speedNum := math.Round(float64(avaAgentsNum) * ts.taskDt.calloutSpeed)
	pOutNum := int(speedNum) - invitingNum
	pOutNum = int(math.Min(float64(pOutNum), float64(ts.taskDt.controlNumber)))
	if pOutNum < 0 {
		pOutNum = 0
	}
	if len(ts.calloutNumberMark) >= 60 {
		ts.calloutNumberMark = ts.calloutNumberMark[1:]
	}
	ts.calloutNumberMark = append(ts.calloutNumberMark, pOutNum)

	log.LOGGER.Info("compId[%s] taskId[%s] make outline dailmodel[3] invitingNum[%d] avaAgentsNum[%d] calloutSpeed[%v] controlNumber[%d] pOutNum[%d] calloutPerMinutes[%d]",
		ts.compId, ts.taskId, invitingNum, avaAgentsNum, ts.taskDt.calloutSpeed, ts.taskDt.controlNumber, pOutNum, tools.SliceSum(ts.calloutNumberMark))
	if pOutNum == 0 {
		return
	}
	hit := 0
	for i := 0; i < pOutNum; i++ {
		if len(ts.deviceWaitAvaAgent) > 0 {
			//mutexTask.Lock()
			device := ts.deviceWaitAvaAgent[0]
			ts.deviceWaitAvaAgent = ts.deviceWaitAvaAgent[1:]
			//mutexTask.Unlock()
			ts.tryMarkPrediction(device)
			hit++
		} else {
			break
		}
	}
	pOutNum = pOutNum - hit
	if pOutNum <= 0 {
		return
	}
	for _, device := range ts.getCustomerGroup(pOutNum) {
		callId := tools.NewCallId()
		inviteInfo := new(TaskInvite)
		inviteInfo.Device = device
		inviteInfo.InviteTime = time.Now().Unix()
		ts.invites.set(callId, inviteInfo)
		call := initCall(callId, ts.compId, ts.taskId)
		message := ts.prepareCalloutParam(device.CusId, device.DeviceId, device.DeviceType, device.ShowNum, device.Param)
		call.MsgQueue <- message
		ts.calloutTime = time.Now().Unix()
	}
}

func (ts *Task) getCustomerGroup(number int) (list []*Device) {
	log.LOGGER.Info("compId[%s] taskId[%s] get customer phoneurl[%s]", ts.compId, ts.taskId, ts.taskDt.phoneUrl)

	if ts.deviceFinish == true {
		ts.tryCloseTask()
		return
	}

	bodyJson := ts.getDeviceMember(number)
	if bodyJson != nil {
		array, _ := bodyJson.Get("data").Array()
		for _, item := range array {
			if item_map, ok := item.(map[string]interface{}); ok {
				_id := item_map["id"].(string)
				param := item_map["param"].(string)
				isfinish := item_map["isfinish"].(string)
				encrypt := item_map["encrypt"].(string)
				phone := item_map["phone"].(string)

				if isfinish == "true" {
					ts.deviceFinish = true
				}
				if encrypt == "true" {
					key := tools.AddToBytes(ts.compId, 16)
					ret, err := tools.AESDecrypt(key, key, phone)
					if err != nil {
						log.LOGGER.Error("compId[%s] taskId[%s] decrypt error[%s] phone[%s]",
							ts.compId, ts.taskId, err, phone)
						phone = ""
					} else {
						log.LOGGER.Info("compId[%s] taskId[%s] decrypt value[%s] phone[%s]",
							ts.compId, ts.taskId, ret, phone)
						phone = ret
					}
				}
				if phone != "" {
					device := new(Device)
					device.CusId = _id
					device.Param = param
					device.DeviceId = phone
					device.DeviceType = "3"
					device.ShowNum = ts.taskDt.showNum
					list = append(list, device)
				}
			}
		}
	}
	return list
}

func (ts *Task) getDeviceMember(count int) (bodyJson *simplejson.Json) {
	params := make(map[string]string)
	params["action"] = "number"
	params["taskId"] = ts.taskId
	params["compId"] = ts.compId
	params["count"] = strconv.Itoa(count)
	url := ts.taskDt.phoneUrl
	bodyJson = HttpRequestGet(url, params)
	return bodyJson
}

// 精准预览标记
func (ts *Task) tryMarkPreview(device *Device) {
	agentId := ts.getAvailableAgent()
	if agentId != "" {
		ret := ts.markPreviewAgent(agentId, device)
		if ret == false {
			ts.tryMarkPreview(device)
		}
	} else {
		ts.deviceQueue <- device
	}
}

// 精准预览 标记空闲坐席是否可以拨打
func (ts *Task) markPreviewAgent(agentId string, device *Device) (ret bool) {
	log.LOGGER.Info("compId[%s] taskId[%s] got available agentId[%s] mark it", ts.compId, ts.taskId, agentId)
	toServerId := getAgentCtiServerid(ts.compId, agentId)
	if toServerId == "" {
		return false
	}
	callId := tools.NewCallId()
	device.CallId = callId
	onMarkResponse := func(message *simplejson.Json) {
		ret, _ := message.Get("msgInfo").Get("result").String()
		if ret == "no" {
			ts.tryMarkPreview(device)
		} else if ret == "yes" {
			agent := new(Agent)
			agent.CusId = device.CusId
			agent.Param = device.Param
			agent.DeviceId = agentId
			agent.DeviceType = "1"
			agent.ShowNum = device.DeviceId
			//callId := tools.NewCallId()
			inviteInfo := new(TaskInvite)
			inviteInfo.Device = device
			inviteInfo.Agent = agent
			inviteInfo.InviteTime = time.Now().Unix()
			//mutexTask.Lock()
			ts.invites.set(callId, inviteInfo)
			//mutexTask.Unlock()
			call := initCall(callId, ts.compId, ts.taskId)
			message := ts.prepareCalloutParam(agent.CusId, agent.DeviceId, agent.DeviceType, agent.ShowNum, agent.Param)
			call.MsgQueue <- message
		}
	}
	noRegAgentRecheck := func(_agentId string) {
		time.Sleep(5 * time.Second)
		log.LOGGER.Info("compId[%s] taskId[%s] agent[%s] has not registered recheck, add to ava_agents", ts.compId, ts.taskId, _agentId)
		if !ts.avaAgents.has(_agentId) && tools.IsExistInArray(_agentId, ts.allAgents) {
			log.LOGGER.Info("add")
			ts.avaAgents.set(agentId, time.Now().Unix())
		}
	}

	trunkid, _, _ := getTrunkId(agentId, "1", ts.compId)
	if trunkid == "" {
		log.LOGGER.Error("compId[%s] taskId[%s] agent[%s] has not registered", ts.compId, ts.taskId, agentId)
		go noRegAgentRecheck(agentId)
		return false
	}
	requestId, _ := tools.NewUUid()
	ts.agentMarkCallback.set(requestId, onMarkResponse)
	amqp.TryMarkAgent(toServerId, ts.compId, agentId, requestId, device.CallId, ts.taskId)
	return true
}

// 渐进拨号 获取坐席去标记
func (ts *Task) tryMarkPrediction(device *Device) {
	agentId := ts.getAvailableAgent()
	if agentId != "" {
		result := ts.markPredictionAgent(agentId, device)
		if result == false {
			ts.tryMarkPrediction(device)
		}
		// 客户排队
	} else {
		call := getCall(device.CallId)
		if call != nil {
			media := ts.taskDt.waitVoice
			//mutexTask.Lock()
			ts.deviceWaitAvaAgent = append(ts.deviceWaitAvaAgent, device)
			//mutexTask.Unlock()
			if media != "" {
				message := ts.preparePlayParam(device.DeviceId, media)
				call.MsgQueue <- message
			} else {
				log.LOGGER.Error("compId[%s] taskId[%s] waiting media is None[%v] callid[%s]",
					ts.compId, ts.taskId, ts.taskDt.waitVoice, call.callId)
			}
		}
	}
}

// 渐进拨号 标记空闲坐席是否可以拨打
func (ts *Task) markPredictionAgent(agentId string, device *Device) (result bool) {
	log.LOGGER.Info("compId[%s] taskId[%s] got available agentId[%s] mark it ", ts.compId, ts.taskId, agentId)
	toServerId := getAgentCtiServerid(ts.compId, agentId)
	if toServerId == "" {
		return false
	}
	onMarkResponse := func(message *simplejson.Json) {
		ret, _ := message.Get("msgInfo").Get("result").String()
		if ret == "no" {
			ts.tryMarkPrediction(device)
		} else if ret == "yes" {
			agent := new(Agent)
			agent.CusId = device.CusId
			agent.Param = device.Param
			agent.DeviceId = agentId
			agent.DeviceType = "1"
			agent.ShowNum = device.DeviceId
			call := getCall(device.CallId)
			if call != nil {
				message := ts.prepareCalloutParam(agent.CusId, agent.DeviceId, agent.DeviceType, agent.ShowNum, agent.Param)
				call.MsgQueue <- message
				return
			} else {
				log.LOGGER.Error("callId[%s] taskId[%s] application does not exist", device.CallId, ts.taskId)
				amqp.CancelMarkedAgent(toServerId, ts.compId, agentId)
			}
		}
	}
	noRegAgentRecheck := func(_agentId string) {
		time.Sleep(5 * time.Second)
		log.LOGGER.Info("compId[%s] taskId[%s] agent[%s] has not registered recheck, add to ava_agents", ts.compId, ts.taskId, _agentId)
		if !ts.avaAgents.has(_agentId) && tools.IsExistInArray(_agentId, ts.allAgents) {
			log.LOGGER.Info("add")
			ts.avaAgents.set(agentId, time.Now().Unix())
		}
	}

	trunkid, _, _ := getTrunkId(agentId, "1", ts.compId)
	if trunkid == "" {
		log.LOGGER.Error("compId[%s] taskId[%s] agent[%s] has not registered", ts.compId, ts.taskId, agentId)
		go noRegAgentRecheck(agentId)
		return false
	}
	requestId, _ := tools.NewUUid()
	ts.agentMarkCallback.set(requestId, onMarkResponse)
	amqp.TryMarkAgent(toServerId, ts.compId, agentId, requestId, device.CallId, ts.taskId)
	return true
}

// 获取空闲坐席 最长空闲时长
func (ts *Task) getAvailableAgent() (agentId string) {
	log.LOGGER.Info("compId[%s] taskId[%s] get available agent avaAgents[%v]", ts.compId, ts.taskId, ts.avaAgents.len())
	ts.avaAgents.Mutex.RLock()
	agentId = tools.GetMinMapValue(ts.avaAgents.Data)
	ts.avaAgents.Mutex.RUnlock()
	ts.avaAgents.delete(agentId)
	return
}

// 尝试结束任务
func (ts *Task) tryCloseTask() (ret bool) {
	invitesNum := ts.invites.len()
	callsNum := ts.calls.len()
	log.LOGGER.Info("compId[%s] taskId[%s] try close task invitesNum[%d] callsNum[%d]", ts.compId, ts.taskId, invitesNum, callsNum)
	if invitesNum == 0 && callsNum == 0 {
		ts.postTaskFinish()
		ts.toAgentTaskFinish(ts.allAgents)
		ts.unsubscribe(ts.allAgents)
		ts.taskState = false
		delTask(ts.taskId)
		delTaskDt(ts.taskId)
		ts.taskDt.closeTag = true
		ts.taskDt = nil
		ts.compDt = nil
		close(ts.deviceQueue)
		close(ts.closeTask)
		log.LOGGER.Info("compId[%s] taskId[%s] close success", ts.compId, ts.taskId)
		return true
	}
	return false
}

func (ts *Task) reservePauseTask() {
	for {
		if ts.taskState == false {
			return
		}
		timeNow := time.Now().Unix()
		diffTime := timeNow - ts.calloutTime
		if diffTime > int64(ts.compDt.autoTaskTimeliness*60*60) {
			ts.taskState = false
			log.LOGGER.Info("compId[%s] taskId[%s] pause success", ts.compId, ts.taskId)
			ts.postTaskPause()
			return
		}
		time.Sleep(10 * 60 * time.Second)
	}
}

func (ts *Task) prepareCalloutParam(cusId string, deviceId string, deviceType string, showNum string, param string) (message *simplejson.Json) {
	calloutParam := new(CalloutParam)
	calloutParam.ActionType = "callout"
	calloutParam.CusId = cusId
	calloutParam.DeviceId = deviceId
	calloutParam.DeviceType = deviceType
	calloutParam.ShowNum = showNum
	calloutParam.Param = param
	jsons, err := json.Marshal(calloutParam) //转换成JSON返回的是byte[]
	if err != nil {
		log.LOGGER.Error("%v",err.Error())
	}
	res, err := simplejson.NewJson([]byte(jsons))
	//println(res)
	if err != nil {
		log.LOGGER.Error("%v",err.Error())
		return
	}
	message = res
	return
}

func (ts *Task) preparePlayParam(deviceId string, media string) (message *simplejson.Json) {
	palyParam := new(PlayMediaParam)
	palyParam.ActionType = "playmedia"
	palyParam.DeviceId = deviceId
	palyParam.Media = media
	jsons, err := json.Marshal(palyParam) //转换成JSON返回的是byte[]
	if err != nil {
		log.LOGGER.Error("%v",err.Error())
	}
	res, err := simplejson.NewJson([]byte(jsons))
	//println(res)
	if err != nil {
		log.LOGGER.Error("%v",err.Error())
		return
	}
	message = res
	return
}

func (ts *Task) prepareHangupParam() (message *simplejson.Json) {
	hangupParam := new(HangupParam)
	hangupParam.ActionType = "hangup"
	jsons, err := json.Marshal(hangupParam) //转换成JSON返回的是byte[]
	if err != nil {
		log.LOGGER.Error("%v",err.Error())
	}
	res, err := simplejson.NewJson([]byte(jsons))
	//println(res)
	if err != nil {
		log.LOGGER.Error("%v",err.Error())
		return
	}
	message = res
	return
}

func (ts *Task) callResult(status string, _type string, deviceId string, callId string) {
	log.LOGGER.Info("task call result status[%s] type[%s] deviceId[%s] callId[%s] ", status, _type, deviceId, callId)
	// 先坐席后客户
	if ts.taskDt.dailModel == 2 {
		// 坐席
		if _type == "1" {
			// 成功
			if status == "1" {
				invite, ok := ts.invites.get(callId)
				if ok {
					device := invite.Device
					ts.invites.delete(callId)
					call := getCall(callId)
					// 呼叫客户
					message := ts.prepareCalloutParam(device.CusId, device.DeviceId, device.DeviceType, device.ShowNum, device.Param)
					call.MsgQueue <- message
					ts.calls.set(callId, &TaskCall{invite.Device, invite.Agent,time.Now().Unix()})
				}
				// 失败
			} else if status == "0" {
				invite, ok := ts.invites.get(callId)
				if ok {
					device := invite.Device
					ts.invites.delete(callId)
					ts.deviceQueue <- device
				}
			}
			// 客户
		} else if _type == "3" {
			// 成功
			if status == "1" {
				// 失败
			} else if status == "0" {
				ts.calls.delete(callId)
			}
		}
		// 先客户后坐席
	} else if ts.taskDt.dailModel == 3 {
		// 客户
		if _type == "3" {
			// 成功
			if status == "1" {
				invite, ok := ts.invites.get(callId)
				if ok {
					device := invite.Device
					//mutexTask.Lock()
					ts.invites.delete(callId)
					//mutexTask.Unlock()
					device.CallId = callId
					call := new(TaskCall)
					call.Device = device
					call.CallTime = time.Now().Unix()
					ts.calls.set(callId, call)
					ts.tryMarkPrediction(device)
				}
				// 失败
			} else if status == "0" {
				ts.invites.delete(callId)
			}
			// 坐席
		} else if _type == "1" {
			// 成功
			if status == "1" {
				// do nothing
			} else if status == "0" {
				// do nothing
			}
		}
	}
}

func (ts *Task) callDestroy(callId string, deviceId string) {
	call := getCall(callId)
	if call != nil && len(ts.deviceWaitAvaAgent) > 0 {
		index := -1
		//mutexTask.Lock()
		for _index, device := range ts.deviceWaitAvaAgent {
			if deviceId == device.DeviceId {
				index = _index
				break
			}
		}
		if index != -1 {
			ts.deviceWaitAvaAgent = append(ts.deviceWaitAvaAgent[:index], ts.deviceWaitAvaAgent[index+1:]...)
		}
		//mutexTask.Unlock()
	}
	ts.calls.delete(callId)
}

func (ts *Task) agentStatusPush(status string, agentId string) {
	data := make(map[string]string)
	data["action"] = "agent_status"
	data["actionId"], _ = tools.NewUUid()
	data["compId"] = ts.compId
	data["taskId"] = ts.taskId
	data["status"] = status
	data["agentId"] = agentId
	data["date"] = string(time.Now().Unix())
	HttpRequestPost(ts.taskDt.agentNoticeUrl, data)
}

func taskExecute(message *simplejson.Json, msgId string, compId string) {
	var (
		taskId string
	)

	defer func() {
		if r := recover(); r != nil {
			log.LOGGER.Error("%v", r)
		}
	}()

	TaskMsgIdArray := []string{TASK_START_MSGID, TASK_STOP_MSGID, AGENT_JOIN_MSGID, AGENT_JOIN2_MSGID, AGENT_EXIT_MSGID, MARK_AGENT_MSGID}
	CallMsgIdArray := []string{CALL_BREAK_MSGID, CALL_HANGUP_MSGID, CALL_INSERT_MSGID, CALL_INTERCEPE_MSGID, CALL_LISTEN_MSGID, CALL_RESET_MSGID}

	// 任务相关消息
	if tools.IsExistInArray(msgId, TaskMsgIdArray) {
		if msgInfo, _ := message.Get("msgInfo").Map(); msgInfo != nil {
			// taskid字段整合
			if msgInfo["taskid"] != nil {
				if value, ok := msgInfo["taskid"].(string); ok {
					taskId = value
				}
			} else {
				if value, ok := msgInfo["taskId"].(string); ok {
					taskId = value
				}
			}
		}
		task := getTask(taskId)
		if task == nil {
			if msgId != TASK_START_MSGID {
				log.LOGGER.Error("compId[%s] taskId[%s] has not been created ", compId, taskId)
				return
			} else {
				task = newTask(taskId, compId)
				setTask(taskId, task)
			}
		}
		task.MsgQueue <- message
		// 坐席状态回复
	} else if msgId == AGENT_STATE_CHANGE_MSGID {
		taskInstance := taskInstanceDeepCopy()
		for _, ts := range taskInstance {
			if ts.compId == compId {
				ts.MsgQueue <- message
			}
		}
		// 订阅状态回复
	} else if msgId == AGENT_SUBSCRIBE_MSGID {
		if taskId, err := message.Get("grpId").String(); err == nil {
			ts := getTask(taskId)
			if ts != nil {
				ts.MsgQueue <- message
			}
		}
		// 挂断、拦截、强插、强拆、监听
	} else if tools.IsExistInArray(msgId, CallMsgIdArray) {
		callId, _ := message.Get("msgInfo").Get("callId").String()
		call := getCall(callId)
		if call != nil {
			message.Set("action_type", "cti")
			call.MsgQueue <- message
		}
	} else {
		log.LOGGER.Warn("miss command or action in message ")
	}
}

// 获取task实例
func getTask(taskId string) (task *Task) {
	mutexTask.RLock()
	defer mutexTask.RUnlock()
	instance, _ := taskIdToTask[taskId]
	return instance
}

// 生成task实例
func newTask(taskId string, compId string) (task *Task) {
	task = new(Task)
	task.taskId = taskId
	task.compId = compId
	task.taskDt = initTaskDt(taskId, compId)
	task.compDt = initCompDt(compId)
	task.agentToCti = &AgentToCti{Data: make(map[string]string)}
	task.agentMarkCallback = &AgentMarkCallback{Data: make(map[string]func(data *simplejson.Json))}
	task.avaAgents = &AvaAgent{Data: make(map[string]int64)}
	task.invites = &Invites{Data: make(map[string]*TaskInvite)}
	task.calls = &Calls{Data: make(map[string]*TaskCall)}
	task.deviceQueue = make(chan *Device, 2000)
	task.MsgQueue = make(chan *simplejson.Json, 2000)
	task.closeTask = make(chan interface{}, 1)
	go task.onCtiEvent()
	return task
}

// 绑定task实例
func setTask(taskId string, task *Task) {
	mutexTask.Lock()
	defer mutexTask.Unlock()
	if task != nil {
		taskIdToTask[taskId] = task
	}
}

// 删除call实例
func delTask(taskId string) {
	mutexTask.Lock()
	defer mutexTask.Unlock()
	delete(taskIdToTask, taskId)
}

func taskInstanceDeepCopy() (dst map[string]*Task) {
	dst = make(map[string]*Task)
	mutexTask.RLock()
	for key, value := range taskIdToTask {
		dst[key] = value
	}
	mutexTask.RUnlock()
	return dst
}
