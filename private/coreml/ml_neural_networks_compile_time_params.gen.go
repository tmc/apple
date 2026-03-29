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

// The class instance for the [MLNeuralNetworksCompileTimeParams] class.
var (
	_MLNeuralNetworksCompileTimeParamsClass     MLNeuralNetworksCompileTimeParamsClass
	_MLNeuralNetworksCompileTimeParamsClassOnce sync.Once
)

func getMLNeuralNetworksCompileTimeParamsClass() MLNeuralNetworksCompileTimeParamsClass {
	_MLNeuralNetworksCompileTimeParamsClassOnce.Do(func() {
		_MLNeuralNetworksCompileTimeParamsClass = MLNeuralNetworksCompileTimeParamsClass{class: objc.GetClass("MLNeuralNetworksCompileTimeParams")}
	})
	return _MLNeuralNetworksCompileTimeParamsClass
}

// GetMLNeuralNetworksCompileTimeParamsClass returns the class object for MLNeuralNetworksCompileTimeParams.
func GetMLNeuralNetworksCompileTimeParamsClass() MLNeuralNetworksCompileTimeParamsClass {
	return getMLNeuralNetworksCompileTimeParamsClass()
}

type MLNeuralNetworksCompileTimeParamsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNeuralNetworksCompileTimeParamsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNeuralNetworksCompileTimeParamsClass) Alloc() MLNeuralNetworksCompileTimeParams {
	rv := objc.Send[MLNeuralNetworksCompileTimeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLNeuralNetworksCompileTimeParams.EncodeWithCoder]
//   - [MLNeuralNetworksCompileTimeParams.LossParameters]
//   - [MLNeuralNetworksCompileTimeParams.SetLossParameters]
//   - [MLNeuralNetworksCompileTimeParams.LossType]
//   - [MLNeuralNetworksCompileTimeParams.SetLossType]
//   - [MLNeuralNetworksCompileTimeParams.OptimizerParameters]
//   - [MLNeuralNetworksCompileTimeParams.SetOptimizerParameters]
//   - [MLNeuralNetworksCompileTimeParams.OptimizerType]
//   - [MLNeuralNetworksCompileTimeParams.SetOptimizerType]
//   - [MLNeuralNetworksCompileTimeParams.TrainableLayerNames]
//   - [MLNeuralNetworksCompileTimeParams.SetTrainableLayerNames]
//   - [MLNeuralNetworksCompileTimeParams.UnarchiveUpdatableParamsAtURLError]
//   - [MLNeuralNetworksCompileTimeParams.UpdateParameters]
//   - [MLNeuralNetworksCompileTimeParams.SetUpdateParameters]
//   - [MLNeuralNetworksCompileTimeParams.WriteUpdatableParamsToURLError]
//   - [MLNeuralNetworksCompileTimeParams.InitWithCoder]
//   - [MLNeuralNetworksCompileTimeParams.InitWithLossTypeOptimizerTypeOptimizerParametersLossParametersTrainableLayerNamesUpdateParameters]
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworksCompileTimeParams
type MLNeuralNetworksCompileTimeParams struct {
	objectivec.Object
}

// MLNeuralNetworksCompileTimeParamsFromID constructs a [MLNeuralNetworksCompileTimeParams] from an objc.ID.
func MLNeuralNetworksCompileTimeParamsFromID(id objc.ID) MLNeuralNetworksCompileTimeParams {
	return MLNeuralNetworksCompileTimeParams{objectivec.Object{ID: id}}
}
// Ensure MLNeuralNetworksCompileTimeParams implements IMLNeuralNetworksCompileTimeParams.
var _ IMLNeuralNetworksCompileTimeParams = MLNeuralNetworksCompileTimeParams{}

// An interface definition for the [MLNeuralNetworksCompileTimeParams] class.
//
// # Methods
//
//   - [IMLNeuralNetworksCompileTimeParams.EncodeWithCoder]
//   - [IMLNeuralNetworksCompileTimeParams.LossParameters]
//   - [IMLNeuralNetworksCompileTimeParams.SetLossParameters]
//   - [IMLNeuralNetworksCompileTimeParams.LossType]
//   - [IMLNeuralNetworksCompileTimeParams.SetLossType]
//   - [IMLNeuralNetworksCompileTimeParams.OptimizerParameters]
//   - [IMLNeuralNetworksCompileTimeParams.SetOptimizerParameters]
//   - [IMLNeuralNetworksCompileTimeParams.OptimizerType]
//   - [IMLNeuralNetworksCompileTimeParams.SetOptimizerType]
//   - [IMLNeuralNetworksCompileTimeParams.TrainableLayerNames]
//   - [IMLNeuralNetworksCompileTimeParams.SetTrainableLayerNames]
//   - [IMLNeuralNetworksCompileTimeParams.UnarchiveUpdatableParamsAtURLError]
//   - [IMLNeuralNetworksCompileTimeParams.UpdateParameters]
//   - [IMLNeuralNetworksCompileTimeParams.SetUpdateParameters]
//   - [IMLNeuralNetworksCompileTimeParams.WriteUpdatableParamsToURLError]
//   - [IMLNeuralNetworksCompileTimeParams.InitWithCoder]
//   - [IMLNeuralNetworksCompileTimeParams.InitWithLossTypeOptimizerTypeOptimizerParametersLossParametersTrainableLayerNamesUpdateParameters]
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworksCompileTimeParams
type IMLNeuralNetworksCompileTimeParams interface {
	objectivec.IObject

	// Topic: Methods

	EncodeWithCoder(coder foundation.INSCoder)
	LossParameters() foundation.INSDictionary
	SetLossParameters(value foundation.INSDictionary)
	LossType() int64
	SetLossType(value int64)
	OptimizerParameters() foundation.INSDictionary
	SetOptimizerParameters(value foundation.INSDictionary)
	OptimizerType() int64
	SetOptimizerType(value int64)
	TrainableLayerNames() foundation.INSArray
	SetTrainableLayerNames(value foundation.INSArray)
	UnarchiveUpdatableParamsAtURLError(url foundation.INSURL) (objectivec.IObject, error)
	UpdateParameters() foundation.INSDictionary
	SetUpdateParameters(value foundation.INSDictionary)
	WriteUpdatableParamsToURLError(url foundation.INSURL) (bool, error)
	InitWithCoder(coder foundation.INSCoder) MLNeuralNetworksCompileTimeParams
	InitWithLossTypeOptimizerTypeOptimizerParametersLossParametersTrainableLayerNamesUpdateParameters(type_ int64, type_2 int64, parameters objectivec.IObject, parameters2 objectivec.IObject, names objectivec.IObject, parameters3 objectivec.IObject) MLNeuralNetworksCompileTimeParams
}

// Init initializes the instance.
func (n MLNeuralNetworksCompileTimeParams) Init() MLNeuralNetworksCompileTimeParams {
	rv := objc.Send[MLNeuralNetworksCompileTimeParams](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MLNeuralNetworksCompileTimeParams) Autorelease() MLNeuralNetworksCompileTimeParams {
	rv := objc.Send[MLNeuralNetworksCompileTimeParams](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNeuralNetworksCompileTimeParams creates a new MLNeuralNetworksCompileTimeParams instance.
func NewMLNeuralNetworksCompileTimeParams() MLNeuralNetworksCompileTimeParams {
	class := getMLNeuralNetworksCompileTimeParamsClass()
	rv := objc.Send[MLNeuralNetworksCompileTimeParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworksCompileTimeParams/initWithCoder:
func NewNeuralNetworksCompileTimeParamsWithCoder(coder objectivec.IObject) MLNeuralNetworksCompileTimeParams {
	instance := getMLNeuralNetworksCompileTimeParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLNeuralNetworksCompileTimeParamsFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworksCompileTimeParams/initWithLossType:optimizerType:optimizerParameters:lossParameters:trainableLayerNames:updateParameters:
func NewNeuralNetworksCompileTimeParamsWithLossTypeOptimizerTypeOptimizerParametersLossParametersTrainableLayerNamesUpdateParameters(type_ int64, type_2 int64, parameters objectivec.IObject, parameters2 objectivec.IObject, names objectivec.IObject, parameters3 objectivec.IObject) MLNeuralNetworksCompileTimeParams {
	instance := getMLNeuralNetworksCompileTimeParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLossType:optimizerType:optimizerParameters:lossParameters:trainableLayerNames:updateParameters:"), type_, type_2, parameters, parameters2, names, parameters3)
	return MLNeuralNetworksCompileTimeParamsFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworksCompileTimeParams/encodeWithCoder:
func (n MLNeuralNetworksCompileTimeParams) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworksCompileTimeParams/unarchiveUpdatableParamsAtURL:error:
func (n MLNeuralNetworksCompileTimeParams) UnarchiveUpdatableParamsAtURLError(url foundation.INSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](n.ID, objc.Sel("unarchiveUpdatableParamsAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworksCompileTimeParams/writeUpdatableParamsToURL:error:
func (n MLNeuralNetworksCompileTimeParams) WriteUpdatableParamsToURLError(url foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](n.ID, objc.Sel("writeUpdatableParamsToURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeUpdatableParamsToURL:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworksCompileTimeParams/initWithCoder:
func (n MLNeuralNetworksCompileTimeParams) InitWithCoder(coder foundation.INSCoder) MLNeuralNetworksCompileTimeParams {
	rv := objc.Send[MLNeuralNetworksCompileTimeParams](n.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworksCompileTimeParams/initWithLossType:optimizerType:optimizerParameters:lossParameters:trainableLayerNames:updateParameters:
func (n MLNeuralNetworksCompileTimeParams) InitWithLossTypeOptimizerTypeOptimizerParametersLossParametersTrainableLayerNamesUpdateParameters(type_ int64, type_2 int64, parameters objectivec.IObject, parameters2 objectivec.IObject, names objectivec.IObject, parameters3 objectivec.IObject) MLNeuralNetworksCompileTimeParams {
	rv := objc.Send[MLNeuralNetworksCompileTimeParams](n.ID, objc.Sel("initWithLossType:optimizerType:optimizerParameters:lossParameters:trainableLayerNames:updateParameters:"), type_, type_2, parameters, parameters2, names, parameters3)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworksCompileTimeParams/supportsSecureCoding
func (_MLNeuralNetworksCompileTimeParamsClass MLNeuralNetworksCompileTimeParamsClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLNeuralNetworksCompileTimeParamsClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworksCompileTimeParams/lossParameters
func (n MLNeuralNetworksCompileTimeParams) LossParameters() foundation.INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("lossParameters"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (n MLNeuralNetworksCompileTimeParams) SetLossParameters(value foundation.INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setLossParameters:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworksCompileTimeParams/lossType
func (n MLNeuralNetworksCompileTimeParams) LossType() int64 {
	rv := objc.Send[int64](n.ID, objc.Sel("lossType"))
	return rv
}
func (n MLNeuralNetworksCompileTimeParams) SetLossType(value int64) {
	objc.Send[struct{}](n.ID, objc.Sel("setLossType:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworksCompileTimeParams/optimizerParameters
func (n MLNeuralNetworksCompileTimeParams) OptimizerParameters() foundation.INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("optimizerParameters"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (n MLNeuralNetworksCompileTimeParams) SetOptimizerParameters(value foundation.INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setOptimizerParameters:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworksCompileTimeParams/optimizerType
func (n MLNeuralNetworksCompileTimeParams) OptimizerType() int64 {
	rv := objc.Send[int64](n.ID, objc.Sel("optimizerType"))
	return rv
}
func (n MLNeuralNetworksCompileTimeParams) SetOptimizerType(value int64) {
	objc.Send[struct{}](n.ID, objc.Sel("setOptimizerType:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworksCompileTimeParams/trainableLayerNames
func (n MLNeuralNetworksCompileTimeParams) TrainableLayerNames() foundation.INSArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("trainableLayerNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (n MLNeuralNetworksCompileTimeParams) SetTrainableLayerNames(value foundation.INSArray) {
	objc.Send[struct{}](n.ID, objc.Sel("setTrainableLayerNames:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworksCompileTimeParams/updateParameters
func (n MLNeuralNetworksCompileTimeParams) UpdateParameters() foundation.INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("updateParameters"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (n MLNeuralNetworksCompileTimeParams) SetUpdateParameters(value foundation.INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setUpdateParameters:"), value)
}

