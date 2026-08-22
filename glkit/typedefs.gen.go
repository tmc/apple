// Code generated from Apple documentation. DO NOT EDIT.

package glkit

import (
	"unsafe"

	"github.com/tmc/apple/foundation"
)

// See: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyPrvPtr
type GLKEffectPropertyPrvPtr = uintptr

// GLKMatrix2 is a `2x2` matrix.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrix2
// GLKMatrix2 is opaque storage with the size and alignment C gives GLKMatrix2:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type GLKMatrix2 [4]uint32

// GLKMatrix3 is a `3x3` matrix stored in column-major order.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrix3
// GLKMatrix3 is opaque storage with the size and alignment C gives GLKMatrix3:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type GLKMatrix3 [9]uint32

// GLKMatrix4 is a `4x4` matrix.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrix4
// GLKMatrix4 is an unresolved C aggregate typedef.
type GLKMatrix4 unsafe.Pointer

// GLKMatrixStackRef is an opaque type that represents a stack of 4 x 4 matrices, providing support for hierarchical transform modeling and similar tasks.
//
// See: https://developer.apple.com/documentation/GLKit/GLKMatrixStack
type GLKMatrixStackRef uintptr

// GLKQuaternion is a representation of a quaternion.
//
// See: https://developer.apple.com/documentation/GLKit/GLKQuaternion
// GLKQuaternion is an unresolved C aggregate typedef.
type GLKQuaternion unsafe.Pointer

// GLKTextureLoaderCallback is signature for the block executed after an asynchronous texture loading operation completes.
//
// See: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderCallback
type GLKTextureLoaderCallback = func(*uintptr, foundation.NSError)

// GLKVector2 is a representation of a 2-component vector.
//
// See: https://developer.apple.com/documentation/GLKit/GLKVector2
// GLKVector2 is opaque storage with the size and alignment C gives GLKVector2:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type GLKVector2 [1]uint64

// GLKVector3 is a representation of a 3-component vector.
//
// See: https://developer.apple.com/documentation/GLKit/GLKVector3
// GLKVector3 is opaque storage with the size and alignment C gives GLKVector3:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type GLKVector3 [3]uint32

// GLKVector4 is a representation of a 4-component vector.
//
// See: https://developer.apple.com/documentation/GLKit/GLKVector4
// GLKVector4 is an unresolved C aggregate typedef.
type GLKVector4 unsafe.Pointer

// See: https://developer.apple.com/documentation/GLKit/GLKVertexAttributeParameters
// GLKVertexAttributeParameters is opaque storage with the size and alignment C gives GLKVertexAttributeParameters:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type GLKVertexAttributeParameters [3]uint32
