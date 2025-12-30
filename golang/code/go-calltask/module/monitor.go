package module

import (
	log "go-calltask/log"
	"time"
)

type Monitor struct {
}

func (m *Monitor) Execute() {
	defer func() {
		if r := recover(); r != nil {
			log.LOGGER.Error("%v", r)
			go m.Execute()
		}
	}()
	for {
		// 呼叫未响应异常
		_taskInstance := taskInstanceDeepCopy()
		for _, task := range _taskInstance {
			invites := task.invites.deepCopy()
			for callId, taskInvite := range invites {
				if taskInvite != nil {
					inviteDuration := time.Now().Unix() - taskInvite.InviteTime
					if inviteDuration > 120 {
						log.LOGGER.Info("taskId[%s] delete invites by callId[%s]", task.taskId, callId)
						task.invites.delete(callId)
						call := getCall(callId)
						if call != nil {
							call.hangup()
						}
					}
				}
			}

			calls := task.calls.deepCopy()
			for callId, taskCall := range calls {
				if taskCall != nil {
					callDuration := time.Now().Unix() - taskCall.CallTime
					if callDuration > 60*60*8 {
						log.LOGGER.Info("taskId[%s] delete calls by callId[%s]", task.taskId, callId)
						task.calls.delete(callId)
						call := getCall(callId)
						if call != nil {
							call.firstMem = nil
							call.secondMem = nil
							call.thirdMem = nil
							call.hangup()
						}
					}
				}
			}
		}
		log.LOGGER.Info("------current calls[%v]------", len(callIdToCall))
		time.Sleep(time.Second * 10)
	}
}
