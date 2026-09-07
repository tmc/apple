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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// BoolHandler handles completion with a primitive value.
//
// Used by:
//   - [ANERequest.SetCompletionHandler]
type BoolHandler = func(bool)

// NewBoolBlock wraps a Go [BoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [ANERequest.SetCompletionHandler]
func NewBoolBlock(handler BoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// BoolINSDictionaryErrorHandler handles completion with primitive and object results.
//
// Used by:
//   - [ANECompilerServiceProtocol.CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathMaxModelMemorySizeWithReply]
type BoolINSDictionaryErrorHandler = func(bool, foundation.INSDictionary, error)

// NewBoolINSDictionaryErrorBlock wraps a Go [BoolINSDictionaryErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [ANECompilerServiceProtocol.CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathMaxModelMemorySizeWithReply]
func NewBoolINSDictionaryErrorBlock(handler BoolINSDictionaryErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive bool, extra0ID objc.ID, errID objc.ID) {
		var extra0 foundation.INSDictionary
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			extra0 = foundation.NSDictionaryFromID(extra0ID)
		}
		handler(primitive, extra0, foundation.SafeErrorFrom(errID))
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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}
