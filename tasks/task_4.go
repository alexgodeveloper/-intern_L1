package tasks

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func workers(n uint, c <-chan int) {
	var i uint
	for i < n {
		go func(i uint) {
			for j := range c {
				fmt.Println(j)
			}
		}(i)
		i++
	}
}

func Task4() {
	ch := make(chan int)
	var w uint
	fmt.Println("Выберите число воркеров")
	fmt.Scan(&w)

	osSign := make(chan os.Signal)
	signal.Notify(osSign, syscall.SIGINT) // SIGINT - сигнал прерывания

	workers(w, ch)

	go func() {
		for {
			r := rand.Int()
			ch <- r
		}
	}()

	select {
	case <-osSign:
		fmt.Println("Принудительное завершение работы")
	}

}
