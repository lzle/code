package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-commweb/config"
	"os"
	"strings"
	"time"
)

func ResponseJson(c *gin.Context, httpCode int, dataCode int, reason string, data interface{}) {
	dataMap := make(map[string]interface{})
	dataMap["code"] = dataCode
	dataMap["reason"] = reason
	if data != nil && data != "" {
		dataMap["result"] = data
	}
	c.JSON(httpCode, dataMap)
}

// 2019-10-16 12:21:49
func DateTime() (datetime string){
	t := time.Now()
	datetime = fmt.Sprintf("%4d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	return
}

// 获取时间戳
func GetTimestap(timeStr string) int64 {
	time, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err == nil {
		return time.Unix()
	}
	return 0
}

// 删除切片中的元素
func SliceDelete(slice []string, item string) (newSlice []string){
	for _, Logger := range slice{
		if Logger != item{
			newSlice = append(newSlice, Logger)
		}
	}
	return newSlice
}

// 判断字符串中是否包含
func StringContain (s, str string) bool {
	if res := strings.Index(s, str); res != -1 {
		return true
	}
	return false
}

func CreateMD5(compId, taskId string) string {
	m5 := md5.New()
	m5.Write([]byte(config.ConfigParam.MD5Config.Salt))
	m5.Write([]byte(compId))
	m5.Write([]byte(taskId))
	st := m5.Sum(nil)
	return hex.EncodeToString(st)
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else {
		if os.IsNotExist(err) {
			return false
		} else {
			panic(err.Error())
		}
	}
}

func ConvertTimeToLogTimeSuffix(time time.Time) string {
	return time.Local().Format("2006-01-02-15:00:00")
}


func LastHour() time.Time{
	now := time.Now()
	lh := now.Add(time.Hour * -1)
	return lh
}