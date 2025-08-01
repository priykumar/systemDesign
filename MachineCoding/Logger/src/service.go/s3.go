package service

import (
	"fmt"

	"github.com/priykumar/MachineCoding/Logger/src/model"
)

func LogToS3(ts model.TsFormat, level string, content interface{}) {
	currTime := getCurrTime(ts)
	fmt.Println("  S3: ", currTime, level, content)
}
