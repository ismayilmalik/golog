package golog

type LogAppender interface {
	ID() string
	Write([]byte) (int, error)
	Struct() error
	Dispose()
}