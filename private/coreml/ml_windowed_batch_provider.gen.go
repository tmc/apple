// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLWindowedBatchProvider] class.
var (
	_MLWindowedBatchProviderClass     MLWindowedBatchProviderClass
	_MLWindowedBatchProviderClassOnce sync.Once
)

func getMLWindowedBatchProviderClass() MLWindowedBatchProviderClass {
	_MLWindowedBatchProviderClassOnce.Do(func() {
		_MLWindowedBatchProviderClass = MLWindowedBatchProviderClass{class: objc.GetClass("MLWindowedBatchProvider")}
	})
	return _MLWindowedBatchProviderClass
}

// GetMLWindowedBatchProviderClass returns the class object for MLWindowedBatchProvider.
func GetMLWindowedBatchProviderClass() MLWindowedBatchProviderClass {
	return getMLWindowedBatchProviderClass()
}

type MLWindowedBatchProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLWindowedBatchProviderClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLWindowedBatchProviderClass) Alloc() MLWindowedBatchProvider {
	rv := objc.Send[MLWindowedBatchProvider](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLWindowedBatchProvider.Count]
//   - [MLWindowedBatchProvider.FeaturesAtIndex]
//   - [MLWindowedBatchProvider.FullBatch]
//   - [MLWindowedBatchProvider.SetFullBatch]
//   - [MLWindowedBatchProvider.StartIndex]
//   - [MLWindowedBatchProvider.SetStartIndex]
//   - [MLWindowedBatchProvider.WindowLength]
//   - [MLWindowedBatchProvider.SetWindowLength]
//   - [MLWindowedBatchProvider.InitWithBatchStartIndexWindowLengthError]
//
// See: https://developer.apple.com/documentation/CoreML/MLWindowedBatchProvider
type MLWindowedBatchProvider struct {
	objectivec.Object
}

// MLWindowedBatchProviderFromID constructs a [MLWindowedBatchProvider] from an objc.ID.
func MLWindowedBatchProviderFromID(id objc.ID) MLWindowedBatchProvider {
	return MLWindowedBatchProvider{objectivec.Object{ID: id}}
}

// Ensure MLWindowedBatchProvider implements IMLWindowedBatchProvider.
var _ IMLWindowedBatchProvider = MLWindowedBatchProvider{}

// An interface definition for the [MLWindowedBatchProvider] class.
//
// # Methods
//
//   - [IMLWindowedBatchProvider.Count]
//   - [IMLWindowedBatchProvider.FeaturesAtIndex]
//   - [IMLWindowedBatchProvider.FullBatch]
//   - [IMLWindowedBatchProvider.SetFullBatch]
//   - [IMLWindowedBatchProvider.StartIndex]
//   - [IMLWindowedBatchProvider.SetStartIndex]
//   - [IMLWindowedBatchProvider.WindowLength]
//   - [IMLWindowedBatchProvider.SetWindowLength]
//   - [IMLWindowedBatchProvider.InitWithBatchStartIndexWindowLengthError]
//
// See: https://developer.apple.com/documentation/CoreML/MLWindowedBatchProvider
type IMLWindowedBatchProvider interface {
	objectivec.IObject

	// Topic: Methods

	Count() int64
	FeaturesAtIndex(index int64) objectivec.IObject
	FullBatch() objectivec.IObject
	SetFullBatch(value objectivec.IObject)
	StartIndex() int64
	SetStartIndex(value int64)
	WindowLength() int64
	SetWindowLength(value int64)
	InitWithBatchStartIndexWindowLengthError(batch objectivec.IObject, index int64, length int64) (MLWindowedBatchProvider, error)
}

// Init initializes the instance.
func (w MLWindowedBatchProvider) Init() MLWindowedBatchProvider {
	rv := objc.Send[MLWindowedBatchProvider](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w MLWindowedBatchProvider) Autorelease() MLWindowedBatchProvider {
	rv := objc.Send[MLWindowedBatchProvider](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLWindowedBatchProvider creates a new MLWindowedBatchProvider instance.
func NewMLWindowedBatchProvider() MLWindowedBatchProvider {
	class := getMLWindowedBatchProviderClass()
	rv := objc.Send[MLWindowedBatchProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLWindowedBatchProvider/initWithBatch:startIndex:windowLength:error:
func NewWindowedBatchProviderWithBatchStartIndexWindowLengthError(batch objectivec.IObject, index int64, length int64) (MLWindowedBatchProvider, error) {
	var errorPtr objc.ID
	instance := getMLWindowedBatchProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBatch:startIndex:windowLength:error:"), batch, index, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLWindowedBatchProvider{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLWindowedBatchProviderFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLWindowedBatchProvider/featuresAtIndex:
func (w MLWindowedBatchProvider) FeaturesAtIndex(index int64) objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("featuresAtIndex:"), index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLWindowedBatchProvider/initWithBatch:startIndex:windowLength:error:
func (w MLWindowedBatchProvider) InitWithBatchStartIndexWindowLengthError(batch objectivec.IObject, index int64, length int64) (MLWindowedBatchProvider, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](w.ID, objc.Sel("initWithBatch:startIndex:windowLength:error:"), batch, index, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLWindowedBatchProvider{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLWindowedBatchProviderFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLWindowedBatchProvider/count
func (w MLWindowedBatchProvider) Count() int64 {
	rv := objc.Send[int64](w.ID, objc.Sel("count"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLWindowedBatchProvider/fullBatch
func (w MLWindowedBatchProvider) FullBatch() objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("fullBatch"))
	return objectivec.Object{ID: rv}
}
func (w MLWindowedBatchProvider) SetFullBatch(value objectivec.IObject) {
	objc.Send[struct{}](w.ID, objc.Sel("setFullBatch:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLWindowedBatchProvider/startIndex
func (w MLWindowedBatchProvider) StartIndex() int64 {
	rv := objc.Send[int64](w.ID, objc.Sel("startIndex"))
	return rv
}
func (w MLWindowedBatchProvider) SetStartIndex(value int64) {
	objc.Send[struct{}](w.ID, objc.Sel("setStartIndex:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLWindowedBatchProvider/windowLength
func (w MLWindowedBatchProvider) WindowLength() int64 {
	rv := objc.Send[int64](w.ID, objc.Sel("windowLength"))
	return rv
}
func (w MLWindowedBatchProvider) SetWindowLength(value int64) {
	objc.Send[struct{}](w.ID, objc.Sel("setWindowLength:"), value)
}
