package tasks

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

func MySleep(dur time.Duration) {
	<-time.After(dur)
}

func Task25() {
	MySleep(5 * time.Second)
	fmt.Println("Проснулись")
}
