package parkinglot

type ParkingStrategy interface{
	GetTheSpot([]*ParkingFloor, Vehicle) (*ParkingSpot, error)
}

type GetSpotFromStart struct {

}
func (gs GetSpotFromStart)GetTheSpot(pf []*ParkingFloor, v Vehicle)(*ParkingSpot, error){
	if pf == nil {
		return nil,nil;
	}

	if v == nil {
		return nil, nil
	}

	for i:=0;i<len(pf);i++ {
		if pf[i].parkingspots == nil {
			continue
		}

		for j:=0;j<len(pf[i].parkingspots);j++ {
			if(pf[i].parkingspots[j].IsAvailable() && pf[i].parkingspots[j].vehicleSize == v.GetVehicleSize()){
				return pf[i].parkingspots[j] , nil
			}
		}
	}
	return nil,nil	
}

type GetSpotFromEnd struct {

}
func (gs *GetSpotFromEnd)GetTheSpot(pf []*ParkingFloor, v Vehicle)(*ParkingSpot, error){
	if pf == nil {
		return nil,nil;
	}
	for i:=len(pf)-1;i>=0;i-- {
		if pf[i].parkingspots == nil {
			continue
		}

		for j:=0;j<len(pf[i].parkingspots);j++ {
			if(pf[i].parkingspots[j].IsAvailable()){
				return pf[i].parkingspots[j] , nil
			}
		}
	}
	return nil,nil	
}