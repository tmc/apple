// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNDArrayLUTQuantizationDescriptor] class.
var (
	_MPSNDArrayLUTQuantizationDescriptorClass     MPSNDArrayLUTQuantizationDescriptorClass
	_MPSNDArrayLUTQuantizationDescriptorClassOnce sync.Once
)

func getMPSNDArrayLUTQuantizationDescriptorClass() MPSNDArrayLUTQuantizationDescriptorClass {
	_MPSNDArrayLUTQuantizationDescriptorClassOnce.Do(func() {
		_MPSNDArrayLUTQuantizationDescriptorClass = MPSNDArrayLUTQuantizationDescriptorClass{class: objc.GetClass("MPSNDArrayLUTQuantizationDescriptor")}
	})
	return _MPSNDArrayLUTQuantizationDescriptorClass
}

// GetMPSNDArrayLUTQuantizationDescriptorClass returns the class object for MPSNDArrayLUTQuantizationDescriptor.
func GetMPSNDArrayLUTQuantizationDescriptorClass() MPSNDArrayLUTQuantizationDescriptorClass {
	return getMPSNDArrayLUTQuantizationDescriptorClass()
}

type MPSNDArrayLUTQuantizationDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayLUTQuantizationDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayLUTQuantizationDescriptorClass) Alloc() MPSNDArrayLUTQuantizationDescriptor {
	rv := objc.Send[MPSNDArrayLUTQuantizationDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNDArrayLUTQuantizationDescriptor.InitWithDataType]
//   - [MPSNDArrayLUTQuantizationDescriptor.InitWithDataTypeVectorAxis]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayLUTQuantizationDescriptor
type MPSNDArrayLUTQuantizationDescriptor struct {
	MPSNDArrayQuantizationDescriptor
}

// MPSNDArrayLUTQuantizationDescriptorFromID constructs a [MPSNDArrayLUTQuantizationDescriptor] from an objc.ID.
func MPSNDArrayLUTQuantizationDescriptorFromID(id objc.ID) MPSNDArrayLUTQuantizationDescriptor {
	return MPSNDArrayLUTQuantizationDescriptor{MPSNDArrayQuantizationDescriptor: MPSNDArrayQuantizationDescriptorFromID(id)}
}

// NOTE: MPSNDArrayLUTQuantizationDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayLUTQuantizationDescriptor] class.
//
// # Initializers
//
//   - [IMPSNDArrayLUTQuantizationDescriptor.InitWithDataType]
//   - [IMPSNDArrayLUTQuantizationDescriptor.InitWithDataTypeVectorAxis]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayLUTQuantizationDescriptor
type IMPSNDArrayLUTQuantizationDescriptor interface {
	IMPSNDArrayQuantizationDescriptor

	// Topic: Initializers

	InitWithDataType(quantizationDataType MPSDataType) MPSNDArrayLUTQuantizationDescriptor
	InitWithDataTypeVectorAxis(quantizationDataType MPSDataType, vectorAxis uint) MPSNDArrayLUTQuantizationDescriptor
}

// Init initializes the instance.
func (n MPSNDArrayLUTQuantizationDescriptor) Init() MPSNDArrayLUTQuantizationDescriptor {
	rv := objc.Send[MPSNDArrayLUTQuantizationDescriptor](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayLUTQuantizationDescriptor) Autorelease() MPSNDArrayLUTQuantizationDescriptor {
	rv := objc.Send[MPSNDArrayLUTQuantizationDescriptor](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayLUTQuantizationDescriptor creates a new MPSNDArrayLUTQuantizationDescriptor instance.
func NewMPSNDArrayLUTQuantizationDescriptor() MPSNDArrayLUTQuantizationDescriptor {
	class := getMPSNDArrayLUTQuantizationDescriptorClass()
	rv := objc.Send[MPSNDArrayLUTQuantizationDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayLUTQuantizationDescriptor/init(dataType:)
func NewNDArrayLUTQuantizationDescriptorWithDataType(quantizationDataType MPSDataType) MPSNDArrayLUTQuantizationDescriptor {
	instance := getMPSNDArrayLUTQuantizationDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDataType:"), quantizationDataType)
	return MPSNDArrayLUTQuantizationDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayLUTQuantizationDescriptor/init(dataType:vectorAxis:)
func NewNDArrayLUTQuantizationDescriptorWithDataTypeVectorAxis(quantizationDataType MPSDataType, vectorAxis uint) MPSNDArrayLUTQuantizationDescriptor {
	instance := getMPSNDArrayLUTQuantizationDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDataType:vectorAxis:"), quantizationDataType, vectorAxis)
	return MPSNDArrayLUTQuantizationDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayLUTQuantizationDescriptor/init(dataType:)
func (n MPSNDArrayLUTQuantizationDescriptor) InitWithDataType(quantizationDataType MPSDataType) MPSNDArrayLUTQuantizationDescriptor {
	rv := objc.Send[MPSNDArrayLUTQuantizationDescriptor](n.ID, objc.Sel("initWithDataType:"), quantizationDataType)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayLUTQuantizationDescriptor/init(dataType:vectorAxis:)
func (n MPSNDArrayLUTQuantizationDescriptor) InitWithDataTypeVectorAxis(quantizationDataType MPSDataType, vectorAxis uint) MPSNDArrayLUTQuantizationDescriptor {
	rv := objc.Send[MPSNDArrayLUTQuantizationDescriptor](n.ID, objc.Sel("initWithDataType:vectorAxis:"), quantizationDataType, vectorAxis)
	return rv
}
