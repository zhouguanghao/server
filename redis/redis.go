package redis

import (
	"fmt"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"github.com/mediocregopher/radix.v2/redis"
	_ "github.com/davyxu/cellnet/peer/redix"
	"time"
)

func init() {
	p := peer.NewGenericPeer("redix.Connector", "redis", "127.0.0.1:6379", nil)
	p.(cellnet.RedisConnector).SetPassword("zgh1625347")
	p.Start()

	for i := 0; i < 5; i++ {
		if p.(cellnet.PeerReadyChecker).IsReady() {
			break
		}
		time.Sleep(time.Millisecond * 200)
	}

	if !p.(cellnet.PeerReadyChecker).IsReady() {
		fmt.Println("Redis connector fail")
		return
	}

	op := p.(cellnet.RedisPoolOperator)
	op.Operate(func (rawClient interface{}) interface{} {
		client := rawClient.(*redis.Client)
		client.Cmd("SET", "mykey", "myvalue")

		v,err := client.Cmd("GET", "mykey").Str()
		if err != nil {
			panic(err)
		}
		fmt.Println("Value is :", v)
		return nil
	})
}
