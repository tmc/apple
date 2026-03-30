// Code generated from Apple documentation. DO NOT EDIT.

package appleneuralengine

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// BoolErrorHandler handles completion with a boolean result and optional error
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
type BoolErrorHandler = func(bool, error)

// NewBoolErrorBlock wraps a Go [BoolErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewBoolErrorBlock(handler BoolErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// DictionaryErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [ANECompilerServiceProtocol.CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathMaxModelMemorySizeWithReply]
type DictionaryErrorHandler = func(*foundation.INSDictionary, error)

// ErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [ANERequest.SetCompletionHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [ANERequest.SetCompletionHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler is the signature for a completion handler block.
//
// Used by:
//   - [ANECompilerServiceProtocol.CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathWithReply]
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
//   - [ANEDaemonProtocol.CompileModelSandboxExtensionOptionsQosWithReply]
//   - [ANEDaemonProtocol.CompiledModelExistsForWithReply]
//   - [ANEDaemonProtocol.CompiledModelExistsMatchingHashWithReply]
//   - [ANEDaemonProtocol.LoadModelNewInstanceOptionsModelInstParamsQosWithReply]
//   - [ANEDaemonProtocol.LoadModelSandboxExtensionOptionsQosWithReply]
//   - [ANEDaemonProtocol.PrepareChainingWithModelOptionsChainingReqQosWithReply]
//   - [ANEDaemonProtocol.PurgeCompiledModelMatchingHashWithReply]
//   - [ANEDaemonProtocol.PurgeCompiledModelWithReply]
//   - [ANEDaemonProtocol.UnloadModelOptionsQosWithReply]
//   - [ANEStorageMaintainerProtocol.PurgeDanglingModelsAtWithReply]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [ANECompilerServiceProtocol.CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathWithReply]
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
//   - [ANEDaemonProtocol.CompileModelSandboxExtensionOptionsQosWithReply]
//   - [ANEDaemonProtocol.CompiledModelExistsForWithReply]
//   - [ANEDaemonProtocol.CompiledModelExistsMatchingHashWithReply]
//   - [ANEDaemonProtocol.LoadModelNewInstanceOptionsModelInstParamsQosWithReply]
//   - [ANEDaemonProtocol.LoadModelSandboxExtensionOptionsQosWithReply]
//   - [ANEDaemonProtocol.PrepareChainingWithModelOptionsChainingReqQosWithReply]
//   - [ANEDaemonProtocol.PurgeCompiledModelMatchingHashWithReply]
//   - [ANEDaemonProtocol.PurgeCompiledModelWithReply]
//   - [ANEDaemonProtocol.UnloadModelOptionsQosWithReply]
//   - [ANEStorageMaintainerProtocol.PurgeDanglingModelsAtWithReply]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}
