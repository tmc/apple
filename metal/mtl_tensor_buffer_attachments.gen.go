// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLTensorBufferAttachments] class.
var (
	_MTLTensorBufferAttachmentsClass     MTLTensorBufferAttachmentsClass
	_MTLTensorBufferAttachmentsClassOnce sync.Once
)

func getMTLTensorBufferAttachmentsClass() MTLTensorBufferAttachmentsClass {
	_MTLTensorBufferAttachmentsClassOnce.Do(func() {
		_MTLTensorBufferAttachmentsClass = MTLTensorBufferAttachmentsClass{class: objc.GetClass("MTLTensorBufferAttachments")}
	})
	return _MTLTensorBufferAttachmentsClass
}

// GetMTLTensorBufferAttachmentsClass returns the class object for MTLTensorBufferAttachments.
func GetMTLTensorBufferAttachmentsClass() MTLTensorBufferAttachmentsClass {
	return getMTLTensorBufferAttachmentsClass()
}

type MTLTensorBufferAttachmentsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLTensorBufferAttachmentsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLTensorBufferAttachmentsClass) Alloc() MTLTensorBufferAttachments {
	rv := objc.Send[MTLTensorBufferAttachments](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that associates each plane of a tensor with a buffer and byte
// offset for buffer-backed tensor creation.
//
// # Instance Methods
//
//   - [MTLTensorBufferAttachments.BufferForPlane]: Returns the buffer backing the given plane, or `nil` if none has been set.
//   - [MTLTensorBufferAttachments.OffsetForPlane]: Returns the byte offset into the buffer for the given plane.
//   - [MTLTensorBufferAttachments.Reset]: Empties the container of all its elements.
//   - [MTLTensorBufferAttachments.SetBufferOffsetForPlane]: Sets the buffer and byte offset to use as backing storage for the given plane.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorBufferAttachments
type MTLTensorBufferAttachments struct {
	objectivec.Object
}

// MTLTensorBufferAttachmentsFromID constructs a [MTLTensorBufferAttachments] from an objc.ID.
//
// An object that associates each plane of a tensor with a buffer and byte
// offset for buffer-backed tensor creation.
func MTLTensorBufferAttachmentsFromID(id objc.ID) MTLTensorBufferAttachments {
	return MTLTensorBufferAttachments{objectivec.Object{ID: id}}
}

// NOTE: MTLTensorBufferAttachments adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLTensorBufferAttachments] class.
//
// # Instance Methods
//
//   - [IMTLTensorBufferAttachments.BufferForPlane]: Returns the buffer backing the given plane, or `nil` if none has been set.
//   - [IMTLTensorBufferAttachments.OffsetForPlane]: Returns the byte offset into the buffer for the given plane.
//   - [IMTLTensorBufferAttachments.Reset]: Empties the container of all its elements.
//   - [IMTLTensorBufferAttachments.SetBufferOffsetForPlane]: Sets the buffer and byte offset to use as backing storage for the given plane.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorBufferAttachments
type IMTLTensorBufferAttachments interface {
	objectivec.IObject

	// Topic: Instance Methods

	// Returns the buffer backing the given plane, or `nil` if none has been set.
	BufferForPlane(plane MTLTensorPlaneType) MTLBuffer
	// Returns the byte offset into the buffer for the given plane.
	OffsetForPlane(plane MTLTensorPlaneType) uint
	// Empties the container of all its elements.
	Reset()
	// Sets the buffer and byte offset to use as backing storage for the given plane.
	SetBufferOffsetForPlane(buffer MTLBuffer, offset uint, plane MTLTensorPlaneType)
}

// Init initializes the instance.
func (t MTLTensorBufferAttachments) Init() MTLTensorBufferAttachments {
	rv := objc.Send[MTLTensorBufferAttachments](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MTLTensorBufferAttachments) Autorelease() MTLTensorBufferAttachments {
	rv := objc.Send[MTLTensorBufferAttachments](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLTensorBufferAttachments creates a new MTLTensorBufferAttachments instance.
func NewMTLTensorBufferAttachments() MTLTensorBufferAttachments {
	class := getMTLTensorBufferAttachmentsClass()
	rv := objc.Send[MTLTensorBufferAttachments](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the buffer backing the given plane, or `nil` if none has been set.
//
// plane: The plane type to look up.
//
// # Return Value
//
// The buffer for the given plane, or `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorBufferAttachments/buffer(for:)
func (t MTLTensorBufferAttachments) BufferForPlane(plane MTLTensorPlaneType) MTLBuffer {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("bufferForPlane:"), plane)
	return MTLBufferObjectFromID(rv)
}

// Returns the byte offset into the buffer for the given plane.
//
// plane: The plane type to look up.
//
// # Return Value
//
// The byte offset for the given plane.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorBufferAttachments/offset(for:)
func (t MTLTensorBufferAttachments) OffsetForPlane(plane MTLTensorPlaneType) uint {
	rv := objc.Send[uint](t.ID, objc.Sel("offsetForPlane:"), plane)
	return rv
}

// Empties the container of all its elements.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorBufferAttachments/reset()
func (t MTLTensorBufferAttachments) Reset() {
	objc.Send[objc.ID](t.ID, objc.Sel("reset"))
}

// Sets the buffer and byte offset to use as backing storage for the given
// plane.
//
// buffer: The buffer to back the plane.
//
// offset: The byte offset into the buffer.
//
// plane: The plane type to associate the buffer with.
//
// # Discussion
//
// The offset needs to be aligned to 128 bytes if the plane uses
// [MTLTensorDataType.int2], [MTLTensorDataType.uint2],
// [MTLTensorDataTypeInt4], [MTLTensorDataTypeUInt4],
// [MTLTensorDataType.metalFloat4e2m1], [MTLTensorDataType.metalFloat8e4m3],
// [MTLTensorDataType.metalFloat8e5m2], or
// [MTLTensorDataType.metalFloat8ue8m0], otherwise it needs to be aligned to
// the size of the plane’s data type in bytes.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorBufferAttachments/setBuffer(_:offset:for:)
//
// [MTLTensorDataType.int2]: https://developer.apple.com/documentation/Metal/MTLTensorDataType/int2
// [MTLTensorDataType.metalFloat4e2m1]: https://developer.apple.com/documentation/Metal/MTLTensorDataType/metalFloat4e2m1
// [MTLTensorDataType.metalFloat8e4m3]: https://developer.apple.com/documentation/Metal/MTLTensorDataType/metalFloat8e4m3
// [MTLTensorDataType.metalFloat8e5m2]: https://developer.apple.com/documentation/Metal/MTLTensorDataType/metalFloat8e5m2
// [MTLTensorDataType.metalFloat8ue8m0]: https://developer.apple.com/documentation/Metal/MTLTensorDataType/metalFloat8ue8m0
// [MTLTensorDataType.uint2]: https://developer.apple.com/documentation/Metal/MTLTensorDataType/uint2
func (t MTLTensorBufferAttachments) SetBufferOffsetForPlane(buffer MTLBuffer, offset uint, plane MTLTensorPlaneType) {
	objc.Send[objc.ID](t.ID, objc.Sel("setBuffer:offset:forPlane:"), buffer, offset, plane)
}
