package skillGroupHandle

import (
	"go-commweb/redis"
)

func redisHget (keys, field string) (res string) {
	if keys != "" && field != "" {
		redisdb := redis.GetLocalRedisDb()
		res = redisdb.HashGet(keys, field)
	}
	return
}

func redisKeys (keys string) (res interface{}) {
	if keys != "" {
		redisdb := redis.GetLocalRedisDb()
		res = redisdb.Keys(keys)
	}
	return
}