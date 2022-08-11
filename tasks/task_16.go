package tasks

import "fmt"

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func Quicksort(slice []int) {
	if len(slice) <= 1 {
		return
	}

	split := partition(slice)

	Quicksort(slice[:split])
	Quicksort(slice[split:])
}

func partition(slice []int) int {
	pivot := slice[len(slice)/2]

	left := 0
	right := len(slice) - 1

	for {
		for ; slice[left] < pivot; left++ {
		}

		for ; slice[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}

		slice[left], slice[right] = slice[right], slice[left]
	}
}

func Task16() {
	sl := []int{3, 4, 1, 2, 5, 7, -1, 0}
	Quicksort(sl)
	fmt.Println(sl)
}
