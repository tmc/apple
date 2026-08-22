//go:build darwin

package ane

// EvalAsync launches a Go goroutine wrapper around synchronous Eval and
// returns a channel that receives the error result when evaluation completes.
// Note: This is a Go-side background goroutine wrapper, not an asynchronous
// hardware dispatch or polling API.
func (m *Model) EvalAsync() <-chan error {
	ch := make(chan error, 1)
	go func() {
		ch <- m.Eval()
	}()
	return ch
}

// EvalAsyncWithCallback launches a Go goroutine wrapper around synchronous Eval
// and calls fn on completion from an arbitrary goroutine.
func (m *Model) EvalAsyncWithCallback(fn func(error)) {
	go func() {
		fn(m.Eval())
	}()
}
