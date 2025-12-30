package core

type Config interface {
	// log
	LogDir() string
	// record
	RecordDirs() []string
	// web addr
	HttpAddr() string
	// https
	EnTls() bool
	// crt
	Crt() string
	// key
	Key() string
}
