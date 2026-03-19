// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// An instance you use to create, submit, and schedule command buffers to a specific GPU device to run the commands within those buffers.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandQueue
type MTLCommandQueue interface {
	objectivec.IObject

	// Returns a command buffer from the command queue that you configure with a descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandQueue/makeCommandBuffer(descriptor:)
	CommandBufferWithDescriptor(descriptor IMTLCommandBufferDescriptor) MTLCommandBuffer

	// Returns a command buffer from the command queue that maintains strong references to resources.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandQueue/makeCommandBuffer()
	CommandBuffer() MTLCommandBuffer

	// Returns a command buffer from the command queue that doesn’t maintain strong references to resources.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandQueue/makeCommandBufferWithUnretainedReferences()
	CommandBufferWithUnretainedReferences() MTLCommandBuffer

	// Applies a residency set to a queue, which Metal applies to the queue’s command buffers as you commit them.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandQueue/addResidencySet(_:)
	AddResidencySet(residencySet MTLResidencySet)

	// Removes a residency set from a command queue’s list, which means Metal doesn’t apply it to the queue’s command buffers as you commit them.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandQueue/removeResidencySet(_:)
	RemoveResidencySet(residencySet MTLResidencySet)

	// The GPU device that creates the command queue.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandQueue/device
	Device() MTLDevice

	// An optional name that can help you identify the command queue.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandQueue/label
	Label() string

	// Applies multiple residency sets to a queue, which Metal applies to the queue’s command buffers as you commit them.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandQueue/addResidencySets:count:
	AddResidencySetsCount(residencySets []MTLResidencySet, count uint)

	// Removes multiple residency sets from a command queue’s list, which means Metal doesn’t apply them to the queue’s command buffers as you commit them.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandQueue/removeResidencySets:count:
	RemoveResidencySetsCount(residencySets []MTLResidencySet, count uint)

	// An optional name that can help you identify the command queue.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandQueue/label
	SetLabel(value string)
}

// MTLCommandQueueObject wraps an existing Objective-C object that conforms to the MTLCommandQueue protocol.
type MTLCommandQueueObject struct {
	objectivec.Object
}
func (o MTLCommandQueueObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLCommandQueueObjectFromID constructs a [MTLCommandQueueObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLCommandQueueObjectFromID(id objc.ID) MTLCommandQueueObject {
	return MTLCommandQueueObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns a command buffer from the command queue that you configure with a
// descriptor.
//
// descriptor: An [MTLCommandBufferDescriptor] instance that configures the
// [MTLCommandBuffer] the method returns.
//
// # Discussion
// 
// Use this method to create a command buffer that you configure with a
// descriptor. You can configure whether the command buffer retains references
// to resources that its commands refer to by setting the `descriptor`
// parameter’s [RetainedReferences] property. You can also configure whether
// the command buffer saves extra error information, which is useful during
// development, by setting the descriptor’s [ErrorOptions] property.
// 
// Each command queue has a fixed number of command buffers for its lifetime
// (see [NewCommandQueueWithMaxCommandBufferCount]). This method blocks the
// calling CPU thread when the queue doesn’t have any free command buffers,
// and returns after the GPU finishes executing one.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandQueue/makeCommandBuffer(descriptor:)
func (o MTLCommandQueueObject) CommandBufferWithDescriptor(descriptor IMTLCommandBufferDescriptor) MTLCommandBuffer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("commandBufferWithDescriptor:"), descriptor)
	return MTLCommandBufferObjectFromID(rv)
	}
// Returns a command buffer from the command queue that maintains strong
// references to resources.
//
// # Discussion
// 
// The command buffers you create with this method maintain strong references
// to the resources you encode into it, including buffers, textures, samplers,
// and pipeline states. The command buffer releases these references after it
// finishes running on the GPU.
// 
// This method sets the [retainedReferences] property to [true] for the
// command buffer it creates.
// 
// Each command queue has a fixed number of command buffers for its lifetime
// (see [NewCommandQueueWithMaxCommandBufferCount]). This method blocks the
// calling CPU thread when the queue doesn’t have any free command buffers,
// and returns after the GPU finishes executing one.
//
// [retainedReferences]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/retainedReferences
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandQueue/makeCommandBuffer()
func (o MTLCommandQueueObject) CommandBuffer() MTLCommandBuffer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("commandBuffer"))
	return MTLCommandBufferObjectFromID(rv)
	}
// Returns a command buffer from the command queue that doesn’t maintain
// strong references to resources.
//
// # Discussion
// 
// Use this method to create a command buffer that doesn’t retain or release
// any of the resources it needs to run its commands.
// 
// Apps typically create command buffers that don’t maintain references to
// resources for extremely performance-critical situations. Even though the
// runtime cost for retaining or releasing a single resource is trivial, the
// aggregate time savings may be worth it.
// 
// It’s your app’s responsibility to maintain strong references to all the
// resources the command buffer uses until it finishes running on the GPU.
// 
// This method sets the [retainedReferences] property to [false] for the
// command buffer it creates.
// 
// Each command queue has a fixed number of command buffers for its lifetime
// (see [NewCommandQueueWithMaxCommandBufferCount]). This method blocks the
// calling CPU thread when the queue doesn’t have any free command buffers,
// and returns after the GPU finishes executing one.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [retainedReferences]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/retainedReferences
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandQueue/makeCommandBufferWithUnretainedReferences()
func (o MTLCommandQueueObject) CommandBufferWithUnretainedReferences() MTLCommandBuffer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("commandBufferWithUnretainedReferences"))
	return MTLCommandBufferObjectFromID(rv)
	}
// Applies a residency set to a queue, which Metal applies to the queue’s
// command buffers as you commit them.
//
// residencySet: A residency set that contains resource allocations, such as [MTLBuffer],
// [MTLTexture], and [MTLHeap] instances.
//
// # Discussion
// 
// Each command queue can maintain a list of up to 32 different residency
// sets. See [Simplifying GPU resource management with residency sets] and
// [MTLResidencySet] for more information.
//
// [Simplifying GPU resource management with residency sets]: https://developer.apple.com/documentation/Metal/simplifying-gpu-resource-management-with-residency-sets
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandQueue/addResidencySet(_:)
func (o MTLCommandQueueObject) AddResidencySet(residencySet MTLResidencySet) {
	
	objc.Send[struct{}](o.ID, objc.Sel("addResidencySet:"), residencySet)
	}
// Removes a residency set from a command queue’s list, which means Metal
// doesn’t apply it to the queue’s command buffers as you commit them.
//
// residencySet: A residency set that contains resource allocations, such as [MTLBuffer],
// [MTLTexture], and [MTLHeap] instances.
//
// # Discussion
// 
// The method doesn’t remove the residency set from command buffers the
// queue owns with a [Status] property that’s equal to
// [CommandBufferStatusCommitted] or [CommandBufferStatusScheduled].
// 
// See [Simplifying GPU resource management with residency sets] and
// [MTLResidencySet] for more information.
//
// [Simplifying GPU resource management with residency sets]: https://developer.apple.com/documentation/Metal/simplifying-gpu-resource-management-with-residency-sets
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandQueue/removeResidencySet(_:)
func (o MTLCommandQueueObject) RemoveResidencySet(residencySet MTLResidencySet) {
	
	objc.Send[struct{}](o.ID, objc.Sel("removeResidencySet:"), residencySet)
	}
// The GPU device that creates the command queue.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandQueue/device
func (o MTLCommandQueueObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}
// An optional name that can help you identify the command queue.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandQueue/label
func (o MTLCommandQueueObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}
// Applies multiple residency sets to a queue, which Metal applies to the
// queue’s command buffers as you commit them.
//
// residencySets: A C array of residency sets, each of which contains resource allocations,
// such as [MTLBuffer], [MTLTexture], and [MTLHeap] instances.
//
// count: The number of elements in `residencySets`.
//
// # Discussion
// 
// Each command queue can maintain a list of up to 32 different residency
// sets. See [Simplifying GPU resource management with residency sets] and
// [MTLResidencySet] for more information.
//
// [Simplifying GPU resource management with residency sets]: https://developer.apple.com/documentation/Metal/simplifying-gpu-resource-management-with-residency-sets
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandQueue/addResidencySets:count:
func (o MTLCommandQueueObject) AddResidencySetsCount(residencySets []MTLResidencySet, count uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("addResidencySets:count:"), objc.CArray(residencySets), count)
	}
// Removes multiple residency sets from a command queue’s list, which means
// Metal doesn’t apply them to the queue’s command buffers as you commit
// them.
//
// residencySets: A C array of residency sets, each of which contains resource allocations,
// such as [MTLBuffer], [MTLTexture], and [MTLHeap] instances.
//
// count: The number of elements in `residencySets`.
//
// # Discussion
// 
// The method doesn’t remove the residency sets from command buffers the
// queue owns with a [Status] property that’s equal to
// [CommandBufferStatusCommitted] or [CommandBufferStatusScheduled].
// 
// See [Simplifying GPU resource management with residency sets] and
// [MTLResidencySet] for more information.
//
// [Simplifying GPU resource management with residency sets]: https://developer.apple.com/documentation/Metal/simplifying-gpu-resource-management-with-residency-sets
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandQueue/removeResidencySets:count:
func (o MTLCommandQueueObject) RemoveResidencySetsCount(residencySets []MTLResidencySet, count uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("removeResidencySets:count:"), objc.CArray(residencySets), count)
	}

func (o MTLCommandQueueObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}

