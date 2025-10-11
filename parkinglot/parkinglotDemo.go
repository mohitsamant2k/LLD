package parkinglot

import (
	"fmt"
	"time"
)

func RunDemo() {
	park := GetSpotFromStart{}
	fee := SimpleFee{
		perSecondfee: 2,
	}

	instance, _ := GetInstance(fee, park)

	f1 := &ParkingFloor{parkingspots: []*ParkingSpot{
		{spotID: "F1-S1", vehicleSize: Small},
		{spotID: "F1-S2", vehicleSize: Medium},
		{spotID: "F1-S3", vehicleSize: Large},
	}}
	f2 := &ParkingFloor{parkingspots: []*ParkingSpot{
		{spotID: "F2-S1", vehicleSize: Small},
		{spotID: "F2-S2", vehicleSize: Small},
		{spotID: "F2-S3", vehicleSize: Medium},
		{spotID: "F2-S4", vehicleSize: Large},
	}}
	f3 := &ParkingFloor{parkingspots: []*ParkingSpot{
		{spotID: "F3-S1", vehicleSize: Medium},
		{spotID: "F3-S2", vehicleSize: Medium},
		{spotID: "F3-S3", vehicleSize: Large},
	}}

	instance.AddFloor(f1)
	instance.AddFloor(f2)
	instance.AddFloor(f3)

	v1 := GetCar("c1", 4)
	v2 := GetBike("b1", 4)
	v3 := GetTruck("t1", 4)
	v4 := GetCar("c2", 4)
	v5 := GetBike("b2", 4)
	v6 := GetTruck("t2", 4)
	v7 := GetTruck("t3", 4)
	v8 := GetTruck("t4", 4)

	t1 := instance.ParkVehicle(v1)
	t2 := instance.ParkVehicle(v2)
	t3 := instance.ParkVehicle(v3)
	t4 := instance.ParkVehicle(v4)
	t5 := instance.ParkVehicle(v5)
	t6 := instance.ParkVehicle(v6)
	t7 := instance.ParkVehicle(v7)
	t8 := instance.ParkVehicle(v8)

	fmt.Printf("%+v", t1)
	fmt.Println(t2)
	fmt.Println(t3)
	fmt.Println(t4)
	fmt.Println(t5)
	fmt.Println(t6)
	fmt.Println(t7)
	fmt.Println(t8)

	time.Sleep(5 * time.Second)
	x1, y1 := instance.ExitVehicle(t1)
	fmt.Printf("%v , %+v", x1, y1)

	x7, y7 := instance.ExitVehicle(t7)
	fmt.Printf("%v , %+v", x7, y7)

	t8 = instance.ParkVehicle(v8)
	fmt.Println(t8)

}
