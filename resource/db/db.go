package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var (
	db    *gorm.DB
	ordDB *gorm.DB
)

func Init(user, password, host, port, ordName, indexerName string) {
	log.Println("connecting MySQL ... ", host)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, ordName)
	mdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
				Colorful: false,
			},
		),
	})
	if err != nil {
		panic(err)
	}
	if mdb == nil {
		panic("failed to connect database")
	}
	log.Println("connected to " + ordName)
	ordDB = mdb

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, indexerName)
	mdb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
				Colorful: false,
			},
		),
	})
	if err != nil {
		panic(err)
	}
	if mdb == nil {
		panic("failed to connect database")
	}
	log.Println("connected to " + indexerName)
	db = mdb

	return
}

func GetDB() *gorm.DB {
	return db
}

func GetOrdDB() *gorm.DB {
	return ordDB
}
