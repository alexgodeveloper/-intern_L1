package tasks

import (
	"fmt"
	"unicode"
)

//26.	Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.
//Например:
//abcd — true
//abCdefAaf — false
//	aabcd — false

func UniqueSymbols(s string) bool {

	m := make(map[rune]string)
	for _, r := range s {

		// Приводим к одному регистру
		r = unicode.ToLower(r)

		// Если такой символ уже есть во множестве, возвращаем false
		if _, ok := m[r]; ok {
			return false
		}
		// Если символ ранее не встречался, добавляем его во множество
		m[r] = string(r)
	}
	return true
}

func Task26() {
	fmt.Println(UniqueSymbols("aaasdsasd"))
	fmt.Println(UniqueSymbols("asd"))
	fmt.Println(UniqueSymbols("321"))
	fmt.Println(UniqueSymbols("123123123"))

}
