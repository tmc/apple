// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[MLLazyUnionBatchProvider](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLLazyUnionBatchProvider.Count]
//   - [MLLazyUnionBatchProvider.FeaturesAtIndex]
//   - [MLLazyUnionBatchProvider.First]
//   - [MLLazyUnionBatchProvider.SetFirst]
//   - [MLLazyUnionBatchProvider.Second]
//   - [MLLazyUnionBatchProvider.SetSecond]
//   - [MLLazyUnionBatchProvider.InitWithFeaturesFromAddedToFeaturesFromError]
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
type IMLLazyUnionBatchProvider interface {
	objectivec.IObject

	// Topic: Methods

	Count() int64
	FeaturesAtIndex(index int64) objectivec.IObject
	First() unsafe.Pointer
	SetFirst(value unsafe.Pointer)
	Second() unsafe.Pointer
	SetSecond(value unsafe.Pointer)
	InitWithFeaturesFromAddedToFeaturesFromError(from objectivec.IObject, from2 objectivec.IObject) (MLLazyUnionBatchProvider, error)
}

// Init initializes the instance.
func (m MLLazyUnionBatchProvider) Init() MLLazyUnionBatchProvider {
	rv := objc.SendIfResponds[MLLazyUnionBatchProvider](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLLazyUnionBatchProvider) Autorelease() MLLazyUnionBatchProvider {
	rv := objc.SendIfResponds[MLLazyUnionBatchProvider](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLLazyUnionBatchProvider creates a new MLLazyUnionBatchProvider instance.
func NewMLLazyUnionBatchProvider() MLLazyUnionBatchProvider {
	class := getMLLazyUnionBatchProviderClass()
	rv := objc.SendIfResponds[MLLazyUnionBatchProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewLazyUnionBatchProviderWithFeaturesFromAddedToFeaturesFromError(from objectivec.IObject, from2 objectivec.IObject) (MLLazyUnionBatchProvider, error) {
	var errorPtr objc.ID
	instance := getMLLazyUnionBatchProviderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithFeaturesFrom:addedToFeaturesFrom:error:"), from, from2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLLazyUnionBatchProvider{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLLazyUnionBatchProvider{}, objc.ErrInitFailed
	}
	return MLLazyUnionBatchProviderFromID(rv), nil
}

func (m MLLazyUnionBatchProvider) FeaturesAtIndex(index int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("featuresAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (m MLLazyUnionBatchProvider) InitWithFeaturesFromAddedToFeaturesFromError(from objectivec.IObject, from2 objectivec.IObject) (MLLazyUnionBatchProvider, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithFeaturesFrom:addedToFeaturesFrom:error:"), from, from2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLLazyUnionBatchProvider{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLLazyUnionBatchProviderFromID(rv), nil

}

func (m MLLazyUnionBatchProvider) Count() int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("count"))
	return rv
}
func (m MLLazyUnionBatchProvider) First() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("first"))
	return rv
}
func (m MLLazyUnionBatchProvider) SetFirst(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setFirst:"), value)
}
func (m MLLazyUnionBatchProvider) Second() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("second"))
	return rv
}
func (m MLLazyUnionBatchProvider) SetSecond(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setSecond:"), value)
}
