package parkinglot

import(
	
)

type ParkingFloor struct{
	parkingspots []*ParkingSpot
}

func(pf *ParkingFloor) GetAvailableSpots(v Vehicle) []ParkingSpot {
	ret := []ParkingSpot{}
	for i:=0; i< len(pf.parkingspots);i++ {
		if (pf.parkingspots[i].IsAvailable() && v.GetVehicleSize() == pf.parkingspots[i].vehicleSize) {
			ret = append(ret ,*pf.parkingspots[i])
		}
	}
	return ret 
}