// Code generated from Apple documentation. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// ErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [VZGraphicsDevice._attachDisplayCompletionHandler]
//   - [VZGraphicsDevice._detachDisplayCompletionHandler]
//   - [VZGraphicsDisplay._takeScreenshotWithCompletionHandler]
//   - [VZMacOSRestoreImage._fetchAvailableImagesWithCompletionHandler]
//   - [VZMacOSRestoreImage._fetchLatestSupportedWithOptionsCompletionHandler]
//   - [VZMacOSRestoreImage._loadCatalogWithOptionsCompletionHandler]
//   - [VZMacOSRestoreImage._loadFileURLDeviceClassParserCompletionHandler]
//   - [VZStorageDevice._setAttachmentCompletionHandler]
//   - [VZUSBController._capturePassthroughDevicesWithCompletionHandler]
//   - [VZVirtualMachine._createCoreWithCompletionHandler]
//   - [VZVirtualMachine._createCoresWithCompletionHandler]
//   - [VZVirtualMachine._enterRestrictedModeWithCompletionHandler]
//   - [VZVirtualMachine._getUSBControllerLocationIDWithCompletionHandler]
//   - [VZVirtualMachine._resetWithTypeCompletionHandler]
//   - [VZVirtualMachine._saveMachineStateToURLOptionsCompletionHandler]
//   - [VZXHCIController.AttachDeviceCompletionHandler]
//   - [VZXHCIController.DetachDeviceCompletionHandler]
//   - [VZXHCIController._capturePassthroughDevicesWithCompletionHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [VZGraphicsDevice._attachDisplayCompletionHandler]
//   - [VZGraphicsDevice._detachDisplayCompletionHandler]
//   - [VZGraphicsDisplay._takeScreenshotWithCompletionHandler]
//   - [VZMacOSRestoreImage._fetchAvailableImagesWithCompletionHandler]
//   - [VZMacOSRestoreImage._fetchLatestSupportedWithOptionsCompletionHandler]
//   - [VZMacOSRestoreImage._loadCatalogWithOptionsCompletionHandler]
//   - [VZMacOSRestoreImage._loadFileURLDeviceClassParserCompletionHandler]
//   - [VZStorageDevice._setAttachmentCompletionHandler]
//   - [VZUSBController._capturePassthroughDevicesWithCompletionHandler]
//   - [VZVirtualMachine._createCoreWithCompletionHandler]
//   - [VZVirtualMachine._createCoresWithCompletionHandler]
//   - [VZVirtualMachine._enterRestrictedModeWithCompletionHandler]
//   - [VZVirtualMachine._getUSBControllerLocationIDWithCompletionHandler]
//   - [VZVirtualMachine._resetWithTypeCompletionHandler]
//   - [VZVirtualMachine._saveMachineStateToURLOptionsCompletionHandler]
//   - [VZXHCIController.AttachDeviceCompletionHandler]
//   - [VZXHCIController.DetachDeviceCompletionHandler]
//   - [VZXHCIController._capturePassthroughDevicesWithCompletionHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

