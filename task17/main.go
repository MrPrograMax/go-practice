package main

import "fmt"

func BinarySearch(arr []int, target, left, right int) int {
	//По условию для бин. поиска элементы должны быть отсортированы
	//Следовательно если target меньше чем левая граница
	//или больше чем правая граница
	//То элемента, точно нет в нашем массиве
	if target < arr[left] || arr[right] < target {
		return -1
	}

	mid := (left + right) / 2
	if arr[mid] == target {
		return mid
	}
	if arr[mid] > target {
		return BinarySearch(arr, target, left, mid-1)
	}

	if right < left {
		return -1
	}

	return BinarySearch(arr, target, mid+1, right)
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 5
	left, right := 0, len(arr)-1

	fmt.Println("Index of target:", BinarySearch(arr, target, left, right))
}
