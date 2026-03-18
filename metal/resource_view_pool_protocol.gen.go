// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// Contains views over resources of a specific type, and allows you to manage those views.
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceViewPool
type MTLResourceViewPool interface {
	objectivec.IObject

	// Obtains the resource ID corresponding to the resource view at index 0 in this resource view pool.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResourceViewPool/baseResourceID
	BaseResourceID() MTLResourceID

	// Obtains a reference to the GPU device this pool belongs to.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResourceViewPool/device
	Device() MTLDevice

	// Queries the optional debug label of this resource view pool.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResourceViewPool/label
	Label() string

	// Queries the number of resource views that this pool contains.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResourceViewPool/resourceViewCount
	ResourceViewCount() uint

	// Copies a range of resource views from a source view pool to a destination location in this view pool.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResourceViewPool/copyResourceViewsFromPool:sourceRange:destinationIndex:
	CopyResourceViewsFromPoolSourceRangeDestinationIndex(sourcePool MTLResourceViewPool, sourceRange foundation.NSRange, destinationIndex uint) MTLResourceID
}

// MTLResourceViewPoolObject wraps an existing Objective-C object that conforms to the MTLResourceViewPool protocol.
type MTLResourceViewPoolObject struct {
	objectivec.Object
}
func (o MTLResourceViewPoolObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLResourceViewPoolObjectFromID constructs a [MTLResourceViewPoolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLResourceViewPoolObjectFromID(id objc.ID) MTLResourceViewPoolObject {
	return MTLResourceViewPoolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Obtains the resource ID corresponding to the resource view at index 0 in
// this resource view pool.
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceViewPool/baseResourceID

func (o MTLResourceViewPoolObject) BaseResourceID() MTLResourceID {
	
	rv := objc.Send[MTLResourceID](o.ID, objc.Sel("baseResourceID"))
	return rv
	}

// Obtains a reference to the GPU device this pool belongs to.
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceViewPool/device

func (o MTLResourceViewPoolObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}

// Queries the optional debug label of this resource view pool.
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceViewPool/label

func (o MTLResourceViewPoolObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}

// Queries the number of resource views that this pool contains.
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceViewPool/resourceViewCount

func (o MTLResourceViewPoolObject) ResourceViewCount() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("resourceViewCount"))
	return rv
	}

// Copies a range of resource views from a source view pool to a destination
// location in this view pool.
//
// sourcePool: Resource view pool from which to copy resource views.
//
// sourceRange: The range in the source resource view pool to copy.
//
// destinationIndex: The starting index in this destination view pool into which to copy the
// source range of resource views.
//
// # Return Value
// 
// The [MTLResourceID] of the resource view corresponding to
// `destinationIndex` of the copy in this resource view pool.
//
// [MTLResourceID]: https://developer.apple.com/documentation/Metal/MTLResourceID
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceViewPool/copyResourceViewsFromPool:sourceRange:destinationIndex:

func (o MTLResourceViewPoolObject) CopyResourceViewsFromPoolSourceRangeDestinationIndex(sourcePool MTLResourceViewPool, sourceRange foundation.NSRange, destinationIndex uint) MTLResourceID {
	
	rv := objc.Send[MTLResourceID](o.ID, objc.Sel("copyResourceViewsFromPool:sourceRange:destinationIndex:"), sourcePool, sourceRange, destinationIndex)
	return rv
	}

