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

// The class instance for the [MLE5ExecutionStreamOperation] class.
var (
	_MLE5ExecutionStreamOperationClass     MLE5ExecutionStreamOperationClass
	_MLE5ExecutionStreamOperationClassOnce sync.Once
)

func getMLE5ExecutionStreamOperationClass() MLE5ExecutionStreamOperationClass {
	_MLE5ExecutionStreamOperationClassOnce.Do(func() {
		_MLE5ExecutionStreamOperationClass = MLE5ExecutionStreamOperationClass{class: objc.GetClass("MLE5ExecutionStreamOperation")}
	})
	return _MLE5ExecutionStreamOperationClass
}

// GetMLE5ExecutionStreamOperationClass returns the class object for MLE5ExecutionStreamOperation.
func GetMLE5ExecutionStreamOperationClass() MLE5ExecutionStreamOperationClass {
	return getMLE5ExecutionStreamOperationClass()
}

type MLE5ExecutionStreamOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLE5ExecutionStreamOperationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLE5ExecutionStreamOperationClass) Alloc() MLE5ExecutionStreamOperation {
	rv := objc.Send[MLE5ExecutionStreamOperation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLE5ExecutionStreamOperation._bindCompletionSyncPointDirectlyIfPossible]
//   - [MLE5ExecutionStreamOperation._bindEventToWaitForCopyingInputFeaturesAfterSyncPoints]
//   - [MLE5ExecutionStreamOperation._bindInputFeaturesAndWaitEventsOptionsError]
//   - [MLE5ExecutionStreamOperation._bindNewCompletionEventsDirectlyWithCompletionSyncPoint]
//   - [MLE5ExecutionStreamOperation._bindNewWaitEventsDirectlyWithWaitSyncPoints]
//   - [MLE5ExecutionStreamOperation._bindOutputPortsWithOptionsError]
//   - [MLE5ExecutionStreamOperation._bindWaitEventsDirectly]
//   - [MLE5ExecutionStreamOperation._copyInputFeaturesError]
//   - [MLE5ExecutionStreamOperation._createOperationAndReturnError]
//   - [MLE5ExecutionStreamOperation._createOperationWithRetryCountError]
//   - [MLE5ExecutionStreamOperation._directlyBoundFeatureNamesForPorts]
//   - [MLE5ExecutionStreamOperation._inoutPortNames]
//   - [MLE5ExecutionStreamOperation._inputPortNames]
//   - [MLE5ExecutionStreamOperation._multiArrayFeatureFromStateFeature]
//   - [MLE5ExecutionStreamOperation._newArrayOfInoutPortsFeatureDescriptionsByNameError]
//   - [MLE5ExecutionStreamOperation._newArrayOfInputPortsFeatureDescriptionsByNameError]
//   - [MLE5ExecutionStreamOperation._newArrayOfOutputPortsFeatureDescriptionsByNameError]
//   - [MLE5ExecutionStreamOperation._outputPortNames]
//   - [MLE5ExecutionStreamOperation._prepareInputPortsForFeaturesError]
//   - [MLE5ExecutionStreamOperation._reusableForCompletionSyncPointAllOutputBackingsUseDirectBinding]
//   - [MLE5ExecutionStreamOperation._reusableForWaitSyncPointsAllInputsUseDirectBinding]
//   - [MLE5ExecutionStreamOperation._updateCompletionEventFutureValuesWithCompletionSyncPoint]
//   - [MLE5ExecutionStreamOperation._updateWaitEventFutureValuesWithWaitSyncPoints]
//   - [MLE5ExecutionStreamOperation.AsyncSubmissionError]
//   - [MLE5ExecutionStreamOperation.SetAsyncSubmissionError]
//   - [MLE5ExecutionStreamOperation.CompletionSharedEventBoundToESOP]
//   - [MLE5ExecutionStreamOperation.SetCompletionSharedEventBoundToESOP]
//   - [MLE5ExecutionStreamOperation.DebugLabel]
//   - [MLE5ExecutionStreamOperation.DirectlyBoundInputFeatureNames]
//   - [MLE5ExecutionStreamOperation.DirectlyBoundOutputFeatureNames]
//   - [MLE5ExecutionStreamOperation.FunctionName]
//   - [MLE5ExecutionStreamOperation.InputPorts]
//   - [MLE5ExecutionStreamOperation.SetInputPorts]
//   - [MLE5ExecutionStreamOperation.ModelConfiguration]
//   - [MLE5ExecutionStreamOperation.ModelDescription]
//   - [MLE5ExecutionStreamOperation.ModelSignpostId]
//   - [MLE5ExecutionStreamOperation.OperationHandle]
//   - [MLE5ExecutionStreamOperation.SetOperationHandle]
//   - [MLE5ExecutionStreamOperation.OutputFeatures]
//   - [MLE5ExecutionStreamOperation.OutputPorts]
//   - [MLE5ExecutionStreamOperation.SetOutputPorts]
//   - [MLE5ExecutionStreamOperation.PixelBufferPool]
//   - [MLE5ExecutionStreamOperation.SetPixelBufferPool]
//   - [MLE5ExecutionStreamOperation.PreloadAndReturnError]
//   - [MLE5ExecutionStreamOperation.PrepareAsyncSubmissionForInputFeaturesOptionsError]
//   - [MLE5ExecutionStreamOperation.PrepareForInputFeaturesOptionsError]
//   - [MLE5ExecutionStreamOperation.ProgramLibrary]
//   - [MLE5ExecutionStreamOperation.Reset]
//   - [MLE5ExecutionStreamOperation.ReusableForInputFeaturesOptions]
//   - [MLE5ExecutionStreamOperation.SerializeInferenceFrameDataForOptionsError]
//   - [MLE5ExecutionStreamOperation.ShapeHash]
//   - [MLE5ExecutionStreamOperation.SetShapeHash]
//   - [MLE5ExecutionStreamOperation.State]
//   - [MLE5ExecutionStreamOperation.SetState]
//   - [MLE5ExecutionStreamOperation.StatePorts]
//   - [MLE5ExecutionStreamOperation.SetStatePorts]
//   - [MLE5ExecutionStreamOperation.WaitEventListener]
//   - [MLE5ExecutionStreamOperation.WaitSharedEventsBoundToESOP]
//   - [MLE5ExecutionStreamOperation.SetWaitSharedEventsBoundToESOP]
//   - [MLE5ExecutionStreamOperation.InitWithProgramLibraryFunctionNameModelDescriptionConfigurationDebugLabelModelSignpostId]
type MLE5ExecutionStreamOperation struct {
	objectivec.Object
}

// MLE5ExecutionStreamOperationFromID constructs a [MLE5ExecutionStreamOperation] from an objc.ID.
func MLE5ExecutionStreamOperationFromID(id objc.ID) MLE5ExecutionStreamOperation {
	return MLE5ExecutionStreamOperation{objectivec.Object{ID: id}}
}

// Ensure MLE5ExecutionStreamOperation implements IMLE5ExecutionStreamOperation.
var _ IMLE5ExecutionStreamOperation = MLE5ExecutionStreamOperation{}

// An interface definition for the [MLE5ExecutionStreamOperation] class.
//
// # Methods
//
//   - [IMLE5ExecutionStreamOperation._bindCompletionSyncPointDirectlyIfPossible]
//   - [IMLE5ExecutionStreamOperation._bindEventToWaitForCopyingInputFeaturesAfterSyncPoints]
//   - [IMLE5ExecutionStreamOperation._bindInputFeaturesAndWaitEventsOptionsError]
//   - [IMLE5ExecutionStreamOperation._bindNewCompletionEventsDirectlyWithCompletionSyncPoint]
//   - [IMLE5ExecutionStreamOperation._bindNewWaitEventsDirectlyWithWaitSyncPoints]
//   - [IMLE5ExecutionStreamOperation._bindOutputPortsWithOptionsError]
//   - [IMLE5ExecutionStreamOperation._bindWaitEventsDirectly]
//   - [IMLE5ExecutionStreamOperation._copyInputFeaturesError]
//   - [IMLE5ExecutionStreamOperation._createOperationAndReturnError]
//   - [IMLE5ExecutionStreamOperation._createOperationWithRetryCountError]
//   - [IMLE5ExecutionStreamOperation._directlyBoundFeatureNamesForPorts]
//   - [IMLE5ExecutionStreamOperation._inoutPortNames]
//   - [IMLE5ExecutionStreamOperation._inputPortNames]
//   - [IMLE5ExecutionStreamOperation._multiArrayFeatureFromStateFeature]
//   - [IMLE5ExecutionStreamOperation._newArrayOfInoutPortsFeatureDescriptionsByNameError]
//   - [IMLE5ExecutionStreamOperation._newArrayOfInputPortsFeatureDescriptionsByNameError]
//   - [IMLE5ExecutionStreamOperation._newArrayOfOutputPortsFeatureDescriptionsByNameError]
//   - [IMLE5ExecutionStreamOperation._outputPortNames]
//   - [IMLE5ExecutionStreamOperation._prepareInputPortsForFeaturesError]
//   - [IMLE5ExecutionStreamOperation._reusableForCompletionSyncPointAllOutputBackingsUseDirectBinding]
//   - [IMLE5ExecutionStreamOperation._reusableForWaitSyncPointsAllInputsUseDirectBinding]
//   - [IMLE5ExecutionStreamOperation._updateCompletionEventFutureValuesWithCompletionSyncPoint]
//   - [IMLE5ExecutionStreamOperation._updateWaitEventFutureValuesWithWaitSyncPoints]
//   - [IMLE5ExecutionStreamOperation.AsyncSubmissionError]
//   - [IMLE5ExecutionStreamOperation.SetAsyncSubmissionError]
//   - [IMLE5ExecutionStreamOperation.CompletionSharedEventBoundToESOP]
//   - [IMLE5ExecutionStreamOperation.SetCompletionSharedEventBoundToESOP]
//   - [IMLE5ExecutionStreamOperation.DebugLabel]
//   - [IMLE5ExecutionStreamOperation.DirectlyBoundInputFeatureNames]
//   - [IMLE5ExecutionStreamOperation.DirectlyBoundOutputFeatureNames]
//   - [IMLE5ExecutionStreamOperation.FunctionName]
//   - [IMLE5ExecutionStreamOperation.InputPorts]
//   - [IMLE5ExecutionStreamOperation.SetInputPorts]
//   - [IMLE5ExecutionStreamOperation.ModelConfiguration]
//   - [IMLE5ExecutionStreamOperation.ModelDescription]
//   - [IMLE5ExecutionStreamOperation.ModelSignpostId]
//   - [IMLE5ExecutionStreamOperation.OperationHandle]
//   - [IMLE5ExecutionStreamOperation.SetOperationHandle]
//   - [IMLE5ExecutionStreamOperation.OutputFeatures]
//   - [IMLE5ExecutionStreamOperation.OutputPorts]
//   - [IMLE5ExecutionStreamOperation.SetOutputPorts]
//   - [IMLE5ExecutionStreamOperation.PixelBufferPool]
//   - [IMLE5ExecutionStreamOperation.SetPixelBufferPool]
//   - [IMLE5ExecutionStreamOperation.PreloadAndReturnError]
//   - [IMLE5ExecutionStreamOperation.PrepareAsyncSubmissionForInputFeaturesOptionsError]
//   - [IMLE5ExecutionStreamOperation.PrepareForInputFeaturesOptionsError]
//   - [IMLE5ExecutionStreamOperation.ProgramLibrary]
//   - [IMLE5ExecutionStreamOperation.Reset]
//   - [IMLE5ExecutionStreamOperation.ReusableForInputFeaturesOptions]
//   - [IMLE5ExecutionStreamOperation.SerializeInferenceFrameDataForOptionsError]
//   - [IMLE5ExecutionStreamOperation.ShapeHash]
//   - [IMLE5ExecutionStreamOperation.SetShapeHash]
//   - [IMLE5ExecutionStreamOperation.State]
//   - [IMLE5ExecutionStreamOperation.SetState]
//   - [IMLE5ExecutionStreamOperation.StatePorts]
//   - [IMLE5ExecutionStreamOperation.SetStatePorts]
//   - [IMLE5ExecutionStreamOperation.WaitEventListener]
//   - [IMLE5ExecutionStreamOperation.WaitSharedEventsBoundToESOP]
//   - [IMLE5ExecutionStreamOperation.SetWaitSharedEventsBoundToESOP]
//   - [IMLE5ExecutionStreamOperation.InitWithProgramLibraryFunctionNameModelDescriptionConfigurationDebugLabelModelSignpostId]
type IMLE5ExecutionStreamOperation interface {
	objectivec.IObject

	// Topic: Methods

	_bindCompletionSyncPointDirectlyIfPossible(possible objectivec.IObject)
	_bindEventToWaitForCopyingInputFeaturesAfterSyncPoints(features objectivec.IObject, points objectivec.IObject)
	_bindInputFeaturesAndWaitEventsOptionsError(events objectivec.IObject, options objectivec.IObject) (bool, error)
	_bindNewCompletionEventsDirectlyWithCompletionSyncPoint(point objectivec.IObject)
	_bindNewWaitEventsDirectlyWithWaitSyncPoints(points objectivec.IObject)
	_bindOutputPortsWithOptionsError(options objectivec.IObject) (bool, error)
	_bindWaitEventsDirectly(directly objectivec.IObject)
	_copyInputFeaturesError(features objectivec.IObject) (bool, error)
	_createOperationAndReturnError() (E5rtExecutionStreamOperationRef, error)
	_createOperationWithRetryCountError(count int64) (E5rtExecutionStreamOperationRef, error)
	_directlyBoundFeatureNamesForPorts(ports objectivec.IObject) objectivec.IObject
	_inoutPortNames() objectivec.IObject
	_inputPortNames() objectivec.IObject
	_multiArrayFeatureFromStateFeature(feature objectivec.IObject) objectivec.IObject
	_newArrayOfInoutPortsFeatureDescriptionsByNameError(ports objectivec.IObject, name objectivec.IObject) (objectivec.IObject, error)
	_newArrayOfInputPortsFeatureDescriptionsByNameError(ports objectivec.IObject, name objectivec.IObject) (objectivec.IObject, error)
	_newArrayOfOutputPortsFeatureDescriptionsByNameError(ports objectivec.IObject, name objectivec.IObject) (objectivec.IObject, error)
	_outputPortNames() objectivec.IObject
	_prepareInputPortsForFeaturesError(features objectivec.IObject) (bool, error)
	_reusableForCompletionSyncPointAllOutputBackingsUseDirectBinding(point objectivec.IObject, binding bool) bool
	_reusableForWaitSyncPointsAllInputsUseDirectBinding(points objectivec.IObject, binding bool) bool
	_updateCompletionEventFutureValuesWithCompletionSyncPoint(point objectivec.IObject)
	_updateWaitEventFutureValuesWithWaitSyncPoints(points objectivec.IObject)
	AsyncSubmissionError() foundation.NSError
	SetAsyncSubmissionError(value foundation.NSError)
	CompletionSharedEventBoundToESOP() unsafe.Pointer
	SetCompletionSharedEventBoundToESOP(value unsafe.Pointer)
	DebugLabel() string
	DirectlyBoundInputFeatureNames() foundation.INSArray
	DirectlyBoundOutputFeatureNames() foundation.INSArray
	FunctionName() string
	InputPorts() foundation.INSArray
	SetInputPorts(value foundation.INSArray)
	ModelConfiguration() IMLModelConfiguration
	ModelDescription() IMLModelDescription
	ModelSignpostId() uint64
	OperationHandle() E5rtExecutionStreamOperationRef
	SetOperationHandle(value E5rtExecutionStreamOperationRef)
	OutputFeatures() unsafe.Pointer
	OutputPorts() foundation.INSArray
	SetOutputPorts(value foundation.INSArray)
	PixelBufferPool() IMLPixelBufferPool
	SetPixelBufferPool(value IMLPixelBufferPool)
	PreloadAndReturnError() (bool, error)
	PrepareAsyncSubmissionForInputFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (bool, error)
	PrepareForInputFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (bool, error)
	ProgramLibrary() IMLE5ProgramLibrary
	Reset()
	ReusableForInputFeaturesOptions(features objectivec.IObject, options objectivec.IObject) bool
	SerializeInferenceFrameDataForOptionsError(options objectivec.IObject) (bool, error)
	ShapeHash() string
	SetShapeHash(value string)
	State() int64
	SetState(value int64)
	StatePorts() foundation.INSArray
	SetStatePorts(value foundation.INSArray)
	WaitEventListener() unsafe.Pointer
	WaitSharedEventsBoundToESOP() foundation.INSArray
	SetWaitSharedEventsBoundToESOP(value foundation.INSArray)
	InitWithProgramLibraryFunctionNameModelDescriptionConfigurationDebugLabelModelSignpostId(library objectivec.IObject, name objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject, label objectivec.IObject, id uint64) MLE5ExecutionStreamOperation
}

// Init initializes the instance.
func (m MLE5ExecutionStreamOperation) Init() MLE5ExecutionStreamOperation {
	rv := objc.Send[MLE5ExecutionStreamOperation](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLE5ExecutionStreamOperation) Autorelease() MLE5ExecutionStreamOperation {
	rv := objc.Send[MLE5ExecutionStreamOperation](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLE5ExecutionStreamOperation creates a new MLE5ExecutionStreamOperation instance.
func NewMLE5ExecutionStreamOperation() MLE5ExecutionStreamOperation {
	class := getMLE5ExecutionStreamOperationClass()
	rv := objc.Send[MLE5ExecutionStreamOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewE5ExecutionStreamOperationWithProgramLibraryFunctionNameModelDescriptionConfigurationDebugLabelModelSignpostId(library objectivec.IObject, name objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject, label objectivec.IObject, id uint64) MLE5ExecutionStreamOperation {
	instance := getMLE5ExecutionStreamOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProgramLibrary:functionName:modelDescription:configuration:debugLabel:modelSignpostId:"), library, name, description, configuration, label, id)
	return MLE5ExecutionStreamOperationFromID(rv)
}

func (m MLE5ExecutionStreamOperation) _bindCompletionSyncPointDirectlyIfPossible(possible objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("_bindCompletionSyncPointDirectlyIfPossible:"), possible)
}

// BindCompletionSyncPointDirectlyIfPossible is an exported wrapper for the private method _bindCompletionSyncPointDirectlyIfPossible.
func (m MLE5ExecutionStreamOperation) BindCompletionSyncPointDirectlyIfPossible(possible objectivec.IObject) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_bindCompletionSyncPointDirectlyIfPossible:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_bindCompletionSyncPointDirectlyIfPossible:"}
		return err
	}
	m._bindCompletionSyncPointDirectlyIfPossible(possible)
	return nil
}

// CanBindCompletionSyncPointDirectlyIfPossible reports whether the receiver responds to the private selector _bindCompletionSyncPointDirectlyIfPossible:.
func (m MLE5ExecutionStreamOperation) CanBindCompletionSyncPointDirectlyIfPossible() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_bindCompletionSyncPointDirectlyIfPossible:"))
}
func (m MLE5ExecutionStreamOperation) _bindEventToWaitForCopyingInputFeaturesAfterSyncPoints(features objectivec.IObject, points objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("_bindEventToWaitForCopyingInputFeatures:afterSyncPoints:"), features, points)
}

// BindEventToWaitForCopyingInputFeaturesAfterSyncPoints is an exported wrapper for the private method _bindEventToWaitForCopyingInputFeaturesAfterSyncPoints.
func (m MLE5ExecutionStreamOperation) BindEventToWaitForCopyingInputFeaturesAfterSyncPoints(features objectivec.IObject, points objectivec.IObject) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_bindEventToWaitForCopyingInputFeatures:afterSyncPoints:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_bindEventToWaitForCopyingInputFeatures:afterSyncPoints:"}
		return err
	}
	m._bindEventToWaitForCopyingInputFeaturesAfterSyncPoints(features, points)
	return nil
}

// CanBindEventToWaitForCopyingInputFeaturesAfterSyncPoints reports whether the receiver responds to the private selector _bindEventToWaitForCopyingInputFeatures:afterSyncPoints:.
func (m MLE5ExecutionStreamOperation) CanBindEventToWaitForCopyingInputFeaturesAfterSyncPoints() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_bindEventToWaitForCopyingInputFeatures:afterSyncPoints:"))
}
func (m MLE5ExecutionStreamOperation) _bindInputFeaturesAndWaitEventsOptionsError(events objectivec.IObject, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_bindInputFeaturesAndWaitEvents:options:error:"), events, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_bindInputFeaturesAndWaitEvents:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// BindInputFeaturesAndWaitEventsOptionsError is an exported wrapper for the private method _bindInputFeaturesAndWaitEventsOptionsError.
func (m MLE5ExecutionStreamOperation) BindInputFeaturesAndWaitEventsOptionsError(events objectivec.IObject, options objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_bindInputFeaturesAndWaitEvents:options:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_bindInputFeaturesAndWaitEvents:options:error:"}
		return false, err
	}
	return m._bindInputFeaturesAndWaitEventsOptionsError(events, options)
}

// CanBindInputFeaturesAndWaitEventsOptionsError reports whether the receiver responds to the private selector _bindInputFeaturesAndWaitEvents:options:error:.
func (m MLE5ExecutionStreamOperation) CanBindInputFeaturesAndWaitEventsOptionsError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_bindInputFeaturesAndWaitEvents:options:error:"))
}
func (m MLE5ExecutionStreamOperation) _bindNewCompletionEventsDirectlyWithCompletionSyncPoint(point objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("_bindNewCompletionEventsDirectlyWithCompletionSyncPoint:"), point)
}

// BindNewCompletionEventsDirectlyWithCompletionSyncPoint is an exported wrapper for the private method _bindNewCompletionEventsDirectlyWithCompletionSyncPoint.
func (m MLE5ExecutionStreamOperation) BindNewCompletionEventsDirectlyWithCompletionSyncPoint(point objectivec.IObject) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_bindNewCompletionEventsDirectlyWithCompletionSyncPoint:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_bindNewCompletionEventsDirectlyWithCompletionSyncPoint:"}
		return err
	}
	m._bindNewCompletionEventsDirectlyWithCompletionSyncPoint(point)
	return nil
}

// CanBindNewCompletionEventsDirectlyWithCompletionSyncPoint reports whether the receiver responds to the private selector _bindNewCompletionEventsDirectlyWithCompletionSyncPoint:.
func (m MLE5ExecutionStreamOperation) CanBindNewCompletionEventsDirectlyWithCompletionSyncPoint() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_bindNewCompletionEventsDirectlyWithCompletionSyncPoint:"))
}
func (m MLE5ExecutionStreamOperation) _bindNewWaitEventsDirectlyWithWaitSyncPoints(points objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("_bindNewWaitEventsDirectlyWithWaitSyncPoints:"), points)
}

// BindNewWaitEventsDirectlyWithWaitSyncPoints is an exported wrapper for the private method _bindNewWaitEventsDirectlyWithWaitSyncPoints.
func (m MLE5ExecutionStreamOperation) BindNewWaitEventsDirectlyWithWaitSyncPoints(points objectivec.IObject) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_bindNewWaitEventsDirectlyWithWaitSyncPoints:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_bindNewWaitEventsDirectlyWithWaitSyncPoints:"}
		return err
	}
	m._bindNewWaitEventsDirectlyWithWaitSyncPoints(points)
	return nil
}

// CanBindNewWaitEventsDirectlyWithWaitSyncPoints reports whether the receiver responds to the private selector _bindNewWaitEventsDirectlyWithWaitSyncPoints:.
func (m MLE5ExecutionStreamOperation) CanBindNewWaitEventsDirectlyWithWaitSyncPoints() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_bindNewWaitEventsDirectlyWithWaitSyncPoints:"))
}
func (m MLE5ExecutionStreamOperation) _bindOutputPortsWithOptionsError(options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_bindOutputPortsWithOptions:error:"), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_bindOutputPortsWithOptions:error: returned NO with nil NSError")
	}
	return rv, nil

}

// BindOutputPortsWithOptionsError is an exported wrapper for the private method _bindOutputPortsWithOptionsError.
func (m MLE5ExecutionStreamOperation) BindOutputPortsWithOptionsError(options objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_bindOutputPortsWithOptions:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_bindOutputPortsWithOptions:error:"}
		return false, err
	}
	return m._bindOutputPortsWithOptionsError(options)
}

// CanBindOutputPortsWithOptionsError reports whether the receiver responds to the private selector _bindOutputPortsWithOptions:error:.
func (m MLE5ExecutionStreamOperation) CanBindOutputPortsWithOptionsError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_bindOutputPortsWithOptions:error:"))
}
func (m MLE5ExecutionStreamOperation) _bindWaitEventsDirectly(directly objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("_bindWaitEventsDirectly:"), directly)
}

// BindWaitEventsDirectly is an exported wrapper for the private method _bindWaitEventsDirectly.
func (m MLE5ExecutionStreamOperation) BindWaitEventsDirectly(directly objectivec.IObject) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_bindWaitEventsDirectly:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_bindWaitEventsDirectly:"}
		return err
	}
	m._bindWaitEventsDirectly(directly)
	return nil
}

// CanBindWaitEventsDirectly reports whether the receiver responds to the private selector _bindWaitEventsDirectly:.
func (m MLE5ExecutionStreamOperation) CanBindWaitEventsDirectly() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_bindWaitEventsDirectly:"))
}
func (m MLE5ExecutionStreamOperation) _copyInputFeaturesError(features objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_copyInputFeatures:error:"), features, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_copyInputFeatures:error: returned NO with nil NSError")
	}
	return rv, nil

}

// CopyInputFeaturesError is an exported wrapper for the private method _copyInputFeaturesError.
func (m MLE5ExecutionStreamOperation) CopyInputFeaturesError(features objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_copyInputFeatures:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_copyInputFeatures:error:"}
		return false, err
	}
	return m._copyInputFeaturesError(features)
}

// CanCopyInputFeaturesError reports whether the receiver responds to the private selector _copyInputFeatures:error:.
func (m MLE5ExecutionStreamOperation) CanCopyInputFeaturesError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_copyInputFeatures:error:"))
}
func (m MLE5ExecutionStreamOperation) _createOperationAndReturnError() (E5rtExecutionStreamOperationRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_createOperationAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(E5rtExecutionStreamOperationRef), foundation.NSErrorFrom(errorPtr)
	}
	return E5rtExecutionStreamOperationRef(rv), nil

}

// CreateOperationAndReturnError is an exported wrapper for the private method _createOperationAndReturnError.
func (m MLE5ExecutionStreamOperation) CreateOperationAndReturnError() (E5rtExecutionStreamOperationRef, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_createOperationAndReturnError:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_createOperationAndReturnError:"}
		return *new(E5rtExecutionStreamOperationRef), err
	}
	return m._createOperationAndReturnError()
}

// CanCreateOperationAndReturnError reports whether the receiver responds to the private selector _createOperationAndReturnError:.
func (m MLE5ExecutionStreamOperation) CanCreateOperationAndReturnError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_createOperationAndReturnError:"))
}
func (m MLE5ExecutionStreamOperation) _createOperationWithRetryCountError(count int64) (E5rtExecutionStreamOperationRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_createOperationWithRetryCount:error:"), count, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(E5rtExecutionStreamOperationRef), foundation.NSErrorFrom(errorPtr)
	}
	return E5rtExecutionStreamOperationRef(rv), nil

}

// CreateOperationWithRetryCountError is an exported wrapper for the private method _createOperationWithRetryCountError.
func (m MLE5ExecutionStreamOperation) CreateOperationWithRetryCountError(count int64) (E5rtExecutionStreamOperationRef, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_createOperationWithRetryCount:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_createOperationWithRetryCount:error:"}
		return *new(E5rtExecutionStreamOperationRef), err
	}
	return m._createOperationWithRetryCountError(count)
}

// CanCreateOperationWithRetryCountError reports whether the receiver responds to the private selector _createOperationWithRetryCount:error:.
func (m MLE5ExecutionStreamOperation) CanCreateOperationWithRetryCountError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_createOperationWithRetryCount:error:"))
}
func (m MLE5ExecutionStreamOperation) _directlyBoundFeatureNamesForPorts(ports objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_directlyBoundFeatureNamesForPorts:"), ports)
	return objectivec.Object{ID: rv}
}

// DirectlyBoundFeatureNamesForPorts is an exported wrapper for the private method _directlyBoundFeatureNamesForPorts.
func (m MLE5ExecutionStreamOperation) DirectlyBoundFeatureNamesForPorts(ports objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_directlyBoundFeatureNamesForPorts:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_directlyBoundFeatureNamesForPorts:"}
		return nil, err
	}
	return m._directlyBoundFeatureNamesForPorts(ports), nil
}

// CanDirectlyBoundFeatureNamesForPorts reports whether the receiver responds to the private selector _directlyBoundFeatureNamesForPorts:.
func (m MLE5ExecutionStreamOperation) CanDirectlyBoundFeatureNamesForPorts() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_directlyBoundFeatureNamesForPorts:"))
}
func (m MLE5ExecutionStreamOperation) _inoutPortNames() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_inoutPortNames"))
	return objectivec.Object{ID: rv}
}

// InoutPortNames is an exported wrapper for the private method _inoutPortNames.
func (m MLE5ExecutionStreamOperation) InoutPortNames() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_inoutPortNames")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_inoutPortNames"}
		return nil, err
	}
	return m._inoutPortNames(), nil
}

// CanInoutPortNames reports whether the receiver responds to the private selector _inoutPortNames.
func (m MLE5ExecutionStreamOperation) CanInoutPortNames() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_inoutPortNames"))
}
func (m MLE5ExecutionStreamOperation) _inputPortNames() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_inputPortNames"))
	return objectivec.Object{ID: rv}
}

// InputPortNames is an exported wrapper for the private method _inputPortNames.
func (m MLE5ExecutionStreamOperation) InputPortNames() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_inputPortNames")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_inputPortNames"}
		return nil, err
	}
	return m._inputPortNames(), nil
}

// CanInputPortNames reports whether the receiver responds to the private selector _inputPortNames.
func (m MLE5ExecutionStreamOperation) CanInputPortNames() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_inputPortNames"))
}
func (m MLE5ExecutionStreamOperation) _multiArrayFeatureFromStateFeature(feature objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_multiArrayFeatureFromStateFeature:"), feature)
	return objectivec.Object{ID: rv}
}

// MultiArrayFeatureFromStateFeature is an exported wrapper for the private method _multiArrayFeatureFromStateFeature.
func (m MLE5ExecutionStreamOperation) MultiArrayFeatureFromStateFeature(feature objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_multiArrayFeatureFromStateFeature:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_multiArrayFeatureFromStateFeature:"}
		return nil, err
	}
	return m._multiArrayFeatureFromStateFeature(feature), nil
}

// CanMultiArrayFeatureFromStateFeature reports whether the receiver responds to the private selector _multiArrayFeatureFromStateFeature:.
func (m MLE5ExecutionStreamOperation) CanMultiArrayFeatureFromStateFeature() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_multiArrayFeatureFromStateFeature:"))
}
func (m MLE5ExecutionStreamOperation) _newArrayOfInoutPortsFeatureDescriptionsByNameError(ports objectivec.IObject, name objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_newArrayOfInoutPorts:featureDescriptionsByName:error:"), ports, name, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// NewArrayOfInoutPortsFeatureDescriptionsByNameError is an exported wrapper for the private method _newArrayOfInoutPortsFeatureDescriptionsByNameError.
func (m MLE5ExecutionStreamOperation) NewArrayOfInoutPortsFeatureDescriptionsByNameError(ports objectivec.IObject, name objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_newArrayOfInoutPorts:featureDescriptionsByName:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_newArrayOfInoutPorts:featureDescriptionsByName:error:"}
		return nil, err
	}
	return m._newArrayOfInoutPortsFeatureDescriptionsByNameError(ports, name)
}

// CanNewArrayOfInoutPortsFeatureDescriptionsByNameError reports whether the receiver responds to the private selector _newArrayOfInoutPorts:featureDescriptionsByName:error:.
func (m MLE5ExecutionStreamOperation) CanNewArrayOfInoutPortsFeatureDescriptionsByNameError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_newArrayOfInoutPorts:featureDescriptionsByName:error:"))
}
func (m MLE5ExecutionStreamOperation) _newArrayOfInputPortsFeatureDescriptionsByNameError(ports objectivec.IObject, name objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_newArrayOfInputPorts:featureDescriptionsByName:error:"), ports, name, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// NewArrayOfInputPortsFeatureDescriptionsByNameError is an exported wrapper for the private method _newArrayOfInputPortsFeatureDescriptionsByNameError.
func (m MLE5ExecutionStreamOperation) NewArrayOfInputPortsFeatureDescriptionsByNameError(ports objectivec.IObject, name objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_newArrayOfInputPorts:featureDescriptionsByName:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_newArrayOfInputPorts:featureDescriptionsByName:error:"}
		return nil, err
	}
	return m._newArrayOfInputPortsFeatureDescriptionsByNameError(ports, name)
}

// CanNewArrayOfInputPortsFeatureDescriptionsByNameError reports whether the receiver responds to the private selector _newArrayOfInputPorts:featureDescriptionsByName:error:.
func (m MLE5ExecutionStreamOperation) CanNewArrayOfInputPortsFeatureDescriptionsByNameError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_newArrayOfInputPorts:featureDescriptionsByName:error:"))
}
func (m MLE5ExecutionStreamOperation) _newArrayOfOutputPortsFeatureDescriptionsByNameError(ports objectivec.IObject, name objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_newArrayOfOutputPorts:featureDescriptionsByName:error:"), ports, name, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// NewArrayOfOutputPortsFeatureDescriptionsByNameError is an exported wrapper for the private method _newArrayOfOutputPortsFeatureDescriptionsByNameError.
func (m MLE5ExecutionStreamOperation) NewArrayOfOutputPortsFeatureDescriptionsByNameError(ports objectivec.IObject, name objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_newArrayOfOutputPorts:featureDescriptionsByName:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_newArrayOfOutputPorts:featureDescriptionsByName:error:"}
		return nil, err
	}
	return m._newArrayOfOutputPortsFeatureDescriptionsByNameError(ports, name)
}

// CanNewArrayOfOutputPortsFeatureDescriptionsByNameError reports whether the receiver responds to the private selector _newArrayOfOutputPorts:featureDescriptionsByName:error:.
func (m MLE5ExecutionStreamOperation) CanNewArrayOfOutputPortsFeatureDescriptionsByNameError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_newArrayOfOutputPorts:featureDescriptionsByName:error:"))
}
func (m MLE5ExecutionStreamOperation) _outputPortNames() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_outputPortNames"))
	return objectivec.Object{ID: rv}
}

// OutputPortNames is an exported wrapper for the private method _outputPortNames.
func (m MLE5ExecutionStreamOperation) OutputPortNames() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_outputPortNames")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_outputPortNames"}
		return nil, err
	}
	return m._outputPortNames(), nil
}

// CanOutputPortNames reports whether the receiver responds to the private selector _outputPortNames.
func (m MLE5ExecutionStreamOperation) CanOutputPortNames() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_outputPortNames"))
}
func (m MLE5ExecutionStreamOperation) _prepareInputPortsForFeaturesError(features objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_prepareInputPortsForFeatures:error:"), features, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_prepareInputPortsForFeatures:error: returned NO with nil NSError")
	}
	return rv, nil

}

// PrepareInputPortsForFeaturesError is an exported wrapper for the private method _prepareInputPortsForFeaturesError.
func (m MLE5ExecutionStreamOperation) PrepareInputPortsForFeaturesError(features objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_prepareInputPortsForFeatures:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_prepareInputPortsForFeatures:error:"}
		return false, err
	}
	return m._prepareInputPortsForFeaturesError(features)
}

// CanPrepareInputPortsForFeaturesError reports whether the receiver responds to the private selector _prepareInputPortsForFeatures:error:.
func (m MLE5ExecutionStreamOperation) CanPrepareInputPortsForFeaturesError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_prepareInputPortsForFeatures:error:"))
}
func (m MLE5ExecutionStreamOperation) _reusableForCompletionSyncPointAllOutputBackingsUseDirectBinding(point objectivec.IObject, binding bool) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_reusableForCompletionSyncPoint:allOutputBackingsUseDirectBinding:"), point, binding)
	return rv
}

// ReusableForCompletionSyncPointAllOutputBackingsUseDirectBinding is an exported wrapper for the private method _reusableForCompletionSyncPointAllOutputBackingsUseDirectBinding.
func (m MLE5ExecutionStreamOperation) ReusableForCompletionSyncPointAllOutputBackingsUseDirectBinding(point objectivec.IObject, binding bool) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_reusableForCompletionSyncPoint:allOutputBackingsUseDirectBinding:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_reusableForCompletionSyncPoint:allOutputBackingsUseDirectBinding:"}
		return false, err
	}
	return m._reusableForCompletionSyncPointAllOutputBackingsUseDirectBinding(point, binding), nil
}

// CanReusableForCompletionSyncPointAllOutputBackingsUseDirectBinding reports whether the receiver responds to the private selector _reusableForCompletionSyncPoint:allOutputBackingsUseDirectBinding:.
func (m MLE5ExecutionStreamOperation) CanReusableForCompletionSyncPointAllOutputBackingsUseDirectBinding() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_reusableForCompletionSyncPoint:allOutputBackingsUseDirectBinding:"))
}
func (m MLE5ExecutionStreamOperation) _reusableForWaitSyncPointsAllInputsUseDirectBinding(points objectivec.IObject, binding bool) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_reusableForWaitSyncPoints:allInputsUseDirectBinding:"), points, binding)
	return rv
}

// ReusableForWaitSyncPointsAllInputsUseDirectBinding is an exported wrapper for the private method _reusableForWaitSyncPointsAllInputsUseDirectBinding.
func (m MLE5ExecutionStreamOperation) ReusableForWaitSyncPointsAllInputsUseDirectBinding(points objectivec.IObject, binding bool) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_reusableForWaitSyncPoints:allInputsUseDirectBinding:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_reusableForWaitSyncPoints:allInputsUseDirectBinding:"}
		return false, err
	}
	return m._reusableForWaitSyncPointsAllInputsUseDirectBinding(points, binding), nil
}

// CanReusableForWaitSyncPointsAllInputsUseDirectBinding reports whether the receiver responds to the private selector _reusableForWaitSyncPoints:allInputsUseDirectBinding:.
func (m MLE5ExecutionStreamOperation) CanReusableForWaitSyncPointsAllInputsUseDirectBinding() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_reusableForWaitSyncPoints:allInputsUseDirectBinding:"))
}
func (m MLE5ExecutionStreamOperation) _updateCompletionEventFutureValuesWithCompletionSyncPoint(point objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("_updateCompletionEventFutureValuesWithCompletionSyncPoint:"), point)
}

// UpdateCompletionEventFutureValuesWithCompletionSyncPoint is an exported wrapper for the private method _updateCompletionEventFutureValuesWithCompletionSyncPoint.
func (m MLE5ExecutionStreamOperation) UpdateCompletionEventFutureValuesWithCompletionSyncPoint(point objectivec.IObject) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_updateCompletionEventFutureValuesWithCompletionSyncPoint:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_updateCompletionEventFutureValuesWithCompletionSyncPoint:"}
		return err
	}
	m._updateCompletionEventFutureValuesWithCompletionSyncPoint(point)
	return nil
}

// CanUpdateCompletionEventFutureValuesWithCompletionSyncPoint reports whether the receiver responds to the private selector _updateCompletionEventFutureValuesWithCompletionSyncPoint:.
func (m MLE5ExecutionStreamOperation) CanUpdateCompletionEventFutureValuesWithCompletionSyncPoint() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_updateCompletionEventFutureValuesWithCompletionSyncPoint:"))
}
func (m MLE5ExecutionStreamOperation) _updateWaitEventFutureValuesWithWaitSyncPoints(points objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("_updateWaitEventFutureValuesWithWaitSyncPoints:"), points)
}

// UpdateWaitEventFutureValuesWithWaitSyncPoints is an exported wrapper for the private method _updateWaitEventFutureValuesWithWaitSyncPoints.
func (m MLE5ExecutionStreamOperation) UpdateWaitEventFutureValuesWithWaitSyncPoints(points objectivec.IObject) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_updateWaitEventFutureValuesWithWaitSyncPoints:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_updateWaitEventFutureValuesWithWaitSyncPoints:"}
		return err
	}
	m._updateWaitEventFutureValuesWithWaitSyncPoints(points)
	return nil
}

// CanUpdateWaitEventFutureValuesWithWaitSyncPoints reports whether the receiver responds to the private selector _updateWaitEventFutureValuesWithWaitSyncPoints:.
func (m MLE5ExecutionStreamOperation) CanUpdateWaitEventFutureValuesWithWaitSyncPoints() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_updateWaitEventFutureValuesWithWaitSyncPoints:"))
}
func (m MLE5ExecutionStreamOperation) PreloadAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("preloadAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("preloadAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLE5ExecutionStreamOperation) PrepareAsyncSubmissionForInputFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("prepareAsyncSubmissionForInputFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("prepareAsyncSubmissionForInputFeatures:options:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLE5ExecutionStreamOperation) PrepareForInputFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("prepareForInputFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("prepareForInputFeatures:options:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLE5ExecutionStreamOperation) Reset() {
	objc.Send[objc.ID](m.ID, objc.Sel("reset"))
}
func (m MLE5ExecutionStreamOperation) ReusableForInputFeaturesOptions(features objectivec.IObject, options objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("reusableForInputFeatures:options:"), features, options)
	return rv
}
func (m MLE5ExecutionStreamOperation) SerializeInferenceFrameDataForOptionsError(options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("serializeInferenceFrameDataForOptions:error:"), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("serializeInferenceFrameDataForOptions:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLE5ExecutionStreamOperation) InitWithProgramLibraryFunctionNameModelDescriptionConfigurationDebugLabelModelSignpostId(library objectivec.IObject, name objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject, label objectivec.IObject, id uint64) MLE5ExecutionStreamOperation {
	rv := objc.Send[MLE5ExecutionStreamOperation](m.ID, objc.Sel("initWithProgramLibrary:functionName:modelDescription:configuration:debugLabel:modelSignpostId:"), library, name, description, configuration, label, id)
	return rv
}

func (m MLE5ExecutionStreamOperation) AsyncSubmissionError() foundation.NSError {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("asyncSubmissionError"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
func (m MLE5ExecutionStreamOperation) SetAsyncSubmissionError(value foundation.NSError) {
	objc.Send[struct{}](m.ID, objc.Sel("setAsyncSubmissionError:"), value)
}
func (m MLE5ExecutionStreamOperation) CompletionSharedEventBoundToESOP() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("completionSharedEventBoundToESOP"))
	return rv
}
func (m MLE5ExecutionStreamOperation) SetCompletionSharedEventBoundToESOP(value unsafe.Pointer) {
	objc.Send[struct{}](m.ID, objc.Sel("setCompletionSharedEventBoundToESOP:"), value)
}
func (m MLE5ExecutionStreamOperation) DebugLabel() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugLabel"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLE5ExecutionStreamOperation) DirectlyBoundInputFeatureNames() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("directlyBoundInputFeatureNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLE5ExecutionStreamOperation) DirectlyBoundOutputFeatureNames() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("directlyBoundOutputFeatureNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLE5ExecutionStreamOperation) FunctionName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("functionName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLE5ExecutionStreamOperation) InputPorts() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputPorts"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLE5ExecutionStreamOperation) SetInputPorts(value foundation.INSArray) {
	objc.Send[struct{}](m.ID, objc.Sel("setInputPorts:"), value)
}
func (m MLE5ExecutionStreamOperation) ModelConfiguration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelConfiguration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}
func (m MLE5ExecutionStreamOperation) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}
func (m MLE5ExecutionStreamOperation) ModelSignpostId() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("modelSignpostId"))
	return rv
}
func (m MLE5ExecutionStreamOperation) OperationHandle() E5rtExecutionStreamOperationRef {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("operationHandle"))
	return E5rtExecutionStreamOperationRef(rv)
}
func (m MLE5ExecutionStreamOperation) SetOperationHandle(value E5rtExecutionStreamOperationRef) {
	objc.Send[struct{}](m.ID, objc.Sel("setOperationHandle:"), value)
}
func (m MLE5ExecutionStreamOperation) OutputFeatures() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("outputFeatures"))
	return rv
}
func (m MLE5ExecutionStreamOperation) OutputPorts() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputPorts"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLE5ExecutionStreamOperation) SetOutputPorts(value foundation.INSArray) {
	objc.Send[struct{}](m.ID, objc.Sel("setOutputPorts:"), value)
}
func (m MLE5ExecutionStreamOperation) PixelBufferPool() IMLPixelBufferPool {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("pixelBufferPool"))
	return MLPixelBufferPoolFromID(objc.ID(rv))
}
func (m MLE5ExecutionStreamOperation) SetPixelBufferPool(value IMLPixelBufferPool) {
	objc.Send[struct{}](m.ID, objc.Sel("setPixelBufferPool:"), value)
}
func (m MLE5ExecutionStreamOperation) ProgramLibrary() IMLE5ProgramLibrary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("programLibrary"))
	return MLE5ProgramLibraryFromID(objc.ID(rv))
}
func (m MLE5ExecutionStreamOperation) ShapeHash() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("shapeHash"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLE5ExecutionStreamOperation) SetShapeHash(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setShapeHash:"), objc.String(value))
}
func (m MLE5ExecutionStreamOperation) State() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("state"))
	return rv
}
func (m MLE5ExecutionStreamOperation) SetState(value int64) {
	objc.Send[struct{}](m.ID, objc.Sel("setState:"), value)
}
func (m MLE5ExecutionStreamOperation) StatePorts() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("statePorts"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLE5ExecutionStreamOperation) SetStatePorts(value foundation.INSArray) {
	objc.Send[struct{}](m.ID, objc.Sel("setStatePorts:"), value)
}
func (m MLE5ExecutionStreamOperation) WaitEventListener() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("waitEventListener"))
	return rv
}
func (m MLE5ExecutionStreamOperation) WaitSharedEventsBoundToESOP() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("waitSharedEventsBoundToESOP"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLE5ExecutionStreamOperation) SetWaitSharedEventsBoundToESOP(value foundation.INSArray) {
	objc.Send[struct{}](m.ID, objc.Sel("setWaitSharedEventsBoundToESOP:"), value)
}
