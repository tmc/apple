package vzkit

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objc"
)

// ErrorCompletionHandler captures an optional NSError from an Objective-C
// completion handler callback (void (^)(NSError *)).
//
// Usage:
//
//	h := vmkit.NewErrorCompletionHandler()
//	objcMethod(h.Block())
//	h.Wait()
//	if err := h.Error(); err != nil { ... }
type ErrorCompletionHandler struct {
	sema   dispatch.Semaphore
	block  objc.Block
	done   atomic.Bool
	mu     sync.Mutex
	errMsg string
}

// NewErrorCompletionHandler creates a completion handler for error-only callbacks.
func NewErrorCompletionHandler() *ErrorCompletionHandler {
	h := &ErrorCompletionHandler{
		sema: dispatch.SemaphoreCreate(0),
	}
	h.block = objc.NewBlock(func(_ objc.Block, err objc.ID) {
		if err != 0 {
			msg := ExtractNSErrorMessage(err)
			h.mu.Lock()
			h.errMsg = msg
			h.mu.Unlock()
		}
		h.done.Store(true)
		h.sema.Signal()
	})
	return h
}

// Block returns the underlying objc.Block for passing to Objective-C methods.
func (h *ErrorCompletionHandler) Block() objc.Block {
	return h.block
}

// Wait blocks until the completion handler is called.
func (h *ErrorCompletionHandler) Wait() {
	h.sema.Wait(dispatch.TimeForever)
}

// Done reports whether the completion handler has been called.
func (h *ErrorCompletionHandler) Done() bool {
	return h.done.Load()
}

// Error returns the error if one was received, or nil.
func (h *ErrorCompletionHandler) Error() error {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.errMsg == "" {
		return nil
	}
	return fmt.Errorf("%s", h.errMsg)
}
