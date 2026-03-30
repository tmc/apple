// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A pipeline state that you can use with machine-learning encoder instances.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineState
type MTL4MachineLearningPipelineState interface {
	objectivec.IObject
	MTLAllocation

	// Returns the device the pipeline state belongs to.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineState/device
	Device() MTLDevice

	// Obtain the size of the heap, in bytes, this pipeline requires during the execution.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineState/intermediatesHeapSize
	IntermediatesHeapSize() uint

	// Queries the string that helps identify this object.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineState/label
	Label() string

	// Returns reflection information for this machine learning pipeline state.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineState/reflection
	Reflection() IMTL4MachineLearningPipelineReflection
}

// MTL4MachineLearningPipelineStateObject wraps an existing Objective-C object that conforms to the MTL4MachineLearningPipelineState protocol.
type MTL4MachineLearningPipelineStateObject struct {
	objectivec.Object
}

func (o MTL4MachineLearningPipelineStateObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTL4MachineLearningPipelineStateObjectFromID constructs a [MTL4MachineLearningPipelineStateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTL4MachineLearningPipelineStateObjectFromID(id objc.ID) MTL4MachineLearningPipelineStateObject {
	return MTL4MachineLearningPipelineStateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the device the pipeline state belongs to.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineState/device
func (o MTL4MachineLearningPipelineStateObject) Device() MTLDevice {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
}

// Obtain the size of the heap, in bytes, this pipeline requires during the
// execution.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineState/intermediatesHeapSize
func (o MTL4MachineLearningPipelineStateObject) IntermediatesHeapSize() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("intermediatesHeapSize"))
	return rv
}

// Queries the string that helps identify this object.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineState/label
func (o MTL4MachineLearningPipelineStateObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

// Returns reflection information for this machine learning pipeline state.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineState/reflection
func (o MTL4MachineLearningPipelineStateObject) Reflection() IMTL4MachineLearningPipelineReflection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("reflection"))
	return MTL4MachineLearningPipelineReflectionFromID(rv)
}

// The amount of memory, in byes, a resource consumes, such as for a buffer,
// texture, or heap.
//
// See: https://developer.apple.com/documentation/Metal/MTLAllocation/allocatedSize
func (o MTL4MachineLearningPipelineStateObject) AllocatedSize() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("allocatedSize"))
	return rv
}
