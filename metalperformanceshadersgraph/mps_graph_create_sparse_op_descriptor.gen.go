// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphCreateSparseOpDescriptor] class.
var (
	_MPSGraphCreateSparseOpDescriptorClass     MPSGraphCreateSparseOpDescriptorClass
	_MPSGraphCreateSparseOpDescriptorClassOnce sync.Once
)

func getMPSGraphCreateSparseOpDescriptorClass() MPSGraphCreateSparseOpDescriptorClass {
	_MPSGraphCreateSparseOpDescriptorClassOnce.Do(func() {
		_MPSGraphCreateSparseOpDescriptorClass = MPSGraphCreateSparseOpDescriptorClass{class: objc.GetClass("MPSGraphCreateSparseOpDescriptor")}
	})
	return _MPSGraphCreateSparseOpDescriptorClass
}

// GetMPSGraphCreateSparseOpDescriptorClass returns the class object for MPSGraphCreateSparseOpDescriptor.
func GetMPSGraphCreateSparseOpDescriptorClass() MPSGraphCreateSparseOpDescriptorClass {
	return getMPSGraphCreateSparseOpDescriptorClass()
}

type MPSGraphCreateSparseOpDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphCreateSparseOpDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphCreateSparseOpDescriptorClass) Alloc() MPSGraphCreateSparseOpDescriptor {
	rv := objc.Send[MPSGraphCreateSparseOpDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class that describes the properties of a create sparse operation.
//
// # Instance Properties
//
//   - [MPSGraphCreateSparseOpDescriptor.DataType]: Defines the datatype of the sparse tensor.
//   - [MPSGraphCreateSparseOpDescriptor.SetDataType]
//   - [MPSGraphCreateSparseOpDescriptor.SparseStorageType]: Defines the storage format of the sparse tensor.
//   - [MPSGraphCreateSparseOpDescriptor.SetSparseStorageType]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCreateSparseOpDescriptor
type MPSGraphCreateSparseOpDescriptor struct {
	MPSGraphObject
}

// MPSGraphCreateSparseOpDescriptorFromID constructs a [MPSGraphCreateSparseOpDescriptor] from an objc.ID.
//
// A class that describes the properties of a create sparse operation.
func MPSGraphCreateSparseOpDescriptorFromID(id objc.ID) MPSGraphCreateSparseOpDescriptor {
	return MPSGraphCreateSparseOpDescriptor{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphCreateSparseOpDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphCreateSparseOpDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSGraphCreateSparseOpDescriptor.DataType]: Defines the datatype of the sparse tensor.
//   - [IMPSGraphCreateSparseOpDescriptor.SetDataType]
//   - [IMPSGraphCreateSparseOpDescriptor.SparseStorageType]: Defines the storage format of the sparse tensor.
//   - [IMPSGraphCreateSparseOpDescriptor.SetSparseStorageType]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCreateSparseOpDescriptor
type IMPSGraphCreateSparseOpDescriptor interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// Defines the datatype of the sparse tensor.
	DataType() uint32
	SetDataType(value uint32)
	// Defines the storage format of the sparse tensor.
	SparseStorageType() MPSGraphSparseStorageType
	SetSparseStorageType(value MPSGraphSparseStorageType)
}

// Init initializes the instance.
func (g MPSGraphCreateSparseOpDescriptor) Init() MPSGraphCreateSparseOpDescriptor {
	rv := objc.Send[MPSGraphCreateSparseOpDescriptor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphCreateSparseOpDescriptor) Autorelease() MPSGraphCreateSparseOpDescriptor {
	rv := objc.Send[MPSGraphCreateSparseOpDescriptor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphCreateSparseOpDescriptor creates a new MPSGraphCreateSparseOpDescriptor instance.
func NewMPSGraphCreateSparseOpDescriptor() MPSGraphCreateSparseOpDescriptor {
	class := getMPSGraphCreateSparseOpDescriptorClass()
	rv := objc.Send[MPSGraphCreateSparseOpDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a descriptor for a sparse tensor.
//
// sparseStorageType: A sparseStorageType.
//
// dataType: A dataType of the sparse tensor.
//
// # Return Value
//
// The descriptor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCreateSparseOpDescriptor/sparseDescriptor(descriptorWithStorageType:dataType:)
func (_MPSGraphCreateSparseOpDescriptorClass MPSGraphCreateSparseOpDescriptorClass) DescriptorWithStorageTypeDataType(sparseStorageType MPSGraphSparseStorageType, dataType uint32) MPSGraphCreateSparseOpDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MPSGraphCreateSparseOpDescriptorClass.class), objc.Sel("descriptorWithStorageType:dataType:"), sparseStorageType, dataType)
	return MPSGraphCreateSparseOpDescriptorFromID(rv)
}

// Defines the datatype of the sparse tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCreateSparseOpDescriptor/dataType
func (g MPSGraphCreateSparseOpDescriptor) DataType() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("dataType"))
	return rv
}
func (g MPSGraphCreateSparseOpDescriptor) SetDataType(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setDataType:"), value)
}

// Defines the storage format of the sparse tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCreateSparseOpDescriptor/sparseStorageType
func (g MPSGraphCreateSparseOpDescriptor) SparseStorageType() MPSGraphSparseStorageType {
	rv := objc.Send[MPSGraphSparseStorageType](g.ID, objc.Sel("sparseStorageType"))
	return MPSGraphSparseStorageType(rv)
}
func (g MPSGraphCreateSparseOpDescriptor) SetSparseStorageType(value MPSGraphSparseStorageType) {
	objc.Send[struct{}](g.ID, objc.Sel("setSparseStorageType:"), value)
}
