package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"server/config"
	"time"
)

var Redisdb *redis.Client

func init() {
	cfg := new(config.Config)
	cfg.InitConfig("server.config")
	redis_addr := cfg.Read("redis", "address")
	Redisdb = redis.NewClient(&redis.Options{
		Addr: redis_addr,
		Password: "zgh1625347",
		DialTimeout: 10 * time.Second,
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize: 10,
		PoolTimeout: 30 * time.Second,
	})
	defer Redisdb.Close()
	pong,err := Redisdb.Ping().Result()
	fmt.Println(pong, err)
	Redisdb.Set("mykey", "zhou", 3 * time.Second)
	v,err := Redisdb.Get("mykey").Result()
	fmt.Println(v)
}

