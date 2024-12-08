package homework

import (
	"testing"
)

// const N = 5
const TOTAL = 100

func TestSemaphore(t *testing.T) {
	sem := NewSemaphore(N)
	done := make(chan bool)
	for i := 1; i <= TOTAL; i++ {
		sem.Acquire(3)
		go func(v int) {
			defer sem.Release(1)
			defer sem.Release(2)
			if v == TOTAL {
				done <- true
			}
		}(i)
	}
	<-done
}
