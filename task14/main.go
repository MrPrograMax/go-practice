package main

import "fmt"

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

func GetType(value interface{}) {
	switch value.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("unknown")
	}

}

func main() {
	var value interface{}
	value = 10
	GetType(value)

	value = 10.5
	GetType(value)

	value = "hello"
	GetType(value)

	value = true
	GetType(value)

	value = nil
	GetType(value)
}
