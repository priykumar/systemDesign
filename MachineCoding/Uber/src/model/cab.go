package model

type VehicleType int

const (
	AUTO VehicleType = iota
	SEDAN
	SUV
)

type Cab struct {
	Id           int
	CabNumber    string
	CabType      VehicleType
	IsAvailable  bool
	CurrLocation Location
}
