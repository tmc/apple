// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VTFrameProcessor] class.
var (
	_VTFrameProcessorClass     VTFrameProcessorClass
	_VTFrameProcessorClassOnce sync.Once
)

func getVTFrameProcessorClass() VTFrameProcessorClass {
	_VTFrameProcessorClassOnce.Do(func() {
		_VTFrameProcessorClass = VTFrameProcessorClass{class: objc.GetClass("VTFrameProcessor")}
	})
	return _VTFrameProcessorClass
}

// GetVTFrameProcessorClass returns the class object for VTFrameProcessor.
func GetVTFrameProcessorClass() VTFrameProcessorClass {
	return getVTFrameProcessorClass()
}

type VTFrameProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VTFrameProcessorClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VTFrameProcessorClass) Alloc() VTFrameProcessor {
	rv := objc.Send[VTFrameProcessor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A class that creates a new frame processor for the configured video effect.
//
// # Overview
//
// Use this class to perform frame by frame processing on your video. Start by
// specifying a video effect by passing a [VTFrameProcessorConfiguration]
// object to the [VTFrameProcessor.StartSessionWithConfigurationError] call.
// Once the session is created,
// [VTFrameProcessor.ProcessWithParametersCompletionHandler] is called in a
// loop to process your video’s frames one at a time. Once all the frames
// are processed, call an [VTFrameProcessor.EndSession] to finish all pending
// processing.
//
// For successful processing, the caller needs to ensure that all buffers
// passed to the processWithParameters interface are unmodified (including
// attachments) until the function returns or the callback is received in the
// case of asynchronous mode.
//
// # Processing frames
//
//   - [VTFrameProcessor.StartSessionWithConfigurationError]: Starts a new session and configures the processor pipeline.
//   - [VTFrameProcessor.ProcessWithParametersCompletionHandler]: Asynchronously performs the video effect specified in the start session.
//   - [VTFrameProcessor.ProcessWithCommandBufferParameters]: Asynchronously performs the video effect specified in the start session specifically for Metal.
//   - [VTFrameProcessor.EndSession]: Performs all necessary tasks to end the session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessor
type VTFrameProcessor struct {
	objectivec.Object
}

// VTFrameProcessorFromID constructs a [VTFrameProcessor] from an objc.ID.
//
// A class that creates a new frame processor for the configured video effect.
func VTFrameProcessorFromID(id objc.ID) VTFrameProcessor {
	return VTFrameProcessor{objectivec.Object{ID: id}}
}

// NOTE: VTFrameProcessor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VTFrameProcessor] class.
//
// # Processing frames
//
//   - [IVTFrameProcessor.StartSessionWithConfigurationError]: Starts a new session and configures the processor pipeline.
//   - [IVTFrameProcessor.ProcessWithParametersCompletionHandler]: Asynchronously performs the video effect specified in the start session.
//   - [IVTFrameProcessor.ProcessWithCommandBufferParameters]: Asynchronously performs the video effect specified in the start session specifically for Metal.
//   - [IVTFrameProcessor.EndSession]: Performs all necessary tasks to end the session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessor
type IVTFrameProcessor interface {
	objectivec.IObject

	// Topic: Processing frames

	// Starts a new session and configures the processor pipeline.
	StartSessionWithConfigurationError(configuration VTFrameProcessorConfiguration) (bool, error)
	// Asynchronously performs the video effect specified in the start session.
	ProcessWithParametersCompletionHandler(parameters VTFrameProcessorParameters, completionHandler VTFrameProcessorParametersErrorHandler)
	// Asynchronously performs the video effect specified in the start session specifically for Metal.
	ProcessWithCommandBufferParameters(commandBuffer metal.MTLCommandBuffer, parameters VTFrameProcessorParameters)
	// Performs all necessary tasks to end the session.
	EndSession()

	// Synchronously performs the configured video effect.
	ProcessWithParametersError(parameters VTFrameProcessorParameters) (bool, error)
}

// Init initializes the instance.
func (v VTFrameProcessor) Init() VTFrameProcessor {
	rv := objc.Send[VTFrameProcessor](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VTFrameProcessor) Autorelease() VTFrameProcessor {
	rv := objc.Send[VTFrameProcessor](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVTFrameProcessor creates a new VTFrameProcessor instance.
func NewVTFrameProcessor() VTFrameProcessor {
	class := getVTFrameProcessorClass()
	rv := objc.Send[VTFrameProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Starts a new session and configures the processor pipeline.
//
// configuration: A configuration object for the video effect that will be applied in the
// subsequent processing calls.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessor/startSession(configuration:)
func (v VTFrameProcessor) StartSessionWithConfigurationError(configuration VTFrameProcessorConfiguration) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("startSessionWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("startSessionWithConfiguration:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Asynchronously performs the video effect specified in the start session.
//
// parameters: A [VTFrameProcessorParameters] object to specify additional parameters to
// use during processing. It needs to match the configuration type used during
// start session.
//
// completionHandler: This completion handler is called when the frame processing is completed.
// The completion handler will receive the same parameters object that was
// provided to the original call, as well as an NSError which will contain an
// error code if processing was not successful.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessor/process(parameters:completionHandler:)
func (v VTFrameProcessor) ProcessWithParametersCompletionHandler(parameters VTFrameProcessorParameters, completionHandler VTFrameProcessorParametersErrorHandler) {
	_block1, _ := NewVTFrameProcessorParametersErrorBlock(completionHandler)
	objc.Send[objc.ID](v.ID, objc.Sel("processWithParameters:completionHandler:"), parameters, _block1)
}

// Asynchronously performs the video effect specified in the start session
// specifically for Metal.
//
// commandBuffer: An existing Metal command buffer where the frame processing will be
// inserted.
//
// parameters: A VTFrameProcessorParameters based object to specify additional frame based
// parameters to be used during processing. It needs to match the
// configuration type used during start session.
//
// # Discussion
//
// This function allows you to add the effect to an existing Metal command
// buffer. This can be used by clients that have an existing Metal pipeline
// and want to add this effect to it.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessor/process(with:parameters:)
func (v VTFrameProcessor) ProcessWithCommandBufferParameters(commandBuffer metal.MTLCommandBuffer, parameters VTFrameProcessorParameters) {
	objc.Send[objc.ID](v.ID, objc.Sel("processWithCommandBuffer:parameters:"), commandBuffer, parameters)
}

// Performs all necessary tasks to end the session.
//
// # Discussion
//
// After this call completes, no new frames can be processed unless
// [VTFrameProcessor.StartSessionWithConfigurationError] is called again.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessor/endSession()
func (v VTFrameProcessor) EndSession() {
	objc.Send[objc.ID](v.ID, objc.Sel("endSession"))
}

// Synchronously performs the configured video effect.
//
// parameters: Frame processing parameters to specify additional frame based parameters to
// be used during processing. The parameters need to match the configuration
// type used during start session.
//
// # Discussion
//
// Frame level settings and frame level input/output parameters are passed by
// using the respective [VTFrameProcessorParameters] for the effect that the
// frame processor is configured for.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessor/processWithParameters:error:
func (v VTFrameProcessor) ProcessWithParametersError(parameters VTFrameProcessorParameters) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("processWithParameters:error:"), parameters, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("processWithParameters:error: returned NO with nil NSError")
	}
	return rv, nil

}

// ProcessWithParameters is a synchronous wrapper around [VTFrameProcessor.ProcessWithParametersCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VTFrameProcessor) ProcessWithParameters(ctx context.Context, parameters VTFrameProcessorParameters) (VTFrameProcessorParameters, error) {
	type result struct {
		val VTFrameProcessorParameters
		err error
	}
	done := make(chan result, 1)
	v.ProcessWithParametersCompletionHandler(parameters, func(val VTFrameProcessorParameters, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
