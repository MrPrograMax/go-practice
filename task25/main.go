package main

import (
	"fmt"
	"time"
)

// sleep использует функцию time.After
// Она нужна для создания канала, который отправляет сигнал по истечении заданного времени.
func sleep(dur time.Duration) {
	<-time.After(dur)
}

func main() {
	fmt.Println("How long did it take?")
	sleep(5 * time.Second)
	fmt.Println("5 seconds...")

}
