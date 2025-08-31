//go:build ignore

// Решение 2.
// Контекст в воркерах

// Идея: воркеры в select слушают и jobs, и ctx.Done(). По Ctrl+C бросают оставшиеся задачи и сразу выходят.
// Плюсы: мгновенная остановка.
// Минусы: возможна потеря необработанных задач из буфера.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): // мгновенно выходим по сигналу
			fmt.Printf("worker %d: stop\n", id)
			return
		case j, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Printf("worker %d: %d\n", id, j)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	workersNum := flag.Int("workers", 3, "number of workers")
	flag.Parse()

	jobs := make(chan int, 100)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	var wg sync.WaitGroup
	for i := 1; i <= *workersNum; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	go func() {
		defer close(jobs)
		i := 0
		for {
			select {
			case <-ctx.Done():
				return
			case jobs <- i:
				i++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	<-ctx.Done()
	wg.Wait()
}
