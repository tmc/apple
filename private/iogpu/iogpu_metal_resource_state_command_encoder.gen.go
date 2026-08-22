// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalResourceStateCommandEncoder] class.
var (
	_IOGPUMetalResourceStateCommandEncoderClass     IOGPUMetalResourceStateCommandEncoderClass
	_IOGPUMetalResourceStateCommandEncoderClassOnce sync.Once
)

func getIOGPUMetalResourceStateCommandEncoderClass() IOGPUMetalResourceStateCommandEncoderClass {
	_IOGPUMetalResourceStateCommandEncoderClassOnce.Do(func() {
		_IOGPUMetalResourceStateCommandEncoderClass = IOGPUMetalResourceStateCommandEncoderClass{class: objc.GetClass("IOGPUMetalResourceStateCommandEncoder")}
	})
	return _IOGPUMetalResourceStateCommandEncoderClass
}

// GetIOGPUMetalResourceStateCommandEncoderClass returns the class object for IOGPUMetalResourceStateCommandEncoder.
func GetIOGPUMetalResourceStateCommandEncoderClass() IOGPUMetalResourceStateCommandEncoderClass {
	return getIOGPUMetalResourceStateCommandEncoderClass()
}

type IOGPUMetalResourceStateCommandEncoderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalResourceStateCommandEncoderClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalResourceStateCommandEncoderClass) Alloc() IOGPUMetalResourceStateCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalResourceStateCommandEncoder](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalResourceStateCommandEncoder.GetType]
//   - [IOGPUMetalResourceStateCommandEncoder.MoveTextureMappingsFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin]
//   - [IOGPUMetalResourceStateCommandEncoder.UpdateTextureMappingModeIndirectBufferIndirectBufferOffset]
//   - [IOGPUMetalResourceStateCommandEncoder.UpdateTextureMappingModeRegionMipLevelSlice]
//   - [IOGPUMetalResourceStateCommandEncoder.UpdateTextureMappingsModeRegionsMipLevelsSlicesNumRegions]
type IOGPUMetalResourceStateCommandEncoder struct {
	IOGPUMetalCommandEncoder
}

// IOGPUMetalResourceStateCommandEncoderFromID constructs a [IOGPUMetalResourceStateCommandEncoder] from an objc.ID.
func IOGPUMetalResourceStateCommandEncoderFromID(id objc.ID) IOGPUMetalResourceStateCommandEncoder {
	return IOGPUMetalResourceStateCommandEncoder{IOGPUMetalCommandEncoder: IOGPUMetalCommandEncoderFromID(id)}
}

// Ensure IOGPUMetalResourceStateCommandEncoder implements IIOGPUMetalResourceStateCommandEncoder.
var _ IIOGPUMetalResourceStateCommandEncoder = IOGPUMetalResourceStateCommandEncoder{}

// An interface definition for the [IOGPUMetalResourceStateCommandEncoder] class.
//
// # Methods
//
//   - [IIOGPUMetalResourceStateCommandEncoder.GetType]
//   - [IIOGPUMetalResourceStateCommandEncoder.MoveTextureMappingsFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin]
//   - [IIOGPUMetalResourceStateCommandEncoder.UpdateTextureMappingModeIndirectBufferIndirectBufferOffset]
//   - [IIOGPUMetalResourceStateCommandEncoder.UpdateTextureMappingModeRegionMipLevelSlice]
//   - [IIOGPUMetalResourceStateCommandEncoder.UpdateTextureMappingsModeRegionsMipLevelsSlicesNumRegions]
type IIOGPUMetalResourceStateCommandEncoder interface {
	IIOGPUMetalCommandEncoder

	// Topic: Methods

	GetType() uint64
	MoveTextureMappingsFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin(texture objectivec.IObject, slice uint64, level uint64, origin unsafe.Pointer, size unsafe.Pointer, texture2 objectivec.IObject, slice2 uint64, level2 uint64, origin2 unsafe.Pointer)
	UpdateTextureMappingModeIndirectBufferIndirectBufferOffset(mapping objectivec.IObject, mode uint64, buffer objectivec.IObject, offset uint64)
	UpdateTextureMappingModeRegionMipLevelSlice(mapping objectivec.IObject, mode uint64, region unsafe.Pointer, level uint64, slice uint64)
	UpdateTextureMappingsModeRegionsMipLevelsSlicesNumRegions(mappings objectivec.IObject, mode uint64, regions unsafe.Pointer, levels *uint64, slices *uint64, regions2 uint64)
}

// Init initializes the instance.
func (i IOGPUMetalResourceStateCommandEncoder) Init() IOGPUMetalResourceStateCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalResourceStateCommandEncoder](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalResourceStateCommandEncoder) Autorelease() IOGPUMetalResourceStateCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalResourceStateCommandEncoder](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalResourceStateCommandEncoder creates a new IOGPUMetalResourceStateCommandEncoder instance.
func NewIOGPUMetalResourceStateCommandEncoder() IOGPUMetalResourceStateCommandEncoder {
	class := getIOGPUMetalResourceStateCommandEncoderClass()
	rv := objc.SendIfResponds[IOGPUMetalResourceStateCommandEncoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalResourceStateCommandEncoderWithCommandBuffer(buffer objectivec.IObject) IOGPUMetalResourceStateCommandEncoder {
	instance := getIOGPUMetalResourceStateCommandEncoderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCommandBuffer:"), buffer)
	return IOGPUMetalResourceStateCommandEncoderFromID(rv)
}

func (i IOGPUMetalResourceStateCommandEncoder) GetType() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("getType"))
	return rv
}
func (i IOGPUMetalResourceStateCommandEncoder) MoveTextureMappingsFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin(texture objectivec.IObject, slice uint64, level uint64, origin unsafe.Pointer, size unsafe.Pointer, texture2 objectivec.IObject, slice2 uint64, level2 uint64, origin2 unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("moveTextureMappingsFromTexture:sourceSlice:sourceLevel:sourceOrigin:sourceSize:toTexture:destinationSlice:destinationLevel:destinationOrigin:"), texture, slice, level, origin, size, texture2, slice2, level2, origin2)
}
func (i IOGPUMetalResourceStateCommandEncoder) UpdateTextureMappingModeIndirectBufferIndirectBufferOffset(mapping objectivec.IObject, mode uint64, buffer objectivec.IObject, offset uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("updateTextureMapping:mode:indirectBuffer:indirectBufferOffset:"), mapping, mode, buffer, offset)
}
func (i IOGPUMetalResourceStateCommandEncoder) UpdateTextureMappingModeRegionMipLevelSlice(mapping objectivec.IObject, mode uint64, region unsafe.Pointer, level uint64, slice uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("updateTextureMapping:mode:region:mipLevel:slice:"), mapping, mode, region, level, slice)
}
func (i IOGPUMetalResourceStateCommandEncoder) UpdateTextureMappingsModeRegionsMipLevelsSlicesNumRegions(mappings objectivec.IObject, mode uint64, regions unsafe.Pointer, levels *uint64, slices *uint64, regions2 uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("updateTextureMappings:mode:regions:mipLevels:slices:numRegions:"), mappings, mode, regions, unsafe.Pointer(levels), unsafe.Pointer(slices), regions2)
}
