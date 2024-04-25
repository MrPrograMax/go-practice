package main

import "fmt"

func main() {
	a, b := 10, 15
	ClassicSwap(a, b)
	PlusMinusSwap(a, b)
	MultiplicationDivisionSwap(a, b)
}

func ClassicSwap(a, b int) {
	fmt.Println("Classic swap:")
	fmt.Printf("Before swap: a = %d, b = %d\n", a, b)
	a, b = b, a
	fmt.Printf("After swap: a = %d, b = %d\n\n", a, b)
}

func PlusMinusSwap(a, b int) {
	fmt.Println("Plus and minus swap:")
	fmt.Printf("Before swap: a = %d, b = %d\n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("After swap: a = %d, b = %d\n\n", a, b)
}

func MultiplicationDivisionSwap(a, b int) {
	fmt.Println("Multiplication and division swap:")
	fmt.Printf("Before swap: a = %d, b = %d\n", a, b)
	a = a * b
	b = a / b
	a = a / b
	fmt.Printf("After swap: a = %d, b = %d\n\n", a, b)
}
