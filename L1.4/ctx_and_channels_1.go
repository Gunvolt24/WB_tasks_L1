//go:build ignore

// Решение 1.
// Контекст + закрытие jobs (graceful shutdown)

// Идея: по Ctrl+C останавливаем продюсеров и закрываем jobs. Воркеры читают range jobs и завершаются, когда канал закрыт.
// Плюсы: завершает все текущие задачи, корректно “дочищает” буфер.
// Минусы: завершает работу воркеров не мгновенно — дожидается обработки оставшихся задач.

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

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("worker %d: working on a job %d\n", id, job)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	workersNum := flag.Int("workers", 3, "number of workers")
	interval := flag.Duration("interval", 100*time.Millisecond, "time interval")
	flag.Parse()

	jobs := make(chan int, 100)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	var wg sync.WaitGroup
	for i := 1; i <= *workersNum; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	ticker := time.NewTicker(*interval)
	defer ticker.Stop()

	next := 0
	for {
		select {
		case <-ctx.Done(): // делаем graceful shutdown
			close(jobs)
			wg.Wait()
			fmt.Println("graceful shutdown")
			return
		case <-ticker.C:
			select {
			case jobs <- next:
				next++
			case <-ctx.Done():
				close(jobs)
				wg.Wait()
				return
			}
		}
	}
}
