package service

import (
	"fmt"
	"time"

	"github.com/priykumar/MachineCoding/Uber/src/model"
)

var Cab_Id int

// isAvailable -> VehicleType -> CabId -> Cab
var IsAvailable_Type_Cabs map[bool]map[model.VehicleType]map[int]model.Cab

func cabIdGenerator() int {
	Cab_Id += 1
	return Cab_Id
}

func NewCab(number string, location model.Location, cabType model.VehicleType, lat, long float64) {
	id := cabIdGenerator()
	newCab := model.Cab{
		Id:          id,
		CabNumber:   number,
		IsAvailable: false,
		CurrLocation: model.Location{
			Lat:  lat,
			Long: long,
		},
	}

	IsAvailable_Type_Cabs[false][cabType][id] = newCab
}

func InitiaseCab() {
	IsAvailable_Type_Cabs = make(map[bool]map[model.VehicleType]map[int]model.Cab)
}

func RequestCab(id int, cabResp chan DriverRespDetail, done chan bool) {
	if id%7 == 0 {
		cabResp <- DriverRespDetail{"ACCEPTED", id}
		done <- true

	} else if id%2 == 0 {
		cabResp <- DriverRespDetail{"REJECTED", id}
	}
}

func SearchCab(vType model.VehicleType, srclat, srclong float64, cabResp chan DriverRespDetail) {
	// send Request to Cabs
	count := 0
	batchSize := 5

	done := make(chan bool, 10)
	for id, c := range IsAvailable_Type_Cabs[true][vType] {
		if CalculateDistance(srclat, srclong, c.CurrLocation.Lat, c.CurrLocation.Long) <= 5 {
			go RequestCab(id, cabResp, done)
			count++

			if count == batchSize {
			forloop:
				for {
					select {
					case val := <-done:
						if val {
							fmt.Println("Cab driver has accespted the ride")
							return
						}
					case <-time.After(10 * time.Second):
						fmt.Println("Time for new set of batch")
						count = 0
						break forloop
					}
				}
			}
		}
	}
}
