// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETLayerInitializationParameters] class.
var (
	_ETLayerInitializationParametersClass     ETLayerInitializationParametersClass
	_ETLayerInitializationParametersClassOnce sync.Once
)

func getETLayerInitializationParametersClass() ETLayerInitializationParametersClass {
	_ETLayerInitializationParametersClassOnce.Do(func() {
		_ETLayerInitializationParametersClass = ETLayerInitializationParametersClass{class: objc.GetClass("ETLayerInitializationParameters")}
	})
	return _ETLayerInitializationParametersClass
}

// GetETLayerInitializationParametersClass returns the class object for ETLayerInitializationParameters.
func GetETLayerInitializationParametersClass() ETLayerInitializationParametersClass {
	return getETLayerInitializationParametersClass()
}

type ETLayerInitializationParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETLayerInitializationParametersClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETLayerInitializationParametersClass) Alloc() ETLayerInitializationParameters {
	rv := objc.Send[ETLayerInitializationParameters](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ETLayerInitializationParameters.Parameters]
//   - [ETLayerInitializationParameters.InitWithModeParametersError]
//
// See: https://developer.apple.com/documentation/Espresso/ETLayerInitializationParameters
type ETLayerInitializationParameters struct {
	objectivec.Object
}

// ETLayerInitializationParametersFromID constructs a [ETLayerInitializationParameters] from an objc.ID.
func ETLayerInitializationParametersFromID(id objc.ID) ETLayerInitializationParameters {
	return ETLayerInitializationParameters{objectivec.Object{ID: id}}
}

// Ensure ETLayerInitializationParameters implements IETLayerInitializationParameters.
var _ IETLayerInitializationParameters = ETLayerInitializationParameters{}

// An interface definition for the [ETLayerInitializationParameters] class.
//
// # Methods
//
//   - [IETLayerInitializationParameters.Parameters]
//   - [IETLayerInitializationParameters.InitWithModeParametersError]
//
// See: https://developer.apple.com/documentation/Espresso/ETLayerInitializationParameters
type IETLayerInitializationParameters interface {
	objectivec.IObject

	// Topic: Methods

	Parameters() foundation.INSDictionary
	InitWithModeParametersError(mode uint64, parameters objectivec.IObject) (ETLayerInitializationParameters, error)
}

// Init initializes the instance.
func (e ETLayerInitializationParameters) Init() ETLayerInitializationParameters {
	rv := objc.Send[ETLayerInitializationParameters](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETLayerInitializationParameters) Autorelease() ETLayerInitializationParameters {
	rv := objc.Send[ETLayerInitializationParameters](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETLayerInitializationParameters creates a new ETLayerInitializationParameters instance.
func NewETLayerInitializationParameters() ETLayerInitializationParameters {
	class := getETLayerInitializationParametersClass()
	rv := objc.Send[ETLayerInitializationParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETLayerInitializationParameters/initWithMode:parameters:error:
func NewETLayerInitializationParametersWithModeParametersError(mode uint64, parameters objectivec.IObject) (ETLayerInitializationParameters, error) {
	var errorPtr objc.ID
	instance := getETLayerInitializationParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMode:parameters:error:"), mode, parameters, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETLayerInitializationParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return ETLayerInitializationParametersFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Espresso/ETLayerInitializationParameters/initWithMode:parameters:error:
func (e ETLayerInitializationParameters) InitWithModeParametersError(mode uint64, parameters objectivec.IObject) (ETLayerInitializationParameters, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("initWithMode:parameters:error:"), mode, parameters, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETLayerInitializationParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return ETLayerInitializationParametersFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Espresso/ETLayerInitializationParameters/parameters
func (e ETLayerInitializationParameters) Parameters() foundation.INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("parameters"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
