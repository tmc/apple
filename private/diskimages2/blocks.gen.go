// Code generated from Apple documentation. DO NOT EDIT.

package diskimages2

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// ErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [DIController2ClientProtocol.AttachCompletedWithHandleReply]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [DIController2ClientProtocol.AttachCompletedWithHandleReply]
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
//   - [DIController2ClientDelegate.AttachCompletedWithHandleReply]
//   - [DIConvertParams.ConvertWithCompletionBlock]
//   - [DiskImageCreatorFromFolder.CreateImageWithSrcFolderCompletionBlock]
//   - [DiskImages2.ConvertWithParamsCompletionBlock]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [DIController2ClientDelegate.AttachCompletedWithHandleReply]
//   - [DIConvertParams.ConvertWithCompletionBlock]
//   - [DiskImageCreatorFromFolder.CreateImageWithSrcFolderCompletionBlock]
//   - [DiskImages2.ConvertWithParamsCompletionBlock]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}
