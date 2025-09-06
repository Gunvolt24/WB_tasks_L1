package main

import (
	"fmt"
)

// ReversedString возвращает строку в обратном порядке
func ReversedString(s string) string {
	arr := []rune(s)                                               // Преобразуем строку в срез рун для корректной работы с юникодом
	for i, j := 0, len(arr)-1; i < len(arr)/2; i, j = i+1, j-1 {  // Меняем элементы местами, начиная с концов среза и двигаясь к центру
		arr[i], arr[j] = arr[j], arr[i] 								 // Меняем местами элементы с индексами i и j
	}
	return string(arr)
}

func main() {
	s := "главрыба 🐟. рабрыба 🐠, рыба-ёж 🐡, акула 🦈, суши 🍥"
	fmt.Println(ReversedString(s))
}
