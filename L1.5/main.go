package main

import (
	"fmt"
	"time"
)

func sender(ch chan int) {
	for i := range 10 {
		ch <- i * i
		ticker := time.NewTicker(time.Second)
		<-ticker.C
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go sender(ch)

	timeout := time.After(5 * time.Second)

	for {
		select {
		case val := <-ch:
			fmt.Println("Получено:", val)
		case <-timeout:
			fmt.Println("Время вышло")
			return
		}
	}
}
