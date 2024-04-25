package shipFabric

import (
	"fmt"
	"wb-tech-l1/task21/carFabric"
)

// ShipFactory фабрика, которая производит корабли
// Она не умеет создавать машины на прямую
type ShipFactory struct {
}

// NewShipFactory конструктор для ShipFactory
func NewShipFactory() *ShipFactory {
	return &ShipFactory{}
}

// ConvertToCar метод, который "конвертирует" корабль в машину
func (s *ShipFactory) ConvertToCar() {}

// ShipAdapter адаптер, который использует
type ShipAdapter struct {
	shipFactory *ShipFactory
}

// NewShipAdapter конструктор для ShipAdapter
func NewShipAdapter() carFabric.Creator {
	return &ShipAdapter{}
}

// CreateCar реализация метода интерфейса для ShipFactory, который не умеет напрямую создавать машины, через ShipAdapter
func (adapter *ShipAdapter) CreateCar() {
	adapter.shipFactory.ConvertToCar()
	fmt.Println("I created a car! (adapted)")
}
