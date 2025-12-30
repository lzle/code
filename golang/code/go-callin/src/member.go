package src

import (
	"go-callin/core"
	"go-callin/utils"
)

const (
	STATE_INIT   = 0 // 初始化
	STATE_HOLD   = 1 // 保持
	STATE_ANSWER = 2 // 应答
	STATE_HANGUP = 3 // 挂断
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
	// 播放id
	PlayId string
	// first second
	MemType string
	// 随路数据
	TransParam string
	// 外显
	ShowNum string
	// 底层serverId
	EslServerId string
	// member 通话状态
	CallState int
	// 监听销毁通话标识
	destoryCall bool
	// 状态回调
	StateCallback func(mem *Member, result int)
	// 放音回调
	PlayCallback map[string]func()
}

func (mem *Member) variables() *Cdr {
	v := new(Cdr)
	v.CallType = "1"
	v.CallId = mem.CallId
	v.CompId = mem.CompId
	v.RequestId = mem.TransParam
	v.Caller = mem.ShowNum

	if mem.DeviceType == "3" {
		v.GroupType = "2"
		v.Dh = mem.ShowNum
		v.Callee = mem.DeviceId
	} else if mem.DeviceType == "1" {
		v.GroupType = "1"
		v.Callee = mem.CompId + mem.DeviceId
		v.AgentId = mem.DeviceId
	} else {
		v.GroupType = "1"
		v.Callee = mem.DeviceId
	}

	v.App = mem.MemType
	if mem.MemType == "first" {
		v.Caller = mem.DeviceId
		v.Callee = mem.PeerId
		v.Cdrtype = "1"
	} else {
		v.Cdrtype = "2"
	}

	v.ServerId = core.CONFIG.ServerId()
	core.LOGGER.Info("set variables %v", v)
	return v
}

func (mem *Member) set(variables Variables) {
	message := &Set{
		Command: mem.command(),
		Variables: variables,
	}
	message.Action = "set"
	SendMsgToAmqp(mem.EslServerId, message)
}

func (mem *Member) command() Command {
	return Command{
		CallId:   mem.CallId,
		MemId:    mem.MemId,
		ServerId: core.CONFIG.ServerId(),
		ActionId: utils.NewUUid(),
	}
}

func (mem *Member) SetCdr() {
	mem.set(mem.variables())
}

func (mem *Member) SetRecord(path string) {
	variables := &Record{
		Path: path,
	}
	mem.set(variables)
}

func (mem *Member) Answer() {
	if mem.CallState == STATE_INIT {
		message := &Answer{
			Command: mem.command(),
		}
		message.Action = "answer"
		SendMsgToAmqp(mem.EslServerId, message)
	}
}

func (mem *Member) Hangup() {
	if mem.CallState != STATE_HANGUP {
		message := &Hangup{
			Command: mem.command(),
		}
		message.Action = "hangup"
		mem.CallState = STATE_HANGUP
		SendMsgToAmqp(mem.EslServerId, message)
	}
}

func (mem *Member) PlayMedia(media string, playCallBack func()) {
	mem.StopPlay()
	mem.PlayId = utils.NewUUid()
	message := &Play{
		Command:  mem.command(),
		PlayId:   mem.PlayId,
		Data:     media,
		DataType: "2",
	}
	message.Action = "play"
	SendMsgToAmqp(mem.EslServerId, message)
	if playCallBack != nil {
		mem.PlayCallback[mem.PlayId] = playCallBack
	}
}

func (mem *Member) StopPlay() {
	if mem.PlayId != "" {
		message := &StopPlay{
			Command: mem.command(),
			PlayId:	mem.PlayId,
		}
		message.Action = "stopplay"
		SendMsgToAmqp(mem.EslServerId, message)
	}
}


func (mem *Member) Speak (media string, playCallBack func()) {
	mem.StopPlay()
	mem.PlayId = utils.NewUUid()
	message := &Speak{
		Command:  mem.command(),
		SpeakId:   mem.PlayId,
		Data:     media,
		TTS:	"unimrcpserver-mrcp2",
	}
	message.Action = "speak"
	SendMsgToAmqp(mem.EslServerId, message)
	if playCallBack != nil {
		mem.PlayCallback[mem.PlayId] = playCallBack
	}
}