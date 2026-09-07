// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLBatchProvider protocol.
type MLBatchProvider interface {
	objectivec.IObject

	// Count protocol.
	Count() int64

	// FeaturesAtIndex protocol.
	FeaturesAtIndex(index int64) objectivec.IObject
}

// MLBatchProviderObject wraps an existing Objective-C object that conforms to the MLBatchProvider protocol.
type MLBatchProviderObject struct {
	objectivec.Object
}

func (o MLBatchProviderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLBatchProviderObjectFromID constructs a [MLBatchProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLBatchProviderObjectFromID(id objc.ID) MLBatchProviderObject {
	return MLBatchProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MLBatchProviderObject) Count() int64 {
	rv := objc.SendIfResponds[int64](o.ID, objc.Sel("count"))
	return rv
}
func (o MLBatchProviderObject) FeaturesAtIndex(index int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("featuresAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
