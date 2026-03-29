// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNeuralNetworkV1Engine] class.
var (
	_MLNeuralNetworkV1EngineClass     MLNeuralNetworkV1EngineClass
	_MLNeuralNetworkV1EngineClassOnce sync.Once
)

func getMLNeuralNetworkV1EngineClass() MLNeuralNetworkV1EngineClass {
	_MLNeuralNetworkV1EngineClassOnce.Do(func() {
		_MLNeuralNetworkV1EngineClass = MLNeuralNetworkV1EngineClass{class: objc.GetClass("MLNeuralNetworkV1Engine")}
	})
	return _MLNeuralNetworkV1EngineClass
}

// GetMLNeuralNetworkV1EngineClass returns the class object for MLNeuralNetworkV1Engine.
func GetMLNeuralNetworkV1EngineClass() MLNeuralNetworkV1EngineClass {
	return getMLNeuralNetworkV1EngineClass()
}

type MLNeuralNetworkV1EngineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNeuralNetworkV1EngineClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNeuralNetworkV1EngineClass) Alloc() MLNeuralNetworkV1Engine {
	rv := objc.Send[MLNeuralNetworkV1Engine](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkV1Engine
type MLNeuralNetworkV1Engine struct {
	MLNeuralNetworkEngine
}

// MLNeuralNetworkV1EngineFromID constructs a [MLNeuralNetworkV1Engine] from an objc.ID.
func MLNeuralNetworkV1EngineFromID(id objc.ID) MLNeuralNetworkV1Engine {
	return MLNeuralNetworkV1Engine{MLNeuralNetworkEngine: MLNeuralNetworkEngineFromID(id)}
}
// Ensure MLNeuralNetworkV1Engine implements IMLNeuralNetworkV1Engine.
var _ IMLNeuralNetworkV1Engine = MLNeuralNetworkV1Engine{}

// An interface definition for the [MLNeuralNetworkV1Engine] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkV1Engine
type IMLNeuralNetworkV1Engine interface {
	IMLNeuralNetworkEngine
}

// Init initializes the instance.
func (n MLNeuralNetworkV1Engine) Init() MLNeuralNetworkV1Engine {
	rv := objc.Send[MLNeuralNetworkV1Engine](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MLNeuralNetworkV1Engine) Autorelease() MLNeuralNetworkV1Engine {
	rv := objc.Send[MLNeuralNetworkV1Engine](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNeuralNetworkV1Engine creates a new MLNeuralNetworkV1Engine instance.
func NewMLNeuralNetworkV1Engine() MLNeuralNetworkV1Engine {
	class := getMLNeuralNetworkV1EngineClass()
	rv := objc.Send[MLNeuralNetworkV1Engine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkEngine/initWithContainer:configuration:error:
func NewNeuralNetworkV1EngineWithContainerConfigurationError(container objectivec.IObject, configuration objectivec.IObject) (MLNeuralNetworkV1Engine, error) {
	var errorPtr objc.ID
	instance := getMLNeuralNetworkV1EngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainer:configuration:error:"), container, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkV1Engine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkV1EngineFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkEngine/initWithContainer:error:
func NewNeuralNetworkV1EngineWithContainerError(container objectivec.IObject) (MLNeuralNetworkV1Engine, error) {
	var errorPtr objc.ID
	instance := getMLNeuralNetworkV1EngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainer:error:"), container, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkV1Engine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkV1EngineFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithDescription:configuration:
func NewNeuralNetworkV1EngineWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLNeuralNetworkV1Engine {
	instance := getMLNeuralNetworkV1EngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLNeuralNetworkV1EngineFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewNeuralNetworkV1EngineWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLNeuralNetworkV1Engine {
	instance := getMLNeuralNetworkV1EngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLNeuralNetworkV1EngineFromID(rv)
}

