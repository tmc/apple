// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A command queue that schedules input/output commands for reading files in the file system, and writing to GPU resources and memory.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandQueue
type MTLIOCommandQueue interface {
	objectivec.IObject

	// Creates an input/output command buffer for the command queue.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandQueue/makeCommandBuffer()
	CommandBuffer() MTLIOCommandBuffer

	// Creates an input/output command buffer for the command queue that doesn’t retain the instances you pass to its methods.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandQueue/makeCommandBufferWithUnretainedReferences()
	CommandBufferWithUnretainedReferences() MTLIOCommandBuffer

	// Appends a barrier that tells the input/output command queue to finish running all in-flight command buffers before running any new command buffers.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandQueue/enqueueBarrier()
	EnqueueBarrier()

	// An optional name for the input/output command queue.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandQueue/label
	Label() string

	// An optional name for the input/output command queue.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOCommandQueue/label
	SetLabel(value string)
}

// MTLIOCommandQueueObject wraps an existing Objective-C object that conforms to the MTLIOCommandQueue protocol.
type MTLIOCommandQueueObject struct {
	objectivec.Object
}
func (o MTLIOCommandQueueObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLIOCommandQueueObjectFromID constructs a [MTLIOCommandQueueObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLIOCommandQueueObjectFromID(id objc.ID) MTLIOCommandQueueObject {
	return MTLIOCommandQueueObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Creates an input/output command buffer for the command queue.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandQueue/makeCommandBuffer()

func (o MTLIOCommandQueueObject) CommandBuffer() MTLIOCommandBuffer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("commandBuffer"))
	return MTLIOCommandBufferObjectFromID(rv)
	}

// Creates an input/output command buffer for the command queue that doesn’t
// retain the instances you pass to its methods.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandQueue/makeCommandBufferWithUnretainedReferences()

func (o MTLIOCommandQueueObject) CommandBufferWithUnretainedReferences() MTLIOCommandBuffer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("commandBufferWithUnretainedReferences"))
	return MTLIOCommandBufferObjectFromID(rv)
	}

// Appends a barrier that tells the input/output command queue to finish
// running all in-flight command buffers before running any new command
// buffers.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandQueue/enqueueBarrier()

func (o MTLIOCommandQueueObject) EnqueueBarrier() {
	
	objc.Send[struct{}](o.ID, objc.Sel("enqueueBarrier"))
	}

// An optional name for the input/output command queue.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandQueue/label

func (o MTLIOCommandQueueObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}

func (o MTLIOCommandQueueObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}

