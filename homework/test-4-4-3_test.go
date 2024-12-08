package homework

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func hashHex(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func TestWorkerPool(t *testing.T) {
	tasks := make(chan string)
	out := make(chan string, 100)

	var wg sync.WaitGroup
	var workers []*Worker

	for i := 0; i < 4; i++ {
		wg.Add(1)
		worker := NewWorker(tasks, &wg, out)
		go worker.Run()
		workers = append(workers, worker)
	}

	results := make(map[string]string)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("task-%d", i)
		tasks <- key
		results[hashHex(key)] = key
	}

	close(tasks)

	wg.Wait()
	close(out)

	for _, w := range workers {
		for val := range w.out {
			if _, ok := results[val]; ok {
				delete(results, val)
				continue
			}
			t.Errorf("hash %s was not expected (task-0...task-99)", val)
		}
	}

	require.Equal(t, 0, len(results), fmt.Sprintf("Found orphan unchecked records: %s", results))
}
