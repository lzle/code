package module

import (
	"github.com/bitly/go-simplejson"
	"go-calltask/amqp"
	"go-calltask/config"
	log "go-calltask/log"
	"sync"
)

// 客户
type Device struct {
	CusId string
	Param string
	DeviceId string
	DeviceType string
	ShowNum string
	CallId string
}

// 坐席
type Agent struct {
	CusId string
	Param string
	DeviceId string
	DeviceType string
	ShowNum string
}


// 邀请信息
type TaskInvite struct {
	Device *Device
	Agent *Agent
	InviteTime int64
}

// 通话信息
type TaskCall struct {
	Device *Device
	Agent *Agent
	CallTime int64
}

type AvaAgent struct {
	Data map[string]int64
	Mutex sync.RWMutex
}

type AgentToCti struct {
	Data map[string]string
	Mutex sync.RWMutex
}

type Invites struct {
	Data map[string]*TaskInvite
	Mutex sync.RWMutex
}

type Calls struct {
	Data map[string]*TaskCall
	Mutex sync.RWMutex
}

type AgentMarkCallback struct {
	Data map[string]func(data *simplejson.Json)
	Mutex sync.RWMutex
}

func (ava *AvaAgent) set(key string, value int64){
	ava.Mutex.Lock()
	defer ava.Mutex.Unlock()
	ava.Data[key] = value
}

func (ava *AvaAgent) get(key string) (value int64, ok bool){
	ava.Mutex.RLock()
	defer ava.Mutex.RUnlock()
	value,ok = ava.Data[key]
	return
}

func (ava *AvaAgent) delete(key string){
	ava.Mutex.Lock()
	defer ava.Mutex.Unlock()
	delete(ava.Data,key)
}

func (ava *AvaAgent) has(key string) (ok bool){
	_, ok = ava.get(key)
	return ok
}

func (ava *AvaAgent) len() (count int){
	ava.Mutex.RLock()
	defer ava.Mutex.RUnlock()
	return len(ava.Data)
}

func (atc *AgentToCti) set(key string, value string){
	atc.Mutex.Lock()
	defer atc.Mutex.Unlock()
	atc.Data[key] = value
}

func (atc *AgentToCti) get(key string) (value string, ok bool){
	atc.Mutex.RLock()
	defer atc.Mutex.RUnlock()
	value,ok = atc.Data[key]
	return
}

func (amc *AgentMarkCallback) set(key string, value func(data *simplejson.Json)){
	amc.Mutex.Lock()
	defer amc.Mutex.Unlock()
	amc.Data[key] = value
}

func (amc *AgentMarkCallback) get(key string) (value func(data *simplejson.Json), ok bool){
	amc.Mutex.RLock()
	defer amc.Mutex.RUnlock()
	value,ok = amc.Data[key]
	return
}

func (amc *AgentMarkCallback) delete(key string){
	amc.Mutex.Lock()
	defer amc.Mutex.Unlock()
	delete(amc.Data,key)
}

func (inv *Invites) set(key string, value *TaskInvite){
	inv.Mutex.Lock()
	defer inv.Mutex.Unlock()
	inv.Data[key] = value
}

func (inv *Invites) get(key string) (value *TaskInvite, ok bool){
	inv.Mutex.RLock()
	defer inv.Mutex.RUnlock()
	value,ok = inv.Data[key]
	return
}

func (inv *Invites) delete(key string){
	inv.Mutex.Lock()
	defer inv.Mutex.Unlock()
	delete(inv.Data,key)
}

func (inv *Invites) has(key string) (ok bool){
	inv.Mutex.RLock()
	defer inv.Mutex.RUnlock()
	_, ok = inv.get(key)
	return ok
}

func (inv *Invites) len() (count int){
	inv.Mutex.RLock()
	defer inv.Mutex.RUnlock()
	return len(inv.Data)
}

func (call *Calls) set(key string, value *TaskCall){
	call.Mutex.Lock()
	defer call.Mutex.Unlock()
	call.Data[key] = value
}

func (call *Calls) get(key string) (value *TaskCall, ok bool){
	call.Mutex.RLock()
	defer call.Mutex.RUnlock()
	value,ok = call.Data[key]
	return
}

func (call *Calls) delete(key string){
	call.Mutex.Lock()
	defer call.Mutex.Unlock()
	delete(call.Data,key)
}

func (call *Calls) has(key string) (ok bool){
	call.Mutex.RLock()
	defer call.Mutex.RUnlock()
	_, ok = call.get(key)
	return ok
}

func (call *Calls) len() (count int){
	call.Mutex.RLock()
	defer call.Mutex.RUnlock()
	return len(call.Data)
}


func (inv *Invites) deepCopy() (dst map[string]*TaskInvite){
	dst = make(map[string]*TaskInvite)
	inv.Mutex.RLock()
	defer inv.Mutex.RUnlock()
	for key,value := range inv.Data{
		dst[key] =value
	}
	//inv.Mutex.RUnlock()
	return dst
}

func (call *Calls) deepCopy() (dst map[string]*TaskCall){
	dst = make(map[string]*TaskCall)
	call.Mutex.RLock()
	defer call.Mutex.RUnlock()
	for key,value := range call.Data{
		dst[key] =value
	}
	//call.Mutex.RUnlock()
	return dst
}


type CalloutParam struct {
	ActionType string `json:"action_type"`
	CusId string `json:"cusid"`
	DeviceId string `json:"deviceid"`
	DeviceType string `json:"device_type"`
	ShowNum string `json:"shownum"`
	Param string `json:"transparam"`
}

type PlayMediaParam struct {
	ActionType string `json:"action_type"`
	DeviceId string `json:"deviceid"`
	Media string `json:"media"`
}

type HangupParam struct {
	ActionType string `json:"action_type"`
}

func LoadActiveTask() {
	log.LOGGER.Info("load active task url[%s]", config.TaskConfig.ActiveTaskUrl)
	params := make(map[string]string)
	params["serverId"] = config.AmqpConfig.ServerId
	url := config.TaskConfig.ActiveTaskUrl
	bodyJson := HttpRequestGet(url, params)
	if bodyJson == nil {
		return
	}
	array, _ := bodyJson.Get("data").Array()
	for _, item := range array {
		if item_map, ok := item.(map[string]interface{}); ok{
			compId,_ := item_map["compid"].(string)
			taskId,_ := item_map["taskidentity"].(string)
			if compId != "" && taskId != "" {
				amqp.TaskStart(config.AmqpConfig.ServerId, compId, taskId)
			}
		}
	}
}