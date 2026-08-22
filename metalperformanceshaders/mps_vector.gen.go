// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSVector] class.
var (
	_MPSVectorClass     MPSVectorClass
	_MPSVectorClassOnce sync.Once
)

func getMPSVectorClass() MPSVectorClass {
	_MPSVectorClassOnce.Do(func() {
		_MPSVectorClass = MPSVectorClass{class: objc.GetClass("MPSVector")}
	})
	return _MPSVectorClass
}

// GetMPSVectorClass returns the class object for MPSVector.
func GetMPSVectorClass() MPSVectorClass {
	return getMPSVectorClass()
}

type MPSVectorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSVectorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSVectorClass) Alloc() MPSVector {
	rv := objc.Send[MPSVector](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A 1D array of data that stores the data’s values.
//
// # Initializers
//
//   - [MPSVector.InitWithBufferDescriptor]
//   - [MPSVector.InitWithBufferOffsetDescriptor]
//   - [MPSVector.InitWithDeviceDescriptor]
//
// # Instance Properties
//
//   - [MPSVector.Data]
//   - [MPSVector.DataType]
//   - [MPSVector.Device]
//   - [MPSVector.Length]
//   - [MPSVector.VectorBytes]
//   - [MPSVector.Vectors]
//   - [MPSVector.Offset]
//
// # Instance Methods
//
//   - [MPSVector.ResourceSize]
//   - [MPSVector.SynchronizeOnCommandBuffer]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVector
type MPSVector struct {
	objectivec.Object
}

// MPSVectorFromID constructs a [MPSVector] from an objc.ID.
//
// A 1D array of data that stores the data’s values.
func MPSVectorFromID(id objc.ID) MPSVector {
	return MPSVector{objectivec.Object{ID: id}}
}

// NOTE: MPSVector adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSVector] class.
//
// # Initializers
//
//   - [IMPSVector.InitWithBufferDescriptor]
//   - [IMPSVector.InitWithBufferOffsetDescriptor]
//   - [IMPSVector.InitWithDeviceDescriptor]
//
// # Instance Properties
//
//   - [IMPSVector.Data]
//   - [IMPSVector.DataType]
//   - [IMPSVector.Device]
//   - [IMPSVector.Length]
//   - [IMPSVector.VectorBytes]
//   - [IMPSVector.Vectors]
//   - [IMPSVector.Offset]
//
// # Instance Methods
//
//   - [IMPSVector.ResourceSize]
//   - [IMPSVector.SynchronizeOnCommandBuffer]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVector
type IMPSVector interface {
	objectivec.IObject

	// Topic: Initializers

	InitWithBufferDescriptor(buffer metal.MTLBuffer, descriptor IMPSVectorDescriptor) MPSVector
	InitWithBufferOffsetDescriptor(buffer metal.MTLBuffer, offset uint, descriptor IMPSVectorDescriptor) MPSVector
	InitWithDeviceDescriptor(device metal.MTLDevice, descriptor IMPSVectorDescriptor) MPSVector

	// Topic: Instance Properties

	Data() metal.MTLBuffer
	DataType() MPSDataType
	Device() metal.MTLDevice
	Length() uint
	VectorBytes() uint
	Vectors() uint
	Offset() uint

	// Topic: Instance Methods

	ResourceSize() uint
	SynchronizeOnCommandBuffer(commandBuffer metal.MTLCommandBuffer)
}

// Init initializes the instance.
func (v MPSVector) Init() MPSVector {
	rv := objc.Send[MPSVector](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v MPSVector) Autorelease() MPSVector {
	rv := objc.Send[MPSVector](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSVector creates a new MPSVector instance.
func NewMPSVector() MPSVector {
	class := getMPSVectorClass()
	rv := objc.Send[MPSVector](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVector/init(buffer:descriptor:)
func NewVectorWithBufferDescriptor(buffer metal.MTLBuffer, descriptor IMPSVectorDescriptor) MPSVector {
	instance := getMPSVectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBuffer:descriptor:"), buffer, descriptor)
	return MPSVectorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVector/init(buffer:offset:descriptor:)
func NewVectorWithBufferOffsetDescriptor(buffer metal.MTLBuffer, offset uint, descriptor IMPSVectorDescriptor) MPSVector {
	instance := getMPSVectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBuffer:offset:descriptor:"), buffer, offset, descriptor)
	return MPSVectorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVector/init(device:descriptor:)
func NewVectorWithDeviceDescriptor(device metal.MTLDevice, descriptor IMPSVectorDescriptor) MPSVector {
	instance := getMPSVectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	return MPSVectorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVector/init(buffer:descriptor:)
func (v MPSVector) InitWithBufferDescriptor(buffer metal.MTLBuffer, descriptor IMPSVectorDescriptor) MPSVector {
	rv := objc.Send[MPSVector](v.ID, objc.Sel("initWithBuffer:descriptor:"), buffer, descriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVector/init(buffer:offset:descriptor:)
func (v MPSVector) InitWithBufferOffsetDescriptor(buffer metal.MTLBuffer, offset uint, descriptor IMPSVectorDescriptor) MPSVector {
	rv := objc.Send[MPSVector](v.ID, objc.Sel("initWithBuffer:offset:descriptor:"), buffer, offset, descriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVector/init(device:descriptor:)
func (v MPSVector) InitWithDeviceDescriptor(device metal.MTLDevice, descriptor IMPSVectorDescriptor) MPSVector {
	rv := objc.Send[MPSVector](v.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVector/resourceSize()
func (v MPSVector) ResourceSize() uint {
	rv := objc.Send[uint](v.ID, objc.Sel("resourceSize"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVector/synchronize(on:)
func (v MPSVector) SynchronizeOnCommandBuffer(commandBuffer metal.MTLCommandBuffer) {
	objc.Send[objc.ID](v.ID, objc.Sel("synchronizeOnCommandBuffer:"), commandBuffer)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVector/data
func (v MPSVector) Data() metal.MTLBuffer {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("data"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVector/dataType
func (v MPSVector) DataType() MPSDataType {
	rv := objc.Send[MPSDataType](v.ID, objc.Sel("dataType"))
	return MPSDataType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVector/device
func (v MPSVector) Device() metal.MTLDevice {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("device"))
	return metal.MTLDeviceObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVector/length
func (v MPSVector) Length() uint {
	rv := objc.Send[uint](v.ID, objc.Sel("length"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVector/vectorBytes
func (v MPSVector) VectorBytes() uint {
	rv := objc.Send[uint](v.ID, objc.Sel("vectorBytes"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVector/vectors
func (v MPSVector) Vectors() uint {
	rv := objc.Send[uint](v.ID, objc.Sel("vectors"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVector/offset
func (v MPSVector) Offset() uint {
	rv := objc.Send[uint](v.ID, objc.Sel("offset"))
	return rv
}
