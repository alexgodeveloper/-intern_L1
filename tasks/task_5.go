package tasks

import (
	"fmt"
	"math/rand"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func sendInChan(c chan<- int) {
	go func() {
		for {
			r := rand.Int()
			c <- r
		}
	}()
}

func receiveFromChan(c <-chan int) {
	go func() {
		for d := range c {
			fmt.Println(d)
		}
	}()
}

func Task5() {
	c := make(chan int)
	var sec time.Duration
	fmt.Println("Установите время работы в секундах:")
	fmt.Scan(&sec)

	sendInChan(c)
	receiveFromChan(c)

	select {
	case <-time.After(sec * time.Second):
	}
}
