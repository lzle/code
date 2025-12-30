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
	// set
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Delete(ctx, "sample_key", )
	cancel()
	if err != nil {
		// handle error!
	}
	fmt.Println(resp)
}