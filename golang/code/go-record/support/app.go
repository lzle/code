package support

import (
	"go-record/src"
	"runtime"
)

type App struct {

}


func (a *App) Start()  {
	// 可以接收参数
	a.Handle()
}

func (a *App) Handle()  {
	config := new(Config)
	config.Init()

	logger := new(Logger)
	logger.Init()

	web := new(src.Web)
	web.Run()

	runtime.Goexit()
}


