// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLComputeDataProvider protocol.
type MLComputeDataProvider interface {
	objectivec.IObject

	// BatchAtIndexError protocol.
	BatchAtIndexError(index uint64) (objectivec.IObject, error)

	// NumberOfBatches protocol.
	NumberOfBatches() uint64

	// SizeOfBatchAtIndex protocol.
	SizeOfBatchAtIndex(index uint64) uint64
}

// MLComputeDataProviderObject wraps an existing Objective-C object that conforms to the MLComputeDataProvider protocol.
type MLComputeDataProviderObject struct {
	objectivec.Object
}

func (o MLComputeDataProviderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLComputeDataProviderObjectFromID constructs a [MLComputeDataProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLComputeDataProviderObjectFromID(id objc.ID) MLComputeDataProviderObject {
	return MLComputeDataProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MLComputeDataProviderObject) BatchAtIndexError(index uint64) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("batchAtIndex:error:"), index)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
func (o MLComputeDataProviderObject) NumberOfBatches() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("numberOfBatches"))
	return rv
}
func (o MLComputeDataProviderObject) SizeOfBatchAtIndex(index uint64) uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("sizeOfBatchAtIndex:"), index)
	return rv
}
