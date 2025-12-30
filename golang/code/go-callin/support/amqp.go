package support

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/streadway/amqp"
	"go-callin/core"
	"time"
)

const (
	ExchangeDirect     = amqp.ExchangeDirect
	ExchangeDurable    = false
	ExchangeAutoDelete = false
	QueueDurable       = false
	QueueAutoDelete    = true
	QueueExclusive     = true
	AutoAck            = true
)

type Amqp struct {
	// rabbitmq连接
	conn *amqp.Connection

	// 声明的channel
	ch *amqp.Channel

	// 连接关闭时会收到Error消息
	notifyChan chan *amqp.Error

	// 接收Delivery消息
	messages <-chan amqp.Delivery

	// 接收recvMsgChan中消息，转换为rabbitmq消息发送出去
	recv chan *simplejson.Json

	// 接收rabbitmq消息，put到sendMsgChan中
	send chan *core.AmapMessage

	// 连接关闭时，触发线程退出
	closeChan chan byte
}

// 建立连接
func (cm *Amqp) connect() error {
	url := core.CONFIG.AmqpUrl()

	conn, err := amqp.Dial(url)
	if err != nil {
		core.LOGGER.Error("failed to connect to RabbitMQ")
		return err
	}
	cm.conn = conn

	ch, err := conn.Channel()
	if err != nil {
		core.LOGGER.Error("failed to open a channel")
		return err
	}
	cm.ch = ch

	if err = cm.exchangeDelare(); err != nil {
		return err
	}

	if err = cm.queueDelare(); err != nil {
		return err
	}

	// 注册closeChan，异常关闭接收数据
	cm.notifyChan = make(chan *amqp.Error)
	cm.ch.NotifyClose(cm.notifyChan)

	cm.closeChan = make(chan byte, 1)
	return nil
}

// 重连
func (cm *Amqp) reconnect() {
	for {
		core.LOGGER.Info("connect to rabbitmq server ...")
		err := cm.connect()
		if err == nil {
			core.LOGGER.Info("connect rabbitmq success")
			break
		} else {
			time.Sleep(time.Second * 1)
			continue
		}
	}
	go cm.consuming()
	go cm.producing()
	go cm.keepalive()
}

// 创建exchange
func (cm *Amqp) exchangeDelare() error {
	err := cm.ch.ExchangeDeclare(core.CONFIG.Exchange(), ExchangeDirect, ExchangeDurable, ExchangeAutoDelete, false, false, nil)
	if err != nil {
		core.LOGGER.Error("Failed to declare an exchange " + err.Error())
		return err
	}
	return nil
}

// 声明queue队列、绑定
func (cm *Amqp) queueDelare() error {
	var (
		queueName string
	)
	queueName = fmt.Sprintf("queue_%s", core.CONFIG.ServerId())

	q, err := cm.ch.QueueDeclare(queueName, QueueDurable, QueueAutoDelete, QueueExclusive, false, nil)
	if err != nil {
		core.LOGGER.Error("failed to declare a queue " + err.Error())
		return err
	}

	// 绑定
	err = cm.ch.QueueBind(q.Name, core.CONFIG.ServerId(), core.CONFIG.Exchange(), false, nil)
	if err != nil {
		core.LOGGER.Error("failed to bind a queue " + err.Error())
		return err
	}

	// consume
	msgs, err := cm.ch.Consume(q.Name, "", AutoAck, false, false, false, nil)
	if err != nil {
		core.LOGGER.Error("failed to register a consumer " + err.Error())
		return err
	}
	cm.messages = msgs
	return nil
}

// 关闭connect、channel
func (cm *Amqp) close() {
	close(cm.closeChan)
	_ = cm.conn.Close()
	_ = cm.ch.Close()
}

func (cm *Amqp) keepalive() {
	for {
		select {
		case err := <-cm.notifyChan:
			core.LOGGER.Error("connetion is closed consuming" + err.Error())
			cm.close()
			cm.reconnect()
			return
		}
	}
}

// 发送消息
func (cm *Amqp) producing() {
	for {
		select {
		case message := <-cm.send:
			toServerid := message.ToServerid
			data, err := json.Marshal(message.Msg)
			if err != nil {
				core.LOGGER.Error("%s: %s[%s]", "failed to JSON Marshal ", err.Error(), message.Msg)
				continue
			}
			core.LOGGER.Info("prepare to send toServerid[%s] message[%v] queues[%v]", toServerid, string(data), len(cm.send))
			err = cm.ch.Publish(
				core.CONFIG.Exchange(), // exchange
				toServerid,             // routing key
				false,                  // mandatory
				false,                  // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(data),
				})

			if err != nil {
				core.LOGGER.Error("%s: %s[%s]", "failed to send message", err.Error(), string(data))
				cm.send <- message
			} else {
				core.LOGGER.Info("send amqp message success")
			}
		case <-cm.closeChan:
			core.LOGGER.Info("producing return")
			return
		}
	}
}

// 接收消息
func (cm *Amqp) consuming() {
	for {
		select {
		case data := <-cm.messages:
			core.LOGGER.Info("recv %s", data.Body)
			result, err := simplejson.NewJson(data.Body)
			if err != nil {
				core.LOGGER.Error("%s: %s[%s]", "failed to JSON data.Body", err.Error(), data.Body)
			}
			cm.recv <- result

		case <-cm.closeChan:
			core.LOGGER.Info("producing return")
			return
		}
	}
}

func (cm *Amqp) Init() {
	cm.send = make(chan *core.AmapMessage,2000)
	cm.recv = make(chan *simplejson.Json,2000)
	cm.reconnect()
	core.AMQP = cm
}

func (cm *Amqp) GetSend() chan *core.AmapMessage {
	return cm.send
}

func (cm *Amqp) GetRecv() chan *simplejson.Json {
	return cm.recv
}

