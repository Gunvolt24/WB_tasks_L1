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
	for i := range 100 {
		select {
		case <-ctx.Done(): // Через 3 секунды останавливаем горутину.
			return
		default:
			println(i) // Пока таймер не истекает продолжаем отправлять данные.
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// Способ выхода из горутины №2: остановка горутины с использованием контекста отмены.
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go sender(ctx, &wg)

	time.AfterFunc(3*time.Second, cancel) // Задаем время работы функции и останавливаем ее по контексту.
	wg.Wait()
}
