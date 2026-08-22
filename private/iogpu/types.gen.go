// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// C struct types

// IOGPUAddressRange
type IOGPUAddressRange struct {
	Field1 uint64
	Field2 uint64
}

// IOGPUClientSharedGlobalRO
type IOGPUClientSharedGlobalRO struct {
}

// IOGPUClientSharedRO
type IOGPUClientSharedRO struct {
}

// IOGPUCommandQueueCommandBufferArgs
type IOGPUCommandQueueCommandBufferArgs struct {
	Field1 uint32
	Field2 uint32
	Field3 uint32
	Field4 unsafe.Pointer
	Field5 unsafe.Pointer
	Field6 uint32
	Field7 uint64
}

// IOGPUDeviceCommPage
type IOGPUDeviceCommPage struct {
}

// IOGPUDeviceConfigData
type IOGPUDeviceConfigData struct {
	Field1  uint32
	Field2  uint32
	Field3  uint64
	Field4  uint64
	Field5  uint32
	Field6  uint32
	Field7  uint64
	Field8  uint64
	Field9  uint64
	Field10 uint32
	Field11 uint32
}

// IOGPUDeviceNewCommandQueueArgs
type IOGPUDeviceNewCommandQueueArgs struct {
	Field1 [1024]int8
	Field2 int32
	Field3 uint8
	Field4 uint8
	Field5 [2]uint8
}

// IOGPUDrawableSurfaceConfig
type IOGPUDrawableSurfaceConfig struct {
	Window_mode uint64
	Width       int32
	Height      int32
	Sfc_width   int32
	Sfc_height  int32
	Config_bits uint32
	Status_bits uint32
	Texture_ram uint64
}

// IOGPUIOKernelCommandListHeader
type IOGPUIOKernelCommandListHeader struct {
	Field1 uint32
	Field2 uint32
}

// IOGPUKernelCommandSignalEventAgentArgs
type IOGPUKernelCommandSignalEventAgentArgs struct {
	Field1 uint32
	Field2 uint32
	Field3 uint64
	Field4 uint64
}

// IOGPUKernelCommandSignalOrWaitEventArgs
type IOGPUKernelCommandSignalOrWaitEventArgs struct {
	Field1 uint32
	Field2 uint32
	Field3 uint64
}

// IOGPUMTLIdKey
type IOGPUMTLIdKey struct {
}

// IOGPUMetalCommandBufferResourceInfo
type IOGPUMetalCommandBufferResourceInfo struct {
	Field1 uint64
	Field2 *uint32
	Field3 *uint32
	Field4 *uint32
	Field5 objectivec.Object
	Field6 IOGPUResourceInfo
}

// IOGPUMetalCommandBufferSidebandBuffer
type IOGPUMetalCommandBufferSidebandBuffer struct {
	Field1 objectivec.Object
	Field2 *byte
	Field3 *byte
	Field4 *byte
}

// IOGPUMetalCommandBufferStorage
type IOGPUMetalCommandBufferStorage struct {
	Field1  objectivec.Object
	Field2  IOGPUMetalCommandBufferStoragePoolRef
	Field3  [2]uint64
	Field4  objectivec.Object
	Field5  *byte
	Field6  *byte
	Field7  *byte
	Field8  IOGPUMetalCommandBufferSidebandBuffer
	Field9  objectivec.Object
	Field10 *byte
	Field11 *byte
	Field12 IOGPUSegmentListHeaderRef
	Field13 IOGPUSegmentResourceListHeaderRef
	Field14 IOGPUSegmentResourceDescriptorGroupRef
	Field15 IOGPUResourceList
	Field16 objectivec.Object
	Field17 objectivec.Object
	Field18 uint64
	Field19 uint64
	Field20 objc.ID
	Field21 IOGPUMetalCommandBufferResourceInfoRef
	Field22 uint64
	Field23 objc.ID
	Field24 uint32
	Field25 uint64
	Field26 IOGPUSegmentListShmemHeaderRef
	Field27 IOGPUSegmentKernelCommmandListHeaderRef
	Field28 uint32
	Field29 uint32
	Field30 int32
	Field31 int32
	Field32 uint32
	Field33 IOGPUSegmentResourceListHeaderRef
	Field34 IOGPUSegmentResourceDescriptorGroupRef
	Field35 uint32
	Field36 objectivec.Object
	Field37 *byte
	Field38 *byte
	Field39 SIOGPUKernelCommandSetResourceGroupsArgsRef
	Field40 [2]objectivec.Object
	Field41 bool
}

// IOGPUMetalCommandBufferStoragePool
type IOGPUMetalCommandBufferStoragePool struct {
	Field1 GpuStorageQueue
	Field2 OSUnfairLockS
	Field3 int32
	Field4 int32
	Field5 int32
	Field6 objectivec.Object
	Field7 unsafe.Pointer
}

// IOGPUMetalDeviceShmemPoolPrivate
type IOGPUMetalDeviceShmemPoolPrivate struct {
	Queue      ShmemPoolQueue
	Lock       OSUnfairLockS
	Count      int32
	ShmemClass objc.Class
	Device     *IOGPUMetalDevice
	ShmemSize  uint32
	ShmemType  int32
}

// IOGPUMetalDeviceShmemPrivate
type IOGPUMetalDeviceShmemPrivate struct {
	Pool               *IOGPUMetalDeviceShmemPool
	Entry              [2]uint64
	Time_added         uint64
	Trim_level         int64
	Used_history       [8]int64
	Used_history_index int32
}

// IOGPUMetalResourcePoolPrivate
type IOGPUMetalResourcePoolPrivate struct {
	VolatileQueue    IOGPUMetalResourceQueue
	NonvolatileQueue IOGPUMetalResourceQueue
	Lock             OSUnfairLockS
	Count            int32
}

// IOGPUMetalResourcePrivate
type IOGPUMetalResourcePrivate struct {
	Pool            *IOGPUMetalResourcePool
	Entry           [2]uint64
	Time_added      uint64
	Pool_generation uint32
}

// IOGPUMetalResourceQueue
type IOGPUMetalResourceQueue struct {
	Tqh_first *IOGPUMetalPooledResource
	Tqh_last  objc.ID
}

// IOGPUMetalSuballocator
type IOGPUMetalSuballocator struct {
}

// IOGPUNewResourceArgs
type IOGPUNewResourceArgs struct {
	Field1 IOGPUNewResourceData
}

// IOGPUNewResourceData
type IOGPUNewResourceData struct {
	Field1  uint32
	Field2  uint32
	Field3  uint16
	Field4  uint16
	Field5  uint16
	Field6  uint16
	Field7  uint8
	Field8  uint8
	Field9  uint8
	Field10 uint8
	Field11 uint32
	Field12 uint64
	Field13 uint64
	Field14 uint64
	Field15 uint32
	Field16 uint32
	Field17 unsafe.Pointer
}

// IOGPUResourceInfo
type IOGPUResourceInfo struct {
	Field1 unsafe.Pointer
	Field2 objectivec.Object
	Field3 objectivec.Object
	Field4 uint32
}

// IOGPUResourceList
type IOGPUResourceList struct {
	Field1  [114]uint32
	Field2  *uint32
	Field3  *uint64
	Field4  unsafe.Pointer
	Field5  uint32
	Field6  uint32
	Field7  uint32
	Field8  uint32
	Field9  uint32
	Field10 uint32
	Field11 uint32
	Field12 uint32
	Field13 IOGPUSegmentResourceDescriptorGroupRef
	Field14 uint64
	Field15 uint64
	Field16 uint32
	Field17 uint32
	Field18 uint32
	Field19 uint32
	Field20 uint32
	Field21 uint32
	Field22 uint32
	Field23 unsafe.Pointer
	Field24 unsafe.Pointer
}

// IOGPUSegmentKernelCommmandListHeader
type IOGPUSegmentKernelCommmandListHeader struct {
}

// IOGPUSegmentListHeader
type IOGPUSegmentListHeader struct {
	Field1 uint64
	Field2 uint32
	Field3 uint32
	Field4 unsafe.Pointer
}

// IOGPUSegmentListShmemHeader
type IOGPUSegmentListShmemHeader struct {
}

// IOGPUSegmentResourceDescriptorGroup
type IOGPUSegmentResourceDescriptorGroup struct {
	Field1 [6]uint32
	Field2 [6]uint32
	Field3 [6]uint16
	Field4 uint16
	Field5 uint16
}

// IOGPUSegmentResourceListHeader
type IOGPUSegmentResourceListHeader struct {
	Field1 uint64
	Field2 uint32
	Field3 uint32
	Field4 uint32
	Field5 uint32
	Field6 uint32
	Field7 uint32
	Field8 unsafe.Pointer
}

// MTLIndirectCommandBufferHeader
type MTLIndirectCommandBufferHeader struct {
	HeaderSize                          uint64
	CommandTypes                        uint32
	InheritPipelineState                bool
	InheritBuffers                      bool
	MaxVertexBufferBindCount            uint8
	MaxFragmentBufferBindCount          uint8
	MaxKernelBufferBindCount            uint8
	MaxObjectBufferBindCount            uint8
	MaxMeshBufferBindCount              uint8
	SupportRayTracing                   bool
	SupportDynamicAttributeStride       bool
	MaxKernelThreadgroupMemoryBindCount uint8
	MaxObjectThreadgroupMemoryBindCount uint8
	MaxScissorRectCount                 uint8
	MaxViewportCount                    uint8
	InheritDepthStencilState            bool
	InheritDepthBias                    bool
	InheritStencilReferenceValues       bool
	InheritDepthClipMode                bool
	InheritCullMode                     bool
	InheritFrontFacingWinding           bool
	InheritTriangleFillMode             bool
	InheritDepthTestBounds              bool
	InheritScissorRects                 bool
	InheritViewports                    bool
	InheritBlendColor                   bool
	AllowOverrideRenderStates           int64
	Size                                uint64
}

// MTLRangeAllocator
type MTLRangeAllocator struct {
	Elements             MTLRangeAllocatorElementRef
	NumElements          uint32
	Capacity             uint64
	CapacityIncrement    uint64
	DefaultAlignmentMask uint64
}

// MTLRangeAllocatorElement
type MTLRangeAllocatorElement struct {
}

// MTLTensorSlice
type MTLTensorSlice struct {
	Field1 objectivec.Object
	Field2 objectivec.Object
}

// CFArray
type CFArray struct {
}

// CFDictionary
type CFDictionary struct {
}

// CFRuntimeBase
type CFRuntimeBase struct {
	Field1 uint64
	Field2 uint64
}

// CFString
type CFString struct {
}

// IOGPUCommandQueue
type IOGPUCommandQueue struct {
}

// IOGPUDevice
type IOGPUDevice struct {
	Field1  CFRuntimeBase
	Field2  uint32
	Field3  uint32
	Field4  *uint32
	Field5  *uint32
	Field6  int32
	Field7  uint32
	Field8  uint32
	Field9  uint32
	Field10 uint32
	Field11 OSUnfairLockS
	Field12 uint64
	Field13 uint64
	Field14 uint64
	Field15 uint64
	Field16 IOGPUDeviceCommPageRef
	Field17 IOGPUClientSharedGlobalRORef
	Field18 func()
	Field19 OSUnfairLockS
	Field20 ShmemlogList
	Field21 uint64
	Field22 uint64
	Field23 uint64
	Field24 IOGPUDeviceConfigData
}

// IOGPUGLDrawable
type IOGPUGLDrawable struct {
}

// IOGPUIOCommandQueue
type IOGPUIOCommandQueue struct {
}

// IOGPUResource
type IOGPUResource struct {
	Field1  CFRuntimeBase
	Field2  IOGPUDeviceRef
	Field3  unsafe.Pointer
	Field4  uint64
	Field5  uint64
	Field6  uint32
	Field7  uint32
	Field8  uint64
	Field9  uint64
	Field10 IOGPUClientSharedRORef
	Field11 uint64
	Field12 uint64
	Field13 uint64
	Field14 uint64
	Field15 unsafe.Pointer
	Field16 unsafe.Pointer
}

// IOSurface
type IOSurface struct {
}

// CommandBufferStorageBusyQueue
type CommandBufferStorageBusyQueue struct {
	Tqh_first IOGPUMetalCommandBufferStorageRef
	Tqh_last  *uintptr
}

// GpuStorageQueue
type GpuStorageQueue struct {
	Field1 IOGPUMetalCommandBufferStorageRef
	Field2 *uintptr
}

// MachRightSend
type MachRightSend struct {
	Mrs_name uint32
}

// Mach_right_send is a type alias for MachRightSend for use in objc.Send[T] calls.
type Mach_right_send = MachRightSend

// OpaquePthreadMutex
type OpaquePthreadMutex struct {
	__sig    int64
	__opaque [56]int8
}

// Opaque_pthread_mutex_t is a type alias for OpaquePthreadMutex for use in objc.Send[T] calls.
type Opaque_pthread_mutex_t = OpaquePthreadMutex

// OSUnfairLockS
type OSUnfairLockS struct {
	Field1 uint32
}

// Os_unfair_lock_s is a type alias for OSUnfairLockS for use in objc.Send[T] calls.
type Os_unfair_lock_s = OSUnfairLockS

// SIOGPUKernelCommandSetResourceGroupsArgs
type SIOGPUKernelCommandSetResourceGroupsArgs struct {
}

// SShmemlog
type SShmemlog struct {
}

// S_shmemlog_ is a type alias for SShmemlog for use in objc.Send[T] calls.
type S_shmemlog_ = SShmemlog

// ShmemPoolQueue
type ShmemPoolQueue struct {
	Tqh_first *IOGPUMetalDeviceShmem
	Tqh_last  objc.ID
}

// ShmemlogList
type ShmemlogList struct {
	Field1 SShmemlogRef
}

// Shmemlog_list is a type alias for ShmemlogList for use in objc.Send[T] calls.
type Shmemlog_list = ShmemlogList
