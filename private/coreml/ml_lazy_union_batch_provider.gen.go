// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLLazyUnionBatchProvider] class.
var (
	_MLLazyUnionBatchProviderClass     MLLazyUnionBatchProviderClass
	_MLLazyUnionBatchProviderClassOnce sync.Once
)

func getMLLazyUnionBatchProviderClass() MLLazyUnionBatchProviderClass {
	_MLLazyUnionBatchProviderClassOnce.Do(func() {
		_MLLazyUnionBatchProviderClass = MLLazyUnionBatchProviderClass{class: objc.GetClass("MLLazyUnionBatchProvider")}
	})
	return _MLLazyUnionBatchProviderClass
}

// GetMLLazyUnionBatchProviderClass returns the class object for MLLazyUnionBatchProvider.
func GetMLLazyUnionBatchProviderClass() MLLazyUnionBatchProviderClass {
	return getMLLazyUnionBatchProviderClass()
}

type MLLazyUnionBatchProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLLazyUnionBatchProviderClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLLazyUnionBatchProviderClass) Alloc() MLLazyUnionBatchProvider {
	rv := objc.Send[MLLazyUnionBatchProvider](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLLazyUnionBatchProvider.Count]
//   - [MLLazyUnionBatchProvider.FeaturesAtIndex]
//   - [MLLazyUnionBatchProvider.First]
//   - [MLLazyUnionBatchProvider.SetFirst]
//   - [MLLazyUnionBatchProvider.Second]
//   - [MLLazyUnionBatchProvider.SetSecond]
//   - [MLLazyUnionBatchProvider.InitWithFeaturesFromAddedToFeaturesFromError]
// See: https://developer.apple.com/documentation/CoreML/MLLazyUnionBatchProvider
type MLLazyUnionBatchProvider struct {
	objectivec.Object
}

// MLLazyUnionBatchProviderFromID constructs a [MLLazyUnionBatchProvider] from an objc.ID.
func MLLazyUnionBatchProviderFromID(id objc.ID) MLLazyUnionBatchProvider {
	return MLLazyUnionBatchProvider{objectivec.Object{ID: id}}
}
// Ensure MLLazyUnionBatchProvider implements IMLLazyUnionBatchProvider.
var _ IMLLazyUnionBatchProvider = MLLazyUnionBatchProvider{}

// An interface definition for the [MLLazyUnionBatchProvider] class.
//
// # Methods
//
//   - [IMLLazyUnionBatchProvider.Count]
//   - [IMLLazyUnionBatchProvider.FeaturesAtIndex]
//   - [IMLLazyUnionBatchProvider.First]
//   - [IMLLazyUnionBatchProvider.SetFirst]
//   - [IMLLazyUnionBatchProvider.Second]
//   - [IMLLazyUnionBatchProvider.SetSecond]
//   - [IMLLazyUnionBatchProvider.InitWithFeaturesFromAddedToFeaturesFromError]
//
// See: https://developer.apple.com/documentation/CoreML/MLLazyUnionBatchProvider
type IMLLazyUnionBatchProvider interface {
	objectivec.IObject

	// Topic: Methods

	Count() int64
	FeaturesAtIndex(index int64) objectivec.IObject
	First() objectivec.IObject
	SetFirst(value objectivec.IObject)
	Second() objectivec.IObject
	SetSecond(value objectivec.IObject)
	InitWithFeaturesFromAddedToFeaturesFromError(from objectivec.IObject, from2 objectivec.IObject) (MLLazyUnionBatchProvider, error)
}

// Init initializes the instance.
func (l MLLazyUnionBatchProvider) Init() MLLazyUnionBatchProvider {
	rv := objc.Send[MLLazyUnionBatchProvider](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MLLazyUnionBatchProvider) Autorelease() MLLazyUnionBatchProvider {
	rv := objc.Send[MLLazyUnionBatchProvider](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLLazyUnionBatchProvider creates a new MLLazyUnionBatchProvider instance.
func NewMLLazyUnionBatchProvider() MLLazyUnionBatchProvider {
	class := getMLLazyUnionBatchProviderClass()
	rv := objc.Send[MLLazyUnionBatchProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLLazyUnionBatchProvider/initWithFeaturesFrom:addedToFeaturesFrom:error:
func NewLazyUnionBatchProviderWithFeaturesFromAddedToFeaturesFromError(from objectivec.IObject, from2 objectivec.IObject) (MLLazyUnionBatchProvider, error) {
	var errorPtr objc.ID
	instance := getMLLazyUnionBatchProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFeaturesFrom:addedToFeaturesFrom:error:"), from, from2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLLazyUnionBatchProvider{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLLazyUnionBatchProviderFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLLazyUnionBatchProvider/featuresAtIndex:
func (l MLLazyUnionBatchProvider) FeaturesAtIndex(index int64) objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("featuresAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLLazyUnionBatchProvider/initWithFeaturesFrom:addedToFeaturesFrom:error:
func (l MLLazyUnionBatchProvider) InitWithFeaturesFromAddedToFeaturesFromError(from objectivec.IObject, from2 objectivec.IObject) (MLLazyUnionBatchProvider, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](l.ID, objc.Sel("initWithFeaturesFrom:addedToFeaturesFrom:error:"), from, from2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLLazyUnionBatchProvider{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLLazyUnionBatchProviderFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLLazyUnionBatchProvider/count
func (l MLLazyUnionBatchProvider) Count() int64 {
	rv := objc.Send[int64](l.ID, objc.Sel("count"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLLazyUnionBatchProvider/first
func (l MLLazyUnionBatchProvider) First() objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("first"))
	return objectivec.Object{ID: rv}
}
func (l MLLazyUnionBatchProvider) SetFirst(value objectivec.IObject) {
	objc.Send[struct{}](l.ID, objc.Sel("setFirst:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLLazyUnionBatchProvider/second
func (l MLLazyUnionBatchProvider) Second() objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("second"))
	return objectivec.Object{ID: rv}
}
func (l MLLazyUnionBatchProvider) SetSecond(value objectivec.IObject) {
	objc.Send[struct{}](l.ID, objc.Sel("setSecond:"), value)
}

