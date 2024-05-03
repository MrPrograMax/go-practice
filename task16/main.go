package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Реализовать quicksort.
*/

// Быстрая сортировка с использованием 3 индексов
func quickSort(arr []int, l, r int) {
	if l < r {
		//Разделяем массив на две части и вызываем сортировку для каждой
		E, G := partition(arr, l, r)
		quickSort(arr, l, E)
		quickSort(arr, G, r)
	}
}

func partition(arr []int, l, r int) (int, int) {
	//Выбираем опорный элемент
	pivot := arr[(l+r)/2]
	E := l // Указатель на первый элемент равный X
	G := r // Указатель на первый элемент больший X
	N := l // Указатель на текущий элемент. Итератор

	for N <= G {
		// У нас бывает только 3 случая
		// Элемент < X
		// Элемент = X
		// Элемент > X
		if arr[N] < pivot {
			arr[N], arr[E] = arr[E], arr[N]
			E += 1
			N += 1
		} else if arr[N] > pivot {
			arr[N], arr[G] = arr[G], arr[N]
			G -= 1
		} else {
			N += 1
		}
	}

	return E, G + 1
}

func main() {
	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Генерируем массива из 10 случайных чисел
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.Intn(100) // Генерируем случайное число от 0 до 99
	}

	// Выводим исходный массив
	fmt.Printf("arr:        %v\n", arr)
	// Сортируем
	quickSort(arr, 0, len(arr)-1)
	// Выводим результат
	fmt.Printf("sorted arr: %v\n", arr)
}
