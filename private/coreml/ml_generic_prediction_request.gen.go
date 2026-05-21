// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLGenericPredictionRequest] class.
var (
	_MLGenericPredictionRequestClass     MLGenericPredictionRequestClass
	_MLGenericPredictionRequestClassOnce sync.Once
)

func getMLGenericPredictionRequestClass() MLGenericPredictionRequestClass {
	_MLGenericPredictionRequestClassOnce.Do(func() {
		_MLGenericPredictionRequestClass = MLGenericPredictionRequestClass{class: objc.GetClass("MLGenericPredictionRequest")}
	})
	return _MLGenericPredictionRequestClass
}

// GetMLGenericPredictionRequestClass returns the class object for MLGenericPredictionRequest.
func GetMLGenericPredictionRequestClass() MLGenericPredictionRequestClass {
	return getMLGenericPredictionRequestClass()
}

type MLGenericPredictionRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLGenericPredictionRequestClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLGenericPredictionRequestClass) Alloc() MLGenericPredictionRequest {
	rv := objc.Send[MLGenericPredictionRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLGenericPredictionRequest.Cancel]
//   - [MLGenericPredictionRequest.InputFeatures]
//   - [MLGenericPredictionRequest.IsCancelled]
//   - [MLGenericPredictionRequest.Model]
//   - [MLGenericPredictionRequest.PredictionOptions]
//   - [MLGenericPredictionRequest.SubmitWithCompletionHandler]
//   - [MLGenericPredictionRequest.InitForModelInputFeaturesOptions]
//   - [MLGenericPredictionRequest.DebugDescription]
//   - [MLGenericPredictionRequest.Description]
//   - [MLGenericPredictionRequest.Hash]
//   - [MLGenericPredictionRequest.Superclass]
type MLGenericPredictionRequest struct {
	objectivec.Object
}

// MLGenericPredictionRequestFromID constructs a [MLGenericPredictionRequest] from an objc.ID.
func MLGenericPredictionRequestFromID(id objc.ID) MLGenericPredictionRequest {
	return MLGenericPredictionRequest{objectivec.Object{ID: id}}
}

// Ensure MLGenericPredictionRequest implements IMLGenericPredictionRequest.
var _ IMLGenericPredictionRequest = MLGenericPredictionRequest{}

// An interface definition for the [MLGenericPredictionRequest] class.
//
// # Methods
//
//   - [IMLGenericPredictionRequest.Cancel]
//   - [IMLGenericPredictionRequest.InputFeatures]
//   - [IMLGenericPredictionRequest.IsCancelled]
//   - [IMLGenericPredictionRequest.Model]
//   - [IMLGenericPredictionRequest.PredictionOptions]
//   - [IMLGenericPredictionRequest.SubmitWithCompletionHandler]
//   - [IMLGenericPredictionRequest.InitForModelInputFeaturesOptions]
//   - [IMLGenericPredictionRequest.DebugDescription]
//   - [IMLGenericPredictionRequest.Description]
//   - [IMLGenericPredictionRequest.Hash]
//   - [IMLGenericPredictionRequest.Superclass]
type IMLGenericPredictionRequest interface {
	objectivec.IObject

	// Topic: Methods

	Cancel()
	InputFeatures() unsafe.Pointer
	IsCancelled() bool
	Model() IMLModel
	PredictionOptions() IMLPredictionOptions
	SubmitWithCompletionHandler(handler ErrorHandler)
	InitForModelInputFeaturesOptions(model objectivec.IObject, features objectivec.IObject, options objectivec.IObject) MLGenericPredictionRequest
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLGenericPredictionRequest) Init() MLGenericPredictionRequest {
	rv := objc.Send[MLGenericPredictionRequest](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLGenericPredictionRequest) Autorelease() MLGenericPredictionRequest {
	rv := objc.Send[MLGenericPredictionRequest](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLGenericPredictionRequest creates a new MLGenericPredictionRequest instance.
func NewMLGenericPredictionRequest() MLGenericPredictionRequest {
	class := getMLGenericPredictionRequestClass()
	rv := objc.Send[MLGenericPredictionRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGenericPredictionRequestForModelInputFeaturesOptions(model objectivec.IObject, features objectivec.IObject, options objectivec.IObject) MLGenericPredictionRequest {
	instance := getMLGenericPredictionRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initForModel:inputFeatures:options:"), model, features, options)
	return MLGenericPredictionRequestFromID(rv)
}

func (m MLGenericPredictionRequest) Cancel() {
	objc.Send[objc.ID](m.ID, objc.Sel("cancel"))
}
func (m MLGenericPredictionRequest) SubmitWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("submitWithCompletionHandler:"), _block0)
}
func (m MLGenericPredictionRequest) InitForModelInputFeaturesOptions(model objectivec.IObject, features objectivec.IObject, options objectivec.IObject) MLGenericPredictionRequest {
	rv := objc.Send[MLGenericPredictionRequest](m.ID, objc.Sel("initForModel:inputFeatures:options:"), model, features, options)
	return rv
}

func (m MLGenericPredictionRequest) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLGenericPredictionRequest) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLGenericPredictionRequest) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLGenericPredictionRequest) InputFeatures() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("inputFeatures"))
	return rv
}
func (m MLGenericPredictionRequest) IsCancelled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isCancelled"))
	return rv
}
func (m MLGenericPredictionRequest) Model() IMLModel {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("model"))
	return MLModelFromID(objc.ID(rv))
}
func (m MLGenericPredictionRequest) PredictionOptions() IMLPredictionOptions {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionOptions"))
	return MLPredictionOptionsFromID(objc.ID(rv))
}
func (m MLGenericPredictionRequest) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

// Submit is a synchronous wrapper around [MLGenericPredictionRequest.SubmitWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLGenericPredictionRequest) Submit(ctx context.Context) error {
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
