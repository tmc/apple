// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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
	rv := objc.Send[MLProbabilityDictionary](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLProbabilityDictionary.ClassLabelOfMaxProbability]
//   - [MLProbabilityDictionary.LabelIndexMap]
//   - [MLProbabilityDictionary.Storage]
//   - [MLProbabilityDictionary.InitWithLabelIndexMapStorage]
//   - [MLProbabilityDictionary.InitWithLabelsProbabilities]
//   - [MLProbabilityDictionary.InitWithLabelsProbabilityArray]
//   - [MLProbabilityDictionary.InitWithSharedKeySetProbabilities]
//   - [MLProbabilityDictionary.InitWithSharedKeySetProbabilityArray]
//   - [MLProbabilityDictionary.InitWithSharedKeySetProbabilityMultiArray]
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionary
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
//   - [IMLProbabilityDictionary.InitWithSharedKeySetProbabilities]
//   - [IMLProbabilityDictionary.InitWithSharedKeySetProbabilityArray]
//   - [IMLProbabilityDictionary.InitWithSharedKeySetProbabilityMultiArray]
//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionary
type IMLProbabilityDictionary interface {
	foundation.INSDictionary

	// Topic: Methods

	ClassLabelOfMaxProbability() objectivec.IObject
	LabelIndexMap() IMLProbabilityDictionarySharedKeySet
	Storage() objectivec.IObject
	InitWithLabelIndexMapStorage(map_ objectivec.IObject, storage objectivec.IObject) MLProbabilityDictionary
	InitWithLabelsProbabilities(labels objectivec.IObject, probabilities []float64) MLProbabilityDictionary
	InitWithLabelsProbabilityArray(labels objectivec.IObject, array objectivec.IObject) MLProbabilityDictionary
	InitWithSharedKeySetProbabilities(set objectivec.IObject, probabilities []float64) MLProbabilityDictionary
	InitWithSharedKeySetProbabilityArray(set objectivec.IObject, array objectivec.IObject) MLProbabilityDictionary
	InitWithSharedKeySetProbabilityMultiArray(set objectivec.IObject, array objectivec.IObject) MLProbabilityDictionary
}

// Init initializes the instance.
func (p MLProbabilityDictionary) Init() MLProbabilityDictionary {
	rv := objc.Send[MLProbabilityDictionary](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLProbabilityDictionary) Autorelease() MLProbabilityDictionary {
	rv := objc.Send[MLProbabilityDictionary](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLProbabilityDictionary creates a new MLProbabilityDictionary instance.
func NewMLProbabilityDictionary() MLProbabilityDictionary {
	class := getMLProbabilityDictionaryClass()
	rv := objc.Send[MLProbabilityDictionary](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionary/initWithLabelIndexMap:storage:
func NewProbabilityDictionaryWithLabelIndexMapStorage(map_ objectivec.IObject, storage objectivec.IObject) MLProbabilityDictionary {
	instance := getMLProbabilityDictionaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLabelIndexMap:storage:"), map_, storage)
	return MLProbabilityDictionaryFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionary/initWithLabels:probabilities:
func NewProbabilityDictionaryWithLabelsProbabilities(labels objectivec.IObject, probabilities []float64) MLProbabilityDictionary {
	instance := getMLProbabilityDictionaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLabels:probabilities:"), labels, probabilities)
	return MLProbabilityDictionaryFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionary/initWithLabels:probabilityArray:
func NewProbabilityDictionaryWithLabelsProbabilityArray(labels objectivec.IObject, array objectivec.IObject) MLProbabilityDictionary {
	instance := getMLProbabilityDictionaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLabels:probabilityArray:"), labels, array)
	return MLProbabilityDictionaryFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionary/initWithSharedKeySet:probabilities:
func NewProbabilityDictionaryWithSharedKeySetProbabilities(set objectivec.IObject, probabilities []float64) MLProbabilityDictionary {
	instance := getMLProbabilityDictionaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSharedKeySet:probabilities:"), set, probabilities)
	return MLProbabilityDictionaryFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionary/initWithSharedKeySet:probabilityArray:
func NewProbabilityDictionaryWithSharedKeySetProbabilityArray(set objectivec.IObject, array objectivec.IObject) MLProbabilityDictionary {
	instance := getMLProbabilityDictionaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSharedKeySet:probabilityArray:"), set, array)
	return MLProbabilityDictionaryFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionary/initWithSharedKeySet:probabilityMultiArray:
func NewProbabilityDictionaryWithSharedKeySetProbabilityMultiArray(set objectivec.IObject, array objectivec.IObject) MLProbabilityDictionary {
	instance := getMLProbabilityDictionaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSharedKeySet:probabilityMultiArray:"), set, array)
	return MLProbabilityDictionaryFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionary/classLabelOfMaxProbability
func (p MLProbabilityDictionary) ClassLabelOfMaxProbability() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("classLabelOfMaxProbability"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionary/initWithLabelIndexMap:storage:
func (p MLProbabilityDictionary) InitWithLabelIndexMapStorage(map_ objectivec.IObject, storage objectivec.IObject) MLProbabilityDictionary {
	rv := objc.Send[MLProbabilityDictionary](p.ID, objc.Sel("initWithLabelIndexMap:storage:"), map_, storage)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionary/initWithLabels:probabilities:
func (p MLProbabilityDictionary) InitWithLabelsProbabilities(labels objectivec.IObject, probabilities []float64) MLProbabilityDictionary {
	rv := objc.Send[MLProbabilityDictionary](p.ID, objc.Sel("initWithLabels:probabilities:"), labels, probabilities)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionary/initWithLabels:probabilityArray:
func (p MLProbabilityDictionary) InitWithLabelsProbabilityArray(labels objectivec.IObject, array objectivec.IObject) MLProbabilityDictionary {
	rv := objc.Send[MLProbabilityDictionary](p.ID, objc.Sel("initWithLabels:probabilityArray:"), labels, array)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionary/initWithSharedKeySet:probabilities:
func (p MLProbabilityDictionary) InitWithSharedKeySetProbabilities(set objectivec.IObject, probabilities []float64) MLProbabilityDictionary {
	rv := objc.Send[MLProbabilityDictionary](p.ID, objc.Sel("initWithSharedKeySet:probabilities:"), set, probabilities)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionary/initWithSharedKeySet:probabilityArray:
func (p MLProbabilityDictionary) InitWithSharedKeySetProbabilityArray(set objectivec.IObject, array objectivec.IObject) MLProbabilityDictionary {
	rv := objc.Send[MLProbabilityDictionary](p.ID, objc.Sel("initWithSharedKeySet:probabilityArray:"), set, array)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionary/initWithSharedKeySet:probabilityMultiArray:
func (p MLProbabilityDictionary) InitWithSharedKeySetProbabilityMultiArray(set objectivec.IObject, array objectivec.IObject) MLProbabilityDictionary {
	rv := objc.Send[MLProbabilityDictionary](p.ID, objc.Sel("initWithSharedKeySet:probabilityMultiArray:"), set, array)
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionary/sharedKeySetForLabels:
func (_MLProbabilityDictionaryClass MLProbabilityDictionaryClass) SharedKeySetForLabels(labels objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLProbabilityDictionaryClass.class), objc.Sel("sharedKeySetForLabels:"), labels)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionary/labelIndexMap
func (p MLProbabilityDictionary) LabelIndexMap() IMLProbabilityDictionarySharedKeySet {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("labelIndexMap"))
	return MLProbabilityDictionarySharedKeySetFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionary/storage
func (p MLProbabilityDictionary) Storage() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("storage"))
	return objectivec.Object{ID: rv}
}

