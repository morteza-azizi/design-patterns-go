package builder

type Builder interface {
	SetWheels() Builder
	SetSeats() Builder
	SetStructure() Builder
	GetVehicle() VehicleProduct
}

type Director struct {
	builder Builder
}

func (director *Director) Construct() {
	director.
		builder.
		SetSeats().
		SetStructure().
		SetWheels()
}

func (director *Director) SetBuilder(builder Builder) {
	director.builder = builder
}

// Product
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// Car Concrete Builder
type CarBuilder struct {
	vehicle VehicleProduct
}

func (car *CarBuilder) SetWheels() Builder {
	car.vehicle.Wheels = 4
	return car
}

func (car *CarBuilder) SetSeats() Builder {
	car.vehicle.Seats = 5
	return car
}

func (car *CarBuilder) SetStructure() Builder {
	car.vehicle.Structure = "Car"
	return car
}

func (car *CarBuilder) GetVehicle() VehicleProduct {
	return car.vehicle
}

// Bus Concrete Builder
type BusBuilder struct {
	vehicle VehicleProduct
}

func (bus *BusBuilder) SetWheels() Builder {
	bus.vehicle.Wheels = 8
	return bus
}

func (bus *BusBuilder) SetSeats() Builder {
	bus.vehicle.Seats = 30
	return bus
}

func (bus *BusBuilder) SetStructure() Builder {
	bus.vehicle.Structure = "Bus"
	return bus
}

func (bus *BusBuilder) GetVehicle() VehicleProduct {
	return bus.vehicle
}
