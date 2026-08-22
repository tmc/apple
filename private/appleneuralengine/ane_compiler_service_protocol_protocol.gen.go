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
	CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathMaxModelMemorySizeWithReply(at objectivec.IObject, identity objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, directory objectivec.IObject, directory2 objectivec.IObject, url foundation.NSURL, path objectivec.IObject, size uint64, reply BoolHandler)
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
func (o ANECompilerServiceProtocolObject) CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathMaxModelMemorySizeWithReply(at objectivec.IObject, identity objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, directory objectivec.IObject, directory2 objectivec.IObject, url foundation.NSURL, path objectivec.IObject, size uint64, reply BoolHandler) {
	_block9, _cleanup9 := NewBoolBlock(reply)
	defer _cleanup9()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("compileModelAt:csIdentity:sandboxExtension:options:tempDirectory:cloneDirectory:outputURL:aotModelBinaryPath:maxModelMemorySize:withReply:"), at, identity, extension, options, directory, directory2, url, path, size, objc.ID(_block9))
}
