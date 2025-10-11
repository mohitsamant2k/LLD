package parkinglot

import (
	"time"
	"math"
)

type FeeStrategy interface{
	GetPrice(*ParkingTicket) (int, *ParkingTicket) 
}

type SimpleFee struct{
	perSecondfee int 
}

func (sf SimpleFee)GetPrice(p *ParkingTicket)(int, *ParkingTicket){
	if p == nil{
		return 0,p
	}

	currentTime := time.Now();
	duration := currentTime.Sub(p.startTime)
	p.endTime = currentTime
	numberOfSecond := int(math.Ceil(duration.Seconds()));
	return numberOfSecond*sf.perSecondfee, p
}

