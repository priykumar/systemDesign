package service

import (
	"time"

	"github.com/priykumar/MachineCoding/Logger/src/model"
)

func getCurrTime(ts model.TsFormat) string {
	now := time.Now()
	switch ts {
	case model.DD_MM_YYYY:
		return now.Format("02-01-2006")
	case model.DD_MM_YYYY_HH_MM_SS:
		return now.Format("02-01-2006:15-04-05")
	case model.YYYY_MM_DD_HH_MM_SS:
		return now.Format("2006-01-02:15-04-05")
	case model.HH_MM_SS:
		return now.Format("15-04-05")
	}
	return now.Format("15-04-05")
}
