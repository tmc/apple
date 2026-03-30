// Command queues demonstrates a serial queue, a group, and a semaphore.
package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/tmc/apple/dispatch"
)

func main() {
	// Create a serial queue and dispatch work.
	queue := dispatch.QueueCreate("com.example.serial")
	var mu sync.Mutex
	var order []string

	group := dispatch.GroupCreate()
	for i := 0; i < 3; i++ {
		i := i
		group.Async(queue, func() {
			mu.Lock()
			order = append(order, fmt.Sprintf("task-%d", i))
			mu.Unlock()
		})
	}
	group.Wait(dispatch.TimeForever)

	fmt.Printf("queue=%s\n", queue.Label())
	fmt.Printf("order=%s\n", strings.Join(order, ", "))

	// Use a semaphore as a signal.
	signal := dispatch.SemaphoreCreate(0)
	queue.Async(func() {
		signal.Signal()
	})
	ok := signal.Wait(dispatch.TimeForever)
	fmt.Printf("signal-received=%t\n", ok)
}
