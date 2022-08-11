package tasks

import (
	"fmt"
)

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	Name string
}

func (h Human) Sing() {
	fmt.Println(h.Name + " is singing")
}

// Структура Action со встроенными структурами
type Action struct {
	Human
}

func Task1() {
	h := Human{
		Name: "Olya",
	}

	a := Action{
		Human: Human{
			Name: "Action-Olya",
		},
	}

	h.Sing()
	a.Sing()

}
