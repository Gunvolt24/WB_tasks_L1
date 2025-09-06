package main

import "fmt"

// binarySearch выполняет бинарный поиск числа i в отсортированном срезе numbers итеративно.
// Возвращает индекс элемента, если он найден, или -1, если элемента нет в списке.
func binarySearch(numbers []int, i int) int {
	// Инициализируем границы поиска
	low := 0
	high := len(numbers) - 1

	// Пока границы не пересеклись
	for low <= high {
		// Вычисляем средний индекс и сравниваем значение с искомым
		midIdx := (low + high) / 2
		midVal := numbers[midIdx]
		// В зависимости от сравнения сдвигаем границы поиска
		switch {
		case midVal == i:
			return midIdx
		case midVal < i:
			low = midIdx + 1
		case midVal > i:
			high = midIdx - 1
		}
	}
	return -1
}

// binarySearchRecursive выполняет бинарный поиск числа i в отсортированном срезе numbers рекурсивно.
// Возвращает индекс элемента, если он найден, или -1, если элемента нет в списке.
func binarySearchRecursive(numbers []int, i, low, high int) int {
	// Базовый случай: если диапазон поиска недействителен, возвращаем -1
	if low > high {
		return -1
	}
	// Вычисляем средний индекс и сравниваем значение с искомым
	midIdx := (low + high) / 2
	midVal := numbers[midIdx]
	// Если найден, возвращаем индекс
	if midVal == i {
		return midIdx
	}
	// Рекурсивно ищем в левой или правой половине
	if numbers[midIdx] < i {
		return binarySearchRecursive(numbers, i, midIdx+1, high)
	} else {
		return binarySearchRecursive(numbers, i, low, midIdx-1)
	}
}

func main() {
	list := []int{12, 13, 56, 66, 78, 99, 124, 245, 542}
	// 5 нету в списке
	fmt.Println("Индекс элемента в списке:", (binarySearch(list, 5)))
	// 124 есть в списке и его индекс: 6
	fmt.Println("Индекс элемента в списке:", (binarySearch(list, 124)))

	fmt.Println("Индекс элемента в списке (рекурсивный метод):", (binarySearchRecursive(list, 5, 0, len(list)-1)))
	fmt.Println("Индекс элемента в списке (рекурсивный метод):", (binarySearchRecursive(list, 542, 0, len(list)-1)))
}
