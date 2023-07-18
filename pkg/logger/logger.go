package logger

import (
	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger() (*logrus.Logger, error) {
	var log = logrus.New()
	// log.Formatter = new(logrus.JSONFormatter)
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true}) //default
	// log.Level = logrus.TraceLevel
	// log.Out = os.Stdout
	log.Out = &lumberjack.Logger{
		// Filename:   "/var/log/template2/app.log",
		Filename:   "./logs/app.log",
		MaxSize:    200, // megabytes
		MaxBackups: 0,
		MaxAge:     0,    // days
		Compress:   true, // disabled by default
	}

	// file, err := os.OpenFile("logs/logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	// if err == nil {
	// 	log.Out = file
	// } else {
	// 	log.Info("Failed to log to file, using default stderr")
	// 	return nil, err
	// }

	return log, nil
}
