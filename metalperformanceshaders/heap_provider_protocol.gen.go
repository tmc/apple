// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MPSHeapProvider protocol.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSHeapProvider
type MPSHeapProvider interface {
	objectivec.IObject

	// NewHeapWithDescriptor protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSHeapProvider/newHeap(with:)
	NewHeapWithDescriptor(descriptor metal.MTLHeapDescriptor) metal.MTLHeap
}

// MPSHeapProviderObject wraps an existing Objective-C object that conforms to the MPSHeapProvider protocol.
type MPSHeapProviderObject struct {
	objectivec.Object
}

func (o MPSHeapProviderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MPSHeapProviderObjectFromID constructs a [MPSHeapProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MPSHeapProviderObjectFromID(id objc.ID) MPSHeapProviderObject {
	return MPSHeapProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSHeapProvider/newHeap(with:)
func (o MPSHeapProviderObject) NewHeapWithDescriptor(descriptor metal.MTLHeapDescriptor) metal.MTLHeap {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newHeapWithDescriptor:"), descriptor)
	return metal.MTLHeapObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSHeapProvider/retire(_:cacheDelay:)
func (o MPSHeapProviderObject) RetireHeapCacheDelay(heap metal.MTLHeap, seconds float64) {
	objc.Send[struct{}](o.ID, objc.Sel("retireHeap:cacheDelay:"), heap, seconds)
}
