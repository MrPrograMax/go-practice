package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров,
которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/

func main() {
	// Считываем ввод пользователя с консоли
	var workersCount int
	fmt.Print("Count of workers: ")
	_, err := fmt.Scan(&workersCount)
	if err != nil {
		fmt.Printf("input error: %v", err)
		return
	}

	// Используем контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup              // Синхронизации для групп горутин
	ch := make(chan int, workersCount) // Создаем канал для записи и чтения

	wg.Add(1)
	go producer(ctx, ch, &wg)

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go consumer(ctx, ch, &wg, i)
	}

	// Создаём канал для обработки сигнала к завершению работы программы
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

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
func consumer(ctx context.Context, ch chan int, wg *sync.WaitGroup, id int) {
	defer wg.Done() // Уменьшаем счетчик горутин
	for {
		// select позволяет ждать нескольких операций на канале.
		select {
		case <-ctx.Done(): // Закрытие канала при завершении работы
			fmt.Printf("Worker [%d] has completed its work!\n", id)
			return
		case item, ok := <-ch: // считывает значение и проверяет статус канала
			if !ok { //Если канал закрыт
				return
			}
			fmt.Printf("Worker: [%d] value: %d\n", id, item)
		}
	}
}
