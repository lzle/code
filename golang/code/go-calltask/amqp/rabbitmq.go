package amqp

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/streadway/amqp"
	"go-calltask/config"
	log "go-calltask/log"
	"time"
)

var (
	AmqpClientInstance *AmqpClient
	AmqpConfig = config.AmqpConfig
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

type AmqpClient struct {
	// rabbitmq连接
	conn *amqp.Connection

	// 声明的channel
	ch *amqp.Channel

	// 连接关闭时会收到Error消息
	notifyChan chan *amqp.Error

	// 接收Delivery消息
	messages <-chan amqp.Delivery

	// 接收recvMsgChan中消息，转换为rabbitmq消息发送出去
	recvMsgChan chan *simplejson.Json

	// 接收rabbitmq消息，put到sendMsgChan中
	sendMsgChan chan *SendAmapMessage

	// 连接关闭时，触发线程退出
	closeChan chan byte
}

// 建立连接
func (cm *AmqpClient) connect() error {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/", AmqpConfig.User, AmqpConfig.Passwd, AmqpConfig.Host, AmqpConfig.Port)
	log.LOGGER.Info("connect to rabbitmq server %s", url)
	conn, err := amqp.Dial(url)
	if err != nil {
		log.LOGGER.Error("failed to connect to RabbitMQ %s", err.Error())
		return err
	}
	cm.conn = conn

	ch, err := conn.Channel()
	if err != nil {
		log.LOGGER.Error("failed to open a channel %s", err.Error())
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
func (cm *AmqpClient) reconnect() {
	for {
		err := cm.connect()
		if err == nil {
			log.LOGGER.Info("connect rabbitmq success")
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
func (cm *AmqpClient) exchangeDelare() error {
	err := cm.ch.ExchangeDeclare(AmqpConfig.Exchange, ExchangeDirect, ExchangeDurable, ExchangeAutoDelete, false, false, nil)
	if err != nil {
		log.LOGGER.Error("failed to declare an exchange %s", err.Error())
		return err
	}
	return nil
}

// 声明queue队列、绑定
func (cm *AmqpClient) queueDelare() error {
	var (
		queueName string
	)
	queueName = fmt.Sprintf("queue_%s", AmqpConfig.ServerId)

	q, err := cm.ch.QueueDeclare(queueName, QueueDurable, QueueAutoDelete, QueueExclusive, false, nil)
	if err != nil {
		log.LOGGER.Error("failed to declare a queue %s", err.Error())
		return err
	}

	// 绑定
	err = cm.ch.QueueBind(q.Name, AmqpConfig.ServerId, AmqpConfig.Exchange, false, nil)
	if err != nil {
		log.LOGGER.Error("failed to bind a queue %s", err.Error())
		return err
	}

	// consume
	msgs, err := cm.ch.Consume(q.Name, "", AutoAck, false, false, false, nil)
	if err != nil {
		log.LOGGER.Error("failed to register a consumer %s", err.Error())
		return err
	}
	cm.messages = msgs
	return nil
}

// 关闭connect、channel
func (cm *AmqpClient) close() {
	close(cm.closeChan)
	_ = cm.conn.Close()
	_ = cm.ch.Close()
}

func (cm *AmqpClient) keepalive(){
	for {
		select {
		case err := <-cm.notifyChan:
			log.LOGGER.Error("connetion is closed consuming %s", err.Error())
			cm.close()
			cm.reconnect()
			return
		}
	}
}

// 发送消息
func (cm *AmqpClient) producing() {
	for {
		select {
		case sendAmapMessage := <-cm.sendMsgChan:
			toServerid := sendAmapMessage.ToServerid
			data, err := json.Marshal(sendAmapMessage.Msg)
			if err != nil {
				log.LOGGER.Error("failed to JSON Marshal %s %v", err.Error(), sendAmapMessage.Msg)
				continue
			}
			log.LOGGER.Info("prepare to send toServerid[%s] message[%v] queues[%v]", toServerid, string(data), len(cm.sendMsgChan))
			err = cm.ch.Publish(
				AmqpConfig.Exchange, // exchange
				toServerid,       // routing key
				false,               // mandatory
				false,               // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(data),
				})

			if err != nil {
				log.LOGGER.Error("%s: %s[%s]", "failed to send message", err.Error(), string(data))
				cm.sendMsgChan <- sendAmapMessage
			} else {
				log.LOGGER.Info("send amqp message success")
			}
		//case confirm := <-cm.confirmChan:
		//	if confirm.Ack {
		//		log.Logger.Printf(" [INFO] %s", "Push confirmed!", )
		//	}
		case <-cm.closeChan:
			log.LOGGER.Info("producing return")
			return
		}
	}
}

// 接收消息
func (cm *AmqpClient) consuming() {
	for {
		select {
		case data := <-cm.messages:
			log.LOGGER.Info("recv %s", data.Body)
			result,err := simplejson.NewJson(data.Body)
			if err != nil{
				log.LOGGER.Error("%s: %s[%s]", "failed to JSON data.Body", err.Error(), data.Body)
			}
			cm.recvMsgChan <- result

		case <-cm.closeChan:
			log.LOGGER.Info("producing return")
			return
		}
	}
}

func GetAmqpClient() (*AmqpClient) {
	return AmqpClientInstance
}

func Run(recvMsgChan chan *simplejson.Json, sendMsgChan chan *SendAmapMessage) {
	AmqpConfig = config.AmqpConfig
	amqpClient := new(AmqpClient)
	amqpClient.sendMsgChan = sendMsgChan
	amqpClient.recvMsgChan = recvMsgChan
	amqpClient.reconnect()
	AmqpClientInstance = amqpClient
}
