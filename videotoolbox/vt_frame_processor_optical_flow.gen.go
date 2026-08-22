// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"sync"

	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VTFrameProcessorOpticalFlow] class.
var (
	_VTFrameProcessorOpticalFlowClass     VTFrameProcessorOpticalFlowClass
	_VTFrameProcessorOpticalFlowClassOnce sync.Once
)

func getVTFrameProcessorOpticalFlowClass() VTFrameProcessorOpticalFlowClass {
	_VTFrameProcessorOpticalFlowClassOnce.Do(func() {
		_VTFrameProcessorOpticalFlowClass = VTFrameProcessorOpticalFlowClass{class: objc.GetClass("VTFrameProcessorOpticalFlow")}
	})
	return _VTFrameProcessorOpticalFlowClass
}

// GetVTFrameProcessorOpticalFlowClass returns the class object for VTFrameProcessorOpticalFlow.
func GetVTFrameProcessorOpticalFlowClass() VTFrameProcessorOpticalFlowClass {
	return getVTFrameProcessorOpticalFlowClass()
}

type VTFrameProcessorOpticalFlowClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VTFrameProcessorOpticalFlowClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VTFrameProcessorOpticalFlowClass) Alloc() VTFrameProcessorOpticalFlow {
	rv := objc.Send[VTFrameProcessorOpticalFlow](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A class to wrap bidirectional optical flow to send to the processor.
//
// # Overview
//
// Instances retain the buffers backing them.
//
// # Creating an optical flow configuration
//
//   - [VTFrameProcessorOpticalFlow.InitWithForwardFlowBackwardFlow]: Creates an object with forward and backward optical flow pixel buffers.
//
// # Inspecting the configuration
//
//   - [VTFrameProcessorOpticalFlow.BackwardFlow]: The backward optical flow pixel buffer that was provided when the object was created.
//   - [VTFrameProcessorOpticalFlow.ForwardFlow]: The forward optical flow pixel that was provided when the object was created.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorOpticalFlow
type VTFrameProcessorOpticalFlow struct {
	objectivec.Object
}

// VTFrameProcessorOpticalFlowFromID constructs a [VTFrameProcessorOpticalFlow] from an objc.ID.
//
// A class to wrap bidirectional optical flow to send to the processor.
func VTFrameProcessorOpticalFlowFromID(id objc.ID) VTFrameProcessorOpticalFlow {
	return VTFrameProcessorOpticalFlow{objectivec.Object{ID: id}}
}

// NOTE: VTFrameProcessorOpticalFlow adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VTFrameProcessorOpticalFlow] class.
//
// # Creating an optical flow configuration
//
//   - [IVTFrameProcessorOpticalFlow.InitWithForwardFlowBackwardFlow]: Creates an object with forward and backward optical flow pixel buffers.
//
// # Inspecting the configuration
//
//   - [IVTFrameProcessorOpticalFlow.BackwardFlow]: The backward optical flow pixel buffer that was provided when the object was created.
//   - [IVTFrameProcessorOpticalFlow.ForwardFlow]: The forward optical flow pixel that was provided when the object was created.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorOpticalFlow
type IVTFrameProcessorOpticalFlow interface {
	objectivec.IObject

	// Topic: Creating an optical flow configuration

	// Creates an object with forward and backward optical flow pixel buffers.
	InitWithForwardFlowBackwardFlow(forwardFlow corevideo.CVImageBufferRef, backwardFlow corevideo.CVImageBufferRef) VTFrameProcessorOpticalFlow

	// Topic: Inspecting the configuration

	// The backward optical flow pixel buffer that was provided when the object was created.
	BackwardFlow() corevideo.CVImageBufferRef
	// The forward optical flow pixel that was provided when the object was created.
	ForwardFlow() corevideo.CVImageBufferRef
}

// Init initializes the instance.
func (v VTFrameProcessorOpticalFlow) Init() VTFrameProcessorOpticalFlow {
	rv := objc.Send[VTFrameProcessorOpticalFlow](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VTFrameProcessorOpticalFlow) Autorelease() VTFrameProcessorOpticalFlow {
	rv := objc.Send[VTFrameProcessorOpticalFlow](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVTFrameProcessorOpticalFlow creates a new VTFrameProcessorOpticalFlow instance.
func NewVTFrameProcessorOpticalFlow() VTFrameProcessorOpticalFlow {
	class := getVTFrameProcessorOpticalFlowClass()
	rv := objc.Send[VTFrameProcessorOpticalFlow](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an object with forward and backward optical flow pixel buffers.
//
// forwardFlow: A pixel buffer that contains forward optical flow. This value must be
// non-NULL and IOSurface backed.
//
// backwardFlow: A pixel buffer that contains backward optical flow. his value must be
// non-NULL and IOSurface backed.
//
// # Discussion
//
// Instances retain the buffers backing them. Returns NULL if a NULL
// CVPixelBuffer is provided or if CVPixelBuffers are not IOSurface backed.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorOpticalFlow/init(forwardFlow:backwardFlow:)
func NewVTFrameProcessorOpticalFlowWithForwardFlowBackwardFlow(forwardFlow corevideo.CVImageBufferRef, backwardFlow corevideo.CVImageBufferRef) VTFrameProcessorOpticalFlow {
	instance := getVTFrameProcessorOpticalFlowClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithForwardFlow:backwardFlow:"), forwardFlow, backwardFlow)
	return VTFrameProcessorOpticalFlowFromID(rv)
}

// Creates an object with forward and backward optical flow pixel buffers.
//
// forwardFlow: A pixel buffer that contains forward optical flow. This value must be
// non-NULL and IOSurface backed.
//
// backwardFlow: A pixel buffer that contains backward optical flow. his value must be
// non-NULL and IOSurface backed.
//
// # Discussion
//
// Instances retain the buffers backing them. Returns NULL if a NULL
// CVPixelBuffer is provided or if CVPixelBuffers are not IOSurface backed.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorOpticalFlow/init(forwardFlow:backwardFlow:)
func (v VTFrameProcessorOpticalFlow) InitWithForwardFlowBackwardFlow(forwardFlow corevideo.CVImageBufferRef, backwardFlow corevideo.CVImageBufferRef) VTFrameProcessorOpticalFlow {
	rv := objc.Send[VTFrameProcessorOpticalFlow](v.ID, objc.Sel("initWithForwardFlow:backwardFlow:"), forwardFlow, backwardFlow)
	return rv
}

// The backward optical flow pixel buffer that was provided when the object
// was created.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorOpticalFlow/backwardFlow
func (v VTFrameProcessorOpticalFlow) BackwardFlow() corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](v.ID, objc.Sel("backwardFlow"))
	return corevideo.CVImageBufferRef(rv)
}

// The forward optical flow pixel that was provided when the object was
// created.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameProcessorOpticalFlow/forwardFlow
func (v VTFrameProcessorOpticalFlow) ForwardFlow() corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](v.ID, objc.Sel("forwardFlow"))
	return corevideo.CVImageBufferRef(rv)
}
