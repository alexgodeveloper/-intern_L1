package tasks

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

// Структура
type Point struct {
	x float64
	y float64
}

//Конструктор
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

//Функция расчета расстояния по теореме Пифагора
func Distance(a, b *Point) float64 {
	x := a.x - b.x
	y := a.y - b.y
	return math.Sqrt(x*x + y*y)
}

func Task24() {
	a := NewPoint(10, 7)
	b := NewPoint(-25, 16)

	fmt.Println(Distance(a, b))
}
