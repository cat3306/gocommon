package goredisutil

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisClientPool struct {
	clients []*redis.Client
}
type ClientConf struct {
	Options *redis.Options
	DB      []int
}

func NewRedisClients(c *ClientConf) *RedisClientPool {
	r := RedisClientPool{
		clients: make([]*redis.Client, 0, len(c.DB)),
	}

	for i := 0; i < len(c.DB); i++ {
		client := redis.NewClient(&redis.Options{
			Network:            c.Options.Network,
			Addr:               c.Options.Addr,
			Dialer:             c.Options.Dialer,
			OnConnect:          c.Options.OnConnect,
			Username:           c.Options.Username,
			Password:           c.Options.Password,
			DB:                 c.DB[i],
			MaxRetries:         c.Options.MaxRetries,
			MinRetryBackoff:    c.Options.MinRetryBackoff,
			MaxRetryBackoff:    c.Options.MaxRetryBackoff,
			DialTimeout:        c.Options.DialTimeout,
			ReadTimeout:        c.Options.ReadTimeout,
			WriteTimeout:       c.Options.WriteTimeout,
			PoolFIFO:           c.Options.PoolFIFO,
			PoolSize:           c.Options.PoolSize,
			MinIdleConns:       c.Options.MinIdleConns,
			MaxConnAge:         c.Options.MaxConnAge,
			PoolTimeout:        c.Options.PoolTimeout,
			IdleTimeout:        c.Options.IdleTimeout,
			IdleCheckFrequency: c.Options.IdleCheckFrequency,
			TLSConfig:          c.Options.TLSConfig,
			Limiter:            c.Options.Limiter,
		})
		if client == nil {
			panic("client nil")
		}
		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			panic(err)
		}
		r.clients = append(r.clients, client)
	}
	return &r
}
func (r *RedisClientPool) Select(args ...int) *redis.Client {
	db := 0
	if len(args) != 0 {
		db = args[0]
	}
	if db < 0 || db > len(r.clients)-1 {
		panic("db invalid")
	}
	return r.clients[db]
}
