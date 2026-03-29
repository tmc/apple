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

// The class instance for the [MLModelAssetModelVendor] class.
var (
	_MLModelAssetModelVendorClass     MLModelAssetModelVendorClass
	_MLModelAssetModelVendorClassOnce sync.Once
)

func getMLModelAssetModelVendorClass() MLModelAssetModelVendorClass {
	_MLModelAssetModelVendorClassOnce.Do(func() {
		_MLModelAssetModelVendorClass = MLModelAssetModelVendorClass{class: objc.GetClass("MLModelAssetModelVendor")}
	})
	return _MLModelAssetModelVendorClass
}

// GetMLModelAssetModelVendorClass returns the class object for MLModelAssetModelVendor.
func GetMLModelAssetModelVendorClass() MLModelAssetModelVendorClass {
	return getMLModelAssetModelVendorClass()
}

type MLModelAssetModelVendorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelAssetModelVendorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelAssetModelVendorClass) Alloc() MLModelAssetModelVendor {
	rv := objc.Send[MLModelAssetModelVendor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLModelAssetModelVendor.CachedConfiguration]
//   - [MLModelAssetModelVendor.SetCachedConfiguration]
//   - [MLModelAssetModelVendor.CachedModel]
//   - [MLModelAssetModelVendor.SetCachedModel]
//   - [MLModelAssetModelVendor.ModelWithConfigurationCompletionHandler]
//   - [MLModelAssetModelVendor.ModelWithConfigurationError]
//   - [MLModelAssetModelVendor.ResourceFactory]
//   - [MLModelAssetModelVendor.InitWithResourceFactory]
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetModelVendor
type MLModelAssetModelVendor struct {
	objectivec.Object
}

// MLModelAssetModelVendorFromID constructs a [MLModelAssetModelVendor] from an objc.ID.
func MLModelAssetModelVendorFromID(id objc.ID) MLModelAssetModelVendor {
	return MLModelAssetModelVendor{objectivec.Object{ID: id}}
}
// Ensure MLModelAssetModelVendor implements IMLModelAssetModelVendor.
var _ IMLModelAssetModelVendor = MLModelAssetModelVendor{}

// An interface definition for the [MLModelAssetModelVendor] class.
//
// # Methods
//
//   - [IMLModelAssetModelVendor.CachedConfiguration]
//   - [IMLModelAssetModelVendor.SetCachedConfiguration]
//   - [IMLModelAssetModelVendor.CachedModel]
//   - [IMLModelAssetModelVendor.SetCachedModel]
//   - [IMLModelAssetModelVendor.ModelWithConfigurationCompletionHandler]
//   - [IMLModelAssetModelVendor.ModelWithConfigurationError]
//   - [IMLModelAssetModelVendor.ResourceFactory]
//   - [IMLModelAssetModelVendor.InitWithResourceFactory]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetModelVendor
type IMLModelAssetModelVendor interface {
	objectivec.IObject

	// Topic: Methods

	CachedConfiguration() IMLModelConfiguration
	SetCachedConfiguration(value IMLModelConfiguration)
	CachedModel() IMLModel
	SetCachedModel(value IMLModel)
	ModelWithConfigurationCompletionHandler(configuration objectivec.IObject, handler ErrorHandler)
	ModelWithConfigurationError(configuration objectivec.IObject) (objectivec.IObject, error)
	ResourceFactory() IMLModelAssetResourceFactory
	InitWithResourceFactory(factory objectivec.IObject) MLModelAssetModelVendor
}

// Init initializes the instance.
func (m MLModelAssetModelVendor) Init() MLModelAssetModelVendor {
	rv := objc.Send[MLModelAssetModelVendor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelAssetModelVendor) Autorelease() MLModelAssetModelVendor {
	rv := objc.Send[MLModelAssetModelVendor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelAssetModelVendor creates a new MLModelAssetModelVendor instance.
func NewMLModelAssetModelVendor() MLModelAssetModelVendor {
	class := getMLModelAssetModelVendorClass()
	rv := objc.Send[MLModelAssetModelVendor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetModelVendor/initWithResourceFactory:
func NewModelAssetModelVendorWithResourceFactory(factory objectivec.IObject) MLModelAssetModelVendor {
	instance := getMLModelAssetModelVendorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResourceFactory:"), factory)
	return MLModelAssetModelVendorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetModelVendor/modelWithConfiguration:completionHandler:
func (m MLModelAssetModelVendor) ModelWithConfigurationCompletionHandler(configuration objectivec.IObject, handler ErrorHandler) {
_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("modelWithConfiguration:completionHandler:"), configuration, _block1)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetModelVendor/modelWithConfiguration:error:
func (m MLModelAssetModelVendor) ModelWithConfigurationError(configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetModelVendor/initWithResourceFactory:
func (m MLModelAssetModelVendor) InitWithResourceFactory(factory objectivec.IObject) MLModelAssetModelVendor {
	rv := objc.Send[MLModelAssetModelVendor](m.ID, objc.Sel("initWithResourceFactory:"), factory)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelAssetModelVendor/cachedConfiguration
func (m MLModelAssetModelVendor) CachedConfiguration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("cachedConfiguration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}
func (m MLModelAssetModelVendor) SetCachedConfiguration(value IMLModelConfiguration) {
	objc.Send[struct{}](m.ID, objc.Sel("setCachedConfiguration:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetModelVendor/cachedModel
func (m MLModelAssetModelVendor) CachedModel() IMLModel {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("cachedModel"))
	return MLModelFromID(objc.ID(rv))
}
func (m MLModelAssetModelVendor) SetCachedModel(value IMLModel) {
	objc.Send[struct{}](m.ID, objc.Sel("setCachedModel:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLModelAssetModelVendor/resourceFactory
func (m MLModelAssetModelVendor) ResourceFactory() IMLModelAssetResourceFactory {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("resourceFactory"))
	return MLModelAssetResourceFactoryFromID(objc.ID(rv))
}

// ModelWithConfiguration is a synchronous wrapper around [MLModelAssetModelVendor.ModelWithConfigurationCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModelAssetModelVendor) ModelWithConfiguration(ctx context.Context, configuration objectivec.IObject) error {
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

