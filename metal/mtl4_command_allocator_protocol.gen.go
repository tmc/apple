// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// Manages the memory backing the encoding of GPU commands into command buffers.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandAllocator
type MTL4CommandAllocator interface {
	objectivec.IObject

	// Returns the GPU device that this command allocator belongs to.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandAllocator/device
	Device() MTLDevice

	// Provides the optional label you specify at creation time for debug purposes.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandAllocator/label
	Label() string

	// Queries the size of the internal memory heaps of this command allocator that support encoding commands into command buffers.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandAllocator/allocatedSize()
	AllocatedSize() uint64

	// Marks the command allocator’s heaps for reuse.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandAllocator/reset()
	Reset()
}

// MTL4CommandAllocatorObject wraps an existing Objective-C object that conforms to the MTL4CommandAllocator protocol.
type MTL4CommandAllocatorObject struct {
	objectivec.Object
}
func (o MTL4CommandAllocatorObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTL4CommandAllocatorObjectFromID constructs a [MTL4CommandAllocatorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTL4CommandAllocatorObjectFromID(id objc.ID) MTL4CommandAllocatorObject {
	return MTL4CommandAllocatorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the GPU device that this command allocator belongs to.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandAllocator/device

func (o MTL4CommandAllocatorObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}

// Provides the optional label you specify at creation time for debug
// purposes.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandAllocator/label

func (o MTL4CommandAllocatorObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}

// Queries the size of the internal memory heaps of this command allocator
// that support encoding commands into command buffers.
//
// # Return Value
// 
// A size in bytes.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandAllocator/allocatedSize()

func (o MTL4CommandAllocatorObject) AllocatedSize() uint64 {
	
	rv := objc.Send[uint64](o.ID, objc.Sel("allocatedSize"))
	return rv
	}

// Marks the command allocator’s heaps for reuse.
//
// # Discussion
// 
// Calling this method allows new [MTL4CommandBuffer] to reuse its existing
// internal memory heaps to encode new GPU commands.
// 
// You are responsible to ensure that all command buffers with memory
// originating from this allocator instance are complete before calling
// resetting it.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandAllocator/reset()

func (o MTL4CommandAllocatorObject) Reset() {
	
	objc.Send[struct{}](o.ID, objc.Sel("reset"))
	}

