package job

import (
	"sync"
)

func InitJob() {
	var wg sync.WaitGroup
	wg.Add(1)

	emailJob := NewEmailJob()
	go func() {
		defer wg.Done()
		emailJob.handle()
	}()

	wg.Wait()
}
