package redis

import (
	"fmt"
	"sync"

	"github.com/gomodule/redigo/redis"

	"go-commweb/config"
	log "go-commweb/log"
)

var (
	RedisDBInstance      *RedisDB
	LocalRedisDBInstance *LocalRedisDB
	RedisConfig          config.Redis
	LocalRedisConfig     config.Redis
	mutexRedis           sync.Mutex
)

type RedisDB struct {
	addr string
	password string
	conn redis.Conn
}

type LocalRedisDB struct {
	RedisDB
}

func (rdb *RedisDB)  connect(netWork, address, password string) (error) {
	conn, err := redis.Dial(netWork, address,redis.DialPassword(password))
	if err != nil {
		log.LOGGER.Error("%v",err.Error())
		return err
	}
	rdb.addr = address
	rdb.conn = conn
	rdb.password=password
	return nil
}

func (rdb *RedisDB) reconnect() {
	mutexRedis.Lock()
	rdb.conn.Close()
	err := rdb.connect("tcp", rdb.addr, rdb.password)
	mutexRedis.Unlock()
	if err != nil {
		return
	}
}

func (rdb *RedisDB) HashGet(key string, field string) (result string) {
	mutexRedis.Lock()
	res, err := rdb.conn.Do("HGET", key, field)
	mutexRedis.Unlock()
	if err != nil {
		log.LOGGER.Error("%v",err.Error())
		rdb.reconnect()
		return
	}
	ascii_res, ok := res.([]uint8)
	if ok == true{
		result = string(ascii_res)
		return result
	} else {
		return ""
	}
}

func (rdb *RedisDB) HashMSet(key string, content map[string]string) (result interface{}) {
	var args = []interface{}{key}
	for k, v := range content {
		args = append(args, k, v)
	}
	mutexRedis.Lock()
	res, err := rdb.conn.Do("HMSET", args...)
	mutexRedis.Unlock()
	if err != nil {
		log.LOGGER.Error("%v",err.Error())
		rdb.reconnect()
		return res
	}
	return res
}

func (rdb *RedisDB) Keys (key string) (result interface{}) {
	mutexRedis.Lock()
	res, err := rdb.conn.Do("KEYS", key)
	mutexRedis.Unlock()
	if err != nil {
		log.LOGGER.Error("%v",err.Error())
		rdb.reconnect()
		return nil
	}
	return res

}

func (rdb *RedisDB) Del(key string) (result interface{}) {
	mutexRedis.Lock()
	res, err := rdb.conn.Do("DEL", key)
	mutexRedis.Unlock()
	if err != nil {
		log.LOGGER.Error("%v",err.Error())
		rdb.reconnect()
		return res
	}
	return res
}

func GetRedisDb() (*RedisDB) {
	return RedisDBInstance
}

func GetLocalRedisDb() (*LocalRedisDB) {
	return LocalRedisDBInstance
}

func Run() {
	RedisConfig = config.ConfigParam.RedisConfig
	redisdb := new(RedisDB)
	url := fmt.Sprintf("%s:%d", RedisConfig.Host, RedisConfig.Port)
	err := redisdb.connect("tcp", url,RedisConfig.Passwd)
	if err != nil {
		return
	}
	log.LOGGER.Info("connect redis success addr[%v]", url)
	RedisDBInstance = redisdb

	LocalRedisConfig = config.ConfigParam.LocalRedisConfig
	localRedisdb := new(LocalRedisDB)
	localUrl:= fmt.Sprintf("%s:%d", LocalRedisConfig.Host, LocalRedisConfig.Port)
	err = localRedisdb.connect("tcp", localUrl, LocalRedisConfig.Passwd)
	if err != nil {
		return
	}
	log.LOGGER.Info("connect local redis success addr[%v]", localUrl)
	LocalRedisDBInstance = localRedisdb

}
