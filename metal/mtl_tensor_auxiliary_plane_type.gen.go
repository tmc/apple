// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLTensorAuxiliaryPlaneType] class.
var (
	_MTLTensorAuxiliaryPlaneTypeClass     MTLTensorAuxiliaryPlaneTypeClass
	_MTLTensorAuxiliaryPlaneTypeClassOnce sync.Once
)

func getMTLTensorAuxiliaryPlaneTypeClass() MTLTensorAuxiliaryPlaneTypeClass {
	_MTLTensorAuxiliaryPlaneTypeClassOnce.Do(func() {
		_MTLTensorAuxiliaryPlaneTypeClass = MTLTensorAuxiliaryPlaneTypeClass{class: objc.GetClass("MTLTensorAuxiliaryPlaneType")}
	})
	return _MTLTensorAuxiliaryPlaneTypeClass
}

// GetMTLTensorAuxiliaryPlaneTypeClass returns the class object for MTLTensorAuxiliaryPlaneType.
func GetMTLTensorAuxiliaryPlaneTypeClass() MTLTensorAuxiliaryPlaneTypeClass {
	return getMTLTensorAuxiliaryPlaneTypeClass()
}

type MTLTensorAuxiliaryPlaneTypeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLTensorAuxiliaryPlaneTypeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLTensorAuxiliaryPlaneTypeClass) Alloc() MTLTensorAuxiliaryPlaneType {
	rv := objc.Send[MTLTensorAuxiliaryPlaneType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An auxiliary plane that a shader’s tensor argument requires.
//
// # Instance Properties
//
//   - [MTLTensorAuxiliaryPlaneType.BlockFactors]: The number of data plane elements that correspond to one element in this plane.
//   - [MTLTensorAuxiliaryPlaneType.DataType]: The data format of all elements in the plane.
//   - [MTLTensorAuxiliaryPlaneType.PlaneType]: The type of information this plane stores.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlaneType
type MTLTensorAuxiliaryPlaneType struct {
	objectivec.Object
}

// MTLTensorAuxiliaryPlaneTypeFromID constructs a [MTLTensorAuxiliaryPlaneType] from an objc.ID.
//
// An auxiliary plane that a shader’s tensor argument requires.
func MTLTensorAuxiliaryPlaneTypeFromID(id objc.ID) MTLTensorAuxiliaryPlaneType {
	return MTLTensorAuxiliaryPlaneType{objectivec.Object{ID: id}}
}

// NOTE: MTLTensorAuxiliaryPlaneType adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLTensorAuxiliaryPlaneType] class.
//
// # Instance Properties
//
//   - [IMTLTensorAuxiliaryPlaneType.BlockFactors]: The number of data plane elements that correspond to one element in this plane.
//   - [IMTLTensorAuxiliaryPlaneType.DataType]: The data format of all elements in the plane.
//   - [IMTLTensorAuxiliaryPlaneType.PlaneType]: The type of information this plane stores.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlaneType
type IMTLTensorAuxiliaryPlaneType interface {
	objectivec.IObject

	// Topic: Instance Properties

	// The number of data plane elements that correspond to one element in this plane.
	BlockFactors() IMTLTensorExtents
	// The data format of all elements in the plane.
	DataType() MTLTensorDataType
	// The type of information this plane stores.
	PlaneType() MTLTensorPlaneType
}

// Init initializes the instance.
func (t MTLTensorAuxiliaryPlaneType) Init() MTLTensorAuxiliaryPlaneType {
	rv := objc.Send[MTLTensorAuxiliaryPlaneType](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MTLTensorAuxiliaryPlaneType) Autorelease() MTLTensorAuxiliaryPlaneType {
	rv := objc.Send[MTLTensorAuxiliaryPlaneType](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLTensorAuxiliaryPlaneType creates a new MTLTensorAuxiliaryPlaneType instance.
func NewMTLTensorAuxiliaryPlaneType() MTLTensorAuxiliaryPlaneType {
	class := getMTLTensorAuxiliaryPlaneTypeClass()
	rv := objc.Send[MTLTensorAuxiliaryPlaneType](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The number of data plane elements that correspond to one element in this
// plane.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlaneType/blockFactors
func (t MTLTensorAuxiliaryPlaneType) BlockFactors() IMTLTensorExtents {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("blockFactors"))
	return MTLTensorExtentsFromID(objc.ID(rv))
}

// The data format of all elements in the plane.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlaneType/dataType
func (t MTLTensorAuxiliaryPlaneType) DataType() MTLTensorDataType {
	rv := objc.Send[MTLTensorDataType](t.ID, objc.Sel("dataType"))
	return MTLTensorDataType(rv)
}

// The type of information this plane stores.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlaneType/planeType
func (t MTLTensorAuxiliaryPlaneType) PlaneType() MTLTensorPlaneType {
	rv := objc.Send[MTLTensorPlaneType](t.ID, objc.Sel("planeType"))
	return MTLTensorPlaneType(rv)
}
