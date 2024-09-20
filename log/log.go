package log

import (
	"github.com/dawsonalex/ms-macrod/build"
	"github.com/sirupsen/logrus"
)

func New() *Logger {
	logger := &Logger{
		Logger: logrus.New(),
	}
	return logger
}

type Logger struct {
	*logrus.Logger
}

func (l *Logger) WithBuildInfo() *Entry {
	b := build.Info()
	return l.WithFields(Fields{
		"version":     b.Version.Sprint(),
		"commit":      b.Commit,
		"branch":      b.Branch,
		"host":        b.Host,
		"environment": b.Environment,
	})
}
