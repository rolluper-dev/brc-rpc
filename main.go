package main

import (
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/roullper-dev/brc-rpc/config"
	"github.com/roullper-dev/brc-rpc/resource/db"
	"github.com/roullper-dev/brc-rpc/resource/log"
	"github.com/roullper-dev/brc-rpc/router"
)

func main() {
	config.Init()
	log.Init(viper.GetString("env"), viper.GetString("logDir"))
	dbConf := viper.GetStringMapString("database")
	db.Init(dbConf["user"], dbConf["password"], dbConf["host"], dbConf["port"], dbConf["ord"], dbConf["indexer"])
	redisConf := viper.GetStringMapString("redis")
	db.InitRedis(redisConf["host"], redisConf["port"], redisConf["password"])

	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           2 * time.Minute,
	}))

	router.Register(engine)
	endless.ListenAndServe(viper.GetString("address"), engine)
}
