// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalTexture] class.
var (
	_IOGPUMetalTextureClass     IOGPUMetalTextureClass
	_IOGPUMetalTextureClassOnce sync.Once
)

func getIOGPUMetalTextureClass() IOGPUMetalTextureClass {
	_IOGPUMetalTextureClassOnce.Do(func() {
		_IOGPUMetalTextureClass = IOGPUMetalTextureClass{class: objc.GetClass("IOGPUMetalTexture")}
	})
	return _IOGPUMetalTextureClass
}

// GetIOGPUMetalTextureClass returns the class object for IOGPUMetalTexture.
func GetIOGPUMetalTextureClass() IOGPUMetalTextureClass {
	return getIOGPUMetalTextureClass()
}

type IOGPUMetalTextureClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalTextureClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalTextureClass) Alloc() IOGPUMetalTexture {
	rv := objc.SendIfResponds[IOGPUMetalTexture](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalTexture.AllowGPUOptimizedContents]
//   - [IOGPUMetalTexture.ArrayLength]
//   - [IOGPUMetalTexture.Buffer]
//   - [IOGPUMetalTexture.BufferBytesPerRow]
//   - [IOGPUMetalTexture.BufferOffset]
//   - [IOGPUMetalTexture.CompressionFeedback]
//   - [IOGPUMetalTexture.CopyFromPixelsRowBytesImageBytesToSliceMipmapLevelOriginSize]
//   - [IOGPUMetalTexture.CopyFromSliceMipmapLevelOriginSizeToPixelsRowBytesImageBytes]
//   - [IOGPUMetalTexture.Depth]
//   - [IOGPUMetalTexture.DidModifyData]
//   - [IOGPUMetalTexture.FormattedDescription]
//   - [IOGPUMetalTexture.GetBytesBytesPerRowFromRegionMipmapLevel]
//   - [IOGPUMetalTexture.GpuHandle]
//   - [IOGPUMetalTexture.GpuResourceID]
//   - [IOGPUMetalTexture.Height]
//   - [IOGPUMetalTexture.Iosurface]
//   - [IOGPUMetalTexture.IosurfacePlane]
//   - [IOGPUMetalTexture.IsCompressed]
//   - [IOGPUMetalTexture.IsDrawable]
//   - [IOGPUMetalTexture.IsFramebufferOnly]
//   - [IOGPUMetalTexture.IsShareable]
//   - [IOGPUMetalTexture.IsSparse]
//   - [IOGPUMetalTexture.MipmapLevelCount]
//   - [IOGPUMetalTexture.NewRemoteTextureViewForDevice]
//   - [IOGPUMetalTexture.NewSharedTextureHandle]
//   - [IOGPUMetalTexture.NumFaces]
//   - [IOGPUMetalTexture.ParentRelativeLevel]
//   - [IOGPUMetalTexture.ParentRelativeSlice]
//   - [IOGPUMetalTexture.ParentTexture]
//   - [IOGPUMetalTexture.PixelFormat]
//   - [IOGPUMetalTexture.PlacementSparsePageSize]
//   - [IOGPUMetalTexture.RemoteStorageTexture]
//   - [IOGPUMetalTexture.ReplaceRegionMipmapLevelWithBytesBytesPerRow]
//   - [IOGPUMetalTexture.RootResource]
//   - [IOGPUMetalTexture.RootResourceIsSuballocatedBuffer]
//   - [IOGPUMetalTexture.Rotation]
//   - [IOGPUMetalTexture.SampleCount]
//   - [IOGPUMetalTexture.Swizzle]
//   - [IOGPUMetalTexture.SwizzleKey]
//   - [IOGPUMetalTexture.TextureType]
//   - [IOGPUMetalTexture.UniqueIdentifier]
//   - [IOGPUMetalTexture.Usage]
//   - [IOGPUMetalTexture.Width]
//   - [IOGPUMetalTexture.InitWithBufferDescriptorOffsetBytesPerRow]
//   - [IOGPUMetalTexture.InitWithBufferDescriptorSysMemOffsetSysMemRowBytesVidMemSizeVidMemRowBytesArgsArgsSize]
//   - [IOGPUMetalTexture.InitWithBufferDescriptorSysMemOffsetSysMemRowBytesVidMemSizeVidMemRowBytesArgsArgsSizeIsStrideTexture]
//   - [IOGPUMetalTexture.InitWithCompressedTexturePixelFormatTextureTypeLevelSlice]
//   - [IOGPUMetalTexture.InitWithDeviceDescriptorIosurfacePlaneFieldArgsArgsSize]
//   - [IOGPUMetalTexture.InitWithDeviceDescriptorPlacementSparseBytesPlacementSparsePageSizePlacementSparseMetaDataBytesArgsArgsSize]
//   - [IOGPUMetalTexture.InitWithDeviceDescriptorPlacementSparseBytesPlacementSparsePageSizePlacementSparseMetaDataBytesPlacementSparseResidencyBytesArgsArgsSize]
//   - [IOGPUMetalTexture.InitWithDeviceDescriptorSysMemPointerSysMemSizeSysMemLengthSysMemRowBytesArgsArgsSizeDeallocator]
//   - [IOGPUMetalTexture.InitWithDeviceDescriptorSysMemSizeSysMemRowBytesVidMemSizeVidMemRowBytesArgsArgsSize]
//   - [IOGPUMetalTexture.InitWithDeviceRemoteStorageTextureArgsArgsSize]
//   - [IOGPUMetalTexture.InitWithHeapResourceOffsetLengthDeviceDescriptor]
//   - [IOGPUMetalTexture.InitWithPrimaryBufferHeapIndexBufferIndexBufferOffsetLengthDescriptorSysMemRowBytesVidMemSizeVidMemRowBytesArgsArgsSize]
//   - [IOGPUMetalTexture.InitWithTexturePixelFormat]
//   - [IOGPUMetalTexture.InitWithTexturePixelFormatTextureTypeLevelsSlices]
//   - [IOGPUMetalTexture.InitWithTexturePixelFormatTextureTypeLevelsSlicesSwizzle]
//   - [IOGPUMetalTexture.InitWithTextureInternalPixelFormatTextureTypeLevelsSlicesSwizzleCompressedView]
//   - [IOGPUMetalTexture.FramebufferOnly]
//   - [IOGPUMetalTexture.Shareable]
type IOGPUMetalTexture struct {
	IOGPUMetalResource
}

// IOGPUMetalTextureFromID constructs a [IOGPUMetalTexture] from an objc.ID.
func IOGPUMetalTextureFromID(id objc.ID) IOGPUMetalTexture {
	return IOGPUMetalTexture{IOGPUMetalResource: IOGPUMetalResourceFromID(id)}
}

// Ensure IOGPUMetalTexture implements IIOGPUMetalTexture.
var _ IIOGPUMetalTexture = IOGPUMetalTexture{}

// An interface definition for the [IOGPUMetalTexture] class.
//
// # Methods
//
//   - [IIOGPUMetalTexture.AllowGPUOptimizedContents]
//   - [IIOGPUMetalTexture.ArrayLength]
//   - [IIOGPUMetalTexture.Buffer]
//   - [IIOGPUMetalTexture.BufferBytesPerRow]
//   - [IIOGPUMetalTexture.BufferOffset]
//   - [IIOGPUMetalTexture.CompressionFeedback]
//   - [IIOGPUMetalTexture.CopyFromPixelsRowBytesImageBytesToSliceMipmapLevelOriginSize]
//   - [IIOGPUMetalTexture.CopyFromSliceMipmapLevelOriginSizeToPixelsRowBytesImageBytes]
//   - [IIOGPUMetalTexture.Depth]
//   - [IIOGPUMetalTexture.DidModifyData]
//   - [IIOGPUMetalTexture.FormattedDescription]
//   - [IIOGPUMetalTexture.GetBytesBytesPerRowFromRegionMipmapLevel]
//   - [IIOGPUMetalTexture.GpuHandle]
//   - [IIOGPUMetalTexture.GpuResourceID]
//   - [IIOGPUMetalTexture.Height]
//   - [IIOGPUMetalTexture.Iosurface]
//   - [IIOGPUMetalTexture.IosurfacePlane]
//   - [IIOGPUMetalTexture.IsCompressed]
//   - [IIOGPUMetalTexture.IsDrawable]
//   - [IIOGPUMetalTexture.IsFramebufferOnly]
//   - [IIOGPUMetalTexture.IsShareable]
//   - [IIOGPUMetalTexture.IsSparse]
//   - [IIOGPUMetalTexture.MipmapLevelCount]
//   - [IIOGPUMetalTexture.NewRemoteTextureViewForDevice]
//   - [IIOGPUMetalTexture.NewSharedTextureHandle]
//   - [IIOGPUMetalTexture.NumFaces]
//   - [IIOGPUMetalTexture.ParentRelativeLevel]
//   - [IIOGPUMetalTexture.ParentRelativeSlice]
//   - [IIOGPUMetalTexture.ParentTexture]
//   - [IIOGPUMetalTexture.PixelFormat]
//   - [IIOGPUMetalTexture.PlacementSparsePageSize]
//   - [IIOGPUMetalTexture.RemoteStorageTexture]
//   - [IIOGPUMetalTexture.ReplaceRegionMipmapLevelWithBytesBytesPerRow]
//   - [IIOGPUMetalTexture.RootResource]
//   - [IIOGPUMetalTexture.RootResourceIsSuballocatedBuffer]
//   - [IIOGPUMetalTexture.Rotation]
//   - [IIOGPUMetalTexture.SampleCount]
//   - [IIOGPUMetalTexture.Swizzle]
//   - [IIOGPUMetalTexture.SwizzleKey]
//   - [IIOGPUMetalTexture.TextureType]
//   - [IIOGPUMetalTexture.UniqueIdentifier]
//   - [IIOGPUMetalTexture.Usage]
//   - [IIOGPUMetalTexture.Width]
//   - [IIOGPUMetalTexture.InitWithBufferDescriptorOffsetBytesPerRow]
//   - [IIOGPUMetalTexture.InitWithBufferDescriptorSysMemOffsetSysMemRowBytesVidMemSizeVidMemRowBytesArgsArgsSize]
//   - [IIOGPUMetalTexture.InitWithBufferDescriptorSysMemOffsetSysMemRowBytesVidMemSizeVidMemRowBytesArgsArgsSizeIsStrideTexture]
//   - [IIOGPUMetalTexture.InitWithCompressedTexturePixelFormatTextureTypeLevelSlice]
//   - [IIOGPUMetalTexture.InitWithDeviceDescriptorIosurfacePlaneFieldArgsArgsSize]
//   - [IIOGPUMetalTexture.InitWithDeviceDescriptorPlacementSparseBytesPlacementSparsePageSizePlacementSparseMetaDataBytesArgsArgsSize]
//   - [IIOGPUMetalTexture.InitWithDeviceDescriptorPlacementSparseBytesPlacementSparsePageSizePlacementSparseMetaDataBytesPlacementSparseResidencyBytesArgsArgsSize]
//   - [IIOGPUMetalTexture.InitWithDeviceDescriptorSysMemPointerSysMemSizeSysMemLengthSysMemRowBytesArgsArgsSizeDeallocator]
//   - [IIOGPUMetalTexture.InitWithDeviceDescriptorSysMemSizeSysMemRowBytesVidMemSizeVidMemRowBytesArgsArgsSize]
//   - [IIOGPUMetalTexture.InitWithDeviceRemoteStorageTextureArgsArgsSize]
//   - [IIOGPUMetalTexture.InitWithHeapResourceOffsetLengthDeviceDescriptor]
//   - [IIOGPUMetalTexture.InitWithPrimaryBufferHeapIndexBufferIndexBufferOffsetLengthDescriptorSysMemRowBytesVidMemSizeVidMemRowBytesArgsArgsSize]
//   - [IIOGPUMetalTexture.InitWithTexturePixelFormat]
//   - [IIOGPUMetalTexture.InitWithTexturePixelFormatTextureTypeLevelsSlices]
//   - [IIOGPUMetalTexture.InitWithTexturePixelFormatTextureTypeLevelsSlicesSwizzle]
//   - [IIOGPUMetalTexture.InitWithTextureInternalPixelFormatTextureTypeLevelsSlicesSwizzleCompressedView]
//   - [IIOGPUMetalTexture.FramebufferOnly]
//   - [IIOGPUMetalTexture.Shareable]
type IIOGPUMetalTexture interface {
	IIOGPUMetalResource

	// Topic: Methods

	AllowGPUOptimizedContents() bool
	ArrayLength() uint64
	Buffer() IIOGPUMetalBuffer
	BufferBytesPerRow() uint64
	BufferOffset() uint64
	CompressionFeedback() int64
	CopyFromPixelsRowBytesImageBytesToSliceMipmapLevelOriginSize(pixels unsafe.Pointer, bytes uint64, bytes2 uint64, slice uint64, level uint64, origin unsafe.Pointer, size unsafe.Pointer)
	CopyFromSliceMipmapLevelOriginSizeToPixelsRowBytesImageBytes(slice uint64, level uint64, origin unsafe.Pointer, size unsafe.Pointer, pixels unsafe.Pointer, bytes uint64, bytes2 uint64)
	Depth() uint64
	DidModifyData()
	FormattedDescription(description uint64) objectivec.IObject
	GetBytesBytesPerRowFromRegionMipmapLevel(bytes unsafe.Pointer, row uint64, region unsafe.Pointer, level uint64)
	GpuHandle() uint64
	GpuResourceID() metal.MTLResourceID
	Height() uint64
	Iosurface() iosurface.IOSurfaceRef
	IosurfacePlane() uint64
	IsCompressed() bool
	IsDrawable() bool
	IsFramebufferOnly() bool
	IsShareable() bool
	IsSparse() bool
	MipmapLevelCount() uint64
	NewRemoteTextureViewForDevice(device objectivec.IObject) objectivec.IObject
	NewSharedTextureHandle() objectivec.IObject
	NumFaces() uint64
	ParentRelativeLevel() uint64
	ParentRelativeSlice() uint64
	ParentTexture() IIOGPUMetalTexture
	PixelFormat() uint64
	PlacementSparsePageSize() int64
	RemoteStorageTexture() unsafe.Pointer
	ReplaceRegionMipmapLevelWithBytesBytesPerRow(region unsafe.Pointer, level uint64, bytes unsafe.Pointer, row uint64)
	RootResource() IIOGPUMetalResource
	RootResourceIsSuballocatedBuffer() bool
	Rotation() uint64
	SampleCount() uint64
	Swizzle() unsafe.Pointer
	SwizzleKey() uint32
	TextureType() uint64
	UniqueIdentifier() uint64
	Usage() uint64
	Width() uint64
	InitWithBufferDescriptorOffsetBytesPerRow(buffer objectivec.IObject, descriptor objectivec.IObject, offset uint64, row uint64) IOGPUMetalTexture
	InitWithBufferDescriptorSysMemOffsetSysMemRowBytesVidMemSizeVidMemRowBytesArgsArgsSize(buffer objectivec.IObject, descriptor objectivec.IObject, offset uint64, bytes uint64, size uint64, bytes2 uint64, args *IOGPUNewResourceArgs, size2 uint32) IOGPUMetalTexture
	InitWithBufferDescriptorSysMemOffsetSysMemRowBytesVidMemSizeVidMemRowBytesArgsArgsSizeIsStrideTexture(buffer objectivec.IObject, descriptor objectivec.IObject, offset uint64, bytes uint64, size uint64, bytes2 uint64, args *IOGPUNewResourceArgs, size2 uint32, texture bool) IOGPUMetalTexture
	InitWithCompressedTexturePixelFormatTextureTypeLevelSlice(texture objectivec.IObject, format uint64, type_ uint64, level uint64, slice uint64) IOGPUMetalTexture
	InitWithDeviceDescriptorIosurfacePlaneFieldArgsArgsSize(device objectivec.IObject, descriptor objectivec.IObject, iosurface iosurface.IOSurfaceRef, plane uint32, field uint32, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalTexture
	InitWithDeviceDescriptorPlacementSparseBytesPlacementSparsePageSizePlacementSparseMetaDataBytesArgsArgsSize(device objectivec.IObject, descriptor objectivec.IObject, bytes uint64, size int64, bytes2 uint64, args *IOGPUNewResourceArgs, size2 uint32) IOGPUMetalTexture
	InitWithDeviceDescriptorPlacementSparseBytesPlacementSparsePageSizePlacementSparseMetaDataBytesPlacementSparseResidencyBytesArgsArgsSize(device objectivec.IObject, descriptor objectivec.IObject, bytes uint64, size int64, bytes2 uint64, bytes3 uint64, args *IOGPUNewResourceArgs, size2 uint32) IOGPUMetalTexture
	InitWithDeviceDescriptorSysMemPointerSysMemSizeSysMemLengthSysMemRowBytesArgsArgsSizeDeallocator(device objectivec.IObject, descriptor objectivec.IObject, pointer unsafe.Pointer, size uint64, length uint64, bytes uint64, args *IOGPUNewResourceArgs, size2 uint32, deallocator VoidHandler) IOGPUMetalTexture
	InitWithDeviceDescriptorSysMemSizeSysMemRowBytesVidMemSizeVidMemRowBytesArgsArgsSize(device objectivec.IObject, descriptor objectivec.IObject, size uint64, bytes uint64, size2 uint64, bytes2 uint64, args *IOGPUNewResourceArgs, size3 uint32) IOGPUMetalTexture
	InitWithDeviceRemoteStorageTextureArgsArgsSize(device objectivec.IObject, texture objectivec.IObject, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalTexture
	InitWithHeapResourceOffsetLengthDeviceDescriptor(heap objectivec.IObject, resource objectivec.IObject, offset uint64, length uint64, device objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalTexture
	InitWithPrimaryBufferHeapIndexBufferIndexBufferOffsetLengthDescriptorSysMemRowBytesVidMemSizeVidMemRowBytesArgsArgsSize(buffer objectivec.IObject, index int16, index2 int16, offset uint64, length uint64, descriptor objectivec.IObject, bytes uint64, size uint64, bytes2 uint64, args *IOGPUNewResourceArgs, size2 uint32) IOGPUMetalTexture
	InitWithTexturePixelFormat(texture objectivec.IObject, format uint64) IOGPUMetalTexture
	InitWithTexturePixelFormatTextureTypeLevelsSlices(texture objectivec.IObject, format uint64, type_ uint64, levels foundation.NSRange, slices foundation.NSRange) IOGPUMetalTexture
	InitWithTexturePixelFormatTextureTypeLevelsSlicesSwizzle(texture objectivec.IObject, format uint64, type_ uint64, levels foundation.NSRange, slices foundation.NSRange, swizzle unsafe.Pointer) IOGPUMetalTexture
	InitWithTextureInternalPixelFormatTextureTypeLevelsSlicesSwizzleCompressedView(internal objectivec.IObject, format uint64, type_ uint64, levels foundation.NSRange, slices foundation.NSRange, swizzle unsafe.Pointer, view bool) IOGPUMetalTexture
	FramebufferOnly() bool
	Shareable() bool
}

// Init initializes the instance.
func (i IOGPUMetalTexture) Init() IOGPUMetalTexture {
	rv := objc.SendIfResponds[IOGPUMetalTexture](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalTexture) Autorelease() IOGPUMetalTexture {
	rv := objc.SendIfResponds[IOGPUMetalTexture](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalTexture creates a new IOGPUMetalTexture instance.
func NewIOGPUMetalTexture() IOGPUMetalTexture {
	class := getIOGPUMetalTextureClass()
	rv := objc.SendIfResponds[IOGPUMetalTexture](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalTextureMemorylessDescriptor(memoryless objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalTexture {
	instance := getIOGPUMetalTextureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initMemoryless:descriptor:"), memoryless, descriptor)
	return IOGPUMetalTextureFromID(rv)
}

func NewGPUMetalTextureStandinWithDevice(device objectivec.IObject) IOGPUMetalTexture {
	instance := getIOGPUMetalTextureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initStandinWithDevice:"), device)
	return IOGPUMetalTextureFromID(rv)
}

func NewGPUMetalTextureWithBufferDescriptorOffsetBytesPerRow(buffer objectivec.IObject, descriptor objectivec.IObject, offset uint64, row uint64) IOGPUMetalTexture {
	instance := getIOGPUMetalTextureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBuffer:descriptor:offset:bytesPerRow:"), buffer, descriptor, offset, row)
	return IOGPUMetalTextureFromID(rv)
}

func NewGPUMetalTextureWithBufferDescriptorSysMemOffsetSysMemRowBytesVidMemSizeVidMemRowBytesArgsArgsSize(buffer objectivec.IObject, descriptor objectivec.IObject, offset uint64, bytes uint64, size uint64, bytes2 uint64, args *IOGPUNewResourceArgs, size2 uint32) IOGPUMetalTexture {
	instance := getIOGPUMetalTextureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBuffer:descriptor:sysMemOffset:sysMemRowBytes:vidMemSize:vidMemRowBytes:args:argsSize:"), buffer, descriptor, offset, bytes, size, bytes2, unsafe.Pointer(args), size2)
	return IOGPUMetalTextureFromID(rv)
}

func NewGPUMetalTextureWithBufferDescriptorSysMemOffsetSysMemRowBytesVidMemSizeVidMemRowBytesArgsArgsSizeIsStrideTexture(buffer objectivec.IObject, descriptor objectivec.IObject, offset uint64, bytes uint64, size uint64, bytes2 uint64, args *IOGPUNewResourceArgs, size2 uint32, texture bool) IOGPUMetalTexture {
	instance := getIOGPUMetalTextureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBuffer:descriptor:sysMemOffset:sysMemRowBytes:vidMemSize:vidMemRowBytes:args:argsSize:isStrideTexture:"), buffer, descriptor, offset, bytes, size, bytes2, unsafe.Pointer(args), size2, texture)
	return IOGPUMetalTextureFromID(rv)
}

func NewGPUMetalTextureWithCompressedTexturePixelFormatTextureTypeLevelSlice(texture objectivec.IObject, format uint64, type_ uint64, level uint64, slice uint64) IOGPUMetalTexture {
	instance := getIOGPUMetalTextureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCompressedTexture:pixelFormat:textureType:level:slice:"), texture, format, type_, level, slice)
	return IOGPUMetalTextureFromID(rv)
}

func NewGPUMetalTextureWithDeviceDescriptorIosurfacePlaneFieldArgsArgsSize(device objectivec.IObject, descriptor objectivec.IObject, iosurface iosurface.IOSurfaceRef, plane uint32, field uint32, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalTexture {
	instance := getIOGPUMetalTextureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:descriptor:iosurface:plane:field:args:argsSize:"), device, descriptor, iosurface, plane, field, unsafe.Pointer(args), size)
	return IOGPUMetalTextureFromID(rv)
}

func NewGPUMetalTextureWithDeviceDescriptorPlacementSparseBytesPlacementSparsePageSizePlacementSparseMetaDataBytesArgsArgsSize(device objectivec.IObject, descriptor objectivec.IObject, bytes uint64, size int64, bytes2 uint64, args *IOGPUNewResourceArgs, size2 uint32) IOGPUMetalTexture {
	instance := getIOGPUMetalTextureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:descriptor:placementSparseBytes:placementSparsePageSize:placementSparseMetaDataBytes:args:argsSize:"), device, descriptor, bytes, size, bytes2, unsafe.Pointer(args), size2)
	return IOGPUMetalTextureFromID(rv)
}

func NewGPUMetalTextureWithDeviceDescriptorPlacementSparseBytesPlacementSparsePageSizePlacementSparseMetaDataBytesPlacementSparseResidencyBytesArgsArgsSize(device objectivec.IObject, descriptor objectivec.IObject, bytes uint64, size int64, bytes2 uint64, bytes3 uint64, args *IOGPUNewResourceArgs, size2 uint32) IOGPUMetalTexture {
	instance := getIOGPUMetalTextureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:descriptor:placementSparseBytes:placementSparsePageSize:placementSparseMetaDataBytes:placementSparseResidencyBytes:args:argsSize:"), device, descriptor, bytes, size, bytes2, bytes3, unsafe.Pointer(args), size2)
	return IOGPUMetalTextureFromID(rv)
}

func NewGPUMetalTextureWithDeviceDescriptorSysMemSizeSysMemRowBytesVidMemSizeVidMemRowBytesArgsArgsSize(device objectivec.IObject, descriptor objectivec.IObject, size uint64, bytes uint64, size2 uint64, bytes2 uint64, args *IOGPUNewResourceArgs, size3 uint32) IOGPUMetalTexture {
	instance := getIOGPUMetalTextureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:descriptor:sysMemSize:sysMemRowBytes:vidMemSize:vidMemRowBytes:args:argsSize:"), device, descriptor, size, bytes, size2, bytes2, unsafe.Pointer(args), size3)
	return IOGPUMetalTextureFromID(rv)
}

func NewGPUMetalTextureWithDeviceOptionsArgsArgsSize(device objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalTexture {
	instance := getIOGPUMetalTextureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:options:args:argsSize:"), device, options, unsafe.Pointer(args), size)
	return IOGPUMetalTextureFromID(rv)
}

func NewGPUMetalTextureWithDeviceRemoteStorageResourceOptionsArgsArgsSize(device objectivec.IObject, resource objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalTexture {
	instance := getIOGPUMetalTextureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:remoteStorageResource:options:args:argsSize:"), device, resource, options, unsafe.Pointer(args), size)
	return IOGPUMetalTextureFromID(rv)
}

func NewGPUMetalTextureWithDeviceRemoteStorageTextureArgsArgsSize(device objectivec.IObject, texture objectivec.IObject, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalTexture {
	instance := getIOGPUMetalTextureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:remoteStorageTexture:args:argsSize:"), device, texture, unsafe.Pointer(args), size)
	return IOGPUMetalTextureFromID(rv)
}

func NewGPUMetalTextureWithHeapResourceOffsetLengthDeviceDescriptor(heap objectivec.IObject, resource objectivec.IObject, offset uint64, length uint64, device objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalTexture {
	instance := getIOGPUMetalTextureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithHeap:resource:offset:length:device:descriptor:"), heap, resource, offset, length, device, descriptor)
	return IOGPUMetalTextureFromID(rv)
}

func NewGPUMetalTextureWithPrimaryBufferHeapIndexBufferIndexBufferOffsetLengthDescriptorSysMemRowBytesVidMemSizeVidMemRowBytesArgsArgsSize(buffer objectivec.IObject, index int16, index2 int16, offset uint64, length uint64, descriptor objectivec.IObject, bytes uint64, size uint64, bytes2 uint64, args *IOGPUNewResourceArgs, size2 uint32) IOGPUMetalTexture {
	instance := getIOGPUMetalTextureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithPrimaryBuffer:heapIndex:bufferIndex:bufferOffset:length:descriptor:sysMemRowBytes:vidMemSize:vidMemRowBytes:args:argsSize:"), buffer, index, index2, offset, length, descriptor, bytes, size, bytes2, unsafe.Pointer(args), size2)
	return IOGPUMetalTextureFromID(rv)
}

func NewGPUMetalTextureWithResource(resource objectivec.IObject) IOGPUMetalTexture {
	instance := getIOGPUMetalTextureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return IOGPUMetalTextureFromID(rv)
}

func NewGPUMetalTextureWithTextureInternalPixelFormatTextureTypeLevelsSlicesSwizzleCompressedView(internal objectivec.IObject, format uint64, type_ uint64, levels foundation.NSRange, slices foundation.NSRange, swizzle unsafe.Pointer, view bool) IOGPUMetalTexture {
	instance := getIOGPUMetalTextureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithTextureInternal:pixelFormat:textureType:levels:slices:swizzle:compressedView:"), internal, format, type_, levels, slices, swizzle, view)
	return IOGPUMetalTextureFromID(rv)
}

func NewGPUMetalTextureWithTexturePixelFormat(texture objectivec.IObject, format uint64) IOGPUMetalTexture {
	instance := getIOGPUMetalTextureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithTexture:pixelFormat:"), texture, format)
	return IOGPUMetalTextureFromID(rv)
}

func NewGPUMetalTextureWithTexturePixelFormatTextureTypeLevelsSlices(texture objectivec.IObject, format uint64, type_ uint64, levels foundation.NSRange, slices foundation.NSRange) IOGPUMetalTexture {
	instance := getIOGPUMetalTextureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithTexture:pixelFormat:textureType:levels:slices:"), texture, format, type_, levels, slices)
	return IOGPUMetalTextureFromID(rv)
}

func NewGPUMetalTextureWithTexturePixelFormatTextureTypeLevelsSlicesSwizzle(texture objectivec.IObject, format uint64, type_ uint64, levels foundation.NSRange, slices foundation.NSRange, swizzle unsafe.Pointer) IOGPUMetalTexture {
	instance := getIOGPUMetalTextureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithTexture:pixelFormat:textureType:levels:slices:swizzle:"), texture, format, type_, levels, slices, swizzle)
	return IOGPUMetalTextureFromID(rv)
}

func (i IOGPUMetalTexture) CopyFromPixelsRowBytesImageBytesToSliceMipmapLevelOriginSize(pixels unsafe.Pointer, bytes uint64, bytes2 uint64, slice uint64, level uint64, origin unsafe.Pointer, size unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("copyFromPixels:rowBytes:imageBytes:toSlice:mipmapLevel:origin:size:"), pixels, bytes, bytes2, slice, level, origin, size)
}
func (i IOGPUMetalTexture) CopyFromSliceMipmapLevelOriginSizeToPixelsRowBytesImageBytes(slice uint64, level uint64, origin unsafe.Pointer, size unsafe.Pointer, pixels unsafe.Pointer, bytes uint64, bytes2 uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("copyFromSlice:mipmapLevel:origin:size:toPixels:rowBytes:imageBytes:"), slice, level, origin, size, pixels, bytes, bytes2)
}
func (i IOGPUMetalTexture) DidModifyData() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("didModifyData"))
}
func (i IOGPUMetalTexture) FormattedDescription(description uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("formattedDescription:"), description)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalTexture) GetBytesBytesPerRowFromRegionMipmapLevel(bytes unsafe.Pointer, row uint64, region unsafe.Pointer, level uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("getBytes:bytesPerRow:fromRegion:mipmapLevel:"), bytes, row, region, level)
}
func (i IOGPUMetalTexture) IsFramebufferOnly() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("isFramebufferOnly"))
	return rv
}
func (i IOGPUMetalTexture) IsShareable() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("isShareable"))
	return rv
}
func (i IOGPUMetalTexture) IsSparse() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("isSparse"))
	return rv
}
func (i IOGPUMetalTexture) NewRemoteTextureViewForDevice(device objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newRemoteTextureViewForDevice:"), device)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalTexture) NewSharedTextureHandle() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newSharedTextureHandle"))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalTexture) ReplaceRegionMipmapLevelWithBytesBytesPerRow(region unsafe.Pointer, level uint64, bytes unsafe.Pointer, row uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("replaceRegion:mipmapLevel:withBytes:bytesPerRow:"), region, level, bytes, row)
}
func (i IOGPUMetalTexture) InitWithBufferDescriptorOffsetBytesPerRow(buffer objectivec.IObject, descriptor objectivec.IObject, offset uint64, row uint64) IOGPUMetalTexture {
	rv := objc.SendIfResponds[IOGPUMetalTexture](i.ID, objc.Sel("initWithBuffer:descriptor:offset:bytesPerRow:"), buffer, descriptor, offset, row)
	return rv
}
func (i IOGPUMetalTexture) InitWithBufferDescriptorSysMemOffsetSysMemRowBytesVidMemSizeVidMemRowBytesArgsArgsSize(buffer objectivec.IObject, descriptor objectivec.IObject, offset uint64, bytes uint64, size uint64, bytes2 uint64, args *IOGPUNewResourceArgs, size2 uint32) IOGPUMetalTexture {
	rv := objc.SendIfResponds[IOGPUMetalTexture](i.ID, objc.Sel("initWithBuffer:descriptor:sysMemOffset:sysMemRowBytes:vidMemSize:vidMemRowBytes:args:argsSize:"), buffer, descriptor, offset, bytes, size, bytes2, unsafe.Pointer(args), size2)
	return rv
}
func (i IOGPUMetalTexture) InitWithBufferDescriptorSysMemOffsetSysMemRowBytesVidMemSizeVidMemRowBytesArgsArgsSizeIsStrideTexture(buffer objectivec.IObject, descriptor objectivec.IObject, offset uint64, bytes uint64, size uint64, bytes2 uint64, args *IOGPUNewResourceArgs, size2 uint32, texture bool) IOGPUMetalTexture {
	rv := objc.SendIfResponds[IOGPUMetalTexture](i.ID, objc.Sel("initWithBuffer:descriptor:sysMemOffset:sysMemRowBytes:vidMemSize:vidMemRowBytes:args:argsSize:isStrideTexture:"), buffer, descriptor, offset, bytes, size, bytes2, unsafe.Pointer(args), size2, texture)
	return rv
}
func (i IOGPUMetalTexture) InitWithCompressedTexturePixelFormatTextureTypeLevelSlice(texture objectivec.IObject, format uint64, type_ uint64, level uint64, slice uint64) IOGPUMetalTexture {
	rv := objc.SendIfResponds[IOGPUMetalTexture](i.ID, objc.Sel("initWithCompressedTexture:pixelFormat:textureType:level:slice:"), texture, format, type_, level, slice)
	return rv
}
func (i IOGPUMetalTexture) InitWithDeviceDescriptorIosurfacePlaneFieldArgsArgsSize(device objectivec.IObject, descriptor objectivec.IObject, iosurface iosurface.IOSurfaceRef, plane uint32, field uint32, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalTexture {
	rv := objc.SendIfResponds[IOGPUMetalTexture](i.ID, objc.Sel("initWithDevice:descriptor:iosurface:plane:field:args:argsSize:"), device, descriptor, iosurface, plane, field, unsafe.Pointer(args), size)
	return rv
}
func (i IOGPUMetalTexture) InitWithDeviceDescriptorPlacementSparseBytesPlacementSparsePageSizePlacementSparseMetaDataBytesArgsArgsSize(device objectivec.IObject, descriptor objectivec.IObject, bytes uint64, size int64, bytes2 uint64, args *IOGPUNewResourceArgs, size2 uint32) IOGPUMetalTexture {
	rv := objc.SendIfResponds[IOGPUMetalTexture](i.ID, objc.Sel("initWithDevice:descriptor:placementSparseBytes:placementSparsePageSize:placementSparseMetaDataBytes:args:argsSize:"), device, descriptor, bytes, size, bytes2, unsafe.Pointer(args), size2)
	return rv
}
func (i IOGPUMetalTexture) InitWithDeviceDescriptorPlacementSparseBytesPlacementSparsePageSizePlacementSparseMetaDataBytesPlacementSparseResidencyBytesArgsArgsSize(device objectivec.IObject, descriptor objectivec.IObject, bytes uint64, size int64, bytes2 uint64, bytes3 uint64, args *IOGPUNewResourceArgs, size2 uint32) IOGPUMetalTexture {
	rv := objc.SendIfResponds[IOGPUMetalTexture](i.ID, objc.Sel("initWithDevice:descriptor:placementSparseBytes:placementSparsePageSize:placementSparseMetaDataBytes:placementSparseResidencyBytes:args:argsSize:"), device, descriptor, bytes, size, bytes2, bytes3, unsafe.Pointer(args), size2)
	return rv
}

var _iogpumetaltexture_initwithdevice_descriptor_sysmempointer_sysmemsize_sysmemlength_sysmemrowbytes_args_argssize_deallocator_p8_key byte

func (i IOGPUMetalTexture) InitWithDeviceDescriptorSysMemPointerSysMemSizeSysMemLengthSysMemRowBytesArgsArgsSizeDeallocator(device objectivec.IObject, descriptor objectivec.IObject, pointer unsafe.Pointer, size uint64, length uint64, bytes uint64, args *IOGPUNewResourceArgs, size2 uint32, deallocator VoidHandler) IOGPUMetalTexture {
	_block8, _ := NewVoidBlock(deallocator)
	rv := objc.SendIfResponds[IOGPUMetalTexture](i.ID, objc.Sel("initWithDevice:descriptor:sysMemPointer:sysMemSize:sysMemLength:sysMemRowBytes:args:argsSize:deallocator:"), device, descriptor, pointer, size, length, bytes, args, size2, _block8)
	return rv
}
func (i IOGPUMetalTexture) InitWithDeviceDescriptorSysMemSizeSysMemRowBytesVidMemSizeVidMemRowBytesArgsArgsSize(device objectivec.IObject, descriptor objectivec.IObject, size uint64, bytes uint64, size2 uint64, bytes2 uint64, args *IOGPUNewResourceArgs, size3 uint32) IOGPUMetalTexture {
	rv := objc.SendIfResponds[IOGPUMetalTexture](i.ID, objc.Sel("initWithDevice:descriptor:sysMemSize:sysMemRowBytes:vidMemSize:vidMemRowBytes:args:argsSize:"), device, descriptor, size, bytes, size2, bytes2, unsafe.Pointer(args), size3)
	return rv
}
func (i IOGPUMetalTexture) InitWithDeviceRemoteStorageTextureArgsArgsSize(device objectivec.IObject, texture objectivec.IObject, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalTexture {
	rv := objc.SendIfResponds[IOGPUMetalTexture](i.ID, objc.Sel("initWithDevice:remoteStorageTexture:args:argsSize:"), device, texture, unsafe.Pointer(args), size)
	return rv
}
func (i IOGPUMetalTexture) InitWithHeapResourceOffsetLengthDeviceDescriptor(heap objectivec.IObject, resource objectivec.IObject, offset uint64, length uint64, device objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalTexture {
	rv := objc.SendIfResponds[IOGPUMetalTexture](i.ID, objc.Sel("initWithHeap:resource:offset:length:device:descriptor:"), heap, resource, offset, length, device, descriptor)
	return rv
}
func (i IOGPUMetalTexture) InitWithPrimaryBufferHeapIndexBufferIndexBufferOffsetLengthDescriptorSysMemRowBytesVidMemSizeVidMemRowBytesArgsArgsSize(buffer objectivec.IObject, index int16, index2 int16, offset uint64, length uint64, descriptor objectivec.IObject, bytes uint64, size uint64, bytes2 uint64, args *IOGPUNewResourceArgs, size2 uint32) IOGPUMetalTexture {
	rv := objc.SendIfResponds[IOGPUMetalTexture](i.ID, objc.Sel("initWithPrimaryBuffer:heapIndex:bufferIndex:bufferOffset:length:descriptor:sysMemRowBytes:vidMemSize:vidMemRowBytes:args:argsSize:"), buffer, index, index2, offset, length, descriptor, bytes, size, bytes2, unsafe.Pointer(args), size2)
	return rv
}
func (i IOGPUMetalTexture) InitWithTexturePixelFormat(texture objectivec.IObject, format uint64) IOGPUMetalTexture {
	rv := objc.SendIfResponds[IOGPUMetalTexture](i.ID, objc.Sel("initWithTexture:pixelFormat:"), texture, format)
	return rv
}
func (i IOGPUMetalTexture) InitWithTexturePixelFormatTextureTypeLevelsSlices(texture objectivec.IObject, format uint64, type_ uint64, levels foundation.NSRange, slices foundation.NSRange) IOGPUMetalTexture {
	rv := objc.SendIfResponds[IOGPUMetalTexture](i.ID, objc.Sel("initWithTexture:pixelFormat:textureType:levels:slices:"), texture, format, type_, levels, slices)
	return rv
}
func (i IOGPUMetalTexture) InitWithTexturePixelFormatTextureTypeLevelsSlicesSwizzle(texture objectivec.IObject, format uint64, type_ uint64, levels foundation.NSRange, slices foundation.NSRange, swizzle unsafe.Pointer) IOGPUMetalTexture {
	rv := objc.SendIfResponds[IOGPUMetalTexture](i.ID, objc.Sel("initWithTexture:pixelFormat:textureType:levels:slices:swizzle:"), texture, format, type_, levels, slices, swizzle)
	return rv
}
func (i IOGPUMetalTexture) InitWithTextureInternalPixelFormatTextureTypeLevelsSlicesSwizzleCompressedView(internal objectivec.IObject, format uint64, type_ uint64, levels foundation.NSRange, slices foundation.NSRange, swizzle unsafe.Pointer, view bool) IOGPUMetalTexture {
	rv := objc.SendIfResponds[IOGPUMetalTexture](i.ID, objc.Sel("initWithTextureInternal:pixelFormat:textureType:levels:slices:swizzle:compressedView:"), internal, format, type_, levels, slices, swizzle, view)
	return rv
}

func (_IOGPUMetalTextureClass IOGPUMetalTextureClass) InitNewTextureDataWithDeviceDescriptorSysMemSizeSysMemRowBytesVidMemSizeVidMemRowBytesArgs(device objectivec.IObject, descriptor objectivec.IObject, size uint64, bytes uint64, size2 uint64, bytes2 uint64, args *IOGPUNewResourceArgs) IOGPUMetalTexture {
	rv := objc.SendIfResponds[IOGPUMetalTexture](objc.ID(_IOGPUMetalTextureClass.class), objc.Sel("initNewTextureDataWithDevice:descriptor:sysMemSize:sysMemRowBytes:vidMemSize:vidMemRowBytes:args:"), device, descriptor, size, bytes, size2, bytes2, unsafe.Pointer(args))
	return rv
}

func (i IOGPUMetalTexture) AllowGPUOptimizedContents() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("allowGPUOptimizedContents"))
	return rv
}
func (i IOGPUMetalTexture) ArrayLength() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("arrayLength"))
	return rv
}
func (i IOGPUMetalTexture) Buffer() IIOGPUMetalBuffer {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("buffer"))
	return IOGPUMetalBufferFromID(objc.ID(rv))
}
func (i IOGPUMetalTexture) BufferBytesPerRow() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("bufferBytesPerRow"))
	return rv
}
func (i IOGPUMetalTexture) BufferOffset() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("bufferOffset"))
	return rv
}
func (i IOGPUMetalTexture) CompressionFeedback() int64 {
	rv := objc.SendIfResponds[int64](i.ID, objc.Sel("compressionFeedback"))
	return rv
}
func (i IOGPUMetalTexture) Depth() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("depth"))
	return rv
}
func (i IOGPUMetalTexture) FramebufferOnly() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("framebufferOnly"))
	return rv
}
func (i IOGPUMetalTexture) GpuHandle() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("gpuHandle"))
	return rv
}
func (i IOGPUMetalTexture) GpuResourceID() metal.MTLResourceID {
	rv := objc.SendIfResponds[metal.MTLResourceID](i.ID, objc.Sel("gpuResourceID"))
	return metal.MTLResourceID(rv)
}
func (i IOGPUMetalTexture) Height() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("height"))
	return rv
}
func (i IOGPUMetalTexture) Iosurface() iosurface.IOSurfaceRef {
	rv := objc.SendIfResponds[iosurface.IOSurfaceRef](i.ID, objc.Sel("iosurface"))
	return iosurface.IOSurfaceRef(rv)
}
func (i IOGPUMetalTexture) IosurfacePlane() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("iosurfacePlane"))
	return rv
}
func (i IOGPUMetalTexture) IsCompressed() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("isCompressed"))
	return rv
}
func (i IOGPUMetalTexture) IsDrawable() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("isDrawable"))
	return rv
}
func (i IOGPUMetalTexture) MipmapLevelCount() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("mipmapLevelCount"))
	return rv
}
func (i IOGPUMetalTexture) NumFaces() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("numFaces"))
	return rv
}
func (i IOGPUMetalTexture) ParentRelativeLevel() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("parentRelativeLevel"))
	return rv
}
func (i IOGPUMetalTexture) ParentRelativeSlice() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("parentRelativeSlice"))
	return rv
}
func (i IOGPUMetalTexture) ParentTexture() IIOGPUMetalTexture {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("parentTexture"))
	return IOGPUMetalTextureFromID(objc.ID(rv))
}
func (i IOGPUMetalTexture) PixelFormat() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("pixelFormat"))
	return rv
}
func (i IOGPUMetalTexture) PlacementSparsePageSize() int64 {
	rv := objc.SendIfResponds[int64](i.ID, objc.Sel("placementSparsePageSize"))
	return rv
}
func (i IOGPUMetalTexture) RemoteStorageTexture() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("remoteStorageTexture"))
	return rv
}
func (i IOGPUMetalTexture) RootResource() IIOGPUMetalResource {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("rootResource"))
	return IOGPUMetalResourceFromID(objc.ID(rv))
}
func (i IOGPUMetalTexture) RootResourceIsSuballocatedBuffer() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("rootResourceIsSuballocatedBuffer"))
	return rv
}
func (i IOGPUMetalTexture) Rotation() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("rotation"))
	return rv
}
func (i IOGPUMetalTexture) SampleCount() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("sampleCount"))
	return rv
}
func (i IOGPUMetalTexture) Shareable() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("shareable"))
	return rv
}
func (i IOGPUMetalTexture) Swizzle() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("swizzle"))
	return rv
}
func (i IOGPUMetalTexture) SwizzleKey() uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("swizzleKey"))
	return rv
}
func (i IOGPUMetalTexture) TextureType() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("textureType"))
	return rv
}
func (i IOGPUMetalTexture) UniqueIdentifier() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("uniqueIdentifier"))
	return rv
}
func (i IOGPUMetalTexture) Usage() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("usage"))
	return rv
}
func (i IOGPUMetalTexture) Width() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("width"))
	return rv
}

// InitWithDeviceDescriptorSysMemPointerSysMemSizeSysMemLengthSysMemRowBytesArgsArgsSizeDeallocatorSync is a synchronous wrapper around [IOGPUMetalTexture.InitWithDeviceDescriptorSysMemPointerSysMemSizeSysMemLengthSysMemRowBytesArgsArgsSizeDeallocator].
// It blocks until the completion handler fires or the context is cancelled.
func (i IOGPUMetalTexture) InitWithDeviceDescriptorSysMemPointerSysMemSizeSysMemLengthSysMemRowBytesArgsArgsSizeDeallocatorSync(ctx context.Context, device objectivec.IObject, descriptor objectivec.IObject, pointer unsafe.Pointer, size uint64, length uint64, bytes uint64, args *IOGPUNewResourceArgs, size2 uint32) error {
	done := make(chan struct{}, 1)
	i.InitWithDeviceDescriptorSysMemPointerSysMemSizeSysMemLengthSysMemRowBytesArgsArgsSizeDeallocator(device, descriptor, pointer, size, length, bytes, args, size2, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
