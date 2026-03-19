// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol your app implements that wraps a Metal buffer instance to serve as scratch memory for an input/output command queue.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOScratchBuffer
type MTLIOScratchBuffer interface {
	objectivec.IObject

	// A Metal buffer that serves as scratch memory for an input/output command queue.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOScratchBuffer/buffer
	Buffer() MTLBuffer
}

// MTLIOScratchBufferObject wraps an existing Objective-C object that conforms to the MTLIOScratchBuffer protocol.
type MTLIOScratchBufferObject struct {
	objectivec.Object
}
func (o MTLIOScratchBufferObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLIOScratchBufferObjectFromID constructs a [MTLIOScratchBufferObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLIOScratchBufferObjectFromID(id objc.ID) MTLIOScratchBufferObject {
	return MTLIOScratchBufferObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A Metal buffer that serves as scratch memory for an input/output command
// queue.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOScratchBuffer/buffer
func (o MTLIOScratchBufferObject) Buffer() MTLBuffer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("buffer"))
	return MTLBufferObjectFromID(rv)
	}

