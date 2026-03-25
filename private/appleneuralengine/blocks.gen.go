// Code generated from Apple documentation. DO NOT EDIT.

package appleneuralengine

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// BoolErrorHandler handles completion with primitive result and optional error.
//
// Used by:
//   - [ANERequest.SetCompletionHandler]
type BoolErrorHandler = func(bool, error)

// NewBoolErrorBlock wraps a Go [BoolErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [ANERequest.SetCompletionHandler]
func NewBoolErrorBlock(handler BoolErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool, errID objc.ID) {
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(primitiveVal, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler is the signature for a completion handler block.
//
// Used by:
//   - [ANEDaemonConnection.BeginRealTimeTaskWithReply]
//   - [ANEDaemonConnection.CompileModelSandboxExtensionOptionsQosWithReply]
//   - [ANEDaemonConnection.CompiledModelExistsForWithReply]
//   - [ANEDaemonConnection.CompiledModelExistsMatchingHashWithReply]
//   - [ANEDaemonConnection.EchoWithReply]
//   - [ANEDaemonConnection.EndRealTimeTaskWithReply]
//   - [ANEDaemonConnection.LoadModelNewInstanceOptionsModelInstParamsQosWithReply]
//   - [ANEDaemonConnection.LoadModelSandboxExtensionOptionsQosWithReply]
//   - [ANEDaemonConnection.PrepareChainingWithModelOptionsChainingReqQosWithReply]
//   - [ANEDaemonConnection.PurgeCompiledModelMatchingHashWithReply]
//   - [ANEDaemonConnection.PurgeCompiledModelWithReply]
//   - [ANEDaemonConnection.UnloadModelOptionsQosWithReply]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [ANEDaemonConnection.BeginRealTimeTaskWithReply]
//   - [ANEDaemonConnection.CompileModelSandboxExtensionOptionsQosWithReply]
//   - [ANEDaemonConnection.CompiledModelExistsForWithReply]
//   - [ANEDaemonConnection.CompiledModelExistsMatchingHashWithReply]
//   - [ANEDaemonConnection.EchoWithReply]
//   - [ANEDaemonConnection.EndRealTimeTaskWithReply]
//   - [ANEDaemonConnection.LoadModelNewInstanceOptionsModelInstParamsQosWithReply]
//   - [ANEDaemonConnection.LoadModelSandboxExtensionOptionsQosWithReply]
//   - [ANEDaemonConnection.PrepareChainingWithModelOptionsChainingReqQosWithReply]
//   - [ANEDaemonConnection.PurgeCompiledModelMatchingHashWithReply]
//   - [ANEDaemonConnection.PurgeCompiledModelWithReply]
//   - [ANEDaemonConnection.UnloadModelOptionsQosWithReply]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

