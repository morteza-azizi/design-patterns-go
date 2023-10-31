package prototype

import "fmt"

type Car interface {
	Clone() Car
	Info() string
}

type ConcreteCar struct {
	Brand      string
	Model      string
	Year       int
	Price      float64
	Colour     string
	IsElectric bool
}

func (c *ConcreteCar) Clone() Car {
	return &ConcreteCar{
		Brand:      c.Brand,
		Model:      c.Model,
		Year:       c.Year,
		Price:      c.Price,
		Colour:     c.Colour,
		IsElectric: c.IsElectric,
	}
}

func (c *ConcreteCar) Info() string {
	return fmt.Sprintf("Car: %s %s, Year: %d, Price: $%.2f, Colour: %s, Electric: %v",
		c.Brand, c.Model, c.Year, c.Price, c.Colour, c.IsElectric)
}
