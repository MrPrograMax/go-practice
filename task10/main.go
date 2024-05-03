package main

import (
	"fmt"
	"sort"
)

/*
Дана последовательность температурных колебаний:
-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc

*/

func main() {
	//Массив проверяемых элементов
	items := []float64{-25.4, -27.0, 0.321, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	//Словарь для хранения данных с шагом 10
	dic := make(map[int][]float64)
	for _, item := range items {
		key := (int(item) / 10) * 10
		if _, ok := dic[key]; ok {
			dic[key] = append(dic[key], item)
		} else {
			dic[key] = []float64{item}
		}
	}

	//Обычно значения не приходят к нам в рандомном порядке
	//Отсортируем ключи из словаря и выведем значения
	keys := make([]int, 0, len(dic))

	for key := range dic {
		keys = append(keys, key)
	}

	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("%4d: ", k)
		for _, item := range dic[k] {
			fmt.Printf("%7.1f ", item)
		}
		fmt.Print("\n")

	}
}
