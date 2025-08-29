package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Println("worker", id, "working on a job", job)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	workersNum := flag.Int("workers", 30, "number of workers")
	interval := flag.Duration("interval", 100*time.Millisecond, "time interval")
	flag.Parse()

	jobs := make(chan int, 100)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup
	for i := 1; i <= *workersNum; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	ticker := time.NewTicker(*interval)
	defer ticker.Stop()

	i := 0
	for {
		select {
		case <-ctx.Done():
			close(jobs)
			wg.Wait()
			return
		case <-ticker.C:
			select {
			case jobs <- i:
				i++
			case <-ctx.Done():
				close(jobs)
				wg.Wait()
				return
			}
		}
	}
}
