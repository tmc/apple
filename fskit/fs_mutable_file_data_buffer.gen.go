// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSMutableFileDataBuffer] class.
var (
	_FSMutableFileDataBufferClass     FSMutableFileDataBufferClass
	_FSMutableFileDataBufferClassOnce sync.Once
)

func getFSMutableFileDataBufferClass() FSMutableFileDataBufferClass {
	_FSMutableFileDataBufferClassOnce.Do(func() {
		_FSMutableFileDataBufferClass = FSMutableFileDataBufferClass{class: objc.GetClass("FSMutableFileDataBuffer")}
	})
	return _FSMutableFileDataBufferClass
}

// GetFSMutableFileDataBufferClass returns the class object for FSMutableFileDataBuffer.
func GetFSMutableFileDataBufferClass() FSMutableFileDataBufferClass {
	return getFSMutableFileDataBufferClass()
}

type FSMutableFileDataBufferClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSMutableFileDataBufferClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSMutableFileDataBufferClass) Alloc() FSMutableFileDataBuffer {
	rv := objc.Send[FSMutableFileDataBuffer](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A wrapper object for a data buffer.
//
// # Overview
//
// This object provides a “zero-copy” buffer, for use when reading data
// from files. By not requiring additional buffer copying, this object reduces
// the extension’s memory footprint and improves performance. The
// [FSMutableFileDataBuffer] behaves similarly to a `uio` in the kernel.
//
// # Accessing buffer properties
//
//   - [FSMutableFileDataBuffer.Length]: The data length of the buffer.
//
// See: https://developer.apple.com/documentation/FSKit/FSMutableFileDataBuffer
type FSMutableFileDataBuffer struct {
	objectivec.Object
}

// FSMutableFileDataBufferFromID constructs a [FSMutableFileDataBuffer] from an objc.ID.
//
// A wrapper object for a data buffer.
func FSMutableFileDataBufferFromID(id objc.ID) FSMutableFileDataBuffer {
	return FSMutableFileDataBuffer{objectivec.Object{ID: id}}
}

// NOTE: FSMutableFileDataBuffer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSMutableFileDataBuffer] class.
//
// # Accessing buffer properties
//
//   - [IFSMutableFileDataBuffer.Length]: The data length of the buffer.
//
// See: https://developer.apple.com/documentation/FSKit/FSMutableFileDataBuffer
type IFSMutableFileDataBuffer interface {
	objectivec.IObject

	// Topic: Accessing buffer properties

	// The data length of the buffer.
	Length() uint

	// The byte data.
	MutableBytes() unsafe.Pointer
}

// Init initializes the instance.
func (m FSMutableFileDataBuffer) Init() FSMutableFileDataBuffer {
	rv := objc.Send[FSMutableFileDataBuffer](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m FSMutableFileDataBuffer) Autorelease() FSMutableFileDataBuffer {
	rv := objc.Send[FSMutableFileDataBuffer](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSMutableFileDataBuffer creates a new FSMutableFileDataBuffer instance.
func NewFSMutableFileDataBuffer() FSMutableFileDataBuffer {
	class := getFSMutableFileDataBufferClass()
	rv := objc.Send[FSMutableFileDataBuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The byte data.
//
// See: https://developer.apple.com/documentation/FSKit/FSMutableFileDataBuffer/mutableBytes
func (m FSMutableFileDataBuffer) MutableBytes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("mutableBytes"))
	return rv
}

// The data length of the buffer.
//
// See: https://developer.apple.com/documentation/FSKit/FSMutableFileDataBuffer/length
func (m FSMutableFileDataBuffer) Length() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("length"))
	return rv
}
