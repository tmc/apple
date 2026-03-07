// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEDaemonConnection] class.
var (
	_ANEDaemonConnectionClass     ANEDaemonConnectionClass
	_ANEDaemonConnectionClassOnce sync.Once
)

func getANEDaemonConnectionClass() ANEDaemonConnectionClass {
	_ANEDaemonConnectionClassOnce.Do(func() {
		_ANEDaemonConnectionClass = ANEDaemonConnectionClass{class: objc.GetClass("_ANEDaemonConnection")}
	})
	return _ANEDaemonConnectionClass
}

// GetANEDaemonConnectionClass returns the class object for _ANEDaemonConnection.
func GetANEDaemonConnectionClass() ANEDaemonConnectionClass {
	return getANEDaemonConnectionClass()
}

type ANEDaemonConnectionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEDaemonConnectionClass) Alloc() ANEDaemonConnection {
	rv := objc.Send[ANEDaemonConnection](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [ANEDaemonConnection.BeginRealTimeTaskWithReply]
//   - [ANEDaemonConnection.CompileModelSandboxExtensionOptionsQosWithReply]
//   - [ANEDaemonConnection.CompiledModelExistsForWithReply]
//   - [ANEDaemonConnection.CompiledModelExistsMatchingHashWithReply]
//   - [ANEDaemonConnection.DaemonConnection]
//   - [ANEDaemonConnection.EchoWithReply]
//   - [ANEDaemonConnection.EndRealTimeTaskWithReply]
//   - [ANEDaemonConnection.LoadModelSandboxExtensionOptionsQosWithReply]
//   - [ANEDaemonConnection.LoadModelNewInstanceOptionsModelInstParamsQosWithReply]
//   - [ANEDaemonConnection.PrepareChainingWithModelOptionsChainingReqQosWithReply]
//   - [ANEDaemonConnection.PurgeCompiledModelWithReply]
//   - [ANEDaemonConnection.PurgeCompiledModelMatchingHashWithReply]
//   - [ANEDaemonConnection.ReportTelemetryToPPSPlayload]
//   - [ANEDaemonConnection.Restricted]
//   - [ANEDaemonConnection.UnloadModelOptionsQosWithReply]
//   - [ANEDaemonConnection.InitWithMachServiceNameRestricted]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonConnection
type ANEDaemonConnection struct {
	objectivec.Object
}

// ANEDaemonConnectionFromID constructs a [ANEDaemonConnection] from an objc.ID.
func ANEDaemonConnectionFromID(id objc.ID) ANEDaemonConnection {
	return ANEDaemonConnection{objectivec.Object{id}}
}
// Ensure ANEDaemonConnection implements IANEDaemonConnection.
var _ IANEDaemonConnection = ANEDaemonConnection{}





// An interface definition for the [ANEDaemonConnection] class.
//
// # Methods
//
//   - [IANEDaemonConnection.BeginRealTimeTaskWithReply]
//   - [IANEDaemonConnection.CompileModelSandboxExtensionOptionsQosWithReply]
//   - [IANEDaemonConnection.CompiledModelExistsForWithReply]
//   - [IANEDaemonConnection.CompiledModelExistsMatchingHashWithReply]
//   - [IANEDaemonConnection.DaemonConnection]
//   - [IANEDaemonConnection.EchoWithReply]
//   - [IANEDaemonConnection.EndRealTimeTaskWithReply]
//   - [IANEDaemonConnection.LoadModelSandboxExtensionOptionsQosWithReply]
//   - [IANEDaemonConnection.LoadModelNewInstanceOptionsModelInstParamsQosWithReply]
//   - [IANEDaemonConnection.PrepareChainingWithModelOptionsChainingReqQosWithReply]
//   - [IANEDaemonConnection.PurgeCompiledModelWithReply]
//   - [IANEDaemonConnection.PurgeCompiledModelMatchingHashWithReply]
//   - [IANEDaemonConnection.ReportTelemetryToPPSPlayload]
//   - [IANEDaemonConnection.Restricted]
//   - [IANEDaemonConnection.UnloadModelOptionsQosWithReply]
//   - [IANEDaemonConnection.InitWithMachServiceNameRestricted]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonConnection
type IANEDaemonConnection interface {
	objectivec.IObject

	// Topic: Methods

	BeginRealTimeTaskWithReply(reply VoidHandler)
	CompileModelSandboxExtensionOptionsQosWithReply(model objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, qos uint32, reply VoidHandler)
	CompiledModelExistsForWithReply(for_ objectivec.IObject, reply VoidHandler)
	CompiledModelExistsMatchingHashWithReply(hash objectivec.IObject, reply VoidHandler)
	DaemonConnection() foundation.NSXPCConnection
	EchoWithReply(echo objectivec.IObject, reply VoidHandler)
	EndRealTimeTaskWithReply(reply VoidHandler)
	LoadModelSandboxExtensionOptionsQosWithReply(model objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, qos uint32, reply VoidHandler)
	LoadModelNewInstanceOptionsModelInstParamsQosWithReply(instance objectivec.IObject, options objectivec.IObject, params objectivec.IObject, qos uint32, reply VoidHandler)
	PrepareChainingWithModelOptionsChainingReqQosWithReply(model objectivec.IObject, options objectivec.IObject, req objectivec.IObject, qos uint32, reply VoidHandler)
	PurgeCompiledModelWithReply(model objectivec.IObject, reply VoidHandler)
	PurgeCompiledModelMatchingHashWithReply(hash objectivec.IObject, reply VoidHandler)
	ReportTelemetryToPPSPlayload(pps objectivec.IObject, playload objectivec.IObject)
	Restricted() bool
	UnloadModelOptionsQosWithReply(model objectivec.IObject, options objectivec.IObject, qos uint32, reply VoidHandler)
	InitWithMachServiceNameRestricted(name objectivec.IObject, restricted bool) ANEDaemonConnection
}





// Init initializes the instance.
func (a ANEDaemonConnection) Init() ANEDaemonConnection {
	rv := objc.Send[ANEDaemonConnection](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEDaemonConnection) Autorelease() ANEDaemonConnection {
	rv := objc.Send[ANEDaemonConnection](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEDaemonConnection creates a new ANEDaemonConnection instance.
func NewANEDaemonConnection() ANEDaemonConnection {
	class := getANEDaemonConnectionClass()
	rv := objc.Send[ANEDaemonConnection](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonConnection/initWithMachServiceName:restricted:
func NewANEDaemonConnectionWithMachServiceNameRestricted(name objectivec.IObject, restricted bool) ANEDaemonConnection {
	instance := getANEDaemonConnectionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMachServiceName:restricted:"), name, restricted)
	return ANEDaemonConnectionFromID(rv)
}







//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonConnection/beginRealTimeTaskWithReply:
func (a ANEDaemonConnection) BeginRealTimeTaskWithReply(reply VoidHandler) {
		_block0, _cleanup0 := NewVoidBlock(reply)
	defer _cleanup0()
		objc.Send[objc.ID](a.ID, objc.Sel("beginRealTimeTaskWithReply:"), _block0)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonConnection/compileModel:sandboxExtension:options:qos:withReply:
func (a ANEDaemonConnection) CompileModelSandboxExtensionOptionsQosWithReply(model objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, qos uint32, reply VoidHandler) {
		_block4, _cleanup4 := NewVoidBlock(reply)
	defer _cleanup4()
		objc.Send[objc.ID](a.ID, objc.Sel("compileModel:sandboxExtension:options:qos:withReply:"), model, extension, options, qos, _block4)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonConnection/compiledModelExistsFor:withReply:
func (a ANEDaemonConnection) CompiledModelExistsForWithReply(for_ objectivec.IObject, reply VoidHandler) {
		_block1, _cleanup1 := NewVoidBlock(reply)
	defer _cleanup1()
		objc.Send[objc.ID](a.ID, objc.Sel("compiledModelExistsFor:withReply:"), for_, _block1)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonConnection/compiledModelExistsMatchingHash:withReply:
func (a ANEDaemonConnection) CompiledModelExistsMatchingHashWithReply(hash objectivec.IObject, reply VoidHandler) {
		_block1, _cleanup1 := NewVoidBlock(reply)
	defer _cleanup1()
		objc.Send[objc.ID](a.ID, objc.Sel("compiledModelExistsMatchingHash:withReply:"), hash, _block1)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonConnection/echo:withReply:
func (a ANEDaemonConnection) EchoWithReply(echo objectivec.IObject, reply VoidHandler) {
		_block1, _cleanup1 := NewVoidBlock(reply)
	defer _cleanup1()
		objc.Send[objc.ID](a.ID, objc.Sel("echo:withReply:"), echo, _block1)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonConnection/endRealTimeTaskWithReply:
func (a ANEDaemonConnection) EndRealTimeTaskWithReply(reply VoidHandler) {
		_block0, _cleanup0 := NewVoidBlock(reply)
	defer _cleanup0()
		objc.Send[objc.ID](a.ID, objc.Sel("endRealTimeTaskWithReply:"), _block0)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonConnection/loadModel:sandboxExtension:options:qos:withReply:
func (a ANEDaemonConnection) LoadModelSandboxExtensionOptionsQosWithReply(model objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, qos uint32, reply VoidHandler) {
		_block4, _cleanup4 := NewVoidBlock(reply)
	defer _cleanup4()
		objc.Send[objc.ID](a.ID, objc.Sel("loadModel:sandboxExtension:options:qos:withReply:"), model, extension, options, qos, _block4)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonConnection/loadModelNewInstance:options:modelInstParams:qos:withReply:
func (a ANEDaemonConnection) LoadModelNewInstanceOptionsModelInstParamsQosWithReply(instance objectivec.IObject, options objectivec.IObject, params objectivec.IObject, qos uint32, reply VoidHandler) {
		_block4, _cleanup4 := NewVoidBlock(reply)
	defer _cleanup4()
		objc.Send[objc.ID](a.ID, objc.Sel("loadModelNewInstance:options:modelInstParams:qos:withReply:"), instance, options, params, qos, _block4)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonConnection/prepareChainingWithModel:options:chainingReq:qos:withReply:
func (a ANEDaemonConnection) PrepareChainingWithModelOptionsChainingReqQosWithReply(model objectivec.IObject, options objectivec.IObject, req objectivec.IObject, qos uint32, reply VoidHandler) {
		_block4, _cleanup4 := NewVoidBlock(reply)
	defer _cleanup4()
		objc.Send[objc.ID](a.ID, objc.Sel("prepareChainingWithModel:options:chainingReq:qos:withReply:"), model, options, req, qos, _block4)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonConnection/purgeCompiledModel:withReply:
func (a ANEDaemonConnection) PurgeCompiledModelWithReply(model objectivec.IObject, reply VoidHandler) {
		_block1, _cleanup1 := NewVoidBlock(reply)
	defer _cleanup1()
		objc.Send[objc.ID](a.ID, objc.Sel("purgeCompiledModel:withReply:"), model, _block1)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonConnection/purgeCompiledModelMatchingHash:withReply:
func (a ANEDaemonConnection) PurgeCompiledModelMatchingHashWithReply(hash objectivec.IObject, reply VoidHandler) {
		_block1, _cleanup1 := NewVoidBlock(reply)
	defer _cleanup1()
		objc.Send[objc.ID](a.ID, objc.Sel("purgeCompiledModelMatchingHash:withReply:"), hash, _block1)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonConnection/reportTelemetryToPPS:playload:
func (a ANEDaemonConnection) ReportTelemetryToPPSPlayload(pps objectivec.IObject, playload objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("reportTelemetryToPPS:playload:"), pps, playload)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonConnection/unloadModel:options:qos:withReply:
func (a ANEDaemonConnection) UnloadModelOptionsQosWithReply(model objectivec.IObject, options objectivec.IObject, qos uint32, reply VoidHandler) {
		_block3, _cleanup3 := NewVoidBlock(reply)
	defer _cleanup3()
		objc.Send[objc.ID](a.ID, objc.Sel("unloadModel:options:qos:withReply:"), model, options, qos, _block3)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonConnection/initWithMachServiceName:restricted:
func (a ANEDaemonConnection) InitWithMachServiceNameRestricted(name objectivec.IObject, restricted bool) ANEDaemonConnection {
	rv := objc.Send[ANEDaemonConnection](a.ID, objc.Sel("initWithMachServiceName:restricted:"), name, restricted)
	return rv
}





// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonConnection/daemonConnectionRestricted
func (_ANEDaemonConnectionClass ANEDaemonConnectionClass) DaemonConnectionRestricted() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEDaemonConnectionClass.class), objc.Sel("daemonConnectionRestricted"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonConnection/userDaemonConnection
func (_ANEDaemonConnectionClass ANEDaemonConnectionClass) UserDaemonConnection() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEDaemonConnectionClass.class), objc.Sel("userDaemonConnection"))
	return objectivec.Object{ID: rv}
}








// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonConnection/daemonConnection
func (a ANEDaemonConnection) DaemonConnection() foundation.NSXPCConnection {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("daemonConnection"))
	return foundation.NSXPCConnectionFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEDaemonConnection/restricted
func (a ANEDaemonConnection) Restricted() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("restricted"))
	return rv
}












// BeginRealTimeTaskWithReplySync is a synchronous wrapper around [ANEDaemonConnection.BeginRealTimeTaskWithReply].
// It blocks until the completion handler fires or the context is cancelled.
func (a ANEDaemonConnection) BeginRealTimeTaskWithReplySync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	a.BeginRealTimeTaskWithReply(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// CompileModelSandboxExtensionOptionsQosWithReplySync is a synchronous wrapper around [ANEDaemonConnection.CompileModelSandboxExtensionOptionsQosWithReply].
// It blocks until the completion handler fires or the context is cancelled.
func (a ANEDaemonConnection) CompileModelSandboxExtensionOptionsQosWithReplySync(ctx context.Context, model objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, qos uint32) error {
	done := make(chan struct{}, 1)
	a.CompileModelSandboxExtensionOptionsQosWithReply(model, extension, options, qos, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// CompiledModelExistsForWithReplySync is a synchronous wrapper around [ANEDaemonConnection.CompiledModelExistsForWithReply].
// It blocks until the completion handler fires or the context is cancelled.
func (a ANEDaemonConnection) CompiledModelExistsForWithReplySync(ctx context.Context, for_ objectivec.IObject) error {
	done := make(chan struct{}, 1)
	a.CompiledModelExistsForWithReply(for_, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// CompiledModelExistsMatchingHashWithReplySync is a synchronous wrapper around [ANEDaemonConnection.CompiledModelExistsMatchingHashWithReply].
// It blocks until the completion handler fires or the context is cancelled.
func (a ANEDaemonConnection) CompiledModelExistsMatchingHashWithReplySync(ctx context.Context, hash objectivec.IObject) error {
	done := make(chan struct{}, 1)
	a.CompiledModelExistsMatchingHashWithReply(hash, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EchoWithReplySync is a synchronous wrapper around [ANEDaemonConnection.EchoWithReply].
// It blocks until the completion handler fires or the context is cancelled.
func (a ANEDaemonConnection) EchoWithReplySync(ctx context.Context, echo objectivec.IObject) error {
	done := make(chan struct{}, 1)
	a.EchoWithReply(echo, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EndRealTimeTaskWithReplySync is a synchronous wrapper around [ANEDaemonConnection.EndRealTimeTaskWithReply].
// It blocks until the completion handler fires or the context is cancelled.
func (a ANEDaemonConnection) EndRealTimeTaskWithReplySync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	a.EndRealTimeTaskWithReply(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// LoadModelSandboxExtensionOptionsQosWithReplySync is a synchronous wrapper around [ANEDaemonConnection.LoadModelSandboxExtensionOptionsQosWithReply].
// It blocks until the completion handler fires or the context is cancelled.
func (a ANEDaemonConnection) LoadModelSandboxExtensionOptionsQosWithReplySync(ctx context.Context, model objectivec.IObject, extension objectivec.IObject, options objectivec.IObject, qos uint32) error {
	done := make(chan struct{}, 1)
	a.LoadModelSandboxExtensionOptionsQosWithReply(model, extension, options, qos, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// LoadModelNewInstanceOptionsModelInstParamsQosWithReplySync is a synchronous wrapper around [ANEDaemonConnection.LoadModelNewInstanceOptionsModelInstParamsQosWithReply].
// It blocks until the completion handler fires or the context is cancelled.
func (a ANEDaemonConnection) LoadModelNewInstanceOptionsModelInstParamsQosWithReplySync(ctx context.Context, instance objectivec.IObject, options objectivec.IObject, params objectivec.IObject, qos uint32) error {
	done := make(chan struct{}, 1)
	a.LoadModelNewInstanceOptionsModelInstParamsQosWithReply(instance, options, params, qos, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PrepareChainingWithModelOptionsChainingReqQosWithReplySync is a synchronous wrapper around [ANEDaemonConnection.PrepareChainingWithModelOptionsChainingReqQosWithReply].
// It blocks until the completion handler fires or the context is cancelled.
func (a ANEDaemonConnection) PrepareChainingWithModelOptionsChainingReqQosWithReplySync(ctx context.Context, model objectivec.IObject, options objectivec.IObject, req objectivec.IObject, qos uint32) error {
	done := make(chan struct{}, 1)
	a.PrepareChainingWithModelOptionsChainingReqQosWithReply(model, options, req, qos, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PurgeCompiledModelWithReplySync is a synchronous wrapper around [ANEDaemonConnection.PurgeCompiledModelWithReply].
// It blocks until the completion handler fires or the context is cancelled.
func (a ANEDaemonConnection) PurgeCompiledModelWithReplySync(ctx context.Context, model objectivec.IObject) error {
	done := make(chan struct{}, 1)
	a.PurgeCompiledModelWithReply(model, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PurgeCompiledModelMatchingHashWithReplySync is a synchronous wrapper around [ANEDaemonConnection.PurgeCompiledModelMatchingHashWithReply].
// It blocks until the completion handler fires or the context is cancelled.
func (a ANEDaemonConnection) PurgeCompiledModelMatchingHashWithReplySync(ctx context.Context, hash objectivec.IObject) error {
	done := make(chan struct{}, 1)
	a.PurgeCompiledModelMatchingHashWithReply(hash, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// UnloadModelOptionsQosWithReplySync is a synchronous wrapper around [ANEDaemonConnection.UnloadModelOptionsQosWithReply].
// It blocks until the completion handler fires or the context is cancelled.
func (a ANEDaemonConnection) UnloadModelOptionsQosWithReplySync(ctx context.Context, model objectivec.IObject, options objectivec.IObject, qos uint32) error {
	done := make(chan struct{}, 1)
	a.UnloadModelOptionsQosWithReply(model, options, qos, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}






