package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.188.105:2379", "192.168.188.105:22379", "192.168.188.105:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer cli.Close()

	// 创建两个单独的会话用来演示锁竞争
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		fmt.Println(err)
	}
	defer s1.Close()
	m1 := concurrency.NewMutex(s1, "/my-lock/")

	//s2, err := concurrency.NewSession(cli)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer s2.Close()
	//m2 := concurrency.NewMutex(s2, "/my-lock/")

	if err := m1.Lock(context.TODO()); err != nil {
		fmt.Println(err)
	}
	fmt.Println("acquired lock for s1")

	//m2Locked := make(chan struct{})
	//go func() {
	//	defer close(m2Locked)
	//	// 等待直到会话s1释放了/my-lock/的锁
	//	if err := m2.Lock(context.TODO()); err != nil {
	//		fmt.Println(err)
	//	}
	//}()
	time.Sleep(time.Second*10)
	if err := m1.Unlock(context.TODO()); err != nil {
		fmt.Println(err)
	}
	fmt.Println("released lock for s1")

	//<-m2Locked
	//fmt.Println("acquired lock for s2")
}