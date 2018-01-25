package golog

import (
	"log"
	"sync"
)

type Logger interface {
	Trace(format string, args ...interface{})
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
}

type defaultLogger struct{
	lock sync.Mutex
	appender LogAppender
	logger *log.Logger
}

func (l *defaultLogger) Trace(format string, args ...interface{})  {
	l.log(TRACE, format, args...)
}

func (l *defaultLogger) Debug(format string, args ...interface{})  {
	l.log(DEBUG, format, args...)
}

func (l *defaultLogger) Info(format string, args ...interface{})  {
	l.log(INFO, format, args...)
}

func (l *defaultLogger) Warn(format string, args ...interface{})  {
	l.log(WARNING, format, args...)
}

func (l *defaultLogger) Error(format string, args ...interface{})  {
	l.log(ERROR, format, args...)
}

func (l *defaultLogger) Fatal(format string, args ...interface{})  {
	l.log(FATAL, format, args...)
}

func (l *defaultLogger) log(level LogLevel, format string, args ...interface{})  {
	
}


