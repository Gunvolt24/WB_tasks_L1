//go:build ignore

package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

// Способ выхода из горутины №8: остановка горутины с использованием errgroup + context.
// Появление ошибки в одной горутине - отменяет все остальные горутины - возвращается первая non-nil ошибка.
func main() {
	// Создаем группу и производный контекст
	errs, ctx := errgroup.WithContext(context.Background())

	errs.Go(func() error {
		// Данная горутина не сможет выполниться, потому что к моменту выполнения контекст был отменен
		select {
		case <-ctx.Done():
			// Handle Cancelation
			return ctx.Err()
		case <-time.After(5 * time.Second):
			fmt.Println("Длинная пауза не успела выполниться")
			return nil
		}
	})

	errs.Go(func() error {
		// Данная горутина успеет выполниться, потому что она возвращает nil
		select {
		case <-ctx.Done():
			// Handle Cancelation
			return ctx.Err()
		case <-time.After(10 * time.Millisecond):
			fmt.Println("Короткая пауза успела выполниться")
			return nil // возвращаем nil
		}
	})
	errs.Go(func() error {
		// Данная горутина не успеет выполниться, потому что она возвращает non-nil
		select {
		case <-ctx.Done():
			// Handle Cancelation
			return ctx.Err()
		case <-time.After(100 * time.Millisecond):
			return fmt.Errorf("получена non-nil ошибка") // возвращаем non-nil
		}
	})

	// Wait ждет пока все горутины завершатся и возвращает первую non-nil ошибку
	if err := errs.Wait(); err != nil {
		fmt.Printf("Получена первая ошибка: %v", err)
	}
}
