// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLResidencySetDescriptor] class.
var (
	_MTLResidencySetDescriptorClass     MTLResidencySetDescriptorClass
	_MTLResidencySetDescriptorClassOnce sync.Once
)

func getMTLResidencySetDescriptorClass() MTLResidencySetDescriptorClass {
	_MTLResidencySetDescriptorClassOnce.Do(func() {
		_MTLResidencySetDescriptorClass = MTLResidencySetDescriptorClass{class: objc.GetClass("MTLResidencySetDescriptor")}
	})
	return _MTLResidencySetDescriptorClass
}

// GetMTLResidencySetDescriptorClass returns the class object for MTLResidencySetDescriptor.
func GetMTLResidencySetDescriptorClass() MTLResidencySetDescriptorClass {
	return getMTLResidencySetDescriptorClass()
}

type MTLResidencySetDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLResidencySetDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLResidencySetDescriptorClass) Alloc() MTLResidencySetDescriptor {
	rv := objc.Send[MTLResidencySetDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A configuration that customizes the behavior for a residency set.
//
// # Overview
// 
// Make an [MTLResidencySet] by creating and configuring an
// [MTLResidencySetDescriptor] instance and pass it to the
// [NewResidencySetWithDescriptorError] method of an [MTLDevice] instance.
// 
// See [Simplifying GPU resource management with residency sets] for more
// information.
//
// [Simplifying GPU resource management with residency sets]: https://developer.apple.com/documentation/Metal/simplifying-gpu-resource-management-with-residency-sets
//
// # Configuring the residency set
//
//   - [MTLResidencySetDescriptor.Label]: An optional name that can help you identify a residency set you create with the descriptor.
//   - [MTLResidencySetDescriptor.SetLabel]
//   - [MTLResidencySetDescriptor.InitialCapacity]: The number of allocations a new residency set can store without reallocating memory.
//   - [MTLResidencySetDescriptor.SetInitialCapacity]
//
// See: https://developer.apple.com/documentation/Metal/MTLResidencySetDescriptor
type MTLResidencySetDescriptor struct {
	objectivec.Object
}

// MTLResidencySetDescriptorFromID constructs a [MTLResidencySetDescriptor] from an objc.ID.
//
// A configuration that customizes the behavior for a residency set.
func MTLResidencySetDescriptorFromID(id objc.ID) MTLResidencySetDescriptor {
	return MTLResidencySetDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLResidencySetDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLResidencySetDescriptor] class.
//
// # Configuring the residency set
//
//   - [IMTLResidencySetDescriptor.Label]: An optional name that can help you identify a residency set you create with the descriptor.
//   - [IMTLResidencySetDescriptor.SetLabel]
//   - [IMTLResidencySetDescriptor.InitialCapacity]: The number of allocations a new residency set can store without reallocating memory.
//   - [IMTLResidencySetDescriptor.SetInitialCapacity]
//
// See: https://developer.apple.com/documentation/Metal/MTLResidencySetDescriptor
type IMTLResidencySetDescriptor interface {
	objectivec.IObject

	// Topic: Configuring the residency set

	// An optional name that can help you identify a residency set you create with the descriptor.
	Label() string
	SetLabel(value string)
	// The number of allocations a new residency set can store without reallocating memory.
	InitialCapacity() uint
	SetInitialCapacity(value uint)
}

// Init initializes the instance.
func (r MTLResidencySetDescriptor) Init() MTLResidencySetDescriptor {
	rv := objc.Send[MTLResidencySetDescriptor](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MTLResidencySetDescriptor) Autorelease() MTLResidencySetDescriptor {
	rv := objc.Send[MTLResidencySetDescriptor](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLResidencySetDescriptor creates a new MTLResidencySetDescriptor instance.
func NewMTLResidencySetDescriptor() MTLResidencySetDescriptor {
	class := getMTLResidencySetDescriptorClass()
	rv := objc.Send[MTLResidencySetDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An optional name that can help you identify a residency set you create with
// the descriptor.
//
// # Discussion
// 
// Metal applies the value of this property to the [Label] property of an
// [MTLResidencySet] that you create by passing the descriptor to
// [NewResidencySetWithDescriptorError].
//
// See: https://developer.apple.com/documentation/Metal/MTLResidencySetDescriptor/label
func (r MTLResidencySetDescriptor) Label() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (r MTLResidencySetDescriptor) SetLabel(value string) {
	objc.Send[struct{}](r.ID, objc.Sel("setLabel:"), objc.String(value))
}
// The number of allocations a new residency set can store without
// reallocating memory.
//
// # Discussion
// 
// Reduce the memory reallocations the set needs to make by setting the
// property to a value large enough to hold the allocations you expect. You
// can leave the property at its default value of `0`, which tells Metal to
// give the residency set the standard starting capacity.
//
// See: https://developer.apple.com/documentation/Metal/MTLResidencySetDescriptor/initialCapacity
func (r MTLResidencySetDescriptor) InitialCapacity() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("initialCapacity"))
	return rv
}
func (r MTLResidencySetDescriptor) SetInitialCapacity(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setInitialCapacity:"), value)
}

