package main

import (
	"fmt"
	"sync"
	"time"
)

// Функция sleep останавливает выполнение горутины на указанное время
// Использует таймер с использованием канала "C"
func sleep(duration time.Duration) {
	timer := time.NewTimer(duration) 
	<-timer.C
}

func main() {
	wg := sync.WaitGroup{}

	wg.Add(10)
	for i := range 10 {
		go func(i int) {
			defer wg.Done()
			sleep(3 * time.Second)
			fmt.Println("Горутина:", i, "завершена!")
		}(i)
	}

	fmt.Println("Ожидание завершения горутин...")
	wg.Wait()
	fmt.Println("Завершено!")
}
