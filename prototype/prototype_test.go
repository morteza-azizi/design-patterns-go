package prototype

import (
	"testing"
)

func TestCarCloning(t *testing.T) {

	carPrototype := &ConcreteCar{
		Brand:      "Toyota",
		Model:      "Camry",
		Year:       2023,
		Price:      25000.0,
		Colour:     "Black",
		IsElectric: false,
	}

	carClone := carPrototype.Clone().(*ConcreteCar)

	carClone.Brand = "Honda"
	carClone.Price = 28000.0

	if carClone == carPrototype {
		t.Error("Clone is not a separate instance")
	}

	if carClone.Brand != "Honda" {
		t.Errorf("Expected brand 'Honda', got %s", carClone.Brand)
	}
	if carClone.Price != 28000.0 {
		t.Errorf("Expected price $28000.00, got $%.2f", carClone.Price)
	}

	if carPrototype.Brand != "Toyota" {
		t.Errorf("Expected brand 'Toyota', got %s", carPrototype.Brand)
	}
	if carPrototype.Price != 25000.0 {
		t.Errorf("Expected price $25000.00, got $%.2f", carPrototype.Price)
	}
	if carPrototype.Colour != "Black" {
		t.Errorf("Expected colour 'Black', got %s", carPrototype.Brand)
	}
}

func TestCarInfo(t *testing.T) {
	car := &ConcreteCar{
		Brand:      "Ford",
		Model:      "F-150",
		Year:       2023,
		Price:      35000.0,
		Colour:     "Black",
		IsElectric: false,
	}

	expectedInfo := "Car: Ford F-150, Year: 2023, Price: $35000.00, Colour: Black, Electric: false"

	if info := car.Info(); info != expectedInfo {
		t.Errorf("Expected info: %s, got: %s", expectedInfo, info)
	}
}

func TestCarCloneNil(t *testing.T) {
	var car Car
	if car == nil {
		t.Skip("Skipping test for nil car")
	}

	clone := car.Clone()

	if clone != nil {
		t.Errorf("Expected cloned car to be nil, got %v", clone)
	}
}

func TestCarCloneType(t *testing.T) {
	carPrototype := &ConcreteCar{
		Brand:      "Toyota",
		Model:      "Camry",
		Year:       2023,
		Price:      25000.0,
		IsElectric: false,
	}

	clone := carPrototype.Clone()

	_, isConcreteCar := clone.(*ConcreteCar)

	if !isConcreteCar {
		t.Error("Expected cloned object to be of type *ConcreteCar")
	}
}
