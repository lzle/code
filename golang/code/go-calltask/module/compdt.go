package module

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"go-calltask/config"
	log "go-calltask/log"
	"go-calltask/tools"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	// compid与compDt绑定
	compIdToCompDt = make(map[string]*compDt)
	mutexCompDt                sync.RWMutex
)

type compDt struct {
	// 企业id
	compId string
	// 无号码外呼时，间隔时长关闭任务
	autoTaskTimeliness int
	// 企业过期  true过期  false正常
	isCompExpriy bool
	// 企业欠费  true欠费  false正常
	isCompForbidden bool
	// 企业启用  true启用  false禁用
	isCompEnable bool
	// 任务前缀
	taskPrefix string
	// 录音格式  0 mp3  1 wav
	format string
	// 匿名  true匿名  false正常
	isAnonymous bool
	// 底层serverId
	eslServerIds []string
	// 短信serverId
	smsServerIds []string
}

func (cd *compDt) getDt() (bodyJson *simplejson.Json) {
	params := make(map[string]string)
	params["compId"] = cd.compId
	url := config.CompConfig.UpdateCompDetailUrl
	bodyJson = HttpRequestGet(url, params)
	return bodyJson
}

func (cd *compDt) update() {
	log.LOGGER.Info("prepare update company detail compId[%s] ", cd.compId)
	defer func() {
		if r := recover(); r != nil {
			log.LOGGER.Error("%v", r)
		}
	}()
	if bodyJson := cd.getDt(); bodyJson != nil {
		if code, _ := bodyJson.Get("code").Int(); code != tools.HttpRespBodyNormalCode {
			return
		}
		data, _ := bodyJson.Get("data").Map()
		if data != nil {
			if autoTaskTimeliness, err := strconv.Atoi(string(data["autotask_timeliness"].(json.Number))); err == nil {
				cd.autoTaskTimeliness = autoTaskTimeliness
			}
			if compExpriy, err := strconv.Atoi(string(data["compexpriy"].(json.Number))); err == nil {
				if compExpriy == 0 {
					cd.isCompExpriy = true
				} else {
					cd.isCompExpriy = false
				}
			}
			if compForbidden, err := strconv.Atoi(string(data["compforbidden"].(json.Number))); err == nil {
				if compForbidden == 0 {
					cd.isCompForbidden = true
				} else {
					cd.isCompForbidden = false
				}
			}
			//if compEnable, err := strconv.Atoi(string(data["compenable"].(json.Number))); err == nil {
			//	if compEnable == 0 {
			//		cd.IsCompEnable = false
			//	} else {
			//		cd.IsCompEnable = true
			//	}
			//}
			if format, err := strconv.Atoi(string(data["format"].(json.Number))); err == nil {
				if format == 0 {
					cd.format = ".mp3"
				} else {
					cd.format = ".wav"
				}
			}
			if taskPrefix, ok := data["taskprefix"].(string); ok == true {
				cd.taskPrefix = taskPrefix
			}
			if anonymous, err := strconv.Atoi(string(data["anonymous"].(json.Number))); err == nil {
				if anonymous == 1 {
					cd.isAnonymous = true
				} else {
					cd.isAnonymous = false
				}
			}
		}
	}

	eslServerIds := getEslServerIds(cd.compId)
	log.LOGGER.Info("eslServerIds %s ", eslServerIds)
	if len(eslServerIds) > 0 {
		cd.eslServerIds = eslServerIds
	}

	smsServerIds := getSmsServerIds(cd.compId)
	log.LOGGER.Info("smsServerIds %s ", smsServerIds)
	if len(smsServerIds) > 0 {
		cd.smsServerIds = smsServerIds
	}
}

func (cd *compDt) getEslServerId() (serverId string) {
	if len(cd.eslServerIds) > 0 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return cd.eslServerIds[r.Intn(len(cd.eslServerIds))]
	}
	return
}

func (cd *compDt) getSmsServerId() (serverId string) {
	if len(cd.smsServerIds) > 0 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return cd.smsServerIds[r.Intn(len(cd.smsServerIds))]
	}
	return
}

// 循环更新
func (cd *compDt) runLoop() {
	for {
		time.Sleep(time.Second * 60)
		cd.update()
	}
}

// 获取taskdetail实例
func getCompDt(compId string) (cd *compDt) {
	mutexCompDt.RLock()
	defer mutexCompDt.RUnlock()
	cd, _ = compIdToCompDt[compId]
	return cd
}

// 生成taskdetail实例
func newCompDt(compId string) (cd *compDt) {
	cd = new(compDt)
	cd.compId = compId
	cd.isCompEnable = true
	return cd
}

// 绑定taskdetail实例
func setCompDt(taskId string, cd *compDt) {
	mutexCompDt.Lock()
	defer mutexCompDt.Unlock()
	if cd != nil {
		compIdToCompDt[taskId] = cd
	}
}

// 初始化
func initCompDt(compId string) (cd *compDt) {
	cd = getCompDt(compId)
	if cd != nil {
		return cd
	}
	cd = newCompDt(compId)
	setCompDt(compId, cd)
	cd.update()
	go cd.runLoop()
	return cd
}
