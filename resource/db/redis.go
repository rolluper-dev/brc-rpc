package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/roullper-dev/brc-rpc/constant"
	"log"
	"time"
)

var (
	rds                *redis.Client
	handlerBlockHeight string
)

func InitRedis(addr, port, pwd string) {
	rds = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", addr, port),
		Password: pwd,
	})
	result, err := rds.Ping().Result()
	if err != nil {
		panic("failed to connect redis" + err.Error())
	}
	log.Println("connect redis", "ping", result)
	go func() {
		for {
			rlt, err := rds.Get(constant.BlockHandlerHeight).Result()
			if err != nil {
				log.Println("getHandlerBlockHeight failed", "err", err)
				time.Sleep(time.Minute)
			}
			handlerBlockHeight = rlt
			time.Sleep(time.Minute)
		}
	}()
}

func GetDb() *redis.Client {
	return rds
}

func GetHandlerBlockHeight() string {
	return handlerBlockHeight
}
