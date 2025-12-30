package amqp

const (
	taskStartMsgId = "313"
	taskStopMsgId  = "314"
	tasklistenMsgId  = "308"
	taskinterceptMsgId  = "304"
	setAgentTranslate = "470"
	cancelAgentTranslate = "471"
)

type Message interface {
}

type TaskInfo struct {
	TaskId string  `json:"taskid"`
}

type MessageTask struct {
	CtiServerId string         `json:"ctiServerid"`
	MsgFrom     string         `json:"msgFrom"`
	MsgFromType string         `json:"msgFromType"`
	MsgId       string         `json:"msgId"`
	CompId      string         `json:"compId"`
	MsgInfo     *TaskInfo      `json:"msgInfo"`
}

type MessageTaskListen struct {
	CompId      string           `json:"compId"`
	CtiServerId string           `json:"ctiServerid"`
	MsgFrom     string           `json:"msgFrom"`
	MsgFromType string           `json:"msgFromType"`
	MsgInfo     *ListenTaskInfo  `json:"msgInfo"`
	MsgId       string           `json:"msgId"`
}

type ListenTaskInfo struct {
	CallId     string    `json:"callId"`
	DeviceType string    `json:"deviceType"`
	DeviceId   string    `json:"deviceId"`
}

type MessageTaskIntercept struct {
	CompId      string              `json:"compId"`
	CtiServerId string              `json:"ctiServerid"`
	MsgFrom     string              `json:"msgFrom"`
	MsgFromType string              `json:"msgFromType"`
	MsgInfo     *InterceptTaskInfo  `json:"msgInfo"`
	MsgId       string              `json:"msgId"`
}

type InterceptTaskInfo struct {
	CallId     string    `json:"callId"`
	DeviceType string    `json:"deviceType"`
	DeviceId   string    `json:"deviceId"`
}

type MessageAgentTranslate struct {
	MsgId       string              `json:"msgId"`
	MsgTo       string              `json:"msgTo"`
	CompId      string              `json:"compid"`
}

type SendAmapMessage struct {
	ToServerid string  `json:"to_serverid"`
	Msg        Message `json:"msg"`
}


func SendTaskStart (taskId, compId, serverId, msgFrom, msgFromType string) {
	startTaskMsg := new(MessageTask)
	startTaskMsgInfo := new(TaskInfo)
	startTaskMsgInfo.TaskId = taskId
	startTaskMsg.MsgId = taskStartMsgId
	startTaskMsg.CompId = compId
	startTaskMsg.MsgFrom = msgFrom
	startTaskMsg.MsgFromType = msgFromType
	startTaskMsg.CtiServerId = ""
	startTaskMsg.MsgInfo = startTaskMsgInfo
	SendMsgToAmqp(serverId, startTaskMsg)
}

func SendTaskStop (taskId, compId, serverId, msgFrom, msgFromType string)  {
	startTaskMsg := new(MessageTask)
	startTaskMsgInfo := new(TaskInfo)
	startTaskMsgInfo.TaskId = taskId
	startTaskMsg.MsgId = taskStopMsgId
	startTaskMsg.CompId = compId
	startTaskMsg.MsgFrom = msgFrom
	startTaskMsg.MsgFromType = msgFromType
	startTaskMsg.CtiServerId = ""
	startTaskMsg.MsgInfo = startTaskMsgInfo
	SendMsgToAmqp(serverId, startTaskMsg)
}

func SendTaskListen (serverId, compId, callId, taskId, from, fromType, to, toType string) {
	listenTaskMsg := new(MessageTaskListen)
	listenTaskMsgInfo := new(ListenTaskInfo)
	listenTaskMsg.MsgId = tasklistenMsgId
	listenTaskMsg.CompId = compId
	listenTaskMsg.MsgFrom = from
	listenTaskMsg.MsgFromType = fromType
	listenTaskMsgInfo.CallId = callId
	listenTaskMsgInfo.DeviceId = to
	listenTaskMsgInfo.DeviceType = toType
	listenTaskMsg.MsgInfo = listenTaskMsgInfo
	SendMsgToAmqp(serverId, listenTaskMsg)
}

func SendTaskIntercept(serverId, compId, callId, taskId, from, fromType, to, toType string)  {
	listenTaskMsg := new(MessageTaskIntercept)
	listenTaskMsgInfo := new(InterceptTaskInfo)
	listenTaskMsg.MsgId = taskinterceptMsgId
	listenTaskMsg.CompId = compId
	listenTaskMsg.MsgFrom = from
	listenTaskMsg.MsgFromType = fromType
	listenTaskMsgInfo.CallId = callId
	listenTaskMsgInfo.DeviceId = to
	listenTaskMsgInfo.DeviceType = toType
	listenTaskMsg.MsgInfo = listenTaskMsgInfo
	SendMsgToAmqp(serverId, listenTaskMsg)
}

func SendSetAgentTranslate(serverId, agentId, compId string)  {
	setAgentTranslateMessage := new(MessageAgentTranslate)
	setAgentTranslateMessage.MsgId = setAgentTranslate
	setAgentTranslateMessage.MsgTo = agentId
	setAgentTranslateMessage.CompId = compId
	SendMsgToAmqp(serverId, setAgentTranslateMessage)
}

func SendCancelAgentTranslate(serverId, agentId, compId string)  {
	setAgentTranslateMessage := new(MessageAgentTranslate)
	setAgentTranslateMessage.MsgId = cancelAgentTranslate
	setAgentTranslateMessage.MsgTo = agentId
	setAgentTranslateMessage.CompId = compId
	SendMsgToAmqp(serverId, setAgentTranslateMessage)
}

// 发送到消息 rabbitmq
func SendMsgToAmqp(toServerid string, message interface{}) {
	amqpClient := GetAmqpClient()
	sendMessage := SendAmapMessage{toServerid, message}
	amqpClient.sendMsgChan <- &sendMessage
}


