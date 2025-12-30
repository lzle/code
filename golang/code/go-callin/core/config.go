package core

type Config interface {
	// mysql
	MysqlUrl() string
	// rabbitmq
	AmqpUrl() string
	Exchange() string
	ServerId() string
	// cdr
	CdrDir() string
	// tts
	TTSUrl() string
	// log
	LogDir() string
}
