package parkinglot
import(

)
type ParkingSpot struct {
	spotID        string
	currentVehicle Vehicle
	vehicleSize VehicleSize
}

func (ps *ParkingSpot) IsAvailable() bool {
	return ps.currentVehicle == nil
}

