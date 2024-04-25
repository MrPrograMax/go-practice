package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет,
что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.
*/

func main() {
	//Тест 1
	str := "abcd"
	fmt.Println(IsAllLettersUnique(str))

	//Тест 2
	str = "abCdefAaf"
	fmt.Println(IsAllLettersUnique(str))

	//Тест 3
	str = "aabcd"
	fmt.Println(IsAllLettersUnique(str))
}

func IsAllLettersUnique(str string) bool {
	//Т.к функция регистронезависима перевод все в нижний регистр
	str = strings.ToLower(str)

	//Инициализируем словарь букв
	var letters = make(map[rune]bool)

	//Проходимся по каждой букве в строке
	for _, letter := range str {
		//Проверяем присутствие буквы в словаре, если да, то буква уже встречалась ранее
		if _, status := letters[letter]; status {
			return false
		} else {
			//В противном случае добавляем букву в словарь
			letters[letter] = true
		}
	}

	//Если дошли до сюда значит все буквы уникальны
	return true
}
