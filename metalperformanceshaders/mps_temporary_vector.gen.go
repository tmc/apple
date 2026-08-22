// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSTemporaryVector] class.
var (
	_MPSTemporaryVectorClass     MPSTemporaryVectorClass
	_MPSTemporaryVectorClassOnce sync.Once
)

func getMPSTemporaryVectorClass() MPSTemporaryVectorClass {
	_MPSTemporaryVectorClassOnce.Do(func() {
		_MPSTemporaryVectorClass = MPSTemporaryVectorClass{class: objc.GetClass("MPSTemporaryVector")}
	})
	return _MPSTemporaryVectorClass
}

// GetMPSTemporaryVectorClass returns the class object for MPSTemporaryVector.
func GetMPSTemporaryVectorClass() MPSTemporaryVectorClass {
	return getMPSTemporaryVectorClass()
}

type MPSTemporaryVectorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSTemporaryVectorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSTemporaryVectorClass) Alloc() MPSTemporaryVector {
	rv := objc.Send[MPSTemporaryVector](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A vector allocated on GPU private memory.
//
// # Instance Properties
//
//   - [MPSTemporaryVector.ReadCount]
//   - [MPSTemporaryVector.SetReadCount]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryVector
type MPSTemporaryVector struct {
	MPSVector
}

// MPSTemporaryVectorFromID constructs a [MPSTemporaryVector] from an objc.ID.
//
// A vector allocated on GPU private memory.
func MPSTemporaryVectorFromID(id objc.ID) MPSTemporaryVector {
	return MPSTemporaryVector{MPSVector: MPSVectorFromID(id)}
}

// NOTE: MPSTemporaryVector adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSTemporaryVector] class.
//
// # Instance Properties
//
//   - [IMPSTemporaryVector.ReadCount]
//   - [IMPSTemporaryVector.SetReadCount]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryVector
type IMPSTemporaryVector interface {
	IMPSVector

	// Topic: Instance Properties

	ReadCount() uint
	SetReadCount(value uint)
}

// Init initializes the instance.
func (t MPSTemporaryVector) Init() MPSTemporaryVector {
	rv := objc.Send[MPSTemporaryVector](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MPSTemporaryVector) Autorelease() MPSTemporaryVector {
	rv := objc.Send[MPSTemporaryVector](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSTemporaryVector creates a new MPSTemporaryVector instance.
func NewMPSTemporaryVector() MPSTemporaryVector {
	class := getMPSTemporaryVectorClass()
	rv := objc.Send[MPSTemporaryVector](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVector/init(buffer:descriptor:)
func NewTemporaryVectorWithBufferDescriptor(buffer metal.MTLBuffer, descriptor IMPSVectorDescriptor) MPSTemporaryVector {
	instance := getMPSTemporaryVectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBuffer:descriptor:"), buffer, descriptor)
	return MPSTemporaryVectorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVector/init(buffer:offset:descriptor:)
func NewTemporaryVectorWithBufferOffsetDescriptor(buffer metal.MTLBuffer, offset uint, descriptor IMPSVectorDescriptor) MPSTemporaryVector {
	instance := getMPSTemporaryVectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBuffer:offset:descriptor:"), buffer, offset, descriptor)
	return MPSTemporaryVectorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryVector/init(commandBuffer:descriptor:)
func NewTemporaryVectorWithCommandBufferDescriptor(commandBuffer metal.MTLCommandBuffer, descriptor IMPSVectorDescriptor) MPSTemporaryVector {
	rv := objc.Send[objc.ID](objc.ID(getMPSTemporaryVectorClass().class), objc.Sel("temporaryVectorWithCommandBuffer:descriptor:"), commandBuffer, descriptor)
	return MPSTemporaryVectorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVector/init(device:descriptor:)
func NewTemporaryVectorWithDeviceDescriptor(device metal.MTLDevice, descriptor IMPSVectorDescriptor) MPSTemporaryVector {
	instance := getMPSTemporaryVectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	return MPSTemporaryVectorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryVector/prefetchStorage(with:descriptorList:)
func (_MPSTemporaryVectorClass MPSTemporaryVectorClass) PrefetchStorageWithCommandBufferDescriptorList(commandBuffer metal.MTLCommandBuffer, descriptorList []MPSVectorDescriptor) {
	objc.Send[objc.ID](objc.ID(_MPSTemporaryVectorClass.class), objc.Sel("prefetchStorageWithCommandBuffer:descriptorList:"), commandBuffer, objectivec.IObjectSliceToNSArray(descriptorList))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryVector/readCount
func (t MPSTemporaryVector) ReadCount() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("readCount"))
	return rv
}
func (t MPSTemporaryVector) SetReadCount(value uint) {
	objc.Send[struct{}](t.ID, objc.Sel("setReadCount:"), value)
}
