// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLComputeDataProvider protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLComputeDataProvider
type MLComputeDataProvider interface {
	objectivec.IObject

	// NumberOfBatches protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLComputeDataProvider/numberOfBatches
	NumberOfBatches() uint64

	// SizeOfBatchAtIndex protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLComputeDataProvider/sizeOfBatchAtIndex:
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

// See: https://developer.apple.com/documentation/CoreML/MLComputeDataProvider/batchAtIndex:error:
func (o MLComputeDataProviderObject) BatchAtIndexError(index uint64) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("batchAtIndex:error:"), index)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}

// See: https://developer.apple.com/documentation/CoreML/MLComputeDataProvider/numberOfBatches
func (o MLComputeDataProviderObject) NumberOfBatches() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("numberOfBatches"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLComputeDataProvider/sizeOfBatchAtIndex:
func (o MLComputeDataProviderObject) SizeOfBatchAtIndex(index uint64) uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("sizeOfBatchAtIndex:"), index)
	return rv
}
