package util

import (
	"github.com/sirupsen/logrus"
)



var log = logrus.New()

func GetLogger() *logrus.Logger {
	return log
}