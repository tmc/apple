// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLMLOdieEngine] class.
var (
	_CoreMLMLOdieEngineClass     CoreMLMLOdieEngineClass
	_CoreMLMLOdieEngineClassOnce sync.Once
)

func getCoreMLMLOdieEngineClass() CoreMLMLOdieEngineClass {
	_CoreMLMLOdieEngineClassOnce.Do(func() {
		_CoreMLMLOdieEngineClass = CoreMLMLOdieEngineClass{class: objc.GetClass("CoreML.MLOdieEngine")}
	})
	return _CoreMLMLOdieEngineClass
}

// GetCoreMLMLOdieEngineClass returns the class object for CoreML.MLOdieEngine.
func GetCoreMLMLOdieEngineClass() CoreMLMLOdieEngineClass {
	return getCoreMLMLOdieEngineClass()
}

type CoreMLMLOdieEngineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLMLOdieEngineClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLMLOdieEngineClass) Alloc() CoreMLMLOdieEngine {
	rv := objc.Send[CoreMLMLOdieEngine](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [CoreMLMLOdieEngine.NewRequestForModelInputFeaturesUsingStateOptionsError]
//   - [CoreMLMLOdieEngine.NewStateWithClientBuffers]
//   - [CoreMLMLOdieEngine.PredictionFromFeaturesUsingStateOptionsError]
// See: https://developer.apple.com/documentation/CoreML/CoreML.MLOdieEngine
type CoreMLMLOdieEngine struct {
	MLModelSwiftEngine
}

// CoreMLMLOdieEngineFromID constructs a [CoreMLMLOdieEngine] from an objc.ID.
func CoreMLMLOdieEngineFromID(id objc.ID) CoreMLMLOdieEngine {
	return CoreMLMLOdieEngine{MLModelSwiftEngine: MLModelSwiftEngineFromID(id)}
}
// Ensure CoreMLMLOdieEngine implements ICoreMLMLOdieEngine.
var _ ICoreMLMLOdieEngine = CoreMLMLOdieEngine{}

// An interface definition for the [CoreMLMLOdieEngine] class.
//
// # Methods
//
//   - [ICoreMLMLOdieEngine.NewRequestForModelInputFeaturesUsingStateOptionsError]
//   - [ICoreMLMLOdieEngine.NewStateWithClientBuffers]
//   - [ICoreMLMLOdieEngine.PredictionFromFeaturesUsingStateOptionsError]
//
// See: https://developer.apple.com/documentation/CoreML/CoreML.MLOdieEngine
type ICoreMLMLOdieEngine interface {
	IMLModelSwiftEngine

	// Topic: Methods

	NewRequestForModelInputFeaturesUsingStateOptionsError(model objectivec.IObject, features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	NewStateWithClientBuffers(buffers objectivec.IObject) objectivec.IObject
	PredictionFromFeaturesUsingStateOptionsError(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
}

// Init initializes the instance.
func (c CoreMLMLOdieEngine) Init() CoreMLMLOdieEngine {
	rv := objc.Send[CoreMLMLOdieEngine](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLMLOdieEngine) Autorelease() CoreMLMLOdieEngine {
	rv := objc.Send[CoreMLMLOdieEngine](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLMLOdieEngine creates a new CoreMLMLOdieEngine instance.
func NewCoreMLMLOdieEngine() CoreMLMLOdieEngine {
	class := getCoreMLMLOdieEngineClass()
	rv := objc.Send[CoreMLMLOdieEngine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/CoreML.MLOdieEngine/initWithDescription:configuration:
func NewCoreMLMLOdieEngineWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) CoreMLMLOdieEngine {
	instance := getCoreMLMLOdieEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return CoreMLMLOdieEngineFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewCoreMLMLOdieEngineWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) CoreMLMLOdieEngine {
	instance := getCoreMLMLOdieEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return CoreMLMLOdieEngineFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/CoreML.MLOdieEngine/newRequestForModel:inputFeatures:usingState:options:error:
func (c CoreMLMLOdieEngine) NewRequestForModelInputFeaturesUsingStateOptionsError(model objectivec.IObject, features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("newRequestForModel:inputFeatures:usingState:options:error:"), model, features, state, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/CoreML.MLOdieEngine/newStateWithClientBuffers:
func (c CoreMLMLOdieEngine) NewStateWithClientBuffers(buffers objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("newStateWithClientBuffers:"), buffers)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/CoreML.MLOdieEngine/predictionFromFeatures:usingState:options:error:
func (c CoreMLMLOdieEngine) PredictionFromFeaturesUsingStateOptionsError(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("predictionFromFeatures:usingState:options:error:"), features, state, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

