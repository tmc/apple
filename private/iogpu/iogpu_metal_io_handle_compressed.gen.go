// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalIOHandleCompressed] class.
var (
	_IOGPUMetalIOHandleCompressedClass     IOGPUMetalIOHandleCompressedClass
	_IOGPUMetalIOHandleCompressedClassOnce sync.Once
)

func getIOGPUMetalIOHandleCompressedClass() IOGPUMetalIOHandleCompressedClass {
	_IOGPUMetalIOHandleCompressedClassOnce.Do(func() {
		_IOGPUMetalIOHandleCompressedClass = IOGPUMetalIOHandleCompressedClass{class: objc.GetClass("IOGPUMetalIOHandleCompressed")}
	})
	return _IOGPUMetalIOHandleCompressedClass
}

// GetIOGPUMetalIOHandleCompressedClass returns the class object for IOGPUMetalIOHandleCompressed.
func GetIOGPUMetalIOHandleCompressedClass() IOGPUMetalIOHandleCompressedClass {
	return getIOGPUMetalIOHandleCompressedClass()
}

type IOGPUMetalIOHandleCompressedClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalIOHandleCompressedClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalIOHandleCompressedClass) Alloc() IOGPUMetalIOHandleCompressed {
	rv := objc.SendIfResponds[IOGPUMetalIOHandleCompressed](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalIOHandleCompressed.GlobalTraceObjectID]
//   - [IOGPUMetalIOHandleCompressed.SetLabel]
//   - [IOGPUMetalIOHandleCompressed.VnioID]
//   - [IOGPUMetalIOHandleCompressed.InitWithDevicePathCompressionTypeErrorUncached]
type IOGPUMetalIOHandleCompressed struct {
	objectivec.Object
}

// IOGPUMetalIOHandleCompressedFromID constructs a [IOGPUMetalIOHandleCompressed] from an objc.ID.
func IOGPUMetalIOHandleCompressedFromID(id objc.ID) IOGPUMetalIOHandleCompressed {
	return IOGPUMetalIOHandleCompressed{objectivec.Object{ID: id}}
}

// NOTE: IOGPUMetalIOHandleCompressed embeds objectivec.Object because the parent type is
// unavailable, but IIOGPUMetalIOHandleCompressed embeds IMTLIOHandleCompressed, which that fallback
// cannot satisfy; skip compile-time assertion.

// An interface definition for the [IOGPUMetalIOHandleCompressed] class.
//
// # Methods
//
//   - [IIOGPUMetalIOHandleCompressed.GlobalTraceObjectID]
//   - [IIOGPUMetalIOHandleCompressed.SetLabel]
//   - [IIOGPUMetalIOHandleCompressed.VnioID]
//   - [IIOGPUMetalIOHandleCompressed.InitWithDevicePathCompressionTypeErrorUncached]
type IIOGPUMetalIOHandleCompressed interface {
	IMTLIOHandleCompressed

	// Topic: Methods

	GlobalTraceObjectID() uint64
	SetLabel(label objectivec.IObject)
	VnioID() uint32
	InitWithDevicePathCompressionTypeErrorUncached(device objectivec.IObject, path string, type_ int64, error_ []objectivec.IObject, uncached bool) IOGPUMetalIOHandleCompressed
}

// Init initializes the instance.
func (i IOGPUMetalIOHandleCompressed) Init() IOGPUMetalIOHandleCompressed {
	rv := objc.SendIfResponds[IOGPUMetalIOHandleCompressed](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalIOHandleCompressed) Autorelease() IOGPUMetalIOHandleCompressed {
	rv := objc.SendIfResponds[IOGPUMetalIOHandleCompressed](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalIOHandleCompressed creates a new IOGPUMetalIOHandleCompressed instance.
func NewIOGPUMetalIOHandleCompressed() IOGPUMetalIOHandleCompressed {
	class := getIOGPUMetalIOHandleCompressedClass()
	rv := objc.SendIfResponds[IOGPUMetalIOHandleCompressed](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalIOHandleCompressedWithDevicePathCompressionTypeErrorUncached(device objectivec.IObject, path string, type_ int64, error_ []objectivec.IObject, uncached bool) IOGPUMetalIOHandleCompressed {
	instance := getIOGPUMetalIOHandleCompressedClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:path:compressionType:error:uncached:"), device, unsafe.Pointer(unsafe.StringData(path+"\x00")), type_, objectivec.IObjectSliceToNSArray(error_), uncached)
	return IOGPUMetalIOHandleCompressedFromID(rv)
}

func (i IOGPUMetalIOHandleCompressed) GlobalTraceObjectID() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("globalTraceObjectID"))
	return rv
}
func (i IOGPUMetalIOHandleCompressed) SetLabel(label objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setLabel:"), label)
}
func (i IOGPUMetalIOHandleCompressed) InitWithDevicePathCompressionTypeErrorUncached(device objectivec.IObject, path string, type_ int64, error_ []objectivec.IObject, uncached bool) IOGPUMetalIOHandleCompressed {
	rv := objc.SendIfResponds[IOGPUMetalIOHandleCompressed](i.ID, objc.Sel("initWithDevice:path:compressionType:error:uncached:"), device, unsafe.Pointer(unsafe.StringData(path+"\x00")), type_, objectivec.IObjectSliceToNSArray(error_), uncached)
	return rv
}

func (i IOGPUMetalIOHandleCompressed) VnioID() uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("vnioID"))
	return rv
}
