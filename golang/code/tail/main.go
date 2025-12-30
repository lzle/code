package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main() {
	filename := "../log/my.log"
	tailFile, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,
		Follow:    true,
		// 每次从文件末尾读取
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})

	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	for true {
		msg, ok := <- tailFile.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename: %s\n", tailFile.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		//time.Sleep(time.Second)
		fmt.Println("msg:", msg.Text)
	}
}
