// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
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

// # Methods
//
//   - [MLState.Backings]
//   - [MLState.FeatureProviderRepresentation]
//   - [MLState.GetMultiArrayWithHandler]
//   - [MLState.InternalGetMultiArrayWithHandler]
//   - [MLState.InitWithBackings]
//
// See: https://developer.apple.com/documentation/CoreML/MLState
type MLState struct {
	objectivec.Object
}

// MLStateFromID constructs a [MLState] from an objc.ID.
func MLStateFromID(id objc.ID) MLState {
	return MLState{objectivec.Object{ID: id}}
}

// Ensure MLState implements IMLState.
var _ IMLState = MLState{}

// An interface definition for the [MLState] class.
//
// # Methods
//
//   - [IMLState.Backings]
//   - [IMLState.FeatureProviderRepresentation]
//   - [IMLState.GetMultiArrayWithHandler]
//   - [IMLState.InternalGetMultiArrayWithHandler]
//   - [IMLState.InitWithBackings]
//
// See: https://developer.apple.com/documentation/CoreML/MLState
type IMLState interface {
	objectivec.IObject

	// Topic: Methods

	Backings() foundation.INSDictionary
	FeatureProviderRepresentation() objectivec.IObject
	GetMultiArrayWithHandler(handler VoidHandler)
	InternalGetMultiArrayWithHandler(handler VoidHandler)
	InitWithBackings(backings objectivec.IObject) MLState
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

// See: https://developer.apple.com/documentation/CoreML/MLState/initWithBackings:
func NewStateWithBackings(backings objectivec.IObject) MLState {
	instance := getMLStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackings:"), backings)
	return MLStateFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLState/getMultiArrayWithHandler:
func (s MLState) GetMultiArrayWithHandler(handler VoidHandler) {
	_block0, _ := NewVoidBlock(handler)
	objc.Send[objc.ID](s.ID, objc.Sel("getMultiArrayWithHandler:"), _block0)
}

// See: https://developer.apple.com/documentation/CoreML/MLState/internalGetMultiArrayWithHandler:
func (s MLState) InternalGetMultiArrayWithHandler(handler VoidHandler) {
	_block0, _ := NewVoidBlock(handler)
	objc.Send[objc.ID](s.ID, objc.Sel("internalGetMultiArrayWithHandler:"), _block0)
}

// See: https://developer.apple.com/documentation/CoreML/MLState/initWithBackings:
func (s MLState) InitWithBackings(backings objectivec.IObject) MLState {
	rv := objc.Send[MLState](s.ID, objc.Sel("initWithBackings:"), backings)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLState/emptyState
func (_MLStateClass MLStateClass) EmptyState() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLStateClass.class), objc.Sel("emptyState"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLState/backings
func (s MLState) Backings() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("backings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLState/featureProviderRepresentation
func (s MLState) FeatureProviderRepresentation() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("featureProviderRepresentation"))
	return objectivec.Object{ID: rv}
}

// GetMultiArrayWithHandlerSync is a synchronous wrapper around [MLState.GetMultiArrayWithHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s MLState) GetMultiArrayWithHandlerSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	s.GetMultiArrayWithHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// InternalGetMultiArrayWithHandlerSync is a synchronous wrapper around [MLState.InternalGetMultiArrayWithHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s MLState) InternalGetMultiArrayWithHandlerSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	s.InternalGetMultiArrayWithHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
