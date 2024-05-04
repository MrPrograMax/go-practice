package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
	"wb-tech-l1/task7/mymap"
)

/*
Реализовать конкурентную запись данных в map.
*/

func producer(ctx context.Context, id int, myMap *mymap.MyMap, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("Producer [%d] started its work\n", id)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Producer [%d] has completed its work\n", id)
			return
		default:
			value := rand.Intn(1000)
			myMap.SetValue(fmt.Sprintf("key_%d", id), value)
			time.Sleep(time.Millisecond * 500)
		}
	}

}

func main() {
	// Создаем словарь
	myMap := mymap.NewMyMap()

	// Создаем контекст и синхронизацию
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	// Запускаем производителей
	prodCount := 5
	wg.Add(prodCount)
	for i := 1; i <= prodCount; i++ {
		go producer(ctx, i, myMap, &wg)
	}

	time.Sleep(time.Second * 1)
	seconds := 3
	for i := seconds; i >= 0; i-- {
		fmt.Printf("\rTimer: %d", i)
		time.Sleep(time.Second * 1)
	}
	fmt.Println()

	//Завершаем контекст
	cancel()
	//Ожидаем завершения горутин
	wg.Wait()

	// Выводим словарь
	myMap.PrintMyMap()
}
