package main

import (
	"log"
	"fmt"
)

const (
	LevelError   = iota
	LevelWarning
	LevelInfo
	LevelDebug
)

type Logger struct {
	level   int
	handler *log.Logger
}

func (logger *Logger) Error(format string, v ...interface{}) {

	if LevelError > logger.level {
		return
	}

	msg := fmt.Sprintf("[E] "+format, v...)

	logger.handler.Printf(msg)
}

func main() {

}
