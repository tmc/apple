// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Represents an opaque, driver-controlled section of memory that can store GPU counter data.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CounterHeap
type MTL4CounterHeap interface {
	objectivec.IObject

	// Queries the number of entries in the heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CounterHeap/count
	Count() uint

	// Assigns a label for later inspection or visualization.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CounterHeap/label
	Label() string

	// Queries the type of the heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CounterHeap/type
	Type() MTL4CounterHeapType

	// Invalidates a range of entries in this counter heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CounterHeap/invalidateCounterRange:
	InvalidateCounterRange(range_ foundation.NSRange)

	// Resolves heap data on the CPU timeline.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CounterHeap/resolveCounterRange:
	ResolveCounterRange(range_ foundation.NSRange) foundation.INSData

	// Assigns a label for later inspection or visualization.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CounterHeap/label
	SetLabel(value string)
}

// MTL4CounterHeapObject wraps an existing Objective-C object that conforms to the MTL4CounterHeap protocol.
type MTL4CounterHeapObject struct {
	objectivec.Object
}

func (o MTL4CounterHeapObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTL4CounterHeapObjectFromID constructs a [MTL4CounterHeapObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTL4CounterHeapObjectFromID(id objc.ID) MTL4CounterHeapObject {
	return MTL4CounterHeapObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Queries the number of entries in the heap.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CounterHeap/count
func (o MTL4CounterHeapObject) Count() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("count"))
	return rv
}

// Assigns a label for later inspection or visualization.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CounterHeap/label
func (o MTL4CounterHeapObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

// Queries the type of the heap.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CounterHeap/type
func (o MTL4CounterHeapObject) Type() MTL4CounterHeapType {
	rv := objc.Send[MTL4CounterHeapType](o.ID, objc.Sel("type"))
	return rv
}

// Invalidates a range of entries in this counter heap.
//
// range: A heap index range to invalidate.
//
// # Discussion
//
// The effect of this call is immediate on the CPU timeline. You are
// responsible for ensuring that this counter heap is not currently in use on
// the GPU.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CounterHeap/invalidateCounterRange:
func (o MTL4CounterHeapObject) InvalidateCounterRange(range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("invalidateCounterRange:"), range_)
}

// Resolves heap data on the CPU timeline.
//
// range: The range in the heap to resolve.
//
// # Discussion
//
// This method resolves heap data in the CPU timeline. Your app needs to
// ensure the GPU work has completed in order to retrieve the data correctly.
// You can alternatively resolve the heap data in the GPU timeline by calling
// [ResolveCounterHeapWithRangeIntoBufferWaitFenceUpdateFence].
//
// - Returns a newly allocated autoreleased NSData containing tightly packed
// resolved heap counter values.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CounterHeap/resolveCounterRange:
func (o MTL4CounterHeapObject) ResolveCounterRange(range_ foundation.NSRange) foundation.INSData {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("resolveCounterRange:"), range_)
	return foundation.NSDataFromID(rv)
}

// Assigns a label for later inspection or visualization.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CounterHeap/label
func (o MTL4CounterHeapObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}
