package service

import (
	"fmt"
	"sync"

	"github.com/priykumar/MachineCoding/Logger/src/model"
)

var mu sync.Mutex

func LogToFile(ts model.TsFormat, level string, content interface{}) {
	currTime := getCurrTime(ts)
	fmt.Println("FILE: ", currTime, level, content)
}

func SyncLogger(content interface{}, Ts model.TsFormat, level string, sType model.SinkType) {
	switch sType {
	case model.S3:
		LogToS3(Ts, level, content)
	case model.FILE:
		LogToFile(Ts, level, content)
	}
}

func AsynLogger() {
	for {
		val := <-model.AsynChannel
		switch val.Mode {
		case model.S3:
			LogToS3(val.Ts, val.Level, val.Content)
		case model.FILE:
			LogToFile(val.Ts, val.Level, val.Content)
		}
	}
}

func ConcurrentLogger(input model.S) {
	mu.Lock()
	defer mu.Unlock()

	switch input.Mode {
	case model.S3:
		LogToS3(input.Ts, input.Level, input.Content)
	case model.FILE:
		LogToFile(input.Ts, input.Level, input.Content)
	}
}
