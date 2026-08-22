//go:build darwin

package engineattest

import (
	"fmt"
	"sync"

	"github.com/tmc/apple/metal"
)

// Queue wraps a Metal command queue so that the command buffers it vends
// are tracked and can be attested by [Queue.GPU]. Code under test uses
// the Queue exactly like the [metal.MTLCommandQueue] it wraps.
type Queue struct {
	q metal.MTLCommandQueue

	mu     sync.Mutex
	vended []metal.MTLCommandBuffer
}

// NewQueue wraps q. The zero Queue is not usable.
func NewQueue(q metal.MTLCommandQueue) *Queue {
	return &Queue{q: q}
}

// CommandBuffer vends a tracked command buffer from the wrapped queue.
func (q *Queue) CommandBuffer() metal.MTLCommandBuffer {
	cb := q.q.CommandBuffer()
	q.mu.Lock()
	q.vended = append(q.vended, cb)
	q.mu.Unlock()
	return cb
}

// Underlying returns the wrapped queue. Work submitted through it
// directly is invisible to [Queue.GPU].
func (q *Queue) Underlying() metal.MTLCommandQueue { return q.q }

// GPU runs fn and returns an error unless at least one command buffer
// vended by this Queue during fn completed on the GPU with a nonzero
// scheduling window (GPUEndTime > GPUStartTime > 0). Buffers that were
// committed but not yet finished are waited on before judging.
//
// If fn returns an error, GPU returns it unchanged without judging the
// claim.
func (q *Queue) GPU(fn func() error) error {
	q.mu.Lock()
	start := len(q.vended)
	q.mu.Unlock()

	if err := fn(); err != nil {
		return err
	}

	q.mu.Lock()
	during := q.vended[start:len(q.vended):len(q.vended)]
	q.mu.Unlock()

	if len(during) == 0 {
		return fmt.Errorf("engineattest: no command buffer was created during the attested region: %w", ErrDidNotRun)
	}
	committed, completed := 0, 0
	for _, cb := range during {
		if cb.Status() < metal.MTLCommandBufferStatusCommitted {
			continue
		}
		committed++
		cb.WaitUntilCompleted()
		if cb.Status() != metal.MTLCommandBufferStatusCompleted {
			continue
		}
		completed++
		if end, start := cb.GPUEndTime(), cb.GPUStartTime(); start > 0 && end > start {
			return nil
		}
	}
	switch {
	case committed == 0:
		return fmt.Errorf("engineattest: %d command buffer(s) created but none committed: %w", len(during), ErrDidNotRun)
	case completed == 0:
		return fmt.Errorf("engineattest: %d command buffer(s) committed but none completed cleanly: %w", committed, ErrDidNotRun)
	default:
		return fmt.Errorf("engineattest: %d command buffer(s) completed but none reported GPU time: %w", completed, ErrUnattestable)
	}
}
