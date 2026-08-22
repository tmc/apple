// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSVectorDescriptor] class.
var (
	_MPSVectorDescriptorClass     MPSVectorDescriptorClass
	_MPSVectorDescriptorClassOnce sync.Once
)

func getMPSVectorDescriptorClass() MPSVectorDescriptorClass {
	_MPSVectorDescriptorClassOnce.Do(func() {
		_MPSVectorDescriptorClass = MPSVectorDescriptorClass{class: objc.GetClass("MPSVectorDescriptor")}
	})
	return _MPSVectorDescriptorClass
}

// GetMPSVectorDescriptorClass returns the class object for MPSVectorDescriptor.
func GetMPSVectorDescriptorClass() MPSVectorDescriptorClass {
	return getMPSVectorDescriptorClass()
}

type MPSVectorDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSVectorDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSVectorDescriptorClass) Alloc() MPSVectorDescriptor {
	rv := objc.Send[MPSVectorDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of the length and data type of a vector.
//
// # Instance Properties
//
//   - [MPSVectorDescriptor.DataType]
//   - [MPSVectorDescriptor.SetDataType]
//   - [MPSVectorDescriptor.Length]
//   - [MPSVectorDescriptor.SetLength]
//   - [MPSVectorDescriptor.VectorBytes]
//   - [MPSVectorDescriptor.Vectors]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVectorDescriptor
type MPSVectorDescriptor struct {
	objectivec.Object
}

// MPSVectorDescriptorFromID constructs a [MPSVectorDescriptor] from an objc.ID.
//
// A description of the length and data type of a vector.
func MPSVectorDescriptorFromID(id objc.ID) MPSVectorDescriptor {
	return MPSVectorDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MPSVectorDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSVectorDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSVectorDescriptor.DataType]
//   - [IMPSVectorDescriptor.SetDataType]
//   - [IMPSVectorDescriptor.Length]
//   - [IMPSVectorDescriptor.SetLength]
//   - [IMPSVectorDescriptor.VectorBytes]
//   - [IMPSVectorDescriptor.Vectors]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVectorDescriptor
type IMPSVectorDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	DataType() MPSDataType
	SetDataType(value MPSDataType)
	Length() uint
	SetLength(value uint)
	VectorBytes() uint
	Vectors() uint
}

// Init initializes the instance.
func (v MPSVectorDescriptor) Init() MPSVectorDescriptor {
	rv := objc.Send[MPSVectorDescriptor](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v MPSVectorDescriptor) Autorelease() MPSVectorDescriptor {
	rv := objc.Send[MPSVectorDescriptor](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSVectorDescriptor creates a new MPSVectorDescriptor instance.
func NewMPSVectorDescriptor() MPSVectorDescriptor {
	class := getMPSVectorDescriptorClass()
	rv := objc.Send[MPSVectorDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVectorDescriptor/init(length:dataType:)
func NewVectorDescriptorWithLengthDataType(length uint, dataType MPSDataType) MPSVectorDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSVectorDescriptorClass().class), objc.Sel("vectorDescriptorWithLength:dataType:"), length, dataType)
	return MPSVectorDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVectorDescriptor/init(length:vectors:vectorBytes:dataType:)
func NewVectorDescriptorWithLengthVectorsVectorBytesDataType(length uint, vectors uint, vectorBytes uint, dataType MPSDataType) MPSVectorDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSVectorDescriptorClass().class), objc.Sel("vectorDescriptorWithLength:vectors:vectorBytes:dataType:"), length, vectors, vectorBytes, dataType)
	return MPSVectorDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVectorDescriptor/vectorBytes(forLength:dataType:)
func (_MPSVectorDescriptorClass MPSVectorDescriptorClass) VectorBytesForLengthDataType(length uint, dataType MPSDataType) uintptr {
	rv := objc.Send[uintptr](objc.ID(_MPSVectorDescriptorClass.class), objc.Sel("vectorBytesForLength:dataType:"), length, dataType)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVectorDescriptor/dataType
func (v MPSVectorDescriptor) DataType() MPSDataType {
	rv := objc.Send[MPSDataType](v.ID, objc.Sel("dataType"))
	return MPSDataType(rv)
}
func (v MPSVectorDescriptor) SetDataType(value MPSDataType) {
	objc.Send[struct{}](v.ID, objc.Sel("setDataType:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVectorDescriptor/length
func (v MPSVectorDescriptor) Length() uint {
	rv := objc.Send[uint](v.ID, objc.Sel("length"))
	return rv
}
func (v MPSVectorDescriptor) SetLength(value uint) {
	objc.Send[struct{}](v.ID, objc.Sel("setLength:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVectorDescriptor/vectorBytes
func (v MPSVectorDescriptor) VectorBytes() uint {
	rv := objc.Send[uint](v.ID, objc.Sel("vectorBytes"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVectorDescriptor/vectors
func (v MPSVectorDescriptor) Vectors() uint {
	rv := objc.Send[uint](v.ID, objc.Sel("vectors"))
	return rv
}
