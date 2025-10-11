package parkinglot
import (
	"time"
)
type ParkingTicket struct {
	vehicleNumber string
	startTime time.Time
	endTime time.Time
	ticketId string
	parkingSpot *ParkingSpot
}

