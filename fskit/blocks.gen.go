// Code generated from Apple documentation. DO NOT EDIT.

package fskit

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// BoolErrorHandler handles A block or closure to indicate success or failure.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [FSVolumeAccessCheckOperations.CheckAccessToItemRequestedAccessReplyHandler]
type BoolErrorHandler = func(bool, error)

// NewBoolErrorBlock wraps a Go [BoolErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [FSVolumeAccessCheckOperations.CheckAccessToItemRequestedAccessReplyHandler]
func NewBoolErrorBlock(handler BoolErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// DataErrorHandler handles A block or closure to indicate success or failure.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [FSVolumeXattrOperations.GetXattrNamedOfItemReplyHandler]
type DataErrorHandler = func(*foundation.NSData, error)

// NewDataErrorBlock wraps a Go [DataErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [FSVolumeXattrOperations.GetXattrNamedOfItemReplyHandler]
func NewDataErrorBlock(handler DataErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *foundation.NSData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSDataFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// ErrorHandler handles A block or closure that executes after the wipe operation completes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [FSFileSystem.WipeResourceCompletionHandler]
//   - [FSFileSystemBase.WipeResourceCompletionHandler]
//   - [FSUnaryFileSystem.WipeResourceCompletionHandler]
//   - [FSUnaryFileSystemOperations.UnloadResourceOptionsReplyHandler]
//   - [FSVolumeItemDeactivation.DeactivateItemReplyHandler]
//   - [FSVolumeKernelOffloadedIOOperations.BlockmapFileOffsetLengthFlagsOperationIDPackerReplyHandler]
//   - [FSVolumeKernelOffloadedIOOperations.CompleteIOForFileOffsetLengthStatusFlagsOperationIDReplyHandler]
//   - [FSVolumeOpenCloseOperations.CloseItemKeepingModesReplyHandler]
//   - [FSVolumeOpenCloseOperations.OpenItemWithModesReplyHandler]
//   - [FSVolumeOperations.DeactivateWithOptionsReplyHandler]
//   - [FSVolumeOperations.MountWithOptionsReplyHandler]
//   - [FSVolumeOperations.ReclaimItemReplyHandler]
//   - [FSVolumeOperations.RemoveItemNamedFromDirectoryReplyHandler]
//   - [FSVolumeOperations.SynchronizeWithFlagsReplyHandler]
//   - [FSVolumeXattrOperations.SetXattrNamedToDataOnItemPolicyReplyHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [FSFileSystem.WipeResourceCompletionHandler]
//   - [FSFileSystemBase.WipeResourceCompletionHandler]
//   - [FSUnaryFileSystem.WipeResourceCompletionHandler]
//   - [FSUnaryFileSystemOperations.UnloadResourceOptionsReplyHandler]
//   - [FSVolumeItemDeactivation.DeactivateItemReplyHandler]
//   - [FSVolumeKernelOffloadedIOOperations.BlockmapFileOffsetLengthFlagsOperationIDPackerReplyHandler]
//   - [FSVolumeKernelOffloadedIOOperations.CompleteIOForFileOffsetLengthStatusFlagsOperationIDReplyHandler]
//   - [FSVolumeOpenCloseOperations.CloseItemKeepingModesReplyHandler]
//   - [FSVolumeOpenCloseOperations.OpenItemWithModesReplyHandler]
//   - [FSVolumeOperations.DeactivateWithOptionsReplyHandler]
//   - [FSVolumeOperations.MountWithOptionsReplyHandler]
//   - [FSVolumeOperations.ReclaimItemReplyHandler]
//   - [FSVolumeOperations.RemoveItemNamedFromDirectoryReplyHandler]
//   - [FSVolumeOperations.SynchronizeWithFlagsReplyHandler]
//   - [FSVolumeXattrOperations.SetXattrNamedToDataOnItemPolicyReplyHandler]
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

// FSDirectoryVerifierErrorHandler handles A block or closure to indicate success or failure.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [FSVolumeOperations.EnumerateDirectoryStartingAtCookieVerifierProvidingAttributesUsingPackerReplyHandler]
type FSDirectoryVerifierErrorHandler = func(FSDirectoryVerifier, error)

// NewFSDirectoryVerifierErrorBlock wraps a Go [FSDirectoryVerifierErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [FSVolumeOperations.EnumerateDirectoryStartingAtCookieVerifierProvidingAttributesUsingPackerReplyHandler]
func NewFSDirectoryVerifierErrorBlock(handler FSDirectoryVerifierErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal FSDirectoryVerifier, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// FSFileNameArrayErrorHandler handles A block or closure to indicate success or failure.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [FSVolumeXattrOperations.ListXattrsOfItemReplyHandler]
type FSFileNameArrayErrorHandler = func(*[]FSFileName, error)

// NewFSFileNameArrayErrorBlock wraps a Go [FSFileNameArrayErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [FSVolumeXattrOperations.ListXattrsOfItemReplyHandler]
func NewFSFileNameArrayErrorBlock(handler FSFileNameArrayErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *[]FSFileName
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]FSFileName, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = FSFileNameFromID(item.GetID())
			}
			result = &res
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// FSFileNameErrorHandler handles A block or closure to indicate success or failure.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [FSVolumeOperations.CreateLinkToItemNamedInDirectoryReplyHandler]
//   - [FSVolumeOperations.ReadSymbolicLinkReplyHandler]
//   - [FSVolumeOperations.RenameItemInDirectoryNamedToNewNameInDirectoryOverItemReplyHandler]
//   - [FSVolumeRenameOperations.SetVolumeNameReplyHandler]
type FSFileNameErrorHandler = func(*FSFileName, error)

// NewFSFileNameErrorBlock wraps a Go [FSFileNameErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [FSVolumeOperations.CreateLinkToItemNamedInDirectoryReplyHandler]
//   - [FSVolumeOperations.ReadSymbolicLinkReplyHandler]
//   - [FSVolumeOperations.RenameItemInDirectoryNamedToNewNameInDirectoryOverItemReplyHandler]
//   - [FSVolumeRenameOperations.SetVolumeNameReplyHandler]
func NewFSFileNameErrorBlock(handler FSFileNameErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *FSFileName
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := FSFileNameFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// FSItemAttributesErrorHandler handles A block or closure to indicate success or failure.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [FSVolumeOperations.GetAttributesOfItemReplyHandler]
//   - [FSVolumeOperations.SetAttributesOnItemReplyHandler]
type FSItemAttributesErrorHandler = func(*FSItemAttributes, error)

// NewFSItemAttributesErrorBlock wraps a Go [FSItemAttributesErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [FSVolumeOperations.GetAttributesOfItemReplyHandler]
//   - [FSVolumeOperations.SetAttributesOnItemReplyHandler]
func NewFSItemAttributesErrorBlock(handler FSItemAttributesErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *FSItemAttributes
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := FSItemAttributesFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// FSItemErrorHandler handles A block or closure to indicate success or failure.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [FSVolumeOperations.ActivateWithOptionsReplyHandler]
type FSItemErrorHandler = func(*FSItem, error)

// NewFSItemErrorBlock wraps a Go [FSItemErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [FSVolumeOperations.ActivateWithOptionsReplyHandler]
func NewFSItemErrorBlock(handler FSItemErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *FSItem
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := FSItemFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// FSItemFSFileNameErrorHandler handles A block or closure to indicate success or failure.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [FSVolumeKernelOffloadedIOOperations.CreateFileNamedInDirectoryAttributesPackerReplyHandler]
//   - [FSVolumeKernelOffloadedIOOperations.LookupItemNamedInDirectoryPackerReplyHandler]
//   - [FSVolumeOperations.CreateItemNamedTypeInDirectoryAttributesReplyHandler]
//   - [FSVolumeOperations.CreateSymbolicLinkNamedInDirectoryAttributesLinkContentsReplyHandler]
//   - [FSVolumeOperations.LookupItemNamedInDirectoryReplyHandler]
type FSItemFSFileNameErrorHandler = func(*FSItem, *FSFileName, error)

// NewFSItemFSFileNameErrorBlock wraps a Go [FSItemFSFileNameErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [FSVolumeKernelOffloadedIOOperations.CreateFileNamedInDirectoryAttributesPackerReplyHandler]
//   - [FSVolumeKernelOffloadedIOOperations.LookupItemNamedInDirectoryPackerReplyHandler]
//   - [FSVolumeOperations.CreateItemNamedTypeInDirectoryAttributesReplyHandler]
//   - [FSVolumeOperations.CreateSymbolicLinkNamedInDirectoryAttributesLinkContentsReplyHandler]
//   - [FSVolumeOperations.LookupItemNamedInDirectoryReplyHandler]
func NewFSItemFSFileNameErrorBlock(handler FSItemFSFileNameErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, errID objc.ID) {
		var result *FSItem
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := FSItemFromID(resultID)
			result = &v
		}
		var extra0 *FSFileName
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := FSFileNameFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// FSModuleIdentityArrayErrorHandler handles A block or closure that executes when FSKit finishes its fetch process.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [FSClient.FetchInstalledExtensionsWithCompletionHandler]
type FSModuleIdentityArrayErrorHandler = func(*[]FSModuleIdentity, error)

// NewFSModuleIdentityArrayErrorBlock wraps a Go [FSModuleIdentityArrayErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [FSClient.FetchInstalledExtensionsWithCompletionHandler]
func NewFSModuleIdentityArrayErrorBlock(handler FSModuleIdentityArrayErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *[]FSModuleIdentity
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]FSModuleIdentity, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = FSModuleIdentityFromID(item.GetID())
			}
			result = &res
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// FSProbeResultErrorHandler handles A block or closure that your implementation invokes when it finishes the probe or encounters an error.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [FSUnaryFileSystemOperations.ProbeResourceReplyHandler]
type FSProbeResultErrorHandler = func(*FSProbeResult, error)

// NewFSProbeResultErrorBlock wraps a Go [FSProbeResultErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [FSUnaryFileSystemOperations.ProbeResourceReplyHandler]
func NewFSProbeResultErrorBlock(handler FSProbeResultErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *FSProbeResult
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := FSProbeResultFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// FSVolumeErrorHandler handles A block or closure that your implementation invokes when it finishes setting up or encounters an error.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [FSUnaryFileSystemOperations.LoadResourceOptionsReplyHandler]
type FSVolumeErrorHandler = func(*FSVolume, error)

// NewFSVolumeErrorBlock wraps a Go [FSVolumeErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [FSUnaryFileSystemOperations.LoadResourceOptionsReplyHandler]
func NewFSVolumeErrorBlock(handler FSVolumeErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *FSVolume
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := FSVolumeFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// NSErrorVoidHandler is the signature for a completion handler block.
type NSErrorVoidHandler = func() foundation.NSError

// NewNSErrorVoidBlock wraps a Go [NSErrorVoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNSErrorVoidBlock(handler NSErrorVoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) objc.ID {
		return handler().ID
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler handles A block or closure to indicate success or failure.
//
// Used by:
//   - [FSVolumeOperations.UnmountWithReplyHandler]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [FSVolumeOperations.UnmountWithReplyHandler]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// size_tErrorHandler handles A block that executes after the read operation completes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [FSBlockDeviceResource.ReadIntoStartingAtLengthCompletionHandler]
//   - [FSBlockDeviceResource.WriteFromStartingAtLengthCompletionHandler]
//   - [FSVolumeKernelOffloadedIOOperations.PreallocateSpaceForFileAtOffsetLengthFlagsPackerReplyHandler]
//   - [FSVolumePreallocateOperations.PreallocateSpaceForItemAtOffsetLengthFlagsReplyHandler]
//   - [FSVolumeReadWriteOperations.ReadFromFileOffsetLengthIntoBufferReplyHandler]
//   - [FSVolumeReadWriteOperations.WriteContentsToFileAtOffsetReplyHandler]
type size_tErrorHandler = func(uintptr, error)

// Newsize_tErrorBlock wraps a Go [size_tErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [FSBlockDeviceResource.ReadIntoStartingAtLengthCompletionHandler]
//   - [FSBlockDeviceResource.WriteFromStartingAtLengthCompletionHandler]
//   - [FSVolumeKernelOffloadedIOOperations.PreallocateSpaceForFileAtOffsetLengthFlagsPackerReplyHandler]
//   - [FSVolumePreallocateOperations.PreallocateSpaceForItemAtOffsetLengthFlagsReplyHandler]
//   - [FSVolumeReadWriteOperations.ReadFromFileOffsetLengthIntoBufferReplyHandler]
//   - [FSVolumeReadWriteOperations.WriteContentsToFileAtOffsetReplyHandler]
func Newsize_tErrorBlock(handler size_tErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal uintptr, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}
