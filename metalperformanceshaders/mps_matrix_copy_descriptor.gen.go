// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSMatrixCopyDescriptor] class.
var (
	_MPSMatrixCopyDescriptorClass     MPSMatrixCopyDescriptorClass
	_MPSMatrixCopyDescriptorClassOnce sync.Once
)

func getMPSMatrixCopyDescriptorClass() MPSMatrixCopyDescriptorClass {
	_MPSMatrixCopyDescriptorClassOnce.Do(func() {
		_MPSMatrixCopyDescriptorClass = MPSMatrixCopyDescriptorClass{class: objc.GetClass("MPSMatrixCopyDescriptor")}
	})
	return _MPSMatrixCopyDescriptorClass
}

// GetMPSMatrixCopyDescriptorClass returns the class object for MPSMatrixCopyDescriptor.
func GetMPSMatrixCopyDescriptorClass() MPSMatrixCopyDescriptorClass {
	return getMPSMatrixCopyDescriptorClass()
}

type MPSMatrixCopyDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixCopyDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixCopyDescriptorClass) Alloc() MPSMatrixCopyDescriptor {
	rv := objc.Send[MPSMatrixCopyDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of multiple matrix copy operations.
//
// # Initializers
//
//   - [MPSMatrixCopyDescriptor.InitWithDeviceCount]
//   - [MPSMatrixCopyDescriptor.InitWithSourceMatricesDestinationMatricesOffsetVectorOffset]
//
// # Instance Methods
//
//   - [MPSMatrixCopyDescriptor.SetCopyOperationAtIndexSourceMatrixDestinationMatrixOffsets]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyDescriptor
type MPSMatrixCopyDescriptor struct {
	objectivec.Object
}

// MPSMatrixCopyDescriptorFromID constructs a [MPSMatrixCopyDescriptor] from an objc.ID.
//
// A description of multiple matrix copy operations.
func MPSMatrixCopyDescriptorFromID(id objc.ID) MPSMatrixCopyDescriptor {
	return MPSMatrixCopyDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MPSMatrixCopyDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixCopyDescriptor] class.
//
// # Initializers
//
//   - [IMPSMatrixCopyDescriptor.InitWithDeviceCount]
//   - [IMPSMatrixCopyDescriptor.InitWithSourceMatricesDestinationMatricesOffsetVectorOffset]
//
// # Instance Methods
//
//   - [IMPSMatrixCopyDescriptor.SetCopyOperationAtIndexSourceMatrixDestinationMatrixOffsets]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyDescriptor
type IMPSMatrixCopyDescriptor interface {
	objectivec.IObject

	// Topic: Initializers

	InitWithDeviceCount(device metal.MTLDevice, count uint) MPSMatrixCopyDescriptor
	InitWithSourceMatricesDestinationMatricesOffsetVectorOffset(sourceMatrices []MPSMatrix, destinationMatrices []MPSMatrix, offsets IMPSVector, byteOffset uint) MPSMatrixCopyDescriptor

	// Topic: Instance Methods

	SetCopyOperationAtIndexSourceMatrixDestinationMatrixOffsets(index uint, sourceMatrix IMPSMatrix, destinationMatrix IMPSMatrix, offsets MPSMatrixCopyOffsets)
}

// Init initializes the instance.
func (m MPSMatrixCopyDescriptor) Init() MPSMatrixCopyDescriptor {
	rv := objc.Send[MPSMatrixCopyDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixCopyDescriptor) Autorelease() MPSMatrixCopyDescriptor {
	rv := objc.Send[MPSMatrixCopyDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixCopyDescriptor creates a new MPSMatrixCopyDescriptor instance.
func NewMPSMatrixCopyDescriptor() MPSMatrixCopyDescriptor {
	class := getMPSMatrixCopyDescriptorClass()
	rv := objc.Send[MPSMatrixCopyDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyDescriptor/init(device:count:)
func NewMatrixCopyDescriptorWithDeviceCount(device metal.MTLDevice, count uint) MPSMatrixCopyDescriptor {
	instance := getMPSMatrixCopyDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:count:"), device, count)
	return MPSMatrixCopyDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyDescriptor/init(sourceMatrices:destinationMatrices:offsetVector:offset:)
func NewMatrixCopyDescriptorWithSourceMatricesDestinationMatricesOffsetVectorOffset(sourceMatrices []MPSMatrix, destinationMatrices []MPSMatrix, offsets IMPSVector, byteOffset uint) MPSMatrixCopyDescriptor {
	instance := getMPSMatrixCopyDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSourceMatrices:destinationMatrices:offsetVector:offset:"), objectivec.IObjectSliceToNSArray(sourceMatrices), objectivec.IObjectSliceToNSArray(destinationMatrices), offsets, byteOffset)
	return MPSMatrixCopyDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyDescriptor/init(sourceMatrix:destinationMatrix:offsets:)
func NewMatrixCopyDescriptorWithSourceMatrixDestinationMatrixOffsets(sourceMatrix IMPSMatrix, destinationMatrix IMPSMatrix, offsets MPSMatrixCopyOffsets) MPSMatrixCopyDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSMatrixCopyDescriptorClass().class), objc.Sel("descriptorWithSourceMatrix:destinationMatrix:offsets:"), sourceMatrix, destinationMatrix, offsets)
	return MPSMatrixCopyDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyDescriptor/init(device:count:)
func (m MPSMatrixCopyDescriptor) InitWithDeviceCount(device metal.MTLDevice, count uint) MPSMatrixCopyDescriptor {
	rv := objc.Send[MPSMatrixCopyDescriptor](m.ID, objc.Sel("initWithDevice:count:"), device, count)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyDescriptor/init(sourceMatrices:destinationMatrices:offsetVector:offset:)
func (m MPSMatrixCopyDescriptor) InitWithSourceMatricesDestinationMatricesOffsetVectorOffset(sourceMatrices []MPSMatrix, destinationMatrices []MPSMatrix, offsets IMPSVector, byteOffset uint) MPSMatrixCopyDescriptor {
	rv := objc.Send[MPSMatrixCopyDescriptor](m.ID, objc.Sel("initWithSourceMatrices:destinationMatrices:offsetVector:offset:"), objectivec.IObjectSliceToNSArray(sourceMatrices), objectivec.IObjectSliceToNSArray(destinationMatrices), offsets, byteOffset)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyDescriptor/setCopyOperationAt(_:sourceMatrix:destinationMatrix:offsets:)
func (m MPSMatrixCopyDescriptor) SetCopyOperationAtIndexSourceMatrixDestinationMatrixOffsets(index uint, sourceMatrix IMPSMatrix, destinationMatrix IMPSMatrix, offsets MPSMatrixCopyOffsets) {
	objc.Send[objc.ID](m.ID, objc.Sel("setCopyOperationAtIndex:sourceMatrix:destinationMatrix:offsets:"), index, sourceMatrix, destinationMatrix, offsets)
}
