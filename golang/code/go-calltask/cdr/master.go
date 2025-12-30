package cdr

import (
	"encoding/json"
	"fmt"
	"go-calltask/config"
	log "go-calltask/log"
	"os"
)

var (
	CdrInstance *Cdr
)

type CdrMaster struct {
	RequestId string `json:"requestid"`
	CallId string `json:"callid"`
	CallType int `json:"calltype"`
	Direction int `json:"direction"`
	Caller string `json:"caller"`
	Callee string `json:"callee"`
	Dh string `json:"dh"`
	CallResult int `json:"callresult"`
	StimeStamp int64 `json:"stimestamp,omitempty"`
	Stime string `json:"stime"`
	Atime string `json:"atime"`
	Etime string `json:"etime"`
	RecPath string `json:"rec_path"`
	ServerId string `json:"serverid"`
	TaskId string `json:"taskid"`
	AgentId string `json:"agentid"`
	AgentGrpId string `json:"agentgrpid"`
	CompId string `json:"compid"`
	FeeTime int `json:"feetime"`
	TotalTime int64 `json:"totaltime"`
	Fee int `json:"fee"`
	CallerArea string `json:"caller_area"`
	CalleeArea string `json:"callee_area"`
	Flags int `json:"flags"`
	IfJandle int `json:"ifhandle"`
	Hanguper int `json:"hanguper"`
	Userphone string `json:"userphone"`
	DetailCount int `json:"detail_count"`
}

type Cdr struct {
	CdrQueue chan *CdrMaster
}

func (cdr *Cdr) write(cdrMaster *CdrMaster){
	defer func(){
		if r := recover();r !=nil{
			log.LOGGER.Error("%v",r)
		}
	}()

	callId := cdrMaster.CallId
	compId := cdrMaster.CompId
	data,err := json.Marshal(cdrMaster)
	if err != nil {
		log.LOGGER.Error(err.Error())
		return
	}
	log.LOGGER.Info("cdr master %s",string(data))
	fileName := fmt.Sprintf("%s%s_%s",config.CdrConfig.Path,compId,callId)
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		log.LOGGER.Error(err.Error())
		return
	}
	_, err = f.Write(data)
	if err != nil {
		log.LOGGER.Error(err.Error())
		return
	}
	log.LOGGER.Info("master cdr written successfully")
	err = os.Rename(fileName, fileName + ".master")
	if err != nil {
		log.LOGGER.Error(err.Error())
	}
}

func (cdr *Cdr) runLoop(){
	for {
		cdrMaster := <- cdr.CdrQueue
		cdr.write(cdrMaster)
	}
}

func GetCdrInstance() (cdr *Cdr){
	return CdrInstance
}

func Run()  {
	cdr := new(Cdr)
	cdr.CdrQueue = make(chan *CdrMaster,2000)
	CdrInstance = cdr
	go cdr.runLoop()
}