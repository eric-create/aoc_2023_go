package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)

	for i := 100; i <= 200; i++ {
		go myfunction(i, channel)
	}

	for i := 100; i <= 200; i++ {
		fmt.Println(<-channel)
	}
}

func myfunction(count int, channel chan int) {
	time.Sleep(2 * time.Second)
	channel <- count
}
