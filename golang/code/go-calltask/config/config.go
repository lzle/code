package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type http struct {
	Host string `json:"host1"`
	Port int    `json:"port1"`
}

type log struct {
	Path string `json:"log_path"`
}

type cdr struct {
	Path string `json:"cdr_path"`
}

type recognition struct {
	Url string `json:"recognition_url"`
}

type record struct {
	AbsRecordPath string `json:"abs_record_path"`
	RecordDirName string    `json:"record_dirname"`
}

type db struct {
	Host   string `json:"host"`
	Port   int    `json:"port"`
	Db     string `json:"db"`
	User   string `json:"user"`
	Passwd string `json:"passwd"`
}

type amqp struct {
	Host   string `json:"host"`
	Port   int    `json:"port"`
	Vhost  string `json:"vhost"`
	User   string `json:"user"`
	Passwd string `json:"passwd"`
	ServerId string `json:"serverid"`
	Exchange string `json:"exchange"`
}

type redis struct{
	Host   string `json:"host"`
	Port   int    `json:"port"`
	Password   string    `json:"passwd"`
}

type localRedis struct{
	Host   string `json:"host"`
	Port   int    `json:"port"`
	Password   string    `json:"passwd"`
}

type task struct {
	UpdateTaskDetailUrl string`json:"update_task_detail_url"`
	GetTaskAgentUrl string`json:"get_task_agent_url"`
	ActiveTaskUrl string`json:"active_task_url"`
}

type comp struct {
	UpdateCompDetailUrl string`json:"update_comp_detail_url"`
}

type accessNumber struct {
	ApplyNumber string`json:"apply_number"`
	ReleaseNumber string`json:"release_number"`
}

type numberPrefix struct {
	NumberPrefix []string `json:"number_prefix"`
}


type baseConfig struct {
	cdr
	log
	http
	record
	db   `json:"dbConfig"`
	amqp `json:"amqpConfig"`
	task
	comp
	redis `json:"redisConfig"`
	localRedis `json:"localRedis"`
	numberPrefix
	recognition
}

var (
	LogConfig *log
	CdrConfig *cdr
	RecordConfig *record
	DBConfig   *db
	AmqpConfig *amqp
	TaskConfig *task
	CompConfig *comp
	RedisConfig *redis
	LocalRedisConfig *localRedis
	NumberPrefixConfig *numberPrefix
	RecognitionConf    *recognition
)

func InitConfig(filename string) (err error) {
	var (
		content []byte
		conf    baseConfig
	)

	if content, err = ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
		return
	}

	if err = json.Unmarshal(content, &conf); err != nil {
		fmt.Println(err)
		return
	}
	//HttpConfig = &conf.http
	RecordConfig = &conf.record
	CdrConfig = &conf.cdr
	LogConfig = &conf.log
	DBConfig = &conf.db
	AmqpConfig = &conf.amqp
	TaskConfig = &conf.task
	CompConfig = &conf.comp
	RedisConfig = &conf.redis
	LocalRedisConfig = &conf.localRedis
	NumberPrefixConfig = &conf.numberPrefix
	RecognitionConf = &conf.recognition
	return
}
