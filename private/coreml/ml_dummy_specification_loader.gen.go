// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLDummySpecificationLoader] class.
var (
	_MLDummySpecificationLoaderClass     MLDummySpecificationLoaderClass
	_MLDummySpecificationLoaderClassOnce sync.Once
)

func getMLDummySpecificationLoaderClass() MLDummySpecificationLoaderClass {
	_MLDummySpecificationLoaderClassOnce.Do(func() {
		_MLDummySpecificationLoaderClass = MLDummySpecificationLoaderClass{class: objc.GetClass("MLDummySpecificationLoader")}
	})
	return _MLDummySpecificationLoaderClass
}

// GetMLDummySpecificationLoaderClass returns the class object for MLDummySpecificationLoader.
func GetMLDummySpecificationLoaderClass() MLDummySpecificationLoaderClass {
	return getMLDummySpecificationLoaderClass()
}

type MLDummySpecificationLoaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLDummySpecificationLoaderClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLDummySpecificationLoaderClass) Alloc() MLDummySpecificationLoader {
	rv := objc.SendIfResponds[MLDummySpecificationLoader](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLDummySpecificationLoader.DebugDescription]
//   - [MLDummySpecificationLoader.Description]
//   - [MLDummySpecificationLoader.Hash]
//   - [MLDummySpecificationLoader.Superclass]
type MLDummySpecificationLoader struct {
	objectivec.Object
}

// MLDummySpecificationLoaderFromID constructs a [MLDummySpecificationLoader] from an objc.ID.
func MLDummySpecificationLoaderFromID(id objc.ID) MLDummySpecificationLoader {
	return MLDummySpecificationLoader{objectivec.Object{ID: id}}
}

// Ensure MLDummySpecificationLoader implements IMLDummySpecificationLoader.
var _ IMLDummySpecificationLoader = MLDummySpecificationLoader{}

// An interface definition for the [MLDummySpecificationLoader] class.
//
// # Methods
//
//   - [IMLDummySpecificationLoader.DebugDescription]
//   - [IMLDummySpecificationLoader.Description]
//   - [IMLDummySpecificationLoader.Hash]
//   - [IMLDummySpecificationLoader.Superclass]
type IMLDummySpecificationLoader interface {
	objectivec.IObject

	// Topic: Methods

	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLDummySpecificationLoader) Init() MLDummySpecificationLoader {
	rv := objc.SendIfResponds[MLDummySpecificationLoader](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLDummySpecificationLoader) Autorelease() MLDummySpecificationLoader {
	rv := objc.SendIfResponds[MLDummySpecificationLoader](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLDummySpecificationLoader creates a new MLDummySpecificationLoader instance.
func NewMLDummySpecificationLoader() MLDummySpecificationLoader {
	class := getMLDummySpecificationLoaderClass()
	rv := objc.SendIfResponds[MLDummySpecificationLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_MLDummySpecificationLoaderClass MLDummySpecificationLoaderClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLDummySpecificationLoaderClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

func (m MLDummySpecificationLoader) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLDummySpecificationLoader) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLDummySpecificationLoader) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLDummySpecificationLoader) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
