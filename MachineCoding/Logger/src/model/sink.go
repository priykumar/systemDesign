package model

type Sink struct {
	S_Type  SinkType
	L_Level LogLevel
	T_Type  TsFormat
}

func NewSink(sType SinkType, lLevel LogLevel, tType TsFormat) Sink {
	return Sink{
		S_Type:  sType,
		L_Level: lLevel,
		T_Type:  tType,
	}
}
