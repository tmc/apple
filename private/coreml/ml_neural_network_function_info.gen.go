// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNeuralNetworkFunctionInfo] class.
var (
	_MLNeuralNetworkFunctionInfoClass     MLNeuralNetworkFunctionInfoClass
	_MLNeuralNetworkFunctionInfoClassOnce sync.Once
)

func getMLNeuralNetworkFunctionInfoClass() MLNeuralNetworkFunctionInfoClass {
	_MLNeuralNetworkFunctionInfoClassOnce.Do(func() {
		_MLNeuralNetworkFunctionInfoClass = MLNeuralNetworkFunctionInfoClass{class: objc.GetClass("MLNeuralNetworkFunctionInfo")}
	})
	return _MLNeuralNetworkFunctionInfoClass
}

// GetMLNeuralNetworkFunctionInfoClass returns the class object for MLNeuralNetworkFunctionInfo.
func GetMLNeuralNetworkFunctionInfoClass() MLNeuralNetworkFunctionInfoClass {
	return getMLNeuralNetworkFunctionInfoClass()
}

type MLNeuralNetworkFunctionInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNeuralNetworkFunctionInfoClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNeuralNetworkFunctionInfoClass) Alloc() MLNeuralNetworkFunctionInfo {
	rv := objc.Send[MLNeuralNetworkFunctionInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLNeuralNetworkFunctionInfo.ClassLabels]
//   - [MLNeuralNetworkFunctionInfo.ClassScoreVectorName]
//   - [MLNeuralNetworkFunctionInfo.IsClassifier]
//   - [MLNeuralNetworkFunctionInfo.OutputNames]
//   - [MLNeuralNetworkFunctionInfo.InitWithCompiledModelArchiveCompilerVersionInfoError]
type MLNeuralNetworkFunctionInfo struct {
	objectivec.Object
}

// MLNeuralNetworkFunctionInfoFromID constructs a [MLNeuralNetworkFunctionInfo] from an objc.ID.
func MLNeuralNetworkFunctionInfoFromID(id objc.ID) MLNeuralNetworkFunctionInfo {
	return MLNeuralNetworkFunctionInfo{objectivec.Object{ID: id}}
}

// Ensure MLNeuralNetworkFunctionInfo implements IMLNeuralNetworkFunctionInfo.
var _ IMLNeuralNetworkFunctionInfo = MLNeuralNetworkFunctionInfo{}

// An interface definition for the [MLNeuralNetworkFunctionInfo] class.
//
// # Methods
//
//   - [IMLNeuralNetworkFunctionInfo.ClassLabels]
//   - [IMLNeuralNetworkFunctionInfo.ClassScoreVectorName]
//   - [IMLNeuralNetworkFunctionInfo.IsClassifier]
//   - [IMLNeuralNetworkFunctionInfo.OutputNames]
//   - [IMLNeuralNetworkFunctionInfo.InitWithCompiledModelArchiveCompilerVersionInfoError]
type IMLNeuralNetworkFunctionInfo interface {
	objectivec.IObject

	// Topic: Methods

	ClassLabels() foundation.INSArray
	ClassScoreVectorName() string
	IsClassifier() bool
	OutputNames() foundation.INSArray
	InitWithCompiledModelArchiveCompilerVersionInfoError(archive unsafe.Pointer, info objectivec.IObject) (MLNeuralNetworkFunctionInfo, error)
}

// Init initializes the instance.
func (m MLNeuralNetworkFunctionInfo) Init() MLNeuralNetworkFunctionInfo {
	rv := objc.Send[MLNeuralNetworkFunctionInfo](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLNeuralNetworkFunctionInfo) Autorelease() MLNeuralNetworkFunctionInfo {
	rv := objc.Send[MLNeuralNetworkFunctionInfo](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNeuralNetworkFunctionInfo creates a new MLNeuralNetworkFunctionInfo instance.
func NewMLNeuralNetworkFunctionInfo() MLNeuralNetworkFunctionInfo {
	class := getMLNeuralNetworkFunctionInfoClass()
	rv := objc.Send[MLNeuralNetworkFunctionInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewNeuralNetworkFunctionInfoWithCompiledModelArchiveCompilerVersionInfoError(archive unsafe.Pointer, info objectivec.IObject) (MLNeuralNetworkFunctionInfo, error) {
	var errorPtr objc.ID
	instance := getMLNeuralNetworkFunctionInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompiledModelArchive:compilerVersionInfo:error:"), archive, info, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkFunctionInfo{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkFunctionInfoFromID(rv), nil
}

func (m MLNeuralNetworkFunctionInfo) InitWithCompiledModelArchiveCompilerVersionInfoError(archive unsafe.Pointer, info objectivec.IObject) (MLNeuralNetworkFunctionInfo, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithCompiledModelArchive:compilerVersionInfo:error:"), archive, info, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkFunctionInfo{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkFunctionInfoFromID(rv), nil

}

func (m MLNeuralNetworkFunctionInfo) ClassLabels() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classLabels"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLNeuralNetworkFunctionInfo) ClassScoreVectorName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classScoreVectorName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLNeuralNetworkFunctionInfo) IsClassifier() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isClassifier"))
	return rv
}
func (m MLNeuralNetworkFunctionInfo) OutputNames() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
