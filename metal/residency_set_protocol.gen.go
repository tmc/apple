// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A collection of resource allocations that can move in and out of resident memory.
//
// See: https://developer.apple.com/documentation/Metal/MTLResidencySet
type MTLResidencySet interface {
	objectivec.IObject

	// Stages a single resource to join the residency set’s list of allocations.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/addAllocation(_:)
	AddAllocation(allocation MTLAllocation)

	// Stages all the resources in the residency set to leave its list of allocations.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/removeAllAllocations()
	RemoveAllAllocations()

	// Stages a single resource to leave the residency set’s list of allocations.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/removeAllocation(_:)
	RemoveAllocation(allocation MTLAllocation)

	// Applies any pending additions to and removals from the residency set.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/commit()
	Commit()

	// Tells Metal to do as much preparatory work as it can, with the system’s current conditions, to make the set’s resource allocations resident.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/requestResidency()
	RequestResidency()

	// Informs Metal that the residency set’s allocations no longer need to be resident, and that it can reuse the memory for other allocations.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/endResidency()
	EndResidency()

	// An optional name that can help you identify the residency set.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/label
	Label() string

	// The Metal device that owns the residency set.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/device
	Device() MTLDevice

	// Returns a Boolean value that indicates whether the residency set contains a specific resource allocation.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/containsAllocation(_:)
	ContainsAllocation(anAllocation MTLAllocation) bool

	// The residency set’s current list of resource allocations.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/allAllocations
	AllAllocations() []objectivec.IObject

	// The number of resource allocations in the residency set.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/allocationCount
	AllocationCount() uint

	// The amount of resident memory, in bytes, the residency set’s resource allocations consume.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/allocatedSize
	AllocatedSize() uint64

	// Stages multiple resources to join the residency set’s list of allocations.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/addAllocations:count:
	AddAllocationsCount(allocations []MTLAllocation, count uint)

	// Stages multiple resources to leave the residency set’s list of allocations.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/removeAllocations:count:
	RemoveAllocationsCount(allocations []MTLAllocation, count uint)
}

// MTLResidencySetObject wraps an existing Objective-C object that conforms to the MTLResidencySet protocol.
type MTLResidencySetObject struct {
	objectivec.Object
}
func (o MTLResidencySetObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLResidencySetObjectFromID constructs a [MTLResidencySetObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLResidencySetObjectFromID(id objc.ID) MTLResidencySetObject {
	return MTLResidencySetObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Stages a single resource to join the residency set’s list of allocations.
//
// allocation: A resource allocation, such as an [MTLBuffer], [MTLTexture], or [MTLHeap].
//
// # Discussion
// 
// Finalize the inclusion of these resource allocations, and all other changes
// you stage, by calling a residency set’s [Commit] method.
//
// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/addAllocation(_:)
func (o MTLResidencySetObject) AddAllocation(allocation MTLAllocation) {
	objc.Send[struct{}](o.ID, objc.Sel("addAllocation:"), allocation)
	}
// Stages all the resources in the residency set to leave its list of
// allocations.
//
// # Discussion
// 
// Finalize the removal of these resource allocations, and all others changes
// you stage, by calling a residency set’s [Commit] method.
//
// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/removeAllAllocations()
func (o MTLResidencySetObject) RemoveAllAllocations() {
	objc.Send[struct{}](o.ID, objc.Sel("removeAllAllocations"))
	}
// Stages a single resource to leave the residency set’s list of
// allocations.
//
// allocation: A resource allocation, such as an [MTLBuffer], [MTLTexture], or [MTLHeap].
//
// # Discussion
// 
// Finalize the removal of these resource allocations, and all others changes
// you stage, by calling a residency set’s [Commit] method.
//
// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/removeAllocation(_:)
func (o MTLResidencySetObject) RemoveAllocation(allocation MTLAllocation) {
	objc.Send[struct{}](o.ID, objc.Sel("removeAllocation:"), allocation)
	}
// Applies any pending additions to and removals from the residency set.
//
// # Discussion
// 
// Call the method when have no other changes to stage, such as with
// [AddAllocation], [RemoveAllocation], and their sibling methods.
//
// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/commit()
func (o MTLResidencySetObject) Commit() {
	objc.Send[struct{}](o.ID, objc.Sel("commit"))
	}
// Tells Metal to do as much preparatory work as it can, with the system’s
// current conditions, to make the set’s resource allocations resident.
//
// # Discussion
// 
// Call the method anytime after calling a residency set’s [Commit] method,
// ideally well before calling the [Commit] method of any [MTLCommandBuffer]
// that uses it.
// 
// The method may postpone some of the necessary steps to make resources
// resident in scenarios where other apps concurrently need resources in
// residency.
//
// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/requestResidency()
func (o MTLResidencySetObject) RequestResidency() {
	objc.Send[struct{}](o.ID, objc.Sel("requestResidency"))
	}
// Informs Metal that the residency set’s allocations no longer need to be
// resident, and that it can reuse the memory for other allocations.
//
// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/endResidency()
func (o MTLResidencySetObject) EndResidency() {
	objc.Send[struct{}](o.ID, objc.Sel("endResidency"))
	}
// An optional name that can help you identify the residency set.
//
// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/label
func (o MTLResidencySetObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}
// The Metal device that owns the residency set.
//
// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/device
func (o MTLResidencySetObject) Device() MTLDevice {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}
// Returns a Boolean value that indicates whether the residency set contains a
// specific resource allocation.
//
// anAllocation: A resource allocation, such as an [MTLBuffer], [MTLTexture], or [MTLHeap].
//
// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/containsAllocation(_:)
func (o MTLResidencySetObject) ContainsAllocation(anAllocation MTLAllocation) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("containsAllocation:"), anAllocation)
	return rv
	}
// The residency set’s current list of resource allocations.
//
// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/allAllocations
func (o MTLResidencySetObject) AllAllocations() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("allAllocations"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
	}
// The number of resource allocations in the residency set.
//
// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/allocationCount
func (o MTLResidencySetObject) AllocationCount() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("allocationCount"))
	return rv
	}
// The amount of resident memory, in bytes, the residency set’s resource
// allocations consume.
//
// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/allocatedSize
func (o MTLResidencySetObject) AllocatedSize() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("allocatedSize"))
	return rv
	}
// Stages multiple resources to join the residency set’s list of
// allocations.
//
// allocations: A C array of resource allocations, whose elements can be an arbitrarily mix
// of [MTLBuffer], [MTLTexture], and [MTLHeap] instances.
//
// count: The number of elements in `allocations`.
//
// # Discussion
// 
// Finalize the inclusion of these resource allocations, and all other changes
// you stage, by calling a residency set’s [Commit] method.
//
// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/addAllocations:count:
func (o MTLResidencySetObject) AddAllocationsCount(allocations []MTLAllocation, count uint) {
	objc.Send[struct{}](o.ID, objc.Sel("addAllocations:count:"), objc.CArray(allocations), count)
	}
// Stages multiple resources to leave the residency set’s list of
// allocations.
//
// allocations: A C array of resource allocations, whose elements can be an arbitrarily mix
// of [MTLBuffer], [MTLTexture], and [MTLHeap] instances.
//
// count: The number of elements in `allocations`.
//
// # Discussion
// 
// Finalize the removal of these resource allocations, and all other changes
// you stage, by calling a residency set’s [Commit] method.
//
// See: https://developer.apple.com/documentation/Metal/MTLResidencySet/removeAllocations:count:
func (o MTLResidencySetObject) RemoveAllocationsCount(allocations []MTLAllocation, count uint) {
	objc.Send[struct{}](o.ID, objc.Sel("removeAllocations:count:"), objc.CArray(allocations), count)
	}

