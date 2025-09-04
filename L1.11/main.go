package main

import "fmt"

// Функция для нахождения пересечения двух срезов при помощи вложенного цикла
func intersection(a, b []int) []int {
	var result []int
	for _, elem1 := range a {
		for _, elem2 := range b {
			if elem1 == elem2 {
				result = append(result, elem1)
				break
			}
		}
	}
	return result
}

func main() {
	a := []int{6, 7, 8, 9, 10}
	b := []int{11, 9, 7, 12, 6}

	fmt.Println("Пересечение:", intersection(a, b))
}
