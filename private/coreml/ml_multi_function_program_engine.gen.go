// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLMultiFunctionProgramEngine] class.
var (
	_MLMultiFunctionProgramEngineClass     MLMultiFunctionProgramEngineClass
	_MLMultiFunctionProgramEngineClassOnce sync.Once
)

func getMLMultiFunctionProgramEngineClass() MLMultiFunctionProgramEngineClass {
	_MLMultiFunctionProgramEngineClassOnce.Do(func() {
		_MLMultiFunctionProgramEngineClass = MLMultiFunctionProgramEngineClass{class: objc.GetClass("MLMultiFunctionProgramEngine")}
	})
	return _MLMultiFunctionProgramEngineClass
}

// GetMLMultiFunctionProgramEngineClass returns the class object for MLMultiFunctionProgramEngine.
func GetMLMultiFunctionProgramEngineClass() MLMultiFunctionProgramEngineClass {
	return getMLMultiFunctionProgramEngineClass()
}

type MLMultiFunctionProgramEngineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLMultiFunctionProgramEngineClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLMultiFunctionProgramEngineClass) Alloc() MLMultiFunctionProgramEngine {
	rv := objc.Send[MLMultiFunctionProgramEngine](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLMultiFunctionProgramEngine.ClassLabels]
//   - [MLMultiFunctionProgramEngine.ClassifyOptionsError]
//   - [MLMultiFunctionProgramEngine.Container]
//   - [MLMultiFunctionProgramEngine.EvaluateError]
//   - [MLMultiFunctionProgramEngine.EvaluateFunctionArgumentsError]
//   - [MLMultiFunctionProgramEngine.ModelFileBasePath]
//   - [MLMultiFunctionProgramEngine.NewContextAndReturnError]
//   - [MLMultiFunctionProgramEngine.Program]
//   - [MLMultiFunctionProgramEngine.ProgramEngineForFunctionError]
//   - [MLMultiFunctionProgramEngine.RegressOptionsError]
//   - [MLMultiFunctionProgramEngine.RemoveEngineForFunctionName]
//   - [MLMultiFunctionProgramEngine.SerializedMILText]
//   - [MLMultiFunctionProgramEngine.UpdateModelFilePath]
//   - [MLMultiFunctionProgramEngine.VerifyArgumentNamesFunctionNameError]
//   - [MLMultiFunctionProgramEngine.InitWithProgramContainerConfigurationError]
// See: https://developer.apple.com/documentation/CoreML/MLMultiFunctionProgramEngine
type MLMultiFunctionProgramEngine struct {
	MLModelEngine
}

// MLMultiFunctionProgramEngineFromID constructs a [MLMultiFunctionProgramEngine] from an objc.ID.
func MLMultiFunctionProgramEngineFromID(id objc.ID) MLMultiFunctionProgramEngine {
	return MLMultiFunctionProgramEngine{MLModelEngine: MLModelEngineFromID(id)}
}
// Ensure MLMultiFunctionProgramEngine implements IMLMultiFunctionProgramEngine.
var _ IMLMultiFunctionProgramEngine = MLMultiFunctionProgramEngine{}

// An interface definition for the [MLMultiFunctionProgramEngine] class.
//
// # Methods
//
//   - [IMLMultiFunctionProgramEngine.ClassLabels]
//   - [IMLMultiFunctionProgramEngine.ClassifyOptionsError]
//   - [IMLMultiFunctionProgramEngine.Container]
//   - [IMLMultiFunctionProgramEngine.EvaluateError]
//   - [IMLMultiFunctionProgramEngine.EvaluateFunctionArgumentsError]
//   - [IMLMultiFunctionProgramEngine.ModelFileBasePath]
//   - [IMLMultiFunctionProgramEngine.NewContextAndReturnError]
//   - [IMLMultiFunctionProgramEngine.Program]
//   - [IMLMultiFunctionProgramEngine.ProgramEngineForFunctionError]
//   - [IMLMultiFunctionProgramEngine.RegressOptionsError]
//   - [IMLMultiFunctionProgramEngine.RemoveEngineForFunctionName]
//   - [IMLMultiFunctionProgramEngine.SerializedMILText]
//   - [IMLMultiFunctionProgramEngine.UpdateModelFilePath]
//   - [IMLMultiFunctionProgramEngine.VerifyArgumentNamesFunctionNameError]
//   - [IMLMultiFunctionProgramEngine.InitWithProgramContainerConfigurationError]
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiFunctionProgramEngine
type IMLMultiFunctionProgramEngine interface {
	IMLModelEngine

	// Topic: Methods

	ClassLabels() objectivec.IObject
	ClassifyOptionsError(classify objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	Container() IMLMultiFunctionProgramContainer
	EvaluateError(evaluate objectivec.IObject) (objectivec.IObject, error)
	EvaluateFunctionArgumentsError(function objectivec.IObject, arguments objectivec.IObject) (objectivec.IObject, error)
	ModelFileBasePath() string
	NewContextAndReturnError() (objectivec.IObject, error)
	Program() objectivec.IObject
	ProgramEngineForFunctionError(function objectivec.IObject) (objectivec.IObject, error)
	RegressOptionsError(regress objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	RemoveEngineForFunctionName(name objectivec.IObject)
	SerializedMILText() string
	UpdateModelFilePath(path objectivec.IObject)
	VerifyArgumentNamesFunctionNameError(names objectivec.IObject, name objectivec.IObject) (bool, error)
	InitWithProgramContainerConfigurationError(container objectivec.IObject, configuration objectivec.IObject) (MLMultiFunctionProgramEngine, error)
}

// Init initializes the instance.
func (m MLMultiFunctionProgramEngine) Init() MLMultiFunctionProgramEngine {
	rv := objc.Send[MLMultiFunctionProgramEngine](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLMultiFunctionProgramEngine) Autorelease() MLMultiFunctionProgramEngine {
	rv := objc.Send[MLMultiFunctionProgramEngine](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLMultiFunctionProgramEngine creates a new MLMultiFunctionProgramEngine instance.
func NewMLMultiFunctionProgramEngine() MLMultiFunctionProgramEngine {
	class := getMLMultiFunctionProgramEngineClass()
	rv := objc.Send[MLMultiFunctionProgramEngine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithDescription:configuration:
func NewMultiFunctionProgramEngineWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLMultiFunctionProgramEngine {
	instance := getMLMultiFunctionProgramEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLMultiFunctionProgramEngineFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewMultiFunctionProgramEngineWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLMultiFunctionProgramEngine {
	instance := getMLMultiFunctionProgramEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLMultiFunctionProgramEngineFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLMultiFunctionProgramEngine/initWithProgramContainer:configuration:error:
func NewMultiFunctionProgramEngineWithProgramContainerConfigurationError(container objectivec.IObject, configuration objectivec.IObject) (MLMultiFunctionProgramEngine, error) {
	var errorPtr objc.ID
	instance := getMLMultiFunctionProgramEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProgramContainer:configuration:error:"), container, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLMultiFunctionProgramEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLMultiFunctionProgramEngineFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiFunctionProgramEngine/classLabels
func (m MLMultiFunctionProgramEngine) ClassLabels() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classLabels"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiFunctionProgramEngine/classify:options:error:
func (m MLMultiFunctionProgramEngine) ClassifyOptionsError(classify objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classify:options:error:"), classify, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiFunctionProgramEngine/evaluate:error:
func (m MLMultiFunctionProgramEngine) EvaluateError(evaluate objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("evaluate:error:"), evaluate, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiFunctionProgramEngine/evaluateFunction:arguments:error:
func (m MLMultiFunctionProgramEngine) EvaluateFunctionArgumentsError(function objectivec.IObject, arguments objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("evaluateFunction:arguments:error:"), function, arguments, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiFunctionProgramEngine/newContextAndReturnError:
func (m MLMultiFunctionProgramEngine) NewContextAndReturnError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("newContextAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// See: https://developer.apple.com/documentation/CoreML/MLMultiFunctionProgramEngine/program
func (m MLMultiFunctionProgramEngine) Program() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("program"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiFunctionProgramEngine/programEngineForFunction:error:
func (m MLMultiFunctionProgramEngine) ProgramEngineForFunctionError(function objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("programEngineForFunction:error:"), function, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiFunctionProgramEngine/regress:options:error:
func (m MLMultiFunctionProgramEngine) RegressOptionsError(regress objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("regress:options:error:"), regress, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiFunctionProgramEngine/removeEngineForFunctionName:
func (m MLMultiFunctionProgramEngine) RemoveEngineForFunctionName(name objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeEngineForFunctionName:"), name)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiFunctionProgramEngine/updateModelFilePath:
func (m MLMultiFunctionProgramEngine) UpdateModelFilePath(path objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("updateModelFilePath:"), path)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiFunctionProgramEngine/verifyArgumentNames:functionName:error:
func (m MLMultiFunctionProgramEngine) VerifyArgumentNamesFunctionNameError(names objectivec.IObject, name objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("verifyArgumentNames:functionName:error:"), names, name, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("verifyArgumentNames:functionName:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiFunctionProgramEngine/initWithProgramContainer:configuration:error:
func (m MLMultiFunctionProgramEngine) InitWithProgramContainerConfigurationError(container objectivec.IObject, configuration objectivec.IObject) (MLMultiFunctionProgramEngine, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithProgramContainer:configuration:error:"), container, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLMultiFunctionProgramEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLMultiFunctionProgramEngineFromID(rv), nil

}

//
// See: https://developer.apple.com/documentation/CoreML/MLMultiFunctionProgramEngine/loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:
func (_MLMultiFunctionProgramEngineClass MLMultiFunctionProgramEngineClass) LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLMultiFunctionProgramEngineClass.class), objc.Sel("loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLMultiFunctionProgramEngine/container
func (m MLMultiFunctionProgramEngine) Container() IMLMultiFunctionProgramContainer {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("container"))
	return MLMultiFunctionProgramContainerFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLMultiFunctionProgramEngine/modelFileBasePath
func (m MLMultiFunctionProgramEngine) ModelFileBasePath() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelFileBasePath"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLMultiFunctionProgramEngine/serializedMILText
func (m MLMultiFunctionProgramEngine) SerializedMILText() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("serializedMILText"))
	return foundation.NSStringFromID(rv).String()
}

