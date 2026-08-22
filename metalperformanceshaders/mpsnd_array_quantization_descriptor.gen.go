// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNDArrayQuantizationDescriptor] class.
var (
	_MPSNDArrayQuantizationDescriptorClass     MPSNDArrayQuantizationDescriptorClass
	_MPSNDArrayQuantizationDescriptorClassOnce sync.Once
)

func getMPSNDArrayQuantizationDescriptorClass() MPSNDArrayQuantizationDescriptorClass {
	_MPSNDArrayQuantizationDescriptorClassOnce.Do(func() {
		_MPSNDArrayQuantizationDescriptorClass = MPSNDArrayQuantizationDescriptorClass{class: objc.GetClass("MPSNDArrayQuantizationDescriptor")}
	})
	return _MPSNDArrayQuantizationDescriptorClass
}

// GetMPSNDArrayQuantizationDescriptorClass returns the class object for MPSNDArrayQuantizationDescriptor.
func GetMPSNDArrayQuantizationDescriptorClass() MPSNDArrayQuantizationDescriptorClass {
	return getMPSNDArrayQuantizationDescriptorClass()
}

type MPSNDArrayQuantizationDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayQuantizationDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayQuantizationDescriptorClass) Alloc() MPSNDArrayQuantizationDescriptor {
	rv := objc.Send[MPSNDArrayQuantizationDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MPSNDArrayQuantizationDescriptor.QuantizationDataType]
//   - [MPSNDArrayQuantizationDescriptor.QuantizationScheme]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayQuantizationDescriptor
type MPSNDArrayQuantizationDescriptor struct {
	objectivec.Object
}

// MPSNDArrayQuantizationDescriptorFromID constructs a [MPSNDArrayQuantizationDescriptor] from an objc.ID.
func MPSNDArrayQuantizationDescriptorFromID(id objc.ID) MPSNDArrayQuantizationDescriptor {
	return MPSNDArrayQuantizationDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MPSNDArrayQuantizationDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayQuantizationDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSNDArrayQuantizationDescriptor.QuantizationDataType]
//   - [IMPSNDArrayQuantizationDescriptor.QuantizationScheme]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayQuantizationDescriptor
type IMPSNDArrayQuantizationDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	QuantizationDataType() MPSDataType
	QuantizationScheme() MPSNDArrayQuantizationScheme
}

// Init initializes the instance.
func (n MPSNDArrayQuantizationDescriptor) Init() MPSNDArrayQuantizationDescriptor {
	rv := objc.Send[MPSNDArrayQuantizationDescriptor](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayQuantizationDescriptor) Autorelease() MPSNDArrayQuantizationDescriptor {
	rv := objc.Send[MPSNDArrayQuantizationDescriptor](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayQuantizationDescriptor creates a new MPSNDArrayQuantizationDescriptor instance.
func NewMPSNDArrayQuantizationDescriptor() MPSNDArrayQuantizationDescriptor {
	class := getMPSNDArrayQuantizationDescriptorClass()
	rv := objc.Send[MPSNDArrayQuantizationDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayQuantizationDescriptor/quantizationDataType
func (n MPSNDArrayQuantizationDescriptor) QuantizationDataType() MPSDataType {
	rv := objc.Send[MPSDataType](n.ID, objc.Sel("quantizationDataType"))
	return MPSDataType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayQuantizationDescriptor/quantizationScheme
func (n MPSNDArrayQuantizationDescriptor) QuantizationScheme() MPSNDArrayQuantizationScheme {
	rv := objc.Send[MPSNDArrayQuantizationScheme](n.ID, objc.Sel("quantizationScheme"))
	return MPSNDArrayQuantizationScheme(rv)
}
