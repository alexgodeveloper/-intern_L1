package tasks

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов с использованием конкурентных вычислений.

func Task3() {

	var nums = []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup // синкгруппа для горутин
	var sum int
	for _, n := range nums {
		wg.Add(1) //+1 счетчик
		go func(n int) {
			sum += n * n
			wg.Done() // -1 счетчик
		}(n)
	}
	wg.Wait() //ожидание завершения всех
	fmt.Println(sum)
}
