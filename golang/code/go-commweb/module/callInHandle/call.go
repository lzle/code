package callInHandle

import (
	"go-commweb/common"
	log "go-commweb/log"
	"sync"
	"time"
)

var	(
	callIdToCall = make(map[string]*call)
	callMutex  sync.RWMutex
)

type call struct {
	callId      string
	compId      string
	caller      string
	callee      string
	callInTime  string
}

func ClearCalls()  {
	for {
		cTime := common.DateTime()
		for _, call := range callIdToCall {
			durationTime := common.GetTimestap(cTime) - common.GetTimestap(call.callInTime)
			if durationTime > 2 * 60 * 60 {
				log.LOGGER.Warn("clear call[%s] current time [%s] callin time [%s] duration_time[%d]", call.callId, cTime, call.callInTime, durationTime)
				delCall(call.callId)
				company := getCompany(call.compId)
				company.callContain = common.SliceDelete(company.callContain, call.callId)
			}
		}
		time.Sleep(time.Second * 60)
	}
}

func newCall(callId, compId, caller, callee string) {
	call := new(call)
	call.callId = callId
	call.compId = compId
	call.caller = caller
	call.callee = callee
	call.callInTime = common.DateTime()
	setCall(callId, call)
}

func setCall(callId string, call *call) {
	callMutex.Lock()
	defer callMutex.Unlock()
	callIdToCall[callId] = call
}

func getCall (callId string) (call *call) {
	if call, ok := callIdToCall[callId]; ok {
		return call
	}else{
		return nil
	}
}

func delCall (callId string)  {
	callMutex.Lock()
	defer callMutex.Unlock()
	if _, ok := callIdToCall[callId]; ok {
		delete(callIdToCall, callId)
	}
}