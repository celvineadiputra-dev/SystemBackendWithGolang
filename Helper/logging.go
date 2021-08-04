package Helper

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type CreateLogging struct {
	Message  string
	FileName string
	TypeLog  string
}

type CreateLoggingOption func(*CreateLogging)

func NewCreateLogging(message string, fileName string, typeLog string) bool {
	var log = logrus.New()

	log.Out = os.Stdout
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 06666)

	if err == nil {
		log.Out = file
	} else {
		log.Info(err.Error())
	}

	info := "DATE : [" + time.Now().Format("01-02-2006 15:04:05 Monday") + "], Message : [" + message + "]"

	if typeLog == "Info" {
		log.Info(info)
	} else if typeLog == "Error" {
		log.Error(info)
	}

	return true
}
