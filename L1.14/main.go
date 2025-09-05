package main

import (
	"fmt"
	"reflect"
)

// detectType определяет тип переданного значения и выводит его.
func detectType(v any) string {
	switch v.(type) {
	case int:
		fmt.Println("Тип:", reflect.TypeOf(v))
	case string:
		fmt.Println("Тип:", reflect.TypeOf(v))
	case bool:
		fmt.Println("Тип:", reflect.TypeOf(v))
	case chan int:
		fmt.Println("Тип:", reflect.TypeOf(v))
	default:
		fmt.Println("Неизвестный тип")
	}
	return "Ни один из типов не подошёл"
}

func main() {
	detectType(42)
	detectType("hello")
	detectType(true)
	detectType(make(chan int))
}
