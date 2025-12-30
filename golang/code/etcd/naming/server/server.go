package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"go.lzle.cn/etcd/naming/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/naming"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

// First, register new endpoint addresses for a service:
func etcdAdd(c *clientv3.Client, service, addr string) error {
	r := &etcdnaming.GRPCResolver{Client: c}
	return r.Update(c.Ctx(), service, naming.Update{Op: naming.Add, Addr: addr})
}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("request 9000")
	return &pb.HelloReply{Message: "Hello 9000" + in.Name}, nil
}


func main()  {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.188.105:2379", "192.168.188.105:22379", "192.168.188.105:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = etcdAdd(cli,"my-server","127.0.0.1:9000")
	if err != nil {
		fmt.Println(err)
	}

	// 监听本地的8972端口
	lis, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer() // 创建gRPC服务器
	pb.RegisterGreeterServer(s, &server{}) // 在gRPC服务端注册服务

	reflection.Register(s) //在给定的gRPC服务器上注册服务器反射服务
	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	fmt.Println("server start")
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
