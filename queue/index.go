package queue

import (
	"sync"
)

func InitQueue() {
	queueUrlQuantity := NewQueueUrlQuantity()
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		queueUrlQuantity.Worker()
	}()

	wg.Wait()
}
