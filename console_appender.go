package golog

import (
	"os"
	"fmt"
	"sync"
	"errors"
)

type consoleAppender struct {
	lock sync.Mutex
	ready bool
	disposed bool
}

func (c *consoleAppender) ID() string {
	return "ConsoleAppender"
}

func (c *consoleAppender) Construct() error {
	if c.ready {
		return nil
	}

	c.lock.Lock()
	c.ready = true
	c.lock.Unlock()

	return nil
}

func (c *consoleAppender) Write(buffer []byte) (int, error) {
	if !c.ready {
		return -1, errors.New(fmt.Sprintf("%s not ready.", c.ID()))
	}

	if c.disposed {
		return -1, errors.New(fmt.Sprintf("%s is disposed.", c.ID()))
	}

	return os.Stdout.Write(buffer)
}

func (c *consoleAppender) Dispose() error {
	if c.disposed {
		return nil
	}

	c.lock.Lock()
	c.disposed = true
	c.lock.Unlock()

	return nil
}

func NewConsoleAppender() *consoleAppender {
	appender := &consoleAppender{}
	appender.Construct()
	return appender
}