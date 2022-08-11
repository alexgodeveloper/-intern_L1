package tasks

import (
	"fmt"
)

//Разработать конвейер чисел.
//Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

// Канал для записи из массива
func set(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

//Канал для записи результата
func get(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()

	return out
}

func Task9() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	out1 := set(nums)
	out2 := get(out1)

	for r := range out2 {
		fmt.Println(r)
	}
}
