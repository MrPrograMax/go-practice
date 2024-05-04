package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

// Использование канала
func workerWay1Chanel(ch <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker(Chanel) started its work...")

	for {
		select {
		case <-ch:
			fmt.Println("Worker(Chanel) has completed its work!")
			fmt.Println()
			return
		default:
			fmt.Println("Worker(Chanel) doing useful work...")
			time.Sleep(time.Second * 1)
		}
	}
}

// Использования контекста
func workerWay2Context(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker(Context) started its work...")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker(Context) has completed its work!")
			fmt.Println()
			return
		default:
			fmt.Println("Worker(Context) doing useful work...")
			time.Sleep(time.Second * 1)
		}
	}
}

// Использование  флага
func workerWay3Flag(flag *bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker(Flag) started its work!")

	for {
		if !*flag {
			fmt.Println("Worker(Flag) has completed its work!")
			fmt.Println()
			return
		} else {
			fmt.Println("Worker(Flag) doing useful work...")
			time.Sleep(time.Second * 1)
		}
	}
}

// Использование завершения внутри самой горутины
func workerWay4SelfCompletion(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker(SelfCompletion) started its work!")

	for i := 0; i < 3; i++ {
		fmt.Println("Worker(SelfCompletion) doing useful work...")
		time.Sleep(time.Second * 1)
	}

	fmt.Println("Worker(SelfCompletion) has completed its work!")
	fmt.Println()

	return
}

func main() {
	//Используем WaitGroup для синхронизации группы горутин
	wg := sync.WaitGroup{}

	// Way1 Chanel
	log.Println("Way 1. Checking the work using channel...")
	ch := make(chan bool) //Создаем канал

	wg.Add(1)                    //Увеличиваем счетчик горутин
	go workerWay1Chanel(ch, &wg) //Запускаем горутину

	time.Sleep(time.Second * 3) //Ждем некоторое время
	close(ch)                   //Закрываем канал
	wg.Wait()

	// Way2 Context
	log.Println("Way 2. Checking the work using context...")
	ctx, cancel := context.WithCancel(context.Background()) // Создаем контекст

	wg.Add(1)
	go workerWay2Context(ctx, &wg)

	time.Sleep(time.Second * 3)
	cancel() //Закрываем контекст
	wg.Wait()

	// Way3 Flag
	log.Println("Way 3. Checking the work using flag...")
	flag := true

	wg.Add(1)
	go workerWay3Flag(&flag, &wg)

	time.Sleep(time.Second * 3)
	flag = false //Переключаем флан для завершения
	wg.Wait()

	// Way4 Self Completion
	log.Println("Way 4. Checking the work using self completion...")

	wg.Add(1)
	go workerWay4SelfCompletion(&wg)
	wg.Wait()
}
