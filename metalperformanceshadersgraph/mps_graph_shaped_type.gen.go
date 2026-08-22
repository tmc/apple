// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphShapedType] class.
var (
	_MPSGraphShapedTypeClass     MPSGraphShapedTypeClass
	_MPSGraphShapedTypeClassOnce sync.Once
)

func getMPSGraphShapedTypeClass() MPSGraphShapedTypeClass {
	_MPSGraphShapedTypeClassOnce.Do(func() {
		_MPSGraphShapedTypeClass = MPSGraphShapedTypeClass{class: objc.GetClass("MPSGraphShapedType")}
	})
	return _MPSGraphShapedTypeClass
}

// GetMPSGraphShapedTypeClass returns the class object for MPSGraphShapedType.
func GetMPSGraphShapedTypeClass() MPSGraphShapedTypeClass {
	return getMPSGraphShapedTypeClass()
}

type MPSGraphShapedTypeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphShapedTypeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphShapedTypeClass) Alloc() MPSGraphShapedType {
	rv := objc.Send[MPSGraphShapedType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The shaped type class for types on tensors with a shape and data type.
//
// # Initializers
//
//   - [MPSGraphShapedType.InitWithShapeDataType]: Initializes a shaped type.
//
// # Instance Properties
//
//   - [MPSGraphShapedType.DataType]: The data type of the shaped type.
//   - [MPSGraphShapedType.SetDataType]
//   - [MPSGraphShapedType.Shape]: The Shape of the shaped type.
//   - [MPSGraphShapedType.SetShape]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType
type MPSGraphShapedType struct {
	MPSGraphType
}

// MPSGraphShapedTypeFromID constructs a [MPSGraphShapedType] from an objc.ID.
//
// The shaped type class for types on tensors with a shape and data type.
func MPSGraphShapedTypeFromID(id objc.ID) MPSGraphShapedType {
	return MPSGraphShapedType{MPSGraphType: MPSGraphTypeFromID(id)}
}

// NOTE: MPSGraphShapedType adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphShapedType] class.
//
// # Initializers
//
//   - [IMPSGraphShapedType.InitWithShapeDataType]: Initializes a shaped type.
//
// # Instance Properties
//
//   - [IMPSGraphShapedType.DataType]: The data type of the shaped type.
//   - [IMPSGraphShapedType.SetDataType]
//   - [IMPSGraphShapedType.Shape]: The Shape of the shaped type.
//   - [IMPSGraphShapedType.SetShape]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType
type IMPSGraphShapedType interface {
	IMPSGraphType

	// Topic: Initializers

	// Initializes a shaped type.
	InitWithShapeDataType(shape foundation.NSArray, dataType uint32) MPSGraphShapedType

	// Topic: Instance Properties

	// The data type of the shaped type.
	DataType() uint32
	SetDataType(value uint32)
	// The Shape of the shaped type.
	Shape() foundation.NSArray
	SetShape(value foundation.NSArray)
}

// Init initializes the instance.
func (g MPSGraphShapedType) Init() MPSGraphShapedType {
	rv := objc.Send[MPSGraphShapedType](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphShapedType) Autorelease() MPSGraphShapedType {
	rv := objc.Send[MPSGraphShapedType](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphShapedType creates a new MPSGraphShapedType instance.
func NewMPSGraphShapedType() MPSGraphShapedType {
	class := getMPSGraphShapedTypeClass()
	rv := objc.Send[MPSGraphShapedType](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a shaped type.
//
// shape: The shape of the shaped type.
//
// dataType: The dataType of the shaped type.
//
// # Return Value
//
// A valid MPSGraphShapedType, or nil if allocation failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType/init(shape:dataType:)
func NewGraphShapedTypeWithShapeDataType(shape foundation.NSArray, dataType uint32) MPSGraphShapedType {
	instance := getMPSGraphShapedTypeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithShape:dataType:"), shape, dataType)
	return MPSGraphShapedTypeFromID(rv)
}

// Initializes a shaped type.
//
// shape: The shape of the shaped type.
//
// dataType: The dataType of the shaped type.
//
// # Return Value
//
// A valid MPSGraphShapedType, or nil if allocation failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType/init(shape:dataType:)
func (g MPSGraphShapedType) InitWithShapeDataType(shape foundation.NSArray, dataType uint32) MPSGraphShapedType {
	rv := objc.Send[MPSGraphShapedType](g.ID, objc.Sel("initWithShape:dataType:"), shape, dataType)
	return rv
}

// The data type of the shaped type.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType/dataType
func (g MPSGraphShapedType) DataType() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("dataType"))
	return rv
}
func (g MPSGraphShapedType) SetDataType(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setDataType:"), value)
}

// The Shape of the shaped type.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphShapedType/shape
func (g MPSGraphShapedType) Shape() foundation.NSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("shape"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g MPSGraphShapedType) SetShape(value foundation.NSArray) {
	objc.Send[struct{}](g.ID, objc.Sel("setShape:"), value)
}
