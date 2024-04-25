package main

import (
	"fmt"
	"wb-tech-l1/task21/carFabric"
	"wb-tech-l1/task21/shipFabric"
)

/*
Реализовать паттерн «адаптер» на любом примере.
*/

func main() {
	carFactory := carFabric.NewFactory()
	shipFactory := shipFabric.NewShipAdapter()

	fmt.Print("Car Factory: ")
	carFactory.CreateCar()
	fmt.Print("Ship Factory: ")
	shipFactory.CreateCar()

}
