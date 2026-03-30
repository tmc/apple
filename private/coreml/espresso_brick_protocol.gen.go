// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// EspressoBrick protocol.
//
// See: https://developer.apple.com/documentation/CoreML/EspressoBrick
type EspressoBrick interface {
	objectivec.IObject

	// HasDynamicOutputShape protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/EspressoBrick/hasDynamicOutputShape:
	HasDynamicOutputShape(shape uint64) bool

	// HasGPUSupport protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/EspressoBrick/hasGPUSupport
	HasGPUSupport() bool

	// SetMappedWeightsSizeInBytes protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/EspressoBrick/setMappedWeights:sizeInBytes:
	SetMappedWeightsSizeInBytes(weights unsafe.Pointer, bytes uint64)
}

// EspressoBrickObject wraps an existing Objective-C object that conforms to the EspressoBrick protocol.
type EspressoBrickObject struct {
	objectivec.Object
}

func (o EspressoBrickObject) BaseObject() objectivec.Object {
	return o.Object
}

// EspressoBrickObjectFromID constructs a [EspressoBrickObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func EspressoBrickObjectFromID(id objc.ID) EspressoBrickObject {
	return EspressoBrickObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreML/EspressoBrick/computeDynamicOutputShape:
func (o EspressoBrickObject) ComputeDynamicOutputShape(shape objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("computeDynamicOutputShape:"), shape)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/EspressoBrick/computeOnCPUWithInputTensors:outputTensors:
func (o EspressoBrickObject) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}

// See: https://developer.apple.com/documentation/CoreML/EspressoBrick/encodeToMetalCommandBuffer:inputTensors:outputTensors:
func (o EspressoBrickObject) EncodeToMetalCommandBufferInputTensorsOutputTensors(buffer objectivec.IObject, tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("encodeToMetalCommandBuffer:inputTensors:outputTensors:"), buffer, tensors, tensors2)
}

// See: https://developer.apple.com/documentation/CoreML/EspressoBrick/hasDynamicOutputShape:
func (o EspressoBrickObject) HasDynamicOutputShape(shape uint64) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("hasDynamicOutputShape:"), shape)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/EspressoBrick/hasGPUSupport
func (o EspressoBrickObject) HasGPUSupport() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("hasGPUSupport"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/EspressoBrick/setMappedWeights:sizeInBytes:
func (o EspressoBrickObject) SetMappedWeightsSizeInBytes(weights unsafe.Pointer, bytes uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setMappedWeights:sizeInBytes:"), weights, bytes)
}

// See: https://developer.apple.com/documentation/CoreML/EspressoBrick/setupForInputShapes:withParameters:
func (o EspressoBrickObject) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
