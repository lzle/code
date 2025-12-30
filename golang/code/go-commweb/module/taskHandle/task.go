package taskHandle

import (
	"github.com/bitly/go-simplejson"
	"go-commweb/amqp"
	log "go-commweb/log"
)

const (
	startTask     =  "startTask"
	stopTask      =  "stopTask"
	listenTask    =  "listenTask"
	interceptTask =  "interceptTask"
	startCallTask =  "startCallTask"
	stopCallTask  =  "stopCallTask"
)

type Task struct {
	RecvQueue chan *simplejson.Json
	TaskQueue chan *TaskParam
	CloseTask chan interface{}
}

func (task *Task) startTask(message *TaskParam){
	serverId := getServerId(message.CompId)
	if serverId == "" {
		log.LOGGER.Error("ctiserverId is empty")
		return
	}
	checkTask(message, serverId)

	amqp.SendTaskStart(message.TaskId, message.CompId, serverId, "", "")
}

func (task *Task) stopTask (message *TaskParam) {
	serverId := getServerId(message.CompId)
	if serverId == "" {
		log.LOGGER.Error("ctiserverId is empty")
		return
	}
	amqp.SendTaskStop(message.TaskId, message.CompId, serverId, "", "")
}

func (task *Task) listenTask (msg *TaskParam) {
	serverId := getServerId(msg.CompId)
	if serverId == "" {
		log.LOGGER.Error("ctiserverId is empty")
		return
	}
	amqp.SendTaskListen(serverId, msg.CompId, msg.CallId, msg.TaskId, msg.From, msg.FromType, msg.To, msg.ToType)
}

func (task *Task) interceptTask (msg *TaskParam)  {
	serverId := getServerId(msg.CompId)
	if serverId == "" {
		log.LOGGER.Error("ctiserverId is empty")
		return
	}
	amqp.SendTaskIntercept(serverId, msg.CompId, msg.CallId, msg.TaskId, msg.From, msg.FromType, msg.To, msg.ToType)
}

func (task *Task) startCallTask(message *TaskParam){
	serverId := getCallTaskServerId(message.TaskId)
	if serverId == "" {
		log.LOGGER.Error("ctiserverId is empty")
		return
	}

	amqp.SendTaskStart(message.TaskId, message.CompId, serverId, "admin", "4")
}

func (task *Task) stopCallTask(message *TaskParam){
	serverId := getCallTaskServerId(message.TaskId)
	if serverId == "" {
		log.LOGGER.Error("ctiserverId is empty")
		return
	}

	amqp.SendTaskStop(message.TaskId, message.CompId, serverId, "admin", "4")
}

func (task *Task) Execute()  {
	defer func() {
		if r := recover(); r != nil {
			log.LOGGER.Error("%v", r)
		}
	}()
	for {
		select {
		case message := <-task.TaskQueue:
			{
				if message.Mode == startTask {
					task.startTask(message)
				}else if message.Mode == stopTask{
					task.stopTask(message)
				}else if message.Mode == listenTask {
					task.listenTask(message)
				}else if message.Mode == interceptTask {
					task.interceptTask(message)
				}else if message.Mode == startCallTask {
					task.startCallTask(message)
				}else if message.Mode == stopCallTask {
					task.stopCallTask(message)
				}
			}
		case <-task.CloseTask:
			{
				goto CLOSE
			}
		}
	}
CLOSE:
	return
}