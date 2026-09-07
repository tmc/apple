// Code generated from Apple documentation. DO NOT EDIT.

package iogpu

import (
	"github.com/tmc/apple/objc"
)

// VoidHandler is the signature for a completion handler block.
//
// Used by:
//   - [IOGPUMemoryInfo.AddDataSource]
//   - [IOGPUMetal4CommandQueue.SetScheduledHandler]
//   - [IOGPUMetalBuffer.InitStandinWithDeviceBytesNoCopyLengthDeallocator]
//   - [IOGPUMetalBuffer.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocator]
//   - [IOGPUMetalBuffer.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagArgsArgsSizeDeallocator]
//   - [IOGPUMetalBuffer.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizeArgsArgsSizeDeallocator]
//   - [IOGPUMetalBuffer.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizePlacementSparseResidencyBytesArgsArgsSizeDeallocator]
//   - [IOGPUMetalBuffer.InitWithDevicePointerLengthOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocator]
//   - [IOGPUMetalBuffer.InitWithDevicePointerLengthOptionsSysMemSizeVidMemSizeArgsArgsSizeDeallocator]
//   - [IOGPUMetalBuffer.ReplaceBackingWithBytesNoCopyLengthDeallocator]
//   - [IOGPUMetalDevice.NewTiledTextureWithBytesNoCopyLengthDeallocatorDescriptorOffsetBytesPerRow]
//   - [IOGPUMetalMTLLateEvalEvent.NotifyListenerAtValueBlock]
//   - [IOGPUMetalTexture.InitWithDeviceDescriptorSysMemPointerSysMemSizeSysMemLengthSysMemRowBytesArgsArgsSizeDeallocator]
//   - [IOGPUMetalVisibleFunctionTable.InitStandinWithDeviceBytesNoCopyLengthDeallocator]
//   - [IOGPUMetalVisibleFunctionTable.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocator]
//   - [IOGPUMetalVisibleFunctionTable.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagArgsArgsSizeDeallocator]
//   - [IOGPUMetalVisibleFunctionTable.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizeArgsArgsSizeDeallocator]
//   - [IOGPUMetalVisibleFunctionTable.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizePlacementSparseResidencyBytesArgsArgsSizeDeallocator]
//   - [IOGPUMetalVisibleFunctionTable.InitWithDevicePointerLengthOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocator]
//   - [IOGPUMetalVisibleFunctionTable.InitWithDevicePointerLengthOptionsSysMemSizeVidMemSizeArgsArgsSizeDeallocator]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [IOGPUMemoryInfo.AddDataSource]
//   - [IOGPUMetal4CommandQueue.SetScheduledHandler]
//   - [IOGPUMetalBuffer.InitStandinWithDeviceBytesNoCopyLengthDeallocator]
//   - [IOGPUMetalBuffer.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocator]
//   - [IOGPUMetalBuffer.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagArgsArgsSizeDeallocator]
//   - [IOGPUMetalBuffer.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizeArgsArgsSizeDeallocator]
//   - [IOGPUMetalBuffer.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizePlacementSparseResidencyBytesArgsArgsSizeDeallocator]
//   - [IOGPUMetalBuffer.InitWithDevicePointerLengthOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocator]
//   - [IOGPUMetalBuffer.InitWithDevicePointerLengthOptionsSysMemSizeVidMemSizeArgsArgsSizeDeallocator]
//   - [IOGPUMetalBuffer.ReplaceBackingWithBytesNoCopyLengthDeallocator]
//   - [IOGPUMetalDevice.NewTiledTextureWithBytesNoCopyLengthDeallocatorDescriptorOffsetBytesPerRow]
//   - [IOGPUMetalMTLLateEvalEvent.NotifyListenerAtValueBlock]
//   - [IOGPUMetalTexture.InitWithDeviceDescriptorSysMemPointerSysMemSizeSysMemLengthSysMemRowBytesArgsArgsSizeDeallocator]
//   - [IOGPUMetalVisibleFunctionTable.InitStandinWithDeviceBytesNoCopyLengthDeallocator]
//   - [IOGPUMetalVisibleFunctionTable.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocator]
//   - [IOGPUMetalVisibleFunctionTable.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagArgsArgsSizeDeallocator]
//   - [IOGPUMetalVisibleFunctionTable.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizeArgsArgsSizeDeallocator]
//   - [IOGPUMetalVisibleFunctionTable.InitWithDevicePointerLengthAlignmentOptionsSysMemSizeGpuAddressGpuTagPlacementSparsePageSizePlacementSparseResidencyBytesArgsArgsSizeDeallocator]
//   - [IOGPUMetalVisibleFunctionTable.InitWithDevicePointerLengthOptionsSysMemSizeGpuAddressArgsArgsSizeDeallocator]
//   - [IOGPUMetalVisibleFunctionTable.InitWithDevicePointerLengthOptionsSysMemSizeVidMemSizeArgsArgsSizeDeallocator]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}
