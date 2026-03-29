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

// The class instance for the [MLE5EnumeratedShapeExecutionStreamOperationPool] class.
var (
	_MLE5EnumeratedShapeExecutionStreamOperationPoolClass     MLE5EnumeratedShapeExecutionStreamOperationPoolClass
	_MLE5EnumeratedShapeExecutionStreamOperationPoolClassOnce sync.Once
)

func getMLE5EnumeratedShapeExecutionStreamOperationPoolClass() MLE5EnumeratedShapeExecutionStreamOperationPoolClass {
	_MLE5EnumeratedShapeExecutionStreamOperationPoolClassOnce.Do(func() {
		_MLE5EnumeratedShapeExecutionStreamOperationPoolClass = MLE5EnumeratedShapeExecutionStreamOperationPoolClass{class: objc.GetClass("MLE5EnumeratedShapeExecutionStreamOperationPool")}
	})
	return _MLE5EnumeratedShapeExecutionStreamOperationPoolClass
}

// GetMLE5EnumeratedShapeExecutionStreamOperationPoolClass returns the class object for MLE5EnumeratedShapeExecutionStreamOperationPool.
func GetMLE5EnumeratedShapeExecutionStreamOperationPoolClass() MLE5EnumeratedShapeExecutionStreamOperationPoolClass {
	return getMLE5EnumeratedShapeExecutionStreamOperationPoolClass()
}

type MLE5EnumeratedShapeExecutionStreamOperationPoolClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLE5EnumeratedShapeExecutionStreamOperationPoolClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLE5EnumeratedShapeExecutionStreamOperationPoolClass) Alloc() MLE5EnumeratedShapeExecutionStreamOperationPool {
	rv := objc.Send[MLE5EnumeratedShapeExecutionStreamOperationPool](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLE5EnumeratedShapeExecutionStreamOperationPool._putBack]
//   - [MLE5EnumeratedShapeExecutionStreamOperationPool._takeOutOperationForFunctionNameError]
//   - [MLE5EnumeratedShapeExecutionStreamOperationPool._takeOutOperationFromAnyProgramFunction]
//   - [MLE5EnumeratedShapeExecutionStreamOperationPool.CompilerVersionInfo]
//   - [MLE5EnumeratedShapeExecutionStreamOperationPool.Configuration]
//   - [MLE5EnumeratedShapeExecutionStreamOperationPool.FunctionNameToPoolMap]
//   - [MLE5EnumeratedShapeExecutionStreamOperationPool.MilEntryFunctionName]
//   - [MLE5EnumeratedShapeExecutionStreamOperationPool.MilFunctionNames]
//   - [MLE5EnumeratedShapeExecutionStreamOperationPool.ModelDescription]
//   - [MLE5EnumeratedShapeExecutionStreamOperationPool.ModelSignpostId]
//   - [MLE5EnumeratedShapeExecutionStreamOperationPool.PrepareWithInitialPoolSizeError]
//   - [MLE5EnumeratedShapeExecutionStreamOperationPool.ProgramLibrary]
//   - [MLE5EnumeratedShapeExecutionStreamOperationPool.PutBack]
//   - [MLE5EnumeratedShapeExecutionStreamOperationPool.SerialQueue]
//   - [MLE5EnumeratedShapeExecutionStreamOperationPool.TakeOutOperationForFeaturesError]
//   - [MLE5EnumeratedShapeExecutionStreamOperationPool.InitWithProgramLibraryFunctionNameModelDescriptionConfigurationModelSignpostIdCompilerVersionInfo]
//   - [MLE5EnumeratedShapeExecutionStreamOperationPool.DebugDescription]
//   - [MLE5EnumeratedShapeExecutionStreamOperationPool.Description]
//   - [MLE5EnumeratedShapeExecutionStreamOperationPool.Hash]
//   - [MLE5EnumeratedShapeExecutionStreamOperationPool.Superclass]
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool
type MLE5EnumeratedShapeExecutionStreamOperationPool struct {
	objectivec.Object
}

// MLE5EnumeratedShapeExecutionStreamOperationPoolFromID constructs a [MLE5EnumeratedShapeExecutionStreamOperationPool] from an objc.ID.
func MLE5EnumeratedShapeExecutionStreamOperationPoolFromID(id objc.ID) MLE5EnumeratedShapeExecutionStreamOperationPool {
	return MLE5EnumeratedShapeExecutionStreamOperationPool{objectivec.Object{ID: id}}
}
// Ensure MLE5EnumeratedShapeExecutionStreamOperationPool implements IMLE5EnumeratedShapeExecutionStreamOperationPool.
var _ IMLE5EnumeratedShapeExecutionStreamOperationPool = MLE5EnumeratedShapeExecutionStreamOperationPool{}

// An interface definition for the [MLE5EnumeratedShapeExecutionStreamOperationPool] class.
//
// # Methods
//
//   - [IMLE5EnumeratedShapeExecutionStreamOperationPool._putBack]
//   - [IMLE5EnumeratedShapeExecutionStreamOperationPool._takeOutOperationForFunctionNameError]
//   - [IMLE5EnumeratedShapeExecutionStreamOperationPool._takeOutOperationFromAnyProgramFunction]
//   - [IMLE5EnumeratedShapeExecutionStreamOperationPool.CompilerVersionInfo]
//   - [IMLE5EnumeratedShapeExecutionStreamOperationPool.Configuration]
//   - [IMLE5EnumeratedShapeExecutionStreamOperationPool.FunctionNameToPoolMap]
//   - [IMLE5EnumeratedShapeExecutionStreamOperationPool.MilEntryFunctionName]
//   - [IMLE5EnumeratedShapeExecutionStreamOperationPool.MilFunctionNames]
//   - [IMLE5EnumeratedShapeExecutionStreamOperationPool.ModelDescription]
//   - [IMLE5EnumeratedShapeExecutionStreamOperationPool.ModelSignpostId]
//   - [IMLE5EnumeratedShapeExecutionStreamOperationPool.PrepareWithInitialPoolSizeError]
//   - [IMLE5EnumeratedShapeExecutionStreamOperationPool.ProgramLibrary]
//   - [IMLE5EnumeratedShapeExecutionStreamOperationPool.PutBack]
//   - [IMLE5EnumeratedShapeExecutionStreamOperationPool.SerialQueue]
//   - [IMLE5EnumeratedShapeExecutionStreamOperationPool.TakeOutOperationForFeaturesError]
//   - [IMLE5EnumeratedShapeExecutionStreamOperationPool.InitWithProgramLibraryFunctionNameModelDescriptionConfigurationModelSignpostIdCompilerVersionInfo]
//   - [IMLE5EnumeratedShapeExecutionStreamOperationPool.DebugDescription]
//   - [IMLE5EnumeratedShapeExecutionStreamOperationPool.Description]
//   - [IMLE5EnumeratedShapeExecutionStreamOperationPool.Hash]
//   - [IMLE5EnumeratedShapeExecutionStreamOperationPool.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool
type IMLE5EnumeratedShapeExecutionStreamOperationPool interface {
	objectivec.IObject

	// Topic: Methods

	_putBack(back objectivec.IObject)
	_takeOutOperationForFunctionNameError(name objectivec.IObject) (objectivec.IObject, error)
	_takeOutOperationFromAnyProgramFunction() objectivec.IObject
	CompilerVersionInfo() IMLVersionInfo
	Configuration() IMLModelConfiguration
	FunctionNameToPoolMap() foundation.INSDictionary
	MilEntryFunctionName() string
	MilFunctionNames() foundation.INSSet
	ModelDescription() IMLModelDescription
	ModelSignpostId() uint64
	PrepareWithInitialPoolSizeError(size int64) (bool, error)
	ProgramLibrary() IMLE5ProgramLibrary
	PutBack(back objectivec.IObject)
	SerialQueue() objectivec.Object
	TakeOutOperationForFeaturesError(features objectivec.IObject) (objectivec.IObject, error)
	InitWithProgramLibraryFunctionNameModelDescriptionConfigurationModelSignpostIdCompilerVersionInfo(library objectivec.IObject, name objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject, id uint64, info objectivec.IObject) MLE5EnumeratedShapeExecutionStreamOperationPool
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) Init() MLE5EnumeratedShapeExecutionStreamOperationPool {
	rv := objc.Send[MLE5EnumeratedShapeExecutionStreamOperationPool](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) Autorelease() MLE5EnumeratedShapeExecutionStreamOperationPool {
	rv := objc.Send[MLE5EnumeratedShapeExecutionStreamOperationPool](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLE5EnumeratedShapeExecutionStreamOperationPool creates a new MLE5EnumeratedShapeExecutionStreamOperationPool instance.
func NewMLE5EnumeratedShapeExecutionStreamOperationPool() MLE5EnumeratedShapeExecutionStreamOperationPool {
	class := getMLE5EnumeratedShapeExecutionStreamOperationPoolClass()
	rv := objc.Send[MLE5EnumeratedShapeExecutionStreamOperationPool](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool/initWithProgramLibrary:functionName:modelDescription:configuration:modelSignpostId:compilerVersionInfo:
func NewE5EnumeratedShapeExecutionStreamOperationPoolWithProgramLibraryFunctionNameModelDescriptionConfigurationModelSignpostIdCompilerVersionInfo(library objectivec.IObject, name objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject, id uint64, info objectivec.IObject) MLE5EnumeratedShapeExecutionStreamOperationPool {
	instance := getMLE5EnumeratedShapeExecutionStreamOperationPoolClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProgramLibrary:functionName:modelDescription:configuration:modelSignpostId:compilerVersionInfo:"), library, name, description, configuration, id, info)
	return MLE5EnumeratedShapeExecutionStreamOperationPoolFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool/_putBack:
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) _putBack(back objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("_putBack:"), back)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool/_takeOutOperationForFunctionName:error:
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) _takeOutOperationForFunctionNameError(name objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("_takeOutOperationForFunctionName:error:"), name, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// TakeOutOperationForFunctionNameError is an exported wrapper for the private method _takeOutOperationForFunctionNameError.
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) TakeOutOperationForFunctionNameError(name objectivec.IObject) (objectivec.IObject, error) {
	return e._takeOutOperationForFunctionNameError(name)
}
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool/_takeOutOperationFromAnyProgramFunction
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) _takeOutOperationFromAnyProgramFunction() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("_takeOutOperationFromAnyProgramFunction"))
	return objectivec.Object{ID: rv}
}

// TakeOutOperationFromAnyProgramFunction is an exported wrapper for the private method _takeOutOperationFromAnyProgramFunction.
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) TakeOutOperationFromAnyProgramFunction() objectivec.IObject {
	return e._takeOutOperationFromAnyProgramFunction()
}
//
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool/prepareWithInitialPoolSize:error:
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) PrepareWithInitialPoolSizeError(size int64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("prepareWithInitialPoolSize:error:"), size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("prepareWithInitialPoolSize:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool/putBack:
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) PutBack(back objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("putBack:"), back)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool/takeOutOperationForFeatures:error:
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) TakeOutOperationForFeaturesError(features objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("takeOutOperationForFeatures:error:"), features, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool/initWithProgramLibrary:functionName:modelDescription:configuration:modelSignpostId:compilerVersionInfo:
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) InitWithProgramLibraryFunctionNameModelDescriptionConfigurationModelSignpostIdCompilerVersionInfo(library objectivec.IObject, name objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject, id uint64, info objectivec.IObject) MLE5EnumeratedShapeExecutionStreamOperationPool {
	rv := objc.Send[MLE5EnumeratedShapeExecutionStreamOperationPool](e.ID, objc.Sel("initWithProgramLibrary:functionName:modelDescription:configuration:modelSignpostId:compilerVersionInfo:"), library, name, description, configuration, id, info)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool/compilerVersionInfo
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) CompilerVersionInfo() IMLVersionInfo {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("compilerVersionInfo"))
	return MLVersionInfoFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool/configuration
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) Configuration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("configuration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool/debugDescription
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) DebugDescription() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool/description
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) Description() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool/functionNameToPoolMap
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) FunctionNameToPoolMap() foundation.INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("functionNameToPoolMap"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool/hash
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) Hash() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool/milEntryFunctionName
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) MilEntryFunctionName() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("milEntryFunctionName"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool/milFunctionNames
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) MilFunctionNames() foundation.INSSet {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("milFunctionNames"))
	return foundation.NSSetFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool/modelDescription
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool/modelSignpostId
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) ModelSignpostId() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("modelSignpostId"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool/programLibrary
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) ProgramLibrary() IMLE5ProgramLibrary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("programLibrary"))
	return MLE5ProgramLibraryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool/serialQueue
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) SerialQueue() objectivec.Object {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("serialQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLE5EnumeratedShapeExecutionStreamOperationPool/superclass
func (e MLE5EnumeratedShapeExecutionStreamOperationPool) Superclass() objc.Class {
	rv := objc.Send[objc.Class](e.ID, objc.Sel("superclass"))
	return rv
}

