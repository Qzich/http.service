package http_service

import (
	"log"
)

func NewLogger(debug bool) *stderrLogger {
	return &stderrLogger{isDebug: debug}
}

type stderrLogger struct {
	isDebug bool
}

func (*stderrLogger) Info(infos ... interface{}) {
	logMessage("info", infos...)
}

func (stderrLogger *stderrLogger) Debug(warns ...interface{}) {
	if stderrLogger.isDebug {
		logMessage("debug", warns...)
	}
}

func (*stderrLogger) Error(errors ...interface{}) {
	logMessage("error", errors...)
}

func logMessage(prefix string, messages ...interface{}) {
	for _, message := range messages {
		log.Println("["+prefix+"] ", message)
	}
}
