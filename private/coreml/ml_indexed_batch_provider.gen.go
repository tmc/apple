// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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
	rv := objc.Send[MLIndexedBatchProvider](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLIndexedBatchProvider.Count]
//   - [MLIndexedBatchProvider.FeaturesAtIndex]
//   - [MLIndexedBatchProvider.FullBatch]
//   - [MLIndexedBatchProvider.SetFullBatch]
//   - [MLIndexedBatchProvider.Indices]
//   - [MLIndexedBatchProvider.SetIndices]
//   - [MLIndexedBatchProvider.InitWithBatchIndicesError]
// See: https://developer.apple.com/documentation/CoreML/MLIndexedBatchProvider
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
//
// See: https://developer.apple.com/documentation/CoreML/MLIndexedBatchProvider
type IMLIndexedBatchProvider interface {
	objectivec.IObject

	// Topic: Methods

	Count() int64
	FeaturesAtIndex(index int64) objectivec.IObject
	FullBatch() objectivec.IObject
	SetFullBatch(value objectivec.IObject)
	Indices() foundation.INSArray
	SetIndices(value foundation.INSArray)
	InitWithBatchIndicesError(batch objectivec.IObject, indices objectivec.IObject) (MLIndexedBatchProvider, error)
}

// Init initializes the instance.
func (i MLIndexedBatchProvider) Init() MLIndexedBatchProvider {
	rv := objc.Send[MLIndexedBatchProvider](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MLIndexedBatchProvider) Autorelease() MLIndexedBatchProvider {
	rv := objc.Send[MLIndexedBatchProvider](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLIndexedBatchProvider creates a new MLIndexedBatchProvider instance.
func NewMLIndexedBatchProvider() MLIndexedBatchProvider {
	class := getMLIndexedBatchProviderClass()
	rv := objc.Send[MLIndexedBatchProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLIndexedBatchProvider/initWithBatch:indices:error:
func NewIndexedBatchProviderWithBatchIndicesError(batch objectivec.IObject, indices objectivec.IObject) (MLIndexedBatchProvider, error) {
	var errorPtr objc.ID
	instance := getMLIndexedBatchProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBatch:indices:error:"), batch, indices, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLIndexedBatchProvider{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLIndexedBatchProviderFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLIndexedBatchProvider/featuresAtIndex:
func (i MLIndexedBatchProvider) FeaturesAtIndex(index int64) objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("featuresAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLIndexedBatchProvider/initWithBatch:indices:error:
func (i MLIndexedBatchProvider) InitWithBatchIndicesError(batch objectivec.IObject, indices objectivec.IObject) (MLIndexedBatchProvider, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](i.ID, objc.Sel("initWithBatch:indices:error:"), batch, indices, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLIndexedBatchProvider{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLIndexedBatchProviderFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLIndexedBatchProvider/count
func (i MLIndexedBatchProvider) Count() int64 {
	rv := objc.Send[int64](i.ID, objc.Sel("count"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLIndexedBatchProvider/fullBatch
func (i MLIndexedBatchProvider) FullBatch() objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("fullBatch"))
	return objectivec.Object{ID: rv}
}
func (i MLIndexedBatchProvider) SetFullBatch(value objectivec.IObject) {
	objc.Send[struct{}](i.ID, objc.Sel("setFullBatch:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLIndexedBatchProvider/indices
func (i MLIndexedBatchProvider) Indices() foundation.INSArray {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("indices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (i MLIndexedBatchProvider) SetIndices(value foundation.INSArray) {
	objc.Send[struct{}](i.ID, objc.Sel("setIndices:"), value)
}

