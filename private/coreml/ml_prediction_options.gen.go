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

// The class instance for the [MLPredictionOptions] class.
var (
	_MLPredictionOptionsClass     MLPredictionOptionsClass
	_MLPredictionOptionsClassOnce sync.Once
)

func getMLPredictionOptionsClass() MLPredictionOptionsClass {
	_MLPredictionOptionsClassOnce.Do(func() {
		_MLPredictionOptionsClass = MLPredictionOptionsClass{class: objc.GetClass("MLPredictionOptions")}
	})
	return _MLPredictionOptionsClass
}

// GetMLPredictionOptionsClass returns the class object for MLPredictionOptions.
func GetMLPredictionOptionsClass() MLPredictionOptionsClass {
	return getMLPredictionOptionsClass()
}

type MLPredictionOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLPredictionOptionsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLPredictionOptionsClass) Alloc() MLPredictionOptions {
	rv := objc.Send[MLPredictionOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLPredictionOptions._validateDirectBindingExpectationsDirectlyBoundFeatureNamesUnexpectedDirectBindingsUnexpectedCopyBindings]
//   - [MLPredictionOptions.AneExecutionPriority]
//   - [MLPredictionOptions.SetAneExecutionPriority]
//   - [MLPredictionOptions.AneQoS]
//   - [MLPredictionOptions.SetAneQoS]
//   - [MLPredictionOptions.AutomaticOutputBackingMode]
//   - [MLPredictionOptions.SetAutomaticOutputBackingMode]
//   - [MLPredictionOptions.ClassifyTopK]
//   - [MLPredictionOptions.SetClassifyTopK]
//   - [MLPredictionOptions.CompletionSyncPoint]
//   - [MLPredictionOptions.SetCompletionSyncPoint]
//   - [MLPredictionOptions.E5rtStreamReuseExpectation]
//   - [MLPredictionOptions.SetE5rtStreamReuseExpectation]
//   - [MLPredictionOptions.EnablePixelBufferDirectBinding]
//   - [MLPredictionOptions.SetEnablePixelBufferDirectBinding]
//   - [MLPredictionOptions.EncodeWithCoder]
//   - [MLPredictionOptions.HasDirectBindingExpectations]
//   - [MLPredictionOptions.InferenceFrameDataSerialization]
//   - [MLPredictionOptions.SetInferenceFrameDataSerialization]
//   - [MLPredictionOptions.InputDirectBindingExpectations]
//   - [MLPredictionOptions.SetInputDirectBindingExpectations]
//   - [MLPredictionOptions.MaxComputationBatchSize]
//   - [MLPredictionOptions.SetMaxComputationBatchSize]
//   - [MLPredictionOptions.OutputDirectBindingExpectations]
//   - [MLPredictionOptions.SetOutputDirectBindingExpectations]
//   - [MLPredictionOptions.ParentSignpostID]
//   - [MLPredictionOptions.SetParentSignpostID]
//   - [MLPredictionOptions.PredictionUsesCPU]
//   - [MLPredictionOptions.UsesCPUOnly]
//   - [MLPredictionOptions.SetUsesCPUOnly]
//   - [MLPredictionOptions.ValidateExpectationsWithDirectlyBoundInputsOutputsError]
//   - [MLPredictionOptions.WaitSyncPoints]
//   - [MLPredictionOptions.SetWaitSyncPoints]
//   - [MLPredictionOptions.InitWithCoder]
//   - [MLPredictionOptions.InitWithUsesCPUOnly]
//
// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions
type MLPredictionOptions struct {
	objectivec.Object
}

// MLPredictionOptionsFromID constructs a [MLPredictionOptions] from an objc.ID.
func MLPredictionOptionsFromID(id objc.ID) MLPredictionOptions {
	return MLPredictionOptions{objectivec.Object{ID: id}}
}

// Ensure MLPredictionOptions implements IMLPredictionOptions.
var _ IMLPredictionOptions = MLPredictionOptions{}

// An interface definition for the [MLPredictionOptions] class.
//
// # Methods
//
//   - [IMLPredictionOptions._validateDirectBindingExpectationsDirectlyBoundFeatureNamesUnexpectedDirectBindingsUnexpectedCopyBindings]
//   - [IMLPredictionOptions.AneExecutionPriority]
//   - [IMLPredictionOptions.SetAneExecutionPriority]
//   - [IMLPredictionOptions.AneQoS]
//   - [IMLPredictionOptions.SetAneQoS]
//   - [IMLPredictionOptions.AutomaticOutputBackingMode]
//   - [IMLPredictionOptions.SetAutomaticOutputBackingMode]
//   - [IMLPredictionOptions.ClassifyTopK]
//   - [IMLPredictionOptions.SetClassifyTopK]
//   - [IMLPredictionOptions.CompletionSyncPoint]
//   - [IMLPredictionOptions.SetCompletionSyncPoint]
//   - [IMLPredictionOptions.E5rtStreamReuseExpectation]
//   - [IMLPredictionOptions.SetE5rtStreamReuseExpectation]
//   - [IMLPredictionOptions.EnablePixelBufferDirectBinding]
//   - [IMLPredictionOptions.SetEnablePixelBufferDirectBinding]
//   - [IMLPredictionOptions.EncodeWithCoder]
//   - [IMLPredictionOptions.HasDirectBindingExpectations]
//   - [IMLPredictionOptions.InferenceFrameDataSerialization]
//   - [IMLPredictionOptions.SetInferenceFrameDataSerialization]
//   - [IMLPredictionOptions.InputDirectBindingExpectations]
//   - [IMLPredictionOptions.SetInputDirectBindingExpectations]
//   - [IMLPredictionOptions.MaxComputationBatchSize]
//   - [IMLPredictionOptions.SetMaxComputationBatchSize]
//   - [IMLPredictionOptions.OutputDirectBindingExpectations]
//   - [IMLPredictionOptions.SetOutputDirectBindingExpectations]
//   - [IMLPredictionOptions.ParentSignpostID]
//   - [IMLPredictionOptions.SetParentSignpostID]
//   - [IMLPredictionOptions.PredictionUsesCPU]
//   - [IMLPredictionOptions.UsesCPUOnly]
//   - [IMLPredictionOptions.SetUsesCPUOnly]
//   - [IMLPredictionOptions.ValidateExpectationsWithDirectlyBoundInputsOutputsError]
//   - [IMLPredictionOptions.WaitSyncPoints]
//   - [IMLPredictionOptions.SetWaitSyncPoints]
//   - [IMLPredictionOptions.InitWithCoder]
//   - [IMLPredictionOptions.InitWithUsesCPUOnly]
//
// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions
type IMLPredictionOptions interface {
	objectivec.IObject

	// Topic: Methods

	_validateDirectBindingExpectationsDirectlyBoundFeatureNamesUnexpectedDirectBindingsUnexpectedCopyBindings(expectations objectivec.IObject, names objectivec.IObject, bindings []objectivec.IObject, bindings2 []objectivec.IObject) bool
	AneExecutionPriority() string
	SetAneExecutionPriority(value string)
	AneQoS() uint32
	SetAneQoS(value uint32)
	AutomaticOutputBackingMode() foundation.INSDictionary
	SetAutomaticOutputBackingMode(value foundation.INSDictionary)
	ClassifyTopK() uint64
	SetClassifyTopK(value uint64)
	CompletionSyncPoint() IMLPredictionSyncPoint
	SetCompletionSyncPoint(value IMLPredictionSyncPoint)
	E5rtStreamReuseExpectation() string
	SetE5rtStreamReuseExpectation(value string)
	EnablePixelBufferDirectBinding() bool
	SetEnablePixelBufferDirectBinding(value bool)
	EncodeWithCoder(coder foundation.INSCoder)
	HasDirectBindingExpectations() bool
	InferenceFrameDataSerialization() IMLInferenceFrameDataSerialization
	SetInferenceFrameDataSerialization(value IMLInferenceFrameDataSerialization)
	InputDirectBindingExpectations() foundation.INSDictionary
	SetInputDirectBindingExpectations(value foundation.INSDictionary)
	MaxComputationBatchSize() uint64
	SetMaxComputationBatchSize(value uint64)
	OutputDirectBindingExpectations() foundation.INSDictionary
	SetOutputDirectBindingExpectations(value foundation.INSDictionary)
	ParentSignpostID() uint64
	SetParentSignpostID(value uint64)
	PredictionUsesCPU() bool
	UsesCPUOnly() bool
	SetUsesCPUOnly(value bool)
	ValidateExpectationsWithDirectlyBoundInputsOutputsError(inputs objectivec.IObject, outputs objectivec.IObject) (bool, error)
	WaitSyncPoints() foundation.INSArray
	SetWaitSyncPoints(value foundation.INSArray)
	InitWithCoder(coder foundation.INSCoder) MLPredictionOptions
	InitWithUsesCPUOnly(cPUOnly bool) MLPredictionOptions
}

// Init initializes the instance.
func (p MLPredictionOptions) Init() MLPredictionOptions {
	rv := objc.Send[MLPredictionOptions](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLPredictionOptions) Autorelease() MLPredictionOptions {
	rv := objc.Send[MLPredictionOptions](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPredictionOptions creates a new MLPredictionOptions instance.
func NewMLPredictionOptions() MLPredictionOptions {
	class := getMLPredictionOptionsClass()
	rv := objc.Send[MLPredictionOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/initWithCoder:
func NewPredictionOptionsWithCoder(coder objectivec.IObject) MLPredictionOptions {
	instance := getMLPredictionOptionsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLPredictionOptionsFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/initWithUsesCPUOnly:
func NewPredictionOptionsWithUsesCPUOnly(cPUOnly bool) MLPredictionOptions {
	instance := getMLPredictionOptionsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUsesCPUOnly:"), cPUOnly)
	return MLPredictionOptionsFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/_validateDirectBindingExpectations:directlyBoundFeatureNames:unexpectedDirectBindings:unexpectedCopyBindings:
func (p MLPredictionOptions) _validateDirectBindingExpectationsDirectlyBoundFeatureNamesUnexpectedDirectBindingsUnexpectedCopyBindings(expectations objectivec.IObject, names objectivec.IObject, bindings []objectivec.IObject, bindings2 []objectivec.IObject) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("_validateDirectBindingExpectations:directlyBoundFeatureNames:unexpectedDirectBindings:unexpectedCopyBindings:"), expectations, names, objectivec.IObjectSliceToNSArray(bindings), objectivec.IObjectSliceToNSArray(bindings2))
	return rv
}

// ValidateDirectBindingExpectationsDirectlyBoundFeatureNamesUnexpectedDirectBindingsUnexpectedCopyBindings is an exported wrapper for the private method _validateDirectBindingExpectationsDirectlyBoundFeatureNamesUnexpectedDirectBindingsUnexpectedCopyBindings.
func (p MLPredictionOptions) ValidateDirectBindingExpectationsDirectlyBoundFeatureNamesUnexpectedDirectBindingsUnexpectedCopyBindings(expectations objectivec.IObject, names objectivec.IObject, bindings []objectivec.IObject, bindings2 []objectivec.IObject) bool {
	return p._validateDirectBindingExpectationsDirectlyBoundFeatureNamesUnexpectedDirectBindingsUnexpectedCopyBindings(expectations, names, bindings, bindings2)
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/encodeWithCoder:
func (p MLPredictionOptions) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/validateExpectationsWithDirectlyBoundInputs:outputs:error:
func (p MLPredictionOptions) ValidateExpectationsWithDirectlyBoundInputsOutputsError(inputs objectivec.IObject, outputs objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](p.ID, objc.Sel("validateExpectationsWithDirectlyBoundInputs:outputs:error:"), inputs, outputs, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateExpectationsWithDirectlyBoundInputs:outputs:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/initWithCoder:
func (p MLPredictionOptions) InitWithCoder(coder foundation.INSCoder) MLPredictionOptions {
	rv := objc.Send[MLPredictionOptions](p.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/initWithUsesCPUOnly:
func (p MLPredictionOptions) InitWithUsesCPUOnly(cPUOnly bool) MLPredictionOptions {
	rv := objc.Send[MLPredictionOptions](p.ID, objc.Sel("initWithUsesCPUOnly:"), cPUOnly)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/defaultOptions
func (_MLPredictionOptionsClass MLPredictionOptionsClass) DefaultOptions() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLPredictionOptionsClass.class), objc.Sel("defaultOptions"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/supportsSecureCoding
func (_MLPredictionOptionsClass MLPredictionOptionsClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLPredictionOptionsClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/aneExecutionPriority
func (p MLPredictionOptions) AneExecutionPriority() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("aneExecutionPriority"))
	return foundation.NSStringFromID(rv).String()
}
func (p MLPredictionOptions) SetAneExecutionPriority(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setAneExecutionPriority:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/aneQoS
func (p MLPredictionOptions) AneQoS() uint32 {
	rv := objc.Send[uint32](p.ID, objc.Sel("aneQoS"))
	return rv
}
func (p MLPredictionOptions) SetAneQoS(value uint32) {
	objc.Send[struct{}](p.ID, objc.Sel("setAneQoS:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/automaticOutputBackingMode
func (p MLPredictionOptions) AutomaticOutputBackingMode() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("automaticOutputBackingMode"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (p MLPredictionOptions) SetAutomaticOutputBackingMode(value foundation.INSDictionary) {
	objc.Send[struct{}](p.ID, objc.Sel("setAutomaticOutputBackingMode:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/classifyTopK
func (p MLPredictionOptions) ClassifyTopK() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("classifyTopK"))
	return rv
}
func (p MLPredictionOptions) SetClassifyTopK(value uint64) {
	objc.Send[struct{}](p.ID, objc.Sel("setClassifyTopK:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/completionSyncPoint
func (p MLPredictionOptions) CompletionSyncPoint() IMLPredictionSyncPoint {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("completionSyncPoint"))
	return MLPredictionSyncPointFromID(objc.ID(rv))
}
func (p MLPredictionOptions) SetCompletionSyncPoint(value IMLPredictionSyncPoint) {
	objc.Send[struct{}](p.ID, objc.Sel("setCompletionSyncPoint:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/e5rtStreamReuseExpectation
func (p MLPredictionOptions) E5rtStreamReuseExpectation() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("e5rtStreamReuseExpectation"))
	return foundation.NSStringFromID(rv).String()
}
func (p MLPredictionOptions) SetE5rtStreamReuseExpectation(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setE5rtStreamReuseExpectation:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/enablePixelBufferDirectBinding
func (p MLPredictionOptions) EnablePixelBufferDirectBinding() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("enablePixelBufferDirectBinding"))
	return rv
}
func (p MLPredictionOptions) SetEnablePixelBufferDirectBinding(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setEnablePixelBufferDirectBinding:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/hasDirectBindingExpectations
func (p MLPredictionOptions) HasDirectBindingExpectations() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("hasDirectBindingExpectations"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/inferenceFrameDataSerialization
func (p MLPredictionOptions) InferenceFrameDataSerialization() IMLInferenceFrameDataSerialization {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("inferenceFrameDataSerialization"))
	return MLInferenceFrameDataSerializationFromID(objc.ID(rv))
}
func (p MLPredictionOptions) SetInferenceFrameDataSerialization(value IMLInferenceFrameDataSerialization) {
	objc.Send[struct{}](p.ID, objc.Sel("setInferenceFrameDataSerialization:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/inputDirectBindingExpectations
func (p MLPredictionOptions) InputDirectBindingExpectations() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("inputDirectBindingExpectations"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (p MLPredictionOptions) SetInputDirectBindingExpectations(value foundation.INSDictionary) {
	objc.Send[struct{}](p.ID, objc.Sel("setInputDirectBindingExpectations:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/maxComputationBatchSize
func (p MLPredictionOptions) MaxComputationBatchSize() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("maxComputationBatchSize"))
	return rv
}
func (p MLPredictionOptions) SetMaxComputationBatchSize(value uint64) {
	objc.Send[struct{}](p.ID, objc.Sel("setMaxComputationBatchSize:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/outputDirectBindingExpectations
func (p MLPredictionOptions) OutputDirectBindingExpectations() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("outputDirectBindingExpectations"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (p MLPredictionOptions) SetOutputDirectBindingExpectations(value foundation.INSDictionary) {
	objc.Send[struct{}](p.ID, objc.Sel("setOutputDirectBindingExpectations:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/parentSignpostID
func (p MLPredictionOptions) ParentSignpostID() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("parentSignpostID"))
	return rv
}
func (p MLPredictionOptions) SetParentSignpostID(value uint64) {
	objc.Send[struct{}](p.ID, objc.Sel("setParentSignpostID:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/predictionUsesCPU
func (p MLPredictionOptions) PredictionUsesCPU() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("predictionUsesCPU"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/usesCPUOnly
func (p MLPredictionOptions) UsesCPUOnly() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("usesCPUOnly"))
	return rv
}
func (p MLPredictionOptions) SetUsesCPUOnly(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setUsesCPUOnly:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/waitSyncPoints
func (p MLPredictionOptions) WaitSyncPoints() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("waitSyncPoints"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (p MLPredictionOptions) SetWaitSyncPoints(value foundation.INSArray) {
	objc.Send[struct{}](p.ID, objc.Sel("setWaitSyncPoints:"), value)
}
