package core

import (
	"github.com/bitly/go-simplejson"
)

type Amqp interface {
	// 发送数据
	GetSend() chan *AmapMessage
	// 接收数据
	GetRecv() chan *simplejson.Json
}

type Message interface {
}

type AmapMessage struct {
	ToServerid string  `json:"to_serverid"`
	Msg        Message `json:"msg"`
}
