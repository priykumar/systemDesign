package service

import (
	"github.com/priykumar/MachineCoding/Uber/src/model"
)

var Id_Rider map[int]model.Rider

func NewRider(name, phone string, lat, long float64) {
	Id := 0
	newRider := model.Rider{
		Id:    Id,
		Name:  name,
		Phone: phone,
		CurrLocation: model.Location{
			Lat:  lat,
			Long: long,
		},
	}

	Id_Rider[Id] = newRider
}

func GetRider(id int) model.Rider {
	return model.Rider{}
}

func InitiaseRider() {
	Id_Rider = make(map[int]model.Rider)
}
