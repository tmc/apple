// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelAssetModelStructureVendor] class.
var (
	_MLModelAssetModelStructureVendorClass     MLModelAssetModelStructureVendorClass
	_MLModelAssetModelStructureVendorClassOnce sync.Once
)

func getMLModelAssetModelStructureVendorClass() MLModelAssetModelStructureVendorClass {
	_MLModelAssetModelStructureVendorClassOnce.Do(func() {
		_MLModelAssetModelStructureVendorClass = MLModelAssetModelStructureVendorClass{class: objc.GetClass("MLModelAssetModelStructureVendor")}
	})
	return _MLModelAssetModelStructureVendorClass
}

// GetMLModelAssetModelStructureVendorClass returns the class object for MLModelAssetModelStructureVendor.
func GetMLModelAssetModelStructureVendorClass() MLModelAssetModelStructureVendorClass {
	return getMLModelAssetModelStructureVendorClass()
}

type MLModelAssetModelStructureVendorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelAssetModelStructureVendorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelAssetModelStructureVendorClass) Alloc() MLModelAssetModelStructureVendor {
	rv := objc.Send[MLModelAssetModelStructureVendor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLModelAssetModelStructureVendor._modelStructureWithCompletionHandler]
//   - [MLModelAssetModelStructureVendor.ModelStructure]
//   - [MLModelAssetModelStructureVendor.SetModelStructure]
//   - [MLModelAssetModelStructureVendor.ModelStructureWithCompletionHandler]
//   - [MLModelAssetModelStructureVendor.ResourceFactory]
//   - [MLModelAssetModelStructureVendor.InitWithResourceFactory]
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetModelStructureVendor
type MLModelAssetModelStructureVendor struct {
	objectivec.Object
}

// MLModelAssetModelStructureVendorFromID constructs a [MLModelAssetModelStructureVendor] from an objc.ID.
func MLModelAssetModelStructureVendorFromID(id objc.ID) MLModelAssetModelStructureVendor {
	return MLModelAssetModelStructureVendor{objectivec.Object{ID: id}}
}
// Ensure MLModelAssetModelStructureVendor implements IMLModelAssetModelStructureVendor.
var _ IMLModelAssetModelStructureVendor = MLModelAssetModelStructureVendor{}

// An interface definition for the [MLModelAssetModelStructureVendor] class.
//
// # Methods
//
//   - [IMLModelAssetModelStructureVendor._modelStructureWithCompletionHandler]
//   - [IMLModelAssetModelStructureVendor.ModelStructure]
//   - [IMLModelAssetModelStructureVendor.SetModelStructure]
//   - [IMLModelAssetModelStructureVendor.ModelStructureWithCompletionHandler]
//   - [IMLModelAssetModelStructureVendor.ResourceFactory]
//   - [IMLModelAssetModelStructureVendor.InitWithResourceFactory]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetModelStructureVendor
type IMLModelAssetModelStructureVendor interface {
	objectivec.IObject

	// Topic: Methods

	_modelStructureWithCompletionHandler(handler ErrorHandler)
	ModelStructure() IMLModelStructure
	SetModelStructure(value IMLModelStructure)
	ModelStructureWithCompletionHandler(handler ErrorHandler)
	ResourceFactory() IMLModelAssetResourceFactory
	InitWithResourceFactory(factory objectivec.IObject) MLModelAssetModelStructureVendor
}

// Init initializes the instance.
func (m MLModelAssetModelStructureVendor) Init() MLModelAssetModelStructureVendor {
	rv := objc.Send[MLModelAssetModelStructureVendor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelAssetModelStructureVendor) Autorelease() MLModelAssetModelStructureVendor {
	rv := objc.Send[MLModelAssetModelStructureVendor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelAssetModelStructureVendor creates a new MLModelAssetModelStructureVendor instance.
func NewMLModelAssetModelStructureVendor() MLModelAssetModelStructureVendor {
	class := getMLModelAssetModelStructureVendorClass()
	rv := objc.Send[MLModelAssetModelStructureVendor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetModelStructureVendor/initWithResourceFactory:
func NewModelAssetModelStructureVendorWithResourceFactory(factory objectivec.IObject) MLModelAssetModelStructureVendor {
	instance := getMLModelAssetModelStructureVendorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResourceFactory:"), factory)
	return MLModelAssetModelStructureVendorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetModelStructureVendor/_modelStructureWithCompletionHandler:
func (m MLModelAssetModelStructureVendor) _modelStructureWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("_modelStructureWithCompletionHandler:"), _block0)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetModelStructureVendor/modelStructureWithCompletionHandler:
func (m MLModelAssetModelStructureVendor) ModelStructureWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("modelStructureWithCompletionHandler:"), _block0)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetModelStructureVendor/initWithResourceFactory:
func (m MLModelAssetModelStructureVendor) InitWithResourceFactory(factory objectivec.IObject) MLModelAssetModelStructureVendor {
	rv := objc.Send[MLModelAssetModelStructureVendor](m.ID, objc.Sel("initWithResourceFactory:"), factory)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelAssetModelStructureVendor/modelStructure
func (m MLModelAssetModelStructureVendor) ModelStructure() IMLModelStructure {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelStructure"))
	return MLModelStructureFromID(objc.ID(rv))
}
func (m MLModelAssetModelStructureVendor) SetModelStructure(value IMLModelStructure) {
	objc.Send[struct{}](m.ID, objc.Sel("setModelStructure:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetModelStructureVendor/resourceFactory
func (m MLModelAssetModelStructureVendor) ResourceFactory() IMLModelAssetResourceFactory {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("resourceFactory"))
	return MLModelAssetResourceFactoryFromID(objc.ID(rv))
}

// _modelStructure is a synchronous wrapper around [MLModelAssetModelStructureVendor._modelStructureWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModelAssetModelStructureVendor) _modelStructure(ctx context.Context) error {
	done := make(chan error, 1)
	m._modelStructureWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ModelStructureSync is a synchronous wrapper around [MLModelAssetModelStructureVendor.ModelStructureWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModelAssetModelStructureVendor) ModelStructureSync(ctx context.Context) error {
	done := make(chan error, 1)
	m.ModelStructureWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

