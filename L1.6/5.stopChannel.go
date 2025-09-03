//go:build ignore

package main

import (
	"fmt"
	"sync"
	"time"
)

// Способ выхода из горутины №5: остановка горутины с использованием канала
// в котором передается сигнал остановки.
func main() {
	quit := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup, quit <-chan struct{}) {
		defer wg.Done()

		ticker := time.NewTicker(300 * time.Millisecond)
		defer ticker.Stop()

		r := 'a'
		for {
			select {
			case <-quit:
				fmt.Println("quit by channel")
				return
			case <-ticker.C:
				fmt.Printf("letter: %c\n", r)
				if r == 'z' {
					r = 'a'
				} else {
					r++
				}
			}
		}
	}(&wg, quit)

	time.AfterFunc(5*time.Second, func() {
		close(quit)
	})

	wg.Wait()
}
