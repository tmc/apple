// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelSwiftEngine] class.
var (
	_MLModelSwiftEngineClass     MLModelSwiftEngineClass
	_MLModelSwiftEngineClassOnce sync.Once
)

func getMLModelSwiftEngineClass() MLModelSwiftEngineClass {
	_MLModelSwiftEngineClassOnce.Do(func() {
		_MLModelSwiftEngineClass = MLModelSwiftEngineClass{class: objc.GetClass("MLModelSwiftEngine")}
	})
	return _MLModelSwiftEngineClass
}

// GetMLModelSwiftEngineClass returns the class object for MLModelSwiftEngine.
func GetMLModelSwiftEngineClass() MLModelSwiftEngineClass {
	return getMLModelSwiftEngineClass()
}

type MLModelSwiftEngineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelSwiftEngineClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelSwiftEngineClass) Alloc() MLModelSwiftEngine {
	rv := objc.Send[MLModelSwiftEngine](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelSwiftEngine
type MLModelSwiftEngine struct {
	MLModelEngine
}

// MLModelSwiftEngineFromID constructs a [MLModelSwiftEngine] from an objc.ID.
func MLModelSwiftEngineFromID(id objc.ID) MLModelSwiftEngine {
	return MLModelSwiftEngine{MLModelEngine: MLModelEngineFromID(id)}
}

// Ensure MLModelSwiftEngine implements IMLModelSwiftEngine.
var _ IMLModelSwiftEngine = MLModelSwiftEngine{}

// An interface definition for the [MLModelSwiftEngine] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelSwiftEngine
type IMLModelSwiftEngine interface {
	IMLModelEngine
}

// Init initializes the instance.
func (m MLModelSwiftEngine) Init() MLModelSwiftEngine {
	rv := objc.Send[MLModelSwiftEngine](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelSwiftEngine) Autorelease() MLModelSwiftEngine {
	rv := objc.Send[MLModelSwiftEngine](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelSwiftEngine creates a new MLModelSwiftEngine instance.
func NewMLModelSwiftEngine() MLModelSwiftEngine {
	class := getMLModelSwiftEngineClass()
	rv := objc.Send[MLModelSwiftEngine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelSwiftEngine/initWithDescription:configuration:
func NewModelSwiftEngineWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLModelSwiftEngine {
	instance := getMLModelSwiftEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLModelSwiftEngineFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewModelSwiftEngineWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLModelSwiftEngine {
	instance := getMLModelSwiftEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLModelSwiftEngineFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelSwiftEngine/loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:
func (_MLModelSwiftEngineClass MLModelSwiftEngineClass) LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelSwiftEngineClass.class), objc.Sel("loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
