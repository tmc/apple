//go:build darwin

package espresso

import (
	"context"
	"fmt"
	"time"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/private/espresso"
)

// FrameStorage holds reusable multi-frame storage for batch or repeated evaluation.
type FrameStorage struct {
	storage espresso.EspressoDataFrameStorage
}

// NewFrameStorage creates reusable storage from the given frames.
func NewFrameStorage(frames ...*Frame) *FrameStorage {
	storage := espresso.NewEspressoDataFrameStorage()
	arr := foundation.NewMutableArrayWithCapacity(uint(len(frames)))
	for _, f := range frames {
		arr.AddObject(f.df)
	}
	storage.SetDataFrames(arr)
	return &FrameStorage{storage: storage}
}

// Eval executes the network on the given frame synchronously.
// Inputs are bound from the frame, execution runs, and outputs are written back.
func (n *Network) Eval(f *Frame) error {
	return n.EvalWithOptions(f, EvalOptions{})
}

// EvalOptions configures execution behavior.
type EvalOptions struct {
	UseCVPixelBuffers bool          // use CoreVideo pixel buffers for I/O
	Timeout           time.Duration // execution timeout (0 = 30s default)
}

// EvalAsync executes the network asynchronously and returns a channel
// that receives the error (or nil) when execution completes.
func (n *Network) EvalAsync(f *Frame) <-chan error {
	ch := make(chan error, 1)
	go func() {
		ch <- n.Eval(f)
	}()
	return ch
}

// EvalAsyncWithCallback executes the network asynchronously and calls fn
// with the result when execution completes.
//
// Note: This uses a goroutine internally because the underlying ObjC
// completion block binding cleans up the block pointer on return,
// making a true async block callback unsafe. The goroutine runs the
// synchronous eval path and calls fn when done.
func (n *Network) EvalAsyncWithCallback(f *Frame, fn func(error)) {
	go func() {
		fn(n.Eval(f))
	}()
}

// EvalWithStorage executes the network using pre-allocated storage.
// Each frame in the storage has its inputs bound before execution and
// outputs bound after execution.
func (n *Network) EvalWithStorage(fs *FrameStorage, opts EvalOptions) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	if n.closed {
		return ErrClosed
	}

	if opts.UseCVPixelBuffers {
		n.binder.SetUse_cvpixelbuffer(1)
	}

	// Bind inputs for each frame in storage.
	nFrames := fs.storage.NumberOfDataFrames()
	for i := range nFrames {
		obj := fs.storage.DataFrameAtIndex(i)
		if obj == nil || obj.GetID() == 0 {
			continue
		}
		df := espresso.EspressoDataFrameFromID(obj.GetID())
		rc := n.binder.BindInputsFromFrameToNetwork(df, n.net)
		if rc != 0 {
			return fmt.Errorf("%w: bind returned %d for frame %d", ErrBindInput, rc, i)
		}
	}

	timeout := opts.Timeout
	if timeout == 0 {
		timeout = 30 * time.Second
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := n.storExec.ExecuteDataFrameStorageWithNetworkBlockSync(ctx, fs.storage, n.net)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrEval, err)
	}

	// Bind outputs for each frame in storage.
	for i := range nFrames {
		obj := fs.storage.DataFrameAtIndex(i)
		if obj == nil || obj.GetID() == 0 {
			continue
		}
		df := espresso.EspressoDataFrameFromID(obj.GetID())
		rc := n.binder.BindOutputsFromFrameToNetwork(df, n.net)
		if rc != 0 {
			return fmt.Errorf("%w: bind returned %d for frame %d", ErrBindOutput, rc, i)
		}
	}

	return nil
}

// EvalWithOptions executes with custom options.
func (n *Network) EvalWithOptions(f *Frame, opts EvalOptions) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	if n.closed {
		return ErrClosed
	}

	if opts.UseCVPixelBuffers {
		n.binder.SetUse_cvpixelbuffer(1)
	}

	// Bind inputs from frame to network.
	rc := n.binder.BindInputsFromFrameToNetwork(f.df, n.net)
	if rc != 0 {
		return fmt.Errorf("%w: bind returned %d", ErrBindInput, rc)
	}

	// Create a single-frame storage for the storage executor.
	storage := espresso.NewEspressoDataFrameStorage()
	frames := foundation.NewArrayWithObject(f.df)
	storage.SetDataFrames(frames)

	// Execute with timeout.
	timeout := opts.Timeout
	if timeout == 0 {
		timeout = 30 * time.Second
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := n.storExec.ExecuteDataFrameStorageWithNetworkBlockSync(ctx, storage, n.net)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrEval, err)
	}

	// Bind outputs from frame to network after execution.
	rc = n.binder.BindOutputsFromFrameToNetwork(f.df, n.net)
	if rc != 0 {
		return fmt.Errorf("%w: bind returned %d", ErrBindOutput, rc)
	}

	return nil
}
