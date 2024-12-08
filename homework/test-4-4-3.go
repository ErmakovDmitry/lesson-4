package homework

import (
	"crypto/md5"
	"encoding/hex"
	"sync"
)

// Worker represents a worker that can process tasks.
type Worker struct {
	// Channel to receive tasks.
	tasks <-chan string
	// WaitGroup to signal when the worker is done.
	wg *sync.WaitGroup
	// Channel to write results
	out chan string
}

// NewWorker creates a new worker.
func NewWorker(tasks <-chan string, wg *sync.WaitGroup, out chan string) *Worker {
	return &Worker{
		tasks: tasks,
		wg:    wg,
		out:   out,
	}
}

// Run starts the worker.
func (w *Worker) Run() {
	defer w.wg.Done()
	for task := range w.tasks {
		hash := md5.Sum([]byte(task))
		w.out <- hex.EncodeToString(hash[:])
	}
}

// StartWorkerPool initializes and starts a pool of workers.
func StartWorkerPool(numWorkers int, tasks <-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		worker := NewWorker(tasks, &wg, out)
		go worker.Run()
	}

	// Закрыть канал out, когда все воркеры завершат выполнение
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
