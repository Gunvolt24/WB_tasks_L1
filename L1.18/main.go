package main

import (
	"fmt"
	"sync"
)

// Создаем структуру Counter с мьютексом для безопасного инкремента
type Counter struct {
	mu  sync.Mutex
	num int
}

// Increment увеличивает значение счетчика на 1 в потокобезопасном режиме.
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.num++
}

// Value возвращает текущее значение счетчика в потокобезопасном режиме.
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.num
}

func main() {
	var wg sync.WaitGroup

	c := &Counter{}

	// Запускаем 451 горутину, каждая из которых увеличивает счетчик на 1
	for range 451 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("Финальное значение:", c.Value())
}
