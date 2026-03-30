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
	rv := objc.Send[MLSVMLoader](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLSVMLoader.DebugDescription]
//   - [MLSVMLoader.Description]
//   - [MLSVMLoader.Hash]
//   - [MLSVMLoader.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLSVMLoader
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
//
// See: https://developer.apple.com/documentation/CoreML/MLSVMLoader
type IMLSVMLoader interface {
	objectivec.IObject

	// Topic: Methods

	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (s MLSVMLoader) Init() MLSVMLoader {
	rv := objc.Send[MLSVMLoader](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MLSVMLoader) Autorelease() MLSVMLoader {
	rv := objc.Send[MLSVMLoader](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSVMLoader creates a new MLSVMLoader instance.
func NewMLSVMLoader() MLSVMLoader {
	class := getMLSVMLoaderClass()
	rv := objc.Send[MLSVMLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSVMLoader/loadModelFromSpecification:configuration:error:
func (_MLSVMLoaderClass MLSVMLoaderClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLSVMLoaderClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLSVMLoader/debugDescription
func (s MLSVMLoader) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLSVMLoader/description
func (s MLSVMLoader) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLSVMLoader/hash
func (s MLSVMLoader) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSVMLoader/superclass
func (s MLSVMLoader) Superclass() objc.Class {
	rv := objc.Send[objc.Class](s.ID, objc.Sel("superclass"))
	return rv
}
