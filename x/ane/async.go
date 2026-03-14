//go:build darwin

package ane

// EvalAsync starts an asynchronous evaluation and returns a channel
// that receives the result when evaluation completes.
func (m *Model) EvalAsync() <-chan error {
	ch := make(chan error, 1)
	go func() {
		ch <- m.Eval()
	}()
	return ch
}

// EvalAsyncWithCallback starts an async evaluation and calls fn on completion.
// fn is called from an arbitrary goroutine.
func (m *Model) EvalAsyncWithCallback(fn func(error)) {
	go func() {
		fn(m.Eval())
	}()
}
