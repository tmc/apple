// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLProbabilityDictionary] class.
var (
	_MLProbabilityDictionaryClass     MLProbabilityDictionaryClass
	_MLProbabilityDictionaryClassOnce sync.Once
)

func getMLProbabilityDictionaryClass() MLProbabilityDictionaryClass {
	_MLProbabilityDictionaryClassOnce.Do(func() {
		_MLProbabilityDictionaryClass = MLProbabilityDictionaryClass{class: objc.GetClass("MLProbabilityDictionary")}
	})
	return _MLProbabilityDictionaryClass
}

// GetMLProbabilityDictionaryClass returns the class object for MLProbabilityDictionary.
func GetMLProbabilityDictionaryClass() MLProbabilityDictionaryClass {
	return getMLProbabilityDictionaryClass()
}

type MLProbabilityDictionaryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLProbabilityDictionaryClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLProbabilityDictionaryClass) Alloc() MLProbabilityDictionary {
	rv := objc.SendIfResponds[MLProbabilityDictionary](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLProbabilityDictionary.ClassLabelOfMaxProbability]
//   - [MLProbabilityDictionary.LabelIndexMap]
//   - [MLProbabilityDictionary.Storage]
//   - [MLProbabilityDictionary.InitWithLabelIndexMapStorage]
//   - [MLProbabilityDictionary.InitWithLabelsProbabilities]
//   - [MLProbabilityDictionary.InitWithLabelsProbabilityArray]
//   - [MLProbabilityDictionary.InitWithObjectsForKeysCount]
//   - [MLProbabilityDictionary.InitWithSharedKeySetProbabilities]
//   - [MLProbabilityDictionary.InitWithSharedKeySetProbabilityArray]
//   - [MLProbabilityDictionary.InitWithSharedKeySetProbabilityMultiArray]
type MLProbabilityDictionary struct {
	foundation.NSDictionary
}

// MLProbabilityDictionaryFromID constructs a [MLProbabilityDictionary] from an objc.ID.
func MLProbabilityDictionaryFromID(id objc.ID) MLProbabilityDictionary {
	return MLProbabilityDictionary{NSDictionary: foundation.NSDictionaryFromID(id)}
}

// Ensure MLProbabilityDictionary implements IMLProbabilityDictionary.
var _ IMLProbabilityDictionary = MLProbabilityDictionary{}

// An interface definition for the [MLProbabilityDictionary] class.
//
// # Methods
//
//   - [IMLProbabilityDictionary.ClassLabelOfMaxProbability]
//   - [IMLProbabilityDictionary.LabelIndexMap]
//   - [IMLProbabilityDictionary.Storage]
//   - [IMLProbabilityDictionary.InitWithLabelIndexMapStorage]
//   - [IMLProbabilityDictionary.InitWithLabelsProbabilities]
//   - [IMLProbabilityDictionary.InitWithLabelsProbabilityArray]
//   - [IMLProbabilityDictionary.InitWithObjectsForKeysCount]
//   - [IMLProbabilityDictionary.InitWithSharedKeySetProbabilities]
//   - [IMLProbabilityDictionary.InitWithSharedKeySetProbabilityArray]
//   - [IMLProbabilityDictionary.InitWithSharedKeySetProbabilityMultiArray]
type IMLProbabilityDictionary interface {
	foundation.INSDictionary

	// Topic: Methods

	ClassLabelOfMaxProbability() objectivec.IObject
	LabelIndexMap() IMLProbabilityDictionarySharedKeySet
	Storage() unsafe.Pointer
	InitWithLabelIndexMapStorage(map_ objectivec.IObject, storage objectivec.IObject) MLProbabilityDictionary
	InitWithLabelsProbabilities(labels objectivec.IObject, probabilities []float64) MLProbabilityDictionary
	InitWithLabelsProbabilityArray(labels objectivec.IObject, array objectivec.IObject) MLProbabilityDictionary
	InitWithObjectsForKeysCount(objects []objectivec.IObject, keys []objectivec.IObject, count uint64) MLProbabilityDictionary
	InitWithSharedKeySetProbabilities(set objectivec.IObject, probabilities []float64) MLProbabilityDictionary
	InitWithSharedKeySetProbabilityArray(set objectivec.IObject, array objectivec.IObject) MLProbabilityDictionary
	InitWithSharedKeySetProbabilityMultiArray(set objectivec.IObject, array objectivec.IObject) MLProbabilityDictionary
}

// Init initializes the instance.
func (m MLProbabilityDictionary) Init() MLProbabilityDictionary {
	rv := objc.SendIfResponds[MLProbabilityDictionary](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLProbabilityDictionary) Autorelease() MLProbabilityDictionary {
	rv := objc.SendIfResponds[MLProbabilityDictionary](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLProbabilityDictionary creates a new MLProbabilityDictionary instance.
func NewMLProbabilityDictionary() MLProbabilityDictionary {
	class := getMLProbabilityDictionaryClass()
	rv := objc.SendIfResponds[MLProbabilityDictionary](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewProbabilityDictionaryWithLabelIndexMapStorage(map_ objectivec.IObject, storage objectivec.IObject) MLProbabilityDictionary {
	instance := getMLProbabilityDictionaryClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithLabelIndexMap:storage:"), map_, storage)
	return MLProbabilityDictionaryFromID(rv)
}

func NewProbabilityDictionaryWithLabelsProbabilities(labels objectivec.IObject, probabilities []float64) MLProbabilityDictionary {
	instance := getMLProbabilityDictionaryClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithLabels:probabilities:"), labels, probabilities)
	return MLProbabilityDictionaryFromID(rv)
}

func NewProbabilityDictionaryWithLabelsProbabilityArray(labels objectivec.IObject, array objectivec.IObject) MLProbabilityDictionary {
	instance := getMLProbabilityDictionaryClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithLabels:probabilityArray:"), labels, array)
	return MLProbabilityDictionaryFromID(rv)
}

func NewProbabilityDictionaryWithSharedKeySetProbabilities(set objectivec.IObject, probabilities []float64) MLProbabilityDictionary {
	instance := getMLProbabilityDictionaryClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSharedKeySet:probabilities:"), set, probabilities)
	return MLProbabilityDictionaryFromID(rv)
}

func NewProbabilityDictionaryWithSharedKeySetProbabilityArray(set objectivec.IObject, array objectivec.IObject) MLProbabilityDictionary {
	instance := getMLProbabilityDictionaryClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSharedKeySet:probabilityArray:"), set, array)
	return MLProbabilityDictionaryFromID(rv)
}

func NewProbabilityDictionaryWithSharedKeySetProbabilityMultiArray(set objectivec.IObject, array objectivec.IObject) MLProbabilityDictionary {
	instance := getMLProbabilityDictionaryClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSharedKeySet:probabilityMultiArray:"), set, array)
	return MLProbabilityDictionaryFromID(rv)
}

func (m MLProbabilityDictionary) ClassLabelOfMaxProbability() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("classLabelOfMaxProbability"))
	return objectivec.Object{ID: rv}
}
func (m MLProbabilityDictionary) InitWithLabelIndexMapStorage(map_ objectivec.IObject, storage objectivec.IObject) MLProbabilityDictionary {
	rv := objc.SendIfResponds[MLProbabilityDictionary](m.ID, objc.Sel("initWithLabelIndexMap:storage:"), map_, storage)
	return rv
}
func (m MLProbabilityDictionary) InitWithLabelsProbabilities(labels objectivec.IObject, probabilities []float64) MLProbabilityDictionary {
	rv := objc.SendIfResponds[MLProbabilityDictionary](m.ID, objc.Sel("initWithLabels:probabilities:"), labels, probabilities)
	return rv
}
func (m MLProbabilityDictionary) InitWithLabelsProbabilityArray(labels objectivec.IObject, array objectivec.IObject) MLProbabilityDictionary {
	rv := objc.SendIfResponds[MLProbabilityDictionary](m.ID, objc.Sel("initWithLabels:probabilityArray:"), labels, array)
	return rv
}
func (m MLProbabilityDictionary) InitWithObjectsForKeysCount(objects []objectivec.IObject, keys []objectivec.IObject, count uint64) MLProbabilityDictionary {
	rv := objc.SendIfResponds[MLProbabilityDictionary](m.ID, objc.Sel("initWithObjects:forKeys:count:"), objc.CArray(objects), objc.CArray(keys), count)
	return rv
}
func (m MLProbabilityDictionary) InitWithSharedKeySetProbabilities(set objectivec.IObject, probabilities []float64) MLProbabilityDictionary {
	rv := objc.SendIfResponds[MLProbabilityDictionary](m.ID, objc.Sel("initWithSharedKeySet:probabilities:"), set, probabilities)
	return rv
}
func (m MLProbabilityDictionary) InitWithSharedKeySetProbabilityArray(set objectivec.IObject, array objectivec.IObject) MLProbabilityDictionary {
	rv := objc.SendIfResponds[MLProbabilityDictionary](m.ID, objc.Sel("initWithSharedKeySet:probabilityArray:"), set, array)
	return rv
}
func (m MLProbabilityDictionary) InitWithSharedKeySetProbabilityMultiArray(set objectivec.IObject, array objectivec.IObject) MLProbabilityDictionary {
	rv := objc.SendIfResponds[MLProbabilityDictionary](m.ID, objc.Sel("initWithSharedKeySet:probabilityMultiArray:"), set, array)
	return rv
}

func (_MLProbabilityDictionaryClass MLProbabilityDictionaryClass) SharedKeySetForLabels(labels objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLProbabilityDictionaryClass.class), objc.Sel("sharedKeySetForLabels:"), labels)
	return objectivec.Object{ID: rv}
}

func (m MLProbabilityDictionary) LabelIndexMap() IMLProbabilityDictionarySharedKeySet {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("labelIndexMap"))
	return MLProbabilityDictionarySharedKeySetFromID(objc.ID(rv))
}
func (m MLProbabilityDictionary) Storage() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("storage"))
	return rv
}
