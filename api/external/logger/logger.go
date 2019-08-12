package logger

import (
	"log"
	"io"
)

// NewLogger new log manager
func NewLogger(w io.Writer) *logger {
	log.SetOutput(w)
	return &logger{}
}

const (
	LogPrefixError = "[Error]"

	LogPrefixInfo = "[Info]"

	LogPrefixDebug = "[Debug]"
)

type logger struct{}

func (l *logger) Errorf(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixError)
	log.Printf(format, args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixInfo)
	log.Printf(format, args...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixDebug)
	log.Printf(format, args...)
}
