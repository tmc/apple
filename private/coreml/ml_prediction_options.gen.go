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
func (m MLPredictionOptions) Init() MLPredictionOptions {
	rv := objc.Send[MLPredictionOptions](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLPredictionOptions) Autorelease() MLPredictionOptions {
	rv := objc.Send[MLPredictionOptions](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPredictionOptions creates a new MLPredictionOptions instance.
func NewMLPredictionOptions() MLPredictionOptions {
	class := getMLPredictionOptionsClass()
	rv := objc.Send[MLPredictionOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewPredictionOptionsWithCoder(coder objectivec.IObject) MLPredictionOptions {
	instance := getMLPredictionOptionsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLPredictionOptionsFromID(rv)
}

func NewPredictionOptionsWithUsesCPUOnly(cPUOnly bool) MLPredictionOptions {
	instance := getMLPredictionOptionsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUsesCPUOnly:"), cPUOnly)
	return MLPredictionOptionsFromID(rv)
}

func (m MLPredictionOptions) _validateDirectBindingExpectationsDirectlyBoundFeatureNamesUnexpectedDirectBindingsUnexpectedCopyBindings(expectations objectivec.IObject, names objectivec.IObject, bindings []objectivec.IObject, bindings2 []objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_validateDirectBindingExpectations:directlyBoundFeatureNames:unexpectedDirectBindings:unexpectedCopyBindings:"), expectations, names, objectivec.IObjectSliceToNSArray(bindings), objectivec.IObjectSliceToNSArray(bindings2))
	return rv
}

// ValidateDirectBindingExpectationsDirectlyBoundFeatureNamesUnexpectedDirectBindingsUnexpectedCopyBindings is an exported wrapper for the private method _validateDirectBindingExpectationsDirectlyBoundFeatureNamesUnexpectedDirectBindingsUnexpectedCopyBindings.
func (m MLPredictionOptions) ValidateDirectBindingExpectationsDirectlyBoundFeatureNamesUnexpectedDirectBindingsUnexpectedCopyBindings(expectations objectivec.IObject, names objectivec.IObject, bindings []objectivec.IObject, bindings2 []objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_validateDirectBindingExpectations:directlyBoundFeatureNames:unexpectedDirectBindings:unexpectedCopyBindings:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_validateDirectBindingExpectations:directlyBoundFeatureNames:unexpectedDirectBindings:unexpectedCopyBindings:"}
		return false, err
	}
	return m._validateDirectBindingExpectationsDirectlyBoundFeatureNamesUnexpectedDirectBindingsUnexpectedCopyBindings(expectations, names, bindings, bindings2), nil
}

// CanValidateDirectBindingExpectationsDirectlyBoundFeatureNamesUnexpectedDirectBindingsUnexpectedCopyBindings reports whether the receiver responds to the private selector _validateDirectBindingExpectations:directlyBoundFeatureNames:unexpectedDirectBindings:unexpectedCopyBindings:.
func (m MLPredictionOptions) CanValidateDirectBindingExpectationsDirectlyBoundFeatureNamesUnexpectedDirectBindingsUnexpectedCopyBindings() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_validateDirectBindingExpectations:directlyBoundFeatureNames:unexpectedDirectBindings:unexpectedCopyBindings:"))
}
func (m MLPredictionOptions) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (m MLPredictionOptions) ValidateExpectationsWithDirectlyBoundInputsOutputsError(inputs objectivec.IObject, outputs objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("validateExpectationsWithDirectlyBoundInputs:outputs:error:"), inputs, outputs, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateExpectationsWithDirectlyBoundInputs:outputs:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLPredictionOptions) InitWithCoder(coder foundation.INSCoder) MLPredictionOptions {
	rv := objc.Send[MLPredictionOptions](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (m MLPredictionOptions) InitWithUsesCPUOnly(cPUOnly bool) MLPredictionOptions {
	rv := objc.Send[MLPredictionOptions](m.ID, objc.Sel("initWithUsesCPUOnly:"), cPUOnly)
	return rv
}

func (_MLPredictionOptionsClass MLPredictionOptionsClass) DefaultOptions() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLPredictionOptionsClass.class), objc.Sel("defaultOptions"))
	return objectivec.Object{ID: rv}
}
func (_MLPredictionOptionsClass MLPredictionOptionsClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLPredictionOptionsClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (m MLPredictionOptions) AneExecutionPriority() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("aneExecutionPriority"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLPredictionOptions) SetAneExecutionPriority(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setAneExecutionPriority:"), objc.String(value))
}
func (m MLPredictionOptions) AneQoS() uint32 {
	rv := objc.Send[uint32](m.ID, objc.Sel("aneQoS"))
	return rv
}
func (m MLPredictionOptions) SetAneQoS(value uint32) {
	objc.Send[struct{}](m.ID, objc.Sel("setAneQoS:"), value)
}
func (m MLPredictionOptions) AutomaticOutputBackingMode() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("automaticOutputBackingMode"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLPredictionOptions) SetAutomaticOutputBackingMode(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setAutomaticOutputBackingMode:"), value)
}
func (m MLPredictionOptions) ClassifyTopK() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("classifyTopK"))
	return rv
}
func (m MLPredictionOptions) SetClassifyTopK(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("setClassifyTopK:"), value)
}
func (m MLPredictionOptions) CompletionSyncPoint() IMLPredictionSyncPoint {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("completionSyncPoint"))
	return MLPredictionSyncPointFromID(objc.ID(rv))
}
func (m MLPredictionOptions) SetCompletionSyncPoint(value IMLPredictionSyncPoint) {
	objc.Send[struct{}](m.ID, objc.Sel("setCompletionSyncPoint:"), value)
}
func (m MLPredictionOptions) E5rtStreamReuseExpectation() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("e5rtStreamReuseExpectation"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLPredictionOptions) SetE5rtStreamReuseExpectation(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setE5rtStreamReuseExpectation:"), objc.String(value))
}
func (m MLPredictionOptions) EnablePixelBufferDirectBinding() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("enablePixelBufferDirectBinding"))
	return rv
}
func (m MLPredictionOptions) SetEnablePixelBufferDirectBinding(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setEnablePixelBufferDirectBinding:"), value)
}
func (m MLPredictionOptions) HasDirectBindingExpectations() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasDirectBindingExpectations"))
	return rv
}
func (m MLPredictionOptions) InferenceFrameDataSerialization() IMLInferenceFrameDataSerialization {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inferenceFrameDataSerialization"))
	return MLInferenceFrameDataSerializationFromID(objc.ID(rv))
}
func (m MLPredictionOptions) SetInferenceFrameDataSerialization(value IMLInferenceFrameDataSerialization) {
	objc.Send[struct{}](m.ID, objc.Sel("setInferenceFrameDataSerialization:"), value)
}
func (m MLPredictionOptions) InputDirectBindingExpectations() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputDirectBindingExpectations"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLPredictionOptions) SetInputDirectBindingExpectations(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setInputDirectBindingExpectations:"), value)
}
func (m MLPredictionOptions) MaxComputationBatchSize() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("maxComputationBatchSize"))
	return rv
}
func (m MLPredictionOptions) SetMaxComputationBatchSize(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("setMaxComputationBatchSize:"), value)
}
func (m MLPredictionOptions) OutputDirectBindingExpectations() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputDirectBindingExpectations"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLPredictionOptions) SetOutputDirectBindingExpectations(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setOutputDirectBindingExpectations:"), value)
}
func (m MLPredictionOptions) ParentSignpostID() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("parentSignpostID"))
	return rv
}
func (m MLPredictionOptions) SetParentSignpostID(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("setParentSignpostID:"), value)
}
func (m MLPredictionOptions) PredictionUsesCPU() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("predictionUsesCPU"))
	return rv
}
func (m MLPredictionOptions) UsesCPUOnly() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("usesCPUOnly"))
	return rv
}
func (m MLPredictionOptions) SetUsesCPUOnly(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setUsesCPUOnly:"), value)
}
func (m MLPredictionOptions) WaitSyncPoints() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("waitSyncPoints"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLPredictionOptions) SetWaitSyncPoints(value foundation.INSArray) {
	objc.Send[struct{}](m.ID, objc.Sel("setWaitSyncPoints:"), value)
}
