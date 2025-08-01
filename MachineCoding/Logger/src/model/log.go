package model

type AsyncLogger struct {
	SinkList []Sink
}

type ConcurrentLogger struct {
	SinkList []Sink
}
