package mylogger

import (
	"github.com/priykumar/MachineCoding/Logger/src/model"
	"github.com/priykumar/MachineCoding/Logger/src/service.go"
)

type SyncLogger struct {
	SinkList []model.Sink
}

func (s SyncLogger) Debug(content interface{}) {
	for _, sink := range s.SinkList {
		if sink.L_Level <= model.DEBUG {
			service.SyncLogger(content, sink.T_Type, "DEBUG", sink.S_Type)
		}
	}
}
func (s SyncLogger) Info(content interface{}) {

	for _, sink := range s.SinkList {
		if sink.L_Level <= model.INFO {
			service.SyncLogger(content, sink.T_Type, "INFO", sink.S_Type)
		}
	}
}
func (s SyncLogger) Warn(content interface{}) {

	for _, sink := range s.SinkList {
		if sink.L_Level <= model.WARN {
			service.SyncLogger(content, sink.T_Type, "WARN", sink.S_Type)
		}
	}
}
func (s SyncLogger) Error(content interface{}) {

	for _, sink := range s.SinkList {
		if sink.L_Level <= model.ERROR {
			service.SyncLogger(content, sink.T_Type, "ERROR", sink.S_Type)
		}
	}
}
func (s SyncLogger) Fatal(content interface{}) {

	for _, sink := range s.SinkList {
		if sink.L_Level <= model.FATAL {
			service.SyncLogger(content, sink.T_Type, "FATAL", sink.S_Type)
		}
	}
}
