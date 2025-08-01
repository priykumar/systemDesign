package main

import (
	"fmt"
	"time"

	"github.com/priykumar/MachineCoding/Logger/src/model"
	"github.com/priykumar/MachineCoding/Logger/src/mylogger"
	"github.com/priykumar/MachineCoding/Logger/src/service.go"
)

func main() {
	file_sink_info := model.NewSink(model.FILE, model.INFO, model.DD_MM_YYYY)
	s3_sink_info := model.NewSink(model.S3, model.INFO, model.DD_MM_YYYY)
	info_logger_sink := []model.Sink{file_sink_info, s3_sink_info}

	file_sink_fatal := model.NewSink(model.FILE, model.FATAL, model.YYYY_MM_DD_HH_MM_SS)
	s3_sink_fatal := model.NewSink(model.S3, model.FATAL, model.DD_MM_YYYY)
	fatal_logger_sink := []model.Sink{file_sink_fatal, s3_sink_fatal}

	done := make(chan bool)
	go service.AsynLogger()

	var syncMyLogger, asyncMyLogger, concurrentMyLogger mylogger.MyLogger
	cases := [][]model.Sink{info_logger_sink, fatal_logger_sink}

	for i := 0; i < len(cases); i++ {
		syncMyLogger = mylogger.NewLogger("SYNC", cases[i])
		asyncMyLogger = mylogger.NewLogger("ASYNC", cases[i])
		concurrentMyLogger = mylogger.NewLogger("CONCURRENT", cases[i])

		fmt.Println("------------------- SYNC -------------------")
		syncMyLogger.Debug(fmt.Sprintf("Sync Case%d", i))
		syncMyLogger.Info(fmt.Sprintf("Sync Case%d", i))
		syncMyLogger.Warn(fmt.Sprintf("Sync Case%d", i))
		syncMyLogger.Error(fmt.Sprintf("Sync Case%d", i))
		syncMyLogger.Fatal(fmt.Sprintf("Sync Case%d", i))

		fmt.Println("------------------- ASYNC -------------------")
		for j := 0; j < 8; j++ {
			asyncMyLogger.Debug(fmt.Sprintf("Async Case%d", j))
			asyncMyLogger.Info(fmt.Sprintf("Async Case%d", j))
			asyncMyLogger.Warn(fmt.Sprintf("Async Case%d", j))
			asyncMyLogger.Error(fmt.Sprintf("Async Case%d", j))
			asyncMyLogger.Fatal(fmt.Sprintf("Async Case%d", j))
		}
		time.Sleep(5 * time.Second)

		fmt.Println(concurrentMyLogger)
		fmt.Println("------------------- CONCURRENT -------------------")
		for j := 0; j < 8; j++ {
			concurrentMyLogger.Debug(fmt.Sprintf("Concurrent Case%d", j))
			concurrentMyLogger.Info(fmt.Sprintf("Concurrent Case%d", j))
			concurrentMyLogger.Warn(fmt.Sprintf("Concurrent Case%d", j))
			concurrentMyLogger.Error(fmt.Sprintf("Concurrent Case%d", j))
			concurrentMyLogger.Fatal(fmt.Sprintf("Concurrent Case%d", j))
		}
		time.Sleep(5 * time.Second)
		fmt.Println()
		fmt.Println()
	}

	<-done
}
