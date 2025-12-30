package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {
	c1, err := redis.Dial("tcp", "192.168.188.105:7000")
	if err != nil {
		log.Fatalln(err)
	}
	defer c1.Close()

	rec1, err := c1.Do("Get", "name")
	fmt.Println(rec1)
	rec2, err := c1.Do("Get", "name")
	fmt.Println(rec2)
}