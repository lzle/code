package src

import (
	"github.com/bitly/go-simplejson"
	"go-callin/core"
	"sync"
	"time"
)

var (
	compIdToCompDt = make(map[string]*CompDetail)
	mutexCompDt    sync.RWMutex
)

type CompDetail struct {
	// 企业id
	compId string
	// 企业过期  true过期  false正常
	IsCompExpriy bool
	// 企业欠费  true欠费  false正常
	IsCompForbidden bool
	// 企业启用  true启用  false禁用
	IsCompEnable bool
	// 录音格式  0 mp3  1 wav
	Format string
}

func (cd *CompDetail) getDetail() (bodyJson *simplejson.Json) {
	return
}

func (cd *CompDetail) update() {
	core.LOGGER.Info("prepare update comp detail compId[%s]", cd.compId)

}

// 循环更新
func (cd *CompDetail) runLoop() {
	defer func() {
		if r := recover(); r != nil {
			core.LOGGER.Info("%v", r)
		}
		go cd.runLoop()
	}()

	for {
		time.Sleep(time.Second * 60)
		cd.update()
	}
}

func (cd *CompDetail) Init() {
	mutexCompDt.Lock()
	compIdToCompDt[cd.compId] = cd
	mutexCompDt.Unlock()
	cd.update()
	go cd.runLoop()
}

func GetCompDt(compId string) *CompDetail {
	mutexCompDt.RLock()
	cd, ok := compIdToCompDt[compId]
	mutexCompDt.RUnlock()
	if !ok {
		cd = &CompDetail{
			compId: compId,
		}
		cd.Init()
	}
	return cd
}
