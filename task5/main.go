package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Считываем ввод пользователя с консоли
	var countSecond int
	fmt.Print("Count of seconds: ")
	_, err := fmt.Scan(&countSecond)
	if err != nil {
		fmt.Printf("input error: %v", err)
		return
	}

	// Используем контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup // Синхронизации для групп горутин
	ch := make(chan int)  // Создаем канал для записи и чтения

	wg.Add(1)
	go producer(ctx, ch, &wg)

	wg.Add(1)
	go consumer(ctx, ch, &wg)

	// Запускаем таймер для выключения программы
	time.Sleep(time.Second * time.Duration(countSecond))

	fmt.Println("Trying to shutdown...")
	cancel()  // Вызов cancel закрывает канал, возвращаемый методом контекста Done()
	close(ch) // Закрываем канал
	wg.Wait() // Ожидает завершения всех горутин
	fmt.Println("Status: 200 (OK). Everything was closed successfully!")
}

// Функция которая записывает в канал случайное значение
func producer(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик горутин
	for {
		// select позволяет ждать нескольких операций на канале.
		select {
		case <-ctx.Done(): // Закрытие канала при завершении работы
			fmt.Println("Producer has completed its work!")
			return
		default:
			ch <- rand.Intn(100)
			time.Sleep(time.Second * 1)
		}
	}
}

// Функция, которая считывает данные из канала и выводит
func consumer(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик горутин
	for {
		// select позволяет ждать нескольких операций на канале.
		select {
		case <-ctx.Done(): // Закрытие канала при завершении работы
			fmt.Printf("Consumer has completed its work!\n")
			return
		case item, ok := <-ch: // считывает значение и проверяет статус канала
			if !ok { //Если канал закрыт
				return
			}
			fmt.Printf("Consumer value: %d\n", item)
		}
	}
}
