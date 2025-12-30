package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"go.lzle.cn/etcd/naming/pb"
	"google.golang.org/grpc"
	"time"
)

// Dial an RPC service using the etcd gRPC resolver and a gRPC Balancer:
func etcdDial(c *clientv3.Client, service string) (*grpc.ClientConn, error) {
	r := &etcdnaming.GRPCResolver{Client: c}
	b := grpc.RoundRobin(r)
	return grpc.Dial(service, grpc.WithBalancer(b), grpc.WithBlock(), grpc.WithInsecure())
}

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.188.105:2379", "192.168.188.105:22379", "192.168.188.105:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
	}

	conn, err := etcdDial(cli, "my-server")
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	//fmt.Println(conn)
	c := pb.NewGreeterClient(conn)
	fmt.Println(c)
	//// 调用服务端的SayHello
	for i := 0; i < 10; i++ {
		r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "连长"})
		if err != nil {
			fmt.Printf("could not greet: %v", err)
		}
		fmt.Printf("Greeting: %s !\n", r.Message)
		time.Sleep(time.Second)
	}
}
