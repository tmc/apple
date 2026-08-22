// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLTensorAuxiliaryPlaneDescriptor] class.
var (
	_MTLTensorAuxiliaryPlaneDescriptorClass     MTLTensorAuxiliaryPlaneDescriptorClass
	_MTLTensorAuxiliaryPlaneDescriptorClassOnce sync.Once
)

func getMTLTensorAuxiliaryPlaneDescriptorClass() MTLTensorAuxiliaryPlaneDescriptorClass {
	_MTLTensorAuxiliaryPlaneDescriptorClassOnce.Do(func() {
		_MTLTensorAuxiliaryPlaneDescriptorClass = MTLTensorAuxiliaryPlaneDescriptorClass{class: objc.GetClass("MTLTensorAuxiliaryPlaneDescriptor")}
	})
	return _MTLTensorAuxiliaryPlaneDescriptorClass
}

// GetMTLTensorAuxiliaryPlaneDescriptorClass returns the class object for MTLTensorAuxiliaryPlaneDescriptor.
func GetMTLTensorAuxiliaryPlaneDescriptorClass() MTLTensorAuxiliaryPlaneDescriptorClass {
	return getMTLTensorAuxiliaryPlaneDescriptorClass()
}

type MTLTensorAuxiliaryPlaneDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLTensorAuxiliaryPlaneDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLTensorAuxiliaryPlaneDescriptorClass) Alloc() MTLTensorAuxiliaryPlaneDescriptor {
	rv := objc.Send[MTLTensorAuxiliaryPlaneDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A configuration for an auxiliary plane in a multi-plane tensor.
//
// # Overview
//
// Use this descriptor to configure an auxiliary plane’s data type and block
// factors before attaching it to a [MTLTensorDescriptor].
//
// # Instance Properties
//
//   - [MTLTensorAuxiliaryPlaneDescriptor.BlockFactors]: An extents instance that represents the number of data plane elements which correspond to one element in a plane you create with this descriptor.
//   - [MTLTensorAuxiliaryPlaneDescriptor.SetBlockFactors]
//   - [MTLTensorAuxiliaryPlaneDescriptor.DataType]: The data format of all elements in the plane.
//   - [MTLTensorAuxiliaryPlaneDescriptor.SetDataType]
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlaneDescriptor
type MTLTensorAuxiliaryPlaneDescriptor struct {
	objectivec.Object
}

// MTLTensorAuxiliaryPlaneDescriptorFromID constructs a [MTLTensorAuxiliaryPlaneDescriptor] from an objc.ID.
//
// A configuration for an auxiliary plane in a multi-plane tensor.
func MTLTensorAuxiliaryPlaneDescriptorFromID(id objc.ID) MTLTensorAuxiliaryPlaneDescriptor {
	return MTLTensorAuxiliaryPlaneDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MTLTensorAuxiliaryPlaneDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLTensorAuxiliaryPlaneDescriptor] class.
//
// # Instance Properties
//
//   - [IMTLTensorAuxiliaryPlaneDescriptor.BlockFactors]: An extents instance that represents the number of data plane elements which correspond to one element in a plane you create with this descriptor.
//   - [IMTLTensorAuxiliaryPlaneDescriptor.SetBlockFactors]
//   - [IMTLTensorAuxiliaryPlaneDescriptor.DataType]: The data format of all elements in the plane.
//   - [IMTLTensorAuxiliaryPlaneDescriptor.SetDataType]
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlaneDescriptor
type IMTLTensorAuxiliaryPlaneDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	// An extents instance that represents the number of data plane elements which correspond to one element in a plane you create with this descriptor.
	BlockFactors() IMTLTensorExtents
	SetBlockFactors(value IMTLTensorExtents)
	// The data format of all elements in the plane.
	DataType() MTLTensorDataType
	SetDataType(value MTLTensorDataType)
}

// Init initializes the instance.
func (t MTLTensorAuxiliaryPlaneDescriptor) Init() MTLTensorAuxiliaryPlaneDescriptor {
	rv := objc.Send[MTLTensorAuxiliaryPlaneDescriptor](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MTLTensorAuxiliaryPlaneDescriptor) Autorelease() MTLTensorAuxiliaryPlaneDescriptor {
	rv := objc.Send[MTLTensorAuxiliaryPlaneDescriptor](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLTensorAuxiliaryPlaneDescriptor creates a new MTLTensorAuxiliaryPlaneDescriptor instance.
func NewMTLTensorAuxiliaryPlaneDescriptor() MTLTensorAuxiliaryPlaneDescriptor {
	class := getMTLTensorAuxiliaryPlaneDescriptorClass()
	rv := objc.Send[MTLTensorAuxiliaryPlaneDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An extents instance that represents the number of data plane elements which
// correspond to one element in a plane you create with this descriptor.
//
// # Discussion
//
// The number of dimensions in the extents needs to match the number of the
// tensor’s dimensions.
//
// The first element of the block factors needs to be `32`. All remaining
// elements need to be `1`.
//
// The default value is a 1D block size of width `32`.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlaneDescriptor/blockFactors
func (t MTLTensorAuxiliaryPlaneDescriptor) BlockFactors() IMTLTensorExtents {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("blockFactors"))
	return MTLTensorExtentsFromID(objc.ID(rv))
}
func (t MTLTensorAuxiliaryPlaneDescriptor) SetBlockFactors(value IMTLTensorExtents) {
	objc.Send[struct{}](t.ID, objc.Sel("setBlockFactors:"), value)
}

// The data format of all elements in the plane.
//
// # Discussion
//
// The default value of this property is [MTLTensorDataType.metalFloat8ue8m0].
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlaneDescriptor/dataType
//
// [MTLTensorDataType.metalFloat8ue8m0]: https://developer.apple.com/documentation/Metal/MTLTensorDataType/metalFloat8ue8m0
func (t MTLTensorAuxiliaryPlaneDescriptor) DataType() MTLTensorDataType {
	rv := objc.Send[MTLTensorDataType](t.ID, objc.Sel("dataType"))
	return MTLTensorDataType(rv)
}
func (t MTLTensorAuxiliaryPlaneDescriptor) SetDataType(value MTLTensorDataType) {
	objc.Send[struct{}](t.ID, objc.Sel("setDataType:"), value)
}
