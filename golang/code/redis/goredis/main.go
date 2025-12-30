package main
import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()


func main()  {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"192.168.188.105:7000", "192.168.188.105:7001", "192.168.188.105:7002",
				"192.168.188.110:7000", "192.168.188.110:7001", "192.168.188.110:7002"},
	})
	rdb.Ping(ctx)

	err := rdb.Set(ctx, "name", "lzle", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name", val)

	val2, err := rdb.Get(ctx, "missing_key").Result()
	if err == redis.Nil {
		fmt.Println("missing_key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("missing_key", val2)
	}

}
