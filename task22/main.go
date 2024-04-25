package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает
две числовых переменных a,b, значение которых > 2^20.
*/

// Addition - сложение
func Addition(a *big.Int, b *big.Int) *big.Int {
	result := big.NewInt(0)
	return result.Add(a, b)
}

// Subtraction - вычитание
func Subtraction(a *big.Int, b *big.Int) *big.Int {
	result := big.NewInt(0)
	return result.Sub(a, b)
}

// Multiplication - умножение
func Multiplication(a *big.Int, b *big.Int) *big.Int {
	result := big.NewInt(0)
	return result.Mul(a, b)
}

// Division - деление
func Division(a *big.Int, b *big.Int) *big.Int {
	result := big.NewInt(0)
	return result.Div(a, b)
}

func main() {
	a := big.NewInt(0)
	b := big.NewInt(0)

	a.SetString("5000000000000000000000000000000000000", 10)
	b.SetString("1000000000000000000000000000000000000", 10)

	fmt.Println(Addition(a, b))
	fmt.Println(Subtraction(a, b))
	fmt.Println(Multiplication(a, b))
	fmt.Println(Division(a, b))
}
