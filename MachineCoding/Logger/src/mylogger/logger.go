package mylogger

import (
	"fmt"

	"github.com/priykumar/MachineCoding/Logger/src/model"
)

type MyLogger interface {
	Debug(interface{})
	Info(interface{})
	Warn(interface{})
	Error(interface{})
	Fatal(interface{})
}

func NewLogger(loggerType string, sinkList []model.Sink) MyLogger {
	switch loggerType {
	case "SYNC":
		return SyncLogger{sinkList}
	case "ASYNC":
		return AsyncLogger{sinkList}
	case "CONCURRENT":
		return ConcurrentLogger{sinkList}
	default:
		fmt.Println("Provide a valid sync type. Supported [SYNC, ASYNC, CONCURRENT]")
	}

	return nil
}
