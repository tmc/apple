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
	rv := objc.SendIfResponds[MLNeuralNetworksCompileTimeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

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
	UnarchiveUpdatableParamsAtURLError(url foundation.NSURL) (objectivec.IObject, error)
	UpdateParameters() foundation.INSDictionary
	SetUpdateParameters(value foundation.INSDictionary)
	WriteUpdatableParamsToURLError(url foundation.NSURL) (bool, error)
	InitWithCoder(coder foundation.INSCoder) MLNeuralNetworksCompileTimeParams
	InitWithLossTypeOptimizerTypeOptimizerParametersLossParametersTrainableLayerNamesUpdateParameters(type_ int64, type_2 int64, parameters objectivec.IObject, parameters2 objectivec.IObject, names objectivec.IObject, parameters3 objectivec.IObject) MLNeuralNetworksCompileTimeParams
}

// Init initializes the instance.
func (m MLNeuralNetworksCompileTimeParams) Init() MLNeuralNetworksCompileTimeParams {
	rv := objc.SendIfResponds[MLNeuralNetworksCompileTimeParams](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLNeuralNetworksCompileTimeParams) Autorelease() MLNeuralNetworksCompileTimeParams {
	rv := objc.SendIfResponds[MLNeuralNetworksCompileTimeParams](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNeuralNetworksCompileTimeParams creates a new MLNeuralNetworksCompileTimeParams instance.
func NewMLNeuralNetworksCompileTimeParams() MLNeuralNetworksCompileTimeParams {
	class := getMLNeuralNetworksCompileTimeParamsClass()
	rv := objc.SendIfResponds[MLNeuralNetworksCompileTimeParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewNeuralNetworksCompileTimeParamsWithCoder(coder objectivec.IObject) MLNeuralNetworksCompileTimeParams {
	instance := getMLNeuralNetworksCompileTimeParamsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLNeuralNetworksCompileTimeParamsFromID(rv)
}

func NewNeuralNetworksCompileTimeParamsWithLossTypeOptimizerTypeOptimizerParametersLossParametersTrainableLayerNamesUpdateParameters(type_ int64, type_2 int64, parameters objectivec.IObject, parameters2 objectivec.IObject, names objectivec.IObject, parameters3 objectivec.IObject) MLNeuralNetworksCompileTimeParams {
	instance := getMLNeuralNetworksCompileTimeParamsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithLossType:optimizerType:optimizerParameters:lossParameters:trainableLayerNames:updateParameters:"), type_, type_2, parameters, parameters2, names, parameters3)
	return MLNeuralNetworksCompileTimeParamsFromID(rv)
}

func (m MLNeuralNetworksCompileTimeParams) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (m MLNeuralNetworksCompileTimeParams) UnarchiveUpdatableParamsAtURLError(url foundation.NSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("unarchiveUpdatableParamsAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNeuralNetworksCompileTimeParams) WriteUpdatableParamsToURLError(url foundation.NSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("writeUpdatableParamsToURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeUpdatableParamsToURL:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworksCompileTimeParams) InitWithCoder(coder foundation.INSCoder) MLNeuralNetworksCompileTimeParams {
	rv := objc.SendIfResponds[MLNeuralNetworksCompileTimeParams](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (m MLNeuralNetworksCompileTimeParams) InitWithLossTypeOptimizerTypeOptimizerParametersLossParametersTrainableLayerNamesUpdateParameters(type_ int64, type_2 int64, parameters objectivec.IObject, parameters2 objectivec.IObject, names objectivec.IObject, parameters3 objectivec.IObject) MLNeuralNetworksCompileTimeParams {
	rv := objc.SendIfResponds[MLNeuralNetworksCompileTimeParams](m.ID, objc.Sel("initWithLossType:optimizerType:optimizerParameters:lossParameters:trainableLayerNames:updateParameters:"), type_, type_2, parameters, parameters2, names, parameters3)
	return rv
}

func (_MLNeuralNetworksCompileTimeParamsClass MLNeuralNetworksCompileTimeParamsClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLNeuralNetworksCompileTimeParamsClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (m MLNeuralNetworksCompileTimeParams) LossParameters() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("lossParameters"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLNeuralNetworksCompileTimeParams) SetLossParameters(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setLossParameters:"), value)
}
func (m MLNeuralNetworksCompileTimeParams) LossType() int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("lossType"))
	return rv
}
func (m MLNeuralNetworksCompileTimeParams) SetLossType(value int64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setLossType:"), value)
}
func (m MLNeuralNetworksCompileTimeParams) OptimizerParameters() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("optimizerParameters"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLNeuralNetworksCompileTimeParams) SetOptimizerParameters(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setOptimizerParameters:"), value)
}
func (m MLNeuralNetworksCompileTimeParams) OptimizerType() int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("optimizerType"))
	return rv
}
func (m MLNeuralNetworksCompileTimeParams) SetOptimizerType(value int64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setOptimizerType:"), value)
}
func (m MLNeuralNetworksCompileTimeParams) TrainableLayerNames() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("trainableLayerNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLNeuralNetworksCompileTimeParams) SetTrainableLayerNames(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setTrainableLayerNames:"), value)
}
func (m MLNeuralNetworksCompileTimeParams) UpdateParameters() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("updateParameters"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLNeuralNetworksCompileTimeParams) SetUpdateParameters(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setUpdateParameters:"), value)
}
