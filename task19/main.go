package main

import "fmt"

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func ReverseString(str string) string {
	//Граничные случаи когда переворот не требуется
	if len(str) == 0 || len(str) == 1 {
		return str
	}

	//Строки являются неизменяемым типом!
	//Для удобства конвертируем строку в массив рун
	letters := []rune(str)

	//Вводим два индекса, для указания какие элементы нужно поменять местами
	i := 0
	j := len(letters) - 1

	for i < j {
		//Свап элементов
		letters[i], letters[j] = letters[j], letters[i]
		//Смещение индексов
		i++
		j--
	}

	//Собираем массив рун в строку
	return string(letters)
}
func main() {
	CheckFunc("главрыба", 1)
	CheckFunc("aabb", 2)
	CheckFunc("12345", 3)
}

func CheckFunc(str string, num int) {
	fmt.Println("\nВариант", num)
	fmt.Println("Before reverse:", str)
	str = ReverseString(str)
	fmt.Println("After reverse:", str)
}
