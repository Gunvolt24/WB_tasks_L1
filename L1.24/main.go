package main

import (
	"fmt"
	"math"
)

// Point — структура для хранения координат клетки (x — строка, y — столбец).
type Point struct {
	x float64
	y float64
}

// NewPoint создаёт новую точку с координатами (x, y).
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// Distance нахождение расстояния между двумя точками
// Формула вычисления расстояния между двумя точками: sqrt(x2 - x1)^2 + (y2 - y1)^2
func (p *Point) Distance(other Point) float64 {
	distX := other.x - p.x
	distY := other.y - p.y
	return math.Hypot(distX, distY)
}

func main() {
	p1 := NewPoint(20, 30)
	p2 := NewPoint(40, 60)

	dist := p1.Distance(p2)
	fmt.Println("Distance", dist)
}
