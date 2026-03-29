// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioQueueElement] class.
var (
	_VZVirtioQueueElementClass     VZVirtioQueueElementClass
	_VZVirtioQueueElementClassOnce sync.Once
)

func getVZVirtioQueueElementClass() VZVirtioQueueElementClass {
	_VZVirtioQueueElementClassOnce.Do(func() {
		_VZVirtioQueueElementClass = VZVirtioQueueElementClass{class: objc.GetClass("_VZVirtioQueueElement")}
	})
	return _VZVirtioQueueElementClass
}

// GetVZVirtioQueueElementClass returns the class object for _VZVirtioQueueElement.
func GetVZVirtioQueueElementClass() VZVirtioQueueElementClass {
	return getVZVirtioQueueElementClass()
}

type VZVirtioQueueElementClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioQueueElementClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioQueueElementClass) Alloc() VZVirtioQueueElement {
	rv := objc.Send[VZVirtioQueueElement](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZVirtioQueueElement.PeekIntoReadBuffersError]
//   - [VZVirtioQueueElement.ReadBuffers]
//   - [VZVirtioQueueElement.ReadBuffersAvailableByteCount]
//   - [VZVirtioQueueElement.ReadBuffersByteCount]
//   - [VZVirtioQueueElement.ReadBytesError]
//   - [VZVirtioQueueElement.ReadBytesIntoLengthError]
//   - [VZVirtioQueueElement.ReturnToQueue]
//   - [VZVirtioQueueElement.WriteError]
//   - [VZVirtioQueueElement.WriteBuffersAvailableByteCount]
//   - [VZVirtioQueueElement.WriteBuffersByteCount]
//   - [VZVirtioQueueElement.WriteDataLengthError]
//   - [VZVirtioQueueElement.WrittenByteCount]
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioQueueElement
type VZVirtioQueueElement struct {
	objectivec.Object
}

// VZVirtioQueueElementFromID constructs a [VZVirtioQueueElement] from an objc.ID.
func VZVirtioQueueElementFromID(id objc.ID) VZVirtioQueueElement {
	return VZVirtioQueueElement{objectivec.Object{ID: id}}
}
// Ensure VZVirtioQueueElement implements IVZVirtioQueueElement.
var _ IVZVirtioQueueElement = VZVirtioQueueElement{}

// An interface definition for the [VZVirtioQueueElement] class.
//
// # Methods
//
//   - [IVZVirtioQueueElement.PeekIntoReadBuffersError]
//   - [IVZVirtioQueueElement.ReadBuffers]
//   - [IVZVirtioQueueElement.ReadBuffersAvailableByteCount]
//   - [IVZVirtioQueueElement.ReadBuffersByteCount]
//   - [IVZVirtioQueueElement.ReadBytesError]
//   - [IVZVirtioQueueElement.ReadBytesIntoLengthError]
//   - [IVZVirtioQueueElement.ReturnToQueue]
//   - [IVZVirtioQueueElement.WriteError]
//   - [IVZVirtioQueueElement.WriteBuffersAvailableByteCount]
//   - [IVZVirtioQueueElement.WriteBuffersByteCount]
//   - [IVZVirtioQueueElement.WriteDataLengthError]
//   - [IVZVirtioQueueElement.WrittenByteCount]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioQueueElement
type IVZVirtioQueueElement interface {
	objectivec.IObject

	// Topic: Methods

	PeekIntoReadBuffersError(buffers uint64) (objectivec.IObject, error)
	ReadBuffers() objectivec.IObject
	ReadBuffersAvailableByteCount() uint64
	ReadBuffersByteCount() uint64
	ReadBytesError(bytes uint64) (objectivec.IObject, error)
	ReadBytesIntoLengthError(into unsafe.Pointer, length uint64) (bool, error)
	ReturnToQueue()
	WriteError(write objectivec.IObject) (bool, error)
	WriteBuffersAvailableByteCount() uint64
	WriteBuffersByteCount() uint64
	WriteDataLengthError(data unsafe.Pointer, length uint64) (bool, error)
	WrittenByteCount() uint64
}

// Init initializes the instance.
func (v VZVirtioQueueElement) Init() VZVirtioQueueElement {
	rv := objc.Send[VZVirtioQueueElement](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioQueueElement) Autorelease() VZVirtioQueueElement {
	rv := objc.Send[VZVirtioQueueElement](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioQueueElement creates a new VZVirtioQueueElement instance.
func NewVZVirtioQueueElement() VZVirtioQueueElement {
	class := getVZVirtioQueueElementClass()
	rv := objc.Send[VZVirtioQueueElement](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioQueueElement/peekIntoReadBuffers:error:
func (v VZVirtioQueueElement) PeekIntoReadBuffersError(buffers uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("peekIntoReadBuffers:error:"), buffers, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioQueueElement/readBuffers
func (v VZVirtioQueueElement) ReadBuffers() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("readBuffers"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioQueueElement/readBytes:error:
func (v VZVirtioQueueElement) ReadBytesError(bytes uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("readBytes:error:"), bytes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioQueueElement/readBytesInto:length:error:
func (v VZVirtioQueueElement) ReadBytesIntoLengthError(into unsafe.Pointer, length uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("readBytesInto:length:error:"), into, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("readBytesInto:length:error: returned NO with nil NSError")
	}
	return rv, nil

}
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioQueueElement/returnToQueue
func (v VZVirtioQueueElement) ReturnToQueue() {
	objc.Send[objc.ID](v.ID, objc.Sel("returnToQueue"))
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioQueueElement/write:error:
func (v VZVirtioQueueElement) WriteError(write objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("write:error:"), write, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("write:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioQueueElement/writeData:length:error:
func (v VZVirtioQueueElement) WriteDataLengthError(data unsafe.Pointer, length uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("writeData:length:error:"), data, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeData:length:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioQueueElement/readBuffersAvailableByteCount
func (v VZVirtioQueueElement) ReadBuffersAvailableByteCount() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("readBuffersAvailableByteCount"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioQueueElement/readBuffersByteCount
func (v VZVirtioQueueElement) ReadBuffersByteCount() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("readBuffersByteCount"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioQueueElement/writeBuffersAvailableByteCount
func (v VZVirtioQueueElement) WriteBuffersAvailableByteCount() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("writeBuffersAvailableByteCount"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioQueueElement/writeBuffersByteCount
func (v VZVirtioQueueElement) WriteBuffersByteCount() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("writeBuffersByteCount"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioQueueElement/writtenByteCount
func (v VZVirtioQueueElement) WrittenByteCount() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("writtenByteCount"))
	return rv
}

