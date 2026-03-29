// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSVRLoader] class.
var (
	_MLSVRLoaderClass     MLSVRLoaderClass
	_MLSVRLoaderClassOnce sync.Once
)

func getMLSVRLoaderClass() MLSVRLoaderClass {
	_MLSVRLoaderClassOnce.Do(func() {
		_MLSVRLoaderClass = MLSVRLoaderClass{class: objc.GetClass("MLSVRLoader")}
	})
	return _MLSVRLoaderClass
}

// GetMLSVRLoaderClass returns the class object for MLSVRLoader.
func GetMLSVRLoaderClass() MLSVRLoaderClass {
	return getMLSVRLoaderClass()
}

type MLSVRLoaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSVRLoaderClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSVRLoaderClass) Alloc() MLSVRLoader {
	rv := objc.Send[MLSVRLoader](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLSVRLoader.DebugDescription]
//   - [MLSVRLoader.Description]
//   - [MLSVRLoader.Hash]
//   - [MLSVRLoader.Superclass]
// See: https://developer.apple.com/documentation/CoreML/MLSVRLoader
type MLSVRLoader struct {
	objectivec.Object
}

// MLSVRLoaderFromID constructs a [MLSVRLoader] from an objc.ID.
func MLSVRLoaderFromID(id objc.ID) MLSVRLoader {
	return MLSVRLoader{objectivec.Object{ID: id}}
}
// Ensure MLSVRLoader implements IMLSVRLoader.
var _ IMLSVRLoader = MLSVRLoader{}

// An interface definition for the [MLSVRLoader] class.
//
// # Methods
//
//   - [IMLSVRLoader.DebugDescription]
//   - [IMLSVRLoader.Description]
//   - [IMLSVRLoader.Hash]
//   - [IMLSVRLoader.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLSVRLoader
type IMLSVRLoader interface {
	objectivec.IObject

	// Topic: Methods

	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (s MLSVRLoader) Init() MLSVRLoader {
	rv := objc.Send[MLSVRLoader](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MLSVRLoader) Autorelease() MLSVRLoader {
	rv := objc.Send[MLSVRLoader](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSVRLoader creates a new MLSVRLoader instance.
func NewMLSVRLoader() MLSVRLoader {
	class := getMLSVRLoaderClass()
	rv := objc.Send[MLSVRLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLSVRLoader/loadModelFromSpecification:configuration:error:
func (_MLSVRLoaderClass MLSVRLoaderClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLSVRLoaderClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLSVRLoader/debugDescription
func (s MLSVRLoader) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLSVRLoader/description
func (s MLSVRLoader) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLSVRLoader/hash
func (s MLSVRLoader) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLSVRLoader/superclass
func (s MLSVRLoader) Superclass() objc.Class {
	rv := objc.Send[objc.Class](s.ID, objc.Sel("superclass"))
	return rv
}

