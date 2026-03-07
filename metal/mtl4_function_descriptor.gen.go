// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4FunctionDescriptor] class.
var (
	_MTL4FunctionDescriptorClass     MTL4FunctionDescriptorClass
	_MTL4FunctionDescriptorClassOnce sync.Once
)

func getMTL4FunctionDescriptorClass() MTL4FunctionDescriptorClass {
	_MTL4FunctionDescriptorClassOnce.Do(func() {
		_MTL4FunctionDescriptorClass = MTL4FunctionDescriptorClass{class: objc.GetClass("MTL4FunctionDescriptor")}
	})
	return _MTL4FunctionDescriptorClass
}

// GetMTL4FunctionDescriptorClass returns the class object for MTL4FunctionDescriptor.
func GetMTL4FunctionDescriptorClass() MTL4FunctionDescriptorClass {
	return getMTL4FunctionDescriptorClass()
}

type MTL4FunctionDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4FunctionDescriptorClass) Alloc() MTL4FunctionDescriptor {
	rv := objc.Send[MTL4FunctionDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// Base interface for describing a Metal 4 shader function.
//
// See: https://developer.apple.com/documentation/Metal/MTL4FunctionDescriptor
type MTL4FunctionDescriptor struct {
	objectivec.Object
}

// MTL4FunctionDescriptorFromID constructs a [MTL4FunctionDescriptor] from an objc.ID.
//
// Base interface for describing a Metal 4 shader function.
func MTL4FunctionDescriptorFromID(id objc.ID) MTL4FunctionDescriptor {
	return MTL4FunctionDescriptor{objectivec.Object{id}}
}
// NOTE: MTL4FunctionDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTL4FunctionDescriptor] class.
//
// See: https://developer.apple.com/documentation/Metal/MTL4FunctionDescriptor
type IMTL4FunctionDescriptor interface {
	objectivec.IObject
}





// Init initializes the instance.
func (m MTL4FunctionDescriptor) Init() MTL4FunctionDescriptor {
	rv := objc.Send[MTL4FunctionDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4FunctionDescriptor) Autorelease() MTL4FunctionDescriptor {
	rv := objc.Send[MTL4FunctionDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4FunctionDescriptor creates a new MTL4FunctionDescriptor instance.
func NewMTL4FunctionDescriptor() MTL4FunctionDescriptor {
	class := getMTL4FunctionDescriptorClass()
	rv := objc.Send[MTL4FunctionDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}










































