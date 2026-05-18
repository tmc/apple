// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLUpdateProgressHandlers] class.
var (
	_MLUpdateProgressHandlersClass     MLUpdateProgressHandlersClass
	_MLUpdateProgressHandlersClassOnce sync.Once
)

func getMLUpdateProgressHandlersClass() MLUpdateProgressHandlersClass {
	_MLUpdateProgressHandlersClassOnce.Do(func() {
		_MLUpdateProgressHandlersClass = MLUpdateProgressHandlersClass{class: objc.GetClass("MLUpdateProgressHandlers")}
	})
	return _MLUpdateProgressHandlersClass
}

// GetMLUpdateProgressHandlersClass returns the class object for MLUpdateProgressHandlers.
func GetMLUpdateProgressHandlersClass() MLUpdateProgressHandlersClass {
	return getMLUpdateProgressHandlersClass()
}

type MLUpdateProgressHandlersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLUpdateProgressHandlersClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLUpdateProgressHandlersClass) Alloc() MLUpdateProgressHandlers {
	rv := objc.Send[MLUpdateProgressHandlers](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLUpdateProgressHandlers._dispatchUpdateProgressHandlerForEventMetricsParametersErrorOnQueue]
//   - [MLUpdateProgressHandlers.DispatchEpochEndProgressHandlerWithMetricsParametersOnQueue]
//   - [MLUpdateProgressHandlers.DispatchMiniBatchEndProgressHandlerWithMetricsParametersOnQueue]
//   - [MLUpdateProgressHandlers.DispatchTrainingBeginProgressHandlerWithMetricsParametersOnQueue]
//   - [MLUpdateProgressHandlers.DispatchTrainingCompletionHandlerWithErrorOnQueue]
//   - [MLUpdateProgressHandlers.DispatchTrainingCompletionHandlerWithMetricsParametersOnQueue]
//   - [MLUpdateProgressHandlers.InterestedEvents]
//   - [MLUpdateProgressHandlers.SetInterestedEvents]
//   - [MLUpdateProgressHandlers.SetCompletionHandler]
//   - [MLUpdateProgressHandlers.SetProgressHandler]
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateProgressHandlers
type MLUpdateProgressHandlers struct {
	objectivec.Object
}

// MLUpdateProgressHandlersFromID constructs a [MLUpdateProgressHandlers] from an objc.ID.
func MLUpdateProgressHandlersFromID(id objc.ID) MLUpdateProgressHandlers {
	return MLUpdateProgressHandlers{objectivec.Object{ID: id}}
}

// Ensure MLUpdateProgressHandlers implements IMLUpdateProgressHandlers.
var _ IMLUpdateProgressHandlers = MLUpdateProgressHandlers{}

// An interface definition for the [MLUpdateProgressHandlers] class.
//
// # Methods
//
//   - [IMLUpdateProgressHandlers._dispatchUpdateProgressHandlerForEventMetricsParametersErrorOnQueue]
//   - [IMLUpdateProgressHandlers.DispatchEpochEndProgressHandlerWithMetricsParametersOnQueue]
//   - [IMLUpdateProgressHandlers.DispatchMiniBatchEndProgressHandlerWithMetricsParametersOnQueue]
//   - [IMLUpdateProgressHandlers.DispatchTrainingBeginProgressHandlerWithMetricsParametersOnQueue]
//   - [IMLUpdateProgressHandlers.DispatchTrainingCompletionHandlerWithErrorOnQueue]
//   - [IMLUpdateProgressHandlers.DispatchTrainingCompletionHandlerWithMetricsParametersOnQueue]
//   - [IMLUpdateProgressHandlers.InterestedEvents]
//   - [IMLUpdateProgressHandlers.SetInterestedEvents]
//   - [IMLUpdateProgressHandlers.SetCompletionHandler]
//   - [IMLUpdateProgressHandlers.SetProgressHandler]
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateProgressHandlers
type IMLUpdateProgressHandlers interface {
	objectivec.IObject

	// Topic: Methods

	_dispatchUpdateProgressHandlerForEventMetricsParametersErrorOnQueue(event int64, metrics objectivec.IObject, parameters objectivec.IObject, error_ objectivec.IObject, queue objectivec.IObject)
	DispatchEpochEndProgressHandlerWithMetricsParametersOnQueue(metrics objectivec.IObject, parameters objectivec.IObject, queue objectivec.IObject)
	DispatchMiniBatchEndProgressHandlerWithMetricsParametersOnQueue(metrics objectivec.IObject, parameters objectivec.IObject, queue objectivec.IObject)
	DispatchTrainingBeginProgressHandlerWithMetricsParametersOnQueue(metrics objectivec.IObject, parameters objectivec.IObject, queue objectivec.IObject)
	DispatchTrainingCompletionHandlerWithErrorOnQueue(error_ objectivec.IObject, queue objectivec.IObject)
	DispatchTrainingCompletionHandlerWithMetricsParametersOnQueue(metrics objectivec.IObject, parameters objectivec.IObject, queue objectivec.IObject)
	InterestedEvents() int64
	SetInterestedEvents(value int64)
	SetCompletionHandler(handler ErrorHandler)
	SetProgressHandler(handler VoidHandler)
}

// Init initializes the instance.
func (m MLUpdateProgressHandlers) Init() MLUpdateProgressHandlers {
	rv := objc.Send[MLUpdateProgressHandlers](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLUpdateProgressHandlers) Autorelease() MLUpdateProgressHandlers {
	rv := objc.Send[MLUpdateProgressHandlers](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLUpdateProgressHandlers creates a new MLUpdateProgressHandlers instance.
func NewMLUpdateProgressHandlers() MLUpdateProgressHandlers {
	class := getMLUpdateProgressHandlersClass()
	rv := objc.Send[MLUpdateProgressHandlers](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLUpdateProgressHandlers/_dispatchUpdateProgressHandlerForEvent:metrics:parameters:error:onQueue:
func (m MLUpdateProgressHandlers) _dispatchUpdateProgressHandlerForEventMetricsParametersErrorOnQueue(event int64, metrics objectivec.IObject, parameters objectivec.IObject, error_ objectivec.IObject, queue objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("_dispatchUpdateProgressHandlerForEvent:metrics:parameters:error:onQueue:"), event, metrics, parameters, error_, queue)
}

// DispatchUpdateProgressHandlerForEventMetricsParametersErrorOnQueue is an exported wrapper for the private method _dispatchUpdateProgressHandlerForEventMetricsParametersErrorOnQueue.
func (m MLUpdateProgressHandlers) DispatchUpdateProgressHandlerForEventMetricsParametersErrorOnQueue(event int64, metrics objectivec.IObject, parameters objectivec.IObject, error_ objectivec.IObject, queue objectivec.IObject) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_dispatchUpdateProgressHandlerForEvent:metrics:parameters:error:onQueue:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_dispatchUpdateProgressHandlerForEvent:metrics:parameters:error:onQueue:"}
		return err
	}
	m._dispatchUpdateProgressHandlerForEventMetricsParametersErrorOnQueue(event, metrics, parameters, error_, queue)
	return nil
}

// CanDispatchUpdateProgressHandlerForEventMetricsParametersErrorOnQueue reports whether the receiver responds to the private selector _dispatchUpdateProgressHandlerForEvent:metrics:parameters:error:onQueue:.
func (m MLUpdateProgressHandlers) CanDispatchUpdateProgressHandlerForEventMetricsParametersErrorOnQueue() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_dispatchUpdateProgressHandlerForEvent:metrics:parameters:error:onQueue:"))
}

// See: https://developer.apple.com/documentation/CoreML/MLUpdateProgressHandlers/dispatchEpochEndProgressHandlerWithMetrics:parameters:onQueue:
func (m MLUpdateProgressHandlers) DispatchEpochEndProgressHandlerWithMetricsParametersOnQueue(metrics objectivec.IObject, parameters objectivec.IObject, queue objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("dispatchEpochEndProgressHandlerWithMetrics:parameters:onQueue:"), metrics, parameters, queue)
}

// See: https://developer.apple.com/documentation/CoreML/MLUpdateProgressHandlers/dispatchMiniBatchEndProgressHandlerWithMetrics:parameters:onQueue:
func (m MLUpdateProgressHandlers) DispatchMiniBatchEndProgressHandlerWithMetricsParametersOnQueue(metrics objectivec.IObject, parameters objectivec.IObject, queue objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("dispatchMiniBatchEndProgressHandlerWithMetrics:parameters:onQueue:"), metrics, parameters, queue)
}

// See: https://developer.apple.com/documentation/CoreML/MLUpdateProgressHandlers/dispatchTrainingBeginProgressHandlerWithMetrics:parameters:onQueue:
func (m MLUpdateProgressHandlers) DispatchTrainingBeginProgressHandlerWithMetricsParametersOnQueue(metrics objectivec.IObject, parameters objectivec.IObject, queue objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("dispatchTrainingBeginProgressHandlerWithMetrics:parameters:onQueue:"), metrics, parameters, queue)
}

// See: https://developer.apple.com/documentation/CoreML/MLUpdateProgressHandlers/dispatchTrainingCompletionHandlerWithError:onQueue:
func (m MLUpdateProgressHandlers) DispatchTrainingCompletionHandlerWithErrorOnQueue(error_ objectivec.IObject, queue objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("dispatchTrainingCompletionHandlerWithError:onQueue:"), error_, queue)
}

// See: https://developer.apple.com/documentation/CoreML/MLUpdateProgressHandlers/dispatchTrainingCompletionHandlerWithMetrics:parameters:onQueue:
func (m MLUpdateProgressHandlers) DispatchTrainingCompletionHandlerWithMetricsParametersOnQueue(metrics objectivec.IObject, parameters objectivec.IObject, queue objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("dispatchTrainingCompletionHandlerWithMetrics:parameters:onQueue:"), metrics, parameters, queue)
}

// See: https://developer.apple.com/documentation/CoreML/MLUpdateProgressHandlers/setCompletionHandler:
func (m MLUpdateProgressHandlers) SetCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("setCompletionHandler:"), _block0)
}

// See: https://developer.apple.com/documentation/CoreML/MLUpdateProgressHandlers/setProgressHandler:
func (m MLUpdateProgressHandlers) SetProgressHandler(handler VoidHandler) {
	_block0, _ := NewVoidBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("setProgressHandler:"), _block0)
}

// See: https://developer.apple.com/documentation/CoreML/MLUpdateProgressHandlers/interestedEvents
func (m MLUpdateProgressHandlers) InterestedEvents() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("interestedEvents"))
	return rv
}
func (m MLUpdateProgressHandlers) SetInterestedEvents(value int64) {
	objc.Send[struct{}](m.ID, objc.Sel("setInterestedEvents:"), value)
}

// Set is a synchronous wrapper around [MLUpdateProgressHandlers.SetCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLUpdateProgressHandlers) Set(ctx context.Context) error {
	done := make(chan error, 1)
	m.SetCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetProgressHandlerSync is a synchronous wrapper around [MLUpdateProgressHandlers.SetProgressHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLUpdateProgressHandlers) SetProgressHandlerSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	m.SetProgressHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
