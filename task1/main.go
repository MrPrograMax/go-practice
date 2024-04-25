package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Задание 1
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action
от родительской структуры Human (аналог наследования).
*/

func main() {
	human := &Human{name: "Maxim", age: 21}
	actions := &Action{Human: *human}

	actions.sayName()
	actions.sayAge()
	actions.introduceYourself()
}

type Human struct {
	name string
	age  int
}

//region Human methods

// sayName  выводит имя человека
func (h *Human) sayName() {
	fmt.Printf("Hello, my name is %s\n", h.name)
}

// sayAge выводит возраст человека
func (h *Human) sayAge() {
	fmt.Printf("I'm %d!\n", h.age)
}

// doSomething возвращает случайное число действия
func (h *Human) doSomething() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

//endregion

type Action struct {
	Human //Встраиваем методы за счет внедерния структуры Human
}

// introduceYourself выводит имя, возраст и действие человека
func (a *Action) introduceYourself() {
	fmt.Printf("I'm %s and I'm %d, I wanna be the number: %d\n", a.name, a.age, a.doSomething())
}
