package main

import (
	"fmt"
	"sync"
)

func processData(wg *sync.WaitGroup, resultDest *int, data int) {
	defer wg.Done()

	processedData := data * data

	*resultDest = processedData
}

func main() {
	var wg sync.WaitGroup

	input := [5]int{2, 4, 6, 8, 10}   // Исходный массив
	result := make([]int, len(input)) // Инициализируем срез с нужной длиной

	for i, data := range input {
		wg.Add(1)
		go processData(&wg, &result[i], data) // В &result[i] - пердаем адрес элемента массива
	}

	wg.Wait()

	fmt.Println("Массив в квадрате:", result)
}
