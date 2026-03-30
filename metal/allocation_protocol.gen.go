// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A memory allocation from a Metal GPU device, such as a memory heap, texture, or data buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLAllocation
type MTLAllocation interface {
	objectivec.IObject

	// The amount of memory, in byes, a resource consumes, such as for a buffer, texture, or heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLAllocation/allocatedSize
	AllocatedSize() uint
}

// MTLAllocationObject wraps an existing Objective-C object that conforms to the MTLAllocation protocol.
type MTLAllocationObject struct {
	objectivec.Object
}

func (o MTLAllocationObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLAllocationObjectFromID constructs a [MTLAllocationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLAllocationObjectFromID(id objc.ID) MTLAllocationObject {
	return MTLAllocationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The amount of memory, in byes, a resource consumes, such as for a buffer,
// texture, or heap.
//
// See: https://developer.apple.com/documentation/Metal/MTLAllocation/allocatedSize
func (o MTLAllocationObject) AllocatedSize() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("allocatedSize"))
	return rv
}
