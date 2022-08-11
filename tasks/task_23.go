package tasks

import "fmt"

//Удалить i-ый элемент из слайса.

func Task23() {
	fmt.Println("Какой элемент удалть из слайса?")
	var a int
	fmt.Scan(&a)
	users := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	users = append(users[:a-1], users[a:]...)
	fmt.Println(users)
}
