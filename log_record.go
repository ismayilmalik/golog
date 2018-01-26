package golog

import (
	"time"
)

type LogRecord struct {
	Time time.Time
	Level LogLevel
	Message string
}