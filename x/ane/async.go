//go:build darwin

package ane

// EvalAsync starts an asynchronous evaluation and returns a channel
// that receives the result when evaluation completes.
func (k *Kernel) EvalAsync() <-chan error {
	ch := make(chan error, 1)
	go func() {
		ch <- k.Eval()
	}()
	return ch
}

// EvalAsyncWithCallback starts an async evaluation and calls fn on completion.
// fn is called from an arbitrary goroutine.
func (k *Kernel) EvalAsyncWithCallback(fn func(error)) {
	go func() {
		fn(k.Eval())
	}()
}
