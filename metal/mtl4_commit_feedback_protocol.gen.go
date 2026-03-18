// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// Describes an object containing debug information from Metal to your app after completing a workload.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommitFeedback
type MTL4CommitFeedback interface {
	objectivec.IObject

	// A description of an error when the GPU encounters an issue as it runs the committed command buffers.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommitFeedback/error
	Error() foundation.INSError

	// The host time, in seconds, when the GPU finishes execution of the committed command buffers.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommitFeedback/gpuEndTime
	GPUEndTime() float64

	// The host time, in seconds, when the GPU starts execution of the committed command buffers.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommitFeedback/gpuStartTime
	GPUStartTime() float64
}

// MTL4CommitFeedbackObject wraps an existing Objective-C object that conforms to the MTL4CommitFeedback protocol.
type MTL4CommitFeedbackObject struct {
	objectivec.Object
}
func (o MTL4CommitFeedbackObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTL4CommitFeedbackObjectFromID constructs a [MTL4CommitFeedbackObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTL4CommitFeedbackObjectFromID(id objc.ID) MTL4CommitFeedbackObject {
	return MTL4CommitFeedbackObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A description of an error when the GPU encounters an issue as it runs the
// committed command buffers.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommitFeedback/error

func (o MTL4CommitFeedbackObject) Error() foundation.INSError {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(rv)
	}

// The host time, in seconds, when the GPU finishes execution of the committed
// command buffers.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommitFeedback/gpuEndTime

func (o MTL4CommitFeedbackObject) GPUEndTime() float64 {
	
	rv := objc.Send[float64](o.ID, objc.Sel("GPUEndTime"))
	return rv
	}

// The host time, in seconds, when the GPU starts execution of the committed
// command buffers.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommitFeedback/gpuStartTime

func (o MTL4CommitFeedbackObject) GPUStartTime() float64 {
	
	rv := objc.Send[float64](o.ID, objc.Sel("GPUStartTime"))
	return rv
	}

