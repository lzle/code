package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"go.lzle.cn/etcd/naming/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/naming"
	"time"
)

// First, register new endpoint addresses for a service:
func etcdAdd(c *clientv3.Client, service, addr string) error {
	r := &etcdnaming.GRPCResolver{Client: c}
	return r.Update(c.Ctx(), service, naming.Update{Op: naming.Add, Addr: addr})
}

// Dial an RPC service using the etcd gRPC resolver and a gRPC Balancer:
func etcdDial(c *clientv3.Client, service string) (*grpc.ClientConn, error) {
	r := &etcdnaming.GRPCResolver{Client: c}
	b := grpc.RoundRobin(r)
	return grpc.Dial(service, grpc.WithBalancer(b),grpc.WithBlock())
}

//// Optionally, force delete an endpoint:
//func etcdDelete(c *clientv3.Client, service, addr string) error {
//	r := &etcdnaming.GRPCResolver{Client: c}
//	return r.Update(c.Ctx(), "my-service", naming.Update{Op: naming.Delete, Addr: "1.2.3.4"})
//}
//
//// Or register an expiring endpoint with a lease:
//func etcdLeaseAdd(c *clientv3.Client, lid clientv3.LeaseID, service, addr string) error {
//	r := &etcdnaming.GRPCResolver{Client: c}
//	return r.Update(c.Ctx(), service, naming.Update{Op: naming.Add, Addr: addr}, clientv3.WithLease(lid))
//}


func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.188.105:2379", "192.168.188.105:22379", "192.168.188.105:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = etcdAdd(cli,"my-server","127.0.0.1:8972")
	if err != nil {
		fmt.Println(err)
	}

	conn,err := etcdDial(cli,"my-server")
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	//fmt.Println(conn)
	c := pb.NewGreeterClient(conn)
	fmt.Println(c)
	//// 调用服务端的SayHello
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "连长"})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
	}
	fmt.Printf("Greeting: %s !\n", r.Message)

}
