// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSTemporaryMatrix] class.
var (
	_MPSTemporaryMatrixClass     MPSTemporaryMatrixClass
	_MPSTemporaryMatrixClassOnce sync.Once
)

func getMPSTemporaryMatrixClass() MPSTemporaryMatrixClass {
	_MPSTemporaryMatrixClassOnce.Do(func() {
		_MPSTemporaryMatrixClass = MPSTemporaryMatrixClass{class: objc.GetClass("MPSTemporaryMatrix")}
	})
	return _MPSTemporaryMatrixClass
}

// GetMPSTemporaryMatrixClass returns the class object for MPSTemporaryMatrix.
func GetMPSTemporaryMatrixClass() MPSTemporaryMatrixClass {
	return getMPSTemporaryMatrixClass()
}

type MPSTemporaryMatrixClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSTemporaryMatrixClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSTemporaryMatrixClass) Alloc() MPSTemporaryMatrix {
	rv := objc.Send[MPSTemporaryMatrix](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A matrix allocated on GPU private memory.
//
// # Instance Properties
//
//   - [MPSTemporaryMatrix.ReadCount]
//   - [MPSTemporaryMatrix.SetReadCount]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryMatrix
type MPSTemporaryMatrix struct {
	MPSMatrix
}

// MPSTemporaryMatrixFromID constructs a [MPSTemporaryMatrix] from an objc.ID.
//
// A matrix allocated on GPU private memory.
func MPSTemporaryMatrixFromID(id objc.ID) MPSTemporaryMatrix {
	return MPSTemporaryMatrix{MPSMatrix: MPSMatrixFromID(id)}
}

// NOTE: MPSTemporaryMatrix adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSTemporaryMatrix] class.
//
// # Instance Properties
//
//   - [IMPSTemporaryMatrix.ReadCount]
//   - [IMPSTemporaryMatrix.SetReadCount]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryMatrix
type IMPSTemporaryMatrix interface {
	IMPSMatrix

	// Topic: Instance Properties

	ReadCount() uint
	SetReadCount(value uint)
}

// Init initializes the instance.
func (t MPSTemporaryMatrix) Init() MPSTemporaryMatrix {
	rv := objc.Send[MPSTemporaryMatrix](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MPSTemporaryMatrix) Autorelease() MPSTemporaryMatrix {
	rv := objc.Send[MPSTemporaryMatrix](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSTemporaryMatrix creates a new MPSTemporaryMatrix instance.
func NewMPSTemporaryMatrix() MPSTemporaryMatrix {
	class := getMPSTemporaryMatrixClass()
	rv := objc.Send[MPSTemporaryMatrix](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a matrix with a buffer.
//
// buffer: The buffer that stores the matrix data.
//
// descriptor: The matrix descriptor.
//
// # Return Value
//
// A valid [MPSMatrix] object or `nil`, if failure.
//
// # Discussion
//
// The dimensions and stride of the matrix are specified by the
// [MPSMatrixDescriptor] object. The size of the provided [MTLBuffer] object
// must be large enough to store the following amount of bytes:
//
// `(descriptor.Rows()-1) * descriptor.RowBytes() + descriptor.Columns() *
// (element size)`
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrix/init(buffer:descriptor:)
//
// [MTLBuffer]: https://developer.apple.com/documentation/Metal/MTLBuffer
func NewTemporaryMatrixWithBufferDescriptor(buffer metal.MTLBuffer, descriptor IMPSMatrixDescriptor) MPSTemporaryMatrix {
	instance := getMPSTemporaryMatrixClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBuffer:descriptor:"), buffer, descriptor)
	return MPSTemporaryMatrixFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrix/init(buffer:offset:descriptor:)
func NewTemporaryMatrixWithBufferOffsetDescriptor(buffer metal.MTLBuffer, offset uint, descriptor IMPSMatrixDescriptor) MPSTemporaryMatrix {
	instance := getMPSTemporaryMatrixClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBuffer:offset:descriptor:"), buffer, offset, descriptor)
	return MPSTemporaryMatrixFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryMatrix/init(commandBuffer:matrixDescriptor:)
func NewTemporaryMatrixWithCommandBufferMatrixDescriptor(commandBuffer metal.MTLCommandBuffer, matrixDescriptor IMPSMatrixDescriptor) MPSTemporaryMatrix {
	rv := objc.Send[objc.ID](objc.ID(getMPSTemporaryMatrixClass().class), objc.Sel("temporaryMatrixWithCommandBuffer:matrixDescriptor:"), commandBuffer, matrixDescriptor)
	return MPSTemporaryMatrixFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrix/init(device:descriptor:)
func NewTemporaryMatrixWithDeviceDescriptor(device metal.MTLDevice, descriptor IMPSMatrixDescriptor) MPSTemporaryMatrix {
	instance := getMPSTemporaryMatrixClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	return MPSTemporaryMatrixFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryMatrix/prefetchStorage(with:matrixDescriptorList:)
func (_MPSTemporaryMatrixClass MPSTemporaryMatrixClass) PrefetchStorageWithCommandBufferMatrixDescriptorList(commandBuffer metal.MTLCommandBuffer, descriptorList []MPSMatrixDescriptor) {
	objc.Send[objc.ID](objc.ID(_MPSTemporaryMatrixClass.class), objc.Sel("prefetchStorageWithCommandBuffer:matrixDescriptorList:"), commandBuffer, objectivec.IObjectSliceToNSArray(descriptorList))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryMatrix/readCount
func (t MPSTemporaryMatrix) ReadCount() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("readCount"))
	return rv
}
func (t MPSTemporaryMatrix) SetReadCount(value uint) {
	objc.Send[struct{}](t.ID, objc.Sel("setReadCount:"), value)
}
