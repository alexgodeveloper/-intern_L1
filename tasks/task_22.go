package tasks

import (
	"fmt"
	"math"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

func Check(x int) bool {
	if x > int(math.Pow(2, 20)) {
		return true
	}

	return false
}

func Task22() {
	fmt.Println("Введите 1-е число")
	var a int
	fmt.Scan(&a)
	fmt.Println("Введите 2-е число")
	var b int
	fmt.Scan(&b)

	if (!Check(a)) || !Check(b) {
		fmt.Println("Числа меншье 2^20")
		return
	}

	fmt.Printf("Результат умножения: %d\n Результат деления: %g\n Результат сложения: %d\n Результат вычитания: %d", a*b, float32(a)/float32(b), a+b, a-b)
}
