// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol your app implements to provide scratch memory to an input/output command queue.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOScratchBufferAllocator
type MTLIOScratchBufferAllocator interface {
	objectivec.IObject

	// Creates a scratch memory buffer for an input/output command queue.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOScratchBufferAllocator/makeScratchBuffer(minimumSize:)
	NewScratchBufferWithMinimumSize(minimumSize uint) MTLIOScratchBuffer
}

// MTLIOScratchBufferAllocatorObject wraps an existing Objective-C object that conforms to the MTLIOScratchBufferAllocator protocol.
type MTLIOScratchBufferAllocatorObject struct {
	objectivec.Object
}
func (o MTLIOScratchBufferAllocatorObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLIOScratchBufferAllocatorObjectFromID constructs a [MTLIOScratchBufferAllocatorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLIOScratchBufferAllocatorObjectFromID(id objc.ID) MTLIOScratchBufferAllocatorObject {
	return MTLIOScratchBufferAllocatorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Creates a scratch memory buffer for an input/output command queue.
//
// minimumSize: The number of bytes the input/output command buffer needs to successfully
// run a command buffer.
//
// # Return Value
// 
// An [MTLIOScratchBuffer] instance that your app implements or `nil`.
//
// # Discussion
// 
// Your app can reduce additional callbacks from the framework by providing
// additional memory above `minimumSize`. If your implementation returns
// `nil`, the input/output command queue cancels the [MTLIOCommandBuffer]
// instance that needs the scratch buffer memory.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOScratchBufferAllocator/makeScratchBuffer(minimumSize:)
func (o MTLIOScratchBufferAllocatorObject) NewScratchBufferWithMinimumSize(minimumSize uint) MTLIOScratchBuffer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newScratchBufferWithMinimumSize:"), minimumSize)
	return MTLIOScratchBufferObjectFromID(rv)
	}

