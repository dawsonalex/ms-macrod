package log

import "github.com/sirupsen/logrus"

func (l *Logger) Infof(format string, args ...interface{}) *Logger {
	l.Logger.Infof(format, args...)
	return l
}
func (l *Logger) Errorf(format string, args ...interface{}) *Logger {
	l.Logger.Errorf(format, args...)
	return l
}

func (l *Logger) Debugf(format string, args ...interface{}) *Logger {
	l.Logger.Debugf(format, args...)
	return l
}

func (l *Logger) Warnf(format string, args ...interface{}) *Logger {
	l.Logger.Warnf(format, args...)
	return l
}

func (l *Logger) WithFields(fields Fields) *Entry {
	entry := l.Logger.WithFields(logrus.Fields(fields))
	return &Entry{
		Entry: *entry,
	}
}
