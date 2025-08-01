package model

type BookingStatus int

const (
	BOOKED BookingStatus = iota
	INPROGRESS
	COMPLETED
	CANCELLED
)

type Booking struct {
	Id         int
	CustomerId int
	CabId      int
	OTP        string
	PickUp     Location
	Drop       Location
	Fare       float64
	Status     BookingStatus
}
