// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNDArrayAffineQuantizationDescriptor] class.
var (
	_MPSNDArrayAffineQuantizationDescriptorClass     MPSNDArrayAffineQuantizationDescriptorClass
	_MPSNDArrayAffineQuantizationDescriptorClassOnce sync.Once
)

func getMPSNDArrayAffineQuantizationDescriptorClass() MPSNDArrayAffineQuantizationDescriptorClass {
	_MPSNDArrayAffineQuantizationDescriptorClassOnce.Do(func() {
		_MPSNDArrayAffineQuantizationDescriptorClass = MPSNDArrayAffineQuantizationDescriptorClass{class: objc.GetClass("MPSNDArrayAffineQuantizationDescriptor")}
	})
	return _MPSNDArrayAffineQuantizationDescriptorClass
}

// GetMPSNDArrayAffineQuantizationDescriptorClass returns the class object for MPSNDArrayAffineQuantizationDescriptor.
func GetMPSNDArrayAffineQuantizationDescriptorClass() MPSNDArrayAffineQuantizationDescriptorClass {
	return getMPSNDArrayAffineQuantizationDescriptorClass()
}

type MPSNDArrayAffineQuantizationDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayAffineQuantizationDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayAffineQuantizationDescriptorClass) Alloc() MPSNDArrayAffineQuantizationDescriptor {
	rv := objc.Send[MPSNDArrayAffineQuantizationDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNDArrayAffineQuantizationDescriptor.InitWithDataTypeHasZeroPointHasMinValue]
//
// # Instance Properties
//
//   - [MPSNDArrayAffineQuantizationDescriptor.HasMinValue]
//   - [MPSNDArrayAffineQuantizationDescriptor.SetHasMinValue]
//   - [MPSNDArrayAffineQuantizationDescriptor.HasZeroPoint]
//   - [MPSNDArrayAffineQuantizationDescriptor.SetHasZeroPoint]
//   - [MPSNDArrayAffineQuantizationDescriptor.ImplicitZeroPoint]
//   - [MPSNDArrayAffineQuantizationDescriptor.SetImplicitZeroPoint]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineQuantizationDescriptor
type MPSNDArrayAffineQuantizationDescriptor struct {
	MPSNDArrayQuantizationDescriptor
}

// MPSNDArrayAffineQuantizationDescriptorFromID constructs a [MPSNDArrayAffineQuantizationDescriptor] from an objc.ID.
func MPSNDArrayAffineQuantizationDescriptorFromID(id objc.ID) MPSNDArrayAffineQuantizationDescriptor {
	return MPSNDArrayAffineQuantizationDescriptor{MPSNDArrayQuantizationDescriptor: MPSNDArrayQuantizationDescriptorFromID(id)}
}

// NOTE: MPSNDArrayAffineQuantizationDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayAffineQuantizationDescriptor] class.
//
// # Initializers
//
//   - [IMPSNDArrayAffineQuantizationDescriptor.InitWithDataTypeHasZeroPointHasMinValue]
//
// # Instance Properties
//
//   - [IMPSNDArrayAffineQuantizationDescriptor.HasMinValue]
//   - [IMPSNDArrayAffineQuantizationDescriptor.SetHasMinValue]
//   - [IMPSNDArrayAffineQuantizationDescriptor.HasZeroPoint]
//   - [IMPSNDArrayAffineQuantizationDescriptor.SetHasZeroPoint]
//   - [IMPSNDArrayAffineQuantizationDescriptor.ImplicitZeroPoint]
//   - [IMPSNDArrayAffineQuantizationDescriptor.SetImplicitZeroPoint]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineQuantizationDescriptor
type IMPSNDArrayAffineQuantizationDescriptor interface {
	IMPSNDArrayQuantizationDescriptor

	// Topic: Initializers

	InitWithDataTypeHasZeroPointHasMinValue(quantizationDataType MPSDataType, hasZeroPoint bool, hasMinValue bool) MPSNDArrayAffineQuantizationDescriptor

	// Topic: Instance Properties

	HasMinValue() bool
	SetHasMinValue(value bool)
	HasZeroPoint() bool
	SetHasZeroPoint(value bool)
	ImplicitZeroPoint() bool
	SetImplicitZeroPoint(value bool)
}

// Init initializes the instance.
func (n MPSNDArrayAffineQuantizationDescriptor) Init() MPSNDArrayAffineQuantizationDescriptor {
	rv := objc.Send[MPSNDArrayAffineQuantizationDescriptor](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayAffineQuantizationDescriptor) Autorelease() MPSNDArrayAffineQuantizationDescriptor {
	rv := objc.Send[MPSNDArrayAffineQuantizationDescriptor](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayAffineQuantizationDescriptor creates a new MPSNDArrayAffineQuantizationDescriptor instance.
func NewMPSNDArrayAffineQuantizationDescriptor() MPSNDArrayAffineQuantizationDescriptor {
	class := getMPSNDArrayAffineQuantizationDescriptorClass()
	rv := objc.Send[MPSNDArrayAffineQuantizationDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineQuantizationDescriptor/init(dataType:hasZeroPoint:hasMinValue:)
func NewNDArrayAffineQuantizationDescriptorWithDataTypeHasZeroPointHasMinValue(quantizationDataType MPSDataType, hasZeroPoint bool, hasMinValue bool) MPSNDArrayAffineQuantizationDescriptor {
	instance := getMPSNDArrayAffineQuantizationDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDataType:hasZeroPoint:hasMinValue:"), quantizationDataType, hasZeroPoint, hasMinValue)
	return MPSNDArrayAffineQuantizationDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineQuantizationDescriptor/init(dataType:hasZeroPoint:hasMinValue:)
func (n MPSNDArrayAffineQuantizationDescriptor) InitWithDataTypeHasZeroPointHasMinValue(quantizationDataType MPSDataType, hasZeroPoint bool, hasMinValue bool) MPSNDArrayAffineQuantizationDescriptor {
	rv := objc.Send[MPSNDArrayAffineQuantizationDescriptor](n.ID, objc.Sel("initWithDataType:hasZeroPoint:hasMinValue:"), quantizationDataType, hasZeroPoint, hasMinValue)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineQuantizationDescriptor/hasMinValue
func (n MPSNDArrayAffineQuantizationDescriptor) HasMinValue() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasMinValue"))
	return rv
}
func (n MPSNDArrayAffineQuantizationDescriptor) SetHasMinValue(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setHasMinValue:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineQuantizationDescriptor/hasZeroPoint
func (n MPSNDArrayAffineQuantizationDescriptor) HasZeroPoint() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasZeroPoint"))
	return rv
}
func (n MPSNDArrayAffineQuantizationDescriptor) SetHasZeroPoint(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setHasZeroPoint:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineQuantizationDescriptor/implicitZeroPoint
func (n MPSNDArrayAffineQuantizationDescriptor) ImplicitZeroPoint() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("implicitZeroPoint"))
	return rv
}
func (n MPSNDArrayAffineQuantizationDescriptor) SetImplicitZeroPoint(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setImplicitZeroPoint:"), value)
}
