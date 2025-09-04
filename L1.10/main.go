package main

import "fmt"

func main() {
	temp := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаем мапу для группировки температур по диапазонам
	groupByTemp := map[int][]float32{
		-20: {},
		10:  {},
		20:  {},
		30:  {},
	}

	// Группируем температуры по диапазонам
	for _, t := range temp {
		switch {
		case t < -20:
			groupByTemp[-20] = append(groupByTemp[-20], t)
		case t >= 10 && t < 19.9:
			groupByTemp[10] = append(groupByTemp[10], t)
		case t >= 20 && t < 29.9:
			groupByTemp[20] = append(groupByTemp[20], t)
		case t >= 30:
			groupByTemp[30] = append(groupByTemp[30], t)
		}
	}

	// Выводим сгруппированные температуры
	for k, v := range groupByTemp {
		fmt.Printf("Group %d: %.1f\n", k, v)
	}
}
