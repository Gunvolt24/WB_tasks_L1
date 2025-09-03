//go:build ignore

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Способ выхода из горутины №6: остановка горутины с использованием канала
// в котором передается сигнал остановки OS (Graceful shutdown (SIGINT/SIGTERM)).
func main() {
	var wg sync.WaitGroup

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM) // подписываем SIGINT и SIGTERM на канал
	defer signal.Stop(sigCh)

	wg.Add(1)
	go func(wg *sync.WaitGroup, sig <-chan os.Signal) {
		defer wg.Done()

		ticker := time.NewTicker(300 * time.Millisecond)
		defer ticker.Stop()

		i := 0
		for {
			select {
			case <-sig:
				fmt.Println("quit by signal")
				return
			case <-ticker.C:
				fmt.Println(i)
				i++
			}
		}
	}(&wg, sigCh)

	wg.Wait()
}
