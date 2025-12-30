package module

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"go-calltask/config"
	log "go-calltask/log"
	"go-calltask/tools"
	"strconv"
	"sync"
	"time"
)

var (
	// taskid与taskDetail绑定
	taskIdToTaskDt = make(map[string]*taskDt)
	mutextTaskDt	sync.RWMutex
)


type taskDt struct {
	// 任务id
	taskId string
	// 企业id
	compId string
	// 关闭标示
	closeTag bool
	// 任务状态 1 开启
	taskStatus int
	// 过期  0 过期  1 正常
	isExpiry bool
	// 预定  1 正常  2 预定
	isReserve bool
	// 工作时间 0非工作时间  1工作时间
	isWorkTime bool
	// 通话通知推送
	callNoticeUrl string
	// 任务通知推送
	taskNoticeUrl string
	// 坐席通知推送
	agentNoticeUrl string
	// 呼叫系数
	calloutSpeed float64
	// 外呼频率
	controlNumber int
	// 响铃时长
	ringingDuration string
	// 号码获取url
	phoneUrl string
	// 任务模式  2 预览  3 渐进
	dailModel int
	// 坐席模式  1 最长空闲时间 2 随机
	seatallocationModel int
	// 外显号码
	showNum string
	// 排队等待音乐
	waitVoice string
	// 是否是真实号码 0:虚拟 1:真实
	enTruth string
	// 是否是开启并发限制 0:非开启 1:开启
	enConCurrency string
	// 0 不发送  1 通话结束发送  2 双方建立通话发送
	smsMode int
	// smsMode = 2  建立通话超过多少秒发送  默认为0秒
	smsDuration int
}

func (td *taskDt) getDt() (bodyJson *simplejson.Json) {
	params := make(map[string]string)
	params["taskId"] = td.taskId
	params["compId"] = td.compId
	url := config.TaskConfig.UpdateTaskDetailUrl
	bodyJson = HttpRequestGet(url, params)
	return bodyJson
}

func (td *taskDt) update() {
	log.LOGGER.Info("prepare update task detail taskId[%s] compId[%s] ", td.taskId, td.compId)
	defer func() {
		if r:= recover(); r!= nil{
			log.LOGGER.Error("%v", r)
		}
	}()
	if bodyJson := td.getDt(); bodyJson != nil {
		if code, _ := bodyJson.Get("code").Int(); code != tools.HttpRespBodyNormalCode {
			return
		}
		data, _ := bodyJson.Get("data").Map()
		if data != nil {
			if taskStatus, err := strconv.Atoi(string(data["taskstatus"].(json.Number))); err == nil {
				td.taskStatus = taskStatus
			}
			if expiry, err := strconv.Atoi(string(data["expiry"].(json.Number))); err == nil {
				if expiry == 0 {
					td.isExpiry = true
				}else if expiry == 2{
					td.isReserve = true
				}else {
					td.isExpiry = false
					td.isReserve = false
				}
			}
			if worktime, err := strconv.Atoi(string(data["worktime"].(json.Number))); err == nil {
				if worktime == 0 {
					td.isWorkTime = false
				}else {
					td.isWorkTime = true
				}
			}
			if callNoticeUrl, ok := data["callnoticeurl"].(string); ok == true {
				td.callNoticeUrl = callNoticeUrl
			}
			if taskNoticeUrl, ok := data["tasknoticeurl"].(string); ok == true {
				td.taskNoticeUrl = taskNoticeUrl
			}
			if agentNoticeUrl, ok := data["agentnoticeurl"].(string); ok == true {
				td.agentNoticeUrl = agentNoticeUrl
			}
			if calloutSpeed, err := strconv.ParseFloat(string(data["calloutspeed"].(json.Number)), 64); err == nil {
				td.calloutSpeed = calloutSpeed
			}
			if controlNumber, err := strconv.Atoi(string(data["controlnumber"].(json.Number))); err == nil {
				td.controlNumber = controlNumber
			}

			if ringingDuration := string(data["ringingduration"].(json.Number)); ringingDuration != "" {
				td.ringingDuration = ringingDuration
			}
			if phoneUrl, ok := data["phoneurl"].(string); ok == true {
				td.phoneUrl = phoneUrl
			}
			if dailModel, err := strconv.Atoi(string(data["dailmodel"].(json.Number))); err == nil {
				td.dailModel = dailModel
			}
			if seatallocationModel, err := strconv.Atoi(string(data["seatallocationmodel"].(json.Number))); err == nil {
				td.seatallocationModel = seatallocationModel
			}
			if showNum, ok := data["shownum"].(string); ok == true {
				td.showNum = showNum
			}
			if waitVoice, ok := data["waitvoice"].(string); ok == true {
				td.waitVoice = waitVoice
			}
			if enTruth := string(data["entruth"].(json.Number)); enTruth != "" {
				td.enTruth = enTruth
			}
			if enConCurrency := string(data["enconcurrency"].(json.Number)); enConCurrency != "" {
				td.enConCurrency = enConCurrency
			}
			if smsMode,err := strconv.Atoi(string(data["smsmode"].(json.Number))); err == nil {
				td.smsMode = smsMode
			}
			if smsDuration,err := strconv.Atoi(string(data["smsduration"].(json.Number))); err == nil {
				td.smsDuration = smsDuration
			}
		}
	}
}

// 循环更新
func (td *taskDt) runLoop() {
	for {
		time.Sleep(time.Second * 60)
		// 关闭退出
		if td.closeTag == true {
			ins := getTaskDt(td.taskId)
			if ins != nil {
				delTaskDt(td.taskId)
			}
			return
		}
		td.update()
	}
}

// 获取taskdetail实例
func getTaskDt(taskId string) (taskDetail *taskDt) {
	mutextTaskDt.RLock()
	defer mutextTaskDt.RUnlock()
	instance, _ := taskIdToTaskDt[taskId]
	return instance
}

// 生成taskdetail实例
func newTaskDt(taskId string, compId string) (td *taskDt) {
	td = new(taskDt)
	td.taskId = taskId
	td.compId = compId
	td.closeTag = false
	return td
}

// 绑定taskdetail实例
func setTaskDt(taskId string, td *taskDt) {
	mutextTaskDt.Lock()
	defer mutextTaskDt.Unlock()
	if td != nil {
		taskIdToTaskDt[taskId] = td
	}
}

func delTaskDt(taskId string){
	mutextTaskDt.Lock()
	defer mutextTaskDt.Unlock()
	delete(taskIdToTaskDt,taskId)
}

// 初始化
func initTaskDt(taskId string, compId string) (td *taskDt) {
	td = getTaskDt(taskId)
	if td != nil {
		return td
	}
	td = newTaskDt(taskId, compId)
	setTaskDt(taskId, td)
	td.update()
	go td.runLoop()
	return td
}
