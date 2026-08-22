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
	rv := objc.SendIfResponds[MLState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLState.Backings]
//   - [MLState.FeatureProviderRepresentation]
//   - [MLState.GetMultiArrayWithHandler]
//   - [MLState.InternalGetMultiArrayWithHandler]
//   - [MLState.InitWithBackings]
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
type IMLState interface {
	objectivec.IObject

	// Topic: Methods

	Backings() foundation.INSDictionary
	FeatureProviderRepresentation() unsafe.Pointer
	GetMultiArrayWithHandler(handler VoidHandler)
	InternalGetMultiArrayWithHandler(handler VoidHandler)
	InitWithBackings(backings objectivec.IObject) MLState
}

// Init initializes the instance.
func (m MLState) Init() MLState {
	rv := objc.SendIfResponds[MLState](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLState) Autorelease() MLState {
	rv := objc.SendIfResponds[MLState](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLState creates a new MLState instance.
func NewMLState() MLState {
	class := getMLStateClass()
	rv := objc.SendIfResponds[MLState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewStateWithBackings(backings objectivec.IObject) MLState {
	instance := getMLStateClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBackings:"), backings)
	return MLStateFromID(rv)
}

func (m MLState) GetMultiArrayWithHandler(handler VoidHandler) {
	_block0, _ := NewVoidBlock(handler)
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("getMultiArrayWithHandler:"), _block0)
}
func (m MLState) InternalGetMultiArrayWithHandler(handler VoidHandler) {
	_block0, _ := NewVoidBlock(handler)
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("internalGetMultiArrayWithHandler:"), _block0)
}
func (m MLState) InitWithBackings(backings objectivec.IObject) MLState {
	rv := objc.SendIfResponds[MLState](m.ID, objc.Sel("initWithBackings:"), backings)
	return rv
}

func (_MLStateClass MLStateClass) EmptyState() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLStateClass.class), objc.Sel("emptyState"))
	return objectivec.Object{ID: rv}
}

func (m MLState) Backings() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("backings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLState) FeatureProviderRepresentation() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("featureProviderRepresentation"))
	return rv
}

// GetMultiArrayWithHandlerSync is a synchronous wrapper around [MLState.GetMultiArrayWithHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLState) GetMultiArrayWithHandlerSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	m.GetMultiArrayWithHandler(func() {
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
func (m MLState) InternalGetMultiArrayWithHandlerSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	m.InternalGetMultiArrayWithHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
