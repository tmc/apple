// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLE5ExecutionStream] class.
var (
	_MLE5ExecutionStreamClass     MLE5ExecutionStreamClass
	_MLE5ExecutionStreamClassOnce sync.Once
)

func getMLE5ExecutionStreamClass() MLE5ExecutionStreamClass {
	_MLE5ExecutionStreamClassOnce.Do(func() {
		_MLE5ExecutionStreamClass = MLE5ExecutionStreamClass{class: objc.GetClass("MLE5ExecutionStream")}
	})
	return _MLE5ExecutionStreamClass
}

// GetMLE5ExecutionStreamClass returns the class object for MLE5ExecutionStream.
func GetMLE5ExecutionStreamClass() MLE5ExecutionStreamClass {
	return getMLE5ExecutionStreamClass()
}

type MLE5ExecutionStreamClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLE5ExecutionStreamClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLE5ExecutionStreamClass) Alloc() MLE5ExecutionStream {
	rv := objc.Send[MLE5ExecutionStream](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLE5ExecutionStream._executeStreamError]
//   - [MLE5ExecutionStream._prepareForInputFeaturesOptionsError]
//   - [MLE5ExecutionStream._reset]
//   - [MLE5ExecutionStream._reusableForInputFeaturesOptions]
//   - [MLE5ExecutionStream._setANEExecutionPriorityWithOptions]
//   - [MLE5ExecutionStream.CancelPendingReset]
//   - [MLE5ExecutionStream.ExecuteForInputFeaturesOptionsError]
//   - [MLE5ExecutionStream.OperationPool]
//   - [MLE5ExecutionStream.SetOperationPool]
//   - [MLE5ExecutionStream.Operations]
//   - [MLE5ExecutionStream.SetOperations]
//   - [MLE5ExecutionStream.PrepareAsyncSubmissionForInputFeaturesOptionsError]
//   - [MLE5ExecutionStream.ResetAfterLingering]
//   - [MLE5ExecutionStream.ResetQueue]
//   - [MLE5ExecutionStream.ResetTimer]
//   - [MLE5ExecutionStream.SetResetTimer]
//   - [MLE5ExecutionStream.ReusableForAsyncSubmissionForInputFeaturesOptions]
//   - [MLE5ExecutionStream.ReusableForSyncPredictionForInputFeaturesOptions]
//   - [MLE5ExecutionStream.SerializeInferenceFrameDataForOptionsError]
//   - [MLE5ExecutionStream.SetupOperationForInputFeaturesOperationPoolError]
//   - [MLE5ExecutionStream.State]
//   - [MLE5ExecutionStream.SetState]
//   - [MLE5ExecutionStream.StreamHandle]
//   - [MLE5ExecutionStream.SetStreamHandle]
//   - [MLE5ExecutionStream.StreamId]
//   - [MLE5ExecutionStream.SubmitWithCompletionHandler]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream
type MLE5ExecutionStream struct {
	objectivec.Object
}

// MLE5ExecutionStreamFromID constructs a [MLE5ExecutionStream] from an objc.ID.
func MLE5ExecutionStreamFromID(id objc.ID) MLE5ExecutionStream {
	return MLE5ExecutionStream{objectivec.Object{ID: id}}
}

// Ensure MLE5ExecutionStream implements IMLE5ExecutionStream.
var _ IMLE5ExecutionStream = MLE5ExecutionStream{}

// An interface definition for the [MLE5ExecutionStream] class.
//
// # Methods
//
//   - [IMLE5ExecutionStream._executeStreamError]
//   - [IMLE5ExecutionStream._prepareForInputFeaturesOptionsError]
//   - [IMLE5ExecutionStream._reset]
//   - [IMLE5ExecutionStream._reusableForInputFeaturesOptions]
//   - [IMLE5ExecutionStream._setANEExecutionPriorityWithOptions]
//   - [IMLE5ExecutionStream.CancelPendingReset]
//   - [IMLE5ExecutionStream.ExecuteForInputFeaturesOptionsError]
//   - [IMLE5ExecutionStream.OperationPool]
//   - [IMLE5ExecutionStream.SetOperationPool]
//   - [IMLE5ExecutionStream.Operations]
//   - [IMLE5ExecutionStream.SetOperations]
//   - [IMLE5ExecutionStream.PrepareAsyncSubmissionForInputFeaturesOptionsError]
//   - [IMLE5ExecutionStream.ResetAfterLingering]
//   - [IMLE5ExecutionStream.ResetQueue]
//   - [IMLE5ExecutionStream.ResetTimer]
//   - [IMLE5ExecutionStream.SetResetTimer]
//   - [IMLE5ExecutionStream.ReusableForAsyncSubmissionForInputFeaturesOptions]
//   - [IMLE5ExecutionStream.ReusableForSyncPredictionForInputFeaturesOptions]
//   - [IMLE5ExecutionStream.SerializeInferenceFrameDataForOptionsError]
//   - [IMLE5ExecutionStream.SetupOperationForInputFeaturesOperationPoolError]
//   - [IMLE5ExecutionStream.State]
//   - [IMLE5ExecutionStream.SetState]
//   - [IMLE5ExecutionStream.StreamHandle]
//   - [IMLE5ExecutionStream.SetStreamHandle]
//   - [IMLE5ExecutionStream.StreamId]
//   - [IMLE5ExecutionStream.SubmitWithCompletionHandler]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream
type IMLE5ExecutionStream interface {
	objectivec.IObject

	// Topic: Methods

	_executeStreamError(stream E5rtExecutionStreamRef) (bool, error)
	_prepareForInputFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (bool, error)
	_reset()
	_reusableForInputFeaturesOptions(features objectivec.IObject, options objectivec.IObject) bool
	_setANEExecutionPriorityWithOptions(options objectivec.IObject)
	CancelPendingReset()
	ExecuteForInputFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (bool, error)
	OperationPool() objectivec.IObject
	SetOperationPool(value objectivec.IObject)
	Operations() foundation.INSArray
	SetOperations(value foundation.INSArray)
	PrepareAsyncSubmissionForInputFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (bool, error)
	ResetAfterLingering(lingering float64)
	ResetQueue() objectivec.Object
	ResetTimer() objectivec.Object
	SetResetTimer(value objectivec.Object)
	ReusableForAsyncSubmissionForInputFeaturesOptions(features objectivec.IObject, options objectivec.IObject) bool
	ReusableForSyncPredictionForInputFeaturesOptions(features objectivec.IObject, options objectivec.IObject) bool
	SerializeInferenceFrameDataForOptionsError(options objectivec.IObject) (bool, error)
	SetupOperationForInputFeaturesOperationPoolError(features objectivec.IObject, pool objectivec.IObject) (bool, error)
	State() int64
	SetState(value int64)
	StreamHandle() E5rtExecutionStreamRef
	SetStreamHandle(value E5rtExecutionStreamRef)
	StreamId() uint64
	SubmitWithCompletionHandler(handler ErrorHandler)
}

// Init initializes the instance.
func (m MLE5ExecutionStream) Init() MLE5ExecutionStream {
	rv := objc.Send[MLE5ExecutionStream](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLE5ExecutionStream) Autorelease() MLE5ExecutionStream {
	rv := objc.Send[MLE5ExecutionStream](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLE5ExecutionStream creates a new MLE5ExecutionStream instance.
func NewMLE5ExecutionStream() MLE5ExecutionStream {
	class := getMLE5ExecutionStreamClass()
	rv := objc.Send[MLE5ExecutionStream](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream/_executeStream:error:
func (m MLE5ExecutionStream) _executeStreamError(stream E5rtExecutionStreamRef) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_executeStream:error:"), stream, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_executeStream:error: returned NO with nil NSError")
	}
	return rv, nil

}

// ExecuteStreamError is an exported wrapper for the private method _executeStreamError.
func (m MLE5ExecutionStream) ExecuteStreamError(stream E5rtExecutionStreamRef) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_executeStream:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_executeStream:error:"}
		return false, err
	}
	return m._executeStreamError(stream)
}

// CanExecuteStreamError reports whether the receiver responds to the private selector _executeStream:error:.
func (m MLE5ExecutionStream) CanExecuteStreamError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_executeStream:error:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream/_prepareForInputFeatures:options:error:
func (m MLE5ExecutionStream) _prepareForInputFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_prepareForInputFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_prepareForInputFeatures:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// PrepareForInputFeaturesOptionsError is an exported wrapper for the private method _prepareForInputFeaturesOptionsError.
func (m MLE5ExecutionStream) PrepareForInputFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_prepareForInputFeatures:options:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_prepareForInputFeatures:options:error:"}
		return false, err
	}
	return m._prepareForInputFeaturesOptionsError(features, options)
}

// CanPrepareForInputFeaturesOptionsError reports whether the receiver responds to the private selector _prepareForInputFeatures:options:error:.
func (m MLE5ExecutionStream) CanPrepareForInputFeaturesOptionsError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_prepareForInputFeatures:options:error:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream/_reset
func (m MLE5ExecutionStream) _reset() {
	objc.Send[objc.ID](m.ID, objc.Sel("_reset"))
}

// Reset is an exported wrapper for the private method _reset.
func (m MLE5ExecutionStream) Reset() error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_reset")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_reset"}
		return err
	}
	m._reset()
	return nil
}

// CanReset reports whether the receiver responds to the private selector _reset.
func (m MLE5ExecutionStream) CanReset() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_reset"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream/_reusableForInputFeatures:options:
func (m MLE5ExecutionStream) _reusableForInputFeaturesOptions(features objectivec.IObject, options objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_reusableForInputFeatures:options:"), features, options)
	return rv
}

// ReusableForInputFeaturesOptions is an exported wrapper for the private method _reusableForInputFeaturesOptions.
func (m MLE5ExecutionStream) ReusableForInputFeaturesOptions(features objectivec.IObject, options objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_reusableForInputFeatures:options:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_reusableForInputFeatures:options:"}
		return false, err
	}
	return m._reusableForInputFeaturesOptions(features, options), nil
}

// CanReusableForInputFeaturesOptions reports whether the receiver responds to the private selector _reusableForInputFeatures:options:.
func (m MLE5ExecutionStream) CanReusableForInputFeaturesOptions() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_reusableForInputFeatures:options:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream/_setANEExecutionPriorityWithOptions:
func (m MLE5ExecutionStream) _setANEExecutionPriorityWithOptions(options objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setANEExecutionPriorityWithOptions:"), options)
}

// SetANEExecutionPriorityWithOptions is an exported wrapper for the private method _setANEExecutionPriorityWithOptions.
func (m MLE5ExecutionStream) SetANEExecutionPriorityWithOptions(options objectivec.IObject) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_setANEExecutionPriorityWithOptions:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setANEExecutionPriorityWithOptions:"}
		return err
	}
	m._setANEExecutionPriorityWithOptions(options)
	return nil
}

// CanSetANEExecutionPriorityWithOptions reports whether the receiver responds to the private selector _setANEExecutionPriorityWithOptions:.
func (m MLE5ExecutionStream) CanSetANEExecutionPriorityWithOptions() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_setANEExecutionPriorityWithOptions:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream/cancelPendingReset
func (m MLE5ExecutionStream) CancelPendingReset() {
	objc.Send[objc.ID](m.ID, objc.Sel("cancelPendingReset"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream/executeForInputFeatures:options:error:
func (m MLE5ExecutionStream) ExecuteForInputFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("executeForInputFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("executeForInputFeatures:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream/prepareAsyncSubmissionForInputFeatures:options:error:
func (m MLE5ExecutionStream) PrepareAsyncSubmissionForInputFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (bool, error) {
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

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream/resetAfterLingering:
func (m MLE5ExecutionStream) ResetAfterLingering(lingering float64) {
	objc.Send[objc.ID](m.ID, objc.Sel("resetAfterLingering:"), lingering)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream/reusableForAsyncSubmissionForInputFeatures:options:
func (m MLE5ExecutionStream) ReusableForAsyncSubmissionForInputFeaturesOptions(features objectivec.IObject, options objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("reusableForAsyncSubmissionForInputFeatures:options:"), features, options)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream/reusableForSyncPredictionForInputFeatures:options:
func (m MLE5ExecutionStream) ReusableForSyncPredictionForInputFeaturesOptions(features objectivec.IObject, options objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("reusableForSyncPredictionForInputFeatures:options:"), features, options)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream/serializeInferenceFrameDataForOptions:error:
func (m MLE5ExecutionStream) SerializeInferenceFrameDataForOptionsError(options objectivec.IObject) (bool, error) {
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

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream/setupOperationForInputFeatures:operationPool:error:
func (m MLE5ExecutionStream) SetupOperationForInputFeaturesOperationPoolError(features objectivec.IObject, pool objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("setupOperationForInputFeatures:operationPool:error:"), features, pool, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setupOperationForInputFeatures:operationPool:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream/submitWithCompletionHandler:
func (m MLE5ExecutionStream) SubmitWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("submitWithCompletionHandler:"), _block0)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream/operationPool
func (m MLE5ExecutionStream) OperationPool() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("operationPool"))
	return objectivec.Object{ID: rv}
}
func (m MLE5ExecutionStream) SetOperationPool(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setOperationPool:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream/operations
func (m MLE5ExecutionStream) Operations() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("operations"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLE5ExecutionStream) SetOperations(value foundation.INSArray) {
	objc.Send[struct{}](m.ID, objc.Sel("setOperations:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream/resetQueue
func (m MLE5ExecutionStream) ResetQueue() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("resetQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream/resetTimer
func (m MLE5ExecutionStream) ResetTimer() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("resetTimer"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (m MLE5ExecutionStream) SetResetTimer(value objectivec.Object) {
	objc.Send[struct{}](m.ID, objc.Sel("setResetTimer:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream/state
func (m MLE5ExecutionStream) State() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("state"))
	return rv
}
func (m MLE5ExecutionStream) SetState(value int64) {
	objc.Send[struct{}](m.ID, objc.Sel("setState:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream/streamHandle
func (m MLE5ExecutionStream) StreamHandle() E5rtExecutionStreamRef {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("streamHandle"))
	return E5rtExecutionStreamRef(rv)
}
func (m MLE5ExecutionStream) SetStreamHandle(value E5rtExecutionStreamRef) {
	objc.Send[struct{}](m.ID, objc.Sel("setStreamHandle:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStream/streamId
func (m MLE5ExecutionStream) StreamId() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("streamId"))
	return rv
}

// Submit is a synchronous wrapper around [MLE5ExecutionStream.SubmitWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLE5ExecutionStream) Submit(ctx context.Context) error {
	done := make(chan error, 1)
	m.SubmitWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
