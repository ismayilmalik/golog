package golog

type LogAppender interface {
	ID() string
	Write([]byte) (int, error)
	Construct() error
	Dispose() error
}