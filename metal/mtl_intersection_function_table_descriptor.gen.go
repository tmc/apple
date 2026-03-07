// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLIntersectionFunctionTableDescriptor] class.
var (
	_MTLIntersectionFunctionTableDescriptorClass     MTLIntersectionFunctionTableDescriptorClass
	_MTLIntersectionFunctionTableDescriptorClassOnce sync.Once
)

func getMTLIntersectionFunctionTableDescriptorClass() MTLIntersectionFunctionTableDescriptorClass {
	_MTLIntersectionFunctionTableDescriptorClassOnce.Do(func() {
		_MTLIntersectionFunctionTableDescriptorClass = MTLIntersectionFunctionTableDescriptorClass{class: objc.GetClass("MTLIntersectionFunctionTableDescriptor")}
	})
	return _MTLIntersectionFunctionTableDescriptorClass
}

// GetMTLIntersectionFunctionTableDescriptorClass returns the class object for MTLIntersectionFunctionTableDescriptor.
func GetMTLIntersectionFunctionTableDescriptorClass() MTLIntersectionFunctionTableDescriptorClass {
	return getMTLIntersectionFunctionTableDescriptorClass()
}

type MTLIntersectionFunctionTableDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLIntersectionFunctionTableDescriptorClass) Alloc() MTLIntersectionFunctionTableDescriptor {
	rv := objc.Send[MTLIntersectionFunctionTableDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A specification of how to create an intersection function table.
//
// # Configuring the table’s size
//
//   - [MTLIntersectionFunctionTableDescriptor.FunctionCount]: The number of entries in the intersection function table.
//   - [MTLIntersectionFunctionTableDescriptor.SetFunctionCount]
//
// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTableDescriptor
type MTLIntersectionFunctionTableDescriptor struct {
	objectivec.Object
}

// MTLIntersectionFunctionTableDescriptorFromID constructs a [MTLIntersectionFunctionTableDescriptor] from an objc.ID.
//
// A specification of how to create an intersection function table.
func MTLIntersectionFunctionTableDescriptorFromID(id objc.ID) MTLIntersectionFunctionTableDescriptor {
	return MTLIntersectionFunctionTableDescriptor{objectivec.Object{id}}
}
// NOTE: MTLIntersectionFunctionTableDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLIntersectionFunctionTableDescriptor] class.
//
// # Configuring the table’s size
//
//   - [IMTLIntersectionFunctionTableDescriptor.FunctionCount]: The number of entries in the intersection function table.
//   - [IMTLIntersectionFunctionTableDescriptor.SetFunctionCount]
//
// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTableDescriptor
type IMTLIntersectionFunctionTableDescriptor interface {
	objectivec.IObject

	// Topic: Configuring the table’s size

	// The number of entries in the intersection function table.
	FunctionCount() uint
	SetFunctionCount(value uint)
}





// Init initializes the instance.
func (i MTLIntersectionFunctionTableDescriptor) Init() MTLIntersectionFunctionTableDescriptor {
	rv := objc.Send[MTLIntersectionFunctionTableDescriptor](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MTLIntersectionFunctionTableDescriptor) Autorelease() MTLIntersectionFunctionTableDescriptor {
	rv := objc.Send[MTLIntersectionFunctionTableDescriptor](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLIntersectionFunctionTableDescriptor creates a new MTLIntersectionFunctionTableDescriptor instance.
func NewMTLIntersectionFunctionTableDescriptor() MTLIntersectionFunctionTableDescriptor {
	class := getMTLIntersectionFunctionTableDescriptorClass()
	rv := objc.Send[MTLIntersectionFunctionTableDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}














// Creates an intersection function table descriptor.
//
// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTableDescriptor/intersectionFunctionTableDescriptor
func (_MTLIntersectionFunctionTableDescriptorClass MTLIntersectionFunctionTableDescriptorClass) IntersectionFunctionTableDescriptor() MTLIntersectionFunctionTableDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MTLIntersectionFunctionTableDescriptorClass.class), objc.Sel("intersectionFunctionTableDescriptor"))
	return MTLIntersectionFunctionTableDescriptorFromID(rv)
}








// The number of entries in the intersection function table.
//
// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTableDescriptor/functionCount
func (i MTLIntersectionFunctionTableDescriptor) FunctionCount() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("functionCount"))
	return rv
}
func (i MTLIntersectionFunctionTableDescriptor) SetFunctionCount(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setFunctionCount:"), value)
}
























