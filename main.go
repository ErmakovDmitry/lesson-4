package main

import (
	"bufio"
	"fmt"
	"strings"

	hw "lesson-4/homework"
)

func main() {

	// task-4-3-2
	channels := []chan int64{
		make(chan int64),
		make(chan int64),
		make(chan int64),
	}
	go func() {
		channels[0] <- 1
		channels[0] <- 2
		channels[0] <- 3
		close(channels[0])
	}()
	go func() {
		channels[1] <- 4
		channels[1] <- 5
		close(channels[1])
	}()
	go func() {
		channels[2] <- 6
		channels[2] <- 7
		channels[2] <- 8
		close(channels[2])
	}()
	totalSum := hw.SumChannels(channels)
	fmt.Println("Общая сумма:", totalSum)

	// task-4-4-2
	fmt.Printf("\ntask-4-4-2\n")
	input1 := make(chan string)
	go func() {
		defer close(input1)
		scanner := bufio.NewScanner(strings.NewReader("пример  текста.    с  множеством.   Пробелов  и    предложений4!  "))
		for scanner.Scan() {
			input1 <- scanner.Text()
		}
	}()
	// fmt.Println("input1")
	// for message := range input1 {
	// 	fmt.Println(message)
	// }

	output1 := make(chan string)
	go hw.Step1(input1, output1)
	// fmt.Println("output1")
	// for message := range output1 {
	// 	fmt.Println(message)
	// }

	output2 := make(chan string)
	go hw.Step2(output1, output2)
	// fmt.Println("output2")
	// for message := range output2 {
	// 	fmt.Println(message)
	// }

	output3 := hw.Step3(output2)
	// fmt.Println("output3", output3)
	for message := range output3 {
		fmt.Println(message)
	}
	// time.Sleep(time.Second * 3)

	// task-4-4-2
	fmt.Printf("\ntask-4-4-3\n")

	tasks := make(chan string)
	numWorkers := 4

	// Запуск пула воркеров
	results := hw.StartWorkerPool(numWorkers, tasks)

	// Отправить задачи
	go func() {
		defer close(tasks)
		for i := 0; i < 10; i++ {
			tasks <- fmt.Sprintf("Task %d", i)
		}
	}()

	// Чтение и вывод результатов
	i := 0
	for result := range results {
		i++
		fmt.Printf("%-2d MD5: %s\n", i, result)
	}

	// fmt.Printf("\ntask-4-5-2\n")
	// hw.Run452()

	// fmt.Printf("\ntask-4-5-3\n")
	// hw.Run453()

	// fmt.Printf("\ntask-4-6-2\n")
	// hw.Run462()

	// fmt.Printf("\ntask-4-7-2\n")
	// hw.Run472()

	fmt.Printf("\ntask-4-8-1\n")
	hw.Run481()

}
