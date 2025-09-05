package main

import (
	"fmt"
)

// unique - дженерик-функция, которая возвращает срез уникальных элементов из входного множества.
func unique[T comparable](words []T) []T {
	// Используем мапу для хранения уникальных элементов
	unique := make(map[T]struct{}, len(words))
	// Результирующий срез для уникальных элементов
	result := make([]T, 0, len(words))
	for _, word := range words {
		// Проверяем, есть ли элемент в мапе
		if _, ok := unique[word]; !ok {
			// Если нет, добавляем его в мапу и в результат
			unique[word] = struct{}{}
			result = append(result, word)
		}
	}
	return result

}

func main() {
	words := []string{"apple", "banana", "cherry", "date", "date", "fig", "grape", "kiwi", "apple", "banana", "cherry", "date", "fig", "grape", "kiwi"}
	fmt.Println("Длина исходного множества:", len(words))
	fmt.Println("Уникльные элементы:", unique(words))
	fmt.Println("Длина уникального множества:", len(unique(words)))
}
