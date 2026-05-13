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
//
// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation
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
//
// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation
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
	_createOperationAndReturnError() (E5rt_execution_stream_operationRef, error)
	_createOperationWithRetryCountError(count int64) (E5rt_execution_stream_operationRef, error)
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
	AsyncSubmissionError() foundation.INSError
	SetAsyncSubmissionError(value foundation.INSError)
	CompletionSharedEventBoundToESOP() objectivec.IObject
	SetCompletionSharedEventBoundToESOP(value objectivec.IObject)
	DebugLabel() string
	DirectlyBoundInputFeatureNames() foundation.INSArray
	DirectlyBoundOutputFeatureNames() foundation.INSArray
	FunctionName() string
	InputPorts() foundation.INSArray
	SetInputPorts(value foundation.INSArray)
	ModelConfiguration() IMLModelConfiguration
	ModelDescription() IMLModelDescription
	ModelSignpostId() uint64
	OperationHandle() E5rt_execution_stream_operationRef
	SetOperationHandle(value E5rt_execution_stream_operationRef)
	OutputFeatures() objectivec.IObject
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
func (e MLE5ExecutionStreamOperation) Init() MLE5ExecutionStreamOperation {
	rv := objc.Send[MLE5ExecutionStreamOperation](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e MLE5ExecutionStreamOperation) Autorelease() MLE5ExecutionStreamOperation {
	rv := objc.Send[MLE5ExecutionStreamOperation](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLE5ExecutionStreamOperation creates a new MLE5ExecutionStreamOperation instance.
func NewMLE5ExecutionStreamOperation() MLE5ExecutionStreamOperation {
	class := getMLE5ExecutionStreamOperationClass()
	rv := objc.Send[MLE5ExecutionStreamOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/initWithProgramLibrary:functionName:modelDescription:configuration:debugLabel:modelSignpostId:
func NewE5ExecutionStreamOperationWithProgramLibraryFunctionNameModelDescriptionConfigurationDebugLabelModelSignpostId(library objectivec.IObject, name objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject, label objectivec.IObject, id uint64) MLE5ExecutionStreamOperation {
	instance := getMLE5ExecutionStreamOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProgramLibrary:functionName:modelDescription:configuration:debugLabel:modelSignpostId:"), library, name, description, configuration, label, id)
	return MLE5ExecutionStreamOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_bindCompletionSyncPointDirectlyIfPossible:
func (e MLE5ExecutionStreamOperation) _bindCompletionSyncPointDirectlyIfPossible(possible objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("_bindCompletionSyncPointDirectlyIfPossible:"), possible)
}

// BindCompletionSyncPointDirectlyIfPossible is an exported wrapper for the private method _bindCompletionSyncPointDirectlyIfPossible.
func (e MLE5ExecutionStreamOperation) BindCompletionSyncPointDirectlyIfPossible(possible objectivec.IObject) error {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_bindCompletionSyncPointDirectlyIfPossible:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_bindCompletionSyncPointDirectlyIfPossible:"}
		return err
	}
	e._bindCompletionSyncPointDirectlyIfPossible(possible)
	return nil
}

// CanBindCompletionSyncPointDirectlyIfPossible reports whether the receiver responds to the private selector _bindCompletionSyncPointDirectlyIfPossible:.
func (e MLE5ExecutionStreamOperation) CanBindCompletionSyncPointDirectlyIfPossible() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_bindCompletionSyncPointDirectlyIfPossible:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_bindEventToWaitForCopyingInputFeatures:afterSyncPoints:
func (e MLE5ExecutionStreamOperation) _bindEventToWaitForCopyingInputFeaturesAfterSyncPoints(features objectivec.IObject, points objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("_bindEventToWaitForCopyingInputFeatures:afterSyncPoints:"), features, points)
}

// BindEventToWaitForCopyingInputFeaturesAfterSyncPoints is an exported wrapper for the private method _bindEventToWaitForCopyingInputFeaturesAfterSyncPoints.
func (e MLE5ExecutionStreamOperation) BindEventToWaitForCopyingInputFeaturesAfterSyncPoints(features objectivec.IObject, points objectivec.IObject) error {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_bindEventToWaitForCopyingInputFeatures:afterSyncPoints:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_bindEventToWaitForCopyingInputFeatures:afterSyncPoints:"}
		return err
	}
	e._bindEventToWaitForCopyingInputFeaturesAfterSyncPoints(features, points)
	return nil
}

// CanBindEventToWaitForCopyingInputFeaturesAfterSyncPoints reports whether the receiver responds to the private selector _bindEventToWaitForCopyingInputFeatures:afterSyncPoints:.
func (e MLE5ExecutionStreamOperation) CanBindEventToWaitForCopyingInputFeaturesAfterSyncPoints() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_bindEventToWaitForCopyingInputFeatures:afterSyncPoints:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_bindInputFeaturesAndWaitEvents:options:error:
func (e MLE5ExecutionStreamOperation) _bindInputFeaturesAndWaitEventsOptionsError(events objectivec.IObject, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("_bindInputFeaturesAndWaitEvents:options:error:"), events, options, unsafe.Pointer(&errorPtr))
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
func (e MLE5ExecutionStreamOperation) BindInputFeaturesAndWaitEventsOptionsError(events objectivec.IObject, options objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_bindInputFeaturesAndWaitEvents:options:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_bindInputFeaturesAndWaitEvents:options:error:"}
		return false, err
	}
	return e._bindInputFeaturesAndWaitEventsOptionsError(events, options)
}

// CanBindInputFeaturesAndWaitEventsOptionsError reports whether the receiver responds to the private selector _bindInputFeaturesAndWaitEvents:options:error:.
func (e MLE5ExecutionStreamOperation) CanBindInputFeaturesAndWaitEventsOptionsError() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_bindInputFeaturesAndWaitEvents:options:error:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_bindNewCompletionEventsDirectlyWithCompletionSyncPoint:
func (e MLE5ExecutionStreamOperation) _bindNewCompletionEventsDirectlyWithCompletionSyncPoint(point objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("_bindNewCompletionEventsDirectlyWithCompletionSyncPoint:"), point)
}

// BindNewCompletionEventsDirectlyWithCompletionSyncPoint is an exported wrapper for the private method _bindNewCompletionEventsDirectlyWithCompletionSyncPoint.
func (e MLE5ExecutionStreamOperation) BindNewCompletionEventsDirectlyWithCompletionSyncPoint(point objectivec.IObject) error {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_bindNewCompletionEventsDirectlyWithCompletionSyncPoint:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_bindNewCompletionEventsDirectlyWithCompletionSyncPoint:"}
		return err
	}
	e._bindNewCompletionEventsDirectlyWithCompletionSyncPoint(point)
	return nil
}

// CanBindNewCompletionEventsDirectlyWithCompletionSyncPoint reports whether the receiver responds to the private selector _bindNewCompletionEventsDirectlyWithCompletionSyncPoint:.
func (e MLE5ExecutionStreamOperation) CanBindNewCompletionEventsDirectlyWithCompletionSyncPoint() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_bindNewCompletionEventsDirectlyWithCompletionSyncPoint:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_bindNewWaitEventsDirectlyWithWaitSyncPoints:
func (e MLE5ExecutionStreamOperation) _bindNewWaitEventsDirectlyWithWaitSyncPoints(points objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("_bindNewWaitEventsDirectlyWithWaitSyncPoints:"), points)
}

// BindNewWaitEventsDirectlyWithWaitSyncPoints is an exported wrapper for the private method _bindNewWaitEventsDirectlyWithWaitSyncPoints.
func (e MLE5ExecutionStreamOperation) BindNewWaitEventsDirectlyWithWaitSyncPoints(points objectivec.IObject) error {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_bindNewWaitEventsDirectlyWithWaitSyncPoints:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_bindNewWaitEventsDirectlyWithWaitSyncPoints:"}
		return err
	}
	e._bindNewWaitEventsDirectlyWithWaitSyncPoints(points)
	return nil
}

// CanBindNewWaitEventsDirectlyWithWaitSyncPoints reports whether the receiver responds to the private selector _bindNewWaitEventsDirectlyWithWaitSyncPoints:.
func (e MLE5ExecutionStreamOperation) CanBindNewWaitEventsDirectlyWithWaitSyncPoints() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_bindNewWaitEventsDirectlyWithWaitSyncPoints:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_bindOutputPortsWithOptions:error:
func (e MLE5ExecutionStreamOperation) _bindOutputPortsWithOptionsError(options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("_bindOutputPortsWithOptions:error:"), options, unsafe.Pointer(&errorPtr))
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
func (e MLE5ExecutionStreamOperation) BindOutputPortsWithOptionsError(options objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_bindOutputPortsWithOptions:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_bindOutputPortsWithOptions:error:"}
		return false, err
	}
	return e._bindOutputPortsWithOptionsError(options)
}

// CanBindOutputPortsWithOptionsError reports whether the receiver responds to the private selector _bindOutputPortsWithOptions:error:.
func (e MLE5ExecutionStreamOperation) CanBindOutputPortsWithOptionsError() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_bindOutputPortsWithOptions:error:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_bindWaitEventsDirectly:
func (e MLE5ExecutionStreamOperation) _bindWaitEventsDirectly(directly objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("_bindWaitEventsDirectly:"), directly)
}

// BindWaitEventsDirectly is an exported wrapper for the private method _bindWaitEventsDirectly.
func (e MLE5ExecutionStreamOperation) BindWaitEventsDirectly(directly objectivec.IObject) error {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_bindWaitEventsDirectly:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_bindWaitEventsDirectly:"}
		return err
	}
	e._bindWaitEventsDirectly(directly)
	return nil
}

// CanBindWaitEventsDirectly reports whether the receiver responds to the private selector _bindWaitEventsDirectly:.
func (e MLE5ExecutionStreamOperation) CanBindWaitEventsDirectly() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_bindWaitEventsDirectly:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_copyInputFeatures:error:
func (e MLE5ExecutionStreamOperation) _copyInputFeaturesError(features objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("_copyInputFeatures:error:"), features, unsafe.Pointer(&errorPtr))
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
func (e MLE5ExecutionStreamOperation) CopyInputFeaturesError(features objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_copyInputFeatures:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_copyInputFeatures:error:"}
		return false, err
	}
	return e._copyInputFeaturesError(features)
}

// CanCopyInputFeaturesError reports whether the receiver responds to the private selector _copyInputFeatures:error:.
func (e MLE5ExecutionStreamOperation) CanCopyInputFeaturesError() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_copyInputFeatures:error:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_createOperationAndReturnError:
func (e MLE5ExecutionStreamOperation) _createOperationAndReturnError() (E5rt_execution_stream_operationRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[E5rt_execution_stream_operationRef](e.ID, objc.Sel("_createOperationAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// CreateOperationAndReturnError is an exported wrapper for the private method _createOperationAndReturnError.
func (e MLE5ExecutionStreamOperation) CreateOperationAndReturnError() (E5rt_execution_stream_operationRef, error) {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_createOperationAndReturnError:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_createOperationAndReturnError:"}
		return 0, err
	}
	return e._createOperationAndReturnError()
}

// CanCreateOperationAndReturnError reports whether the receiver responds to the private selector _createOperationAndReturnError:.
func (e MLE5ExecutionStreamOperation) CanCreateOperationAndReturnError() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_createOperationAndReturnError:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_createOperationWithRetryCount:error:
func (e MLE5ExecutionStreamOperation) _createOperationWithRetryCountError(count int64) (E5rt_execution_stream_operationRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[E5rt_execution_stream_operationRef](e.ID, objc.Sel("_createOperationWithRetryCount:error:"), count, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// CreateOperationWithRetryCountError is an exported wrapper for the private method _createOperationWithRetryCountError.
func (e MLE5ExecutionStreamOperation) CreateOperationWithRetryCountError(count int64) (E5rt_execution_stream_operationRef, error) {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_createOperationWithRetryCount:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_createOperationWithRetryCount:error:"}
		return 0, err
	}
	return e._createOperationWithRetryCountError(count)
}

// CanCreateOperationWithRetryCountError reports whether the receiver responds to the private selector _createOperationWithRetryCount:error:.
func (e MLE5ExecutionStreamOperation) CanCreateOperationWithRetryCountError() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_createOperationWithRetryCount:error:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_directlyBoundFeatureNamesForPorts:
func (e MLE5ExecutionStreamOperation) _directlyBoundFeatureNamesForPorts(ports objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("_directlyBoundFeatureNamesForPorts:"), ports)
	return objectivec.Object{ID: rv}
}

// DirectlyBoundFeatureNamesForPorts is an exported wrapper for the private method _directlyBoundFeatureNamesForPorts.
func (e MLE5ExecutionStreamOperation) DirectlyBoundFeatureNamesForPorts(ports objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_directlyBoundFeatureNamesForPorts:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_directlyBoundFeatureNamesForPorts:"}
		return nil, err
	}
	return e._directlyBoundFeatureNamesForPorts(ports), nil
}

// CanDirectlyBoundFeatureNamesForPorts reports whether the receiver responds to the private selector _directlyBoundFeatureNamesForPorts:.
func (e MLE5ExecutionStreamOperation) CanDirectlyBoundFeatureNamesForPorts() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_directlyBoundFeatureNamesForPorts:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_inoutPortNames
func (e MLE5ExecutionStreamOperation) _inoutPortNames() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("_inoutPortNames"))
	return objectivec.Object{ID: rv}
}

// InoutPortNames is an exported wrapper for the private method _inoutPortNames.
func (e MLE5ExecutionStreamOperation) InoutPortNames() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_inoutPortNames")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_inoutPortNames"}
		return nil, err
	}
	return e._inoutPortNames(), nil
}

// CanInoutPortNames reports whether the receiver responds to the private selector _inoutPortNames.
func (e MLE5ExecutionStreamOperation) CanInoutPortNames() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_inoutPortNames"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_inputPortNames
func (e MLE5ExecutionStreamOperation) _inputPortNames() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("_inputPortNames"))
	return objectivec.Object{ID: rv}
}

// InputPortNames is an exported wrapper for the private method _inputPortNames.
func (e MLE5ExecutionStreamOperation) InputPortNames() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_inputPortNames")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_inputPortNames"}
		return nil, err
	}
	return e._inputPortNames(), nil
}

// CanInputPortNames reports whether the receiver responds to the private selector _inputPortNames.
func (e MLE5ExecutionStreamOperation) CanInputPortNames() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_inputPortNames"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_multiArrayFeatureFromStateFeature:
func (e MLE5ExecutionStreamOperation) _multiArrayFeatureFromStateFeature(feature objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("_multiArrayFeatureFromStateFeature:"), feature)
	return objectivec.Object{ID: rv}
}

// MultiArrayFeatureFromStateFeature is an exported wrapper for the private method _multiArrayFeatureFromStateFeature.
func (e MLE5ExecutionStreamOperation) MultiArrayFeatureFromStateFeature(feature objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_multiArrayFeatureFromStateFeature:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_multiArrayFeatureFromStateFeature:"}
		return nil, err
	}
	return e._multiArrayFeatureFromStateFeature(feature), nil
}

// CanMultiArrayFeatureFromStateFeature reports whether the receiver responds to the private selector _multiArrayFeatureFromStateFeature:.
func (e MLE5ExecutionStreamOperation) CanMultiArrayFeatureFromStateFeature() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_multiArrayFeatureFromStateFeature:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_newArrayOfInoutPorts:featureDescriptionsByName:error:
func (e MLE5ExecutionStreamOperation) _newArrayOfInoutPortsFeatureDescriptionsByNameError(ports objectivec.IObject, name objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("_newArrayOfInoutPorts:featureDescriptionsByName:error:"), ports, name, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// NewArrayOfInoutPortsFeatureDescriptionsByNameError is an exported wrapper for the private method _newArrayOfInoutPortsFeatureDescriptionsByNameError.
func (e MLE5ExecutionStreamOperation) NewArrayOfInoutPortsFeatureDescriptionsByNameError(ports objectivec.IObject, name objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_newArrayOfInoutPorts:featureDescriptionsByName:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_newArrayOfInoutPorts:featureDescriptionsByName:error:"}
		return nil, err
	}
	return e._newArrayOfInoutPortsFeatureDescriptionsByNameError(ports, name)
}

// CanNewArrayOfInoutPortsFeatureDescriptionsByNameError reports whether the receiver responds to the private selector _newArrayOfInoutPorts:featureDescriptionsByName:error:.
func (e MLE5ExecutionStreamOperation) CanNewArrayOfInoutPortsFeatureDescriptionsByNameError() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_newArrayOfInoutPorts:featureDescriptionsByName:error:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_newArrayOfInputPorts:featureDescriptionsByName:error:
func (e MLE5ExecutionStreamOperation) _newArrayOfInputPortsFeatureDescriptionsByNameError(ports objectivec.IObject, name objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("_newArrayOfInputPorts:featureDescriptionsByName:error:"), ports, name, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// NewArrayOfInputPortsFeatureDescriptionsByNameError is an exported wrapper for the private method _newArrayOfInputPortsFeatureDescriptionsByNameError.
func (e MLE5ExecutionStreamOperation) NewArrayOfInputPortsFeatureDescriptionsByNameError(ports objectivec.IObject, name objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_newArrayOfInputPorts:featureDescriptionsByName:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_newArrayOfInputPorts:featureDescriptionsByName:error:"}
		return nil, err
	}
	return e._newArrayOfInputPortsFeatureDescriptionsByNameError(ports, name)
}

// CanNewArrayOfInputPortsFeatureDescriptionsByNameError reports whether the receiver responds to the private selector _newArrayOfInputPorts:featureDescriptionsByName:error:.
func (e MLE5ExecutionStreamOperation) CanNewArrayOfInputPortsFeatureDescriptionsByNameError() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_newArrayOfInputPorts:featureDescriptionsByName:error:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_newArrayOfOutputPorts:featureDescriptionsByName:error:
func (e MLE5ExecutionStreamOperation) _newArrayOfOutputPortsFeatureDescriptionsByNameError(ports objectivec.IObject, name objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("_newArrayOfOutputPorts:featureDescriptionsByName:error:"), ports, name, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// NewArrayOfOutputPortsFeatureDescriptionsByNameError is an exported wrapper for the private method _newArrayOfOutputPortsFeatureDescriptionsByNameError.
func (e MLE5ExecutionStreamOperation) NewArrayOfOutputPortsFeatureDescriptionsByNameError(ports objectivec.IObject, name objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_newArrayOfOutputPorts:featureDescriptionsByName:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_newArrayOfOutputPorts:featureDescriptionsByName:error:"}
		return nil, err
	}
	return e._newArrayOfOutputPortsFeatureDescriptionsByNameError(ports, name)
}

// CanNewArrayOfOutputPortsFeatureDescriptionsByNameError reports whether the receiver responds to the private selector _newArrayOfOutputPorts:featureDescriptionsByName:error:.
func (e MLE5ExecutionStreamOperation) CanNewArrayOfOutputPortsFeatureDescriptionsByNameError() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_newArrayOfOutputPorts:featureDescriptionsByName:error:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_outputPortNames
func (e MLE5ExecutionStreamOperation) _outputPortNames() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("_outputPortNames"))
	return objectivec.Object{ID: rv}
}

// OutputPortNames is an exported wrapper for the private method _outputPortNames.
func (e MLE5ExecutionStreamOperation) OutputPortNames() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_outputPortNames")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_outputPortNames"}
		return nil, err
	}
	return e._outputPortNames(), nil
}

// CanOutputPortNames reports whether the receiver responds to the private selector _outputPortNames.
func (e MLE5ExecutionStreamOperation) CanOutputPortNames() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_outputPortNames"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_prepareInputPortsForFeatures:error:
func (e MLE5ExecutionStreamOperation) _prepareInputPortsForFeaturesError(features objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("_prepareInputPortsForFeatures:error:"), features, unsafe.Pointer(&errorPtr))
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
func (e MLE5ExecutionStreamOperation) PrepareInputPortsForFeaturesError(features objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_prepareInputPortsForFeatures:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_prepareInputPortsForFeatures:error:"}
		return false, err
	}
	return e._prepareInputPortsForFeaturesError(features)
}

// CanPrepareInputPortsForFeaturesError reports whether the receiver responds to the private selector _prepareInputPortsForFeatures:error:.
func (e MLE5ExecutionStreamOperation) CanPrepareInputPortsForFeaturesError() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_prepareInputPortsForFeatures:error:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_reusableForCompletionSyncPoint:allOutputBackingsUseDirectBinding:
func (e MLE5ExecutionStreamOperation) _reusableForCompletionSyncPointAllOutputBackingsUseDirectBinding(point objectivec.IObject, binding bool) bool {
	rv := objc.Send[bool](e.ID, objc.Sel("_reusableForCompletionSyncPoint:allOutputBackingsUseDirectBinding:"), point, binding)
	return rv
}

// ReusableForCompletionSyncPointAllOutputBackingsUseDirectBinding is an exported wrapper for the private method _reusableForCompletionSyncPointAllOutputBackingsUseDirectBinding.
func (e MLE5ExecutionStreamOperation) ReusableForCompletionSyncPointAllOutputBackingsUseDirectBinding(point objectivec.IObject, binding bool) (bool, error) {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_reusableForCompletionSyncPoint:allOutputBackingsUseDirectBinding:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_reusableForCompletionSyncPoint:allOutputBackingsUseDirectBinding:"}
		return false, err
	}
	return e._reusableForCompletionSyncPointAllOutputBackingsUseDirectBinding(point, binding), nil
}

// CanReusableForCompletionSyncPointAllOutputBackingsUseDirectBinding reports whether the receiver responds to the private selector _reusableForCompletionSyncPoint:allOutputBackingsUseDirectBinding:.
func (e MLE5ExecutionStreamOperation) CanReusableForCompletionSyncPointAllOutputBackingsUseDirectBinding() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_reusableForCompletionSyncPoint:allOutputBackingsUseDirectBinding:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_reusableForWaitSyncPoints:allInputsUseDirectBinding:
func (e MLE5ExecutionStreamOperation) _reusableForWaitSyncPointsAllInputsUseDirectBinding(points objectivec.IObject, binding bool) bool {
	rv := objc.Send[bool](e.ID, objc.Sel("_reusableForWaitSyncPoints:allInputsUseDirectBinding:"), points, binding)
	return rv
}

// ReusableForWaitSyncPointsAllInputsUseDirectBinding is an exported wrapper for the private method _reusableForWaitSyncPointsAllInputsUseDirectBinding.
func (e MLE5ExecutionStreamOperation) ReusableForWaitSyncPointsAllInputsUseDirectBinding(points objectivec.IObject, binding bool) (bool, error) {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_reusableForWaitSyncPoints:allInputsUseDirectBinding:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_reusableForWaitSyncPoints:allInputsUseDirectBinding:"}
		return false, err
	}
	return e._reusableForWaitSyncPointsAllInputsUseDirectBinding(points, binding), nil
}

// CanReusableForWaitSyncPointsAllInputsUseDirectBinding reports whether the receiver responds to the private selector _reusableForWaitSyncPoints:allInputsUseDirectBinding:.
func (e MLE5ExecutionStreamOperation) CanReusableForWaitSyncPointsAllInputsUseDirectBinding() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_reusableForWaitSyncPoints:allInputsUseDirectBinding:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_updateCompletionEventFutureValuesWithCompletionSyncPoint:
func (e MLE5ExecutionStreamOperation) _updateCompletionEventFutureValuesWithCompletionSyncPoint(point objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("_updateCompletionEventFutureValuesWithCompletionSyncPoint:"), point)
}

// UpdateCompletionEventFutureValuesWithCompletionSyncPoint is an exported wrapper for the private method _updateCompletionEventFutureValuesWithCompletionSyncPoint.
func (e MLE5ExecutionStreamOperation) UpdateCompletionEventFutureValuesWithCompletionSyncPoint(point objectivec.IObject) error {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_updateCompletionEventFutureValuesWithCompletionSyncPoint:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_updateCompletionEventFutureValuesWithCompletionSyncPoint:"}
		return err
	}
	e._updateCompletionEventFutureValuesWithCompletionSyncPoint(point)
	return nil
}

// CanUpdateCompletionEventFutureValuesWithCompletionSyncPoint reports whether the receiver responds to the private selector _updateCompletionEventFutureValuesWithCompletionSyncPoint:.
func (e MLE5ExecutionStreamOperation) CanUpdateCompletionEventFutureValuesWithCompletionSyncPoint() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_updateCompletionEventFutureValuesWithCompletionSyncPoint:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/_updateWaitEventFutureValuesWithWaitSyncPoints:
func (e MLE5ExecutionStreamOperation) _updateWaitEventFutureValuesWithWaitSyncPoints(points objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("_updateWaitEventFutureValuesWithWaitSyncPoints:"), points)
}

// UpdateWaitEventFutureValuesWithWaitSyncPoints is an exported wrapper for the private method _updateWaitEventFutureValuesWithWaitSyncPoints.
func (e MLE5ExecutionStreamOperation) UpdateWaitEventFutureValuesWithWaitSyncPoints(points objectivec.IObject) error {
	if !objc.RespondsToSelector(e.ID, objc.Sel("_updateWaitEventFutureValuesWithWaitSyncPoints:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_updateWaitEventFutureValuesWithWaitSyncPoints:"}
		return err
	}
	e._updateWaitEventFutureValuesWithWaitSyncPoints(points)
	return nil
}

// CanUpdateWaitEventFutureValuesWithWaitSyncPoints reports whether the receiver responds to the private selector _updateWaitEventFutureValuesWithWaitSyncPoints:.
func (e MLE5ExecutionStreamOperation) CanUpdateWaitEventFutureValuesWithWaitSyncPoints() bool {
	return objc.RespondsToSelector(e.ID, objc.Sel("_updateWaitEventFutureValuesWithWaitSyncPoints:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/preloadAndReturnError:
func (e MLE5ExecutionStreamOperation) PreloadAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("preloadAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("preloadAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/prepareAsyncSubmissionForInputFeatures:options:error:
func (e MLE5ExecutionStreamOperation) PrepareAsyncSubmissionForInputFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("prepareAsyncSubmissionForInputFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("prepareAsyncSubmissionForInputFeatures:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/prepareForInputFeatures:options:error:
func (e MLE5ExecutionStreamOperation) PrepareForInputFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("prepareForInputFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("prepareForInputFeatures:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/reset
func (e MLE5ExecutionStreamOperation) Reset() {
	objc.Send[objc.ID](e.ID, objc.Sel("reset"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/reusableForInputFeatures:options:
func (e MLE5ExecutionStreamOperation) ReusableForInputFeaturesOptions(features objectivec.IObject, options objectivec.IObject) bool {
	rv := objc.Send[bool](e.ID, objc.Sel("reusableForInputFeatures:options:"), features, options)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/serializeInferenceFrameDataForOptions:error:
func (e MLE5ExecutionStreamOperation) SerializeInferenceFrameDataForOptionsError(options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("serializeInferenceFrameDataForOptions:error:"), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("serializeInferenceFrameDataForOptions:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/initWithProgramLibrary:functionName:modelDescription:configuration:debugLabel:modelSignpostId:
func (e MLE5ExecutionStreamOperation) InitWithProgramLibraryFunctionNameModelDescriptionConfigurationDebugLabelModelSignpostId(library objectivec.IObject, name objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject, label objectivec.IObject, id uint64) MLE5ExecutionStreamOperation {
	rv := objc.Send[MLE5ExecutionStreamOperation](e.ID, objc.Sel("initWithProgramLibrary:functionName:modelDescription:configuration:debugLabel:modelSignpostId:"), library, name, description, configuration, label, id)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/asyncSubmissionError
func (e MLE5ExecutionStreamOperation) AsyncSubmissionError() foundation.INSError {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("asyncSubmissionError"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
func (e MLE5ExecutionStreamOperation) SetAsyncSubmissionError(value foundation.INSError) {
	objc.Send[struct{}](e.ID, objc.Sel("setAsyncSubmissionError:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/completionSharedEventBoundToESOP
func (e MLE5ExecutionStreamOperation) CompletionSharedEventBoundToESOP() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("completionSharedEventBoundToESOP"))
	return objectivec.Object{ID: rv}
}
func (e MLE5ExecutionStreamOperation) SetCompletionSharedEventBoundToESOP(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setCompletionSharedEventBoundToESOP:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/debugLabel
func (e MLE5ExecutionStreamOperation) DebugLabel() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("debugLabel"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/directlyBoundInputFeatureNames
func (e MLE5ExecutionStreamOperation) DirectlyBoundInputFeatureNames() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("directlyBoundInputFeatureNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/directlyBoundOutputFeatureNames
func (e MLE5ExecutionStreamOperation) DirectlyBoundOutputFeatureNames() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("directlyBoundOutputFeatureNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/functionName
func (e MLE5ExecutionStreamOperation) FunctionName() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("functionName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/inputPorts
func (e MLE5ExecutionStreamOperation) InputPorts() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("inputPorts"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e MLE5ExecutionStreamOperation) SetInputPorts(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setInputPorts:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/modelConfiguration
func (e MLE5ExecutionStreamOperation) ModelConfiguration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("modelConfiguration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/modelDescription
func (e MLE5ExecutionStreamOperation) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/modelSignpostId
func (e MLE5ExecutionStreamOperation) ModelSignpostId() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("modelSignpostId"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/operationHandle
func (e MLE5ExecutionStreamOperation) OperationHandle() E5rt_execution_stream_operationRef {
	rv := objc.Send[E5rt_execution_stream_operationRef](e.ID, objc.Sel("operationHandle"))
	return E5rt_execution_stream_operationRef(rv)
}
func (e MLE5ExecutionStreamOperation) SetOperationHandle(value E5rt_execution_stream_operationRef) {
	objc.Send[struct{}](e.ID, objc.Sel("setOperationHandle:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/outputFeatures
func (e MLE5ExecutionStreamOperation) OutputFeatures() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("outputFeatures"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/outputPorts
func (e MLE5ExecutionStreamOperation) OutputPorts() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("outputPorts"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e MLE5ExecutionStreamOperation) SetOutputPorts(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setOutputPorts:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/pixelBufferPool
func (e MLE5ExecutionStreamOperation) PixelBufferPool() IMLPixelBufferPool {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("pixelBufferPool"))
	return MLPixelBufferPoolFromID(objc.ID(rv))
}
func (e MLE5ExecutionStreamOperation) SetPixelBufferPool(value IMLPixelBufferPool) {
	objc.Send[struct{}](e.ID, objc.Sel("setPixelBufferPool:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/programLibrary
func (e MLE5ExecutionStreamOperation) ProgramLibrary() IMLE5ProgramLibrary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("programLibrary"))
	return MLE5ProgramLibraryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/shapeHash
func (e MLE5ExecutionStreamOperation) ShapeHash() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("shapeHash"))
	return foundation.NSStringFromID(rv).String()
}
func (e MLE5ExecutionStreamOperation) SetShapeHash(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setShapeHash:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/state
func (e MLE5ExecutionStreamOperation) State() int64 {
	rv := objc.Send[int64](e.ID, objc.Sel("state"))
	return rv
}
func (e MLE5ExecutionStreamOperation) SetState(value int64) {
	objc.Send[struct{}](e.ID, objc.Sel("setState:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/statePorts
func (e MLE5ExecutionStreamOperation) StatePorts() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("statePorts"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e MLE5ExecutionStreamOperation) SetStatePorts(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setStatePorts:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/waitEventListener
func (e MLE5ExecutionStreamOperation) WaitEventListener() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("waitEventListener"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamOperation/waitSharedEventsBoundToESOP
func (e MLE5ExecutionStreamOperation) WaitSharedEventsBoundToESOP() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("waitSharedEventsBoundToESOP"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e MLE5ExecutionStreamOperation) SetWaitSharedEventsBoundToESOP(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setWaitSharedEventsBoundToESOP:"), value)
}
