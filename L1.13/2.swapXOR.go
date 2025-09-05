//go:build ignore

package main

import "fmt"

// swapXOR меняет значения двух целых чисел местами с помощью побитового XOR.
func swapXOR(a, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}

// Пример решения с помощью XOR.
func main() {
	a, b := 25, 60

	swapXOR(&a, &b)
	fmt.Println(a, b)
}
