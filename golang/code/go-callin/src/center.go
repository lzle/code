package src

import (
	"go-callin/core"
)

type Center struct {
}

// 接收消息，处理
func (ct *Center) Execute() {
	defer func() {
		if r := recover(); r != nil {
			core.LOGGER.Error("error in async method: %v", r)
			ct.Execute()
		}
	}()

	queue := core.AMQP.GetRecv()

	for {
		message := <-queue
		core.LOGGER.Info("callin recv data queue %d", len(queue))
		msgId, _ := message.Get("msgId").String()
		action, _ := message.Get("action").String()
		if msgId != "" {
			// todo
		} else if action != "" {
			callId, _ := message.Get("callId").String()
			// 第一次呼入
			if action == "callin" {
				call := Call{
					compId:    "",
					callId:    callId,
					Queue:     nil,
					firstMem:  nil,
					secondMem: nil,
					close:     nil,
				}
				call.Init()
				go func() {call.Queue <- message}()
				continue
			}
			call := GetCall(callId)
			if call != nil {
				go func() {call.Queue <- message}()
			}
		} else {
			core.LOGGER.Warn("miss msgId and action in message")
		}
	}
}
