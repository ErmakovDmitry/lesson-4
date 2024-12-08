package homework

import (
	"fmt"
	"sync"
	"time"
)

// Принимает слайс каналов, ожидает каждый из них и обрабатывает полученные сообщения
func selectMany(channels []chan int64) chan int64 {
	// Создание итогового канала для результатов
	resultChan := make(chan int64)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	// Для каждого канала создается горутина для чтения данных
	for _, ch := range channels {
		go func(c chan int64) {
			defer wg.Done()
			for val := range c {
				resultChan <- val
			}
		}(ch)
	}

	// Горутина для закрытия итогового канала, когда все данные будут получены
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	return resultChan
}

func Run472() {
	ch1 := make(chan int64)
	ch2 := make(chan int64)

	go func() {
		for i := int64(0); i < 5; i++ {
			fmt.Println("Пишем в ch1:", i)
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := int64(5); i < 10; i++ {
			fmt.Println("Пишем в ch2:", i)
			ch2 <- i
			time.Sleep(100 * time.Millisecond) // Для демонстрации параллельной обработки
		}
		close(ch2)
	}()

	resultChan := selectMany([]chan int64{ch1, ch2})

	for res := range resultChan {
		fmt.Println("res:", res)
	}
}
