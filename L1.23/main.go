package main

import "fmt"

// deleteElem удаляет элемент по индексу
func deleteElem(slice []int, index int) []int {
	// проверяем, что индекс в допустимом диапазоне
	if index < 0 || index >= len(slice) {
		return slice
	}
	copy(slice[index:], slice[index+1:]) // сдвигаем хвост на место удаляемого элемента
	slice[len(slice)-1] = 0              // удаляем последний элемент
	slice = slice[:len(slice)-1]         // сокращаем срез
	return slice
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(deleteElem(slice, 4))

}
