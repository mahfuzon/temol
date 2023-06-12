package libraries

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func NewLogger() *logrus.Logger {
	log := logrus.New()
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customFormatter)
	dateNow := time.Now().Format("2006-01-02")
	file, err := os.OpenFile("logs/"+dateNow+"_log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(errors.New("failed create log file"))
	}

	log.SetOutput(file)

	return log
}
