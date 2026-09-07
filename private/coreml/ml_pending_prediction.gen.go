// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLPendingPrediction] class.
var (
	_MLPendingPredictionClass     MLPendingPredictionClass
	_MLPendingPredictionClassOnce sync.Once
)

func getMLPendingPredictionClass() MLPendingPredictionClass {
	_MLPendingPredictionClassOnce.Do(func() {
		_MLPendingPredictionClass = MLPendingPredictionClass{class: objc.GetClass("MLPendingPrediction")}
	})
	return _MLPendingPredictionClass
}

// GetMLPendingPredictionClass returns the class object for MLPendingPrediction.
func GetMLPendingPredictionClass() MLPendingPredictionClass {
	return getMLPendingPredictionClass()
}

type MLPendingPredictionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLPendingPredictionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLPendingPredictionClass) Alloc() MLPendingPrediction {
	rv := objc.SendIfResponds[MLPendingPrediction](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLPendingPrediction.CompletionHandler]
//   - [MLPendingPrediction.PredictionRequest]
//   - [MLPendingPrediction.InitWithPredictionRequestCompletionHandler]
type MLPendingPrediction struct {
	objectivec.Object
}

// MLPendingPredictionFromID constructs a [MLPendingPrediction] from an objc.ID.
func MLPendingPredictionFromID(id objc.ID) MLPendingPrediction {
	return MLPendingPrediction{objectivec.Object{ID: id}}
}

// Ensure MLPendingPrediction implements IMLPendingPrediction.
var _ IMLPendingPrediction = MLPendingPrediction{}

// An interface definition for the [MLPendingPrediction] class.
//
// # Methods
//
//   - [IMLPendingPrediction.CompletionHandler]
//   - [IMLPendingPrediction.PredictionRequest]
//   - [IMLPendingPrediction.InitWithPredictionRequestCompletionHandler]
type IMLPendingPrediction interface {
	objectivec.IObject

	// Topic: Methods

	CompletionHandler() VoidHandler
	PredictionRequest() unsafe.Pointer
	InitWithPredictionRequestCompletionHandler(request objectivec.IObject, handler ErrorHandler) MLPendingPrediction
}

// Init initializes the instance.
func (m MLPendingPrediction) Init() MLPendingPrediction {
	rv := objc.SendIfResponds[MLPendingPrediction](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLPendingPrediction) Autorelease() MLPendingPrediction {
	rv := objc.SendIfResponds[MLPendingPrediction](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPendingPrediction creates a new MLPendingPrediction instance.
func NewMLPendingPrediction() MLPendingPrediction {
	class := getMLPendingPredictionClass()
	rv := objc.SendIfResponds[MLPendingPrediction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m MLPendingPrediction) InitWithPredictionRequestCompletionHandler(request objectivec.IObject, handler ErrorHandler) MLPendingPrediction {
	_block1, _ := NewErrorBlock(handler)
	rv := objc.SendIfResponds[MLPendingPrediction](m.ID, objc.Sel("initWithPredictionRequest:completionHandler:"), request, _block1)
	return rv
}

func (m MLPendingPrediction) CompletionHandler() VoidHandler {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("completionHandler"))
	_ = rv
	return nil
}
func (m MLPendingPrediction) PredictionRequest() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("predictionRequest"))
	return rv
}

// InitWithPredictionRequest is a synchronous wrapper around [MLPendingPrediction.InitWithPredictionRequestCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLPendingPrediction) InitWithPredictionRequest(ctx context.Context, request objectivec.IObject) error {
	done := make(chan error, 1)
	m.InitWithPredictionRequestCompletionHandler(request, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
