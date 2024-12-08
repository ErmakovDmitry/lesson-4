package homework

import (
	"fmt"
	"sync"
	"time"
)

// Функция для проверки доступности внешнего сервиса
func checkServiceAvailable() bool {

	// Логика проверки доступности сервиса

	return true
}

func Run453() {
	const numWorkers = 5
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)
	stopWorkers := false

	var wg sync.WaitGroup

	// Функция обработки
	worker := func(id int) {
		defer wg.Done()
		for {
			mutex.Lock()
			// Ожидать, пока флаг stopWorkers не станет ложным
			for stopWorkers {
				cond.Wait()
			}
			mutex.Unlock()

			// Имитация получения задачи и обработки
			fmt.Printf("Worker %d is processing a task\n", id)
			time.Sleep(1 * time.Second) // Имитация времени обработки
		}
	}

	// Запуск рабочих горутин
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i)
	}

	// Горутина для проверки доступности сервиса
	go func() {
		for {
			time.Sleep(2 * time.Second) // Период проверки
			mutex.Lock()
			if checkServiceAvailable() {
				fmt.Println("Service is available. Notifying workers...")
				stopWorkers = false
				cond.Broadcast() // Разбудить все горутины
			} else {
				fmt.Println("Service is unavailable. Workers are waiting...")
				stopWorkers = true
			}
			mutex.Unlock()
		}
	}()

	// Ожидание завершения всех рабочих горутин
	wg.Wait()
}
