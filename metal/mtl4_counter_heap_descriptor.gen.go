// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4CounterHeapDescriptor] class.
var (
	_MTL4CounterHeapDescriptorClass     MTL4CounterHeapDescriptorClass
	_MTL4CounterHeapDescriptorClassOnce sync.Once
)

func getMTL4CounterHeapDescriptorClass() MTL4CounterHeapDescriptorClass {
	_MTL4CounterHeapDescriptorClassOnce.Do(func() {
		_MTL4CounterHeapDescriptorClass = MTL4CounterHeapDescriptorClass{class: objc.GetClass("MTL4CounterHeapDescriptor")}
	})
	return _MTL4CounterHeapDescriptorClass
}

// GetMTL4CounterHeapDescriptorClass returns the class object for MTL4CounterHeapDescriptor.
func GetMTL4CounterHeapDescriptorClass() MTL4CounterHeapDescriptorClass {
	return getMTL4CounterHeapDescriptorClass()
}

type MTL4CounterHeapDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4CounterHeapDescriptorClass) Alloc() MTL4CounterHeapDescriptor {
	rv := objc.Send[MTL4CounterHeapDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// Groups together parameters for configuring a counter heap object at
// creation time.
//
// # Instance Properties
//
//   - [MTL4CounterHeapDescriptor.Count]: Assigns the number of entries in the heap.
//   - [MTL4CounterHeapDescriptor.SetCount]
//   - [MTL4CounterHeapDescriptor.Type]: Assigns the type of data that the heap contains.
//   - [MTL4CounterHeapDescriptor.SetType]
//
// See: https://developer.apple.com/documentation/Metal/MTL4CounterHeapDescriptor
type MTL4CounterHeapDescriptor struct {
	objectivec.Object
}

// MTL4CounterHeapDescriptorFromID constructs a [MTL4CounterHeapDescriptor] from an objc.ID.
//
// Groups together parameters for configuring a counter heap object at
// creation time.
func MTL4CounterHeapDescriptorFromID(id objc.ID) MTL4CounterHeapDescriptor {
	return MTL4CounterHeapDescriptor{objectivec.Object{id}}
}
// NOTE: MTL4CounterHeapDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTL4CounterHeapDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4CounterHeapDescriptor.Count]: Assigns the number of entries in the heap.
//   - [IMTL4CounterHeapDescriptor.SetCount]
//   - [IMTL4CounterHeapDescriptor.Type]: Assigns the type of data that the heap contains.
//   - [IMTL4CounterHeapDescriptor.SetType]
//
// See: https://developer.apple.com/documentation/Metal/MTL4CounterHeapDescriptor
type IMTL4CounterHeapDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Assigns the number of entries in the heap.
	Count() uint
	SetCount(value uint)
	// Assigns the type of data that the heap contains.
	Type() MTL4CounterHeapType
	SetType(value MTL4CounterHeapType)

	MTL4CommandQueueErrorDomain() string
}





// Init initializes the instance.
func (m MTL4CounterHeapDescriptor) Init() MTL4CounterHeapDescriptor {
	rv := objc.Send[MTL4CounterHeapDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4CounterHeapDescriptor) Autorelease() MTL4CounterHeapDescriptor {
	rv := objc.Send[MTL4CounterHeapDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4CounterHeapDescriptor creates a new MTL4CounterHeapDescriptor instance.
func NewMTL4CounterHeapDescriptor() MTL4CounterHeapDescriptor {
	class := getMTL4CounterHeapDescriptorClass()
	rv := objc.Send[MTL4CounterHeapDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// Assigns the number of entries in the heap.
//
// # Discussion
// 
// Each entry represents one item in the heap. The size of the individual
// entries depends on the heap type.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CounterHeapDescriptor/count
func (m MTL4CounterHeapDescriptor) Count() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("count"))
	return rv
}
func (m MTL4CounterHeapDescriptor) SetCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setCount:"), value)
}



// Assigns the type of data that the heap contains.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CounterHeapDescriptor/type
func (m MTL4CounterHeapDescriptor) Type() MTL4CounterHeapType {
	rv := objc.Send[MTL4CounterHeapType](m.ID, objc.Sel("type"))
	return MTL4CounterHeapType(rv)
}
func (m MTL4CounterHeapDescriptor) SetType(value MTL4CounterHeapType) {
	objc.Send[struct{}](m.ID, objc.Sel("setType:"), value)
}



// See: https://developer.apple.com/documentation/metal/mtl4commandqueueerrordomain
func (m MTL4CounterHeapDescriptor) MTL4CommandQueueErrorDomain() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("MTL4CommandQueueErrorDomain"))
	return foundation.NSStringFromID(rv).String()
}
























