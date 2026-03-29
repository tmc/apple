// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLProbabilityDictionarySharedKeySet] class.
var (
	_MLProbabilityDictionarySharedKeySetClass     MLProbabilityDictionarySharedKeySetClass
	_MLProbabilityDictionarySharedKeySetClassOnce sync.Once
)

func getMLProbabilityDictionarySharedKeySetClass() MLProbabilityDictionarySharedKeySetClass {
	_MLProbabilityDictionarySharedKeySetClassOnce.Do(func() {
		_MLProbabilityDictionarySharedKeySetClass = MLProbabilityDictionarySharedKeySetClass{class: objc.GetClass("MLProbabilityDictionarySharedKeySet")}
	})
	return _MLProbabilityDictionarySharedKeySetClass
}

// GetMLProbabilityDictionarySharedKeySetClass returns the class object for MLProbabilityDictionarySharedKeySet.
func GetMLProbabilityDictionarySharedKeySetClass() MLProbabilityDictionarySharedKeySetClass {
	return getMLProbabilityDictionarySharedKeySetClass()
}

type MLProbabilityDictionarySharedKeySetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLProbabilityDictionarySharedKeySetClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLProbabilityDictionarySharedKeySetClass) Alloc() MLProbabilityDictionarySharedKeySet {
	rv := objc.Send[MLProbabilityDictionarySharedKeySet](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLProbabilityDictionarySharedKeySet.Count]
//   - [MLProbabilityDictionarySharedKeySet.IndexOfLabel]
//   - [MLProbabilityDictionarySharedKeySet.LabelAtIndex]
//   - [MLProbabilityDictionarySharedKeySet.LabelEnumerator]
//   - [MLProbabilityDictionarySharedKeySet.UniqueLabelCount]
//   - [MLProbabilityDictionarySharedKeySet.InitWithLabels]
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionarySharedKeySet
type MLProbabilityDictionarySharedKeySet struct {
	objectivec.Object
}

// MLProbabilityDictionarySharedKeySetFromID constructs a [MLProbabilityDictionarySharedKeySet] from an objc.ID.
func MLProbabilityDictionarySharedKeySetFromID(id objc.ID) MLProbabilityDictionarySharedKeySet {
	return MLProbabilityDictionarySharedKeySet{objectivec.Object{ID: id}}
}
// Ensure MLProbabilityDictionarySharedKeySet implements IMLProbabilityDictionarySharedKeySet.
var _ IMLProbabilityDictionarySharedKeySet = MLProbabilityDictionarySharedKeySet{}

// An interface definition for the [MLProbabilityDictionarySharedKeySet] class.
//
// # Methods
//
//   - [IMLProbabilityDictionarySharedKeySet.Count]
//   - [IMLProbabilityDictionarySharedKeySet.IndexOfLabel]
//   - [IMLProbabilityDictionarySharedKeySet.LabelAtIndex]
//   - [IMLProbabilityDictionarySharedKeySet.LabelEnumerator]
//   - [IMLProbabilityDictionarySharedKeySet.UniqueLabelCount]
//   - [IMLProbabilityDictionarySharedKeySet.InitWithLabels]
//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionarySharedKeySet
type IMLProbabilityDictionarySharedKeySet interface {
	objectivec.IObject

	// Topic: Methods

	Count() uint64
	IndexOfLabel(label objectivec.IObject) uint64
	LabelAtIndex(index uint64) objectivec.IObject
	LabelEnumerator() foundation.NSEnumerator
	UniqueLabelCount() uint64
	InitWithLabels(labels objectivec.IObject) MLProbabilityDictionarySharedKeySet
}

// Init initializes the instance.
func (p MLProbabilityDictionarySharedKeySet) Init() MLProbabilityDictionarySharedKeySet {
	rv := objc.Send[MLProbabilityDictionarySharedKeySet](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLProbabilityDictionarySharedKeySet) Autorelease() MLProbabilityDictionarySharedKeySet {
	rv := objc.Send[MLProbabilityDictionarySharedKeySet](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLProbabilityDictionarySharedKeySet creates a new MLProbabilityDictionarySharedKeySet instance.
func NewMLProbabilityDictionarySharedKeySet() MLProbabilityDictionarySharedKeySet {
	class := getMLProbabilityDictionarySharedKeySetClass()
	rv := objc.Send[MLProbabilityDictionarySharedKeySet](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionarySharedKeySet/initWithLabels:
func NewProbabilityDictionarySharedKeySetWithLabels(labels objectivec.IObject) MLProbabilityDictionarySharedKeySet {
	instance := getMLProbabilityDictionarySharedKeySetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLabels:"), labels)
	return MLProbabilityDictionarySharedKeySetFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionarySharedKeySet/indexOfLabel:
func (p MLProbabilityDictionarySharedKeySet) IndexOfLabel(label objectivec.IObject) uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("indexOfLabel:"), label)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionarySharedKeySet/labelAtIndex:
func (p MLProbabilityDictionarySharedKeySet) LabelAtIndex(index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("labelAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionarySharedKeySet/initWithLabels:
func (p MLProbabilityDictionarySharedKeySet) InitWithLabels(labels objectivec.IObject) MLProbabilityDictionarySharedKeySet {
	rv := objc.Send[MLProbabilityDictionarySharedKeySet](p.ID, objc.Sel("initWithLabels:"), labels)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionarySharedKeySet/count
func (p MLProbabilityDictionarySharedKeySet) Count() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("count"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionarySharedKeySet/labelEnumerator
func (p MLProbabilityDictionarySharedKeySet) LabelEnumerator() foundation.NSEnumerator {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("labelEnumerator"))
	return foundation.NSEnumeratorFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLProbabilityDictionarySharedKeySet/uniqueLabelCount
func (p MLProbabilityDictionarySharedKeySet) UniqueLabelCount() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("uniqueLabelCount"))
	return rv
}

