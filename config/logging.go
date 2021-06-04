package config

import (
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogging() {
	Log = logrus.New()
	//Format
	Log.SetFormatter(&logrus.TextFormatter{})
	//Log Level
	Log.SetLevel(logrus.Level(ApiConfig.LogLevel))
}
