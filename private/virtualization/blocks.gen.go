// Code generated from Apple documentation. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// ErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [VZCustomVirtioDevice.UpdateDeviceSpecificConfigurationCompletionHandler]
//   - [VZDiskImageFormat.CreateDynamicDiskImageWithURLSuggestedSizeCompletionHandler]
//   - [VZDiskImageFormat.CreateFixedDiskImageWithURLSuggestedSizeCompletionHandler]
//   - [VZGraphicsDevice._attachDisplayCompletionHandler]
//   - [VZGraphicsDevice._detachDisplayCompletionHandler]
//   - [VZGraphicsDisplay._takeScreenshotWithCompletionHandler]
//   - [VZMacOSRestoreImage._fetchAvailableImagesWithCompletionHandler]
//   - [VZMacOSRestoreImage._fetchLatestSupportedWithOptionsCompletionHandler]
//   - [VZMacOSRestoreImage._loadCatalogWithOptionsCompletionHandler]
//   - [VZMacOSRestoreImage._loadFileURLDeviceClassParserCompletionHandler]
//   - [VZStorageDevice._setAttachmentCompletionHandler]
//   - [VZTemporaryRAMStorageDeviceAttachment._getAttachmentWithQueueCompletionHandler]
//   - [VZUSBOpticalDriveDeviceConfiguration._getStorageDeviceWithQueueSessionCompletionHandler]
//   - [VZVirtualMachine._createCoreWithCompletionHandler]
//   - [VZVirtualMachine._createCoresWithCompletionHandler]
//   - [VZVirtualMachine._enterRestrictedModeWithCompletionHandler]
//   - [VZVirtualMachine._getUSBControllerLocationIDWithCompletionHandler]
//   - [VZVirtualMachine._resetWithTypeCompletionHandler]
//   - [VZVirtualMachine._saveMachineStateToURLOptionsCompletionHandler]
//   - [VZXHCIController.AttachDeviceCompletionHandler]
//   - [VZXHCIController.DetachDeviceCompletionHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [VZCustomVirtioDevice.UpdateDeviceSpecificConfigurationCompletionHandler]
//   - [VZDiskImageFormat.CreateDynamicDiskImageWithURLSuggestedSizeCompletionHandler]
//   - [VZDiskImageFormat.CreateFixedDiskImageWithURLSuggestedSizeCompletionHandler]
//   - [VZGraphicsDevice._attachDisplayCompletionHandler]
//   - [VZGraphicsDevice._detachDisplayCompletionHandler]
//   - [VZGraphicsDisplay._takeScreenshotWithCompletionHandler]
//   - [VZMacOSRestoreImage._fetchAvailableImagesWithCompletionHandler]
//   - [VZMacOSRestoreImage._fetchLatestSupportedWithOptionsCompletionHandler]
//   - [VZMacOSRestoreImage._loadCatalogWithOptionsCompletionHandler]
//   - [VZMacOSRestoreImage._loadFileURLDeviceClassParserCompletionHandler]
//   - [VZStorageDevice._setAttachmentCompletionHandler]
//   - [VZTemporaryRAMStorageDeviceAttachment._getAttachmentWithQueueCompletionHandler]
//   - [VZUSBOpticalDriveDeviceConfiguration._getStorageDeviceWithQueueSessionCompletionHandler]
//   - [VZVirtualMachine._createCoreWithCompletionHandler]
//   - [VZVirtualMachine._createCoresWithCompletionHandler]
//   - [VZVirtualMachine._enterRestrictedModeWithCompletionHandler]
//   - [VZVirtualMachine._getUSBControllerLocationIDWithCompletionHandler]
//   - [VZVirtualMachine._resetWithTypeCompletionHandler]
//   - [VZVirtualMachine._saveMachineStateToURLOptionsCompletionHandler]
//   - [VZXHCIController.AttachDeviceCompletionHandler]
//   - [VZXHCIController.DetachDeviceCompletionHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	objc.SetNSErrorBlockSignature(block)
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler is the signature for a completion handler block.
//
// Used by:
//   - [VZFramebuffer._takeScreenshotWithCompletionHandlerImageConversionBlock]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [VZFramebuffer._takeScreenshotWithCompletionHandlerImageConversionBlock]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}
