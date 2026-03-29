// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLProbabilityDictionaryMultiArrayStorage] class.
var (
	_MLProbabilityDictionaryMultiArrayStorageClass     MLProbabilityDictionaryMultiArrayStorageClass
	_MLProbabilityDictionaryMultiArrayStorageClassOnce sync.Once
)

func getMLProbabilityDictionaryMultiArrayStorageClass() MLProbabilityDictionaryMultiArrayStorageClass {
	_MLProbabilityDictionaryMultiArrayStorageClassOnce.Do(func() {
		_MLProbabilityDictionaryMultiArrayStorageClass = MLProbabilityDictionaryMultiArrayStorageClass{class: objc.GetClass("MLProbabilityDictionaryMultiArrayStorage")}
	})
	return _MLProbabilityDictionaryMultiArrayStorageClass
}

// GetMLProbabilityDictionaryMultiArrayStorageClass returns the class object for MLProbabilityDictionaryMultiArrayStorage.
func GetMLProbabilityDictionaryMultiArrayStorageClass() MLProbabilityDictionaryMultiArrayStorageClass {
	return getMLProbabilityDictionaryMultiArrayStorageClass()
}

type MLProbabilityDictionaryMultiArrayStorageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLProbabilityDictionaryMultiArrayStorageClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLProbabilityDictionaryMultiArrayStorageClass) Alloc() MLProbabilityDictionaryMultiArrayStorage {
	rv := objc.Send[MLProbabilityDictionaryMultiArrayStorage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLProbabilityDictionaryMultiArrayStorage.Count]
//   - [MLProbabilityDictionaryMultiArrayStorage.MaxElementIndex]
//   - [MLProbabilityDictionaryMultiArrayStorage.ProbabilityAtIndex]
//   - [MLProbabilityDictionaryMultiArrayStorage.InitWithMultiArray]
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionaryMultiArrayStorage
type MLProbabilityDictionaryMultiArrayStorage struct {
	objectivec.Object
}

// MLProbabilityDictionaryMultiArrayStorageFromID constructs a [MLProbabilityDictionaryMultiArrayStorage] from an objc.ID.
func MLProbabilityDictionaryMultiArrayStorageFromID(id objc.ID) MLProbabilityDictionaryMultiArrayStorage {
	return MLProbabilityDictionaryMultiArrayStorage{objectivec.Object{ID: id}}
}
// Ensure MLProbabilityDictionaryMultiArrayStorage implements IMLProbabilityDictionaryMultiArrayStorage.
var _ IMLProbabilityDictionaryMultiArrayStorage = MLProbabilityDictionaryMultiArrayStorage{}

// An interface definition for the [MLProbabilityDictionaryMultiArrayStorage] class.
//
// # Methods
//
//   - [IMLProbabilityDictionaryMultiArrayStorage.Count]
//   - [IMLProbabilityDictionaryMultiArrayStorage.MaxElementIndex]
//   - [IMLProbabilityDictionaryMultiArrayStorage.ProbabilityAtIndex]
//   - [IMLProbabilityDictionaryMultiArrayStorage.InitWithMultiArray]
//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionaryMultiArrayStorage
type IMLProbabilityDictionaryMultiArrayStorage interface {
	objectivec.IObject

	// Topic: Methods

	Count() uint64
	MaxElementIndex() uint64
	ProbabilityAtIndex(index uint64) objectivec.IObject
	InitWithMultiArray(array objectivec.IObject) MLProbabilityDictionaryMultiArrayStorage
}

// Init initializes the instance.
func (p MLProbabilityDictionaryMultiArrayStorage) Init() MLProbabilityDictionaryMultiArrayStorage {
	rv := objc.Send[MLProbabilityDictionaryMultiArrayStorage](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLProbabilityDictionaryMultiArrayStorage) Autorelease() MLProbabilityDictionaryMultiArrayStorage {
	rv := objc.Send[MLProbabilityDictionaryMultiArrayStorage](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLProbabilityDictionaryMultiArrayStorage creates a new MLProbabilityDictionaryMultiArrayStorage instance.
func NewMLProbabilityDictionaryMultiArrayStorage() MLProbabilityDictionaryMultiArrayStorage {
	class := getMLProbabilityDictionaryMultiArrayStorageClass()
	rv := objc.Send[MLProbabilityDictionaryMultiArrayStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionaryMultiArrayStorage/initWithMultiArray:
func NewProbabilityDictionaryMultiArrayStorageWithMultiArray(array objectivec.IObject) MLProbabilityDictionaryMultiArrayStorage {
	instance := getMLProbabilityDictionaryMultiArrayStorageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMultiArray:"), array)
	return MLProbabilityDictionaryMultiArrayStorageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionaryMultiArrayStorage/count
func (p MLProbabilityDictionaryMultiArrayStorage) Count() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("count"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionaryMultiArrayStorage/maxElementIndex
func (p MLProbabilityDictionaryMultiArrayStorage) MaxElementIndex() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("maxElementIndex"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionaryMultiArrayStorage/probabilityAtIndex:
func (p MLProbabilityDictionaryMultiArrayStorage) ProbabilityAtIndex(index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("probabilityAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionaryMultiArrayStorage/initWithMultiArray:
func (p MLProbabilityDictionaryMultiArrayStorage) InitWithMultiArray(array objectivec.IObject) MLProbabilityDictionaryMultiArrayStorage {
	rv := objc.Send[MLProbabilityDictionaryMultiArrayStorage](p.ID, objc.Sel("initWithMultiArray:"), array)
	return rv
}

