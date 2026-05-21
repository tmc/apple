// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLNearestNeighborsIndex protocol.
type MLNearestNeighborsIndex interface {
	objectivec.IObject

	// DataPointCount protocol.
	DataPointCount() uint64

	// FindNearestNeighborsToIndex protocol.
	FindNearestNeighborsToIndex(neighbors uint64, index uint64) unsafe.Pointer

	// FindNearestNeighborsToQueryPoint protocol.
	FindNearestNeighborsToQueryPoint(neighbors uint64, point unsafe.Pointer) unsafe.Pointer

	// UpdateWithDataError protocol.
	UpdateWithDataError(data unsafe.Pointer) (bool, error)
}

// MLNearestNeighborsIndexObject wraps an existing Objective-C object that conforms to the MLNearestNeighborsIndex protocol.
type MLNearestNeighborsIndexObject struct {
	objectivec.Object
}

func (o MLNearestNeighborsIndexObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLNearestNeighborsIndexObjectFromID constructs a [MLNearestNeighborsIndexObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLNearestNeighborsIndexObjectFromID(id objc.ID) MLNearestNeighborsIndexObject {
	return MLNearestNeighborsIndexObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MLNearestNeighborsIndexObject) DataPointCount() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("dataPointCount"))
	return rv
}
func (o MLNearestNeighborsIndexObject) FindNearestNeighborsToIndex(neighbors uint64, index uint64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("findNearestNeighbors:toIndex:"), neighbors, index)
	return rv
}
func (o MLNearestNeighborsIndexObject) FindNearestNeighborsToQueryPoint(neighbors uint64, point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("findNearestNeighbors:toQueryPoint:"), neighbors, point)
	return rv
}
func (o MLNearestNeighborsIndexObject) UpdateWithDataError(data unsafe.Pointer) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("updateWithData:error:"), data)
	if err != nil {
		return false, err
	}
	return rv, nil
}
