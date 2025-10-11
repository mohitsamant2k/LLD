package parkinglot

import (
	"fmt"
	"sync"
	"time"
)

type ParkingLot struct {
	floor           []*ParkingFloor
	activeTicket    map[string]ParkingTicket
	feestrategy     FeeStrategy
	parkingstrategy ParkingStrategy
}

func (pt *ParkingLot) AddFloor(f *ParkingFloor) {
	pt.floor = append(pt.floor, f)
}

func (pt *ParkingLot) ParkVehicle(v Vehicle) *ParkingTicket {
	spot, _ := pt.parkingstrategy.GetTheSpot(pt.floor, v)
	if spot == nil {
		fmt.Println("no spots for vehicle")
		return nil
	}
	spot.currentVehicle = v
	return &ParkingTicket{
		vehicleNumber: v.GetVehicleNumber(),
		startTime:     time.Now(),
		ticketId:      v.GetVehicleNumber() + time.Now().Format(time.RFC3339),
		parkingSpot:   spot,
	}
}

func (pt *ParkingLot) ExitVehicle(p *ParkingTicket) (int, *ParkingTicket) {
	fee, ticket := pt.feestrategy.GetPrice(p)
	p.parkingSpot.currentVehicle = nil
	return fee, ticket
}

var (
	instance *ParkingLot
	mu       sync.Mutex
)

func GetInstance(f FeeStrategy, p ParkingStrategy) (*ParkingLot, error) {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &ParkingLot{
			floor:           []*ParkingFloor{},
			activeTicket:    make(map[string]ParkingTicket),
			feestrategy:     f,
			parkingstrategy: p,
		}
	}
	return instance, nil
}
