//go:build ignore

package main

import (
	"context"
	"sync"
	"time"
)

// Функция sender отправляет данные в цикле и ожидает контекст.
func sender(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(120 * time.Millisecond) // Таймер с интервалом 120 мс
	defer ticker.Stop()                              // Останавливаем таймер
	for i := range 100 {
		select {
		case <-ctx.Done(): // Через 2 секунды останавливаем горутину.
			return
		case <-ticker.C:
			println(i) // Пока deadline не истекает продолжаем отправлять данные.
		}
	}
}

// Способ выхода из горутины №3: остановка горутины с использованием контекста WithDeadline.
func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go sender(ctx, &wg)

	wg.Wait()
}
