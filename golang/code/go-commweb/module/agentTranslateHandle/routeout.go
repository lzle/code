package agentTranslateHandle

import (
	"fmt"
	"go-commweb/mysql"
	"go-commweb/redis"
	log "go-commweb/log"
)

type BusinessServerId struct {
	ServerId string `json:"serverid"`
}

func getAgentServerId(compId, agentId string) (serverId string) {
	key := fmt.Sprintf("agent:%s:%s", compId, agentId)
	redisdb := redis.GetRedisDb()
	serverId = redisdb.HashGet(key, "serverid")
	return
}

func getServerId(compId, busType, deviceId string) (serverId string) {
	log.LOGGER.Info("get esl serviceId from tx_route_business bustype[%s] compid[%s]", busType, compId)
	sql := "SELECT serverid FROM pbx.tx_route_business where bustype = ?  and compid = ? and deviceid = ? and enable = 1"
	DB := mysql.GetMysqlDb().DB
	rows, err := DB.Query(sql, busType, compId, deviceId)
	if err != nil {
		return
	}
	defer func() {
		if rows != nil {
			rows.Close() //可以关闭掉未scan连接一直占用
		}
	}()
	for rows.Next() {
		businessServerId := new(BusinessServerId)
		err = rows.Scan(&businessServerId.ServerId) //不scan会导致连接不释放
		if err != nil {
			log.LOGGER.Error("scan failed %s", err.Error())
			return
		}
		serverId = businessServerId.ServerId
	}
	if serverId != "" {
		return
	}else{
		sql := "SELECT serverid FROM pbx.tx_route_business where bustype = ?  and compid = ? and deviceid = '' and enable = 1"
		rows, err := DB.Query(sql, busType, compId)
		if err != nil {
			return
		}
		defer func() {
			if rows != nil {
				rows.Close() //可以关闭掉未scan连接一直占用
			}
		}()
		for rows.Next() {
			businessServerId := new(BusinessServerId)
			err = rows.Scan(&businessServerId.ServerId) //不scan会导致连接不释放
			if err != nil {
				log.LOGGER.Error("scan failed %s", err.Error())
				return
			}
			serverId = businessServerId.ServerId
		}
		if serverId != "" {
			return
		}else{
			sql := "SELECT serverid FROM pbx.tx_route_business where bustype = ?  and compid = '' and deviceid = '' and enable = 1"
			rows, err := DB.Query(sql, busType)
			if err != nil {
				return
			}
			defer func() {
				if rows != nil {
					rows.Close() //可以关闭掉未scan连接一直占用
				}
			}()
			for rows.Next() {
				businessServerId := new(BusinessServerId)
				err = rows.Scan(&businessServerId.ServerId) //不scan会导致连接不释放
				if err != nil {
					log.LOGGER.Error("scan failed %s", err.Error())
					return
				}
				serverId = businessServerId.ServerId
			}
			if serverId != "" {
				return
			}
		}
	}
	return
}