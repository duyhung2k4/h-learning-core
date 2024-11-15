package main

import (
	"app/config"
	"app/job"
	"app/queue"
	"app/router"
	"log"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		server := &http.Server{
			Addr:           ":" + config.GetAppPort(),
			Handler:        router.AppRouter(),
			MaxHeaderBytes: 1 << 20,
		}

		log.Fatalln(server.ListenAndServe())
	}()

	// queue
	go func() {
		defer wg.Done()
		queue.InitQueue()
	}()

	// Chan job
	go func() {
		defer wg.Done()
		job.InitJob()
	}()

	wg.Wait()
}
