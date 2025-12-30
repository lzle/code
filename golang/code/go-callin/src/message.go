package src

import "go-callin/core"

type Variables interface {
}

type Cdr struct {
	RequestId string `json:"requestid"`
	CallType  string `json:"calltype"`
	CallId    string `json:"callid"`
	CompId    string `json:"compid"`
	GroupType string `json:"grouptype,omitempty"`
	Callee    string `json:"callee"`
	Dh        string `json:"dh,omitempty"`
	AgentId   string `json:"agentid,omitempty"`
	App       string `json:"app"`
	Cdrtype   string `json:"cdrtype"`
	Caller    string `json:"caller"`
	ServerId  string `json:"serverid"`
}

type Record struct {
	Path string `json:"rec_path"`
}

type Command struct {
	Action   string `json:"action"`
	CallId   string `json:"callId"`
	MemId    string `json:"memId"`
	ServerId string `json:"serverId"`
	ActionId string `json:"actionId"`
}

type Set struct {
	Command
	Variables Variables `json:"variables"`
}

type Answer struct {
	Command
}

type Hangup struct {
	Command
}

type Play struct {
	Command
	PlayId   string `json:"playId"`
	Data     string `json:"data"`
	DataType string `json:"dataType"`
}

type StopPlay struct {
	Command
	PlayId string `json:"playId"`
}

type Speak struct {
	Command
	SpeakId string `json:"speakId"`
	Data    string `json:"data"`
	TTS     string `json:"tts"`
}

// 发送到消息 rabbitmq
func SendMsgToAmqp(serverId string, message interface{}) {
	queue := core.AMQP.GetSend()
	sendMessage := &core.AmapMessage{
		ToServerid: serverId,
		Msg:        message,
	}
	queue <- sendMessage
}
