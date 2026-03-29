// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _ANEDaemonProtocol protocol.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonProtocol
type ANEDaemonProtocol interface {
	objectivec.IObject
}

// ANEDaemonProtocolObject wraps an existing Objective-C object that conforms to the ANEDaemonProtocol protocol.
type ANEDaemonProtocolObject struct {
	objectivec.Object
}
func (o ANEDaemonProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// ANEDaemonProtocolObjectFromID constructs a [ANEDaemonProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ANEDaemonProtocolObjectFromID(id objc.ID) ANEDaemonProtocolObject {
	return ANEDaemonProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonProtocol/compileModel:sandboxExtension:options:qos:withReply:
func (o ANEDaemonProtocolObject) CompileModelSandboxExtensionOptionsQosWithReply(model objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, qos uint32, reply VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("compileModel:sandboxExtension:options:qos:withReply:"), model, extension, options, qos, reply)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonProtocol/compiledModelExistsFor:withReply:
func (o ANEDaemonProtocolObject) CompiledModelExistsForWithReply(for_ objectivec.IObject, reply VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("compiledModelExistsFor:withReply:"), for_, reply)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonProtocol/compiledModelExistsMatchingHash:withReply:
func (o ANEDaemonProtocolObject) CompiledModelExistsMatchingHashWithReply(hash objectivec.IObject, reply VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("compiledModelExistsMatchingHash:withReply:"), hash, reply)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonProtocol/loadModel:sandboxExtension:options:qos:withReply:
func (o ANEDaemonProtocolObject) LoadModelSandboxExtensionOptionsQosWithReply(model objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, qos uint32, reply VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("loadModel:sandboxExtension:options:qos:withReply:"), model, extension, options, qos, reply)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonProtocol/loadModelNewInstance:options:modelInstParams:qos:withReply:
func (o ANEDaemonProtocolObject) LoadModelNewInstanceOptionsModelInstParamsQosWithReply(instance objectivec.IObject, options objectivec.IObject, params objectivec.IObject, qos uint32, reply VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("loadModelNewInstance:options:modelInstParams:qos:withReply:"), instance, options, params, qos, reply)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonProtocol/prepareChainingWithModel:options:chainingReq:qos:withReply:
func (o ANEDaemonProtocolObject) PrepareChainingWithModelOptionsChainingReqQosWithReply(model objectivec.IObject, options objectivec.IObject, req objectivec.IObject, qos uint32, reply VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("prepareChainingWithModel:options:chainingReq:qos:withReply:"), model, options, req, qos, reply)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonProtocol/purgeCompiledModel:withReply:
func (o ANEDaemonProtocolObject) PurgeCompiledModelWithReply(model objectivec.IObject, reply VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("purgeCompiledModel:withReply:"), model, reply)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonProtocol/purgeCompiledModelMatchingHash:withReply:
func (o ANEDaemonProtocolObject) PurgeCompiledModelMatchingHashWithReply(hash objectivec.IObject, reply VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("purgeCompiledModelMatchingHash:withReply:"), hash, reply)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonProtocol/reportTelemetryToPPS:playload:
func (o ANEDaemonProtocolObject) ReportTelemetryToPPSPlayload(pps objectivec.IObject, playload objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("reportTelemetryToPPS:playload:"), pps, playload)
	}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonProtocol/unloadModel:options:qos:withReply:
func (o ANEDaemonProtocolObject) UnloadModelOptionsQosWithReply(model objectivec.IObject, options objectivec.IObject, qos uint32, reply VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("unloadModel:options:qos:withReply:"), model, options, qos, reply)
	}

