package sportscar

import (
	"fmt"
    "github.com/snkydvl/golang-lab/car")

type SportsCar struct {
	*car.Car
	HorsePower int
	TopSpeed   float64
}

func NewSportsCar(name string, year int, brand string, price int, speed float64, model string, doors int, horsePower int, topSpeed float64) *SportsCar {
	return &SportsCar{
		Car:        car.NewCar(name, year, brand, price, speed, model, doors),
		HorsePower: horsePower,
		TopSpeed:   topSpeed,
	}
}

func (s *SportsCar) PrintInfo() {
	fmt.Printf("SportsCar Model: %s\n", s.Model)
	fmt.Printf("Year: %d\n", s.Year)
	fmt.Printf("Brand: %s\n", s.Brand)
	fmt.Printf("Price: %d $\n", s.Price)
	fmt.Printf("Speed: %.1f km/h\n", s.Speed)
	fmt.Printf("Fuel: %.1f %%\n", s.Fuelcap)
	fmt.Printf("Doors: %d\n", s.Doors)
	fmt.Printf("HorsePower: %d hp\n", s.HorsePower)
	fmt.Printf("TopSpeed: %.1f km/h\n", s.TopSpeed)
	fmt.Printf("\n")
}

func (s *SportsCar) Resale() int {
	return int(float64(s.Price) * (1.0 - (0.15 * float64(2026-s.Year))))
}

func (s *SportsCar) PrintResale() {
	fmt.Printf("Resale price: %d $\n", s.Resale())
	fmt.Printf("\n")
}
