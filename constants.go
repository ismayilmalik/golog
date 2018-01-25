package golog

// Log level
type LogLevel int

const (
	TRACE LogLevel = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

var logLevelNames = []string {
	"TRACE",
	"DEBUG",
	"INFO",
	"WARNING",
	"ERROR",
	"FATAL",
}

func (l LogLevel) String() string {
	return logLevelNames[l]
}

