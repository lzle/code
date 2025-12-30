package support

import (
	"encoding/json"
	"fmt"
	"go-callin/core"
	"io/ioutil"
)

type Config struct {
	LOG
	CDR
	TTS
	DB   DB   `json:"db"`
	AMQP AMQP `json:"amqp"`
}

type CDR struct {
	Path string `json:"cdr_path"`
}

type TTS struct {
	Url string `json:"tts_url"`
}

type LOG struct {
	Dir string `json:"log_path"`
}

type DB struct {
	Host   string `json:"host"`
	Port   int    `json:"port"`
	Db     string `json:"db"`
	User   string `json:"user"`
	Passwd string `json:"passwd"`
}

type AMQP struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Vhost    string `json:"vhost"`
	User     string `json:"user"`
	Passwd   string `json:"passwd"`
	ServerId string `json:"serverid"`
	Exchange string `json:"exchange"`
}

func (c *Config) MysqlUrl() string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?autocommit=true&timeout=3s", c.DB.User, c.DB.Passwd, "tcp", c.DB.Host, c.DB.Port, c.DB.Db)
}

func (c *Config) AmqpUrl() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", c.AMQP.User, c.AMQP.Passwd, c.AMQP.Host, c.AMQP.Port)
}

func (c *Config) Exchange() string {
	return c.AMQP.Exchange
}

func (c *Config) ServerId() string {
	return c.AMQP.ServerId
}

func (c *Config) CdrDir() string {
	return c.CDR.Path
}

func (c *Config) TTSUrl() string {
	return c.TTS.Url
}

func (c *Config) LogDir() string {
	return c.LOG.Dir
}


func (c *Config) Init() {
	fileName := c.confFile()
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
		return
	}
	err = json.Unmarshal(content, c)
	if err != nil {
		panic(err.Error())
		return
	}
	core.CONFIG = c
}

func (c *Config) confFile() string {
	return "./config.json"
}
