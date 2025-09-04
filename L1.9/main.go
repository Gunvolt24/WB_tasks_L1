// При решении опирался на материал из статьи:
// https://kovardin.ru/articles/go/modeli-konkurentnosti-v-go/

package main

import (
	"fmt"
	"sync"
)

// Перевый этап.
// convertToChannel представляет собой фунцию, которая преобразует срез целых чисел в канал.
// Внутри этой функции запускается горутина, которая отправляет целые числа в канал и закрывает этот канал, когда все числа отправлены.
func convertToChannel(nums []int, wg *sync.WaitGroup) <-chan int {
	out := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

// Второй этап.
// Функция conveyor забирает числа из канала и возвращает новый канал, который отдает квадрат каждого полученного числа.
// После того как входящий канал закрыт и все значения на этом шаге отправлены в исходящий канал, то исходящий канал закрывается.
func conveyor(in <-chan int, wg *sync.WaitGroup) <-chan int {
	out := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range in {
			out <- num * 2
		}
		close(out)
	}()
	return out
}

func main() {
	wg := sync.WaitGroup{}

	// Создаем каналы и запускаем конвейер
	x := convertToChannel([]int{11, 22, 33, 44, 55}, &wg)
	out := conveyor(x, &wg)

	// Читаем результаты из канала out
	for v := range out {
		fmt.Println(v)
	}

	wg.Wait()
}
