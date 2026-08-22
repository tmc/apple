// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSTemporaryNDArray] class.
var (
	_MPSTemporaryNDArrayClass     MPSTemporaryNDArrayClass
	_MPSTemporaryNDArrayClassOnce sync.Once
)

func getMPSTemporaryNDArrayClass() MPSTemporaryNDArrayClass {
	_MPSTemporaryNDArrayClassOnce.Do(func() {
		_MPSTemporaryNDArrayClass = MPSTemporaryNDArrayClass{class: objc.GetClass("MPSTemporaryNDArray")}
	})
	return _MPSTemporaryNDArrayClass
}

// GetMPSTemporaryNDArrayClass returns the class object for MPSTemporaryNDArray.
func GetMPSTemporaryNDArrayClass() MPSTemporaryNDArrayClass {
	return getMPSTemporaryNDArrayClass()
}

type MPSTemporaryNDArrayClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSTemporaryNDArrayClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSTemporaryNDArrayClass) Alloc() MPSTemporaryNDArray {
	rv := objc.Send[MPSTemporaryNDArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MPSTemporaryNDArray.ReadCount]
//   - [MPSTemporaryNDArray.SetReadCount]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryNDArray
type MPSTemporaryNDArray struct {
	MPSNDArray
}

// MPSTemporaryNDArrayFromID constructs a [MPSTemporaryNDArray] from an objc.ID.
func MPSTemporaryNDArrayFromID(id objc.ID) MPSTemporaryNDArray {
	return MPSTemporaryNDArray{MPSNDArray: MPSNDArrayFromID(id)}
}

// NOTE: MPSTemporaryNDArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSTemporaryNDArray] class.
//
// # Instance Properties
//
//   - [IMPSTemporaryNDArray.ReadCount]
//   - [IMPSTemporaryNDArray.SetReadCount]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryNDArray
type IMPSTemporaryNDArray interface {
	IMPSNDArray

	// Topic: Instance Properties

	ReadCount() uint
	SetReadCount(value uint)
}

// Init initializes the instance.
func (t MPSTemporaryNDArray) Init() MPSTemporaryNDArray {
	rv := objc.Send[MPSTemporaryNDArray](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MPSTemporaryNDArray) Autorelease() MPSTemporaryNDArray {
	rv := objc.Send[MPSTemporaryNDArray](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSTemporaryNDArray creates a new MPSTemporaryNDArray instance.
func NewMPSTemporaryNDArray() MPSTemporaryNDArray {
	class := getMPSTemporaryNDArrayClass()
	rv := objc.Send[MPSTemporaryNDArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/init(buffer:offset:descriptor:)
func NewTemporaryNDArrayWithBufferOffsetDescriptor(buffer metal.MTLBuffer, offset uint, descriptor IMPSNDArrayDescriptor) MPSTemporaryNDArray {
	instance := getMPSTemporaryNDArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBuffer:offset:descriptor:"), buffer, offset, descriptor)
	return MPSTemporaryNDArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryNDArray/init(commandBuffer:descriptor:)
func NewTemporaryNDArrayWithCommandBufferDescriptor(commandBuffer metal.MTLCommandBuffer, descriptor IMPSNDArrayDescriptor) MPSTemporaryNDArray {
	rv := objc.Send[objc.ID](objc.ID(getMPSTemporaryNDArrayClass().class), objc.Sel("temporaryNDArrayWithCommandBuffer:descriptor:"), commandBuffer, descriptor)
	return MPSTemporaryNDArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/init(device:descriptor:)
func NewTemporaryNDArrayWithDeviceDescriptor(device metal.MTLDevice, descriptor IMPSNDArrayDescriptor) MPSTemporaryNDArray {
	instance := getMPSTemporaryNDArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	return MPSTemporaryNDArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/init(device:scalar:)
func NewTemporaryNDArrayWithDeviceScalar(device metal.MTLDevice, value float64) MPSTemporaryNDArray {
	instance := getMPSTemporaryNDArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:scalar:"), device, value)
	return MPSTemporaryNDArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryNDArray/readCount
func (t MPSTemporaryNDArray) ReadCount() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("readCount"))
	return rv
}
func (t MPSTemporaryNDArray) SetReadCount(value uint) {
	objc.Send[struct{}](t.ID, objc.Sel("setReadCount:"), value)
}
