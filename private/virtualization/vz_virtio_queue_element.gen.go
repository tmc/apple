// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioQueueElement] class.
var (
	_VZVirtioQueueElementClass     VZVirtioQueueElementClass
	_VZVirtioQueueElementClassOnce sync.Once
)

func getVZVirtioQueueElementClass() VZVirtioQueueElementClass {
	_VZVirtioQueueElementClassOnce.Do(func() {
		_VZVirtioQueueElementClass = VZVirtioQueueElementClass{class: objc.GetClass("VZVirtioQueueElement")}
	})
	return _VZVirtioQueueElementClass
}

// GetVZVirtioQueueElementClass returns the class object for VZVirtioQueueElement.
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
	rv := objc.SendIfResponds[VZVirtioQueueElement](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZVirtioQueueElement.PeekIntoReadBuffersError]
//   - [VZVirtioQueueElement.PeekIntoReadBuffersWithExactLengthError]
//   - [VZVirtioQueueElement.ReadBuffers]
//   - [VZVirtioQueueElement.ReadBuffersAvailableByteCount]
//   - [VZVirtioQueueElement.ReadBuffersByteCount]
//   - [VZVirtioQueueElement.ReadBytesError]
//   - [VZVirtioQueueElement.ReadBytesIntoLengthError]
//   - [VZVirtioQueueElement.ReadBytesIntoBufferExactLengthError]
//   - [VZVirtioQueueElement.ReadBytesWithExactLengthError]
//   - [VZVirtioQueueElement.ReturnToQueue]
//   - [VZVirtioQueueElement.WriteError]
//   - [VZVirtioQueueElement.WriteBufferExactLengthError]
//   - [VZVirtioQueueElement.WriteBuffersAvailableByteCount]
//   - [VZVirtioQueueElement.WriteBuffersByteCount]
//   - [VZVirtioQueueElement.WriteDataError]
//   - [VZVirtioQueueElement.WriteDataLengthError]
//   - [VZVirtioQueueElement.WrittenByteCount]
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
//   - [IVZVirtioQueueElement.PeekIntoReadBuffersWithExactLengthError]
//   - [IVZVirtioQueueElement.ReadBuffers]
//   - [IVZVirtioQueueElement.ReadBuffersAvailableByteCount]
//   - [IVZVirtioQueueElement.ReadBuffersByteCount]
//   - [IVZVirtioQueueElement.ReadBytesError]
//   - [IVZVirtioQueueElement.ReadBytesIntoLengthError]
//   - [IVZVirtioQueueElement.ReadBytesIntoBufferExactLengthError]
//   - [IVZVirtioQueueElement.ReadBytesWithExactLengthError]
//   - [IVZVirtioQueueElement.ReturnToQueue]
//   - [IVZVirtioQueueElement.WriteError]
//   - [IVZVirtioQueueElement.WriteBufferExactLengthError]
//   - [IVZVirtioQueueElement.WriteBuffersAvailableByteCount]
//   - [IVZVirtioQueueElement.WriteBuffersByteCount]
//   - [IVZVirtioQueueElement.WriteDataError]
//   - [IVZVirtioQueueElement.WriteDataLengthError]
//   - [IVZVirtioQueueElement.WrittenByteCount]
type IVZVirtioQueueElement interface {
	objectivec.IObject

	// Topic: Methods

	PeekIntoReadBuffersError(buffers uint64) (objectivec.IObject, error)
	PeekIntoReadBuffersWithExactLengthError(length uint64) (objectivec.IObject, error)
	ReadBuffers() objectivec.IObject
	ReadBuffersAvailableByteCount() uint64
	ReadBuffersByteCount() uint64
	ReadBytesError(bytes uint64) (objectivec.IObject, error)
	ReadBytesIntoLengthError(into unsafe.Pointer, length uint64) (bool, error)
	ReadBytesIntoBufferExactLengthError(buffer unsafe.Pointer, length uint64) (bool, error)
	ReadBytesWithExactLengthError(length uint64) (objectivec.IObject, error)
	ReturnToQueue()
	WriteError(write objectivec.IObject) (bool, error)
	WriteBufferExactLengthError(buffer unsafe.Pointer, length uint64) (bool, error)
	WriteBuffersAvailableByteCount() uint64
	WriteBuffersByteCount() uint64
	WriteDataError(data objectivec.IObject) (bool, error)
	WriteDataLengthError(data unsafe.Pointer, length uint64) (bool, error)
	WrittenByteCount() uint64
}

// Init initializes the instance.
func (v VZVirtioQueueElement) Init() VZVirtioQueueElement {
	rv := objc.SendIfResponds[VZVirtioQueueElement](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioQueueElement) Autorelease() VZVirtioQueueElement {
	rv := objc.SendIfResponds[VZVirtioQueueElement](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioQueueElement creates a new VZVirtioQueueElement instance.
func NewVZVirtioQueueElement() VZVirtioQueueElement {
	class := getVZVirtioQueueElementClass()
	rv := objc.SendIfResponds[VZVirtioQueueElement](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZVirtioQueueElement) PeekIntoReadBuffersError(buffers uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("peekIntoReadBuffers:error:"), buffers, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (v VZVirtioQueueElement) PeekIntoReadBuffersWithExactLengthError(length uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("peekIntoReadBuffersWithExactLength:error:"), length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (v VZVirtioQueueElement) ReadBuffers() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("readBuffers"))
	return objectivec.Object{ID: rv}
}
func (v VZVirtioQueueElement) ReadBytesError(bytes uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("readBytes:error:"), bytes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
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
func (v VZVirtioQueueElement) ReadBytesIntoBufferExactLengthError(buffer unsafe.Pointer, length uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("readBytesIntoBuffer:exactLength:error:"), buffer, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("readBytesIntoBuffer:exactLength:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (v VZVirtioQueueElement) ReadBytesWithExactLengthError(length uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("readBytesWithExactLength:error:"), length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (v VZVirtioQueueElement) ReturnToQueue() {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("returnToQueue"))
}
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
func (v VZVirtioQueueElement) WriteBufferExactLengthError(buffer unsafe.Pointer, length uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("writeBuffer:exactLength:error:"), buffer, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeBuffer:exactLength:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (v VZVirtioQueueElement) WriteDataError(data objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("writeData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeData:error: returned NO with nil NSError")
	}
	return rv, nil

}
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

func (v VZVirtioQueueElement) ReadBuffersAvailableByteCount() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("readBuffersAvailableByteCount"))
	return rv
}
func (v VZVirtioQueueElement) ReadBuffersByteCount() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("readBuffersByteCount"))
	return rv
}
func (v VZVirtioQueueElement) WriteBuffersAvailableByteCount() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("writeBuffersAvailableByteCount"))
	return rv
}
func (v VZVirtioQueueElement) WriteBuffersByteCount() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("writeBuffersByteCount"))
	return rv
}
func (v VZVirtioQueueElement) WrittenByteCount() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("writtenByteCount"))
	return rv
}
