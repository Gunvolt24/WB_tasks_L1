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
	ticker := time.NewTicker(100 * time.Millisecond) // Таймер с интервалом 100 мс
	defer ticker.Stop()                              // Останавливаем таймер
	for i := range 100 {
		select {
		case <-ctx.Done(): // Через 3 секунды останавливаем горутину.
			return
		case <-ticker.C:
			println(i) // Пока таймаут не истекает продолжаем отправлять данные.
		}
	}
}

// Способ выхода из горутины №4: остановка горутины с использованием контекста WithTimeout.
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go sender(ctx, &wg)

	wg.Wait()
}
