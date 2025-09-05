package main

import (
	"fmt"
	"reflect"
)

// detectType определяет тип переданного значения и выводит его.
func detectType(v any) (n int, err error) {
	switch v.(type) {
	case int:
		return fmt.Println("Тип:", reflect.TypeOf(v))
	case string:
		return fmt.Println("Тип:", reflect.TypeOf(v))
	case bool:
		return fmt.Println("Тип:", reflect.TypeOf(v))
	case chan int:
		return fmt.Println("Тип:", reflect.TypeOf(v))
	}

	if t := reflect.TypeOf(v); t != nil && t.Kind() == reflect.Chan {
		return fmt.Println("Тип:", t)
	}

	return fmt.Println("Неизвестный тип")
}

func main() {
	detectType(42)
	detectType("hello")
	detectType(true)
	detectType(make(chan <- int))
}
