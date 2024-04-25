package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

// ReserveWords метод для переворота слов в строке
func ReserveWords(str string) (string, error) {
	// Проверяем не пришла ли нам на вход пустая строка
	if len(str) == 0 {
		return str, errors.New("string is empty")
	} else if len(str) == 1 { // Граничный слуай когда у нас всего 1 элемент и переворачивать нечего
		return str, nil
	}

	//Разбиваем строку на массив слов
	//в качестве разделителя используются пробел, а так же знаки пунктуации
	words := strings.FieldsFunc(str, func(r rune) bool {
		return unicode.IsSpace(r) || unicode.IsPunct(r)
	})

	//Вводим два индекса, для указания какие элементы нужно поменять местами
	i := 0
	j := len(words) - 1

	for i < j {
		//Свап элементов
		words[i], words[j] = words[j], words[i]
		//Смещение индексов
		i++
		j--
	}

	//Собираем массив слов в единую строку
	return strings.Join(words, " "), nil
}

func main() {
	CheckFunc("snow dog sun", 1)
	CheckFunc("1, 2, 3, 4, 5", 2)
	CheckFunc("белый; синий; красный", 3)
}

func CheckFunc(str string, num int) {
	fmt.Println("\nВариант", num)
	fmt.Println("Before reverse:", str)

	if str, err := ReserveWords(str); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("After reverse:", str)
	}
}
