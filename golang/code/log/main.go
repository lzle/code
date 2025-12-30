package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var LOGGER *Logger

var COUNT = 0


type Logger struct {
	//lock. when maintaining, cannot write
	sync.RWMutex
	//extend logger
	logger *log.Logger
	//log file.
	file *os.File
}

func (l *Logger) Init() {
	l.openFile()
	expression := "*/1 * * * *"
	cronJob := cron.New()
	_,err := cronJob.AddFunc(expression,l.rotation)
	if err != nil {
		fmt.Println("create cron job " + err.Error())
	}
	cronJob.Start()
	l.Info("[cron job] Every hour rotation log file.")
	LOGGER = l
}

func (l *Logger) Log(prefix string, format string, v ...interface{}) {
	content := fmt.Sprintf(format + "\r\n" , v...)
	l.Lock()
	defer l.Unlock()
	l.logger.SetPrefix(prefix)
	err := l.logger.Output(3,content)
	if err != nil {
		fmt.Printf("occur error while logging %s \r\n", err.Error())
	}
}

func (l *Logger) Debug(format string, v ...interface{}) {
	l.Log("[DEBUG]",  format, v...)
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.Log("[INFO ]",  format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.Log("[Error ]",  format, v...)
}

func (l *Logger) openFile() {
	fmt.Printf("use log file %s\r\n", l.logFile())
	f, err := os.OpenFile(l.logFile(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("cannot open log file " + err.Error())
	}

	l.logger = log.New(f,"", log.Ldate| log.Lmicroseconds |log.Lshortfile)
	if l.logger == nil {
		fmt.Printf("Error: cannot create logger \r\n")
	}

	l.file = f
}


func (l *Logger) closeFile() {
	if l.file != nil {
		err := l.file.Close()
		if err != nil {
			fmt.Println("occur error while closing log file: " + err.Error())
		}
	}
}


func (l *Logger) logFile() string {
	return  "./my.log"
}




func (l *Logger) rotation(){
	l.Info("rotation log")
	l.Lock()
	defer l.Unlock()
	COUNT++
	destPath := l.logFile() +"@" + strconv.Itoa(COUNT)
	//rename the log file.
	err := os.Rename(l.logFile(), destPath)
	if err != nil {
		l.Error("occur error while renaming log file %s %s", destPath, err.Error())
	}
	l.closeFile()
	//reopen
	l.openFile()
}


func main()  {
	logger := new(Logger)
	logger.Init()

	i := 0
	for {
		i++
		LOGGER.Info("%d",i)
		time.Sleep(time.Millisecond*1000)
	}
}