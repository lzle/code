package amqp

import (
	"go-calltask/config"
	"go-calltask/tools"
	"strconv"
)

type Message interface {
}

type MessageNotify struct {
	Name    string `json:"name"`
	Count   int    `json:"count,omitempty"`
	MsgInfo *SubscribeAgentMsgInfo
}

type SubscribeAgentMsgInfo struct {
	AgentIds []string `json:"agentIds"`
}

type MessageSubscribeAgent struct {
	MsgId       string                 `json:"msgId"`
	CompId      string                 `json:"compid"`
	GrpId       string                 `json:"grpId"`
	AgentIds    string                 `json:"agentIds"`
	AriFromName string                 `json:"ariFromName"`
	MsgInfo     *SubscribeAgentMsgInfo `json:"msgInfo"`
}

type UnSubscribeAgentMsgInfo struct {
	AgentIds []string `json:"agentIds"`
}

type TaskFinishToAgentMsgInfo struct {
	TaskId string `json:"taskid"`
}

type MessageUnSubscribeAgent struct {
	MsgId       string                   `json:"msgId"`
	CompId      string                   `json:"compid"`
	GrpId       string                   `json:"grpId"`
	AgentIds    string                   `json:"agentIds"`
	AriFromName string                   `json:"ariFromName"`
	MsgInfo     *UnSubscribeAgentMsgInfo `json:"msgInfo"`
}

type MessageTaskFinishToAgent struct {
	MsgId     string                   `json:"msgId"`
	CompId    string                   `json:"compid"`
	MsgTo     string                   `json:"msgTo"`
	MsgToType string                   `json:"msgToType"`
	MsgInfo   *TaskFinishToAgentMsgInfo `json:"msgInfo"`
}

type CtiResponseMsgInfo struct {
	TaskId  string `json:"taskid,omitempty"`
	CallId  string `json:"callid,omitempty"`
	Success string `json:"success"`
	Reason  string `json:"reason"`
}

type MessageCtiResponse struct {
	MsgId     string              `json:"msgId"`
	MsgTo     string              `json:"msgTo"`
	MsgToType string              `json:"msgToType"`
	CompId    string              `json:"compid"`
	MsgInfo   *CtiResponseMsgInfo `json:"msgInfo"`
}

type MessageHangupSms struct {
	CompId    string `json:"compId"`
	Bstype    int    `json:"bstype"`
	AgentId   string `json:"agentId"`
	TaskId    string `json:"taskId"`
	RevNumber string `json:"revNumber"`
	Date      string `json:"date"`
	Param     string `json:"param"`
}

type CallMsgInfo struct {
	AgentId string `json:"agentId"`
	CallId  string `json:"callId"`
}

type MessageCallInfo struct {
	MsgId       string       `json:"msgId"`
	CompId      string       `json:"compid"`
	AriFromName string       `json:"ariFromName"`
	MsgInfo     *CallMsgInfo `json:"msgInfo"`
}

type RecordMsgInfo struct {
	FileName         string `json:"fileName"`
	CallerDevice     string `json:"callerDevice"`
	CalleeDevice     string `json:"calleeDevice"`
	CallType         string `json:"callType"`
	TaskId           string `json:"taskId"`
	TaskType         string `json:"taskType"`
	TransParentParam string `json:"transParentParam"`
}

type MessageRecordInfo struct {
	MsgId       string         `json:"msgId"`
	CompId      string         `json:"compid"`
	MsgTo       string         `json:"msgTo"`
	MsgToType   string         `json:"msgToType"`
	AriFromName string         `json:"ariFromName"`
	MsgInfo     *RecordMsgInfo `json:"msgInfo"`
}

type CustomerProgressInfo struct {
	CallerDevice     string `json:"callerDevice"`
	CalleeDevice     string `json:"calleeDevice"`
	CallType         string `json:"callType"`
	TaskId           string `json:"taskId"`
	TaskType         string `json:"taskType"`
	TransParentParam string `json:"transParentParam"`
}

type MessageCustomerProgress struct {
	MsgId       string                `json:"msgId"`
	CompId      string                `json:"compid"`
	MsgTo       string                `json:"msgTo"`
	MsgToType   string                `json:"msgToType"`
	AriFromName string                `json:"ariFromName"`
	MsgInfo     *CustomerProgressInfo `json:"msgInfo"`
}

type CdrVariables struct {
	CallerId  string `json:"callerid"`
	RequestId string `json:"requestid"`
	CallType  string `json:"calltype"`
	CallId    string `json:"callid"`
	CompId    string `json:"compid"`
	TaskId    string `json:"taskid"`
	GroupType string `json:"grouptype,omitempty"`
	Callee    string `json:"callee"`
	Dh        string `json:"dh,omitempty"`
	AgentId   string `json:"agentid,omitempty"`
	App       string `json:"app"`
	Direction string `json:"direction"`
	Cdrtype   string `json:"cdrtype"`
	Caller    string `json:"caller"`
	Serverid  string `json:"serverid"`
	RingTimeout string `json:"bridge_early_media"`
}

type AgentStateChangeInfo struct {
	State        string `json:"state"`
	CallerDevice string `json:"callerDevice,omitempty"`
	CalleeDevice string `json:"calleeDevice,omitempty"`
	CallType     string `json:"callType,omitempty"`
	TaskId       string `json:"taskId,omitempty"`
	TaskType     string `json:"taskType,omitempty"`
}

type MessageAgentStateChange struct {
	MsgId       string                `json:"msgId"`
	TransParam  string                `json:"transParentParam,omitempty"`
	MsgTo       string                `json:"msgTo"`
	MsgToType   string                `json:"msgToType"`
	CompId      string                `json:"compid"`
	AriFromName string                `json:"ariFromName"`
	MsgInfo     *AgentStateChangeInfo `json:"msgInfo"`
}

type MarkAgentInfo struct {
	AgentId   string `json:"agentId"`
	RequestId string `json:"requestId"`
	CallId    string `json:"callId"`
	TaskId    string `json:"taskId"` /**/
}

type MessageMarkAgent struct {
	MsgId       string         `json:"msgId"`
	CompId      string         `json:"compid"`
	AriFromName string         `json:"ariFromName"`
	MsgInfo     *MarkAgentInfo `json:"msgInfo"`
}

type CancelMarkedAgentInfo struct {
	State string `json:"state"`
}

type MessageCancelMarkedAgent struct {
	MsgId       string                 `json:"msgId"`
	MsgTo       string                 `json:"msgTo"`
	MsgToType   string                 `json:"msgToType"`
	CompId      string                 `json:"compid"`
	AriFromName string                 `json:"ariFromName"`
	MsgInfo     *CancelMarkedAgentInfo `json:"msgInfo"`
}

type MessageCallout struct {
	Action     string        `json:"action"`
	CallId     string        `json:"callId"`
	Caller     string        `json:"caller"`
	Callee     string        `json:"callee"`
	Timeout    string        `json:"timeout"`
	MemId      string        `json:"memId"`
	BridgeId   string        `json:"bridgeId"`
	Webrtc     string        `json:"webrtc"`
	TrunkId    string        `json:"trunkId"`
	Media      string        `json:"media"`
	AutoAnswer string        `json:"autoanswer"`
	ServerId   string        `json:"serverId"`
	ActionId   string        `json:"actionId"`
	Variables  *CdrVariables `json:"variables"`
}

type MessagePlay struct {
	Action   string `json:"action"`
	CallId   string `json:"callId"`
	MemId    string `json:"memId"`
	PlayId   string `json:"playId"`
	Data     string `json:"data"`
	DataType string `json:"dataType"`
	ServerId string `json:"serverId"`
	ActionId string `json:"actionId"`
}

type MessageStopPlay struct {
	Action   string `json:"action"`
	CallId   string `json:"callId"`
	MemId    string `json:"memId"`
	PlayId   string `json:"playId"`
	ServerId string `json:"serverId"`
	ActionId string `json:"actionId"`
}

type MessageAnswer struct {
	Action   string `json:"action"`
	CallId   string `json:"callId"`
	MemId    string `json:"memId"`
	ServerId string `json:"serverId"`
	ActionId string `json:"actionId"`
}

type MessageHangup struct {
	Action   string `json:"action"`
	CallId   string `json:"callId"`
	MemId    string `json:"memId"`
	ServerId string `json:"serverId"`
	ActionId string `json:"actionId"`
}

type VariablesRecordPath struct {
	RecPath string `json:"rec_path"`
}

type MessageSetRecordPath struct {
	Action    string               `json:"action"`
	CallId    string               `json:"callId"`
	MemId     string               `json:"memId"`
	ServerId  string               `json:"serverId"`
	ActionId  string               `json:"actionId"`
	Variables *VariablesRecordPath `json:"variables"`
}

type MessageDestroy struct {
	Action   string `json:"action"`
	CallId   string `json:"callId"`
	ServerId string `json:"serverId"`
	ActionId string `json:"actionId"`
}

type MessageBridgeAdd struct {
	Action   string `json:"action"`
	CallId   string `json:"callId"`
	BridgeId string `json:"bridgeId"`
	MemId    string `json:"memId"`
	ServerId string `json:"serverId"`
	ActionId string `json:"actionId"`
}

type MessageBridgeRecord struct {
	Action   string `json:"action"`
	CallId   string `json:"callId"`
	BridgeId string `json:"bridgeId"`
	RecordId string `json:"recordId"`
	File     string `json:"file"`
	ServerId string `json:"serverId"`
	ActionId string `json:"actionId"`
}

type MessagePick struct {
	Action    string        `json:"action"`
	CallId    string        `json:"callId"`
	BridgeId  string        `json:"bridgeId"`
	MemId     string        `json:"memId"`
	PeerMemId string        `json:"peerMemId"`
	PickMemId string        `json:"pickMemId"`
	Caller    string        `json:"caller"`
	Callee    string        `json:"callee"`
	Webrtc    string        `json:"webrtc"`
	TrunkId   string        `json:"trunkId"`
	Timeout   string        `json:"timeout"`
	ServerId  string        `json:"serverId"`
	ActionId  string        `json:"actionId"`
	Variables *CdrVariables `json:"variables"`
}

type MessageInsert struct {
	Action     string        `json:"action"`
	CallId     string        `json:"callId"`
	MemId      string        `json:"memId"`
	PeerMemId  string        `json:"peerMemId"`
	Caller     string        `json:"caller"`
	Callee     string        `json:"callee"`
	Webrtc     string        `json:"webrtc"`
	TrunkId    string        `json:"trunkId"`
	Timeout    string        `json:"timeout"`
	Media      string        `json:"media"`
	AutoAnswer string        `json:"autoanswer"`
	ServerId   string        `json:"serverId"`
	ActionId   string        `json:"actionId"`
	Variables  *CdrVariables `json:"variables"`
}

type MessageListen struct {
	Action    string        `json:"action"`
	CallId    string        `json:"callId"`
	MemId     string        `json:"memId"`
	PeerMemId string        `json:"peerMemId"`
	Caller    string        `json:"caller"`
	Callee    string        `json:"callee"`
	Webrtc    string        `json:"webrtc"`
	TrunkId   string        `json:"trunkId"`
	Timeout   string        `json:"timeout"`
	Media     string        `json:"media"`
	ServerId  string        `json:"serverId"`
	ActionId  string        `json:"actionId"`
	Variables *CdrVariables `json:"variables"`
}

type TaskStartInfo struct {
	TaskId string `json:"taskid"`
}

type MessageTaskStart struct {
	CtiServerId string         `json:"ctiServerid"`
	MsgFrom     string         `json:"msgFrom"`
	MsgFromType string         `json:"msgFromType"`
	MsgId       string         `json:"msgId"`
	CompId      string         `json:"compId"`
	MsgInfo     *TaskStartInfo `json:"msgInfo"`
}

type SendAmapMessage struct {
	ToServerid string  `json:"to_serverid"`
	Msg        Message `json:"msg"`
}

func SubscribeAgentState(toServerid string, compId string, taskId string, agents []string) {
	message := new(MessageSubscribeAgent)
	msgInfo := new(SubscribeAgentMsgInfo)
	message.MsgId = "456"
	message.CompId = compId
	message.GrpId = taskId
	message.AriFromName = config.AmqpConfig.ServerId
	msgInfo.AgentIds = agents
	message.MsgInfo = msgInfo
	SendMsgToAmqp(toServerid, message)
}

func UnSubscribeAgentState(toServerid string, compId string, taskId string, agents []string) {
	message := new(MessageUnSubscribeAgent)
	msgInfo := new(UnSubscribeAgentMsgInfo)
	message.MsgId = "457"
	message.CompId = compId
	message.GrpId = taskId
	message.AriFromName = config.AmqpConfig.ServerId
	msgInfo.AgentIds = agents
	message.MsgInfo = msgInfo
	SendMsgToAmqp(toServerid, message)
}

func SendTaskFinishToAgent(toServerid string, taskId string, deviceId string, deviceType string, compId string) {
	message := new(MessageTaskFinishToAgent)
	msgInfo := new(TaskFinishToAgentMsgInfo)
	message.MsgId = "462"
	message.MsgTo = deviceId
	message.MsgToType = deviceType
	message.CompId = compId
	msgInfo.TaskId = taskId
	message.MsgInfo = msgInfo
	SendMsgToAmqp(toServerid, message)
}

func CtiResponse(toServerid string, msgId string, deviceId string, deviceType string, compId string, taskId string, callId string, success string, reason string) {
	message := new(MessageCtiResponse)
	msgInfo := new(CtiResponseMsgInfo)
	message.MsgId = msgId
	message.MsgTo = deviceId
	message.MsgToType = deviceType
	message.CompId = compId
	msgInfo.TaskId = taskId
	msgInfo.CallId = callId
	msgInfo.Success = success
	msgInfo.Reason = reason
	message.MsgInfo = msgInfo
	SendMsgToAmqp(toServerid, message)
}

// 发送挂机短信
func SendHangupSms(toServerid string, compId string, agentId string, deviceId string, param string, taskId string) {
	message := new(MessageHangupSms)
	message.CompId = compId
	message.AgentId = agentId
	message.RevNumber = deviceId
	message.Bstype = 3
	message.Param = param
	message.Date = tools.DateTime()
	message.TaskId = taskId
	SendMsgToAmqp(toServerid, message)
}

// 发送呼叫信息  305挂断时携带callid回来
func SendCallInfo(toServerid string, compId string, agentId string, CallId string) {
	message := new(MessageCallInfo)
	msgInfo := new(CallMsgInfo)
	message.MsgId = "459"
	message.CompId = compId
	message.AriFromName = config.AmqpConfig.ServerId
	msgInfo.AgentId = agentId
	msgInfo.CallId = CallId
	message.MsgInfo = msgInfo
	SendMsgToAmqp(toServerid, message)
}

// 推送录音地址
func SendRecordInfo(toServerid string, to string, toType string, compId string, fileName string, caller string, callee string, callType string, taskId string, dailModel string, transPatPam string) {
	message := new(MessageRecordInfo)
	msgInfo := new(RecordMsgInfo)
	message.MsgId = "460"
	message.CompId = compId
	message.MsgTo = to
	message.MsgToType = toType
	message.AriFromName = config.AmqpConfig.ServerId
	msgInfo.FileName = fileName
	msgInfo.CallerDevice = caller
	msgInfo.CalleeDevice = callee
	msgInfo.CallType = callType
	msgInfo.TaskId = taskId
	msgInfo.TaskType = dailModel
	msgInfo.TransParentParam = transPatPam
	message.MsgInfo = msgInfo
	SendMsgToAmqp(toServerid, message)
}

// 客户振铃推送
func SendCustomerProgressInfo(toServerid string, to string, toType string, compId string, caller string, callee string, callType string, taskId string, dailModel int, transPatPam string) {
	message := new(MessageCustomerProgress)
	msgInfo := new(CustomerProgressInfo)
	message.MsgId = "461"
	message.CompId = compId
	message.MsgTo = to
	message.MsgToType = toType
	message.AriFromName = config.AmqpConfig.ServerId
	msgInfo.CallerDevice = caller
	msgInfo.CalleeDevice = callee
	msgInfo.CallType = callType
	msgInfo.TaskId = taskId
	msgInfo.TaskType = strconv.Itoa(dailModel)
	msgInfo.TransParentParam = transPatPam
	message.MsgInfo = msgInfo
	SendMsgToAmqp(toServerid, message)
}

// 坐席通话状态改变的信息
func AgentStateChange(toServerid string, deviceId string, deviceType string, compId string, status string, caller string, callee string, callType string, taskId string, dailmodel int, transParam string) {
	// state 0//振铃  1//call up  2//call down  3//通话中  4//被挂断
	message := new(MessageAgentStateChange)
	msgInfo := new(AgentStateChangeInfo)
	msgInfo.State = status
	if status == "0" {
		msgInfo.CallerDevice = caller
		msgInfo.CalleeDevice = callee
		msgInfo.TaskId = taskId
		msgInfo.TaskType = strconv.Itoa(dailmodel)
		msgInfo.CallType = callType
		message.TransParam = transParam
	}
	message.MsgId = "453"
	message.CompId = compId
	message.MsgTo = deviceId
	message.MsgToType = deviceType
	message.AriFromName = config.AmqpConfig.ServerId
	message.MsgInfo = msgInfo
	SendMsgToAmqp(toServerid, message)
}

// 尝试标记CTI座席
func TryMarkAgent(toServerid string, compId string, agentId string, requestId string, callId string, taskId string) {
	message := new(MessageMarkAgent)
	msgInfo := new(MarkAgentInfo)
	message.MsgId = "458"
	message.CompId = compId
	msgInfo.AgentId = agentId
	msgInfo.RequestId = requestId
	msgInfo.CallId = callId
	if taskId != "" {
		msgInfo.TaskId = taskId
	}
	message.AriFromName = config.AmqpConfig.ServerId
	message.MsgInfo = msgInfo
	SendMsgToAmqp(toServerid, message)
}

// 取消标记CTI座席
func CancelMarkedAgent(toServerid string, compId string, agentId string) {
	message := new(MessageCancelMarkedAgent)
	msgInfo := new(CancelMarkedAgentInfo)
	message.MsgId = "453"
	message.MsgTo = agentId
	message.MsgToType = "1"
	message.CompId = compId
	message.AriFromName = config.AmqpConfig.ServerId
	msgInfo.State = "4"
	message.MsgInfo = msgInfo
	SendMsgToAmqp(toServerid, message)
}

// 加载启动中的任务
func TaskStart(toServerid string, compId string, taskId string) {
	message := new(MessageTaskStart)
	msgInfo := new(TaskStartInfo)
	message.MsgId = "313"
	message.MsgFrom = "admin"
	message.MsgFromType = "4"
	message.CompId = compId
	message.CtiServerId = "none"
	msgInfo.TaskId = taskId
	message.MsgInfo = msgInfo
	SendMsgToAmqp(toServerid, message)
}

// 发送到消息 rabbitmq
func SendMsgToAmqp(toServerid string, message interface{}) {
	amqpClient := GetAmqpClient()
	sendMessage := SendAmapMessage{toServerid, message}
	amqpClient.sendMsgChan <- &sendMessage
}
