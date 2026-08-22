// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[MLE5EnumeratedShapeExecutionStreamOperationPool](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

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
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) Init() MLE5EnumeratedShapeExecutionStreamOperationPool {
	rv := objc.SendIfResponds[MLE5EnumeratedShapeExecutionStreamOperationPool](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) Autorelease() MLE5EnumeratedShapeExecutionStreamOperationPool {
	rv := objc.SendIfResponds[MLE5EnumeratedShapeExecutionStreamOperationPool](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLE5EnumeratedShapeExecutionStreamOperationPool creates a new MLE5EnumeratedShapeExecutionStreamOperationPool instance.
func NewMLE5EnumeratedShapeExecutionStreamOperationPool() MLE5EnumeratedShapeExecutionStreamOperationPool {
	class := getMLE5EnumeratedShapeExecutionStreamOperationPoolClass()
	rv := objc.SendIfResponds[MLE5EnumeratedShapeExecutionStreamOperationPool](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewE5EnumeratedShapeExecutionStreamOperationPoolWithProgramLibraryFunctionNameModelDescriptionConfigurationModelSignpostIdCompilerVersionInfo(library objectivec.IObject, name objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject, id uint64, info objectivec.IObject) MLE5EnumeratedShapeExecutionStreamOperationPool {
	instance := getMLE5EnumeratedShapeExecutionStreamOperationPoolClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithProgramLibrary:functionName:modelDescription:configuration:modelSignpostId:compilerVersionInfo:"), library, name, description, configuration, id, info)
	return MLE5EnumeratedShapeExecutionStreamOperationPoolFromID(rv)
}

func (m MLE5EnumeratedShapeExecutionStreamOperationPool) _putBack(back objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("_putBack:"), back)
}
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) _takeOutOperationForFunctionNameError(name objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_takeOutOperationForFunctionName:error:"), name, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// TakeOutOperationForFunctionNameError is an exported wrapper for the private method _takeOutOperationForFunctionNameError.
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) TakeOutOperationForFunctionNameError(name objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_takeOutOperationForFunctionName:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_takeOutOperationForFunctionName:error:"}
		return nil, err
	}
	return m._takeOutOperationForFunctionNameError(name)
}

// CanTakeOutOperationForFunctionNameError reports whether the receiver responds to the private selector _takeOutOperationForFunctionName:error:.
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) CanTakeOutOperationForFunctionNameError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_takeOutOperationForFunctionName:error:"))
}
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) _takeOutOperationFromAnyProgramFunction() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("_takeOutOperationFromAnyProgramFunction"))
	return objectivec.Object{ID: rv}
}

// TakeOutOperationFromAnyProgramFunction is an exported wrapper for the private method _takeOutOperationFromAnyProgramFunction.
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) TakeOutOperationFromAnyProgramFunction() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_takeOutOperationFromAnyProgramFunction")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_takeOutOperationFromAnyProgramFunction"}
		return nil, err
	}
	return m._takeOutOperationFromAnyProgramFunction(), nil
}

// CanTakeOutOperationFromAnyProgramFunction reports whether the receiver responds to the private selector _takeOutOperationFromAnyProgramFunction.
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) CanTakeOutOperationFromAnyProgramFunction() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_takeOutOperationFromAnyProgramFunction"))
}
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) PrepareWithInitialPoolSizeError(size int64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("prepareWithInitialPoolSize:error:"), size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("prepareWithInitialPoolSize:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) PutBack(back objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("putBack:"), back)
}
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) TakeOutOperationForFeaturesError(features objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("takeOutOperationForFeatures:error:"), features, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) InitWithProgramLibraryFunctionNameModelDescriptionConfigurationModelSignpostIdCompilerVersionInfo(library objectivec.IObject, name objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject, id uint64, info objectivec.IObject) MLE5EnumeratedShapeExecutionStreamOperationPool {
	rv := objc.SendIfResponds[MLE5EnumeratedShapeExecutionStreamOperationPool](m.ID, objc.Sel("initWithProgramLibrary:functionName:modelDescription:configuration:modelSignpostId:compilerVersionInfo:"), library, name, description, configuration, id, info)
	return rv
}

func (m MLE5EnumeratedShapeExecutionStreamOperationPool) CompilerVersionInfo() IMLVersionInfo {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("compilerVersionInfo"))
	return MLVersionInfoFromID(objc.ID(rv))
}
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) Configuration() IMLModelConfiguration {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("configuration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) FunctionNameToPoolMap() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("functionNameToPoolMap"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) MilEntryFunctionName() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("milEntryFunctionName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) MilFunctionNames() foundation.INSSet {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("milFunctionNames"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) ModelDescription() IMLModelDescription {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) ModelSignpostId() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("modelSignpostId"))
	return rv
}
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) ProgramLibrary() IMLE5ProgramLibrary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("programLibrary"))
	return MLE5ProgramLibraryFromID(objc.ID(rv))
}
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) SerialQueue() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("serialQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (m MLE5EnumeratedShapeExecutionStreamOperationPool) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
