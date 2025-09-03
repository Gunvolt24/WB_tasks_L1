package main

import (
	"fmt"
	"strconv"
)

// writeBit устанавливает i-й бит в 0 или 1.
// 1) Создаём маску у которой установлен только i-й бит в 1 -> остальные биты в 0: единственная 1 в маске сдвигается на i позцию,
// влево, поскольку int64(1) - это число, у которого в двоичном виде все биты 0, кроме младшего: ...00000001.
// 2) Делаем нормализацию b к 0 или 1 - берем младший бит параметра b с помощью & 1, чтобы получить строго 0 или 1.
// 2.1) Если b == 1 - делаем побитовое OR: n | mask - это устанавливает бит в 1.
// 2.2) Если b == 0 - используем &^ (AND NOT) - это устанавливает бит в 0.
func writeBit(n int64, pos uint, b uint8) int64 {
	mask := int64(1) << pos
	if (b & 1) == 1 {
		return n | mask // установить бит в 1
	}
	return n &^ mask // установить бит в 0
}

// convertToBinary возвращает строку с двоичным представлением числа фиксированной длины 64 бита.
// Приведение к uint64 позволяет корректно показать отрицательные int64
// в виде их двоичного представления.
func convertToBinary(n int64) string {
	return fmt.Sprintf("%064b", uint64(n))
}

func main() {
	var n int64   // начальное значение (в диапазоне 0..63 бит)
	var pos uint  // позиция бита, используем тип uint для сдвига
	var val uint8 // установка значения бита: 0 или 1

	fmt.Print("Число: ")
	fmt.Scan(&n)

	fmt.Print("Позиция бита (0..63): ")
	fmt.Scan(&pos)

	fmt.Print("Значение бита (0/1): ")
	fmt.Scan(&val)

	if pos > 63 {
		fmt.Println("Ошибка: позиция должна быть в диапазоне 0..63 для int64")
		return
	}

	if val != 0 && val != 1 {
		fmt.Println("Ошибка: значение бита должно быть 0 или 1")
		return
	}

	before := n                    // запомним исходное значение
	after := writeBit(n, pos, val) // установка бита pos в значение val (0 или 1)

	fmt.Printf("До: %s (%s)\n", strconv.FormatInt(before, 10), convertToBinary(before))
	fmt.Printf("После: %s (%s)\n", strconv.FormatInt(after, 10), convertToBinary(after))
}
