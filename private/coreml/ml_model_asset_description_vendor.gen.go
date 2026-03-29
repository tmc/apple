// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelAssetDescriptionVendor] class.
var (
	_MLModelAssetDescriptionVendorClass     MLModelAssetDescriptionVendorClass
	_MLModelAssetDescriptionVendorClassOnce sync.Once
)

func getMLModelAssetDescriptionVendorClass() MLModelAssetDescriptionVendorClass {
	_MLModelAssetDescriptionVendorClassOnce.Do(func() {
		_MLModelAssetDescriptionVendorClass = MLModelAssetDescriptionVendorClass{class: objc.GetClass("MLModelAssetDescriptionVendor")}
	})
	return _MLModelAssetDescriptionVendorClass
}

// GetMLModelAssetDescriptionVendorClass returns the class object for MLModelAssetDescriptionVendor.
func GetMLModelAssetDescriptionVendorClass() MLModelAssetDescriptionVendorClass {
	return getMLModelAssetDescriptionVendorClass()
}

type MLModelAssetDescriptionVendorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelAssetDescriptionVendorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelAssetDescriptionVendorClass) Alloc() MLModelAssetDescriptionVendor {
	rv := objc.Send[MLModelAssetDescriptionVendor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLModelAssetDescriptionVendor._modelAssetDescriptionWithCompletionHandler]
//   - [MLModelAssetDescriptionVendor.FunctionNamesWithCompletionHandler]
//   - [MLModelAssetDescriptionVendor.ModelAssetDescription]
//   - [MLModelAssetDescriptionVendor.SetModelAssetDescription]
//   - [MLModelAssetDescriptionVendor.ModelDescriptionOfFunctionNamedCompletionHandler]
//   - [MLModelAssetDescriptionVendor.ModelDescriptionWithCompletionHandler]
//   - [MLModelAssetDescriptionVendor.ResourceFactory]
//   - [MLModelAssetDescriptionVendor.InitWithResourceFactory]
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetDescriptionVendor
type MLModelAssetDescriptionVendor struct {
	objectivec.Object
}

// MLModelAssetDescriptionVendorFromID constructs a [MLModelAssetDescriptionVendor] from an objc.ID.
func MLModelAssetDescriptionVendorFromID(id objc.ID) MLModelAssetDescriptionVendor {
	return MLModelAssetDescriptionVendor{objectivec.Object{ID: id}}
}
// Ensure MLModelAssetDescriptionVendor implements IMLModelAssetDescriptionVendor.
var _ IMLModelAssetDescriptionVendor = MLModelAssetDescriptionVendor{}

// An interface definition for the [MLModelAssetDescriptionVendor] class.
//
// # Methods
//
//   - [IMLModelAssetDescriptionVendor._modelAssetDescriptionWithCompletionHandler]
//   - [IMLModelAssetDescriptionVendor.FunctionNamesWithCompletionHandler]
//   - [IMLModelAssetDescriptionVendor.ModelAssetDescription]
//   - [IMLModelAssetDescriptionVendor.SetModelAssetDescription]
//   - [IMLModelAssetDescriptionVendor.ModelDescriptionOfFunctionNamedCompletionHandler]
//   - [IMLModelAssetDescriptionVendor.ModelDescriptionWithCompletionHandler]
//   - [IMLModelAssetDescriptionVendor.ResourceFactory]
//   - [IMLModelAssetDescriptionVendor.InitWithResourceFactory]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetDescriptionVendor
type IMLModelAssetDescriptionVendor interface {
	objectivec.IObject

	// Topic: Methods

	_modelAssetDescriptionWithCompletionHandler(handler ErrorHandler)
	FunctionNamesWithCompletionHandler(handler ErrorHandler)
	ModelAssetDescription() IMLModelAssetDescription
	SetModelAssetDescription(value IMLModelAssetDescription)
	ModelDescriptionOfFunctionNamedCompletionHandler(named objectivec.IObject, handler ErrorHandler)
	ModelDescriptionWithCompletionHandler(handler ErrorHandler)
	ResourceFactory() IMLModelAssetResourceFactory
	InitWithResourceFactory(factory objectivec.IObject) MLModelAssetDescriptionVendor
}

// Init initializes the instance.
func (m MLModelAssetDescriptionVendor) Init() MLModelAssetDescriptionVendor {
	rv := objc.Send[MLModelAssetDescriptionVendor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelAssetDescriptionVendor) Autorelease() MLModelAssetDescriptionVendor {
	rv := objc.Send[MLModelAssetDescriptionVendor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelAssetDescriptionVendor creates a new MLModelAssetDescriptionVendor instance.
func NewMLModelAssetDescriptionVendor() MLModelAssetDescriptionVendor {
	class := getMLModelAssetDescriptionVendorClass()
	rv := objc.Send[MLModelAssetDescriptionVendor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetDescriptionVendor/initWithResourceFactory:
func NewModelAssetDescriptionVendorWithResourceFactory(factory objectivec.IObject) MLModelAssetDescriptionVendor {
	instance := getMLModelAssetDescriptionVendorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResourceFactory:"), factory)
	return MLModelAssetDescriptionVendorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetDescriptionVendor/_modelAssetDescriptionWithCompletionHandler:
func (m MLModelAssetDescriptionVendor) _modelAssetDescriptionWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("_modelAssetDescriptionWithCompletionHandler:"), _block0)
}

// ModelAssetDescriptionWithCompletionHandler is an exported wrapper for the private method _modelAssetDescriptionWithCompletionHandler.
func (m MLModelAssetDescriptionVendor) ModelAssetDescriptionWithCompletionHandler(handler ErrorHandler) {
	m._modelAssetDescriptionWithCompletionHandler(handler)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetDescriptionVendor/functionNamesWithCompletionHandler:
func (m MLModelAssetDescriptionVendor) FunctionNamesWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("functionNamesWithCompletionHandler:"), _block0)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetDescriptionVendor/modelDescriptionOfFunctionNamed:completionHandler:
func (m MLModelAssetDescriptionVendor) ModelDescriptionOfFunctionNamedCompletionHandler(named objectivec.IObject, handler ErrorHandler) {
_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("modelDescriptionOfFunctionNamed:completionHandler:"), named, _block1)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetDescriptionVendor/modelDescriptionWithCompletionHandler:
func (m MLModelAssetDescriptionVendor) ModelDescriptionWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("modelDescriptionWithCompletionHandler:"), _block0)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetDescriptionVendor/initWithResourceFactory:
func (m MLModelAssetDescriptionVendor) InitWithResourceFactory(factory objectivec.IObject) MLModelAssetDescriptionVendor {
	rv := objc.Send[MLModelAssetDescriptionVendor](m.ID, objc.Sel("initWithResourceFactory:"), factory)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelAssetDescriptionVendor/modelAssetDescription
func (m MLModelAssetDescriptionVendor) ModelAssetDescription() IMLModelAssetDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelAssetDescription"))
	return MLModelAssetDescriptionFromID(objc.ID(rv))
}
func (m MLModelAssetDescriptionVendor) SetModelAssetDescription(value IMLModelAssetDescription) {
	objc.Send[struct{}](m.ID, objc.Sel("setModelAssetDescription:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetDescriptionVendor/resourceFactory
func (m MLModelAssetDescriptionVendor) ResourceFactory() IMLModelAssetResourceFactory {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("resourceFactory"))
	return MLModelAssetResourceFactoryFromID(objc.ID(rv))
}

// _modelAssetDescription is a synchronous wrapper around [MLModelAssetDescriptionVendor._modelAssetDescriptionWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModelAssetDescriptionVendor) _modelAssetDescription(ctx context.Context) error {
	done := make(chan error, 1)
	m._modelAssetDescriptionWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// FunctionNames is a synchronous wrapper around [MLModelAssetDescriptionVendor.FunctionNamesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModelAssetDescriptionVendor) FunctionNames(ctx context.Context) error {
	done := make(chan error, 1)
	m.FunctionNamesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ModelDescriptionOfFunctionNamed is a synchronous wrapper around [MLModelAssetDescriptionVendor.ModelDescriptionOfFunctionNamedCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModelAssetDescriptionVendor) ModelDescriptionOfFunctionNamed(ctx context.Context, named objectivec.IObject) error {
	done := make(chan error, 1)
	m.ModelDescriptionOfFunctionNamedCompletionHandler(named, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ModelDescription is a synchronous wrapper around [MLModelAssetDescriptionVendor.ModelDescriptionWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModelAssetDescriptionVendor) ModelDescription(ctx context.Context) error {
	done := make(chan error, 1)
	m.ModelDescriptionWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

