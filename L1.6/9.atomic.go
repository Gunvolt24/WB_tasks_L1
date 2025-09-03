//go:build ignore

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Способ выхода из горутины №9: остановка горутины с использованием пакета atomic
// atomic.Bool (или можно использовать atomic.Int32) в качестве стоп-флага.
// Горутина сама завершает функцию, периодически проверяя флаг.
// Важно помнить ограничение: atomic не умеет будить блокирующие операции (реакция на "stop" будет с задержкой до 50 мс (пока сон не закончится).
func main() {

	var ops atomic.Bool

	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			if ops.Load() {
				return
			}

			fmt.Println("working")
			time.Sleep(50 * time.Millisecond)
		}
	}(&wg)

	time.AfterFunc(300*time.Millisecond, func() {
		ops.Store(true)
		fmt.Println("stop")
	})

	wg.Wait()
}
