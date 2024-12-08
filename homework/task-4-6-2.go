package homework

import (
	"fmt"
	"sync"
)

// Программа с race condition
// Запуск go run -race main.go без mu дает предуспреждение "Found 2 data race(s)", а с mu все хорошо
func Run462() {
	counter := 0
	iterations := 10000

	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	// Горутина инкрементирует счетчик
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}()

	// Вторая горутина также инкрементирует счетчик
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
