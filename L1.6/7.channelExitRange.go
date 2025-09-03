//go:build ignore

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func worker(jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Println(job)
	}
	fmt.Println("done")
}

// Способ выхода из горутины №7: остановка горутины с использованием закрытия входного канала
// воркеры сами выходит из range.
func main() {
	var wg sync.WaitGroup

	jobs := make(chan int, 2)

	wg.Add(1)
	go worker(jobs, &wg)

	jobs <- 2 * 5
	jobs <- (10 / 2)
	jobs <- rand.Intn(100)
	jobs <- rand.Intn(20) * 2
	close(jobs)

	wg.Wait()

}
