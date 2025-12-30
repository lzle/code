package support

import (
	"encoding/json"
	"fmt"
	"go-callin/core"
	"go-callin/utils"
	"os"
)

type Cdr struct {
	queue chan *core.Master
}

func (cdr *Cdr) write(cdrMaster *core.Master) {
	defer func() {
		if r := recover(); r != nil {
			core.LOGGER.Error("%v", r)
		}
	}()

	callId := cdrMaster.CallId
	compId := cdrMaster.CompId
	data, err := json.Marshal(cdrMaster)
	if err != nil {
		core.LOGGER.Error("%v", err)
		return
	}
	core.LOGGER.Info("cdr master %s", string(data))
	fileName := fmt.Sprintf("%s%s_%s", cdr.dirPath(), compId, callId)
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		core.LOGGER.Error("%v", err)
		return
	}
	_, err = f.Write(data)
	if err != nil {
		core.LOGGER.Error("%v", err)
		return
	}
	core.LOGGER.Info("[INFO] master cdr written successfully")
	err = os.Rename(fileName, fileName+".master")
	if err != nil {
		core.LOGGER.Error("%v", err)
	}
}

func (cdr *Cdr) dirPath() string {
	cdrPath := core.CONFIG.CdrDir()
	exists := utils.PathExists(cdrPath)

	if !exists {
		err := os.MkdirAll(cdrPath, 0777)
		if err != nil {
			fmt.Println("error while mkdir " + err.Error())
		}
	}
	return cdrPath
}

func (cdr *Cdr) runLoop() {
	for {
		cdrMaster := <-cdr.queue
		cdr.write(cdrMaster)
	}
}

func (cdr *Cdr) GetQueue() chan *core.Master {
	return cdr.queue
}


func (cdr *Cdr) Init() {
	cdr.queue = make(chan *core.Master,1000)
	core.CDR = cdr
	cdr.dirPath()
	go cdr.runLoop()
}
