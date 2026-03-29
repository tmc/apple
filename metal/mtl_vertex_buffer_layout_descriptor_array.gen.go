// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLVertexBufferLayoutDescriptorArray] class.
var (
	_MTLVertexBufferLayoutDescriptorArrayClass     MTLVertexBufferLayoutDescriptorArrayClass
	_MTLVertexBufferLayoutDescriptorArrayClassOnce sync.Once
)

func getMTLVertexBufferLayoutDescriptorArrayClass() MTLVertexBufferLayoutDescriptorArrayClass {
	_MTLVertexBufferLayoutDescriptorArrayClassOnce.Do(func() {
		_MTLVertexBufferLayoutDescriptorArrayClass = MTLVertexBufferLayoutDescriptorArrayClass{class: objc.GetClass("MTLVertexBufferLayoutDescriptorArray")}
	})
	return _MTLVertexBufferLayoutDescriptorArrayClass
}

// GetMTLVertexBufferLayoutDescriptorArrayClass returns the class object for MTLVertexBufferLayoutDescriptorArray.
func GetMTLVertexBufferLayoutDescriptorArrayClass() MTLVertexBufferLayoutDescriptorArrayClass {
	return getMTLVertexBufferLayoutDescriptorArrayClass()
}

type MTLVertexBufferLayoutDescriptorArrayClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLVertexBufferLayoutDescriptorArrayClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLVertexBufferLayoutDescriptorArrayClass) Alloc() MTLVertexBufferLayoutDescriptorArray {
	rv := objc.Send[MTLVertexBufferLayoutDescriptorArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An array of vertex buffer layout descriptor instances.
//
// # Overview
// 
// An [MTLVertexBufferLayoutDescriptorArray] holds an array of vertex buffer
// layout states. The methods of [MTLVertexBufferLayoutDescriptorArray] set
// the vertex buffer layout state in the array or retrieve the state from the
// array.
//
// # Accessing a specified vertex buffer layout
//
//   - [MTLVertexBufferLayoutDescriptorArray.ObjectAtIndexedSubscript]: Returns the state of the specified vertex buffer layout.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptorArray
type MTLVertexBufferLayoutDescriptorArray struct {
	objectivec.Object
}

// MTLVertexBufferLayoutDescriptorArrayFromID constructs a [MTLVertexBufferLayoutDescriptorArray] from an objc.ID.
//
// An array of vertex buffer layout descriptor instances.
func MTLVertexBufferLayoutDescriptorArrayFromID(id objc.ID) MTLVertexBufferLayoutDescriptorArray {
	return MTLVertexBufferLayoutDescriptorArray{objectivec.Object{ID: id}}
}
// NOTE: MTLVertexBufferLayoutDescriptorArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLVertexBufferLayoutDescriptorArray] class.
//
// # Accessing a specified vertex buffer layout
//
//   - [IMTLVertexBufferLayoutDescriptorArray.ObjectAtIndexedSubscript]: Returns the state of the specified vertex buffer layout.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptorArray
type IMTLVertexBufferLayoutDescriptorArray interface {
	objectivec.IObject

	// Topic: Accessing a specified vertex buffer layout

	// Returns the state of the specified vertex buffer layout.
	ObjectAtIndexedSubscript(index uint) IMTLVertexBufferLayoutDescriptor

	MTLBufferLayoutStrideDynamic() int
	// Sets the state of the specified vertex buffer layout.
	SetObjectAtIndexedSubscript(bufferDesc IMTLVertexBufferLayoutDescriptor, index uint)
}

// Init initializes the instance.
func (v MTLVertexBufferLayoutDescriptorArray) Init() MTLVertexBufferLayoutDescriptorArray {
	rv := objc.Send[MTLVertexBufferLayoutDescriptorArray](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v MTLVertexBufferLayoutDescriptorArray) Autorelease() MTLVertexBufferLayoutDescriptorArray {
	rv := objc.Send[MTLVertexBufferLayoutDescriptorArray](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLVertexBufferLayoutDescriptorArray creates a new MTLVertexBufferLayoutDescriptorArray instance.
func NewMTLVertexBufferLayoutDescriptorArray() MTLVertexBufferLayoutDescriptorArray {
	class := getMTLVertexBufferLayoutDescriptorArrayClass()
	rv := objc.Send[MTLVertexBufferLayoutDescriptorArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the state of the specified vertex buffer layout.
//
// index: A specified index in the array of vertex buffer layouts.
//
// # Return Value
// 
// A descriptor that contains vertex buffer layout state.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptorArray/subscript(_:)
func (v MTLVertexBufferLayoutDescriptorArray) ObjectAtIndexedSubscript(index uint) IMTLVertexBufferLayoutDescriptor {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("objectAtIndexedSubscript:"), index)
	return MTLVertexBufferLayoutDescriptorFromID(rv)
}
// Sets the state of the specified vertex buffer layout.
//
// bufferDesc: A descriptor that contains vertex buffer layout state.
//
// index: An index in the array of vertex buffer layouts.
//
// # Discussion
// 
// If this method is called with `nil` for `bufferDesc` for any legal index,
// the [MTLVertexBufferLayoutDescriptor] object in the array is set to the
// default values.
//
// See: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptorArray/setObject:atIndexedSubscript:
func (v MTLVertexBufferLayoutDescriptorArray) SetObjectAtIndexedSubscript(bufferDesc IMTLVertexBufferLayoutDescriptor, index uint) {
	objc.Send[objc.ID](v.ID, objc.Sel("setObject:atIndexedSubscript:"), bufferDesc, index)
}

// See: https://developer.apple.com/documentation/metal/mtlbufferlayoutstridedynamic
func (v MTLVertexBufferLayoutDescriptorArray) MTLBufferLayoutStrideDynamic() int {
	rv := objc.Send[int](v.ID, objc.Sel("MTLBufferLayoutStrideDynamic"))
	return rv
}

