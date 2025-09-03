package main

import (
	"fmt"
	"sync"
)

// SafeMap - это потокобезопасная структура данных для хранения данных.
type SafeMap struct {
	mu      sync.RWMutex
	numbers map[int]int
}

// Set устанавливает значение квадрата числа по ключу.
func (s *SafeMap) Set(key, val int) {
	s.mu.Lock()
	s.numbers[key] = val
	s.mu.Unlock()
}

// Get возвращает значение квадрата числа по ключу и булево значение, указывающее, найдено ли число.
func (s *SafeMap) Get(key int) (int, bool) {
	s.mu.RLock()
	val, ok := s.numbers[key]
	s.mu.RUnlock()
	return val, ok
}

func main() {
	wg := sync.WaitGroup{}
	// Создаем экземпляр SafeMap
	sm := &SafeMap{numbers: map[int]int{}}

	// Запускаем горутины для записи данных
	for key := range 100 {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			sm.Set(key, key*key)
		}(key)
	}

	// Запускаем горутины для чтения данных
	for key := range 100 {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			if val, ok := sm.Get(key); ok {
				fmt.Printf("Число (key): %d, его квадрат (value): %d\n", key, val)
			} else {
				fmt.Printf("Число (key) %d не найдено\n", key)
			}
		}(key)
	}

	wg.Wait()
}
