// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVSampleBufferGeneratorBatch] class.
var (
	_AVSampleBufferGeneratorBatchClass     AVSampleBufferGeneratorBatchClass
	_AVSampleBufferGeneratorBatchClassOnce sync.Once
)

func getAVSampleBufferGeneratorBatchClass() AVSampleBufferGeneratorBatchClass {
	_AVSampleBufferGeneratorBatchClassOnce.Do(func() {
		_AVSampleBufferGeneratorBatchClass = AVSampleBufferGeneratorBatchClass{class: objc.GetClass("AVSampleBufferGeneratorBatch")}
	})
	return _AVSampleBufferGeneratorBatchClass
}

// GetAVSampleBufferGeneratorBatchClass returns the class object for AVSampleBufferGeneratorBatch.
func GetAVSampleBufferGeneratorBatchClass() AVSampleBufferGeneratorBatchClass {
	return getAVSampleBufferGeneratorBatchClass()
}

type AVSampleBufferGeneratorBatchClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSampleBufferGeneratorBatchClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSampleBufferGeneratorBatchClass) Alloc() AVSampleBufferGeneratorBatch {
	rv := objc.Send[AVSampleBufferGeneratorBatch](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that generates sample buffers in a batch.
//
// # Overview
//
// The benefit of batching is it aggregates adjacent I/O requests and overlaps
// them when possible for all sample buffers within the batch.
//
// # Preparing a batch
//
//   - [AVSampleBufferGeneratorBatch.MakeDataReadyWithCompletionHandler]: Loads sample data asynchronously for all sample buffers within a batch.
//
// # Canceling a batch
//
//   - [AVSampleBufferGeneratorBatch.Cancel]: Cancels any I/O for this batch.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGeneratorBatch
type AVSampleBufferGeneratorBatch struct {
	objectivec.Object
}

// AVSampleBufferGeneratorBatchFromID constructs a [AVSampleBufferGeneratorBatch] from an objc.ID.
//
// An object that generates sample buffers in a batch.
func AVSampleBufferGeneratorBatchFromID(id objc.ID) AVSampleBufferGeneratorBatch {
	return AVSampleBufferGeneratorBatch{objectivec.Object{ID: id}}
}

// NOTE: AVSampleBufferGeneratorBatch adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVSampleBufferGeneratorBatch] class.
//
// # Preparing a batch
//
//   - [IAVSampleBufferGeneratorBatch.MakeDataReadyWithCompletionHandler]: Loads sample data asynchronously for all sample buffers within a batch.
//
// # Canceling a batch
//
//   - [IAVSampleBufferGeneratorBatch.Cancel]: Cancels any I/O for this batch.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGeneratorBatch
type IAVSampleBufferGeneratorBatch interface {
	objectivec.IObject

	// Topic: Preparing a batch

	// Loads sample data asynchronously for all sample buffers within a batch.
	MakeDataReadyWithCompletionHandler(completionHandler ErrorHandler)

	// Topic: Canceling a batch

	// Cancels any I/O for this batch.
	Cancel()
}

// Init initializes the instance.
func (s AVSampleBufferGeneratorBatch) Init() AVSampleBufferGeneratorBatch {
	rv := objc.Send[AVSampleBufferGeneratorBatch](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSampleBufferGeneratorBatch) Autorelease() AVSampleBufferGeneratorBatch {
	rv := objc.Send[AVSampleBufferGeneratorBatch](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSampleBufferGeneratorBatch creates a new AVSampleBufferGeneratorBatch instance.
func NewAVSampleBufferGeneratorBatch() AVSampleBufferGeneratorBatch {
	class := getAVSampleBufferGeneratorBatchClass()
	rv := objc.Send[AVSampleBufferGeneratorBatch](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Loads sample data asynchronously for all sample buffers within a batch.
//
// completionHandler: A callback the system invokes once when all sample buffers in the batch are
// data-ready, or when an error occurs.
//
// # Discussion
//
// Calling this method more than once on a batch generates an exception.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGeneratorBatch/makeDataReady(completionHandler:)
func (s AVSampleBufferGeneratorBatch) MakeDataReadyWithCompletionHandler(completionHandler ErrorHandler) {
	_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](s.ID, objc.Sel("makeDataReadyWithCompletionHandler:"), _block0)
}

// Cancels any I/O for this batch.
//
// # Discussion
//
// The system invokes the associated sample buffers data ready handlers with
// an error.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGeneratorBatch/cancel()
func (s AVSampleBufferGeneratorBatch) Cancel() {
	objc.Send[objc.ID](s.ID, objc.Sel("cancel"))
}

// MakeDataReady is a synchronous wrapper around [AVSampleBufferGeneratorBatch.MakeDataReadyWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s AVSampleBufferGeneratorBatch) MakeDataReady(ctx context.Context) error {
	done := make(chan error, 1)
	s.MakeDataReadyWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
