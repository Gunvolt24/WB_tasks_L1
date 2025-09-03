package main

import (
	"fmt"
	"sync"
)

// Способ выхода из горутины №1: остановка горутины возвратом по условию.
func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range 10 {
			fmt.Println(i)
			if i == 8 {
				fmt.Println("done")
				return
			}
		}
	}()

	wg.Wait()
}
