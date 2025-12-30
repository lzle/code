package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"go-commweb/amqp"
	"go-commweb/config"
	log "go-commweb/log"
	"go-commweb/mysql"
	"go-commweb/redis"
	"go-commweb/routers"
)

func main () {
    err := config.LoadConfig("./config.yaml")
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

	// 启动mysql连接
	mysql.Run()

	// 启动redis
	redis.Run()

    route := routers.InitRouter()
	if err := route.Run(config.ConfigParam.AddrConfig.String()); err != nil {
		return
	}
}



