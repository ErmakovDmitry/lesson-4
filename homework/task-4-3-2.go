package homework

import "sync"

// sumChannels
// in - слайс входных каналов, в которые приходят числа
// Признак окончания данных в канале - канал закрыт
func SumChannels(inputs []chan int64) int64 {

	resultChan := make(chan int64)

	// Эта переменная позволит дождаться завершения всех запущенных горутин
	var wg sync.WaitGroup

	// Подготовка и запуск горутин для обработки каждого канала
	for _, ch := range inputs {
		wg.Add(1)
		go worker(ch, resultChan, &wg)
	}

	// Запуск горутины, которая закрывает выходной канал,
	// когда все рабочие горутины завершат работу.
	go func() {
		wg.Wait() // Ждем завершения всех рабочих горутин
		close(resultChan)
	}()

	// Собираем результаты из выходного канала и суммируем их
	var totalSum int64 = 0
	for partialSum := range resultChan {
		totalSum += partialSum
	}

	return totalSum
}

// Рабочая функция, которая считывает данные из входного канала,
// складывает их и записывает результат в выходной канал.
func worker(inputChan <-chan int64, resultChan chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()

	var partialSum int64 = 0
	for num := range inputChan {
		partialSum += num
	}

	resultChan <- partialSum
}
