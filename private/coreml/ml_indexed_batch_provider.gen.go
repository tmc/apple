// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLIndexedBatchProvider] class.
var (
	_MLIndexedBatchProviderClass     MLIndexedBatchProviderClass
	_MLIndexedBatchProviderClassOnce sync.Once
)

func getMLIndexedBatchProviderClass() MLIndexedBatchProviderClass {
	_MLIndexedBatchProviderClassOnce.Do(func() {
		_MLIndexedBatchProviderClass = MLIndexedBatchProviderClass{class: objc.GetClass("MLIndexedBatchProvider")}
	})
	return _MLIndexedBatchProviderClass
}

// GetMLIndexedBatchProviderClass returns the class object for MLIndexedBatchProvider.
func GetMLIndexedBatchProviderClass() MLIndexedBatchProviderClass {
	return getMLIndexedBatchProviderClass()
}

type MLIndexedBatchProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLIndexedBatchProviderClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLIndexedBatchProviderClass) Alloc() MLIndexedBatchProvider {
	rv := objc.SendIfResponds[MLIndexedBatchProvider](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLIndexedBatchProvider.Count]
//   - [MLIndexedBatchProvider.FeaturesAtIndex]
//   - [MLIndexedBatchProvider.FullBatch]
//   - [MLIndexedBatchProvider.SetFullBatch]
//   - [MLIndexedBatchProvider.Indices]
//   - [MLIndexedBatchProvider.SetIndices]
//   - [MLIndexedBatchProvider.InitWithBatchIndicesError]
type MLIndexedBatchProvider struct {
	objectivec.Object
}

// MLIndexedBatchProviderFromID constructs a [MLIndexedBatchProvider] from an objc.ID.
func MLIndexedBatchProviderFromID(id objc.ID) MLIndexedBatchProvider {
	return MLIndexedBatchProvider{objectivec.Object{ID: id}}
}

// Ensure MLIndexedBatchProvider implements IMLIndexedBatchProvider.
var _ IMLIndexedBatchProvider = MLIndexedBatchProvider{}

// An interface definition for the [MLIndexedBatchProvider] class.
//
// # Methods
//
//   - [IMLIndexedBatchProvider.Count]
//   - [IMLIndexedBatchProvider.FeaturesAtIndex]
//   - [IMLIndexedBatchProvider.FullBatch]
//   - [IMLIndexedBatchProvider.SetFullBatch]
//   - [IMLIndexedBatchProvider.Indices]
//   - [IMLIndexedBatchProvider.SetIndices]
//   - [IMLIndexedBatchProvider.InitWithBatchIndicesError]
type IMLIndexedBatchProvider interface {
	objectivec.IObject

	// Topic: Methods

	Count() int64
	FeaturesAtIndex(index int64) objectivec.IObject
	FullBatch() unsafe.Pointer
	SetFullBatch(value unsafe.Pointer)
	Indices() foundation.INSArray
	SetIndices(value foundation.INSArray)
	InitWithBatchIndicesError(batch objectivec.IObject, indices objectivec.IObject) (MLIndexedBatchProvider, error)
}

// Init initializes the instance.
func (m MLIndexedBatchProvider) Init() MLIndexedBatchProvider {
	rv := objc.SendIfResponds[MLIndexedBatchProvider](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLIndexedBatchProvider) Autorelease() MLIndexedBatchProvider {
	rv := objc.SendIfResponds[MLIndexedBatchProvider](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLIndexedBatchProvider creates a new MLIndexedBatchProvider instance.
func NewMLIndexedBatchProvider() MLIndexedBatchProvider {
	class := getMLIndexedBatchProviderClass()
	rv := objc.SendIfResponds[MLIndexedBatchProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewIndexedBatchProviderWithBatchIndicesError(batch objectivec.IObject, indices objectivec.IObject) (MLIndexedBatchProvider, error) {
	var errorPtr objc.ID
	instance := getMLIndexedBatchProviderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBatch:indices:error:"), batch, indices, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLIndexedBatchProvider{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLIndexedBatchProvider{}, objc.ErrInitFailed
	}
	return MLIndexedBatchProviderFromID(rv), nil
}

func (m MLIndexedBatchProvider) FeaturesAtIndex(index int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("featuresAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (m MLIndexedBatchProvider) InitWithBatchIndicesError(batch objectivec.IObject, indices objectivec.IObject) (MLIndexedBatchProvider, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithBatch:indices:error:"), batch, indices, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLIndexedBatchProvider{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLIndexedBatchProviderFromID(rv), nil

}

func (m MLIndexedBatchProvider) Count() int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("count"))
	return rv
}
func (m MLIndexedBatchProvider) FullBatch() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("fullBatch"))
	return rv
}
func (m MLIndexedBatchProvider) SetFullBatch(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setFullBatch:"), value)
}
func (m MLIndexedBatchProvider) Indices() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("indices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLIndexedBatchProvider) SetIndices(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setIndices:"), value)
}
