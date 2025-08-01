package service

import "github.com/priykumar/MachineCoding/Uber/src/model"

var vehicleType_baseFare = map[model.VehicleType]float64{
	model.AUTO:  20,
	model.SEDAN: 27,
	model.SUV:   35,
}

func CalculateFare(vType model.VehicleType, dist float64) float64 {
	return vehicleType_baseFare[vType] * dist
}
