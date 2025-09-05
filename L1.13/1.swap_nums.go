package main

import "fmt"

// Пример решения при помощи свапа двух чисел.
func main() {
	a, b := 25, 72

	a, b = b, a

	fmt.Println(a, b)
}
