package main

import (
	"fmt"
)

/*
Разработать конвейер чисел.
Даны два канала:
в первый пишутся числа (x) из массива,
во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	//Создаем два канала
	c0 := make(chan int)
	c1 := make(chan int)

	//Массив
	arr := []int{1, 2, 4, 6, 8, 10}

	//Запускаем горутины
	go source(arr, c0)
	go degreeMagnifier(c0, c1)
	//Ждем выполнения перед выходом из программы
	//Если сделать go printer(c1) - горутины не успеют отработать
	//Можно использовать WaitGroup, если нужно
	printer(c1)
}

func source(arr []int, down chan int) {
	for _, item := range arr {
		down <- item
	}
	//Закрываем канал, т.к данных больше нет
	close(down)
}

func degreeMagnifier(up, down chan int) {
	//Читаем значения из канала, пока тот не будет закрыт
	//Благодаря range *chan int*
	for item := range up {
		down <- item * item
	}
	//Канал из которого, происходило чтение закрылся
	//Данных больше не будет, закрываем канал
	close(down)
}

func printer(up chan int) {
	//Читаем канал, пока тот не будет закрыт
	for item := range up {
		fmt.Print(item, " ")
	}
	fmt.Println()
}
