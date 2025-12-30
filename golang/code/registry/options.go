package registry

import "time"

type Option func(*Options)


type Options struct {
	Addr []string

	TimeOut time.Duration

	HeartBeat int64

	RegistryPath string
}


func WithAddr(addr []string) Option{
	return func(options *Options) {
		options.Addr = addr
	}
}

func WithTimeOut(timeOut time.Duration) Option {
	return func(options *Options) {
		options.TimeOut = timeOut
	}
}

func WithHeartBeat(heartBeat int64) Option {
	return func(options *Options) {
		options.HeartBeat = heartBeat
	}
}


func WithRegistryPath(path string) Option {
	return func(options *Options) {
		options.RegistryPath = path
	}
}