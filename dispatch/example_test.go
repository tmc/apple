//go:build darwin

package dispatch_test

import (
	"fmt"

	"github.com/tmc/apple/dispatch"
)

func ExampleQueue_Async() {
	queue := dispatch.QueueCreate("com.example.queue")
	fmt.Println("Queue label:", queue.Label())

	done := make(chan int)
	queue.Async(func() {
		done <- 42
	})

	fmt.Println("Async result:", <-done)

	// Output:
	// Queue label: com.example.queue
	// Async result: 42
}

func ExampleQueueCreateConcurrent() {
	queue := dispatch.QueueCreateConcurrent("com.example.concurrent")
	fmt.Println("Queue label:", queue.Label())

	done := make(chan string)
	queue.Async(func() {
		done <- "concurrent task executed"
	})

	fmt.Println(<-done)

	// Output:
	// Queue label: com.example.concurrent
	// concurrent task executed
}

func ExampleTimeFromNow() {
	var val int
	q := dispatch.QueueCreate("com.example.timefromnow")
	item := dispatch.WorkItemCreate(0, func() {
		val = 100
	})
	defer item.Release()
	q.AsyncWorkItem(item)
	item.Wait(dispatch.TimeFromNow(1000000000))
	fmt.Println("Value after wait:", val)

	// Output:
	// Value after wait: 100
}
