package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/namespace"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.188.105:2379", "192.168.188.105:22379", "192.168.188.105:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {

	}

	// Package namespace is a clientv3 wrapper that translates all keys to begin
	// with a given prefix.
	//
	// First, create a client:
	//
	//	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{"localhost:2379"}})
	//	if err != nil {
	//		// handle error!
	//	}
	//
	// Next, override the client interfaces:
	//
	//unprefixedKV := cli.KV
	cli.KV = namespace.NewKV(cli.KV, "my-prefix/")
	cli.Watcher = namespace.NewWatcher(cli.Watcher, "my-prefix/")
	cli.Lease = namespace.NewLease(cli.Lease, "my-prefix/")

	// Now calls using 'cli' will namespace / prefix all keys with "my-prefix/":
	//
	cli.Put(context.TODO(), "sample_key", "1234")
	//resp, _ := unprefixedKV.Get(context.TODO(), "my-prefix/sample_key")
	//fmt.Printf("%s\n", resp.Kvs[0].Value)
	//// Output: 123
	//unprefixedKV.Put(context.TODO(), "my-prefix/sample_key", "4567")
	resp, _ := cli.Get(context.TODO(),"sample_key")
	fmt.Printf("%s\n", resp.Kvs[0].Value)
	//	// Output: 456
	//
}