// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalBuffer] class.
var (
	_IOGPUMetalBufferClass     IOGPUMetalBufferClass
	_IOGPUMetalBufferClassOnce sync.Once
)

func getIOGPUMetalBufferClass() IOGPUMetalBufferClass {
	_IOGPUMetalBufferClassOnce.Do(func() {
		_IOGPUMetalBufferClass = IOGPUMetalBufferClass{class: objc.GetClass("IOGPUMetalBuffer")}
	})
	return _IOGPUMetalBufferClass
}

// GetIOGPUMetalBufferClass returns the class object for IOGPUMetalBuffer.
func GetIOGPUMetalBufferClass() IOGPUMetalBufferClass {
	return getIOGPUMetalBufferClass()
}

type IOGPUMetalBufferClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalBufferClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalBufferClass) Alloc() IOGPUMetalBuffer {
	rv := objc.SendIfResponds[IOGPUMetalBuffer](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalBuffer._aneIOSurface]
//   - [IOGPUMetalBuffer.AddDebugMarkerRange]
//   - [IOGPUMetalBuffer.Contents]
//   - [IOGPUMetalBuffer.DetachBacking]
//   - [IOGPUMetalBuffer.DidModifyRange]
//   - [IOGPUMetalBuffer.FormattedDescription]
//   - [IOGPUMetalBuffer.Iosurface]
//   - [IOGPUMetalBuffer.Length]
//   - [IOGPUMetalBuffer.NewLinearTextureWithDescriptorOffsetBytesPerRowBytesPerImage]
//   - [IOGPUMetalBuffer.NewRemoteBufferViewForDevice]
//   - [IOGPUMetalBuffer.PlacementSparsePageSize]
//   - [IOGPUMetalBuffer.RemoteStorageBuffer]
//   - [IOGPUMetalBuffer.RemoveAllDebugMarkers]
//   - [IOGPUMetalBuffer.ReplaceBackingWithBytesNoCopyLengthDeallocator]
//   - [IOGPUMetalBuffer.ReplaceBackingWithRangesReadOnly]
//   - [IOGPUMetalBuffer.InitStandinWithDeviceBytesNoCopyLengthDeallocator]
//   - [IOGPUMetalBuffer.InitWithDeviceAddressRangesAddressRangeCountLengthOptionsGpuAddressArgsArgsSize]
//   - [IOGPUMetalBuffer.InitWithDeviceIosurfaceArgsArgsSize]
//   - [IOGPUMetalBuffer.InitWithDeviceIosurfaceGpuAddressArgsArgsSize]
//   - [IOGPUMetalBuffer.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocator]
//   - [IOGPUMetalBuffer.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagArgsArgsSizeDeallocator]
//   - [IOGPUMetalBuffer.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizeArgsArgsSizeDeallocator]
//   - [IOGPUMetalBuffer.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizePlacementSparseResidencyBytesArgsArgsSizeDeallocator]
//   - [IOGPUMetalBuffer.InitWithDevicePointerLengthOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocator]
//   - [IOGPUMetalBuffer.InitWithDevicePointerLengthOptionsSysMemSizeVidMemSizeArgsArgsSizeDeallocator]
//   - [IOGPUMetalBuffer.InitWithDeviceRemoteStorageBufferArgsArgsSize]
//   - [IOGPUMetalBuffer.InitWithHeapResourceOffsetLength]
//   - [IOGPUMetalBuffer.InitWithHeapResourceOffsetLengthGpuTag]
//   - [IOGPUMetalBuffer.InitWithPrimaryBufferHeapIndexBufferIndexBufferOffsetLengthArgsArgsSize]
//   - [IOGPUMetalBuffer.InitWithPrimaryBufferHeapIndexBufferIndexBufferOffsetLengthArgsArgsSizeGpuTag]
type IOGPUMetalBuffer struct {
	IOGPUMetalResource
}

// IOGPUMetalBufferFromID constructs a [IOGPUMetalBuffer] from an objc.ID.
func IOGPUMetalBufferFromID(id objc.ID) IOGPUMetalBuffer {
	return IOGPUMetalBuffer{IOGPUMetalResource: IOGPUMetalResourceFromID(id)}
}

// Ensure IOGPUMetalBuffer implements IIOGPUMetalBuffer.
var _ IIOGPUMetalBuffer = IOGPUMetalBuffer{}

// An interface definition for the [IOGPUMetalBuffer] class.
//
// # Methods
//
//   - [IIOGPUMetalBuffer._aneIOSurface]
//   - [IIOGPUMetalBuffer.AddDebugMarkerRange]
//   - [IIOGPUMetalBuffer.Contents]
//   - [IIOGPUMetalBuffer.DetachBacking]
//   - [IIOGPUMetalBuffer.DidModifyRange]
//   - [IIOGPUMetalBuffer.FormattedDescription]
//   - [IIOGPUMetalBuffer.Iosurface]
//   - [IIOGPUMetalBuffer.Length]
//   - [IIOGPUMetalBuffer.NewLinearTextureWithDescriptorOffsetBytesPerRowBytesPerImage]
//   - [IIOGPUMetalBuffer.NewRemoteBufferViewForDevice]
//   - [IIOGPUMetalBuffer.PlacementSparsePageSize]
//   - [IIOGPUMetalBuffer.RemoteStorageBuffer]
//   - [IIOGPUMetalBuffer.RemoveAllDebugMarkers]
//   - [IIOGPUMetalBuffer.ReplaceBackingWithBytesNoCopyLengthDeallocator]
//   - [IIOGPUMetalBuffer.ReplaceBackingWithRangesReadOnly]
//   - [IIOGPUMetalBuffer.InitStandinWithDeviceBytesNoCopyLengthDeallocator]
//   - [IIOGPUMetalBuffer.InitWithDeviceAddressRangesAddressRangeCountLengthOptionsGpuAddressArgsArgsSize]
//   - [IIOGPUMetalBuffer.InitWithDeviceIosurfaceArgsArgsSize]
//   - [IIOGPUMetalBuffer.InitWithDeviceIosurfaceGpuAddressArgsArgsSize]
//   - [IIOGPUMetalBuffer.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocator]
//   - [IIOGPUMetalBuffer.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagArgsArgsSizeDeallocator]
//   - [IIOGPUMetalBuffer.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizeArgsArgsSizeDeallocator]
//   - [IIOGPUMetalBuffer.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizePlacementSparseResidencyBytesArgsArgsSizeDeallocator]
//   - [IIOGPUMetalBuffer.InitWithDevicePointerLengthOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocator]
//   - [IIOGPUMetalBuffer.InitWithDevicePointerLengthOptionsSysMemSizeVidMemSizeArgsArgsSizeDeallocator]
//   - [IIOGPUMetalBuffer.InitWithDeviceRemoteStorageBufferArgsArgsSize]
//   - [IIOGPUMetalBuffer.InitWithHeapResourceOffsetLength]
//   - [IIOGPUMetalBuffer.InitWithHeapResourceOffsetLengthGpuTag]
//   - [IIOGPUMetalBuffer.InitWithPrimaryBufferHeapIndexBufferIndexBufferOffsetLengthArgsArgsSize]
//   - [IIOGPUMetalBuffer.InitWithPrimaryBufferHeapIndexBufferIndexBufferOffsetLengthArgsArgsSizeGpuTag]
type IIOGPUMetalBuffer interface {
	IIOGPUMetalResource

	// Topic: Methods

	_aneIOSurface() iosurface.IOSurfaceRef
	AddDebugMarkerRange(marker objectivec.IObject, range_ foundation.NSRange)
	Contents() unsafe.Pointer
	DetachBacking() bool
	DidModifyRange(range_ foundation.NSRange)
	FormattedDescription(description uint64) objectivec.IObject
	Iosurface() iosurface.IOSurfaceRef
	Length() uint64
	NewLinearTextureWithDescriptorOffsetBytesPerRowBytesPerImage(descriptor objectivec.IObject, offset uint64, row uint64, image uint64) objectivec.IObject
	NewRemoteBufferViewForDevice(device objectivec.IObject) objectivec.IObject
	PlacementSparsePageSize() int64
	RemoteStorageBuffer() unsafe.Pointer
	RemoveAllDebugMarkers()
	ReplaceBackingWithBytesNoCopyLengthDeallocator(copy_ unsafe.Pointer, length uint64, deallocator VoidHandler) bool
	ReplaceBackingWithRangesReadOnly(ranges objectivec.IObject, only bool) bool
	InitStandinWithDeviceBytesNoCopyLengthDeallocator(device objectivec.IObject, copy_ unsafe.Pointer, length uint64, deallocator VoidHandler) IOGPUMetalBuffer
	InitWithDeviceAddressRangesAddressRangeCountLengthOptionsGpuAddressArgsArgsSize(device objectivec.IObject, ranges *IOGPUAddressRange, count uint64, length uint64, options uint64, address uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalBuffer
	InitWithDeviceIosurfaceArgsArgsSize(device objectivec.IObject, iosurface iosurface.IOSurfaceRef, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalBuffer
	InitWithDeviceIosurfaceGpuAddressArgsArgsSize(device objectivec.IObject, iosurface iosurface.IOSurfaceRef, address uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalBuffer
	InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocator(device objectivec.IObject, pointer unsafe.Pointer, length uint64, alignment uint32, options uint64, size uint64, address uint64, args *IOGPUNewResourceArgs, size2 uint32, deallocator VoidHandler) IOGPUMetalBuffer
	InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagArgsArgsSizeDeallocator(device objectivec.IObject, pointer unsafe.Pointer, length uint64, alignment uint32, options uint64, size uint64, address uint64, tag uint64, args *IOGPUNewResourceArgs, size2 uint32, deallocator VoidHandler) IOGPUMetalBuffer
	InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizeArgsArgsSizeDeallocator(device objectivec.IObject, pointer unsafe.Pointer, length uint64, alignment uint32, options uint64, size uint64, address uint64, tag uint64, size2 int64, args *IOGPUNewResourceArgs, size3 uint32, deallocator VoidHandler) IOGPUMetalBuffer
	InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizePlacementSparseResidencyBytesArgsArgsSizeDeallocator(device objectivec.IObject, pointer unsafe.Pointer, length uint64, alignment uint32, options uint64, size uint64, address uint64, tag uint64, size2 int64, bytes uint64, args *IOGPUNewResourceArgs, size3 uint32, deallocator VoidHandler) IOGPUMetalBuffer
	InitWithDevicePointerLengthOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocator(device objectivec.IObject, pointer unsafe.Pointer, length uint64, options uint64, size uint64, address uint64, args *IOGPUNewResourceArgs, size2 uint32, deallocator VoidHandler) IOGPUMetalBuffer
	InitWithDevicePointerLengthOptionsSysMemSizeVidMemSizeArgsArgsSizeDeallocator(device objectivec.IObject, pointer unsafe.Pointer, length uint64, options uint64, size uint64, size2 uint64, args *IOGPUNewResourceArgs, size3 uint32, deallocator VoidHandler) IOGPUMetalBuffer
	InitWithDeviceRemoteStorageBufferArgsArgsSize(device objectivec.IObject, buffer objectivec.IObject, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalBuffer
	InitWithHeapResourceOffsetLength(heap objectivec.IObject, resource objectivec.IObject, offset uint64, length uint64) IOGPUMetalBuffer
	InitWithHeapResourceOffsetLengthGpuTag(heap objectivec.IObject, resource objectivec.IObject, offset uint64, length uint64, tag uint64) IOGPUMetalBuffer
	InitWithPrimaryBufferHeapIndexBufferIndexBufferOffsetLengthArgsArgsSize(buffer objectivec.IObject, index int16, index2 int16, offset uint64, length uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalBuffer
	InitWithPrimaryBufferHeapIndexBufferIndexBufferOffsetLengthArgsArgsSizeGpuTag(buffer objectivec.IObject, index int16, index2 int16, offset uint64, length uint64, args *IOGPUNewResourceArgs, size uint32, tag uint64) IOGPUMetalBuffer
}

// Init initializes the instance.
func (i IOGPUMetalBuffer) Init() IOGPUMetalBuffer {
	rv := objc.SendIfResponds[IOGPUMetalBuffer](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalBuffer) Autorelease() IOGPUMetalBuffer {
	rv := objc.SendIfResponds[IOGPUMetalBuffer](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalBuffer creates a new IOGPUMetalBuffer instance.
func NewIOGPUMetalBuffer() IOGPUMetalBuffer {
	class := getIOGPUMetalBufferClass()
	rv := objc.SendIfResponds[IOGPUMetalBuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalBufferMemorylessDescriptor(memoryless objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalBuffer {
	instance := getIOGPUMetalBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initMemoryless:descriptor:"), memoryless, descriptor)
	return IOGPUMetalBufferFromID(rv)
}

func NewGPUMetalBufferStandinWithDevice(device objectivec.IObject) IOGPUMetalBuffer {
	instance := getIOGPUMetalBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initStandinWithDevice:"), device)
	return IOGPUMetalBufferFromID(rv)
}

func NewGPUMetalBufferWithDeviceAddressRangesAddressRangeCountLengthOptionsGpuAddressArgsArgsSize(device objectivec.IObject, ranges *IOGPUAddressRange, count uint64, length uint64, options uint64, address uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalBuffer {
	instance := getIOGPUMetalBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:addressRanges:addressRangeCount:length:options:gpuAddress:args:argsSize:"), device, unsafe.Pointer(ranges), count, length, options, address, unsafe.Pointer(args), size)
	return IOGPUMetalBufferFromID(rv)
}

func NewGPUMetalBufferWithDeviceIosurfaceArgsArgsSize(device objectivec.IObject, iosurface iosurface.IOSurfaceRef, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalBuffer {
	instance := getIOGPUMetalBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:iosurface:args:argsSize:"), device, iosurface, unsafe.Pointer(args), size)
	return IOGPUMetalBufferFromID(rv)
}

func NewGPUMetalBufferWithDeviceIosurfaceGpuAddressArgsArgsSize(device objectivec.IObject, iosurface iosurface.IOSurfaceRef, address uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalBuffer {
	instance := getIOGPUMetalBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:iosurface:gpuAddress:args:argsSize:"), device, iosurface, address, unsafe.Pointer(args), size)
	return IOGPUMetalBufferFromID(rv)
}

func NewGPUMetalBufferWithDeviceOptionsArgsArgsSize(device objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalBuffer {
	instance := getIOGPUMetalBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:options:args:argsSize:"), device, options, unsafe.Pointer(args), size)
	return IOGPUMetalBufferFromID(rv)
}

func NewGPUMetalBufferWithDeviceRemoteStorageBufferArgsArgsSize(device objectivec.IObject, buffer objectivec.IObject, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalBuffer {
	instance := getIOGPUMetalBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:remoteStorageBuffer:args:argsSize:"), device, buffer, unsafe.Pointer(args), size)
	return IOGPUMetalBufferFromID(rv)
}

func NewGPUMetalBufferWithDeviceRemoteStorageResourceOptionsArgsArgsSize(device objectivec.IObject, resource objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalBuffer {
	instance := getIOGPUMetalBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:remoteStorageResource:options:args:argsSize:"), device, resource, options, unsafe.Pointer(args), size)
	return IOGPUMetalBufferFromID(rv)
}

func NewGPUMetalBufferWithHeapResourceOffsetLength(heap objectivec.IObject, resource objectivec.IObject, offset uint64, length uint64) IOGPUMetalBuffer {
	instance := getIOGPUMetalBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithHeap:resource:offset:length:"), heap, resource, offset, length)
	return IOGPUMetalBufferFromID(rv)
}

func NewGPUMetalBufferWithHeapResourceOffsetLengthGpuTag(heap objectivec.IObject, resource objectivec.IObject, offset uint64, length uint64, tag uint64) IOGPUMetalBuffer {
	instance := getIOGPUMetalBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithHeap:resource:offset:length:gpuTag:"), heap, resource, offset, length, tag)
	return IOGPUMetalBufferFromID(rv)
}

func NewGPUMetalBufferWithPrimaryBufferHeapIndexBufferIndexBufferOffsetLengthArgsArgsSize(buffer objectivec.IObject, index int16, index2 int16, offset uint64, length uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalBuffer {
	instance := getIOGPUMetalBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithPrimaryBuffer:heapIndex:bufferIndex:bufferOffset:length:args:argsSize:"), buffer, index, index2, offset, length, unsafe.Pointer(args), size)
	return IOGPUMetalBufferFromID(rv)
}

func NewGPUMetalBufferWithPrimaryBufferHeapIndexBufferIndexBufferOffsetLengthArgsArgsSizeGpuTag(buffer objectivec.IObject, index int16, index2 int16, offset uint64, length uint64, args *IOGPUNewResourceArgs, size uint32, tag uint64) IOGPUMetalBuffer {
	instance := getIOGPUMetalBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithPrimaryBuffer:heapIndex:bufferIndex:bufferOffset:length:args:argsSize:gpuTag:"), buffer, index, index2, offset, length, unsafe.Pointer(args), size, tag)
	return IOGPUMetalBufferFromID(rv)
}

func NewGPUMetalBufferWithResource(resource objectivec.IObject) IOGPUMetalBuffer {
	instance := getIOGPUMetalBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return IOGPUMetalBufferFromID(rv)
}

func (i IOGPUMetalBuffer) AddDebugMarkerRange(marker objectivec.IObject, range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("addDebugMarker:range:"), marker, range_)
}
func (i IOGPUMetalBuffer) DetachBacking() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("detachBacking"))
	return rv
}
func (i IOGPUMetalBuffer) DidModifyRange(range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("didModifyRange:"), range_)
}
func (i IOGPUMetalBuffer) FormattedDescription(description uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("formattedDescription:"), description)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalBuffer) NewLinearTextureWithDescriptorOffsetBytesPerRowBytesPerImage(descriptor objectivec.IObject, offset uint64, row uint64, image uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newLinearTextureWithDescriptor:offset:bytesPerRow:bytesPerImage:"), descriptor, offset, row, image)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalBuffer) NewRemoteBufferViewForDevice(device objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newRemoteBufferViewForDevice:"), device)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalBuffer) RemoveAllDebugMarkers() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("removeAllDebugMarkers"))
}

var _iogpumetalbuffer_replacebackingwithbytesnocopy_length_deallocator_p2_key byte

func (i IOGPUMetalBuffer) ReplaceBackingWithBytesNoCopyLengthDeallocator(copy_ unsafe.Pointer, length uint64, deallocator VoidHandler) bool {
	_block2, _ := NewVoidBlock(deallocator)
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("replaceBackingWithBytesNoCopy:length:deallocator:"), copy_, length, _block2)
	return rv
}
func (i IOGPUMetalBuffer) ReplaceBackingWithRangesReadOnly(ranges objectivec.IObject, only bool) bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("replaceBackingWithRanges:readOnly:"), ranges, only)
	return rv
}

var _iogpumetalbuffer_initstandinwithdevice_bytesnocopy_length_deallocator_p3_key byte

func (i IOGPUMetalBuffer) InitStandinWithDeviceBytesNoCopyLengthDeallocator(device objectivec.IObject, copy_ unsafe.Pointer, length uint64, deallocator VoidHandler) IOGPUMetalBuffer {
	_block3, _ := NewVoidBlock(deallocator)
	rv := objc.SendIfResponds[IOGPUMetalBuffer](i.ID, objc.Sel("initStandinWithDevice:bytesNoCopy:length:deallocator:"), device, copy_, length, _block3)
	return rv
}
func (i IOGPUMetalBuffer) InitWithDeviceAddressRangesAddressRangeCountLengthOptionsGpuAddressArgsArgsSize(device objectivec.IObject, ranges *IOGPUAddressRange, count uint64, length uint64, options uint64, address uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalBuffer {
	rv := objc.SendIfResponds[IOGPUMetalBuffer](i.ID, objc.Sel("initWithDevice:addressRanges:addressRangeCount:length:options:gpuAddress:args:argsSize:"), device, unsafe.Pointer(ranges), count, length, options, address, unsafe.Pointer(args), size)
	return rv
}
func (i IOGPUMetalBuffer) InitWithDeviceIosurfaceArgsArgsSize(device objectivec.IObject, iosurface iosurface.IOSurfaceRef, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalBuffer {
	rv := objc.SendIfResponds[IOGPUMetalBuffer](i.ID, objc.Sel("initWithDevice:iosurface:args:argsSize:"), device, iosurface, unsafe.Pointer(args), size)
	return rv
}
func (i IOGPUMetalBuffer) InitWithDeviceIosurfaceGpuAddressArgsArgsSize(device objectivec.IObject, iosurface iosurface.IOSurfaceRef, address uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalBuffer {
	rv := objc.SendIfResponds[IOGPUMetalBuffer](i.ID, objc.Sel("initWithDevice:iosurface:gpuAddress:args:argsSize:"), device, iosurface, address, unsafe.Pointer(args), size)
	return rv
}

var _iogpumetalbuffer_initwithdevice_pointer_length_alignment_options_sysmemsize_gpuaddress_args_argssize_deallocator_p9_key byte

func (i IOGPUMetalBuffer) InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocator(device objectivec.IObject, pointer unsafe.Pointer, length uint64, alignment uint32, options uint64, size uint64, address uint64, args *IOGPUNewResourceArgs, size2 uint32, deallocator VoidHandler) IOGPUMetalBuffer {
	_block9, _ := NewVoidBlock(deallocator)
	rv := objc.SendIfResponds[IOGPUMetalBuffer](i.ID, objc.Sel("initWithDevice:pointer:length:alignment:options:sysMemSize:gpuAddress:args:argsSize:deallocator:"), device, pointer, length, alignment, options, size, address, args, size2, _block9)
	return rv
}

var _iogpumetalbuffer_initwithdevice_pointer_length_alignment_options_sysmemsize_gpuaddress_gputag_args_argssize_deallocator_p10_key byte

func (i IOGPUMetalBuffer) InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagArgsArgsSizeDeallocator(device objectivec.IObject, pointer unsafe.Pointer, length uint64, alignment uint32, options uint64, size uint64, address uint64, tag uint64, args *IOGPUNewResourceArgs, size2 uint32, deallocator VoidHandler) IOGPUMetalBuffer {
	_block10, _ := NewVoidBlock(deallocator)
	rv := objc.SendIfResponds[IOGPUMetalBuffer](i.ID, objc.Sel("initWithDevice:pointer:length:alignment:options:sysMemSize:gpuAddress:gpuTag:args:argsSize:deallocator:"), device, pointer, length, alignment, options, size, address, tag, args, size2, _block10)
	return rv
}

var _iogpumetalbuffer_initwithdevice_pointer_length_alignment_options_sysmemsize_gpuaddress_gputag_placementsparsepagesize_args_argssize_deallocator_p11_key byte

func (i IOGPUMetalBuffer) InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizeArgsArgsSizeDeallocator(device objectivec.IObject, pointer unsafe.Pointer, length uint64, alignment uint32, options uint64, size uint64, address uint64, tag uint64, size2 int64, args *IOGPUNewResourceArgs, size3 uint32, deallocator VoidHandler) IOGPUMetalBuffer {
	_block11, _ := NewVoidBlock(deallocator)
	rv := objc.SendIfResponds[IOGPUMetalBuffer](i.ID, objc.Sel("initWithDevice:pointer:length:alignment:options:sysMemSize:gpuAddress:gpuTag:placementSparsePageSize:args:argsSize:deallocator:"), device, pointer, length, alignment, options, size, address, tag, size2, args, size3, _block11)
	return rv
}

var _iogpumetalbuffer_initwithdevice_pointer_length_alignment_options_sysmemsize_gpuaddress_gputag_placementsparsepagesize_placementsparseresidencybytes_args_argssize_deallocator_p12_key byte

func (i IOGPUMetalBuffer) InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizePlacementSparseResidencyBytesArgsArgsSizeDeallocator(device objectivec.IObject, pointer unsafe.Pointer, length uint64, alignment uint32, options uint64, size uint64, address uint64, tag uint64, size2 int64, bytes uint64, args *IOGPUNewResourceArgs, size3 uint32, deallocator VoidHandler) IOGPUMetalBuffer {
	_block12, _ := NewVoidBlock(deallocator)
	rv := objc.SendIfResponds[IOGPUMetalBuffer](i.ID, objc.Sel("initWithDevice:pointer:length:alignment:options:sysMemSize:gpuAddress:gpuTag:placementSparsePageSize:placementSparseResidencyBytes:args:argsSize:deallocator:"), device, pointer, length, alignment, options, size, address, tag, size2, bytes, args, size3, _block12)
	return rv
}

var _iogpumetalbuffer_initwithdevice_pointer_length_options_sysmemsize_gpuaddress_args_argssize_deallocator_p8_key byte

func (i IOGPUMetalBuffer) InitWithDevicePointerLengthOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocator(device objectivec.IObject, pointer unsafe.Pointer, length uint64, options uint64, size uint64, address uint64, args *IOGPUNewResourceArgs, size2 uint32, deallocator VoidHandler) IOGPUMetalBuffer {
	_block8, _ := NewVoidBlock(deallocator)
	rv := objc.SendIfResponds[IOGPUMetalBuffer](i.ID, objc.Sel("initWithDevice:pointer:length:options:sysMemSize:gpuAddress:args:argsSize:deallocator:"), device, pointer, length, options, size, address, args, size2, _block8)
	return rv
}

var _iogpumetalbuffer_initwithdevice_pointer_length_options_sysmemsize_vidmemsize_args_argssize_deallocator_p8_key byte

func (i IOGPUMetalBuffer) InitWithDevicePointerLengthOptionsSysMemSizeVidMemSizeArgsArgsSizeDeallocator(device objectivec.IObject, pointer unsafe.Pointer, length uint64, options uint64, size uint64, size2 uint64, args *IOGPUNewResourceArgs, size3 uint32, deallocator VoidHandler) IOGPUMetalBuffer {
	_block8, _ := NewVoidBlock(deallocator)
	rv := objc.SendIfResponds[IOGPUMetalBuffer](i.ID, objc.Sel("initWithDevice:pointer:length:options:sysMemSize:vidMemSize:args:argsSize:deallocator:"), device, pointer, length, options, size, size2, args, size3, _block8)
	return rv
}
func (i IOGPUMetalBuffer) InitWithDeviceRemoteStorageBufferArgsArgsSize(device objectivec.IObject, buffer objectivec.IObject, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalBuffer {
	rv := objc.SendIfResponds[IOGPUMetalBuffer](i.ID, objc.Sel("initWithDevice:remoteStorageBuffer:args:argsSize:"), device, buffer, unsafe.Pointer(args), size)
	return rv
}
func (i IOGPUMetalBuffer) InitWithHeapResourceOffsetLength(heap objectivec.IObject, resource objectivec.IObject, offset uint64, length uint64) IOGPUMetalBuffer {
	rv := objc.SendIfResponds[IOGPUMetalBuffer](i.ID, objc.Sel("initWithHeap:resource:offset:length:"), heap, resource, offset, length)
	return rv
}
func (i IOGPUMetalBuffer) InitWithHeapResourceOffsetLengthGpuTag(heap objectivec.IObject, resource objectivec.IObject, offset uint64, length uint64, tag uint64) IOGPUMetalBuffer {
	rv := objc.SendIfResponds[IOGPUMetalBuffer](i.ID, objc.Sel("initWithHeap:resource:offset:length:gpuTag:"), heap, resource, offset, length, tag)
	return rv
}
func (i IOGPUMetalBuffer) InitWithPrimaryBufferHeapIndexBufferIndexBufferOffsetLengthArgsArgsSize(buffer objectivec.IObject, index int16, index2 int16, offset uint64, length uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalBuffer {
	rv := objc.SendIfResponds[IOGPUMetalBuffer](i.ID, objc.Sel("initWithPrimaryBuffer:heapIndex:bufferIndex:bufferOffset:length:args:argsSize:"), buffer, index, index2, offset, length, unsafe.Pointer(args), size)
	return rv
}
func (i IOGPUMetalBuffer) InitWithPrimaryBufferHeapIndexBufferIndexBufferOffsetLengthArgsArgsSizeGpuTag(buffer objectivec.IObject, index int16, index2 int16, offset uint64, length uint64, args *IOGPUNewResourceArgs, size uint32, tag uint64) IOGPUMetalBuffer {
	rv := objc.SendIfResponds[IOGPUMetalBuffer](i.ID, objc.Sel("initWithPrimaryBuffer:heapIndex:bufferIndex:bufferOffset:length:args:argsSize:gpuTag:"), buffer, index, index2, offset, length, unsafe.Pointer(args), size, tag)
	return rv
}

func (i IOGPUMetalBuffer) _aneIOSurface() iosurface.IOSurfaceRef {
	rv := objc.SendIfResponds[iosurface.IOSurfaceRef](i.ID, objc.Sel("_aneIOSurface"))
	return iosurface.IOSurfaceRef(rv)
}

// CanAneIOSurface reports whether the receiver responds to the private selector _aneIOSurface.
func (i IOGPUMetalBuffer) CanAneIOSurface() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_aneIOSurface"))
}

// AneIOSurface is an exported wrapper for the private property _aneIOSurface.
func (i IOGPUMetalBuffer) AneIOSurface() (iosurface.IOSurfaceRef, error) {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_aneIOSurface")) {
		return *new(iosurface.IOSurfaceRef), &objc.UnrecognizedSelectorError{Selector: "_aneIOSurface"}
	}
	return i._aneIOSurface(), nil
}
func (i IOGPUMetalBuffer) Contents() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("contents"))
	return rv
}
func (i IOGPUMetalBuffer) Iosurface() iosurface.IOSurfaceRef {
	rv := objc.SendIfResponds[iosurface.IOSurfaceRef](i.ID, objc.Sel("iosurface"))
	return iosurface.IOSurfaceRef(rv)
}
func (i IOGPUMetalBuffer) Length() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("length"))
	return rv
}
func (i IOGPUMetalBuffer) PlacementSparsePageSize() int64 {
	rv := objc.SendIfResponds[int64](i.ID, objc.Sel("placementSparsePageSize"))
	return rv
}
func (i IOGPUMetalBuffer) RemoteStorageBuffer() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("remoteStorageBuffer"))
	return rv
}

// ReplaceBackingWithBytesNoCopyLengthDeallocatorSync is a synchronous wrapper around [IOGPUMetalBuffer.ReplaceBackingWithBytesNoCopyLengthDeallocator].
// It blocks until the completion handler fires or the context is cancelled.
func (i IOGPUMetalBuffer) ReplaceBackingWithBytesNoCopyLengthDeallocatorSync(ctx context.Context, copy_ unsafe.Pointer, length uint64) error {
	done := make(chan struct{}, 1)
	i.ReplaceBackingWithBytesNoCopyLengthDeallocator(copy_, length, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// InitStandinWithDeviceBytesNoCopyLengthDeallocatorSync is a synchronous wrapper around [IOGPUMetalBuffer.InitStandinWithDeviceBytesNoCopyLengthDeallocator].
// It blocks until the completion handler fires or the context is cancelled.
func (i IOGPUMetalBuffer) InitStandinWithDeviceBytesNoCopyLengthDeallocatorSync(ctx context.Context, device objectivec.IObject, copy_ unsafe.Pointer, length uint64) error {
	done := make(chan struct{}, 1)
	i.InitStandinWithDeviceBytesNoCopyLengthDeallocator(device, copy_, length, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocatorSync is a synchronous wrapper around [IOGPUMetalBuffer.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocator].
// It blocks until the completion handler fires or the context is cancelled.
func (i IOGPUMetalBuffer) InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocatorSync(ctx context.Context, device objectivec.IObject, pointer unsafe.Pointer, length uint64, alignment uint32, options uint64, size uint64, address uint64, args *IOGPUNewResourceArgs, size2 uint32) error {
	done := make(chan struct{}, 1)
	i.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocator(device, pointer, length, alignment, options, size, address, args, size2, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagArgsArgsSizeDeallocatorSync is a synchronous wrapper around [IOGPUMetalBuffer.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagArgsArgsSizeDeallocator].
// It blocks until the completion handler fires or the context is cancelled.
func (i IOGPUMetalBuffer) InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagArgsArgsSizeDeallocatorSync(ctx context.Context, device objectivec.IObject, pointer unsafe.Pointer, length uint64, alignment uint32, options uint64, size uint64, address uint64, tag uint64, args *IOGPUNewResourceArgs, size2 uint32) error {
	done := make(chan struct{}, 1)
	i.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagArgsArgsSizeDeallocator(device, pointer, length, alignment, options, size, address, tag, args, size2, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizeArgsArgsSizeDeallocatorSync is a synchronous wrapper around [IOGPUMetalBuffer.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizeArgsArgsSizeDeallocator].
// It blocks until the completion handler fires or the context is cancelled.
func (i IOGPUMetalBuffer) InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizeArgsArgsSizeDeallocatorSync(ctx context.Context, device objectivec.IObject, pointer unsafe.Pointer, length uint64, alignment uint32, options uint64, size uint64, address uint64, tag uint64, size2 int64, args *IOGPUNewResourceArgs, size3 uint32) error {
	done := make(chan struct{}, 1)
	i.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizeArgsArgsSizeDeallocator(device, pointer, length, alignment, options, size, address, tag, size2, args, size3, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizePlacementSparseResidencyBytesArgsArgsSizeDeallocatorSync is a synchronous wrapper around [IOGPUMetalBuffer.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizePlacementSparseResidencyBytesArgsArgsSizeDeallocator].
// It blocks until the completion handler fires or the context is cancelled.
func (i IOGPUMetalBuffer) InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizePlacementSparseResidencyBytesArgsArgsSizeDeallocatorSync(ctx context.Context, device objectivec.IObject, pointer unsafe.Pointer, length uint64, alignment uint32, options uint64, size uint64, address uint64, tag uint64, size2 int64, bytes uint64, args *IOGPUNewResourceArgs, size3 uint32) error {
	done := make(chan struct{}, 1)
	i.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizePlacementSparseResidencyBytesArgsArgsSizeDeallocator(device, pointer, length, alignment, options, size, address, tag, size2, bytes, args, size3, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// InitWithDevicePointerLengthOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocatorSync is a synchronous wrapper around [IOGPUMetalBuffer.InitWithDevicePointerLengthOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocator].
// It blocks until the completion handler fires or the context is cancelled.
func (i IOGPUMetalBuffer) InitWithDevicePointerLengthOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocatorSync(ctx context.Context, device objectivec.IObject, pointer unsafe.Pointer, length uint64, options uint64, size uint64, address uint64, args *IOGPUNewResourceArgs, size2 uint32) error {
	done := make(chan struct{}, 1)
	i.InitWithDevicePointerLengthOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocator(device, pointer, length, options, size, address, args, size2, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// InitWithDevicePointerLengthOptionsSysMemSizeVidMemSizeArgsArgsSizeDeallocatorSync is a synchronous wrapper around [IOGPUMetalBuffer.InitWithDevicePointerLengthOptionsSysMemSizeVidMemSizeArgsArgsSizeDeallocator].
// It blocks until the completion handler fires or the context is cancelled.
func (i IOGPUMetalBuffer) InitWithDevicePointerLengthOptionsSysMemSizeVidMemSizeArgsArgsSizeDeallocatorSync(ctx context.Context, device objectivec.IObject, pointer unsafe.Pointer, length uint64, options uint64, size uint64, size2 uint64, args *IOGPUNewResourceArgs, size3 uint32) error {
	done := make(chan struct{}, 1)
	i.InitWithDevicePointerLengthOptionsSysMemSizeVidMemSizeArgsArgsSizeDeallocator(device, pointer, length, options, size, size2, args, size3, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
