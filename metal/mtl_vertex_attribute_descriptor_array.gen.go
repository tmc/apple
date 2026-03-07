// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLVertexAttributeDescriptorArray] class.
var (
	_MTLVertexAttributeDescriptorArrayClass     MTLVertexAttributeDescriptorArrayClass
	_MTLVertexAttributeDescriptorArrayClassOnce sync.Once
)

func getMTLVertexAttributeDescriptorArrayClass() MTLVertexAttributeDescriptorArrayClass {
	_MTLVertexAttributeDescriptorArrayClassOnce.Do(func() {
		_MTLVertexAttributeDescriptorArrayClass = MTLVertexAttributeDescriptorArrayClass{class: objc.GetClass("MTLVertexAttributeDescriptorArray")}
	})
	return _MTLVertexAttributeDescriptorArrayClass
}

// GetMTLVertexAttributeDescriptorArrayClass returns the class object for MTLVertexAttributeDescriptorArray.
func GetMTLVertexAttributeDescriptorArrayClass() MTLVertexAttributeDescriptorArrayClass {
	return getMTLVertexAttributeDescriptorArrayClass()
}

type MTLVertexAttributeDescriptorArrayClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLVertexAttributeDescriptorArrayClass) Alloc() MTLVertexAttributeDescriptorArray {
	rv := objc.Send[MTLVertexAttributeDescriptorArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// An array of vertex attribute descriptor instances.
//
// # Overview
// 
// An [MTLVertexAttributeDescriptorArray] instance is an array of instances
// that defines how vertex attribute data is formatted and assigned to an
// index in the attribute argument table. The methods of
// [MTLVertexAttributeDescriptorArray] set or retrieve the attribute
// formatting information from the array.
//
// # Accessing a specified vertex attribute
//
//   - [MTLVertexAttributeDescriptorArray.ObjectAtIndexedSubscript]: Returns the state of the specified vertex attribute.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptorArray
type MTLVertexAttributeDescriptorArray struct {
	objectivec.Object
}

// MTLVertexAttributeDescriptorArrayFromID constructs a [MTLVertexAttributeDescriptorArray] from an objc.ID.
//
// An array of vertex attribute descriptor instances.
func MTLVertexAttributeDescriptorArrayFromID(id objc.ID) MTLVertexAttributeDescriptorArray {
	return MTLVertexAttributeDescriptorArray{objectivec.Object{id}}
}
// NOTE: MTLVertexAttributeDescriptorArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLVertexAttributeDescriptorArray] class.
//
// # Accessing a specified vertex attribute
//
//   - [IMTLVertexAttributeDescriptorArray.ObjectAtIndexedSubscript]: Returns the state of the specified vertex attribute.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptorArray
type IMTLVertexAttributeDescriptorArray interface {
	objectivec.IObject

	// Topic: Accessing a specified vertex attribute

	// Returns the state of the specified vertex attribute.
	ObjectAtIndexedSubscript(index uint) IMTLVertexAttributeDescriptor

	MTLBufferLayoutStrideDynamic() int
	// Sets state for the specified vertex attribute.
	SetObjectAtIndexedSubscript(attributeDesc IMTLVertexAttributeDescriptor, index uint)
}





// Init initializes the instance.
func (v MTLVertexAttributeDescriptorArray) Init() MTLVertexAttributeDescriptorArray {
	rv := objc.Send[MTLVertexAttributeDescriptorArray](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v MTLVertexAttributeDescriptorArray) Autorelease() MTLVertexAttributeDescriptorArray {
	rv := objc.Send[MTLVertexAttributeDescriptorArray](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLVertexAttributeDescriptorArray creates a new MTLVertexAttributeDescriptorArray instance.
func NewMTLVertexAttributeDescriptorArray() MTLVertexAttributeDescriptorArray {
	class := getMTLVertexAttributeDescriptorArrayClass()
	rv := objc.Send[MTLVertexAttributeDescriptorArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Returns the state of the specified vertex attribute.
//
// index: A specified index in the array of vertex attribute states.
//
// # Return Value
// 
// A descriptor that contains the vertex attribute state.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptorArray/subscript(_:)
func (v MTLVertexAttributeDescriptorArray) ObjectAtIndexedSubscript(index uint) IMTLVertexAttributeDescriptor {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("objectAtIndexedSubscript:"), index)
	return MTLVertexAttributeDescriptorFromID(rv)
}

// Sets state for the specified vertex attribute.
//
// attributeDesc: A descriptor that contains vertex attribute state.
//
// index: A specified index in the array of vertex attribute states.
//
// # Discussion
// 
// If this method is called with `nil` for `attributeDesc` for any legal
// `index`, its vertex attribute state is set to the default values.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptorArray/setObject:atIndexedSubscript:
func (v MTLVertexAttributeDescriptorArray) SetObjectAtIndexedSubscript(attributeDesc IMTLVertexAttributeDescriptor, index uint) {
	objc.Send[objc.ID](v.ID, objc.Sel("setObject:atIndexedSubscript:"), attributeDesc, index)
}












// See: https://developer.apple.com/documentation/metal/mtlbufferlayoutstridedynamic
func (v MTLVertexAttributeDescriptorArray) MTLBufferLayoutStrideDynamic() int {
	rv := objc.Send[int](v.ID, objc.Sel("MTLBufferLayoutStrideDynamic"))
	return rv
}























