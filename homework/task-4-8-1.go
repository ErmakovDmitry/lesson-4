package homework

import (
	"fmt"
	"sync"
	"time"
)

// Напишите программу, которая реализует broadcast события во все горутины.
// Эта задача похожа на worker pool, но в worker pool только одна горутина получит событие.
// Хотим, чтобы событие было доставлено всем горутинам в рамках одного worker pool.
func Run481() {
	const numWorkers = 5
	broadcastChan := make(chan string)
	doneChan := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	worker := func(id int, broadcastChan <-chan string, doneChan <-chan struct{}) {
		defer wg.Done()
		for {
			select {
			case msg := <-broadcastChan:
				fmt.Printf("Worker %d received message: %s\n", id, msg)
			case <-doneChan:
				fmt.Printf("Worker %d exiting\n", id)
				return
			}
		}
	}

	// Запуск рабочих горутин
	for i := 0; i < numWorkers; i++ {
		go worker(i, broadcastChan, doneChan)
	}

	// Отправка нескольких сообщений
	go func() {
		broadcastMessages := []string{"Event 1", "Event 2", "Event 3"}
		for _, msg := range broadcastMessages {
			broadcastChan <- msg
			time.Sleep(1 * time.Second) // Имитируем задержку между событиями
		}
		close(doneChan) // Завершаем работу всех горутин
	}()

	// Ожидание завершения всех рабочих горутин
	wg.Wait()
	fmt.Println("All workers have exited")
}
