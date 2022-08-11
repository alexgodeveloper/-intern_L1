package tasks

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func CreateSlice(sl []string) map[string][]string {
	m := make(map[string][]string)
	slice := make([]string, 0)
	for _, k := range sl {
		m[k] = append(slice, k)
	}

	return m
}

func Task12() {
	example := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(CreateSlice(example))

}
