// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLShufflingBatchProvider] class.
var (
	_MLShufflingBatchProviderClass     MLShufflingBatchProviderClass
	_MLShufflingBatchProviderClassOnce sync.Once
)

func getMLShufflingBatchProviderClass() MLShufflingBatchProviderClass {
	_MLShufflingBatchProviderClassOnce.Do(func() {
		_MLShufflingBatchProviderClass = MLShufflingBatchProviderClass{class: objc.GetClass("MLShufflingBatchProvider")}
	})
	return _MLShufflingBatchProviderClass
}

// GetMLShufflingBatchProviderClass returns the class object for MLShufflingBatchProvider.
func GetMLShufflingBatchProviderClass() MLShufflingBatchProviderClass {
	return getMLShufflingBatchProviderClass()
}

type MLShufflingBatchProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLShufflingBatchProviderClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLShufflingBatchProviderClass) Alloc() MLShufflingBatchProvider {
	rv := objc.Send[MLShufflingBatchProvider](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLShufflingBatchProvider.BatchProvider]
//   - [MLShufflingBatchProvider.SetBatchProvider]
//   - [MLShufflingBatchProvider.Count]
//   - [MLShufflingBatchProvider.FeaturesAtIndex]
//   - [MLShufflingBatchProvider.Shuffle]
//   - [MLShufflingBatchProvider.InitWithBatchProviderSeed]
type MLShufflingBatchProvider struct {
	objectivec.Object
}

// MLShufflingBatchProviderFromID constructs a [MLShufflingBatchProvider] from an objc.ID.
func MLShufflingBatchProviderFromID(id objc.ID) MLShufflingBatchProvider {
	return MLShufflingBatchProvider{objectivec.Object{ID: id}}
}

// Ensure MLShufflingBatchProvider implements IMLShufflingBatchProvider.
var _ IMLShufflingBatchProvider = MLShufflingBatchProvider{}

// An interface definition for the [MLShufflingBatchProvider] class.
//
// # Methods
//
//   - [IMLShufflingBatchProvider.BatchProvider]
//   - [IMLShufflingBatchProvider.SetBatchProvider]
//   - [IMLShufflingBatchProvider.Count]
//   - [IMLShufflingBatchProvider.FeaturesAtIndex]
//   - [IMLShufflingBatchProvider.Shuffle]
//   - [IMLShufflingBatchProvider.InitWithBatchProviderSeed]
type IMLShufflingBatchProvider interface {
	objectivec.IObject

	// Topic: Methods

	BatchProvider() unsafe.Pointer
	SetBatchProvider(value unsafe.Pointer)
	Count() int64
	FeaturesAtIndex(index int64) objectivec.IObject
	Shuffle()
	InitWithBatchProviderSeed(provider objectivec.IObject, seed objectivec.IObject) MLShufflingBatchProvider
}

// Init initializes the instance.
func (m MLShufflingBatchProvider) Init() MLShufflingBatchProvider {
	rv := objc.Send[MLShufflingBatchProvider](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLShufflingBatchProvider) Autorelease() MLShufflingBatchProvider {
	rv := objc.Send[MLShufflingBatchProvider](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLShufflingBatchProvider creates a new MLShufflingBatchProvider instance.
func NewMLShufflingBatchProvider() MLShufflingBatchProvider {
	class := getMLShufflingBatchProviderClass()
	rv := objc.Send[MLShufflingBatchProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewShufflingBatchProviderWithBatchProviderSeed(provider objectivec.IObject, seed objectivec.IObject) MLShufflingBatchProvider {
	instance := getMLShufflingBatchProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBatchProvider:seed:"), provider, seed)
	return MLShufflingBatchProviderFromID(rv)
}

func (m MLShufflingBatchProvider) FeaturesAtIndex(index int64) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("featuresAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (m MLShufflingBatchProvider) Shuffle() {
	objc.Send[objc.ID](m.ID, objc.Sel("shuffle"))
}
func (m MLShufflingBatchProvider) InitWithBatchProviderSeed(provider objectivec.IObject, seed objectivec.IObject) MLShufflingBatchProvider {
	rv := objc.Send[MLShufflingBatchProvider](m.ID, objc.Sel("initWithBatchProvider:seed:"), provider, seed)
	return rv
}

func (m MLShufflingBatchProvider) BatchProvider() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("batchProvider"))
	return rv
}
func (m MLShufflingBatchProvider) SetBatchProvider(value unsafe.Pointer) {
	objc.Send[struct{}](m.ID, objc.Sel("setBatchProvider:"), value)
}
func (m MLShufflingBatchProvider) Count() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("count"))
	return rv
}
