// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalDebugCommandEncoder] class.
var (
	_IOGPUMetalDebugCommandEncoderClass     IOGPUMetalDebugCommandEncoderClass
	_IOGPUMetalDebugCommandEncoderClassOnce sync.Once
)

func getIOGPUMetalDebugCommandEncoderClass() IOGPUMetalDebugCommandEncoderClass {
	_IOGPUMetalDebugCommandEncoderClassOnce.Do(func() {
		_IOGPUMetalDebugCommandEncoderClass = IOGPUMetalDebugCommandEncoderClass{class: objc.GetClass("IOGPUMetalDebugCommandEncoder")}
	})
	return _IOGPUMetalDebugCommandEncoderClass
}

// GetIOGPUMetalDebugCommandEncoderClass returns the class object for IOGPUMetalDebugCommandEncoder.
func GetIOGPUMetalDebugCommandEncoderClass() IOGPUMetalDebugCommandEncoderClass {
	return getIOGPUMetalDebugCommandEncoderClass()
}

type IOGPUMetalDebugCommandEncoderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalDebugCommandEncoderClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalDebugCommandEncoderClass) Alloc() IOGPUMetalDebugCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalDebugCommandEncoder](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalDebugCommandEncoder.IOLogBytesLength]
//   - [IOGPUMetalDebugCommandEncoder.AddAPIResource]
//   - [IOGPUMetalDebugCommandEncoder.AddDebugResourceListInfoFlag]
//   - [IOGPUMetalDebugCommandEncoder.DebugBytesLengthOutput_type]
//   - [IOGPUMetalDebugCommandEncoder.EndEncoding]
//   - [IOGPUMetalDebugCommandEncoder.KprintfBytesLength]
//   - [IOGPUMetalDebugCommandEncoder.ReserveKernelCommandBufferSpace]
//   - [IOGPUMetalDebugCommandEncoder.RestartDebugPass]
//   - [IOGPUMetalDebugCommandEncoder.InitWithCommandBuffer]
type IOGPUMetalDebugCommandEncoder struct {
	objectivec.Object
}

// IOGPUMetalDebugCommandEncoderFromID constructs a [IOGPUMetalDebugCommandEncoder] from an objc.ID.
func IOGPUMetalDebugCommandEncoderFromID(id objc.ID) IOGPUMetalDebugCommandEncoder {
	return IOGPUMetalDebugCommandEncoder{objectivec.Object{ID: id}}
}

// NOTE: IOGPUMetalDebugCommandEncoder embeds objectivec.Object because the parent type is
// unavailable, but IIOGPUMetalDebugCommandEncoder embeds IMTLDebugCommandEncoder, which that fallback
// cannot satisfy; skip compile-time assertion.

// An interface definition for the [IOGPUMetalDebugCommandEncoder] class.
//
// # Methods
//
//   - [IIOGPUMetalDebugCommandEncoder.IOLogBytesLength]
//   - [IIOGPUMetalDebugCommandEncoder.AddAPIResource]
//   - [IIOGPUMetalDebugCommandEncoder.AddDebugResourceListInfoFlag]
//   - [IIOGPUMetalDebugCommandEncoder.DebugBytesLengthOutput_type]
//   - [IIOGPUMetalDebugCommandEncoder.EndEncoding]
//   - [IIOGPUMetalDebugCommandEncoder.KprintfBytesLength]
//   - [IIOGPUMetalDebugCommandEncoder.ReserveKernelCommandBufferSpace]
//   - [IIOGPUMetalDebugCommandEncoder.RestartDebugPass]
//   - [IIOGPUMetalDebugCommandEncoder.InitWithCommandBuffer]
type IIOGPUMetalDebugCommandEncoder interface {
	IMTLDebugCommandEncoder

	// Topic: Methods

	IOLogBytesLength(bytes string, length uint64)
	AddAPIResource(aPIResource objectivec.IObject)
	AddDebugResourceListInfoFlag(info *IOGPUResourceInfo, flag uint32) uint32
	DebugBytesLengthOutput_type(bytes string, length uint64, output_type uint32)
	EndEncoding()
	KprintfBytesLength(bytes string, length uint64)
	ReserveKernelCommandBufferSpace(space uint64) unsafe.Pointer
	RestartDebugPass()
	InitWithCommandBuffer(buffer objectivec.IObject) IOGPUMetalDebugCommandEncoder
}

// Init initializes the instance.
func (i IOGPUMetalDebugCommandEncoder) Init() IOGPUMetalDebugCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalDebugCommandEncoder](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalDebugCommandEncoder) Autorelease() IOGPUMetalDebugCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalDebugCommandEncoder](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalDebugCommandEncoder creates a new IOGPUMetalDebugCommandEncoder instance.
func NewIOGPUMetalDebugCommandEncoder() IOGPUMetalDebugCommandEncoder {
	class := getIOGPUMetalDebugCommandEncoderClass()
	rv := objc.SendIfResponds[IOGPUMetalDebugCommandEncoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalDebugCommandEncoderWithCommandBuffer(buffer objectivec.IObject) IOGPUMetalDebugCommandEncoder {
	instance := getIOGPUMetalDebugCommandEncoderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCommandBuffer:"), buffer)
	return IOGPUMetalDebugCommandEncoderFromID(rv)
}

func (i IOGPUMetalDebugCommandEncoder) IOLogBytesLength(bytes string, length uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("IOLogBytes:length:"), unsafe.Pointer(unsafe.StringData(bytes+"\x00")), length)
}
func (i IOGPUMetalDebugCommandEncoder) AddAPIResource(aPIResource objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("addAPIResource:"), aPIResource)
}
func (i IOGPUMetalDebugCommandEncoder) AddDebugResourceListInfoFlag(info *IOGPUResourceInfo, flag uint32) uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("addDebugResourceListInfo:flag:"), unsafe.Pointer(info), flag)
	return rv
}
func (i IOGPUMetalDebugCommandEncoder) DebugBytesLengthOutput_type(bytes string, length uint64, output_type uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("debugBytes:length:output_type:"), unsafe.Pointer(unsafe.StringData(bytes+"\x00")), length, output_type)
}
func (i IOGPUMetalDebugCommandEncoder) EndEncoding() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("endEncoding"))
}
func (i IOGPUMetalDebugCommandEncoder) KprintfBytesLength(bytes string, length uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("kprintfBytes:length:"), unsafe.Pointer(unsafe.StringData(bytes+"\x00")), length)
}
func (i IOGPUMetalDebugCommandEncoder) ReserveKernelCommandBufferSpace(space uint64) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("reserveKernelCommandBufferSpace:"), space)
	return rv
}
func (i IOGPUMetalDebugCommandEncoder) RestartDebugPass() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("restartDebugPass"))
}
func (i IOGPUMetalDebugCommandEncoder) InitWithCommandBuffer(buffer objectivec.IObject) IOGPUMetalDebugCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalDebugCommandEncoder](i.ID, objc.Sel("initWithCommandBuffer:"), buffer)
	return rv
}
