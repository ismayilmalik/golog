package main

import (
	"time"
	"github.com/ismayilmalik/golog"
)

func main() {
	logger := golog.DefaultLogger()
	
	logger.Info("Default appender is ConsoleAppender")
	logger.Trace("This is trace log.")
	logger.Debug("This is debug log.")
	logger.Error("This is error log.")
	logger.Info("This is info log.")
	logger.Warn("This is warning log.")

	time.Sleep(time.Second * 5)
}
