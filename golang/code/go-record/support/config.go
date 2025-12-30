package support

import (
	"encoding/json"
	"go-record/core"
	"io/ioutil"
)

type Config struct {
	LOG
	HTTP
	RECORD
}

type LOG struct {
	Dir string `json:"LOG_DIR"`
}

type HTTP struct {
	Addr  string `json:"HTTP_ADDR"`
	EnTls string `json:"EN_TLS"`
	Crt   string `json:"CRT"`
	Key   string `json:"KEY"`
}

type RECORD struct {
	Dirs []string `json:"RECORD_DIRS"`
}

func (c *Config) LogDir() string {
	return c.LOG.Dir
}

func (c *Config) RecordDirs() []string {
	return c.RECORD.Dirs
}

func (c *Config) HttpAddr() string {
	return c.HTTP.Addr
}

func (c *Config) EnTls() bool {
	if c.HTTP.EnTls == "1" {
		return true
	}
	return false
}
func (c *Config) Crt() string {
	return c.HTTP.Crt
}

func (c *Config) Key() string {
	return c.HTTP.Key
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
