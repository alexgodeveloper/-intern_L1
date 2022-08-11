package tasks

import (
	"fmt"
	"sync"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

type Incrementer struct {
	Inc int
	mu  sync.Mutex
}

func (i *Incrementer) Increment(wg *sync.WaitGroup) {
	defer wg.Done() //отложенный -1 из счетчка
	i.mu.Lock()
	defer i.mu.Unlock() // отложенное разблокирование мьютекса
	i.Inc++
}

func Task18() {
	var i int
	wg := sync.WaitGroup{} //синкгруппа
	inc := Incrementer{mu: sync.Mutex{}}
	for i < 9999999 {
		wg.Add(1) //+1 в счетчик
		go inc.Increment(&wg)
		i++
	}
	wg.Wait()
	fmt.Println(inc.Inc)
}
