// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSMatrix] class.
var (
	_MPSMatrixClass     MPSMatrixClass
	_MPSMatrixClassOnce sync.Once
)

func getMPSMatrixClass() MPSMatrixClass {
	_MPSMatrixClassOnce.Do(func() {
		_MPSMatrixClass = MPSMatrixClass{class: objc.GetClass("MPSMatrix")}
	})
	return _MPSMatrixClass
}

// GetMPSMatrixClass returns the class object for MPSMatrix.
func GetMPSMatrixClass() MPSMatrixClass {
	return getMPSMatrixClass()
}

type MPSMatrixClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixClass) Alloc() MPSMatrix {
	rv := objc.Send[MPSMatrix](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A 2D array of data that stores the data’s values.
//
// # Overview
//
// [MPSMatrix] objects serve as inputs and outputs of
// [MPSMatrixMultiplication] objects. Matrix data is assumed to be stored in
// row-major order.
//
// # Methods
//
//   - [MPSMatrix.InitWithBufferDescriptor]: Initializes a matrix with a buffer.
//
// # Properties
//
//   - [MPSMatrix.Device]: The device on which the matrix will be used.
//   - [MPSMatrix.Rows]: The number of rows in the matrix.
//   - [MPSMatrix.Columns]: The number of columns in the matrix.
//   - [MPSMatrix.DataType]: The type of the values in the matrix.
//   - [MPSMatrix.RowBytes]: The stride, in bytes, between corresponding elements of consecutive rows in the matrix.
//   - [MPSMatrix.Data]: The buffer that stores the matrix data.
//   - [MPSMatrix.Matrices]
//   - [MPSMatrix.MatrixBytes]
//
// # Initializers
//
//   - [MPSMatrix.InitWithBufferOffsetDescriptor]
//   - [MPSMatrix.InitWithDeviceDescriptor]
//
// # Instance Properties
//
//   - [MPSMatrix.Offset]
//
// # Instance Methods
//
//   - [MPSMatrix.ResourceSize]
//   - [MPSMatrix.SynchronizeOnCommandBuffer]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrix
type MPSMatrix struct {
	objectivec.Object
}

// MPSMatrixFromID constructs a [MPSMatrix] from an objc.ID.
//
// A 2D array of data that stores the data’s values.
func MPSMatrixFromID(id objc.ID) MPSMatrix {
	return MPSMatrix{objectivec.Object{ID: id}}
}

// NOTE: MPSMatrix adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrix] class.
//
// # Methods
//
//   - [IMPSMatrix.InitWithBufferDescriptor]: Initializes a matrix with a buffer.
//
// # Properties
//
//   - [IMPSMatrix.Device]: The device on which the matrix will be used.
//   - [IMPSMatrix.Rows]: The number of rows in the matrix.
//   - [IMPSMatrix.Columns]: The number of columns in the matrix.
//   - [IMPSMatrix.DataType]: The type of the values in the matrix.
//   - [IMPSMatrix.RowBytes]: The stride, in bytes, between corresponding elements of consecutive rows in the matrix.
//   - [IMPSMatrix.Data]: The buffer that stores the matrix data.
//   - [IMPSMatrix.Matrices]
//   - [IMPSMatrix.MatrixBytes]
//
// # Initializers
//
//   - [IMPSMatrix.InitWithBufferOffsetDescriptor]
//   - [IMPSMatrix.InitWithDeviceDescriptor]
//
// # Instance Properties
//
//   - [IMPSMatrix.Offset]
//
// # Instance Methods
//
//   - [IMPSMatrix.ResourceSize]
//   - [IMPSMatrix.SynchronizeOnCommandBuffer]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrix
type IMPSMatrix interface {
	objectivec.IObject

	// Topic: Methods

	// Initializes a matrix with a buffer.
	InitWithBufferDescriptor(buffer metal.MTLBuffer, descriptor IMPSMatrixDescriptor) MPSMatrix

	// Topic: Properties

	// The device on which the matrix will be used.
	Device() metal.MTLDevice
	// The number of rows in the matrix.
	Rows() uint
	// The number of columns in the matrix.
	Columns() uint
	// The type of the values in the matrix.
	DataType() MPSDataType
	// The stride, in bytes, between corresponding elements of consecutive rows in the matrix.
	RowBytes() uint
	// The buffer that stores the matrix data.
	Data() metal.MTLBuffer
	Matrices() uint
	MatrixBytes() uint

	// Topic: Initializers

	InitWithBufferOffsetDescriptor(buffer metal.MTLBuffer, offset uint, descriptor IMPSMatrixDescriptor) MPSMatrix
	InitWithDeviceDescriptor(device metal.MTLDevice, descriptor IMPSMatrixDescriptor) MPSMatrix

	// Topic: Instance Properties

	Offset() uint

	// Topic: Instance Methods

	ResourceSize() uint
	SynchronizeOnCommandBuffer(commandBuffer metal.MTLCommandBuffer)
}

// Init initializes the instance.
func (m MPSMatrix) Init() MPSMatrix {
	rv := objc.Send[MPSMatrix](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrix) Autorelease() MPSMatrix {
	rv := objc.Send[MPSMatrix](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrix creates a new MPSMatrix instance.
func NewMPSMatrix() MPSMatrix {
	class := getMPSMatrixClass()
	rv := objc.Send[MPSMatrix](objc.ID(class.class), objc.Sel("new"))
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
func NewMatrixWithBufferDescriptor(buffer metal.MTLBuffer, descriptor IMPSMatrixDescriptor) MPSMatrix {
	instance := getMPSMatrixClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBuffer:descriptor:"), buffer, descriptor)
	return MPSMatrixFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrix/init(buffer:offset:descriptor:)
func NewMatrixWithBufferOffsetDescriptor(buffer metal.MTLBuffer, offset uint, descriptor IMPSMatrixDescriptor) MPSMatrix {
	instance := getMPSMatrixClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBuffer:offset:descriptor:"), buffer, offset, descriptor)
	return MPSMatrixFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrix/init(device:descriptor:)
func NewMatrixWithDeviceDescriptor(device metal.MTLDevice, descriptor IMPSMatrixDescriptor) MPSMatrix {
	instance := getMPSMatrixClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	return MPSMatrixFromID(rv)
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
func (m MPSMatrix) InitWithBufferDescriptor(buffer metal.MTLBuffer, descriptor IMPSMatrixDescriptor) MPSMatrix {
	rv := objc.Send[MPSMatrix](m.ID, objc.Sel("initWithBuffer:descriptor:"), buffer, descriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrix/init(buffer:offset:descriptor:)
func (m MPSMatrix) InitWithBufferOffsetDescriptor(buffer metal.MTLBuffer, offset uint, descriptor IMPSMatrixDescriptor) MPSMatrix {
	rv := objc.Send[MPSMatrix](m.ID, objc.Sel("initWithBuffer:offset:descriptor:"), buffer, offset, descriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrix/init(device:descriptor:)
func (m MPSMatrix) InitWithDeviceDescriptor(device metal.MTLDevice, descriptor IMPSMatrixDescriptor) MPSMatrix {
	rv := objc.Send[MPSMatrix](m.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrix/resourceSize()
func (m MPSMatrix) ResourceSize() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("resourceSize"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrix/synchronize(on:)
func (m MPSMatrix) SynchronizeOnCommandBuffer(commandBuffer metal.MTLCommandBuffer) {
	objc.Send[objc.ID](m.ID, objc.Sel("synchronizeOnCommandBuffer:"), commandBuffer)
}

// The device on which the matrix will be used.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrix/device
func (m MPSMatrix) Device() metal.MTLDevice {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("device"))
	return metal.MTLDeviceObjectFromID(rv)
}

// The number of rows in the matrix.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrix/rows
func (m MPSMatrix) Rows() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("rows"))
	return rv
}

// The number of columns in the matrix.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrix/columns
func (m MPSMatrix) Columns() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("columns"))
	return rv
}

// The type of the values in the matrix.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrix/dataType
func (m MPSMatrix) DataType() MPSDataType {
	rv := objc.Send[MPSDataType](m.ID, objc.Sel("dataType"))
	return MPSDataType(rv)
}

// The stride, in bytes, between corresponding elements of consecutive rows in
// the matrix.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrix/rowBytes
func (m MPSMatrix) RowBytes() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("rowBytes"))
	return rv
}

// The buffer that stores the matrix data.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrix/data
func (m MPSMatrix) Data() metal.MTLBuffer {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("data"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrix/matrices
func (m MPSMatrix) Matrices() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("matrices"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrix/matrixBytes
func (m MPSMatrix) MatrixBytes() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("matrixBytes"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrix/offset
func (m MPSMatrix) Offset() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("offset"))
	return rv
}
