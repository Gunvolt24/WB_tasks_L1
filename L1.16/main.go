package main

import "fmt"

// Функция быстрой сортировки
// Быстрая сортировка работает по принципу "разделяй и властвуй".
// Она выбирает опорный элемент (pivot) и разделяет массив на элементы меньше и больше опорного.
func quickSort(list []int) []int {
	if len(list) < 2 { // Базовый случай: если массив пустой или содержит один элемент, он уже отсортирован
		return list
	} else {
		pivot := list[0] // Опорный элемент (list[0] - первый элемент)

		// Разделяй и властвуй. Создаем два списка:
		var less = []int{}             // Список для элементов меньше опорного
		var greater = []int{}          // Список для элементов больше опорного
		for _, num := range list[1:] { // Проходим по всем элементам, кроме опорного
			if pivot > num {
				less = append(less, num)
			} else {
				greater = append(greater, num)
			}
		}

		less = append(quickSort(less), pivot)
		greater = quickSort(greater)
		return append(less, greater...)
	}
}

func main() {
	fmt.Println(quickSort([]int{25, 51, 42, 12, 14, 3, 5, 7, 9, 11, 13, 15}))
}
