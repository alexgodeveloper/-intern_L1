package tasks

import "fmt"

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

//Функция побитового изменения, где x - вводимое число, i - кол-во разрядов, v - значение, в которое нужно установить.
func Task8(x int64, i int, v int) {
	var res int64
	var mask int64 = 00000001
	if v == 1 {
		res = x | mask<<(i-1)
	} else if v == 0 {
		res = x & ^(mask << (i - 1))
	} else {
		fmt.Println("Неправильное значение для битовых операций.(0 или 1)")
	}
	fmt.Println(res)
}
