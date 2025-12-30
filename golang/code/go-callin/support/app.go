package support

import "go-callin/src"

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

	mysql := new(Mysql)
	mysql.Init()

	amqp := new(Amqp)
	amqp.Init()

	cdr := new(Cdr)
	cdr.Init()

	router := new(src.Router)
	router.Init()

	center := new(src.Center)
	center.Execute()

	forever := make(chan bool)
	<- forever
}


