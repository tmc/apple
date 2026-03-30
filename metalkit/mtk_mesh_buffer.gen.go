// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTKMeshBuffer] class.
var (
	_MTKMeshBufferClass     MTKMeshBufferClass
	_MTKMeshBufferClassOnce sync.Once
)

func getMTKMeshBufferClass() MTKMeshBufferClass {
	_MTKMeshBufferClassOnce.Do(func() {
		_MTKMeshBufferClass = MTKMeshBufferClass{class: objc.GetClass("MTKMeshBuffer")}
	})
	return _MTKMeshBufferClass
}

// GetMTKMeshBufferClass returns the class object for MTKMeshBuffer.
func GetMTKMeshBufferClass() MTKMeshBufferClass {
	return getMTKMeshBufferClass()
}

type MTKMeshBufferClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTKMeshBufferClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTKMeshBufferClass) Alloc() MTKMeshBuffer {
	rv := objc.Send[MTKMeshBuffer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A buffer that backs the vertex data of a Model I/O mesh, suitable for use
// in a Metal app.
//
// # Originating Objects
//
//   - [MTKMeshBuffer.Allocator]: The allocator object used to create this mesh buffer.
//   - [MTKMeshBuffer.Type]: The type of data contained in the originating Model I/O buffer.
//
// # Metal Buffer Properties
//
//   - [MTKMeshBuffer.Buffer]: The Metal buffer backing all vertex and index data.
//   - [MTKMeshBuffer.Length]: The logical size of the Metal buffer, in bytes.
//   - [MTKMeshBuffer.Offset]: The byte offset of the data within the Metal buffer.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer
type MTKMeshBuffer struct {
	objectivec.Object
}

// MTKMeshBufferFromID constructs a [MTKMeshBuffer] from an objc.ID.
//
// A buffer that backs the vertex data of a Model I/O mesh, suitable for use
// in a Metal app.
func MTKMeshBufferFromID(id objc.ID) MTKMeshBuffer {
	return MTKMeshBuffer{objectivec.Object{ID: id}}
}

// NOTE: MTKMeshBuffer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTKMeshBuffer] class.
//
// # Originating Objects
//
//   - [IMTKMeshBuffer.Allocator]: The allocator object used to create this mesh buffer.
//   - [IMTKMeshBuffer.Type]: The type of data contained in the originating Model I/O buffer.
//
// # Metal Buffer Properties
//
//   - [IMTKMeshBuffer.Buffer]: The Metal buffer backing all vertex and index data.
//   - [IMTKMeshBuffer.Length]: The logical size of the Metal buffer, in bytes.
//   - [IMTKMeshBuffer.Offset]: The byte offset of the data within the Metal buffer.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer
type IMTKMeshBuffer interface {
	objectivec.IObject

	// Topic: Originating Objects

	// The allocator object used to create this mesh buffer.
	Allocator() IMTKMeshBufferAllocator
	// The type of data contained in the originating Model I/O buffer.
	Type() objectivec.IObject

	// Topic: Metal Buffer Properties

	// The Metal buffer backing all vertex and index data.
	Buffer() metal.MTLBuffer
	// The logical size of the Metal buffer, in bytes.
	Length() uint
	// The byte offset of the data within the Metal buffer.
	Offset() uint
}

// Init initializes the instance.
func (m MTKMeshBuffer) Init() MTKMeshBuffer {
	rv := objc.Send[MTKMeshBuffer](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTKMeshBuffer) Autorelease() MTKMeshBuffer {
	rv := objc.Send[MTKMeshBuffer](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTKMeshBuffer creates a new MTKMeshBuffer instance.
func NewMTKMeshBuffer() MTKMeshBuffer {
	class := getMTKMeshBufferClass()
	rv := objc.Send[MTKMeshBuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The allocator object used to create this mesh buffer.
//
// # Discussion
//
// The allocator uses Model I/O for copy and re-layout operations, such as
// when a new vertex descriptor is applied to an existing vertex buffer.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer/allocator
func (m MTKMeshBuffer) Allocator() IMTKMeshBufferAllocator {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("allocator"))
	return MTKMeshBufferAllocatorFromID(objc.ID(rv))
}

// The type of data contained in the originating Model I/O buffer.
//
// # Discussion
//
// A [MDLMeshBuffer] object can contain Model I/O mesh vertex data or submesh
// index data.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer/type
//
// [MDLMeshBuffer]: https://developer.apple.com/documentation/ModelIO/MDLMeshBuffer
func (m MTKMeshBuffer) Type() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("type"))
	return objectivec.Object{ID: rv}
}

// The Metal buffer backing all vertex and index data.
//
// # Discussion
//
// Many [MTKMeshBuffer] objects may reference the same [MTLBuffer] object, in
// which case each [MTKMeshBuffer] object will have its own unique [Offset]
// value.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer/buffer
//
// [MTLBuffer]: https://developer.apple.com/documentation/Metal/MTLBuffer
func (m MTKMeshBuffer) Buffer() metal.MTLBuffer {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("buffer"))
	return metal.MTLBufferObjectFromID(rv)
}

// The logical size of the Metal buffer, in bytes.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer/length
func (m MTKMeshBuffer) Length() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("length"))
	return rv
}

// The byte offset of the data within the Metal buffer.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMeshBuffer/offset
func (m MTKMeshBuffer) Offset() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("offset"))
	return rv
}
