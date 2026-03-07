package main

import (
	"fmt"
	"sync"

	"github.com/tmc/apple/dispatch"
)

func main() {
	// Create a serial queue and dispatch work.
	queue := dispatch.QueueCreate("com.example.serial")
	var mu sync.Mutex
	var results []string

	group := dispatch.GroupCreate()
	for i := range 3 {
		i := i
		group.Async(queue, func() {
			mu.Lock()
			results = append(results, fmt.Sprintf("task-%d", i))
			mu.Unlock()
		})
	}
	group.Wait(dispatch.TimeForever)

	fmt.Printf("queue=%s\n", queue.Label())
	fmt.Printf("tasks=%d\n", len(results))

	// Use a semaphore as a signal.
	sema := dispatch.SemaphoreCreate(0)
	queue.Async(func() {
		sema.Signal()
	})
	ok := sema.Wait(dispatch.TimeForever)
	fmt.Printf("signal=%t\n", ok)
}
