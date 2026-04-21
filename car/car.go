package car

import (
	"fmt"
	"github.com/snkydvl/vehicle-lab/transport"
)

type Car struct {
	*transport.Transport
	Model string
	Doors int
}

func NewCar(name string, year int, brand string, price int, speed float64, model string, doors int) *Car {
	return &Car{
		Transport: transport.NewTransport(name, year, brand, price, speed),
		Model:     model,
		Doors:     doors,
	}
}

func (c *Car) PrintInfo() {
	fmt.Printf("Car Model: %s\n", c.Model)
	fmt.Printf("Year: %d\n", c.Year)
	fmt.Printf("Brand: %s\n", c.Brand)
	fmt.Printf("Price: %d $\n", c.Price)
	fmt.Printf("Speed: %.1f km/h\n", c.Speed)
	fmt.Printf("Fuel: %.1f %%\n", c.Fuelcap)
	fmt.Printf("Doors: %d\n", c.Doors)
	fmt.Printf("\n")
}

func (c *Car) Resale() int {
	return int(float64(c.Price) * (1.0 - (0.1 * float64(2026-c.Year))))
}

func (c *Car) PrintResale() {
	fmt.Printf("Resale price: %d $\n", c.Resale())
	fmt.Printf("\n")
}
