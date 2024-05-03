package main

import "fmt"

/*
Дана переменная int64.
Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func SetBit(number, bitNumber, bitValue int64) int64 {
	// Устанавливаем значение в i-й бит
	if bitValue == 0 {
		number &^= 1 << bitNumber
	} else {
		number |= 1 << bitNumber
	}
	return number
}

func main() {
	var number int64    //Исходное число
	var bitNumber int64 //Номер бита
	var bitValue int64  //Значение бита

	//Считываем ввод пользователя number
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Printf("input error: %v", err)
		return
	}

	//Считываем ввод пользователя bitNumber
	fmt.Print("Enter the bit number: ")
	_, err = fmt.Scan(&bitNumber)
	if err != nil {
		fmt.Printf("input error: %v", err)
		return
	}

	//Считываем ввод пользователя bitValue
	fmt.Print("Enter a value for the bit (0 or 1): ")
	_, err = fmt.Scan(&bitValue)
	if err != nil {
		fmt.Printf("input error: %v", err)
		return
	}

	if bitValue != 0 && bitValue != 1 {
		fmt.Println("input error! Use only 0 or 1")
		return
	}

	fmt.Println()
	fmt.Printf("Number:     %d [ %b ]\n", number, number)
	number = SetBit(number, bitNumber, bitValue)
	fmt.Printf("New Number: %d [ %b ]\n", number, number)
}
