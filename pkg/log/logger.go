package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

func New() *logrus.Logger {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetFormatter(&logrus.JSONFormatter{})
	return l
}
