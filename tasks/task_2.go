package tasks

//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

import (
	"fmt"
	"sync"
)

func Task2() {

	var nums = []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup // синкгруппа для горутин
	for _, num := range nums {
		wg.Add(1) // +1 в счетчик wg

		go func(num int) {
			square := num * num
			fmt.Println(square)
			wg.Done() // -1 в счетчик wg
		}(num)

	}
	wg.Wait() //ожидание всех горутин

}
