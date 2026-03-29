// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelAssetResourceFactory] class.
var (
	_MLModelAssetResourceFactoryClass     MLModelAssetResourceFactoryClass
	_MLModelAssetResourceFactoryClassOnce sync.Once
)

func getMLModelAssetResourceFactoryClass() MLModelAssetResourceFactoryClass {
	_MLModelAssetResourceFactoryClassOnce.Do(func() {
		_MLModelAssetResourceFactoryClass = MLModelAssetResourceFactoryClass{class: objc.GetClass("MLModelAssetResourceFactory")}
	})
	return _MLModelAssetResourceFactoryClass
}

// GetMLModelAssetResourceFactoryClass returns the class object for MLModelAssetResourceFactory.
func GetMLModelAssetResourceFactoryClass() MLModelAssetResourceFactoryClass {
	return getMLModelAssetResourceFactoryClass()
}

type MLModelAssetResourceFactoryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelAssetResourceFactoryClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelAssetResourceFactoryClass) Alloc() MLModelAssetResourceFactory {
	rv := objc.Send[MLModelAssetResourceFactory](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLModelAssetResourceFactory.CompiledModelURL]
//   - [MLModelAssetResourceFactory.DescriptionLoadQueue]
//   - [MLModelAssetResourceFactory.Impl]
//   - [MLModelAssetResourceFactory.ModelAssetDescriptionWithCompletionHandler]
//   - [MLModelAssetResourceFactory.ModelLoadQueue]
//   - [MLModelAssetResourceFactory.ModelStructureWithCompletionHandler]
//   - [MLModelAssetResourceFactory.ModelWithConfigurationCompletionHandler]
//   - [MLModelAssetResourceFactory.ModelWithConfigurationError]
//   - [MLModelAssetResourceFactory.StructureLoadQueue]
//   - [MLModelAssetResourceFactory.InitWithImpl]
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactory
type MLModelAssetResourceFactory struct {
	objectivec.Object
}

// MLModelAssetResourceFactoryFromID constructs a [MLModelAssetResourceFactory] from an objc.ID.
func MLModelAssetResourceFactoryFromID(id objc.ID) MLModelAssetResourceFactory {
	return MLModelAssetResourceFactory{objectivec.Object{ID: id}}
}
// Ensure MLModelAssetResourceFactory implements IMLModelAssetResourceFactory.
var _ IMLModelAssetResourceFactory = MLModelAssetResourceFactory{}

// An interface definition for the [MLModelAssetResourceFactory] class.
//
// # Methods
//
//   - [IMLModelAssetResourceFactory.CompiledModelURL]
//   - [IMLModelAssetResourceFactory.DescriptionLoadQueue]
//   - [IMLModelAssetResourceFactory.Impl]
//   - [IMLModelAssetResourceFactory.ModelAssetDescriptionWithCompletionHandler]
//   - [IMLModelAssetResourceFactory.ModelLoadQueue]
//   - [IMLModelAssetResourceFactory.ModelStructureWithCompletionHandler]
//   - [IMLModelAssetResourceFactory.ModelWithConfigurationCompletionHandler]
//   - [IMLModelAssetResourceFactory.ModelWithConfigurationError]
//   - [IMLModelAssetResourceFactory.StructureLoadQueue]
//   - [IMLModelAssetResourceFactory.InitWithImpl]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactory
type IMLModelAssetResourceFactory interface {
	objectivec.IObject

	// Topic: Methods

	CompiledModelURL() foundation.INSURL
	DescriptionLoadQueue() objectivec.Object
	Impl() objectivec.IObject
	ModelAssetDescriptionWithCompletionHandler(handler ErrorHandler)
	ModelLoadQueue() objectivec.Object
	ModelStructureWithCompletionHandler(handler ErrorHandler)
	ModelWithConfigurationCompletionHandler(configuration objectivec.IObject, handler ErrorHandler)
	ModelWithConfigurationError(configuration objectivec.IObject) (objectivec.IObject, error)
	StructureLoadQueue() objectivec.Object
	InitWithImpl(impl objectivec.IObject) MLModelAssetResourceFactory
}

// Init initializes the instance.
func (m MLModelAssetResourceFactory) Init() MLModelAssetResourceFactory {
	rv := objc.Send[MLModelAssetResourceFactory](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelAssetResourceFactory) Autorelease() MLModelAssetResourceFactory {
	rv := objc.Send[MLModelAssetResourceFactory](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelAssetResourceFactory creates a new MLModelAssetResourceFactory instance.
func NewMLModelAssetResourceFactory() MLModelAssetResourceFactory {
	class := getMLModelAssetResourceFactoryClass()
	rv := objc.Send[MLModelAssetResourceFactory](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactory/initWithImpl:
func NewModelAssetResourceFactoryWithImpl(impl objectivec.IObject) MLModelAssetResourceFactory {
	instance := getMLModelAssetResourceFactoryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImpl:"), impl)
	return MLModelAssetResourceFactoryFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactory/modelAssetDescriptionWithCompletionHandler:
func (m MLModelAssetResourceFactory) ModelAssetDescriptionWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("modelAssetDescriptionWithCompletionHandler:"), _block0)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactory/modelStructureWithCompletionHandler:
func (m MLModelAssetResourceFactory) ModelStructureWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("modelStructureWithCompletionHandler:"), _block0)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactory/modelWithConfiguration:completionHandler:
func (m MLModelAssetResourceFactory) ModelWithConfigurationCompletionHandler(configuration objectivec.IObject, handler ErrorHandler) {
_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("modelWithConfiguration:completionHandler:"), configuration, _block1)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactory/modelWithConfiguration:error:
func (m MLModelAssetResourceFactory) ModelWithConfigurationError(configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactory/initWithImpl:
func (m MLModelAssetResourceFactory) InitWithImpl(impl objectivec.IObject) MLModelAssetResourceFactory {
	rv := objc.Send[MLModelAssetResourceFactory](m.ID, objc.Sel("initWithImpl:"), impl)
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactory/resourceFactoryWithArchiveData:
func (_MLModelAssetResourceFactoryClass MLModelAssetResourceFactoryClass) ResourceFactoryWithArchiveData(data objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetResourceFactoryClass.class), objc.Sel("resourceFactoryWithArchiveData:"), data)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactory/resourceFactoryWithModelURL:error:
func (_MLModelAssetResourceFactoryClass MLModelAssetResourceFactoryClass) ResourceFactoryWithModelURLError(url foundation.INSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelAssetResourceFactoryClass.class), objc.Sel("resourceFactoryWithModelURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactory/compiledModelURL
func (m MLModelAssetResourceFactory) CompiledModelURL() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("compiledModelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactory/descriptionLoadQueue
func (m MLModelAssetResourceFactory) DescriptionLoadQueue() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("descriptionLoadQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactory/impl
func (m MLModelAssetResourceFactory) Impl() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("impl"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactory/modelLoadQueue
func (m MLModelAssetResourceFactory) ModelLoadQueue() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelLoadQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetResourceFactory/structureLoadQueue
func (m MLModelAssetResourceFactory) StructureLoadQueue() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("structureLoadQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

// ModelAssetDescription is a synchronous wrapper around [MLModelAssetResourceFactory.ModelAssetDescriptionWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModelAssetResourceFactory) ModelAssetDescription(ctx context.Context) error {
	done := make(chan error, 1)
	m.ModelAssetDescriptionWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ModelStructure is a synchronous wrapper around [MLModelAssetResourceFactory.ModelStructureWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModelAssetResourceFactory) ModelStructure(ctx context.Context) error {
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

// ModelWithConfiguration is a synchronous wrapper around [MLModelAssetResourceFactory.ModelWithConfigurationCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModelAssetResourceFactory) ModelWithConfiguration(ctx context.Context, configuration objectivec.IObject) error {
	done := make(chan error, 1)
	m.ModelWithConfigurationCompletionHandler(configuration, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

