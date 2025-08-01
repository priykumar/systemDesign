package mylogger

import (
	"github.com/priykumar/MachineCoding/Logger/src/model"
)

type AsyncLogger struct {
	SinkList []model.Sink
}

func (s AsyncLogger) Debug(content interface{}) {
	for _, sink := range s.SinkList {
		if sink.L_Level <= model.DEBUG {
			model.AsynChannel <- model.S{Content: content, Ts: sink.T_Type, Level: "DEBUG", Mode: sink.S_Type}
		}
	}
}

func (s AsyncLogger) Info(content interface{}) {
	for _, sink := range s.SinkList {
		if sink.L_Level <= model.INFO {
			model.AsynChannel <- model.S{Content: content, Ts: sink.T_Type, Level: "INFO", Mode: sink.S_Type}
		}
	}
}

func (s AsyncLogger) Warn(content interface{}) {
	for _, sink := range s.SinkList {
		if sink.L_Level <= model.WARN {
			model.AsynChannel <- model.S{Content: content, Ts: sink.T_Type, Level: "WARN", Mode: sink.S_Type}
		}
	}
}

func (s AsyncLogger) Error(content interface{}) {
	for _, sink := range s.SinkList {
		if sink.L_Level <= model.ERROR {
			model.AsynChannel <- model.S{Content: content, Ts: sink.T_Type, Level: "ERROR", Mode: sink.S_Type}
		}
	}
}

func (s AsyncLogger) Fatal(content interface{}) {
	for _, sink := range s.SinkList {
		if sink.L_Level <= model.FATAL {
			model.AsynChannel <- model.S{Content: content, Ts: sink.T_Type, Level: "FATAL", Mode: sink.S_Type}
		}
	}
}
