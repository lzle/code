package taskHandle

import (
	log "go-commweb/log"
	"go-commweb/mysql"
	"math/rand"
	"time"
	sql2 "database/sql"
)

const (
	 ctiBusType = 7
)

type BusinessServerId struct {
	ServerId string `json:"serverid"`
}


type TaskInfo struct {
	id string `json:"id"`
}

type TaskServerIdInfo struct {
	ServerId string `json:"serverId"`
}

func getServerId (compId string) (serverId string) {
	log.LOGGER.Info("get serviceid from tx_route_business bustype[%d] compid[%s]", ctiBusType, compId)
	sql := "select serverid from pbx.tx_route_business where bustype=? and compid=? and enable=1;"
	DB := mysql.GetMysqlDb().DB
	rows, err := DB.Query(sql, ctiBusType, compId)
	if err != nil {
		return
	}
	defer func() {
		if rows != nil {
			rows.Close() //可以关闭掉未scan连接一直占用
		}
	}()

	var serverIds []string
	for rows.Next() {
		businessServerId := new(BusinessServerId)
		err = rows.Scan(&businessServerId.ServerId) //不scan会导致连接不释放
		if err != nil {
			log.LOGGER.Error("Scan failed,err:%v", err)
			return
		}
		if businessServerId.ServerId != ""{
			serverIds = append(serverIds,businessServerId.ServerId)
		}
	}
	if len(serverIds) > 0 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return serverIds[r.Intn(len(serverIds))]
	}else{
		sql := "select serverid from pbx.tx_route_business where bustype=? and compid='' and enable=1;"
		_rows, err := DB.Query(sql, ctiBusType)
		if err != nil {
			return
		}
		defer func() {
			if _rows != nil {
				_rows.Close() //可以关闭掉未scan连接一直占用
			}
		}()
		for _rows.Next() {
			businessServerId := new(BusinessServerId)
			err = _rows.Scan(&businessServerId.ServerId) //不scan会导致连接不释放
			if err != nil {
				log.LOGGER.Error("Scan failed,err:%v", err)
				return
			}
			if businessServerId.ServerId != ""{
				serverIds = append(serverIds,businessServerId.ServerId)
			}
		}
		if len(serverIds) > 0 {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			return serverIds[r.Intn(len(serverIds))]
		}
	}
	return
}

func checkTask (message *TaskParam, serverId string){
	sql := "select id from tx_task where taskid=?"
	DB := mysql.GetMysqlDb().DB
	row := DB.QueryRow(sql, message.TaskId)
	taskInfo := new(TaskInfo)
	if err := row.Scan(&taskInfo.id); err == sql2.ErrNoRows {
		sql := "insert into tx_task(taskid, serverid, phoneid) VALUES(?,?,?);"
		_, err := DB.Exec(sql, message.TaskId, serverId, 0)
		if err != nil {
			log.LOGGER.Error("insert task [%s] error [%v]", message.TaskId, err)
			return
		}
		log.LOGGER.Info("execute sql[%s] args[%s, %s, %d]", sql, message.TaskId, serverId, 0)
	}else if err != nil{
		log.LOGGER.Error("scan failed, err:%v", err)
		return
	}else{
		log.LOGGER.Info("taskid already exist value")
	}
}

func getCallTaskServerId(taskId string) (serverId string)  {
	sql := "select serverid from tx_task where taskid=?"
	DB := mysql.GetMysqlDb().DB
	row := DB.QueryRow(sql, taskId)
	taskServerId := new(TaskServerIdInfo)
	err := row.Scan(&taskServerId.ServerId) //不scan会导致连接不释放
	if err != nil {
		log.LOGGER.Error("Scan failed,err:%v", err)
		return ""
	}else{
		return taskServerId.ServerId
	}
}