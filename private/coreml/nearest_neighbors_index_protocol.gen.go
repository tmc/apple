// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLNearestNeighborsIndex protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsIndex
type MLNearestNeighborsIndex interface {
	objectivec.IObject

	// DataPointCount protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsIndex/dataPointCount
	DataPointCount() uint64

	// FindNearestNeighborsToIndex protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsIndex/findNearestNeighbors:toIndex:
	FindNearestNeighborsToIndex(neighbors uint64, index uint64) objectivec.IObject

	// FindNearestNeighborsToQueryPoint protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsIndex/findNearestNeighbors:toQueryPoint:
	FindNearestNeighborsToQueryPoint(neighbors uint64, point unsafe.Pointer) objectivec.IObject

	// UpdateWithDataError protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsIndex/updateWithData:error:
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

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsIndex/dataPointCount
func (o MLNearestNeighborsIndexObject) DataPointCount() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("dataPointCount"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsIndex/findNearestNeighbors:toIndex:
func (o MLNearestNeighborsIndexObject) FindNearestNeighborsToIndex(neighbors uint64, index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("findNearestNeighbors:toIndex:"), neighbors, index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsIndex/findNearestNeighbors:toQueryPoint:
func (o MLNearestNeighborsIndexObject) FindNearestNeighborsToQueryPoint(neighbors uint64, point unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("findNearestNeighbors:toQueryPoint:"), neighbors, point)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsIndex/updateWithData:error:
func (o MLNearestNeighborsIndexObject) UpdateWithDataError(data unsafe.Pointer) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("updateWithData:error:"), data)
	if err != nil {
		return false, err
	}
	return rv, nil
}
