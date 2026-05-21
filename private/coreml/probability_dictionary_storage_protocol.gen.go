// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLProbabilityDictionaryStorage protocol.
type MLProbabilityDictionaryStorage interface {
	objectivec.IObject

	// Count protocol.
	Count() uint64

	// MaxElementIndex protocol.
	MaxElementIndex() uint64
}

// MLProbabilityDictionaryStorageObject wraps an existing Objective-C object that conforms to the MLProbabilityDictionaryStorage protocol.
type MLProbabilityDictionaryStorageObject struct {
	objectivec.Object
}

func (o MLProbabilityDictionaryStorageObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLProbabilityDictionaryStorageObjectFromID constructs a [MLProbabilityDictionaryStorageObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLProbabilityDictionaryStorageObjectFromID(id objc.ID) MLProbabilityDictionaryStorageObject {
	return MLProbabilityDictionaryStorageObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MLProbabilityDictionaryStorageObject) Count() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("count"))
	return rv
}
func (o MLProbabilityDictionaryStorageObject) MaxElementIndex() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("maxElementIndex"))
	return rv
}
func (o MLProbabilityDictionaryStorageObject) ProbabilityAtIndex(index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("probabilityAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
