// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVSampleBufferGenerator] class.
var (
	_AVSampleBufferGeneratorClass     AVSampleBufferGeneratorClass
	_AVSampleBufferGeneratorClassOnce sync.Once
)

func getAVSampleBufferGeneratorClass() AVSampleBufferGeneratorClass {
	_AVSampleBufferGeneratorClassOnce.Do(func() {
		_AVSampleBufferGeneratorClass = AVSampleBufferGeneratorClass{class: objc.GetClass("AVSampleBufferGenerator")}
	})
	return _AVSampleBufferGeneratorClass
}

// GetAVSampleBufferGeneratorClass returns the class object for AVSampleBufferGenerator.
func GetAVSampleBufferGeneratorClass() AVSampleBufferGeneratorClass {
	return getAVSampleBufferGeneratorClass()
}

type AVSampleBufferGeneratorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSampleBufferGeneratorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSampleBufferGeneratorClass) Alloc() AVSampleBufferGenerator {
	rv := objc.Send[AVSampleBufferGenerator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that creates sample buffers.
//
// # Overview
// 
// Each request for [CMSampleBuffer] creation is described in an
// [AVSampleBufferRequest] object. The [CMSampleBuffer] opaque objects are
// returned synchronously. If requested, sample data may be loaded
// asynchronously (depending on file format support).
//
// [CMSampleBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer
//
// # Creating sample buffer generators
//
//   - [AVSampleBufferGenerator.InitWithAssetTimebase]: Creates a new sample buffer generator.
//
// # Creating a sample buffer
//
//   - [AVSampleBufferGenerator.CreateSampleBufferForRequestError]: Creates a sample buffer, and attempts to load its data asynchronously if requested.
//   - [AVSampleBufferGenerator.MakeBatch]: Creates a batch object to handle generating multiple sample buffers.
//   - [AVSampleBufferGenerator.CreateSampleBufferForRequestAddingToBatchError]: Creates a sample buffer and attempts to defer I/O for its data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGenerator
type AVSampleBufferGenerator struct {
	objectivec.Object
}

// AVSampleBufferGeneratorFromID constructs a [AVSampleBufferGenerator] from an objc.ID.
//
// An object that creates sample buffers.
func AVSampleBufferGeneratorFromID(id objc.ID) AVSampleBufferGenerator {
	return AVSampleBufferGenerator{objectivec.Object{ID: id}}
}
// NOTE: AVSampleBufferGenerator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVSampleBufferGenerator] class.
//
// # Creating sample buffer generators
//
//   - [IAVSampleBufferGenerator.InitWithAssetTimebase]: Creates a new sample buffer generator.
//
// # Creating a sample buffer
//
//   - [IAVSampleBufferGenerator.CreateSampleBufferForRequestError]: Creates a sample buffer, and attempts to load its data asynchronously if requested.
//   - [IAVSampleBufferGenerator.MakeBatch]: Creates a batch object to handle generating multiple sample buffers.
//   - [IAVSampleBufferGenerator.CreateSampleBufferForRequestAddingToBatchError]: Creates a sample buffer and attempts to defer I/O for its data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGenerator
type IAVSampleBufferGenerator interface {
	objectivec.IObject

	// Topic: Creating sample buffer generators

	// Creates a new sample buffer generator.
	InitWithAssetTimebase(asset IAVAsset, timebase uintptr) AVSampleBufferGenerator

	// Topic: Creating a sample buffer

	// Creates a sample buffer, and attempts to load its data asynchronously if requested.
	CreateSampleBufferForRequestError(request IAVSampleBufferRequest) (uintptr, error)
	// Creates a batch object to handle generating multiple sample buffers.
	MakeBatch() IAVSampleBufferGeneratorBatch
	// Creates a sample buffer and attempts to defer I/O for its data.
	CreateSampleBufferForRequestAddingToBatchError(request IAVSampleBufferRequest, batch IAVSampleBufferGeneratorBatch) (uintptr, error)
}

// Init initializes the instance.
func (s AVSampleBufferGenerator) Init() AVSampleBufferGenerator {
	rv := objc.Send[AVSampleBufferGenerator](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSampleBufferGenerator) Autorelease() AVSampleBufferGenerator {
	rv := objc.Send[AVSampleBufferGenerator](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSampleBufferGenerator creates a new AVSampleBufferGenerator instance.
func NewAVSampleBufferGenerator() AVSampleBufferGenerator {
	class := getAVSampleBufferGeneratorClass()
	rv := objc.Send[AVSampleBufferGenerator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new sample buffer generator.
//
// asset: The asset.
//
// timebase: If [NULL], requests will be handled synchronously.
//
// # Return Value
// 
// An initialized [AVSampleBufferGenerator] instance.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGenerator/init(asset:timebase:)
func NewSampleBufferGeneratorWithAssetTimebase(asset IAVAsset, timebase uintptr) AVSampleBufferGenerator {
	instance := getAVSampleBufferGeneratorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAsset:timebase:"), asset, timebase)
	return AVSampleBufferGeneratorFromID(rv)
}

// Creates a new sample buffer generator.
//
// asset: The asset.
//
// timebase: If [NULL], requests will be handled synchronously.
//
// # Return Value
// 
// An initialized [AVSampleBufferGenerator] instance.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGenerator/init(asset:timebase:)
func (s AVSampleBufferGenerator) InitWithAssetTimebase(asset IAVAsset, timebase uintptr) AVSampleBufferGenerator {
	rv := objc.Send[AVSampleBufferGenerator](s.ID, objc.Sel("initWithAsset:timebase:"), asset, timebase)
	return rv
}
// Creates a sample buffer, and attempts to load its data asynchronously if
// requested.
//
// request: A sample buffer creation request.
//
// # Return Value
// 
// A sample buffer object.
//
// # Discussion
// 
// If you created the generator with a `nil` timebase, any associated
// [AVSampleBufferRequest] objects default to using a request mode of
// [SampleBufferRequestModeImmediate].
// 
// Call the [NotifyOfDataReadyForSampleBufferCompletionHandler] class method
// to have the system notify you when sample buffer data is available.
// 
// The request may fail based on generator configuration or file format.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGenerator/makeSampleBuffer(for:)
func (s AVSampleBufferGenerator) CreateSampleBufferForRequestError(request IAVSampleBufferRequest) (uintptr, error) {
	var errorPtr objc.ID
	rv := objc.Send[uintptr](s.ID, objc.Sel("createSampleBufferForRequest:error:"), request, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}
// Creates a batch object to handle generating multiple sample buffers.
//
// # Return Value
// 
// An object to batch generate sample buffers.
//
// # Discussion
// 
// Generating sample buffers in batches optimizes performance by allowing the
// system to asynchronously load sample data and optimize I/O when possible.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGenerator/makeBatch()
func (s AVSampleBufferGenerator) MakeBatch() IAVSampleBufferGeneratorBatch {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeBatch"))
	return AVSampleBufferGeneratorBatchFromID(rv)
}
// Creates a sample buffer and attempts to defer I/O for its data.
//
// request: A sample buffer creation request.
//
// batch: A batch object to contain the output sample buffer. You must create this
// object by calling [Batch] on the same instance of [AVSampleBufferGenerator]
// or an error occurs.
//
// # Return Value
// 
// A sample buffer.
//
// # Discussion
// 
// Call the [DataReadyWithCompletionHandler] on [AVSampleBufferGeneratorBatch]
// once to commence I/O and load sample data for all [CMSampleBuffer] objects
// in a batch. After loading commences, any subsequent calls to
// [CreateSampleBufferForRequestAddingToBatchError] throw an exception.
// 
// The generator may defer I/O to fetch sample data depending on the source of
// the sample data and the generator’s timebase
// 
// The request may fail based on generator configuration or file format.
//
// [CMSampleBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGenerator/makeSampleBuffer(for:addTo:)
func (s AVSampleBufferGenerator) CreateSampleBufferForRequestAddingToBatchError(request IAVSampleBufferRequest, batch IAVSampleBufferGeneratorBatch) (uintptr, error) {
	var errorPtr objc.ID
	rv := objc.Send[uintptr](s.ID, objc.Sel("createSampleBufferForRequest:addingToBatch:error:"), request, batch, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// Notifies the sample buffer generator when data is ready for the sample
// buffer reference or an error has occurred.
//
// sbuf: The [CMSampleBufferRef].
//
// completionHandler: A completion block that is called when data is ready for the sample buffer
// or an error occurs. The `dataReady` argument is [true] if data is read for
// the sample buffer. If an error occurs, the `error` argument contains the
// [NSError] object.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferGenerator/notifyOfDataReady(for:completionHandler:)
func (_AVSampleBufferGeneratorClass AVSampleBufferGeneratorClass) NotifyOfDataReadyForSampleBufferCompletionHandler(sbuf uintptr, completionHandler BoolErrorHandler) {
_block1, _ := NewBoolErrorBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_AVSampleBufferGeneratorClass.class), objc.Sel("notifyOfDataReadyForSampleBuffer:completionHandler:"), sbuf, _block1)
}

// NotifyOfDataReadyForSampleBuffer is a synchronous wrapper around [AVSampleBufferGenerator.NotifyOfDataReadyForSampleBufferCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (sc AVSampleBufferGeneratorClass) NotifyOfDataReadyForSampleBuffer(ctx context.Context, sbuf uintptr) (bool, error) {
	type result struct {
		val bool
		err error
	}
	done := make(chan result, 1)
	sc.NotifyOfDataReadyForSampleBufferCompletionHandler(sbuf, func(val bool, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

