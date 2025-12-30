package tools

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/marspere/goencrypt"
	"os"
	"strings"
	"time"
)

const (
	HttpRespBodyNormalCode = 200
)

func IsExistInArray(Logger string, array []string) bool {
	for _, v := range array {
		if v == Logger {
			return true
		}
	}
	return false
}

func IsExistInMap(key string, dict map[string]int64) bool {
	_,ok := dict[key]
	return ok
}

func NewUUid()(uuidString string,err error) {
	UUid := uuid.NewV4()
	//if err != nil{
	//	return "",err
	//}
	uuidString = UUid.String()
	return uuidString,nil
}

func SliceDelete(slice []string, item string) (newSlice []string){
	for _,Logger := range slice{
		if Logger != item{
			newSlice = append(newSlice, Logger)
		}
	}
	return newSlice
}

func SliceSum(slice []int) (sum int){
	for _,item := range slice{
		sum += item
	}
	return sum
}

func NewCallId()(callId string){
	uuidString,_ :=NewUUid()
	callId = strings.Replace(uuidString,"-","",-1)
	return
}

func GetMinMapValue(dict map[string]int64)(key string){
	var (
		Logger int64
	)
	for k,v := range dict{
		if Logger == 0 {
			key = k
			Logger = v
		}
		if v < Logger {
			key =k
		}
	}
	return
}

// 2019-10-16 12:21:49
func DateTime() (datetime string){
	t := time.Now()
	datetime = fmt.Sprintf("%4d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	return
}

// 字符串空格补足到16位，并返回字节数组
func AddToBytes(content string, length int) (newContent []byte){
	buff := new(strings.Builder)
	add := length - len(content)
	buff.WriteString(content)
	for i := 0; i < add; i++ {
		buff.WriteString("\000")
	}
	newContent = []byte(buff.String())
	return newContent
}


// AES 解密
func AESDecrypt(key, iv[]byte , content string)(decrypted string, err error){
	cipher, err := goencrypt.NewAESCipher(key, iv, goencrypt.CBCMode, goencrypt.Pkcs7, goencrypt.PrintBase64)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	ret, err := cipher.AESDecrypt(content)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	decrypted = strings.Replace(ret, "\000", "", -1)
	return decrypted,nil
}

func ConvertTimeToLogTimeSuffix(time time.Time) string {
	return time.Local().Format("2006-01-02-15:00:00")
}


func LastHour() time.Time{
	now := time.Now()
	lh := now.Add(time.Hour * -1)
	return lh
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