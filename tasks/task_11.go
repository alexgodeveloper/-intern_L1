package tasks

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.

//Функция поиска пересечений
func Intersect(a, b []int) []int {
	m := make(map[int]int)
	intersectSlice := make([]int, 0)
	for _, v := range a {
		m[v] += 1 //счетчик числа по ключу +1
	}

	for _, v := range b {
		m[v] += 1 //счетчик числа по ключу +1
	}
	for k := range m {
		if m[k] > 1 { // все что больше 1 есть в обоих множествах
			intersectSlice = append(intersectSlice, k) // добавляем в срез
		}
	}
	return intersectSlice
}

func Task11() {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(Intersect(a, b))
}
