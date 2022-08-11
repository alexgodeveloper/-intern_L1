package tasks

import "fmt"

//Поменять местами два числа без создания временной переменной.

func Task13() {
	// Да, в го так можно))
	var a, b = 10, 5
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}
