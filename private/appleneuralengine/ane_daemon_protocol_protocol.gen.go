// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _ANEDaemonProtocol protocol.
type ANEDaemonProtocol interface {
	objectivec.IObject

	// CompileModelSandboxExtensionOptionsQosWithReply protocol.
	CompileModelSandboxExtensionOptionsQosWithReply(model objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, qos uint32, reply VoidHandler)

	// CompiledModelExistsForWithReply protocol.
	CompiledModelExistsForWithReply(for_ objectivec.IObject, reply VoidHandler)

	// CompiledModelExistsMatchingHashWithReply protocol.
	CompiledModelExistsMatchingHashWithReply(hash objectivec.IObject, reply VoidHandler)

	// LoadModelSandboxExtensionOptionsQosWithReply protocol.
	LoadModelSandboxExtensionOptionsQosWithReply(model objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, qos uint32, reply VoidHandler)

	// LoadModelNewInstanceOptionsModelInstParamsQosWithReply protocol.
	LoadModelNewInstanceOptionsModelInstParamsQosWithReply(instance objectivec.IObject, options objectivec.IObject, params objectivec.IObject, qos uint32, reply VoidHandler)

	// PrepareChainingWithModelOptionsChainingReqQosWithReply protocol.
	PrepareChainingWithModelOptionsChainingReqQosWithReply(model objectivec.IObject, options objectivec.IObject, req objectivec.IObject, qos uint32, reply VoidHandler)

	// PurgeCompiledModelWithReply protocol.
	PurgeCompiledModelWithReply(model objectivec.IObject, reply VoidHandler)

	// PurgeCompiledModelMatchingHashWithReply protocol.
	PurgeCompiledModelMatchingHashWithReply(hash objectivec.IObject, reply VoidHandler)

	// ReportTelemetryToPPSPlayload protocol.
	ReportTelemetryToPPSPlayload(pps objectivec.IObject, playload objectivec.IObject)

	// UnloadModelOptionsQosWithReply protocol.
	UnloadModelOptionsQosWithReply(model objectivec.IObject, options objectivec.IObject, qos uint32, reply VoidHandler)
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

func (o ANEDaemonProtocolObject) CompileModelSandboxExtensionOptionsQosWithReply(model objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, qos uint32, reply VoidHandler) {
	_block4, _cleanup4 := NewVoidBlock(reply)
	defer _cleanup4()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("compileModel:sandboxExtension:options:qos:withReply:"), model, extension, options, qos, objc.ID(_block4))
}
func (o ANEDaemonProtocolObject) CompiledModelExistsForWithReply(for_ objectivec.IObject, reply VoidHandler) {
	_block1, _cleanup1 := NewVoidBlock(reply)
	defer _cleanup1()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("compiledModelExistsFor:withReply:"), for_, objc.ID(_block1))
}
func (o ANEDaemonProtocolObject) CompiledModelExistsMatchingHashWithReply(hash objectivec.IObject, reply VoidHandler) {
	_block1, _cleanup1 := NewVoidBlock(reply)
	defer _cleanup1()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("compiledModelExistsMatchingHash:withReply:"), hash, objc.ID(_block1))
}
func (o ANEDaemonProtocolObject) LoadModelSandboxExtensionOptionsQosWithReply(model objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, qos uint32, reply VoidHandler) {
	_block4, _cleanup4 := NewVoidBlock(reply)
	defer _cleanup4()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("loadModel:sandboxExtension:options:qos:withReply:"), model, extension, options, qos, objc.ID(_block4))
}
func (o ANEDaemonProtocolObject) LoadModelNewInstanceOptionsModelInstParamsQosWithReply(instance objectivec.IObject, options objectivec.IObject, params objectivec.IObject, qos uint32, reply VoidHandler) {
	_block4, _cleanup4 := NewVoidBlock(reply)
	defer _cleanup4()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("loadModelNewInstance:options:modelInstParams:qos:withReply:"), instance, options, params, qos, objc.ID(_block4))
}
func (o ANEDaemonProtocolObject) PrepareChainingWithModelOptionsChainingReqQosWithReply(model objectivec.IObject, options objectivec.IObject, req objectivec.IObject, qos uint32, reply VoidHandler) {
	_block4, _cleanup4 := NewVoidBlock(reply)
	defer _cleanup4()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("prepareChainingWithModel:options:chainingReq:qos:withReply:"), model, options, req, qos, objc.ID(_block4))
}
func (o ANEDaemonProtocolObject) PurgeCompiledModelWithReply(model objectivec.IObject, reply VoidHandler) {
	_block1, _cleanup1 := NewVoidBlock(reply)
	defer _cleanup1()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("purgeCompiledModel:withReply:"), model, objc.ID(_block1))
}
func (o ANEDaemonProtocolObject) PurgeCompiledModelMatchingHashWithReply(hash objectivec.IObject, reply VoidHandler) {
	_block1, _cleanup1 := NewVoidBlock(reply)
	defer _cleanup1()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("purgeCompiledModelMatchingHash:withReply:"), hash, objc.ID(_block1))
}
func (o ANEDaemonProtocolObject) ReportTelemetryToPPSPlayload(pps objectivec.IObject, playload objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("reportTelemetryToPPS:playload:"), pps, playload)
}
func (o ANEDaemonProtocolObject) UnloadModelOptionsQosWithReply(model objectivec.IObject, options objectivec.IObject, qos uint32, reply VoidHandler) {
	_block3, _cleanup3 := NewVoidBlock(reply)
	defer _cleanup3()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("unloadModel:options:qos:withReply:"), model, options, qos, objc.ID(_block3))
}
