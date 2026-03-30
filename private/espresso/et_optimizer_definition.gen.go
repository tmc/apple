// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETOptimizerDefinition] class.
var (
	_ETOptimizerDefinitionClass     ETOptimizerDefinitionClass
	_ETOptimizerDefinitionClassOnce sync.Once
)

func getETOptimizerDefinitionClass() ETOptimizerDefinitionClass {
	_ETOptimizerDefinitionClassOnce.Do(func() {
		_ETOptimizerDefinitionClass = ETOptimizerDefinitionClass{class: objc.GetClass("ETOptimizerDefinition")}
	})
	return _ETOptimizerDefinitionClass
}

// GetETOptimizerDefinitionClass returns the class object for ETOptimizerDefinition.
func GetETOptimizerDefinitionClass() ETOptimizerDefinitionClass {
	return getETOptimizerDefinitionClass()
}

type ETOptimizerDefinitionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETOptimizerDefinitionClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETOptimizerDefinitionClass) Alloc() ETOptimizerDefinition {
	rv := objc.Send[ETOptimizerDefinition](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ETOptimizerDefinition.OptimizationParameters]
//   - [ETOptimizerDefinition.Type]
//   - [ETOptimizerDefinition.InitWithOptimizationAlgorithmParametersError]
//
// See: https://developer.apple.com/documentation/Espresso/ETOptimizerDefinition
type ETOptimizerDefinition struct {
	objectivec.Object
}

// ETOptimizerDefinitionFromID constructs a [ETOptimizerDefinition] from an objc.ID.
func ETOptimizerDefinitionFromID(id objc.ID) ETOptimizerDefinition {
	return ETOptimizerDefinition{objectivec.Object{ID: id}}
}

// Ensure ETOptimizerDefinition implements IETOptimizerDefinition.
var _ IETOptimizerDefinition = ETOptimizerDefinition{}

// An interface definition for the [ETOptimizerDefinition] class.
//
// # Methods
//
//   - [IETOptimizerDefinition.OptimizationParameters]
//   - [IETOptimizerDefinition.Type]
//   - [IETOptimizerDefinition.InitWithOptimizationAlgorithmParametersError]
//
// See: https://developer.apple.com/documentation/Espresso/ETOptimizerDefinition
type IETOptimizerDefinition interface {
	objectivec.IObject

	// Topic: Methods

	OptimizationParameters() foundation.INSDictionary
	Type() int64
	InitWithOptimizationAlgorithmParametersError(algorithm int64, parameters objectivec.IObject) (ETOptimizerDefinition, error)
}

// Init initializes the instance.
func (e ETOptimizerDefinition) Init() ETOptimizerDefinition {
	rv := objc.Send[ETOptimizerDefinition](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETOptimizerDefinition) Autorelease() ETOptimizerDefinition {
	rv := objc.Send[ETOptimizerDefinition](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETOptimizerDefinition creates a new ETOptimizerDefinition instance.
func NewETOptimizerDefinition() ETOptimizerDefinition {
	class := getETOptimizerDefinitionClass()
	rv := objc.Send[ETOptimizerDefinition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETOptimizerDefinition/initWithOptimizationAlgorithm:parameters:error:
func NewETOptimizerDefinitionWithOptimizationAlgorithmParametersError(algorithm int64, parameters objectivec.IObject) (ETOptimizerDefinition, error) {
	var errorPtr objc.ID
	instance := getETOptimizerDefinitionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOptimizationAlgorithm:parameters:error:"), algorithm, parameters, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETOptimizerDefinition{}, foundation.NSErrorFrom(errorPtr)
	}
	return ETOptimizerDefinitionFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Espresso/ETOptimizerDefinition/initWithOptimizationAlgorithm:parameters:error:
func (e ETOptimizerDefinition) InitWithOptimizationAlgorithmParametersError(algorithm int64, parameters objectivec.IObject) (ETOptimizerDefinition, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("initWithOptimizationAlgorithm:parameters:error:"), algorithm, parameters, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETOptimizerDefinition{}, foundation.NSErrorFrom(errorPtr)
	}
	return ETOptimizerDefinitionFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Espresso/ETOptimizerDefinition/optimizationParameters
func (e ETOptimizerDefinition) OptimizationParameters() foundation.INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("optimizationParameters"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Espresso/ETOptimizerDefinition/type
func (e ETOptimizerDefinition) Type() int64 {
	rv := objc.Send[int64](e.ID, objc.Sel("type"))
	return rv
}
