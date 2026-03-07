package vzkit

import (
	"github.com/tmc/apple/dispatch"
)

// Queue wraps a dispatch queue for VM operations.
// The Virtualization framework requires VM operations to happen on a specific queue.
type Queue struct {
	q dispatch.Queue
}

// NewQueue creates a serial dispatch queue with the given label.
func NewQueue(label string) *Queue {
	return &Queue{q: dispatch.QueueCreate(label)}
}

// WrapQueue creates a Queue from an existing dispatch.Queue.
func WrapQueue(q dispatch.Queue) *Queue {
	return &Queue{q: q}
}

// Queue returns the underlying dispatch.Queue.
func (q *Queue) Queue() dispatch.Queue {
	return q.q
}

// Sync executes work synchronously on the dispatch queue.
func (q *Queue) Sync(work func()) {
	q.q.Sync(work)
}

// Async executes work asynchronously on the dispatch queue.
func (q *Queue) Async(work func()) {
	q.q.Async(work)
}
