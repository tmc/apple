// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _ANECompilerServiceProtocol protocol.
type ANECompilerServiceProtocol interface {
	objectivec.IObject

	// CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathWithReply protocol.
	CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathWithReply(at objectivec.IObject, identity objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, directory objectivec.IObject, directory2 objectivec.IObject, url foundation.NSURL, path objectivec.IObject, reply VoidHandler)

	// CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathMaxModelMemorySizeWithReply protocol.
	CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathMaxModelMemorySizeWithReply(at objectivec.IObject, identity objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, directory objectivec.IObject, directory2 objectivec.IObject, url foundation.NSURL, path objectivec.IObject, size uint64, reply BoolINSDictionaryErrorHandler)

	// MoveCachedModelFromSourceToDestinationWithReply protocol.
	MoveCachedModelFromSourceToDestinationWithReply(source objectivec.IObject, destination objectivec.IObject, reply BoolErrorHandler)

	// UpdateSourcePathAtToWithContainerAtWithContainerWithReply protocol.
	UpdateSourcePathAtToWithContainerAtWithContainerWithReply(at objectivec.IObject, to objectivec.IObject, at2 objectivec.IObject, container objectivec.IObject, reply BoolErrorHandler)
}

// ANECompilerServiceProtocolObject wraps an existing Objective-C object that conforms to the ANECompilerServiceProtocol protocol.
type ANECompilerServiceProtocolObject struct {
	objectivec.Object
}

func (o ANECompilerServiceProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// ANECompilerServiceProtocolObjectFromID constructs a [ANECompilerServiceProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ANECompilerServiceProtocolObjectFromID(id objc.ID) ANECompilerServiceProtocolObject {
	return ANECompilerServiceProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o ANECompilerServiceProtocolObject) CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathWithReply(at objectivec.IObject, identity objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, directory objectivec.IObject, directory2 objectivec.IObject, url foundation.NSURL, path objectivec.IObject, reply VoidHandler) {
	_block8, _cleanup8 := NewVoidBlock(reply)
	defer _cleanup8()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("compileModelAt:csIdentity:sandboxExtension:options:tempDirectory:cloneDirectory:outputURL:aotModelBinaryPath:withReply:"), at, identity, extension, options, directory, directory2, url, path, objc.ID(_block8))
}
func (o ANECompilerServiceProtocolObject) CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathMaxModelMemorySizeWithReply(at objectivec.IObject, identity objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, directory objectivec.IObject, directory2 objectivec.IObject, url foundation.NSURL, path objectivec.IObject, size uint64, reply BoolINSDictionaryErrorHandler) {
	_block9, _cleanup9 := NewBoolINSDictionaryErrorBlock(reply)
	defer _cleanup9()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("compileModelAt:csIdentity:sandboxExtension:options:tempDirectory:cloneDirectory:outputURL:aotModelBinaryPath:maxModelMemorySize:withReply:"), at, identity, extension, options, directory, directory2, url, path, size, objc.ID(_block9))
}
func (o ANECompilerServiceProtocolObject) MoveCachedModelFromSourceToDestinationWithReply(source objectivec.IObject, destination objectivec.IObject, reply BoolErrorHandler) {
	_block2, _cleanup2 := NewBoolErrorBlock(reply)
	defer _cleanup2()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("moveCachedModelFromSource:toDestination:withReply:"), source, destination, objc.ID(_block2))
}
func (o ANECompilerServiceProtocolObject) UpdateSourcePathAtToWithContainerAtWithContainerWithReply(at objectivec.IObject, to objectivec.IObject, at2 objectivec.IObject, container objectivec.IObject, reply BoolErrorHandler) {
	_block4, _cleanup4 := NewBoolErrorBlock(reply)
	defer _cleanup4()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("updateSourcePathAt:to:withContainerAt:withContainer:withReply:"), at, to, at2, container, objc.ID(_block4))
}
