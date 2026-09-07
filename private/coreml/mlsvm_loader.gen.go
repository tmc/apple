// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSVMLoader] class.
var (
	_MLSVMLoaderClass     MLSVMLoaderClass
	_MLSVMLoaderClassOnce sync.Once
)

func getMLSVMLoaderClass() MLSVMLoaderClass {
	_MLSVMLoaderClassOnce.Do(func() {
		_MLSVMLoaderClass = MLSVMLoaderClass{class: objc.GetClass("MLSVMLoader")}
	})
	return _MLSVMLoaderClass
}

// GetMLSVMLoaderClass returns the class object for MLSVMLoader.
func GetMLSVMLoaderClass() MLSVMLoaderClass {
	return getMLSVMLoaderClass()
}

type MLSVMLoaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSVMLoaderClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSVMLoaderClass) Alloc() MLSVMLoader {
	rv := objc.SendIfResponds[MLSVMLoader](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLSVMLoader.DebugDescription]
//   - [MLSVMLoader.Description]
//   - [MLSVMLoader.Hash]
//   - [MLSVMLoader.Superclass]
type MLSVMLoader struct {
	objectivec.Object
}

// MLSVMLoaderFromID constructs a [MLSVMLoader] from an objc.ID.
func MLSVMLoaderFromID(id objc.ID) MLSVMLoader {
	return MLSVMLoader{objectivec.Object{ID: id}}
}

// Ensure MLSVMLoader implements IMLSVMLoader.
var _ IMLSVMLoader = MLSVMLoader{}

// An interface definition for the [MLSVMLoader] class.
//
// # Methods
//
//   - [IMLSVMLoader.DebugDescription]
//   - [IMLSVMLoader.Description]
//   - [IMLSVMLoader.Hash]
//   - [IMLSVMLoader.Superclass]
type IMLSVMLoader interface {
	objectivec.IObject

	// Topic: Methods

	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLSVMLoader) Init() MLSVMLoader {
	rv := objc.SendIfResponds[MLSVMLoader](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLSVMLoader) Autorelease() MLSVMLoader {
	rv := objc.SendIfResponds[MLSVMLoader](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSVMLoader creates a new MLSVMLoader instance.
func NewMLSVMLoader() MLSVMLoader {
	class := getMLSVMLoaderClass()
	rv := objc.SendIfResponds[MLSVMLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_MLSVMLoaderClass MLSVMLoaderClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLSVMLoaderClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

func (m MLSVMLoader) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLSVMLoader) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLSVMLoader) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLSVMLoader) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
