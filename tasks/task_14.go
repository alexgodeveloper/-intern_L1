package tasks

import (
	"fmt"
)

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

func identifyType(v interface{}) {
	switch i := v.(type) { //Проверяем тип и ищем сопоставление
	case string:
		fmt.Printf("Интерфейс строка: %s\n", i)
	case int:
		fmt.Printf("Интерфейс число: %v\n", i)
	case bool:
		fmt.Printf("Интерфейс Булево: %v\n", i)
	case chan any:
		fmt.Printf("Интерфейс канал: %v\n", i)
	}
}

func Task14() {
	identifyType("Проверяем функцию")
	identifyType(make(chan any))
	identifyType(25)
	identifyType(true)

}
