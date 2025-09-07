package main

import (
	"fmt"
)

// reverseSlice переворачивает элементы среза s от индекса left до индекса right включительно.
// Использует обобщения для работы с любым типом элементов.
func reverseSlice[T any](s []T, left, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// reverseWords переворачивает порядок слов в строке s.
// Слова считаются последовательностями символов, разделёнными пробелами.
func reverseWords(s string) string {
	// Преобразуем строку в срез рун для корректной работы с юникодом
	runes := []rune(s)
	length := len(runes)
	if length == 0 {
		return s
	}

	// Переворачиваем всю строку
	reverseSlice(runes, 0, length-1)

	// Переворачиваем каждое слово в строке (между пробелами)
	wordStart := 0
	for wordStart < length {
		// Пропускаем пробелы
		for wordStart < length && runes[wordStart] == ' ' {
			wordStart++
		}
		// Если достигли конца строки, выходим из цикла
		if wordStart >= length {
			break
		}

		// Находим конец слова runes[wordStart:wordEnd]
		wordEnd := wordStart
		for wordEnd < length && runes[wordEnd] != ' ' {
			wordEnd++
		}

		// Переворачиваем текущее слово на месте
		reverseSlice(runes, wordStart, wordEnd-1)

		// Переходим к следующему слову
		wordStart = wordEnd
	}

	return string(runes)
}

func main() {
	s := "акула рыба-меч рыба-ёж дельфин пингвин кит морж"
	fmt.Println(reverseWords(s))
}
