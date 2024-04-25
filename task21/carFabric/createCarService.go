package carFabric

import "fmt"

// Creator интерефейс для создания машины
type Creator interface {
	CreateCar()
}

// CarFactory Фабрика, которая производит машины
type CarFactory struct{}

// NewFactory конструктор для CarFactory
func NewFactory() *CarFactory {
	return &CarFactory{}
}

// CreateCar реализация метода интерфейса Creator для создания машин
func (f *CarFactory) CreateCar() {
	fmt.Println("I created a car!")
}
