package transport

import "fmt"

type Vehicle interface {
	PrintInfo()
	Resale() int
	PrintResale()
}

type Transport struct {
	Name    string
	Year    int
	Brand   string
	Price   int
	Speed   float64
	Fuelcap float64
}

func NewTransport(name string, year int, brand string, price int, speed float64) *Transport {
	return &Transport{
		Name:    name,
		Year:    year,
		Brand:   brand,
		Price:   price,
		Speed:   speed,
		Fuelcap: 100.0,
	}
}

func (t *Transport) PrintInfo() {
	fmt.Printf("Transport Model: %s\n", t.Name)
	fmt.Printf("Year: %d\n", t.Year)
	fmt.Printf("Brand: %s\n", t.Brand)
	fmt.Printf("Price: %d $\n", t.Price)
	fmt.Printf("Speed: %.1f km/h\n", t.Speed)
	fmt.Printf("Fuel: %.1f %%\n", t.Fuelcap)
	fmt.Printf("\n")
}

func (t *Transport) Resale() int {
	return int(float64(t.Price) * (1.0 - (0.1 * float64(2026-t.Year))))
}

func (t *Transport) PrintResale() {
	fmt.Printf("Resale price: %d $\n", t.Resale())
	fmt.Printf("\n")
}
