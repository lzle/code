package module

import (
	"github.com/bitly/go-simplejson"
	log "go-calltask/log"
)

type CallCenter struct {
	RecvMsgChan chan *simplejson.Json
}

func (cc *CallCenter) Execute() {
	defer func() {
		if r := recover(); r != nil{
			log.LOGGER.Error("%v", r)
			go cc.Execute()
		}
	}()

	for {
		message := <-cc.RecvMsgChan
		log.LOGGER.Info("callcenter recv data queue %v", len(cc.RecvMsgChan))
		//{'compid': '830068', 'ctiServerId': 'cti_5', 'msgId': '352', 'msgInfo': {'agentId': '1070', 'state': '1'}}
		msgId, _ := message.Get("msgId").String()
		action, _ := message.Get("action").String()
		// task
		if msgId != "" {
			// compid字段转换为compId
			if compId, _ := message.Get("compid").String(); compId != "" {
				message.Set("compId", compId)
				message.Del("compid")
			}
			compId, err := message.Get("compId").String()
			if err != nil {
				log.LOGGER.Error("message get compId error %s", msgId)
				continue
			}
			taskExecute(message, msgId, compId)
		// call
		} else if action != "" {
			if callId, _ := message.Get("callId").String(); callId != "" {
				call := getCall(callId)
				if call != nil {
					action, _ := message.Get("action").String()
					if action != "state"{
						message.Set("state", action)
					}
					message.Set("action_type", "ari")
					call.MsgQueue <- message
				}
			} else {
				log.LOGGER.Warn("match call failed %s", callId)
			}
		} else {
			log.LOGGER.Warn("miss msgId or action in message")
		}
	}
}
