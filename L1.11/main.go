package main

import "fmt"

// Дженерик-функция для нахождения пересечения двух множеств, если T - сравнимый тип.
func intersection[T comparable](a, b []T) []T {
	// Создаем мапу для хранения элементов первого среза
	m := make(map[T]struct{}, len(a))

	// Заполняем мапу элементами из первого среза
	for _, v := range a {
		m[v] = struct{}{}
	}

	// Результирующий срез для пересечения
	result := make([]T, 0)
	// Проверяем элементы второго среза на наличие в мапе
	for _, v := range b {
		if _, found := m[v]; found {
			result = append(result, v) // Добавляем пересекающийся элемент в результат
			delete(m, v)               // Удаляем элемент из мапы, чтобы избежать дубликатов
		}
	}
	return result
}

func main() {
	a := []int{6, 7, 8, 9, 10}
	b := []int{11, 9, 7, 12, 6}

	fmt.Println("Пересечение:", intersection(a, b))
}
