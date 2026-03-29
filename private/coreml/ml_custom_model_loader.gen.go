// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLCustomModelLoader] class.
var (
	_MLCustomModelLoaderClass     MLCustomModelLoaderClass
	_MLCustomModelLoaderClassOnce sync.Once
)

func getMLCustomModelLoaderClass() MLCustomModelLoaderClass {
	_MLCustomModelLoaderClassOnce.Do(func() {
		_MLCustomModelLoaderClass = MLCustomModelLoaderClass{class: objc.GetClass("MLCustomModelLoader")}
	})
	return _MLCustomModelLoaderClass
}

// GetMLCustomModelLoaderClass returns the class object for MLCustomModelLoader.
func GetMLCustomModelLoaderClass() MLCustomModelLoaderClass {
	return getMLCustomModelLoaderClass()
}

type MLCustomModelLoaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLCustomModelLoaderClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLCustomModelLoaderClass) Alloc() MLCustomModelLoader {
	rv := objc.Send[MLCustomModelLoader](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLCustomModelLoader.DebugDescription]
//   - [MLCustomModelLoader.Description]
//   - [MLCustomModelLoader.Hash]
//   - [MLCustomModelLoader.Superclass]
// See: https://developer.apple.com/documentation/CoreML/MLCustomModelLoader
type MLCustomModelLoader struct {
	objectivec.Object
}

// MLCustomModelLoaderFromID constructs a [MLCustomModelLoader] from an objc.ID.
func MLCustomModelLoaderFromID(id objc.ID) MLCustomModelLoader {
	return MLCustomModelLoader{objectivec.Object{ID: id}}
}
// Ensure MLCustomModelLoader implements IMLCustomModelLoader.
var _ IMLCustomModelLoader = MLCustomModelLoader{}

// An interface definition for the [MLCustomModelLoader] class.
//
// # Methods
//
//   - [IMLCustomModelLoader.DebugDescription]
//   - [IMLCustomModelLoader.Description]
//   - [IMLCustomModelLoader.Hash]
//   - [IMLCustomModelLoader.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomModelLoader
type IMLCustomModelLoader interface {
	objectivec.IObject

	// Topic: Methods

	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c MLCustomModelLoader) Init() MLCustomModelLoader {
	rv := objc.Send[MLCustomModelLoader](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLCustomModelLoader) Autorelease() MLCustomModelLoader {
	rv := objc.Send[MLCustomModelLoader](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLCustomModelLoader creates a new MLCustomModelLoader instance.
func NewMLCustomModelLoader() MLCustomModelLoader {
	class := getMLCustomModelLoaderClass()
	rv := objc.Send[MLCustomModelLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLCustomModelLoader/customModelWithName:modelDescription:modelConfiguration:parameterDictionary:error:
func (_MLCustomModelLoaderClass MLCustomModelLoaderClass) CustomModelWithNameModelDescriptionModelConfigurationParameterDictionaryError(name objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject, dictionary objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLCustomModelLoaderClass.class), objc.Sel("customModelWithName:modelDescription:modelConfiguration:parameterDictionary:error:"), name, description, configuration, dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomModelLoader/loadModelFromSpecification:configuration:error:
func (_MLCustomModelLoaderClass MLCustomModelLoaderClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLCustomModelLoaderClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomModelLoader/parametersFromCustomModelSpec:error:
func (_MLCustomModelLoaderClass MLCustomModelLoaderClass) ParametersFromCustomModelSpecError(spec unsafe.Pointer) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLCustomModelLoaderClass.class), objc.Sel("parametersFromCustomModelSpec:error:"), spec, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLCustomModelLoader/debugDescription
func (c MLCustomModelLoader) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLCustomModelLoader/description
func (c MLCustomModelLoader) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLCustomModelLoader/hash
func (c MLCustomModelLoader) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLCustomModelLoader/superclass
func (c MLCustomModelLoader) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}

