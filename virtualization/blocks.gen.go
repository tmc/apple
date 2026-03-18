// Code generated from Apple documentation. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// ErrorHandler handles The completion handler the framework invokes after the request finishes processing.
//   - error: A result type that contains an error object when the VM fails to start.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [VZLinuxRosettaDirectoryShare.InstallRosettaWithCompletionHandler]
//   - [VZMacOSInstaller.InstallWithCompletionHandler]
//   - [VZUSBController.AttachDeviceCompletionHandler]
//   - [VZUSBController.DetachDeviceCompletionHandler]
//   - [VZVirtualMachine.PauseWithCompletionHandler]
//   - [VZVirtualMachine.RestoreMachineStateFromURLCompletionHandler]
//   - [VZVirtualMachine.ResumeWithCompletionHandler]
//   - [VZVirtualMachine.SaveMachineStateToURLCompletionHandler]
//   - [VZVirtualMachine.StartWithCompletionHandler]
//   - [VZVirtualMachine.StartWithOptionsCompletionHandler]
//   - [VZVirtualMachine.StopWithCompletionHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [VZLinuxRosettaDirectoryShare.InstallRosettaWithCompletionHandler]
//   - [VZMacOSInstaller.InstallWithCompletionHandler]
//   - [VZUSBController.AttachDeviceCompletionHandler]
//   - [VZUSBController.DetachDeviceCompletionHandler]
//   - [VZVirtualMachine.PauseWithCompletionHandler]
//   - [VZVirtualMachine.RestoreMachineStateFromURLCompletionHandler]
//   - [VZVirtualMachine.ResumeWithCompletionHandler]
//   - [VZVirtualMachine.SaveMachineStateToURLCompletionHandler]
//   - [VZVirtualMachine.StartWithCompletionHandler]
//   - [VZVirtualMachine.StartWithOptionsCompletionHandler]
//   - [VZVirtualMachine.StopWithCompletionHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// MacOSRestoreImageErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [VZMacOSRestoreImage.FetchLatestSupportedWithCompletionHandler]
//   - [VZMacOSRestoreImage.LoadFileURLCompletionHandler]
type MacOSRestoreImageErrorHandler = func(*VZMacOSRestoreImage, error)

// NewMacOSRestoreImageErrorBlock wraps a Go [MacOSRestoreImageErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [VZMacOSRestoreImage.FetchLatestSupportedWithCompletionHandler]
//   - [VZMacOSRestoreImage.LoadFileURLCompletionHandler]
func NewMacOSRestoreImageErrorBlock(handler MacOSRestoreImageErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *VZMacOSRestoreImage
		if resultID != 0 {
			v := VZMacOSRestoreImageFromID(resultID)
			result = &v
		}
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// VirtioSocketConnectionErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [VZVirtioSocketDevice.ConnectToPortCompletionHandler]
type VirtioSocketConnectionErrorHandler = func(*VZVirtioSocketConnection, error)

// NewVirtioSocketConnectionErrorBlock wraps a Go [VirtioSocketConnectionErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [VZVirtioSocketDevice.ConnectToPortCompletionHandler]
func NewVirtioSocketConnectionErrorBlock(handler VirtioSocketConnectionErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *VZVirtioSocketConnection
		if resultID != 0 {
			v := VZVirtioSocketConnectionFromID(resultID)
			result = &v
		}
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

