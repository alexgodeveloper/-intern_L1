package tasks

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.

func Task6() {
	var wg sync.WaitGroup

	// Читающая из канала рутина завершит работу при закрытии канала
	c := make(chan int)
	wg.Add(1)
	go func() {
		fmt.Println("Старт горутины")
		for i := range c {
			fmt.Printf("Чтение из канала %d\n", i)
		}
		fmt.Println("Горутина остановлена")
		wg.Done()
	}()
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Закрываем канал")
	close(c)
	wg.Wait() // Ожидается закрытие всех рутин и писать в закрытый канал уже нет возможности

	// Отдельный канал, при отправке в который рутина завершает работу
	closed := make(chan bool)
	wg.Add(1)
	go func() {
		fmt.Println("Старт горутины")
		i := 0
		for {
			select {
			case <-closed:
				fmt.Println("Горутина остановлена")
				wg.Done()
				return
			default:
				i++
				fmt.Println(i)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(500 * time.Millisecond)
	fmt.Println("Передаем значение в канал закрытия")
	closed <- true
	wg.Wait() // ожидаем завершения всех рутин

	// Отдельный канал, при закрытии которого рутина завершает работу
	// Так можно завершить работу нескольких рутин, в отличии от предыдущего варианта
	done := make(chan bool)
	f := func(n int) {
		fmt.Printf("Горутина %d стартовала\n", n)
		i := 0
		for {
			select {
			case <-done:
				fmt.Printf("Горутина %d остановлена\n", n)
				wg.Done()
				return
			default:
				i++
				fmt.Printf("Горутина %d: обрабатывает значение: %d\n", n, i)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
	wg.Add(2)
	go f(1)
	go f(2)

	time.Sleep(500 * time.Millisecond)
	fmt.Println("Закрываем канал")
	close(done)
	wg.Wait()

	// Использование контекста
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context) {
		fmt.Println("Горутина стартовала")
		i := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина остановлена")
				wg.Done()
				return
			default:
				i++
				fmt.Println(i)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}(ctx)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Контекст отмены")
	cancel()
	wg.Wait()

	// Завершение работы по условию
	var a int64
	wg.Add(1)
	go func() {
		fmt.Println("Горутина стартовала")
		for a < 1 {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Значение: %d\n", a)
		}
		fmt.Println("Горутина остановлена")
		wg.Done()
	}()

	time.Sleep(500 * time.Millisecond)
	fmt.Println("увеличиваем а")
	//atomic.AddInt64(&a, 1)
	a += 1
	wg.Wait()
}
