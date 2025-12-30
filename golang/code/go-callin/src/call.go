package src

import (
	"github.com/bitly/go-simplejson"
	"go-callin/core"
	"go-callin/utils"
	"strings"
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
	MOD_BLACKLIST        = "blacklist"
)

var (
	callIdToCall = make(map[string]*Call)
	mutexCall    sync.RWMutex
)

type Call struct {
	// 企业
	compId string
	// 标示
	callId string
	// 消息队列
	Queue chan *simplejson.Json
	// first
	firstMem *Member
	// second
	secondMem *Member
	// memId mem绑定
	memIdToMem map[string]*Member
	// 关闭
	close chan interface{}
	// 锁
	sync.Mutex
	// 路由
	router core.Router
	// 主话单
	cdrMaster *core.Master
}

func (c *Call) execute() {
	defer func() {
		if r := recover(); r != nil {
			core.LOGGER.Error("error in async method: %v", r)
			c.execute()
		}
	}()
	for {
		select {
		case message := <-c.Queue:
			core.LOGGER.Info("call[%s] recv %v", c.callId, message)
			msgId, _ := message.Get("msgId").String()
			action, _ := message.Get("action").String()
			if msgId != "" {
				c.onCtiEvent(msgId, message)
			} else if action != "" {
				c.onEslEvent(action, message)
			} else {
				core.LOGGER.Warn("miss msgId and action in message")
			}
		case <-c.close:
			core.LOGGER.Info("callId[%s] close", c.callId)
			return
		}
	}
}

func (c *Call) onCtiEvent(msgId string, message *simplejson.Json) {

}

func (c *Call) onEslEvent(action string, message *simplejson.Json) {
	if action == "callin" {
		utils.RunWithRecovery(func() {
			c.callIn(message)
		})
		return
	}

	state, _ := message.Get("state").String()
	memId, _ := message.Get("memId").String()

	mem,ok := c.memIdToMem[memId]
	if !ok {
		core.LOGGER.Error("callId[%s] memId[%s] does not have mem application",c.callId, memId)
		return
	}

	core.LOGGER.Info("callId[%s] memId[%s] deviceId[%s,%s] memType[%s] action[%s] state[%s]", c.callId, memId, mem.DeviceId, mem.DeviceType, mem.MemType, action, state)
	if action == "playFinished" {
		playId,_ := message.Get("playId").String()
		callback, ok := mem.PlayCallback[playId]
		if ok {
			delete(mem.PlayCallback, playId)
			callback()
		}
		return
	} else if action == "playStarted" {
		return
	}

	callResult := c.callResult
	if mem.StateCallback != nil {
		callResult = mem.StateCallback
	}

	if state == "CHANNEL_CREATE" {
		c.cdrMaster.DetailCount++
	} else if state == "CHANNEL_ANSWER" {
		callResult(mem, CHANSWER)
	} else if state == "CHANNEL_RINGING" {

	} else if state == "CHANNEL_PROGRESS" {

	} else if state == "CHANNEL_HANGUP" {
		causeTxt, _ := message.Get("resp").String()
		core.LOGGER.Info("callId[%s] deviceId[%s,%s] memType[%s] causeTxt[%s]", c.callId, mem.DeviceId, mem.DeviceType, mem.MemType, causeTxt)
		if causeTxt == "480" {
			c.cdrMaster.CallResult = 23
		}
		if mem.CallState == STATE_INIT {
			if c.cdrMaster.CallResult == 40 {
				c.cdrMaster.CallResult = 24
			}
			callResult(mem, CHUNANSWER)
		} else {
			callResult(mem, CHHANGUP)
		}
	}
}

func (c *Call) callIn(message *simplejson.Json) {
	var (
		deviceType string
	)

	caller, _ := message.Get("caller").String()
	callee, _ := message.Get("callee").String()
	memId, _ := message.Get("memId").String()
	serverId, _ := message.Get("serverId").String()
	trunkIp, _ := message.Get("trunkIp").String()
	trunkPort, _ := message.Get("trunkPort").String()

	routeType := c.router.GetRouteType(trunkIp, trunkPort)
	if routeType == 3 {
		deviceType = "3"
	} else {
		deviceType = "2"
	}
	mem := &Member{
		DeviceId:    caller,
		DeviceType:  deviceType,
		PeerId:      callee,
		PeerType:    "",
		MemId:       memId,
		CallId:      c.callId,
		CompId:      c.compId,
		MemType:     "first",
		ShowNum:     callee,
		EslServerId: serverId,
		PlayCallback: make(map[string]func()),
	}
	compId := c.router.GetCompId(callee)
	if compId == "" {
		core.LOGGER.Error("callIn caller[%s] callee[%s] compId is empty", caller, callee)
		mem.Hangup()
		close(c.close)
		c.del()
		return
	}

	c.firstMem = mem
	mem.StateCallback = c.callResult
	mem.CompId = compId
	c.memIdToMem[memId] = mem
	c.cdrMaster = new(core.Master)
	c.cdrMaster.Userphone = caller
	c.cdrMaster.CallId = c.callId
	c.cdrMaster.Dh = callee
	c.cdrMaster.Caller = caller
	c.cdrMaster.Callee = callee
	c.cdrMaster.CallResult = 40
	c.cdrMaster.CompId = compId
	c.cdrMaster.ServerId = core.CONFIG.ServerId()
	c.cdrMaster.Stime = utils.DateTime()

	//compDt := GetCompDt(compId)
	//// 过期
	//if compDt.IsExpriy {
	//	log.Logger.Printf("callIn compId[%s] callId[%s] company expired",compId, callId)
	//	mem.Hangup()
	//	c.cdrMaster.CallResult = 51
	//	return
	//}
	//// 欠费
	//if compDt.IsForbidden {
	//	log.Logger.Printf("callIn compId[%s] callId[%s] company forbidden",compId, callId)
	//	mem.Hangup()
	//	c.cdrMaster.CallResult = 52
	//	return
	//}

	// 请求并发控制
	// Todo

	mem.SetCdr()
	mem.Answer()
	mode, defined := c.router.AccessMode(compId, callee)
	c.onRoute(mem, mode, defined)
}

func (c *Call) onRoute(mem *Member, mode string, defined string) {
	if mode == MOD_BLACKLIST {
		blackId := defined
		if c.matchBlackList(blackId) {
			mode, defined = c.BlackProcess(blackId)
			c.onRoute(mem, mode, defined)
		} else {
			mode, defined := c.router.AccessModeNext(c.compId, c.firstMem.PeerId)
			c.onRoute(mem, mode, defined)
		}
	} else {
		params := make(map[string]string)
		params["callId"] = c.callId
		params["caller"] = mem.DeviceId
		params["callee"] = mem.PeerId
		url := core.CONFIG.TTSUrl()
		bodyJson := HttpRequestGet(url, params)
		if bodyJson != nil {
			data,_ := bodyJson.Get("media").String()
			dataType,_ := bodyJson.Get("type").String()
			playBack := func() {
				time.Sleep(time.Second*2)
				if dataType == "1"{
					mem.PlayMedia(data, mem.Hangup)
				} else {
					mem.Speak(data, mem.Hangup)
				}
			}
			if dataType == "1"{
				mem.PlayMedia(data, playBack)
			} else {
				mem.Speak(data, playBack)
			}
		} else {
			mem.Hangup()
		}
	}
}

func (c *Call) onEvent() {

}

func (c *Call) matchBlackList(blackId string) bool {
	caller := c.firstMem.DeviceId

	for _, item := range c.router.BlackLists(blackId) {
		bType := item[0]
		bValue := item[1]
		core.LOGGER.Info("caller[%s] match blackList type[%s] value[%s]", caller, bType, bValue)
		// 固话
		if strings.HasPrefix(caller, "0") {
			// 固话号码
			if bType == "1" && bValue == caller {
				return true
				// 区号匹配
			} else if bType == "4" && strings.HasPrefix(caller, bValue) {
				return true
			}
			// 手机
		} else {
			// 手机号码
			if bType == "2" && (bValue == "0" || bValue == caller) {
				return true
				// 手机号段
			} else if bType == "3" && strings.HasPrefix(caller, bValue) {
				return true
				// 区号匹配
			} else if bType == "4" && len(caller) > 7 {
				prefix := caller[:7]
				return c.router.MobileArea(prefix, bValue)
			}
		}
	}
	return false
}

func (c *Call) BlackProcess(blackId string) (mode string, defined string) {
	bMode, bDefined := c.router.BlackProcess(blackId, c.compId)
	if bMode == "播放语音后挂机" {
		mode = "playback"
		defined = bDefined
	} else {
		mode = "hangup"
	}
	return
}

// channel状态回调
func (c *Call) callResult(mem *Member, result int) {
	core.LOGGER.Info("call result[%v] callId[%s] compId[%s] deviceId[%s]", result, c.callId, c.compId, mem.DeviceId)
	if result == CHUNANSWER {
		mem.CallState = STATE_HANGUP
		if mem == c.firstMem {
			if c.cdrMaster.Hanguper == 0 {
				c.cdrMaster.Hanguper = 1
			}
			c.firstMem = nil
		} else if mem == c.secondMem {
			if c.cdrMaster.Hanguper == 0 {
				c.cdrMaster.Hanguper = 2
			}
			c.secondMem = nil
		}
		c.hangup()
	} else if result == CHANSWER {
		mem.CallState = STATE_ANSWER
		c.cdrMaster.Atime = utils.DateTime()
	} else if result == CHHANGUP {
		mem.CallState = STATE_HANGUP
		if mem == c.firstMem {
			if c.cdrMaster.Hanguper == 0 {
				c.cdrMaster.Hanguper = 1
			}
			c.firstMem = nil
		} else if mem == c.secondMem {
			if c.cdrMaster.Hanguper == 0 {
				c.cdrMaster.Hanguper = 2
			}
			c.secondMem = nil
		}
		c.hangup()
	}
}

// 结束通话
func (c *Call) hangup() {
	if c.firstMem != nil {
		c.firstMem.Hangup()
	}
	if c.secondMem != nil {
		c.secondMem.Hangup()
	}
	c.tryDestroyCall()
}

// 销毁通话
func (c *Call) tryDestroyCall() {
	if c.firstMem == nil && c.secondMem == nil {
		c.memIdToMem = map[string]*Member{}
		core.LOGGER.Info("call[%s] destroyed", c.callId)
		c.cdrMaster.Etime = utils.DateTime()
		queue := core.CDR.GetQueue()
		queue <- c.cdrMaster
		close(c.close)
		c.del()
	}
}

func GetCall(callId string) *Call {
	mutexCall.RLock()
	defer mutexCall.RUnlock()
	call, _ := callIdToCall[callId]
	return call
}

func (c *Call) set() {
	mutexCall.Lock()
	defer mutexCall.Unlock()
	callIdToCall[c.callId] = c
}

func (c *Call) del() {
	mutexCall.Lock()
	defer mutexCall.Unlock()
	delete(callIdToCall, c.callId)
}

func (c *Call) Init() {
	c.Queue = make(chan *simplejson.Json, 100)
	c.close = make(chan interface{}, 1)
	c.memIdToMem = make(map[string]*Member)
	c.router = core.ROUTER
	c.set()
	go func() { c.execute() }()
}
