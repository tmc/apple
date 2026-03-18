// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4CommandAllocatorDescriptor] class.
var (
	_MTL4CommandAllocatorDescriptorClass     MTL4CommandAllocatorDescriptorClass
	_MTL4CommandAllocatorDescriptorClassOnce sync.Once
)

func getMTL4CommandAllocatorDescriptorClass() MTL4CommandAllocatorDescriptorClass {
	_MTL4CommandAllocatorDescriptorClassOnce.Do(func() {
		_MTL4CommandAllocatorDescriptorClass = MTL4CommandAllocatorDescriptorClass{class: objc.GetClass("MTL4CommandAllocatorDescriptor")}
	})
	return _MTL4CommandAllocatorDescriptorClass
}

// GetMTL4CommandAllocatorDescriptorClass returns the class object for MTL4CommandAllocatorDescriptor.
func GetMTL4CommandAllocatorDescriptorClass() MTL4CommandAllocatorDescriptorClass {
	return getMTL4CommandAllocatorDescriptorClass()
}

type MTL4CommandAllocatorDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4CommandAllocatorDescriptorClass) Alloc() MTL4CommandAllocatorDescriptor {
	rv := objc.Send[MTL4CommandAllocatorDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Groups together parameters for creating a command allocator.
//
// # Instance Properties
//
//   - [MTL4CommandAllocatorDescriptor.Label]: An optional label you can assign to the command allocator to aid debugging.
//   - [MTL4CommandAllocatorDescriptor.SetLabel]
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandAllocatorDescriptor
type MTL4CommandAllocatorDescriptor struct {
	objectivec.Object
}

// MTL4CommandAllocatorDescriptorFromID constructs a [MTL4CommandAllocatorDescriptor] from an objc.ID.
//
// Groups together parameters for creating a command allocator.
func MTL4CommandAllocatorDescriptorFromID(id objc.ID) MTL4CommandAllocatorDescriptor {
	return MTL4CommandAllocatorDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTL4CommandAllocatorDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4CommandAllocatorDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4CommandAllocatorDescriptor.Label]: An optional label you can assign to the command allocator to aid debugging.
//   - [IMTL4CommandAllocatorDescriptor.SetLabel]
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandAllocatorDescriptor
type IMTL4CommandAllocatorDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	// An optional label you can assign to the command allocator to aid debugging.
	Label() string
	SetLabel(value string)

	MTL4CommandQueueErrorDomain() string
}

// Init initializes the instance.
func (m MTL4CommandAllocatorDescriptor) Init() MTL4CommandAllocatorDescriptor {
	rv := objc.Send[MTL4CommandAllocatorDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4CommandAllocatorDescriptor) Autorelease() MTL4CommandAllocatorDescriptor {
	rv := objc.Send[MTL4CommandAllocatorDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4CommandAllocatorDescriptor creates a new MTL4CommandAllocatorDescriptor instance.
func NewMTL4CommandAllocatorDescriptor() MTL4CommandAllocatorDescriptor {
	class := getMTL4CommandAllocatorDescriptorClass()
	rv := objc.Send[MTL4CommandAllocatorDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An optional label you can assign to the command allocator to aid debugging.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandAllocatorDescriptor/label
func (m MTL4CommandAllocatorDescriptor) Label() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (m MTL4CommandAllocatorDescriptor) SetLabel(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setLabel:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/metal/mtl4commandqueueerrordomain
func (m MTL4CommandAllocatorDescriptor) MTL4CommandQueueErrorDomain() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("MTL4CommandQueueErrorDomain"))
	return foundation.NSStringFromID(rv).String()
}

