// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLProbabilityDictionaryFloat64Storage] class.
var (
	_MLProbabilityDictionaryFloat64StorageClass     MLProbabilityDictionaryFloat64StorageClass
	_MLProbabilityDictionaryFloat64StorageClassOnce sync.Once
)

func getMLProbabilityDictionaryFloat64StorageClass() MLProbabilityDictionaryFloat64StorageClass {
	_MLProbabilityDictionaryFloat64StorageClassOnce.Do(func() {
		_MLProbabilityDictionaryFloat64StorageClass = MLProbabilityDictionaryFloat64StorageClass{class: objc.GetClass("MLProbabilityDictionaryFloat64Storage")}
	})
	return _MLProbabilityDictionaryFloat64StorageClass
}

// GetMLProbabilityDictionaryFloat64StorageClass returns the class object for MLProbabilityDictionaryFloat64Storage.
func GetMLProbabilityDictionaryFloat64StorageClass() MLProbabilityDictionaryFloat64StorageClass {
	return getMLProbabilityDictionaryFloat64StorageClass()
}

type MLProbabilityDictionaryFloat64StorageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLProbabilityDictionaryFloat64StorageClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLProbabilityDictionaryFloat64StorageClass) Alloc() MLProbabilityDictionaryFloat64Storage {
	rv := objc.SendIfResponds[MLProbabilityDictionaryFloat64Storage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLProbabilityDictionaryFloat64Storage.Count]
//   - [MLProbabilityDictionaryFloat64Storage.MaxElementIndex]
//   - [MLProbabilityDictionaryFloat64Storage.ProbabilityAtIndex]
//   - [MLProbabilityDictionaryFloat64Storage.InitWithFloat64CArrayCount]
type MLProbabilityDictionaryFloat64Storage struct {
	objectivec.Object
}

// MLProbabilityDictionaryFloat64StorageFromID constructs a [MLProbabilityDictionaryFloat64Storage] from an objc.ID.
func MLProbabilityDictionaryFloat64StorageFromID(id objc.ID) MLProbabilityDictionaryFloat64Storage {
	return MLProbabilityDictionaryFloat64Storage{objectivec.Object{ID: id}}
}

// Ensure MLProbabilityDictionaryFloat64Storage implements IMLProbabilityDictionaryFloat64Storage.
var _ IMLProbabilityDictionaryFloat64Storage = MLProbabilityDictionaryFloat64Storage{}

// An interface definition for the [MLProbabilityDictionaryFloat64Storage] class.
//
// # Methods
//
//   - [IMLProbabilityDictionaryFloat64Storage.Count]
//   - [IMLProbabilityDictionaryFloat64Storage.MaxElementIndex]
//   - [IMLProbabilityDictionaryFloat64Storage.ProbabilityAtIndex]
//   - [IMLProbabilityDictionaryFloat64Storage.InitWithFloat64CArrayCount]
type IMLProbabilityDictionaryFloat64Storage interface {
	objectivec.IObject

	// Topic: Methods

	Count() uint64
	MaxElementIndex() uint64
	ProbabilityAtIndex(index uint64) objectivec.IObject
	InitWithFloat64CArrayCount(float64CArray []float64, count uint64) MLProbabilityDictionaryFloat64Storage
}

// Init initializes the instance.
func (m MLProbabilityDictionaryFloat64Storage) Init() MLProbabilityDictionaryFloat64Storage {
	rv := objc.SendIfResponds[MLProbabilityDictionaryFloat64Storage](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLProbabilityDictionaryFloat64Storage) Autorelease() MLProbabilityDictionaryFloat64Storage {
	rv := objc.SendIfResponds[MLProbabilityDictionaryFloat64Storage](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLProbabilityDictionaryFloat64Storage creates a new MLProbabilityDictionaryFloat64Storage instance.
func NewMLProbabilityDictionaryFloat64Storage() MLProbabilityDictionaryFloat64Storage {
	class := getMLProbabilityDictionaryFloat64StorageClass()
	rv := objc.SendIfResponds[MLProbabilityDictionaryFloat64Storage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m MLProbabilityDictionaryFloat64Storage) Count() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("count"))
	return rv
}
func (m MLProbabilityDictionaryFloat64Storage) MaxElementIndex() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("maxElementIndex"))
	return rv
}
func (m MLProbabilityDictionaryFloat64Storage) ProbabilityAtIndex(index uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("probabilityAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (m MLProbabilityDictionaryFloat64Storage) InitWithFloat64CArrayCount(float64CArray []float64, count uint64) MLProbabilityDictionaryFloat64Storage {
	rv := objc.SendIfResponds[MLProbabilityDictionaryFloat64Storage](m.ID, objc.Sel("initWithFloat64CArray:count:"), objc.CArray(float64CArray), count)
	return rv
}
