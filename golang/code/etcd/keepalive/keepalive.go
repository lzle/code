package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.188.105:2379", "192.168.188.105:22379", "192.168.188.105:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	defer cli.Close()
	// 创建一个5秒的租约
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		// handle error!
	}
	// 5秒钟之后, /nazha/ 这个key就会被移除
	cli.Put(context.TODO(), "sample_key", "lease", clientv3.WithLease(resp.ID))
	// the key 'foo' will be kept forever
	ch, err := cli.KeepAlive(context.TODO(), resp.ID)
	if err != nil {
		fmt.Println(err)
	}
	for {
		ka := <-ch
		fmt.Println("ttl:", ka.TTL)
	}
}