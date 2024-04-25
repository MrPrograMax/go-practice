package main

import (
	"errors"
	"fmt"
)

func RemoveAt(arr []int, index int) ([]int, error) {
	//Проверка на выход индекса за приделы массива
	if index < 0 || len(arr) <= index {
		return arr, errors.New("index out of range")
	}

	//объеденяем части до и после элемента
	arr = append(arr[:index], arr[index+1:]...)
	return arr, nil
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	index := 12
	if arr, err := RemoveAt(arr, index); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(arr)
	}
}
