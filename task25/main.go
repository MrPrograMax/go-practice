package main

import (
	"fmt"
	"time"
)

func sleep(dur time.Duration) {
	<-time.After(dur)
}

func main() {
	fmt.Println("How long did it take?")
	sleep(5 * time.Second)
	fmt.Println("5 seconds...")

}
