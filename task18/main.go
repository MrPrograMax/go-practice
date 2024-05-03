package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	num int
	mx  sync.Mutex
}

// NewCounter - конструктор
func NewCounter() *Counter {
	return &Counter{
		num: 0,
	}
}

func (c *Counter) IncreaseCounter() {
	c.mx.Lock()         // Блокируем мьютекс
	defer c.mx.Unlock() //Разблокируем мьютекс

	c.num++ // Увеличиваем значение счётчика на 1
}

func (c *Counter) GetCounterValue() int {
	return c.num
}

func Increment(counter *Counter, wg *sync.WaitGroup) {
	defer wg.Done()

	counter.IncreaseCounter() //Вызываем увеличение счётчика
}

func main() {
	// Инициализируем счетчик
	counter := NewCounter()
	// Задаем количество горутин
	goCount := 10

	var wg sync.WaitGroup

	for i := 0; i < goCount; i++ {
		wg.Add(1)
		//Вызываем функцию увеличение счётчика в отдельной горутине
		go Increment(counter, &wg)
	}

	// Ждем завершения всех горутин
	wg.Wait()

	// Выводим результат
	fmt.Println(counter.GetCounterValue())
}
