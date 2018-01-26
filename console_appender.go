package golog

import (
	"os"
	"fmt"
	"sync"
	"errors"
)

type ConsoleAppender struct {
	mutex sync.Mutex
	ready bool
	disposed bool
}

func (c *ConsoleAppender) ID() string {
	return "ConsoleAppender"
}

func (c *ConsoleAppender) Struct() error {
	if c.ready {
		return nil
	}

	c.mutex.Lock()
	c.ready = true
	c.mutex.Unlock()

	return nil
}

func (c *ConsoleAppender) Write(buffer []byte) (int, error) {
	if !c.ready {
		return -1, errors.New(fmt.Sprintf("%s not ready.", c.ID()))
	}

	if c.disposed {
		return -1, errors.New(fmt.Sprintf("%s is disposed.", c.ID()))
	}

	return os.Stdout.Write(buffer)
}

func (c *ConsoleAppender) Dispose() error {
	if c.disposed {
		return nil
	}

	c.mutex.Lock()
	c.disposed = true
	c.mutex.Unlock()

	return nil
}