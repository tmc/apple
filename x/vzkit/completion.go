package vzkit

import (
	"fmt"

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
	errMsg string
	block  objc.Block
	done   bool
}

// NewErrorCompletionHandler creates a completion handler for error-only callbacks.
func NewErrorCompletionHandler() *ErrorCompletionHandler {
	h := &ErrorCompletionHandler{
		sema: dispatch.SemaphoreCreate(0),
	}
	h.block = objc.NewBlock(func(_ objc.Block, err objc.ID) {
		if err != 0 {
			h.errMsg = ExtractNSErrorMessage(err)
		}
		h.done = true
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
	return h.done
}

// Error returns the error if one was received, or nil.
func (h *ErrorCompletionHandler) Error() error {
	if h.errMsg == "" {
		return nil
	}
	return fmt.Errorf("%s", h.errMsg)
}
