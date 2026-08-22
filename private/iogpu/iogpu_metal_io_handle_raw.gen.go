// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalIOHandleRaw] class.
var (
	_IOGPUMetalIOHandleRawClass     IOGPUMetalIOHandleRawClass
	_IOGPUMetalIOHandleRawClassOnce sync.Once
)

func getIOGPUMetalIOHandleRawClass() IOGPUMetalIOHandleRawClass {
	_IOGPUMetalIOHandleRawClassOnce.Do(func() {
		_IOGPUMetalIOHandleRawClass = IOGPUMetalIOHandleRawClass{class: objc.GetClass("IOGPUMetalIOHandleRaw")}
	})
	return _IOGPUMetalIOHandleRawClass
}

// GetIOGPUMetalIOHandleRawClass returns the class object for IOGPUMetalIOHandleRaw.
func GetIOGPUMetalIOHandleRawClass() IOGPUMetalIOHandleRawClass {
	return getIOGPUMetalIOHandleRawClass()
}

type IOGPUMetalIOHandleRawClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalIOHandleRawClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalIOHandleRawClass) Alloc() IOGPUMetalIOHandleRaw {
	rv := objc.SendIfResponds[IOGPUMetalIOHandleRaw](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalIOHandleRaw.GlobalTraceObjectID]
//   - [IOGPUMetalIOHandleRaw.SetLabel]
//   - [IOGPUMetalIOHandleRaw.VnioID]
//   - [IOGPUMetalIOHandleRaw.InitWithDevicePathErrorUncached]
type IOGPUMetalIOHandleRaw struct {
	objectivec.Object
}

// IOGPUMetalIOHandleRawFromID constructs a [IOGPUMetalIOHandleRaw] from an objc.ID.
func IOGPUMetalIOHandleRawFromID(id objc.ID) IOGPUMetalIOHandleRaw {
	return IOGPUMetalIOHandleRaw{objectivec.Object{ID: id}}
}

// NOTE: IOGPUMetalIOHandleRaw embeds objectivec.Object because the parent type is
// unavailable, but IIOGPUMetalIOHandleRaw embeds IMTLIOHandleRaw, which that fallback
// cannot satisfy; skip compile-time assertion.

// An interface definition for the [IOGPUMetalIOHandleRaw] class.
//
// # Methods
//
//   - [IIOGPUMetalIOHandleRaw.GlobalTraceObjectID]
//   - [IIOGPUMetalIOHandleRaw.SetLabel]
//   - [IIOGPUMetalIOHandleRaw.VnioID]
//   - [IIOGPUMetalIOHandleRaw.InitWithDevicePathErrorUncached]
type IIOGPUMetalIOHandleRaw interface {
	IMTLIOHandleRaw

	// Topic: Methods

	GlobalTraceObjectID() uint64
	SetLabel(label objectivec.IObject)
	VnioID() uint32
	InitWithDevicePathErrorUncached(device objectivec.IObject, path string, error_ []objectivec.IObject, uncached bool) IOGPUMetalIOHandleRaw
}

// Init initializes the instance.
func (i IOGPUMetalIOHandleRaw) Init() IOGPUMetalIOHandleRaw {
	rv := objc.SendIfResponds[IOGPUMetalIOHandleRaw](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalIOHandleRaw) Autorelease() IOGPUMetalIOHandleRaw {
	rv := objc.SendIfResponds[IOGPUMetalIOHandleRaw](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalIOHandleRaw creates a new IOGPUMetalIOHandleRaw instance.
func NewIOGPUMetalIOHandleRaw() IOGPUMetalIOHandleRaw {
	class := getIOGPUMetalIOHandleRawClass()
	rv := objc.SendIfResponds[IOGPUMetalIOHandleRaw](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalIOHandleRawWithDevicePathErrorUncached(device objectivec.IObject, path string, error_ []objectivec.IObject, uncached bool) IOGPUMetalIOHandleRaw {
	instance := getIOGPUMetalIOHandleRawClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:path:error:uncached:"), device, unsafe.Pointer(unsafe.StringData(path+"\x00")), objectivec.IObjectSliceToNSArray(error_), uncached)
	return IOGPUMetalIOHandleRawFromID(rv)
}

func (i IOGPUMetalIOHandleRaw) GlobalTraceObjectID() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("globalTraceObjectID"))
	return rv
}
func (i IOGPUMetalIOHandleRaw) SetLabel(label objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setLabel:"), label)
}
func (i IOGPUMetalIOHandleRaw) InitWithDevicePathErrorUncached(device objectivec.IObject, path string, error_ []objectivec.IObject, uncached bool) IOGPUMetalIOHandleRaw {
	rv := objc.SendIfResponds[IOGPUMetalIOHandleRaw](i.ID, objc.Sel("initWithDevice:path:error:uncached:"), device, unsafe.Pointer(unsafe.StringData(path+"\x00")), objectivec.IObjectSliceToNSArray(error_), uncached)
	return rv
}

func (i IOGPUMetalIOHandleRaw) VnioID() uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("vnioID"))
	return rv
}
