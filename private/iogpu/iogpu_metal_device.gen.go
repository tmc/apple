// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalDevice] class.
var (
	_IOGPUMetalDeviceClass     IOGPUMetalDeviceClass
	_IOGPUMetalDeviceClassOnce sync.Once
)

func getIOGPUMetalDeviceClass() IOGPUMetalDeviceClass {
	_IOGPUMetalDeviceClassOnce.Do(func() {
		_IOGPUMetalDeviceClass = IOGPUMetalDeviceClass{class: objc.GetClass("IOGPUMetalDevice")}
	})
	return _IOGPUMetalDeviceClass
}

// GetIOGPUMetalDeviceClass returns the class object for IOGPUMetalDevice.
func GetIOGPUMetalDeviceClass() IOGPUMetalDeviceClass {
	return getIOGPUMetalDeviceClass()
}

type IOGPUMetalDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalDeviceClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalDeviceClass) Alloc() IOGPUMetalDevice {
	rv := objc.SendIfResponds[IOGPUMetalDevice](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalDevice._addResource]
//   - [IOGPUMetalDevice._deviceWrapper]
//   - [IOGPUMetalDevice._purgeDevice]
//   - [IOGPUMetalDevice._removeResource]
//   - [IOGPUMetalDevice._setAcceleratorService]
//   - [IOGPUMetalDevice._setDeviceWrapper]
//   - [IOGPUMetalDevice.AcceleratorPort]
//   - [IOGPUMetalDevice.AkPrivateResourceListPool]
//   - [IOGPUMetalDevice.AkResourceListPool]
//   - [IOGPUMetalDevice.AllocBufferSubDataWithLengthOptionsAlignmentHeapIndexBufferIndexBufferOffset]
//   - [IOGPUMetalDevice.AllocBufferSubDataWithLengthOptionsAlignmentHeapIndexBufferIndexBufferOffsetParentAddressParentLength]
//   - [IOGPUMetalDevice.BufferClass]
//   - [IOGPUMetalDevice.CmdBufArgsSize]
//   - [IOGPUMetalDevice.DeallocBufferSubDataHeapIndexBufferIndexBufferOffsetLength]
//   - [IOGPUMetalDevice.DedicatedMemorySize]
//   - [IOGPUMetalDevice.DeviceRef]
//   - [IOGPUMetalDevice.GetBuiltInGPUPropertiesTransferRate]
//   - [IOGPUMetalDevice.GetExternalGPUPropertiesTransferRate]
//   - [IOGPUMetalDevice.GetSlottedGPUPropertiesTransferRate]
//   - [IOGPUMetalDevice.GetSurfaceModeWidthHeightForId]
//   - [IOGPUMetalDevice.HwResourcePoolCount]
//   - [IOGPUMetalDevice.HwResourcePools]
//   - [IOGPUMetalDevice.IndirectArgumentBufferDecodingData]
//   - [IOGPUMetalDevice.InitialDebugBufferShmemSize]
//   - [IOGPUMetalDevice.InitialIOKernelCommandListShmemSize]
//   - [IOGPUMetalDevice.InitialKernelCommandShmemSize]
//   - [IOGPUMetalDevice.InitialSegmentListShmemSize]
//   - [IOGPUMetalDevice.InitialSidebandShmemSize]
//   - [IOGPUMetalDevice.IsBuiltIn]
//   - [IOGPUMetalDevice.IsSlotted]
//   - [IOGPUMetalDevice.LaunchMappingThread]
//   - [IOGPUMetalDevice.MemoryInfo]
//   - [IOGPUMetalDevice.NewAccelerationStructureWithBufferOffset]
//   - [IOGPUMetalDevice.NewAccelerationStructureWithBufferOffsetResourceIndex]
//   - [IOGPUMetalDevice.NewAccelerationStructureWithSizeResourceIndex]
//   - [IOGPUMetalDevice.NewArgumentEncoderWithLayout]
//   - [IOGPUMetalDevice.NewDevicePoolAliasedCommandAllocator]
//   - [IOGPUMetalDevice.NewEventWithOptions]
//   - [IOGPUMetalDevice.NewGLDrawable]
//   - [IOGPUMetalDevice.NewIOHandleWithURLCompressionTypeError]
//   - [IOGPUMetalDevice.NewIndirectArgumentBufferLayoutWithStructType]
//   - [IOGPUMetalDevice.NewIndirectArgumentEncoderWithLayout]
//   - [IOGPUMetalDevice.NewIntersectionFunctionTableWithDescriptor]
//   - [IOGPUMetalDevice.NewLateEvalEvent]
//   - [IOGPUMetalDevice.NewTiledTextureWithBytesNoCopyLengthDeallocatorDescriptorOffsetBytesPerRow]
//   - [IOGPUMetalDevice.NewTiledTextureWithBytesNoCopyLengthDescriptorOffsetBytesPerRow]
//   - [IOGPUMetalDevice.NewUncachedIOFileHandleWithURLCompressionMethodError]
//   - [IOGPUMetalDevice.NewUncachedIOFileHandleWithURLError]
//   - [IOGPUMetalDevice.NewUncachedIOHandleWithURLCompressionTypeError]
//   - [IOGPUMetalDevice.NewUncachedIOHandleWithURLError]
//   - [IOGPUMetalDevice.NumCommandBuffers]
//   - [IOGPUMetalDevice.ReleasePeerConnection]
//   - [IOGPUMetalDevice.RetainPeerConnection]
//   - [IOGPUMetalDevice.SetComputePipelineStateCommandShmemSize]
//   - [IOGPUMetalDevice.SetHwResourcePoolCount]
//   - [IOGPUMetalDevice.SetIndirectArgumentBufferDecodingData]
//   - [IOGPUMetalDevice.SetSegmentListShmemSize]
//   - [IOGPUMetalDevice.SharedMemorySize]
//   - [IOGPUMetalDevice.SupportPriorityBand]
//   - [IOGPUMetalDevice.SupportsBackgroundAppRole]
//   - [IOGPUMetalDevice.SupportsResourceDetachBacking]
//   - [IOGPUMetalDevice.SupportsVertexAmplification]
//   - [IOGPUMetalDevice.UpdateGPUSelectionProperties]
//   - [IOGPUMetalDevice.InitWithAcceleratorPort]
//   - [IOGPUMetalDevice.InitWithAcceleratorPortOptions]
//   - [IOGPUMetalDevice.Headless]
//   - [IOGPUMetalDevice.LowPower]
//   - [IOGPUMetalDevice.Removable]
type IOGPUMetalDevice struct {
	metal.MTLDeviceObject
}

// IOGPUMetalDeviceFromID constructs a [IOGPUMetalDevice] from an objc.ID.
func IOGPUMetalDeviceFromID(id objc.ID) IOGPUMetalDevice {
	return IOGPUMetalDevice{MTLDeviceObject: metal.MTLDeviceObjectFromID(id)}
}

// Ensure IOGPUMetalDevice implements IIOGPUMetalDevice.
var _ IIOGPUMetalDevice = IOGPUMetalDevice{}

// An interface definition for the [IOGPUMetalDevice] class.
//
// # Methods
//
//   - [IIOGPUMetalDevice._addResource]
//   - [IIOGPUMetalDevice._deviceWrapper]
//   - [IIOGPUMetalDevice._purgeDevice]
//   - [IIOGPUMetalDevice._removeResource]
//   - [IIOGPUMetalDevice._setAcceleratorService]
//   - [IIOGPUMetalDevice._setDeviceWrapper]
//   - [IIOGPUMetalDevice.AcceleratorPort]
//   - [IIOGPUMetalDevice.AkPrivateResourceListPool]
//   - [IIOGPUMetalDevice.AkResourceListPool]
//   - [IIOGPUMetalDevice.AllocBufferSubDataWithLengthOptionsAlignmentHeapIndexBufferIndexBufferOffset]
//   - [IIOGPUMetalDevice.AllocBufferSubDataWithLengthOptionsAlignmentHeapIndexBufferIndexBufferOffsetParentAddressParentLength]
//   - [IIOGPUMetalDevice.BufferClass]
//   - [IIOGPUMetalDevice.CmdBufArgsSize]
//   - [IIOGPUMetalDevice.DeallocBufferSubDataHeapIndexBufferIndexBufferOffsetLength]
//   - [IIOGPUMetalDevice.DedicatedMemorySize]
//   - [IIOGPUMetalDevice.DeviceRef]
//   - [IIOGPUMetalDevice.GetBuiltInGPUPropertiesTransferRate]
//   - [IIOGPUMetalDevice.GetExternalGPUPropertiesTransferRate]
//   - [IIOGPUMetalDevice.GetSlottedGPUPropertiesTransferRate]
//   - [IIOGPUMetalDevice.GetSurfaceModeWidthHeightForId]
//   - [IIOGPUMetalDevice.HwResourcePoolCount]
//   - [IIOGPUMetalDevice.HwResourcePools]
//   - [IIOGPUMetalDevice.IndirectArgumentBufferDecodingData]
//   - [IIOGPUMetalDevice.InitialDebugBufferShmemSize]
//   - [IIOGPUMetalDevice.InitialIOKernelCommandListShmemSize]
//   - [IIOGPUMetalDevice.InitialKernelCommandShmemSize]
//   - [IIOGPUMetalDevice.InitialSegmentListShmemSize]
//   - [IIOGPUMetalDevice.InitialSidebandShmemSize]
//   - [IIOGPUMetalDevice.IsBuiltIn]
//   - [IIOGPUMetalDevice.IsSlotted]
//   - [IIOGPUMetalDevice.LaunchMappingThread]
//   - [IIOGPUMetalDevice.MemoryInfo]
//   - [IIOGPUMetalDevice.NewAccelerationStructureWithBufferOffset]
//   - [IIOGPUMetalDevice.NewAccelerationStructureWithBufferOffsetResourceIndex]
//   - [IIOGPUMetalDevice.NewAccelerationStructureWithSizeResourceIndex]
//   - [IIOGPUMetalDevice.NewArgumentEncoderWithLayout]
//   - [IIOGPUMetalDevice.NewDevicePoolAliasedCommandAllocator]
//   - [IIOGPUMetalDevice.NewEventWithOptions]
//   - [IIOGPUMetalDevice.NewGLDrawable]
//   - [IIOGPUMetalDevice.NewIOHandleWithURLCompressionTypeError]
//   - [IIOGPUMetalDevice.NewIndirectArgumentBufferLayoutWithStructType]
//   - [IIOGPUMetalDevice.NewIndirectArgumentEncoderWithLayout]
//   - [IIOGPUMetalDevice.NewIntersectionFunctionTableWithDescriptor]
//   - [IIOGPUMetalDevice.NewLateEvalEvent]
//   - [IIOGPUMetalDevice.NewTiledTextureWithBytesNoCopyLengthDeallocatorDescriptorOffsetBytesPerRow]
//   - [IIOGPUMetalDevice.NewTiledTextureWithBytesNoCopyLengthDescriptorOffsetBytesPerRow]
//   - [IIOGPUMetalDevice.NewUncachedIOFileHandleWithURLCompressionMethodError]
//   - [IIOGPUMetalDevice.NewUncachedIOFileHandleWithURLError]
//   - [IIOGPUMetalDevice.NewUncachedIOHandleWithURLCompressionTypeError]
//   - [IIOGPUMetalDevice.NewUncachedIOHandleWithURLError]
//   - [IIOGPUMetalDevice.NumCommandBuffers]
//   - [IIOGPUMetalDevice.ReleasePeerConnection]
//   - [IIOGPUMetalDevice.RetainPeerConnection]
//   - [IIOGPUMetalDevice.SetComputePipelineStateCommandShmemSize]
//   - [IIOGPUMetalDevice.SetHwResourcePoolCount]
//   - [IIOGPUMetalDevice.SetIndirectArgumentBufferDecodingData]
//   - [IIOGPUMetalDevice.SetSegmentListShmemSize]
//   - [IIOGPUMetalDevice.SharedMemorySize]
//   - [IIOGPUMetalDevice.SupportPriorityBand]
//   - [IIOGPUMetalDevice.SupportsBackgroundAppRole]
//   - [IIOGPUMetalDevice.SupportsResourceDetachBacking]
//   - [IIOGPUMetalDevice.SupportsVertexAmplification]
//   - [IIOGPUMetalDevice.UpdateGPUSelectionProperties]
//   - [IIOGPUMetalDevice.InitWithAcceleratorPort]
//   - [IIOGPUMetalDevice.InitWithAcceleratorPortOptions]
//   - [IIOGPUMetalDevice.Headless]
//   - [IIOGPUMetalDevice.LowPower]
//   - [IIOGPUMetalDevice.Removable]
type IIOGPUMetalDevice interface {
	metal.MTLDevice

	// Topic: Methods

	_addResource(resource objectivec.IObject)
	_deviceWrapper() objectivec.IObject
	_purgeDevice()
	_removeResource(resource objectivec.IObject)
	_setAcceleratorService(service objectivec.IObject)
	_setDeviceWrapper(wrapper objectivec.IObject)
	AcceleratorPort() uint32
	AkPrivateResourceListPool() objectivec.IObject
	AkResourceListPool() objectivec.IObject
	AllocBufferSubDataWithLengthOptionsAlignmentHeapIndexBufferIndexBufferOffset(length uint64, options uint64, alignment uint64, index *int16, index2 *int16, offset *uint64) objectivec.IObject
	AllocBufferSubDataWithLengthOptionsAlignmentHeapIndexBufferIndexBufferOffsetParentAddressParentLength(length uint64, options uint64, alignment uint64, index *int16, index2 *int16, offset *uint64, address uint64, length2 uint64) objectivec.IObject
	BufferClass() objectivec.Class
	CmdBufArgsSize() uint32
	DeallocBufferSubDataHeapIndexBufferIndexBufferOffsetLength(data objectivec.IObject, index int16, index2 int16, offset uint64, length uint64)
	DedicatedMemorySize() uint64
	DeviceRef() uintptr
	GetBuiltInGPUPropertiesTransferRate(gPUProperties *uint64, rate *uint64)
	GetExternalGPUPropertiesTransferRate(gPUProperties *uint64, rate *uint64)
	GetSlottedGPUPropertiesTransferRate(gPUProperties *uint64, rate *uint64)
	GetSurfaceModeWidthHeightForId(mode *uint64, width *uint32, height *uint32, id uint32) int
	HwResourcePoolCount() uint32
	HwResourcePools() []objectivec.IObject
	IndirectArgumentBufferDecodingData() objectivec.IObject
	InitialDebugBufferShmemSize() IOGPUMetalDevice
	InitialIOKernelCommandListShmemSize() IOGPUMetalDevice
	InitialKernelCommandShmemSize() IOGPUMetalDevice
	InitialSegmentListShmemSize() IOGPUMetalDevice
	InitialSidebandShmemSize() IOGPUMetalDevice
	IsBuiltIn() bool
	IsSlotted() bool
	LaunchMappingThread()
	MemoryInfo() IIOGPUMemoryInfo
	NewAccelerationStructureWithBufferOffset(buffer objectivec.IObject, offset uint64) objectivec.IObject
	NewAccelerationStructureWithBufferOffsetResourceIndex(buffer objectivec.IObject, offset uint64, index uint64) objectivec.IObject
	NewAccelerationStructureWithSizeResourceIndex(size uint64, index uint64) objectivec.IObject
	NewArgumentEncoderWithLayout(layout objectivec.IObject) objectivec.IObject
	NewDevicePoolAliasedCommandAllocator() objectivec.IObject
	NewEventWithOptions(options int64) objectivec.IObject
	NewGLDrawable() objectivec.IObject
	NewIOHandleWithURLCompressionTypeError(url foundation.NSURL, type_ int64) (objectivec.IObject, error)
	NewIndirectArgumentBufferLayoutWithStructType(type_ objectivec.IObject) objectivec.IObject
	NewIndirectArgumentEncoderWithLayout(layout objectivec.IObject) objectivec.IObject
	NewIntersectionFunctionTableWithDescriptor(descriptor objectivec.IObject) objectivec.IObject
	NewLateEvalEvent() objectivec.IObject
	NewTiledTextureWithBytesNoCopyLengthDeallocatorDescriptorOffsetBytesPerRow(copy_ unsafe.Pointer, length uint64, deallocator VoidHandler, descriptor objectivec.IObject, offset uint64, row uint64) objectivec.IObject
	NewTiledTextureWithBytesNoCopyLengthDescriptorOffsetBytesPerRow(copy_ unsafe.Pointer, length uint64, descriptor objectivec.IObject, offset uint64, row uint64) objectivec.IObject
	NewUncachedIOFileHandleWithURLCompressionMethodError(url foundation.NSURL, method int64) (objectivec.IObject, error)
	NewUncachedIOFileHandleWithURLError(url foundation.NSURL) (objectivec.IObject, error)
	NewUncachedIOHandleWithURLCompressionTypeError(url foundation.NSURL, type_ int64) (objectivec.IObject, error)
	NewUncachedIOHandleWithURLError(url foundation.NSURL) (objectivec.IObject, error)
	NumCommandBuffers() int
	ReleasePeerConnection(connection objectivec.IObject)
	RetainPeerConnection(connection objectivec.IObject) bool
	SetComputePipelineStateCommandShmemSize(size uint32)
	SetHwResourcePoolCount(pool []objectivec.IObject, count int)
	SetIndirectArgumentBufferDecodingData(data objectivec.IObject)
	SetSegmentListShmemSize(size uint32)
	SharedMemorySize() uint64
	SupportPriorityBand() bool
	SupportsBackgroundAppRole() bool
	SupportsResourceDetachBacking() bool
	SupportsVertexAmplification() bool
	UpdateGPUSelectionProperties()
	InitWithAcceleratorPort(port uint32) IOGPUMetalDevice
	InitWithAcceleratorPortOptions(port uint32, options uint64) IOGPUMetalDevice
	Headless() bool
	LowPower() bool
	Removable() bool
}

// Init initializes the instance.
func (i IOGPUMetalDevice) Init() IOGPUMetalDevice {
	rv := objc.SendIfResponds[IOGPUMetalDevice](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalDevice) Autorelease() IOGPUMetalDevice {
	rv := objc.SendIfResponds[IOGPUMetalDevice](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalDevice creates a new IOGPUMetalDevice instance.
func NewIOGPUMetalDevice() IOGPUMetalDevice {
	class := getIOGPUMetalDeviceClass()
	rv := objc.SendIfResponds[IOGPUMetalDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalDeviceWithAcceleratorPort(port uint32) IOGPUMetalDevice {
	instance := getIOGPUMetalDeviceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithAcceleratorPort:"), port)
	return IOGPUMetalDeviceFromID(rv)
}

func NewGPUMetalDeviceWithAcceleratorPortOptions(port uint32, options uint64) IOGPUMetalDevice {
	instance := getIOGPUMetalDeviceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithAcceleratorPort:options:"), port, options)
	return IOGPUMetalDeviceFromID(rv)
}

func (i IOGPUMetalDevice) _addResource(resource objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("_addResource:"), resource)
}

// AddResource is an exported wrapper for the private method _addResource.
func (i IOGPUMetalDevice) AddResource(resource objectivec.IObject) error {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_addResource:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_addResource:"}
		return err
	}
	i._addResource(resource)
	return nil
}

// CanAddResource reports whether the receiver responds to the private selector _addResource:.
func (i IOGPUMetalDevice) CanAddResource() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_addResource:"))
}
func (i IOGPUMetalDevice) _deviceWrapper() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("_deviceWrapper"))
	return objectivec.Object{ID: rv}
}

// DeviceWrapper is an exported wrapper for the private method _deviceWrapper.
func (i IOGPUMetalDevice) DeviceWrapper() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_deviceWrapper")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_deviceWrapper"}
		return nil, err
	}
	return i._deviceWrapper(), nil
}

// CanDeviceWrapper reports whether the receiver responds to the private selector _deviceWrapper.
func (i IOGPUMetalDevice) CanDeviceWrapper() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_deviceWrapper"))
}
func (i IOGPUMetalDevice) _purgeDevice() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("_purgeDevice"))
}

// PurgeDevice is an exported wrapper for the private method _purgeDevice.
func (i IOGPUMetalDevice) PurgeDevice() error {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_purgeDevice")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_purgeDevice"}
		return err
	}
	i._purgeDevice()
	return nil
}

// CanPurgeDevice reports whether the receiver responds to the private selector _purgeDevice.
func (i IOGPUMetalDevice) CanPurgeDevice() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_purgeDevice"))
}
func (i IOGPUMetalDevice) _removeResource(resource objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("_removeResource:"), resource)
}

// RemoveResource is an exported wrapper for the private method _removeResource.
func (i IOGPUMetalDevice) RemoveResource(resource objectivec.IObject) error {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_removeResource:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_removeResource:"}
		return err
	}
	i._removeResource(resource)
	return nil
}

// CanRemoveResource reports whether the receiver responds to the private selector _removeResource:.
func (i IOGPUMetalDevice) CanRemoveResource() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_removeResource:"))
}
func (i IOGPUMetalDevice) _setAcceleratorService(service objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("_setAcceleratorService:"), service)
}

// SetAcceleratorService is an exported wrapper for the private method _setAcceleratorService.
func (i IOGPUMetalDevice) SetAcceleratorService(service objectivec.IObject) error {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_setAcceleratorService:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setAcceleratorService:"}
		return err
	}
	i._setAcceleratorService(service)
	return nil
}

// CanSetAcceleratorService reports whether the receiver responds to the private selector _setAcceleratorService:.
func (i IOGPUMetalDevice) CanSetAcceleratorService() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_setAcceleratorService:"))
}
func (i IOGPUMetalDevice) _setDeviceWrapper(wrapper objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("_setDeviceWrapper:"), wrapper)
}

// SetDeviceWrapper is an exported wrapper for the private method _setDeviceWrapper.
func (i IOGPUMetalDevice) SetDeviceWrapper(wrapper objectivec.IObject) error {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_setDeviceWrapper:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setDeviceWrapper:"}
		return err
	}
	i._setDeviceWrapper(wrapper)
	return nil
}

// CanSetDeviceWrapper reports whether the receiver responds to the private selector _setDeviceWrapper:.
func (i IOGPUMetalDevice) CanSetDeviceWrapper() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_setDeviceWrapper:"))
}
func (i IOGPUMetalDevice) AkPrivateResourceListPool() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("akPrivateResourceListPool"))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalDevice) AkResourceListPool() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("akResourceListPool"))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalDevice) AllocBufferSubDataWithLengthOptionsAlignmentHeapIndexBufferIndexBufferOffset(length uint64, options uint64, alignment uint64, index *int16, index2 *int16, offset *uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("allocBufferSubDataWithLength:options:alignment:heapIndex:bufferIndex:bufferOffset:"), length, options, alignment, index, index2, unsafe.Pointer(offset))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalDevice) AllocBufferSubDataWithLengthOptionsAlignmentHeapIndexBufferIndexBufferOffsetParentAddressParentLength(length uint64, options uint64, alignment uint64, index *int16, index2 *int16, offset *uint64, address uint64, length2 uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("allocBufferSubDataWithLength:options:alignment:heapIndex:bufferIndex:bufferOffset:parentAddress:parentLength:"), length, options, alignment, index, index2, unsafe.Pointer(offset), address, length2)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalDevice) CmdBufArgsSize() uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("cmdBufArgsSize"))
	return rv
}
func (i IOGPUMetalDevice) DeallocBufferSubDataHeapIndexBufferIndexBufferOffsetLength(data objectivec.IObject, index int16, index2 int16, offset uint64, length uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("deallocBufferSubData:heapIndex:bufferIndex:bufferOffset:length:"), data, index, index2, offset, length)
}
func (i IOGPUMetalDevice) DeviceRef() uintptr {
	rv := objc.SendIfResponds[uintptr](i.ID, objc.Sel("deviceRef"))
	return rv
}
func (i IOGPUMetalDevice) GetBuiltInGPUPropertiesTransferRate(gPUProperties *uint64, rate *uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("getBuiltInGPUProperties:transferRate:"), unsafe.Pointer(gPUProperties), unsafe.Pointer(rate))
}
func (i IOGPUMetalDevice) GetExternalGPUPropertiesTransferRate(gPUProperties *uint64, rate *uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("getExternalGPUProperties:transferRate:"), unsafe.Pointer(gPUProperties), unsafe.Pointer(rate))
}
func (i IOGPUMetalDevice) GetSlottedGPUPropertiesTransferRate(gPUProperties *uint64, rate *uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("getSlottedGPUProperties:transferRate:"), unsafe.Pointer(gPUProperties), unsafe.Pointer(rate))
}
func (i IOGPUMetalDevice) GetSurfaceModeWidthHeightForId(mode *uint64, width *uint32, height *uint32, id uint32) int {
	rv := objc.SendIfResponds[int](i.ID, objc.Sel("getSurfaceMode:width:height:forId:"), unsafe.Pointer(mode), unsafe.Pointer(width), unsafe.Pointer(height), id)
	return rv
}
func (i IOGPUMetalDevice) IndirectArgumentBufferDecodingData() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("indirectArgumentBufferDecodingData"))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalDevice) InitialDebugBufferShmemSize() IOGPUMetalDevice {
	rv := objc.SendIfResponds[IOGPUMetalDevice](i.ID, objc.Sel("initialDebugBufferShmemSize"))
	return rv
}
func (i IOGPUMetalDevice) InitialIOKernelCommandListShmemSize() IOGPUMetalDevice {
	rv := objc.SendIfResponds[IOGPUMetalDevice](i.ID, objc.Sel("initialIOKernelCommandListShmemSize"))
	return rv
}
func (i IOGPUMetalDevice) InitialKernelCommandShmemSize() IOGPUMetalDevice {
	rv := objc.SendIfResponds[IOGPUMetalDevice](i.ID, objc.Sel("initialKernelCommandShmemSize"))
	return rv
}
func (i IOGPUMetalDevice) InitialSegmentListShmemSize() IOGPUMetalDevice {
	rv := objc.SendIfResponds[IOGPUMetalDevice](i.ID, objc.Sel("initialSegmentListShmemSize"))
	return rv
}
func (i IOGPUMetalDevice) InitialSidebandShmemSize() IOGPUMetalDevice {
	rv := objc.SendIfResponds[IOGPUMetalDevice](i.ID, objc.Sel("initialSidebandShmemSize"))
	return rv
}
func (i IOGPUMetalDevice) IsBuiltIn() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("isBuiltIn"))
	return rv
}
func (i IOGPUMetalDevice) IsSlotted() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("isSlotted"))
	return rv
}
func (i IOGPUMetalDevice) LaunchMappingThread() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("launchMappingThread"))
}
func (i IOGPUMetalDevice) NewAccelerationStructureWithBufferOffset(buffer objectivec.IObject, offset uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newAccelerationStructureWithBuffer:offset:"), buffer, offset)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalDevice) NewAccelerationStructureWithBufferOffsetResourceIndex(buffer objectivec.IObject, offset uint64, index uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newAccelerationStructureWithBuffer:offset:resourceIndex:"), buffer, offset, index)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalDevice) NewAccelerationStructureWithSizeResourceIndex(size uint64, index uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newAccelerationStructureWithSize:resourceIndex:"), size, index)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalDevice) NewArgumentEncoderWithLayout(layout objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newArgumentEncoderWithLayout:"), layout)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalDevice) NewDevicePoolAliasedCommandAllocator() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newDevicePoolAliasedCommandAllocator"))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalDevice) NewEventWithOptions(options int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newEventWithOptions:"), options)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalDevice) NewGLDrawable() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newGLDrawable"))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalDevice) NewIOHandleWithURLCompressionTypeError(url foundation.NSURL, type_ int64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](i.ID, objc.Sel("newIOHandleWithURL:compressionType:error:"), url, type_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (i IOGPUMetalDevice) NewIndirectArgumentBufferLayoutWithStructType(type_ objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newIndirectArgumentBufferLayoutWithStructType:"), type_)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalDevice) NewIndirectArgumentEncoderWithLayout(layout objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newIndirectArgumentEncoderWithLayout:"), layout)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalDevice) NewIntersectionFunctionTableWithDescriptor(descriptor objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newIntersectionFunctionTableWithDescriptor:"), descriptor)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalDevice) NewLateEvalEvent() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newLateEvalEvent"))
	return objectivec.Object{ID: rv}
}

var _iogpumetaldevice_newtiledtexturewithbytesnocopy_length_deallocator_descriptor_offset_bytesperrow_p2_key byte

func (i IOGPUMetalDevice) NewTiledTextureWithBytesNoCopyLengthDeallocatorDescriptorOffsetBytesPerRow(copy_ unsafe.Pointer, length uint64, deallocator VoidHandler, descriptor objectivec.IObject, offset uint64, row uint64) objectivec.IObject {
	_block2, _ := NewVoidBlock(deallocator)
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newTiledTextureWithBytesNoCopy:length:deallocator:descriptor:offset:bytesPerRow:"), copy_, length, _block2, descriptor, offset, row)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalDevice) NewTiledTextureWithBytesNoCopyLengthDescriptorOffsetBytesPerRow(copy_ unsafe.Pointer, length uint64, descriptor objectivec.IObject, offset uint64, row uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newTiledTextureWithBytesNoCopy:length:descriptor:offset:bytesPerRow:"), copy_, length, descriptor, offset, row)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalDevice) NewUncachedIOFileHandleWithURLCompressionMethodError(url foundation.NSURL, method int64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](i.ID, objc.Sel("newUncachedIOFileHandleWithURL:compressionMethod:error:"), url, method, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (i IOGPUMetalDevice) NewUncachedIOFileHandleWithURLError(url foundation.NSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](i.ID, objc.Sel("newUncachedIOFileHandleWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (i IOGPUMetalDevice) NewUncachedIOHandleWithURLCompressionTypeError(url foundation.NSURL, type_ int64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](i.ID, objc.Sel("newUncachedIOHandleWithURL:compressionType:error:"), url, type_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (i IOGPUMetalDevice) NewUncachedIOHandleWithURLError(url foundation.NSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](i.ID, objc.Sel("newUncachedIOHandleWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (i IOGPUMetalDevice) ReleasePeerConnection(connection objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("releasePeerConnection:"), connection)
}
func (i IOGPUMetalDevice) RetainPeerConnection(connection objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("retainPeerConnection:"), connection)
	return rv
}
func (i IOGPUMetalDevice) SetComputePipelineStateCommandShmemSize(size uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setComputePipelineStateCommandShmemSize:"), size)
}
func (i IOGPUMetalDevice) SetHwResourcePoolCount(pool []objectivec.IObject, count int) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setHwResourcePool:count:"), objc.CArray(pool), count)
}
func (i IOGPUMetalDevice) SetIndirectArgumentBufferDecodingData(data objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setIndirectArgumentBufferDecodingData:"), data)
}
func (i IOGPUMetalDevice) SetSegmentListShmemSize(size uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setSegmentListShmemSize:"), size)
}
func (i IOGPUMetalDevice) SupportsBackgroundAppRole() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("supportsBackgroundAppRole"))
	return rv
}
func (i IOGPUMetalDevice) SupportsResourceDetachBacking() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("supportsResourceDetachBacking"))
	return rv
}
func (i IOGPUMetalDevice) UpdateGPUSelectionProperties() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("updateGPUSelectionProperties"))
}
func (i IOGPUMetalDevice) InitWithAcceleratorPort(port uint32) IOGPUMetalDevice {
	rv := objc.SendIfResponds[IOGPUMetalDevice](i.ID, objc.Sel("initWithAcceleratorPort:"), port)
	return rv
}
func (i IOGPUMetalDevice) InitWithAcceleratorPortOptions(port uint32, options uint64) IOGPUMetalDevice {
	rv := objc.SendIfResponds[IOGPUMetalDevice](i.ID, objc.Sel("initWithAcceleratorPort:options:"), port, options)
	return rv
}

func (_IOGPUMetalDeviceClass IOGPUMetalDeviceClass) RegisterAcceleratorService(service objectivec.IObject) {
	objc.SendIfResponds[objc.ID](objc.ID(_IOGPUMetalDeviceClass.class), objc.Sel("registerAcceleratorService:"), service)
}
func (_IOGPUMetalDeviceClass IOGPUMetalDeviceClass) RegisterDevices() {
	objc.SendIfResponds[objc.ID](objc.ID(_IOGPUMetalDeviceClass.class), objc.Sel("registerDevices"))
}

func (i IOGPUMetalDevice) AcceleratorPort() uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("acceleratorPort"))
	return rv
}
func (i IOGPUMetalDevice) BufferClass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](i.ID, objc.Sel("bufferClass"))
	return objectivec.Class(rv)
}
func (i IOGPUMetalDevice) DedicatedMemorySize() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("dedicatedMemorySize"))
	return rv
}
func (i IOGPUMetalDevice) Headless() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("headless"))
	return rv
}
func (i IOGPUMetalDevice) HwResourcePoolCount() uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("hwResourcePoolCount"))
	return rv
}
func (i IOGPUMetalDevice) HwResourcePools() []objectivec.IObject {
	rv := objc.SendIfResponds[[]objc.ID](i.ID, objc.Sel("hwResourcePools"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (i IOGPUMetalDevice) LowPower() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("lowPower"))
	return rv
}
func (i IOGPUMetalDevice) MemoryInfo() IIOGPUMemoryInfo {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("memoryInfo"))
	return IOGPUMemoryInfoFromID(objc.ID(rv))
}
func (i IOGPUMetalDevice) NumCommandBuffers() int {
	rv := objc.SendIfResponds[int](i.ID, objc.Sel("numCommandBuffers"))
	return rv
}
func (i IOGPUMetalDevice) Removable() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("removable"))
	return rv
}
func (i IOGPUMetalDevice) SharedMemorySize() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("sharedMemorySize"))
	return rv
}
func (i IOGPUMetalDevice) SupportPriorityBand() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("supportPriorityBand"))
	return rv
}
func (i IOGPUMetalDevice) SupportsVertexAmplification() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("supportsVertexAmplification"))
	return rv
}
