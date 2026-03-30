// Code generated from Apple documentation. DO NOT EDIT.

package diskimages2

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// ErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [DIAttachParams.SetHandleRefCount]
//   - [DIBaseParams.PrepareImageWithXpcHandlerFileModeError]
//   - [DIChpassParams.ChangePassWithXpcHandlerError]
//   - [DIChpassParams.PrepareImageWithXpcHandlerFileModeEncChpassError]
//   - [DIController2ClientProtocol.AttachCompletedWithHandleReply]
//   - [DICreateParams.CreateEncryptionWithXPCHandlerError]
//   - [DICreateParams.SetFolderCopyXPCHandler]
//   - [DIDeviceHandle.SetClient2IOhandler]
//   - [DIDeviceHandle.SetHandleRefCount]
//   - [DIEncryptionChpass.ReplacePassWithXpcHandlerParamsError]
//   - [DIEncryptionCreator.AddPublicKeyEntryWithXpcHandlerError]
//   - [DIEncryptionCreator.CreateWithXpcHandlerError]
//   - [DIEncryptionFrontend.AddPassphraseEntryWithXpcHandlerFlagsUsageError]
//   - [DIEncryptionFrontend.LookupLegacyKeychainWithXpcHandlerError]
//   - [DIEncryptionFrontend.UnlockWithXpcHandlerError]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [DIAttachParams.SetHandleRefCount]
//   - [DIBaseParams.PrepareImageWithXpcHandlerFileModeError]
//   - [DIChpassParams.ChangePassWithXpcHandlerError]
//   - [DIChpassParams.PrepareImageWithXpcHandlerFileModeEncChpassError]
//   - [DIController2ClientProtocol.AttachCompletedWithHandleReply]
//   - [DICreateParams.CreateEncryptionWithXPCHandlerError]
//   - [DICreateParams.SetFolderCopyXPCHandler]
//   - [DIDeviceHandle.SetClient2IOhandler]
//   - [DIDeviceHandle.SetHandleRefCount]
//   - [DIEncryptionChpass.ReplacePassWithXpcHandlerParamsError]
//   - [DIEncryptionCreator.AddPublicKeyEntryWithXpcHandlerError]
//   - [DIEncryptionCreator.CreateWithXpcHandlerError]
//   - [DIEncryptionFrontend.AddPassphraseEntryWithXpcHandlerFlagsUsageError]
//   - [DIEncryptionFrontend.LookupLegacyKeychainWithXpcHandlerError]
//   - [DIEncryptionFrontend.UnlockWithXpcHandlerError]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler is the signature for a completion handler block.
//
// Used by:
//   - [DIController2ClientDelegate.AttachCompletedWithHandleReply]
//   - [DIConvertParams.ConvertWithCompletionBlock]
//   - [DiskImageCreatorFromFolder.CreateImageWithSrcFolderCompletionBlock]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [DIController2ClientDelegate.AttachCompletedWithHandleReply]
//   - [DIConvertParams.ConvertWithCompletionBlock]
//   - [DiskImageCreatorFromFolder.CreateImageWithSrcFolderCompletionBlock]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}
