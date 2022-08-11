package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func Task20() {
	fmt.Println("Введите строку")
	scanner := bufio.NewScanner(os.Stdin) //Получаем введенный текст в консоли
	scanner.Scan()
	text := scanner.Text()
	str := strings.Split(text, " ") //разделяем по пробелу
	var newStr string
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}

	newStr = strings.Join(str, " ") //соединяем с пробелом в строку
	fmt.Println(newStr)
}
