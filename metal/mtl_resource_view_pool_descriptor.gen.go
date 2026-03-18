// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLResourceViewPoolDescriptor] class.
var (
	_MTLResourceViewPoolDescriptorClass     MTLResourceViewPoolDescriptorClass
	_MTLResourceViewPoolDescriptorClassOnce sync.Once
)

func getMTLResourceViewPoolDescriptorClass() MTLResourceViewPoolDescriptorClass {
	_MTLResourceViewPoolDescriptorClassOnce.Do(func() {
		_MTLResourceViewPoolDescriptorClass = MTLResourceViewPoolDescriptorClass{class: objc.GetClass("MTLResourceViewPoolDescriptor")}
	})
	return _MTLResourceViewPoolDescriptorClass
}

// GetMTLResourceViewPoolDescriptorClass returns the class object for MTLResourceViewPoolDescriptor.
func GetMTLResourceViewPoolDescriptorClass() MTLResourceViewPoolDescriptorClass {
	return getMTLResourceViewPoolDescriptorClass()
}

type MTLResourceViewPoolDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLResourceViewPoolDescriptorClass) Alloc() MTLResourceViewPoolDescriptor {
	rv := objc.Send[MTLResourceViewPoolDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Provides parameters for creating a resource view pool.
//
// # Instance Properties
//
//   - [MTLResourceViewPoolDescriptor.Label]: Assigns an optional label you to the resource view pool for debugging purposes.
//   - [MTLResourceViewPoolDescriptor.SetLabel]
//   - [MTLResourceViewPoolDescriptor.ResourceViewCount]: Configures the number of resource views with which Metal creates the resource view pool.
//   - [MTLResourceViewPoolDescriptor.SetResourceViewCount]
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceViewPoolDescriptor
type MTLResourceViewPoolDescriptor struct {
	objectivec.Object
}

// MTLResourceViewPoolDescriptorFromID constructs a [MTLResourceViewPoolDescriptor] from an objc.ID.
//
// Provides parameters for creating a resource view pool.
func MTLResourceViewPoolDescriptorFromID(id objc.ID) MTLResourceViewPoolDescriptor {
	return MTLResourceViewPoolDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLResourceViewPoolDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLResourceViewPoolDescriptor] class.
//
// # Instance Properties
//
//   - [IMTLResourceViewPoolDescriptor.Label]: Assigns an optional label you to the resource view pool for debugging purposes.
//   - [IMTLResourceViewPoolDescriptor.SetLabel]
//   - [IMTLResourceViewPoolDescriptor.ResourceViewCount]: Configures the number of resource views with which Metal creates the resource view pool.
//   - [IMTLResourceViewPoolDescriptor.SetResourceViewCount]
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceViewPoolDescriptor
type IMTLResourceViewPoolDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Assigns an optional label you to the resource view pool for debugging purposes.
	Label() string
	SetLabel(value string)
	// Configures the number of resource views with which Metal creates the resource view pool.
	ResourceViewCount() uint
	SetResourceViewCount(value uint)
}

// Init initializes the instance.
func (r MTLResourceViewPoolDescriptor) Init() MTLResourceViewPoolDescriptor {
	rv := objc.Send[MTLResourceViewPoolDescriptor](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MTLResourceViewPoolDescriptor) Autorelease() MTLResourceViewPoolDescriptor {
	rv := objc.Send[MTLResourceViewPoolDescriptor](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLResourceViewPoolDescriptor creates a new MTLResourceViewPoolDescriptor instance.
func NewMTLResourceViewPoolDescriptor() MTLResourceViewPoolDescriptor {
	class := getMTLResourceViewPoolDescriptorClass()
	rv := objc.Send[MTLResourceViewPoolDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Assigns an optional label you to the resource view pool for debugging
// purposes.
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceViewPoolDescriptor/label
func (r MTLResourceViewPoolDescriptor) Label() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (r MTLResourceViewPoolDescriptor) SetLabel(value string) {
	objc.Send[struct{}](r.ID, objc.Sel("setLabel:"), objc.String(value))
}

// Configures the number of resource views with which Metal creates the
// resource view pool.
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceViewPoolDescriptor/resourceViewCount
func (r MTLResourceViewPoolDescriptor) ResourceViewCount() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("resourceViewCount"))
	return rv
}
func (r MTLResourceViewPoolDescriptor) SetResourceViewCount(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setResourceViewCount:"), value)
}

