package model

type SinkType int
type LogLevel int
type TsFormat int

const (
	STDOUT SinkType = iota
	FILE
	S3
)

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

const (
	DD_MM_YYYY_HH_MM_SS TsFormat = iota
	YYYY_MM_DD_HH_MM_SS
	DD_MM_YYYY
	HH_MM_SS
)

type S struct {
	Content interface{}
	Ts      TsFormat
	Level   string
	Mode    SinkType
}

var AsynChannel = make(chan S, 5)
