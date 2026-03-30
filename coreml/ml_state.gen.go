// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLState] class.
var (
	_MLStateClass     MLStateClass
	_MLStateClassOnce sync.Once
)

func getMLStateClass() MLStateClass {
	_MLStateClassOnce.Do(func() {
		_MLStateClass = MLStateClass{class: objc.GetClass("MLState")}
	})
	return _MLStateClass
}

// GetMLStateClass returns the class object for MLState.
func GetMLStateClass() MLStateClass {
	return getMLStateClass()
}

type MLStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLStateClass) Alloc() MLState {
	rv := objc.Send[MLState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Handle to the state buffers.
//
// # Overview
//
// A stateful model maintains a state from one prediction to another by
// storing the information in the state buffers. To use such a model, the
// client must request the model to create state buffers and get [MLState]
// object, which is the handle to those buffers. Then, at the prediction time,
// pass the [MLState] object in one of the stateful prediction functions.
//
// The object is a handle to the state buffers. The client shall not read or
// write the buffers while a prediction is in-flight.
//
// Each stateful prediction that uses the same [MLState] must be serialized.
// Otherwise, if two such predictions run concurrently, the behavior is
// undefined.
//
// See: https://developer.apple.com/documentation/CoreML/MLState
type MLState struct {
	objectivec.Object
}

// MLStateFromID constructs a [MLState] from an objc.ID.
//
// Handle to the state buffers.
func MLStateFromID(id objc.ID) MLState {
	return MLState{objectivec.Object{ID: id}}
}

// NOTE: MLState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLState] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLState
type IMLState interface {
	objectivec.IObject

	// Gets a mutable view into a state buffer.
	GetMultiArrayForStateNamedHandler(stateName string, handler MLMultiArrayHandler)
}

// Init initializes the instance.
func (s MLState) Init() MLState {
	rv := objc.Send[MLState](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MLState) Autorelease() MLState {
	rv := objc.Send[MLState](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLState creates a new MLState instance.
func NewMLState() MLState {
	class := getMLStateClass()
	rv := objc.Send[MLState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Gets a mutable view into a state buffer.
//
// handler: Block to access the state buffer through [MLMultiArray].
//
// # Discussion
//
// The underlying state buffer’s address can differ for each call; one shall
// not access the state buffer outside of the closure.
//
// See: https://developer.apple.com/documentation/CoreML/MLState/getMultiArrayForStateNamed:handler:
func (s MLState) GetMultiArrayForStateNamedHandler(stateName string, handler MLMultiArrayHandler) {
	_block1, _ := NewMLMultiArrayBlock(handler)
	objc.Send[objc.ID](s.ID, objc.Sel("getMultiArrayForStateNamed:handler:"), objc.String(stateName), _block1)
}

// GetMultiArrayForStateNamedHandlerSync is a synchronous wrapper around [MLState.GetMultiArrayForStateNamedHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s MLState) GetMultiArrayForStateNamedHandlerSync(ctx context.Context, stateName string) (*MLMultiArray, error) {
	done := make(chan *MLMultiArray, 1)
	s.GetMultiArrayForStateNamedHandler(stateName, func(val *MLMultiArray) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
