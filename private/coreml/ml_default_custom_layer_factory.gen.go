// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLDefaultCustomLayerFactory] class.
var (
	_MLDefaultCustomLayerFactoryClass     MLDefaultCustomLayerFactoryClass
	_MLDefaultCustomLayerFactoryClassOnce sync.Once
)

func getMLDefaultCustomLayerFactoryClass() MLDefaultCustomLayerFactoryClass {
	_MLDefaultCustomLayerFactoryClassOnce.Do(func() {
		_MLDefaultCustomLayerFactoryClass = MLDefaultCustomLayerFactoryClass{class: objc.GetClass("MLDefaultCustomLayerFactory")}
	})
	return _MLDefaultCustomLayerFactoryClass
}

// GetMLDefaultCustomLayerFactoryClass returns the class object for MLDefaultCustomLayerFactory.
func GetMLDefaultCustomLayerFactoryClass() MLDefaultCustomLayerFactoryClass {
	return getMLDefaultCustomLayerFactoryClass()
}

type MLDefaultCustomLayerFactoryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLDefaultCustomLayerFactoryClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLDefaultCustomLayerFactoryClass) Alloc() MLDefaultCustomLayerFactory {
	rv := objc.Send[MLDefaultCustomLayerFactory](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLDefaultCustomLayerFactory.CreateCustomLayerWithParametersError]
// See: https://developer.apple.com/documentation/CoreML/MLDefaultCustomLayerFactory
type MLDefaultCustomLayerFactory struct {
	objectivec.Object
}

// MLDefaultCustomLayerFactoryFromID constructs a [MLDefaultCustomLayerFactory] from an objc.ID.
func MLDefaultCustomLayerFactoryFromID(id objc.ID) MLDefaultCustomLayerFactory {
	return MLDefaultCustomLayerFactory{objectivec.Object{ID: id}}
}
// Ensure MLDefaultCustomLayerFactory implements IMLDefaultCustomLayerFactory.
var _ IMLDefaultCustomLayerFactory = MLDefaultCustomLayerFactory{}

// An interface definition for the [MLDefaultCustomLayerFactory] class.
//
// # Methods
//
//   - [IMLDefaultCustomLayerFactory.CreateCustomLayerWithParametersError]
//
// See: https://developer.apple.com/documentation/CoreML/MLDefaultCustomLayerFactory
type IMLDefaultCustomLayerFactory interface {
	objectivec.IObject

	// Topic: Methods

	CreateCustomLayerWithParametersError(layer objectivec.IObject, parameters objectivec.IObject) (objectivec.IObject, error)
}

// Init initializes the instance.
func (d MLDefaultCustomLayerFactory) Init() MLDefaultCustomLayerFactory {
	rv := objc.Send[MLDefaultCustomLayerFactory](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d MLDefaultCustomLayerFactory) Autorelease() MLDefaultCustomLayerFactory {
	rv := objc.Send[MLDefaultCustomLayerFactory](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLDefaultCustomLayerFactory creates a new MLDefaultCustomLayerFactory instance.
func NewMLDefaultCustomLayerFactory() MLDefaultCustomLayerFactory {
	class := getMLDefaultCustomLayerFactoryClass()
	rv := objc.Send[MLDefaultCustomLayerFactory](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLDefaultCustomLayerFactory/createCustomLayer:withParameters:error:
func (d MLDefaultCustomLayerFactory) CreateCustomLayerWithParametersError(layer objectivec.IObject, parameters objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("createCustomLayer:withParameters:error:"), layer, parameters, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

