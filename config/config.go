package config

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

var (
	configFile string
)

func Init() {
	flag.StringVar(&configFile, "config", "config/brc-rpc-dev.yaml", "path of config file")
	flag.Parse()
	if configFile == "" {
		flag.Usage()
	}

	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		errStr := fmt.Sprintf("viper read config is failed, err is %v configFile is %s ", err, configFile)
		panic(errStr)
	}
}
