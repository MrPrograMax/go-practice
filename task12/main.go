package main

import (
	"fmt"
	mySet "wb-tech-l1/task12/set"
)

func main() {
	//Инициализация
	set := mySet.NewSet()
	// исходный слайс
	stringList := []string{"cat", "cat", "dog", "cat", "tree"}
	//Конвертирование массива в set и вывод
	set.AddSlice(stringList)
	fmt.Println(set.GetItems())
}
