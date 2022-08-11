package tasks

import (
	"fmt"
	"strings"
)

//15.	К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
/*var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}*/

var justString string

func createHugeString(i int) string {
	s := ""
	for i > 0 {
		s += "1"
		i--
	}
	return s
}
func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100]) //Как вариант, чтобы убрать лишнее из памяти.
	fmt.Println(justString)
}

func Task15() {
	someFunc()
}
