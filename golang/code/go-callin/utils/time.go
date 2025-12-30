package utils

import (
	"fmt"
	"time"
)

func ConvertTimeToLogTimeSuffix(time time.Time) string {
	return time.Local().Format("2006-01-02-15:00:00")
}


func LastHour() time.Time{
	now := time.Now()
	lh := now.Add(time.Hour * -1)
	return lh
}

// 2019-10-16 12:21:49
func DateTime() (datetime string){
	t := time.Now()
	datetime = fmt.Sprintf("%4d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	return
}