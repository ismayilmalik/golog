package golog

import (
	"bytes"
	"time"
)

type LogRecord struct {
	Time    time.Time
	Level   LogLevel
	Message string
}

func (lr *LogRecord) ToBuffer() *bytes.Buffer {
	buffer := &bytes.Buffer{}

	buffer.WriteString(lr.Time.Format(RFC822))
	buffer.WriteString(" : ")
	buffer.WriteString(lr.Level.String())
	buffer.WriteString(" : ")
	buffer.WriteString(lr.Message)

	return buffer
}
