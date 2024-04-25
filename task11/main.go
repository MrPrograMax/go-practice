package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств
*/

func main() {
	//Инициализируем множества
	arr1 := []int{2, 4, 6, 8, 10}
	arr2 := []int{1, 3, 5}

	//Вызываем функцию
	fmt.Println(GetIntersection(arr1, arr2))
}

func GetIntersection(arr1, arr2 []int) []int {
	//Инициализруем словарь для хранения букв из первого массива
	values := make(map[int]bool)

	//Инициализируем срез, в который будем
	//пересекающиеся значения
	result := make([]int, 0)

	//Проходимся по первому массиву и сохраняем в словарь
	for _, value := range arr1 {
		values[value] = true
	}

	//Смотрим присутвие элемента из массива 2 в словаре
	for _, value := range arr2 {
		if _, status := values[value]; status {
			//Если в словаре значение найдено, то
			//Оно является пересекающимся с массивом 1
			result = append(result, value)
		}
	}

	return result
}
