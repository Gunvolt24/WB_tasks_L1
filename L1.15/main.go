// Пример решения задачи:
// В функции someFunc создается большая строка v размером 1024 байта.
// Затем из этой строки берется срез первых 100 байт и сохраняется в глобальную переменную justString.
// Однако, из-за особенностей работы строк в Go, это приводит к удержанию всей большой строки v в памяти,
// что может вызвать утечку памяти, если someFunc вызывается многократно.
// Чтобы избежать этого, нужно явно скопировать нужные байты в justString.

package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(size int) string {
	return string(make([]byte, size))
}

func someFunc() {
	v := createHugeString(1024)
	fmt.Println("Длина переменной v:", len(v))
	n := min(100, len(v))
	justString = strings.Clone(v[:n])
}

func main() {
	someFunc()
	fmt.Println("Длина переменной justString:", len(justString))
}
