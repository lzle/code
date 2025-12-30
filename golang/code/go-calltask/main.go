package main

import (
	"github.com/bitly/go-simplejson"
	"go-calltask/amqp"
	"go-calltask/cdr"
	"go-calltask/config"
	log "go-calltask/log"
	"go-calltask/module"
	"go-calltask/mysql"
	"go-calltask/redis"
)

func main() {
	err := config.InitConfig("./config.json")
	if err != nil {
		return
	}
	// 设置日志
	logger := new(log.Logger)
	logger.Init()

	// 接收amqp消息
	recvMsgChan := make(chan *simplejson.Json, 2000*1)
	// 发送amqp消息
	sendMsgChan := make(chan *amqp.SendAmapMessage, 2000*1)
	// 启动amqp连接
	amqp.Run(recvMsgChan, sendMsgChan)
	// 启动redis连接
	redis.Run()
	// 启动mysql连接
	mysql.Run()
	// 启动话单
	cdr.Run()

	callcenter := new(module.CallCenter)
	callcenter.RecvMsgChan = recvMsgChan
	go callcenter.Execute()

	monitor := new(module.Monitor)
	go monitor.Execute()

	// 加载任务
	module.LoadActiveTask()

	forever := make(chan bool)
	<-forever
}
