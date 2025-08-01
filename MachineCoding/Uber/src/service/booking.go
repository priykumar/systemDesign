package service

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/priykumar/MachineCoding/Uber/src/model"
)

var Booking_Id int
var Id_Booking map[int]model.Booking

type DriverRespDetail struct {
	Status string
	CabId  int
}

func NewBooking(customerId int, customerLoc model.Location, vehicleId int, drop model.Location) model.Booking {
	otp := "0000"
	// opt := generateOtp()

	return model.Booking{
		Id:         0,
		CustomerId: customerId,
		CabId:      vehicleId,
		OTP:        otp,
		PickUp:     customerLoc,
		Drop:       drop,
	}
}

func bookingIdGenerator() int {
	Booking_Id += 1
	return Booking_Id
}

func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%04d", rand.Intn(10000))
}

func CalculateDistance(srclat, srclong, dstlat, dstlong float64) float64 {
	x_square, y_square := (srclat-dstlat)*(srclat-dstlat), (srclong-dstlong)*(srclong-dstlong)
	return math.Sqrt(x_square + y_square)
}

func ShowCabAvailabiltyAndFare(srclat, srclong, dstlat, dstlong float64) map[model.VehicleType]float64 {
	vType_Fare := map[model.VehicleType]float64{}
	for vType, cabs := range IsAvailable_Type_Cabs[true] {
		for _, c := range cabs {
			if CalculateDistance(srclat, srclong, c.CurrLocation.Lat, c.CurrLocation.Long) <= 5 {
				vType_Fare[vType] = 0
			}
		}
	}

	src_dst_distance := CalculateDistance(srclat, srclong, dstlat, dstlong)
	for vType := range vType_Fare {
		vType_Fare[vType] = CalculateFare(vType, src_dst_distance)
	}

	return vType_Fare
}

func MakeBooking(cId int, srclat, srclong, dstlat, dstlong, fare float64, vType model.VehicleType) string {
	var Id int

	cabResp := make(chan DriverRespDetail)
	go SearchCab(vType, srclat, srclong, cabResp)

	acceptedCabId := -1

	select {
	case resp := <-cabResp:
		if resp.Status == "ACCEPTED" {
			Id = bookingIdGenerator()
			acceptedCabId = resp.CabId
			break
		}
	case <-time.After(time.Minute):
		fmt.Println("Timed out searching for a cab")
		break
	}

	newBooking := model.Booking{
		Id:         Id,
		CustomerId: cId,
		CabId:      acceptedCabId,
		PickUp:     model.Location{Lat: srclat, Long: srclong},
		Drop:       model.Location{Lat: dstlat, Long: dstlong},
		Fare:       fare,
		OTP:        generateOTP(),
		Status:     model.BOOKED,
	}

	Id_Booking[Id] = newBooking

	// Move cabs from Available to NotAvailable
	cabDetails := IsAvailable_Type_Cabs[true][vType][acceptedCabId]
	delete(IsAvailable_Type_Cabs[true][vType], acceptedCabId)
	IsAvailable_Type_Cabs[false][vType][acceptedCabId] = cabDetails

	return newBooking.OTP
}

func TransaitionBookingStatus(Id int, status model.BookingStatus) {
	booking := Id_Booking[Id]
	booking.Status = status

	Id_Booking[Id] = booking
}
