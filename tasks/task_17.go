package tasks

import (
	"fmt"
	"sort"
)

//Реализовать бинарный поиск встроенными методами языка.

func Search(slice []int, x int) int { //Передаем слайс и число которое будем искать
	sort.Ints(slice)
	return sort.Search(len(slice), func(i int) bool {
		return x <= slice[i]
	})
}

func Task17() {
	a := []int{6, 1, 2, 7, 3, 9}
	x := 10
	i := Search(a, x)
	if i < len(a) && a[i] == x {
		fmt.Printf("Найдено %d по индексу %d в %v.\n", x, i, a)
	} else {
		fmt.Printf("Не найдено %d в %v.\n", x, a)
	}
}
