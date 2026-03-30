// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLProgramEngine] class.
var (
	_MLProgramEngineClass     MLProgramEngineClass
	_MLProgramEngineClassOnce sync.Once
)

func getMLProgramEngineClass() MLProgramEngineClass {
	_MLProgramEngineClassOnce.Do(func() {
		_MLProgramEngineClass = MLProgramEngineClass{class: objc.GetClass("MLProgramEngine")}
	})
	return _MLProgramEngineClass
}

// GetMLProgramEngineClass returns the class object for MLProgramEngine.
func GetMLProgramEngineClass() MLProgramEngineClass {
	return getMLProgramEngineClass()
}

type MLProgramEngineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLProgramEngineClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLProgramEngineClass) Alloc() MLProgramEngine {
	rv := objc.Send[MLProgramEngine](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramEngine
type MLProgramEngine struct {
	MLNeuralNetworkEngine
}

// MLProgramEngineFromID constructs a [MLProgramEngine] from an objc.ID.
func MLProgramEngineFromID(id objc.ID) MLProgramEngine {
	return MLProgramEngine{MLNeuralNetworkEngine: MLNeuralNetworkEngineFromID(id)}
}

// Ensure MLProgramEngine implements IMLProgramEngine.
var _ IMLProgramEngine = MLProgramEngine{}

// An interface definition for the [MLProgramEngine] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLProgramEngine
type IMLProgramEngine interface {
	IMLNeuralNetworkEngine
}

// Init initializes the instance.
func (p MLProgramEngine) Init() MLProgramEngine {
	rv := objc.Send[MLProgramEngine](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLProgramEngine) Autorelease() MLProgramEngine {
	rv := objc.Send[MLProgramEngine](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLProgramEngine creates a new MLProgramEngine instance.
func NewMLProgramEngine() MLProgramEngine {
	class := getMLProgramEngineClass()
	rv := objc.Send[MLProgramEngine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkEngine/initWithContainer:configuration:error:
func NewProgramEngineWithContainerConfigurationError(container objectivec.IObject, configuration objectivec.IObject) (MLProgramEngine, error) {
	var errorPtr objc.ID
	instance := getMLProgramEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainer:configuration:error:"), container, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLProgramEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLProgramEngineFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkEngine/initWithContainer:error:
func NewProgramEngineWithContainerError(container objectivec.IObject) (MLProgramEngine, error) {
	var errorPtr objc.ID
	instance := getMLProgramEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainer:error:"), container, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLProgramEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLProgramEngineFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithDescription:configuration:
func NewProgramEngineWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLProgramEngine {
	instance := getMLProgramEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLProgramEngineFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewProgramEngineWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLProgramEngine {
	instance := getMLProgramEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLProgramEngineFromID(rv)
}
