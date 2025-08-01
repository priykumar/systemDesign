package mylogger

import (
	"fmt"

	"github.com/priykumar/MachineCoding/Logger/src/model"
	"github.com/priykumar/MachineCoding/Logger/src/service.go"
)

type ConcurrentLogger struct {
	SinkList []model.Sink
}

func (s ConcurrentLogger) Log() {
	fmt.Println("Logging using CONCURRENT Logger")
}

func (s ConcurrentLogger) Debug(content interface{}) {
	for _, sink := range s.SinkList {
		if sink.L_Level <= model.DEBUG {
			go service.ConcurrentLogger(model.S{Content: content, Ts: sink.T_Type, Level: "DEBUG", Mode: sink.S_Type})
		}
	}
}

func (s ConcurrentLogger) Info(content interface{}) {
	for _, sink := range s.SinkList {
		if sink.L_Level <= model.INFO {
			go service.ConcurrentLogger(model.S{Content: content, Ts: sink.T_Type, Level: "INFO", Mode: sink.S_Type})
		}
	}
}

func (s ConcurrentLogger) Warn(content interface{}) {
	for _, sink := range s.SinkList {
		if sink.L_Level <= model.WARN {
			go service.ConcurrentLogger(model.S{Content: content, Ts: sink.T_Type, Level: "WARN", Mode: sink.S_Type})
		}
	}
}

func (s ConcurrentLogger) Error(content interface{}) {
	for _, sink := range s.SinkList {
		if sink.L_Level <= model.ERROR {
			go service.ConcurrentLogger(model.S{Content: content, Ts: sink.T_Type, Level: "ERROR", Mode: sink.S_Type})
		}
	}
}

func (s ConcurrentLogger) Fatal(content interface{}) {
	for _, sink := range s.SinkList {
		if sink.L_Level <= model.FATAL {
			go service.ConcurrentLogger(model.S{Content: content, Ts: sink.T_Type, Level: "FATAL", Mode: sink.S_Type})
		}
	}
}
