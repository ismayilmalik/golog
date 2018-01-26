package golog

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Logger interface {
	Trace(format string, args ...interface{})
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})

	SetLevel(LogLevel)
	SetAppender(LogAppender)
}

type defaultLogger struct {
	lock     sync.Mutex
	logLevel LogLevel
	appender LogAppender
	logger   *log.Logger
}

var lg Logger

func DefaultLogger() Logger {
	return lg
}

func (l *defaultLogger) Trace(format string, args ...interface{}) {
	l.log(TRACE, format, args...)
}

func (l *defaultLogger) Debug(format string, args ...interface{}) {
	l.log(DEBUG, format, args...)
}

func (l *defaultLogger) Info(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}

func (l *defaultLogger) Warn(format string, args ...interface{}) {
	l.log(WARNING, format, args...)
}

func (l *defaultLogger) Error(format string, args ...interface{}) {
	l.log(ERROR, format, args...)
}

func (l *defaultLogger) Fatal(format string, args ...interface{}) {
	l.log(FATAL, format, args...)
}

func (l *defaultLogger) SetLevel(level LogLevel) {
	l.lock.Lock()
	l.logLevel = level
	l.lock.Unlock()
}

func (l *defaultLogger) SetAppender(appender LogAppender) {
	l.lock.Lock()
	l.appender = appender
	l.logger = constructLogger(appender)
	l.lock.Unlock()
}

func (l *defaultLogger) log(level LogLevel, format string, args ...interface{}) {
	if level < l.logLevel {
		return
	}

	record := &LogRecord{
		Level:   l.logLevel,
		Time:    time.Now(),
		Message: fmt.Sprintf(format, args...),
	}

	go func() {
		switch level {
		case TRACE, DEBUG, INFO, WARNING, ERROR:
			l.logger.Print(record.ToBuffer())
		case FATAL:
			l.logger.Fatal(record.ToBuffer())
		}
	}()
}

func init() {
	a := NewConsoleAppender()
	lg = &defaultLogger{}
	lg.SetLevel(ALL)
	lg.SetAppender(a)
}

func constructLogger(appender LogAppender) *log.Logger {
	return log.New(appender, "-", 0)
}
