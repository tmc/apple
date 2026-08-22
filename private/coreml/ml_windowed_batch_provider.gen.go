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
	rv := objc.SendIfResponds[MLWindowedBatchProvider](objc.ID(mc.class), objc.Sel("alloc"))
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
type IMLWindowedBatchProvider interface {
	objectivec.IObject

	// Topic: Methods

	Count() int64
	FeaturesAtIndex(index int64) objectivec.IObject
	FullBatch() unsafe.Pointer
	SetFullBatch(value unsafe.Pointer)
	StartIndex() int64
	SetStartIndex(value int64)
	WindowLength() int64
	SetWindowLength(value int64)
	InitWithBatchStartIndexWindowLengthError(batch objectivec.IObject, index int64, length int64) (MLWindowedBatchProvider, error)
}

// Init initializes the instance.
func (m MLWindowedBatchProvider) Init() MLWindowedBatchProvider {
	rv := objc.SendIfResponds[MLWindowedBatchProvider](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLWindowedBatchProvider) Autorelease() MLWindowedBatchProvider {
	rv := objc.SendIfResponds[MLWindowedBatchProvider](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLWindowedBatchProvider creates a new MLWindowedBatchProvider instance.
func NewMLWindowedBatchProvider() MLWindowedBatchProvider {
	class := getMLWindowedBatchProviderClass()
	rv := objc.SendIfResponds[MLWindowedBatchProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewWindowedBatchProviderWithBatchStartIndexWindowLengthError(batch objectivec.IObject, index int64, length int64) (MLWindowedBatchProvider, error) {
	var errorPtr objc.ID
	instance := getMLWindowedBatchProviderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBatch:startIndex:windowLength:error:"), batch, index, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLWindowedBatchProvider{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLWindowedBatchProvider{}, objc.ErrInitFailed
	}
	return MLWindowedBatchProviderFromID(rv), nil
}

func (m MLWindowedBatchProvider) FeaturesAtIndex(index int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("featuresAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (m MLWindowedBatchProvider) InitWithBatchStartIndexWindowLengthError(batch objectivec.IObject, index int64, length int64) (MLWindowedBatchProvider, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithBatch:startIndex:windowLength:error:"), batch, index, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLWindowedBatchProvider{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLWindowedBatchProviderFromID(rv), nil

}

func (m MLWindowedBatchProvider) Count() int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("count"))
	return rv
}
func (m MLWindowedBatchProvider) FullBatch() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("fullBatch"))
	return rv
}
func (m MLWindowedBatchProvider) SetFullBatch(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setFullBatch:"), value)
}
func (m MLWindowedBatchProvider) StartIndex() int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("startIndex"))
	return rv
}
func (m MLWindowedBatchProvider) SetStartIndex(value int64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setStartIndex:"), value)
}
func (m MLWindowedBatchProvider) WindowLength() int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("windowLength"))
	return rv
}
func (m MLWindowedBatchProvider) SetWindowLength(value int64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setWindowLength:"), value)
}
