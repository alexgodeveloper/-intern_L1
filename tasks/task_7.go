package tasks

import (
	"fmt"
	"sync"
	"time"
)

//Реализовать конкурентную запись данных в map.

func Task7() {
	m := make(map[string]int)
	mut := sync.Mutex{}
	i := 0

	for i < 10 {
		go func() {
			mut.Lock()
			defer mut.Unlock()
			m["counter"]++
			time.Sleep(1 * time.Second)
			fmt.Println(m["counter"])
		}()
		i++
	}

	fmt.Println(m)

}
