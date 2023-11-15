package log

import (
	"log"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Create a new instance of the logger. You can have any number of instances.
var logger = logrus.New()

const DEV = "dev"

func Init(env, path string) {
	log.Printf("init log, env: %s, path: %s", env, path)
	logger.SetReportCaller(true)

	logger.SetFormatter(&logrus.TextFormatter{
		QuoteEmptyFields: true,
	})

	if env == DEV {
		logger.SetLevel(logrus.DebugLevel)
		return
	}

	logger.SetLevel(logrus.InfoLevel)
	file, err := os.OpenFile(filepath.Join(path, "node.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Warn("Failed to log to file, using default stderr")
	}

	logrus.SetOutput(&lumberjack.Logger{
		Filename:   "brc-rpc.log",
		MaxSize:    100,
		MaxAge:     30,
		MaxBackups: 30,
		LocalTime:  true,
		Compress:   false,
	})

}

func Logger() *logrus.Logger {
	return logger
}
