// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"sync"

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
//
// See: https://developer.apple.com/documentation/CoreML/MLGenericPredictionRequest
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
//
// See: https://developer.apple.com/documentation/CoreML/MLGenericPredictionRequest
type IMLGenericPredictionRequest interface {
	objectivec.IObject

	// Topic: Methods

	Cancel()
	InputFeatures() objectivec.IObject
	IsCancelled() bool
	Model() IMLModel
	PredictionOptions() IMLPredictionOptions
	SubmitWithCompletionHandler(handler ErrorHandler)
	InitForModelInputFeaturesOptions(model objectivec.IObject, features objectivec.IObject, options objectivec.IObject) MLGenericPredictionRequest
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (g MLGenericPredictionRequest) Init() MLGenericPredictionRequest {
	rv := objc.Send[MLGenericPredictionRequest](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MLGenericPredictionRequest) Autorelease() MLGenericPredictionRequest {
	rv := objc.Send[MLGenericPredictionRequest](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLGenericPredictionRequest creates a new MLGenericPredictionRequest instance.
func NewMLGenericPredictionRequest() MLGenericPredictionRequest {
	class := getMLGenericPredictionRequestClass()
	rv := objc.Send[MLGenericPredictionRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLGenericPredictionRequest/initForModel:inputFeatures:options:
func NewGenericPredictionRequestForModelInputFeaturesOptions(model objectivec.IObject, features objectivec.IObject, options objectivec.IObject) MLGenericPredictionRequest {
	instance := getMLGenericPredictionRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initForModel:inputFeatures:options:"), model, features, options)
	return MLGenericPredictionRequestFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLGenericPredictionRequest/cancel
func (g MLGenericPredictionRequest) Cancel() {
	objc.Send[objc.ID](g.ID, objc.Sel("cancel"))
}

// See: https://developer.apple.com/documentation/CoreML/MLGenericPredictionRequest/submitWithCompletionHandler:
func (g MLGenericPredictionRequest) SubmitWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](g.ID, objc.Sel("submitWithCompletionHandler:"), _block0)
}

// See: https://developer.apple.com/documentation/CoreML/MLGenericPredictionRequest/initForModel:inputFeatures:options:
func (g MLGenericPredictionRequest) InitForModelInputFeaturesOptions(model objectivec.IObject, features objectivec.IObject, options objectivec.IObject) MLGenericPredictionRequest {
	rv := objc.Send[MLGenericPredictionRequest](g.ID, objc.Sel("initForModel:inputFeatures:options:"), model, features, options)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLGenericPredictionRequest/debugDescription
func (g MLGenericPredictionRequest) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLGenericPredictionRequest/description
func (g MLGenericPredictionRequest) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLGenericPredictionRequest/hash
func (g MLGenericPredictionRequest) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLGenericPredictionRequest/inputFeatures
func (g MLGenericPredictionRequest) InputFeatures() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("inputFeatures"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLGenericPredictionRequest/isCancelled
func (g MLGenericPredictionRequest) IsCancelled() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isCancelled"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLGenericPredictionRequest/model
func (g MLGenericPredictionRequest) Model() IMLModel {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("model"))
	return MLModelFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLGenericPredictionRequest/predictionOptions
func (g MLGenericPredictionRequest) PredictionOptions() IMLPredictionOptions {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("predictionOptions"))
	return MLPredictionOptionsFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLGenericPredictionRequest/superclass
func (g MLGenericPredictionRequest) Superclass() objc.Class {
	rv := objc.Send[objc.Class](g.ID, objc.Sel("superclass"))
	return rv
}

// Submit is a synchronous wrapper around [MLGenericPredictionRequest.SubmitWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (g MLGenericPredictionRequest) Submit(ctx context.Context) error {
	done := make(chan error, 1)
	g.SubmitWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
