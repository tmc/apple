// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLProbabilityDictionaryArrayStorage] class.
var (
	_MLProbabilityDictionaryArrayStorageClass     MLProbabilityDictionaryArrayStorageClass
	_MLProbabilityDictionaryArrayStorageClassOnce sync.Once
)

func getMLProbabilityDictionaryArrayStorageClass() MLProbabilityDictionaryArrayStorageClass {
	_MLProbabilityDictionaryArrayStorageClassOnce.Do(func() {
		_MLProbabilityDictionaryArrayStorageClass = MLProbabilityDictionaryArrayStorageClass{class: objc.GetClass("MLProbabilityDictionaryArrayStorage")}
	})
	return _MLProbabilityDictionaryArrayStorageClass
}

// GetMLProbabilityDictionaryArrayStorageClass returns the class object for MLProbabilityDictionaryArrayStorage.
func GetMLProbabilityDictionaryArrayStorageClass() MLProbabilityDictionaryArrayStorageClass {
	return getMLProbabilityDictionaryArrayStorageClass()
}

type MLProbabilityDictionaryArrayStorageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLProbabilityDictionaryArrayStorageClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLProbabilityDictionaryArrayStorageClass) Alloc() MLProbabilityDictionaryArrayStorage {
	rv := objc.Send[MLProbabilityDictionaryArrayStorage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLProbabilityDictionaryArrayStorage.Count]
//   - [MLProbabilityDictionaryArrayStorage.MaxElementIndex]
//   - [MLProbabilityDictionaryArrayStorage.ProbabilityAtIndex]
//   - [MLProbabilityDictionaryArrayStorage.InitWithArray]
//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionaryArrayStorage
type MLProbabilityDictionaryArrayStorage struct {
	objectivec.Object
}

// MLProbabilityDictionaryArrayStorageFromID constructs a [MLProbabilityDictionaryArrayStorage] from an objc.ID.
func MLProbabilityDictionaryArrayStorageFromID(id objc.ID) MLProbabilityDictionaryArrayStorage {
	return MLProbabilityDictionaryArrayStorage{objectivec.Object{ID: id}}
}

// Ensure MLProbabilityDictionaryArrayStorage implements IMLProbabilityDictionaryArrayStorage.
var _ IMLProbabilityDictionaryArrayStorage = MLProbabilityDictionaryArrayStorage{}

// An interface definition for the [MLProbabilityDictionaryArrayStorage] class.
//
// # Methods
//
//   - [IMLProbabilityDictionaryArrayStorage.Count]
//   - [IMLProbabilityDictionaryArrayStorage.MaxElementIndex]
//   - [IMLProbabilityDictionaryArrayStorage.ProbabilityAtIndex]
//   - [IMLProbabilityDictionaryArrayStorage.InitWithArray]
//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionaryArrayStorage
type IMLProbabilityDictionaryArrayStorage interface {
	objectivec.IObject

	// Topic: Methods

	Count() uint64
	MaxElementIndex() uint64
	ProbabilityAtIndex(index uint64) objectivec.IObject
	InitWithArray(array objectivec.IObject) MLProbabilityDictionaryArrayStorage
}

// Init initializes the instance.
func (p MLProbabilityDictionaryArrayStorage) Init() MLProbabilityDictionaryArrayStorage {
	rv := objc.Send[MLProbabilityDictionaryArrayStorage](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLProbabilityDictionaryArrayStorage) Autorelease() MLProbabilityDictionaryArrayStorage {
	rv := objc.Send[MLProbabilityDictionaryArrayStorage](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLProbabilityDictionaryArrayStorage creates a new MLProbabilityDictionaryArrayStorage instance.
func NewMLProbabilityDictionaryArrayStorage() MLProbabilityDictionaryArrayStorage {
	class := getMLProbabilityDictionaryArrayStorageClass()
	rv := objc.Send[MLProbabilityDictionaryArrayStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionaryArrayStorage/initWithArray:
func NewProbabilityDictionaryArrayStorageWithArray(array objectivec.IObject) MLProbabilityDictionaryArrayStorage {
	instance := getMLProbabilityDictionaryArrayStorageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithArray:"), array)
	return MLProbabilityDictionaryArrayStorageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionaryArrayStorage/count
func (p MLProbabilityDictionaryArrayStorage) Count() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("count"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionaryArrayStorage/maxElementIndex
func (p MLProbabilityDictionaryArrayStorage) MaxElementIndex() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("maxElementIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionaryArrayStorage/probabilityAtIndex:
func (p MLProbabilityDictionaryArrayStorage) ProbabilityAtIndex(index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("probabilityAtIndex:"), index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionaryArrayStorage/initWithArray:
func (p MLProbabilityDictionaryArrayStorage) InitWithArray(array objectivec.IObject) MLProbabilityDictionaryArrayStorage {
	rv := objc.Send[MLProbabilityDictionaryArrayStorage](p.ID, objc.Sel("initWithArray:"), array)
	return rv
}
