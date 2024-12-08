package homework

import (
	"fmt"
	"sync"
	"time"
)

// semaphore - это канал из пустых структур
type semaphore chan struct{}

// Функция для создания нового семафора
func NewSemaphore(n int) semaphore {
	return make(semaphore, n)
}

// Метод Acquire позволяет занять n мест в семафоре
func (s semaphore) Acquire(n int) {
	for i := 0; i < n; i++ {
		s <- struct{}{}
	}
}

// Метод Release освобождает n мест в семафоре
func (s semaphore) Release(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

func Run452() {
	const (
		totalRequests = 1000 // Общее количество запросов
		rateLimit     = 10   // Ограничение на количество параллельных запросов
	)

	// Создание семафора с лимитом параллельных запросов
	sem := NewSemaphore(rateLimit)

	var wg sync.WaitGroup
	start := time.Now()

	// Функция, имитирующая выполнение запроса
	processRequest := func(reqID int) {
		defer wg.Done()

		// Занимаем место в семафоре
		sem.Acquire(1)

		// Имитация выполнения запроса с помощью Sleep
		time.Sleep(1 * time.Second)

		// Освобождаем место в семафоре
		sem.Release(1)

		// Вывод времени выполнения запроса
		fmt.Printf("Request %d completed at %v\n", reqID, time.Since(start))
	}

	// Запуск выполнения всех запросов
	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		go processRequest(i)
	}

	fmt.Println("All run!")

	// Ожидание завершения всех запросов
	wg.Wait()

	fmt.Println("All finished!")
}
