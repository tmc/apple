//go:build darwin

// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch_test

import (
	"fmt"

	"github.com/tmc/apple/dispatch"
)

func ExampleQueue_Sync() {
	q := dispatch.QueueCreate("com.appledocs.dispatch.example.sync")
	value := 0
	q.Sync(func() { value = 42 })
	fmt.Println(value)
	// Output: 42
}

func ExampleSemaphore() {
	s := dispatch.SemaphoreCreate(0)
	s.Signal()
	fmt.Println(s.Wait(dispatch.TimeNow))
	// Output: true
}

func ExampleGroup() {
	g := dispatch.GroupCreate()
	g.Enter()
	g.Leave()
	fmt.Println(g.Wait(dispatch.TimeNow))
	// Output: true
}
