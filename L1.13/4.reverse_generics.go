//go:build ignore

package main

import "fmt"

// reverseAnything переворачивает срез любого типа T.
func reverseAnything[T any](s []T) {
	// Переменные для указания первых и последних индексов
	first := 0
	last := len(s) - 1

	// Пока первый индекс меньше последнего, меняем элементы местами
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

// Пример решения задачи с использованием дженериков.
func main() {
	ints := []int{1, 8}

	reverseAnything(ints)
	fmt.Println(ints)
}
