// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphVariableOp] class.
var (
	_MPSGraphVariableOpClass     MPSGraphVariableOpClass
	_MPSGraphVariableOpClassOnce sync.Once
)

func getMPSGraphVariableOpClass() MPSGraphVariableOpClass {
	_MPSGraphVariableOpClassOnce.Do(func() {
		_MPSGraphVariableOpClass = MPSGraphVariableOpClass{class: objc.GetClass("MPSGraphVariableOp")}
	})
	return _MPSGraphVariableOpClass
}

// GetMPSGraphVariableOpClass returns the class object for MPSGraphVariableOp.
func GetMPSGraphVariableOpClass() MPSGraphVariableOpClass {
	return getMPSGraphVariableOpClass()
}

type MPSGraphVariableOpClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphVariableOpClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphVariableOpClass) Alloc() MPSGraphVariableOp {
	rv := objc.Send[MPSGraphVariableOp](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The class that defines the parameters for a variable.
//
// # Instance Properties
//
//   - [MPSGraphVariableOp.DataType]: The data type of the variable.
//   - [MPSGraphVariableOp.Shape]: The shape of the variable.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphVariableOp
type MPSGraphVariableOp struct {
	MPSGraphOperation
}

// MPSGraphVariableOpFromID constructs a [MPSGraphVariableOp] from an objc.ID.
//
// The class that defines the parameters for a variable.
func MPSGraphVariableOpFromID(id objc.ID) MPSGraphVariableOp {
	return MPSGraphVariableOp{MPSGraphOperation: MPSGraphOperationFromID(id)}
}

// NOTE: MPSGraphVariableOp adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphVariableOp] class.
//
// # Instance Properties
//
//   - [IMPSGraphVariableOp.DataType]: The data type of the variable.
//   - [IMPSGraphVariableOp.Shape]: The shape of the variable.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphVariableOp
type IMPSGraphVariableOp interface {
	IMPSGraphOperation

	// Topic: Instance Properties

	// The data type of the variable.
	DataType() uint32
	// The shape of the variable.
	Shape() foundation.NSArray
}

// Init initializes the instance.
func (g MPSGraphVariableOp) Init() MPSGraphVariableOp {
	rv := objc.Send[MPSGraphVariableOp](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphVariableOp) Autorelease() MPSGraphVariableOp {
	rv := objc.Send[MPSGraphVariableOp](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphVariableOp creates a new MPSGraphVariableOp instance.
func NewMPSGraphVariableOp() MPSGraphVariableOp {
	class := getMPSGraphVariableOpClass()
	rv := objc.Send[MPSGraphVariableOp](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The data type of the variable.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphVariableOp/dataType
func (g MPSGraphVariableOp) DataType() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("dataType"))
	return rv
}

// The shape of the variable.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphVariableOp/shape
func (g MPSGraphVariableOp) Shape() foundation.NSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("shape"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
