package main

import (
	"fmt"
	"strings"
)

// findUniqueLetters проверяет, являются ли все буквы в строке уникальными
// Возвращает true, если все буквы уникальны, и false в противном случае
// Переводим строку в нижний регистр, чтобы сравнивать символы независимо от регистра
func findUniqueLetters(s string) bool {
	s = strings.ToLower(s)

	// Создаем мапу для хранения уникальных символов
	// Ключи — символы (в rune - для работы с юникодом), значения — флаг, указывающий, был ли символ уже встречен
	checkUniqueness := make(map[rune]bool)

	// Проходим по каждому символу в строке и проверяем, является ли он уникальным
	// Если символ уже встречался, возвращаем false
	for _, letter := range s {
		if checkUniqueness[letter] {
			return false
		}
		checkUniqueness[letter] = true
	}
	return true
}

func main() {
	letters := []string{
		"abcdefghijklmnopqrstuvwxyz",                  // true
		"ABcDeFgHiJkLmNoPqRsTuVwXyZabCdeFgHiJkLmNoPq", // false
		"aabbccdeefgghiijkllmnnoppqrrsttuuvvwwxyzz",   // false
		"abcd",                       // true
		"abBcd",                      // false
		"AbCdEfGhIjKlMnOpQrStUvWxYz", // true
	}
	for _, s := range letters {
		fmt.Println(findUniqueLetters(s))
	}
}
