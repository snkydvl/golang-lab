package main

import (
//	"fmt"
	"os"
	"os/exec"
	
	"github.com/snkydvl/vehicle-lab/transport"
    "github.com/snkydvl/vehicle-lab/car"
    "github.com/snkydvl/vehicle-lab/sportscar"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	clearScreen()

	transport1 := transport.NewTransport("CargoMaster", 2022, "Scania", 150000, 60)
	transport1.PrintInfo()
	transport1.PrintResale()

	car1 := car.NewCar("FamilyCar", 2023, "Toyota", 35000, 120, "Camry", 4)
	car1.PrintInfo()
	car1.PrintResale()

	sportsCar1 := sportscar.NewSportsCar("SuperFast", 2024, "Ferrari", 350000, 180, "F8 Tributo", 2, 720, 340)
	sportsCar1.PrintInfo()
	sportsCar1.PrintResale()

	/*fmt.Println("\nПОЛИМОРФИЗМ")
	fmt.Println("Вызов методов через интерфейс Vehicle:\n")

	vehicles := []transport.Vehicle{transport1, car1, sportsCar1}

	for i, v := range vehicles {
		fmt.Printf("--- Транспорт %d ---\n", i+1)
		v.PrintInfo()
		v.PrintResale()
	} */
}
