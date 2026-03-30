// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETVariablesDefinition] class.
var (
	_ETVariablesDefinitionClass     ETVariablesDefinitionClass
	_ETVariablesDefinitionClassOnce sync.Once
)

func getETVariablesDefinitionClass() ETVariablesDefinitionClass {
	_ETVariablesDefinitionClassOnce.Do(func() {
		_ETVariablesDefinitionClass = ETVariablesDefinitionClass{class: objc.GetClass("ETVariablesDefinition")}
	})
	return _ETVariablesDefinitionClass
}

// GetETVariablesDefinitionClass returns the class object for ETVariablesDefinition.
func GetETVariablesDefinitionClass() ETVariablesDefinitionClass {
	return getETVariablesDefinitionClass()
}

type ETVariablesDefinitionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETVariablesDefinitionClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETVariablesDefinitionClass) Alloc() ETVariablesDefinition {
	rv := objc.Send[ETVariablesDefinition](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ETVariablesDefinition.LayerNames]
//   - [ETVariablesDefinition.SetLayerNames]
//   - [ETVariablesDefinition.InitForLayersError]
//
// See: https://developer.apple.com/documentation/Espresso/ETVariablesDefinition
type ETVariablesDefinition struct {
	objectivec.Object
}

// ETVariablesDefinitionFromID constructs a [ETVariablesDefinition] from an objc.ID.
func ETVariablesDefinitionFromID(id objc.ID) ETVariablesDefinition {
	return ETVariablesDefinition{objectivec.Object{ID: id}}
}

// Ensure ETVariablesDefinition implements IETVariablesDefinition.
var _ IETVariablesDefinition = ETVariablesDefinition{}

// An interface definition for the [ETVariablesDefinition] class.
//
// # Methods
//
//   - [IETVariablesDefinition.LayerNames]
//   - [IETVariablesDefinition.SetLayerNames]
//   - [IETVariablesDefinition.InitForLayersError]
//
// See: https://developer.apple.com/documentation/Espresso/ETVariablesDefinition
type IETVariablesDefinition interface {
	objectivec.IObject

	// Topic: Methods

	LayerNames() foundation.INSArray
	SetLayerNames(value foundation.INSArray)
	InitForLayersError(layers objectivec.IObject) (ETVariablesDefinition, error)
}

// Init initializes the instance.
func (e ETVariablesDefinition) Init() ETVariablesDefinition {
	rv := objc.Send[ETVariablesDefinition](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETVariablesDefinition) Autorelease() ETVariablesDefinition {
	rv := objc.Send[ETVariablesDefinition](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETVariablesDefinition creates a new ETVariablesDefinition instance.
func NewETVariablesDefinition() ETVariablesDefinition {
	class := getETVariablesDefinitionClass()
	rv := objc.Send[ETVariablesDefinition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETVariablesDefinition/initForLayers:error:
func NewETVariablesDefinitionForLayersError(layers objectivec.IObject) (ETVariablesDefinition, error) {
	var errorPtr objc.ID
	instance := getETVariablesDefinitionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initForLayers:error:"), layers, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETVariablesDefinition{}, foundation.NSErrorFrom(errorPtr)
	}
	return ETVariablesDefinitionFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Espresso/ETVariablesDefinition/initForLayers:error:
func (e ETVariablesDefinition) InitForLayersError(layers objectivec.IObject) (ETVariablesDefinition, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("initForLayers:error:"), layers, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETVariablesDefinition{}, foundation.NSErrorFrom(errorPtr)
	}
	return ETVariablesDefinitionFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Espresso/ETVariablesDefinition/layerNames
func (e ETVariablesDefinition) LayerNames() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("layerNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e ETVariablesDefinition) SetLayerNames(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setLayerNames:"), value)
}
