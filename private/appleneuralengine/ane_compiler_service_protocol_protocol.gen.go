// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _ANECompilerServiceProtocol protocol.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerServiceProtocol
type ANECompilerServiceProtocol interface {
	objectivec.IObject
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

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerServiceProtocol/compileModelAt:csIdentity:sandboxExtension:options:tempDirectory:cloneDirectory:outputURL:aotModelBinaryPath:withReply:
func (o ANECompilerServiceProtocolObject) CompileModelAtCsIdentitySandboxExtensionOptionsTempDirectoryCloneDirectoryOutputURLAotModelBinaryPathWithReply(at objectivec.IObject, identity objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, directory objectivec.IObject, directory2 objectivec.IObject, url foundation.INSURL, path objectivec.IObject, reply VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("compileModelAt:csIdentity:sandboxExtension:options:tempDirectory:cloneDirectory:outputURL:aotModelBinaryPath:withReply:"), at, identity, extension, options, directory, directory2, url, path, reply)
}
