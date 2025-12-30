package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var ConfigParam *Config

type Config struct {
	MysqlConfig      Mysql      `yaml:"mysql"`
	LogPathConfig    LogPath    `yaml:"logPath"`
	AddrConfig       Addr       `yaml:"addr"`
	AmqpConfig       Amqp       `yaml:"rabbitmq"`
	RedisConfig      Redis      `yaml:"redis"`
	LocalRedisConfig Redis      `yaml:"localRedis"`
	MD5Config        MD5        `yaml:"md5"`
	UpLoadConfig     UpLoadPath `yaml:"uploadPath"`
}

type Addr struct {
	Host   string
	Port string
}

type Mysql struct {
	Host     string
	Port     int
	Db       string
	User     string
	PassWd   string
}

type Amqp struct {
	Host       string
	Port       int
	VHost      string
	User       string
	PassWd     string
	ServerId   string
	Exchange   string
}

type LogPath struct {
	Path   string
}

type Redis struct {
	Host       string
	Port       int
	Passwd	   string
}

type MD5 struct {
	Salt    string
}

type UpLoadPath struct {
	Path    string
}


func (addr *Addr) String() (s string) {
	return fmt.Sprintf("%s:%s",addr.Host,addr.Port)
}

func LoadConfig (path string) error {
	conf := new(Config)
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("yamlFile.Get err:", err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatal("Unmarshal:", err)
		return err
	}
	ConfigParam = conf
	return err
}