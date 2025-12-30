package core

type Cdr interface {
	GetQueue() chan *Master
}

type Master struct {
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
	IfHandle int `json:"ifhandle"`
	Hanguper int `json:"hanguper"`
	Userphone string `json:"userphone"`
	DetailCount int `json:"detail_count"`
}