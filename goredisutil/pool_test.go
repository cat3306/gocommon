package goredisutil

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

func TestNewRedisClients(t *testing.T) {
	clients := NewRedisClients(&ClientConf{
		Options: &redis.Options{
			Addr: "127.0.0.1:6379",
		},
		DB: []int{0, 1, 2, 3, 4, 5},
	})
	client := clients.Select(1)
	fmt.Println(client.String())
	fmt.Println(client.Get(context.Background(), "haha").Result())
}
