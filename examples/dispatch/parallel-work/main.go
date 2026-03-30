// Parallel-work dispatches concurrent tasks using GCD groups and semaphores.
package main

import (
	"fmt"
	"math"
	"sync"

	"github.com/tmc/apple/dispatch"
)

func main() {
	// Use a concurrent queue for parallel work.
	queue := dispatch.QueueCreateConcurrent("com.example.parallel")
	group := dispatch.GroupCreate()

	var mu sync.Mutex
	results := make(map[int]float64)

	// Dispatch several tasks that compute square roots in parallel.
	const n = 8
	for i := 0; i < n; i++ {
		i := i
		group.Async(queue, func() {
			val := math.Sqrt(float64((i + 1) * 1000))
			mu.Lock()
			results[i] = val
			mu.Unlock()
		})
	}
	group.Wait(dispatch.TimeForever)

	fmt.Printf("Computed %d results on concurrent queue %q:\n", len(results), queue.Label())
	for i := 0; i < n; i++ {
		fmt.Printf("  sqrt(%d) = %.4f\n", (i+1)*1000, results[i])
	}

	// Use group.Notify for completion notification.
	group2 := dispatch.GroupCreate()
	serial := dispatch.QueueCreate("com.example.serial")
	done := make(chan struct{})

	for i := 0; i < 4; i++ {
		i := i
		group2.Async(queue, func() {
			_ = i * i
		})
	}
	group2.Notify(serial, func() {
		fmt.Println("All group2 tasks completed (via Notify).")
		close(done)
	})
	<-done

	// Demonstrate semaphore for resource limiting.
	sema := dispatch.SemaphoreCreate(2)
	group3 := dispatch.GroupCreate()
	var order []int

	for i := 0; i < 6; i++ {
		i := i
		group3.Async(queue, func() {
			sema.Wait(dispatch.TimeForever)
			mu.Lock()
			order = append(order, i)
			mu.Unlock()
			sema.Signal()
		})
	}
	group3.Wait(dispatch.TimeForever)
	fmt.Printf("Semaphore-limited tasks completed: %d\n", len(order))
}
