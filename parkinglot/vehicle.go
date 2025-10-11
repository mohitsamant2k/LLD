package parkinglot

type Vehicle interface {
	GetVehicleNumber() string
	GetVehicleSize() VehicleSize
}

type VehicleSize int

const (
	Small VehicleSize = iota
	Medium
	Large
)

type BaseVehicle struct {
	vehicleNumber string
	vehicleSize   VehicleSize
}

func (bv *BaseVehicle) GetVehicleNumber() string {
	return bv.vehicleNumber
}

func (bv *BaseVehicle) GetVehicleSize() VehicleSize {
	return bv.vehicleSize
}

type Car struct {
	*BaseVehicle
	numDoors int
}

// Car-specific method
func (c *Car) GetNumDoors() int {
	return c.numDoors
}

// Constructor
func GetCar(number string, doors int) Vehicle {
	return &Car{
		BaseVehicle: &BaseVehicle{
			vehicleNumber: number,
			vehicleSize:   Medium,
		},
		numDoors: doors,
	}
}

type Bike struct {
	*BaseVehicle
	numDoors int
}

// Car-specific method
func (c *Bike) GetNumDoors() int {
	return c.numDoors
}

// Constructor
func GetBike(number string, doors int) Vehicle {
	return &Bike{
		BaseVehicle: &BaseVehicle{
			vehicleNumber: number,
			vehicleSize:   Small,
		},
		numDoors: doors,
	}
}

type Truck struct {
	*BaseVehicle
	numDoors int
}

// Car-specific method
func (c *Truck) GetNumDoors() int {
	return c.numDoors
}

// Constructor
func GetTruck(number string, doors int) Vehicle {
	return &Truck{
		BaseVehicle: &BaseVehicle{
			vehicleNumber: number,
			vehicleSize:   Large,
		},
		numDoors: doors,
	}
}

// 3️⃣ Key differences summarized
// Feature	Value embedding (BaseVehicle)	Pointer embedding (*BaseVehicle)
// Memory	Copies the struct inside	Holds a pointer, no copy
// Method calls	Only value receiver methods directly	Both value and pointer receivers work
// Field mutation	Changes affect only the copy	Changes affect the original
// Initialization	Easy, no pointer needed	Must initialize pointer (&BaseVehicle)
