// Code generated from Apple documentation for remotecoreml. DO NOT EDIT.

package remotecoreml

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLServer] class.
var (
	_MLServerClass     MLServerClass
	_MLServerClassOnce sync.Once
)

func getMLServerClass() MLServerClass {
	_MLServerClassOnce.Do(func() {
		_MLServerClass = MLServerClass{class: objc.GetClass("_MLServer")}
	})
	return _MLServerClass
}

// GetMLServerClass returns the class object for _MLServer.
func GetMLServerClass() MLServerClass {
	return getMLServerClass()
}

type MLServerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLServerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLServerClass) Alloc() MLServer {
	rv := objc.Send[MLServer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLServer.DoReceiveContextIsCompleteError]
//   - [MLServer.NwObj]
//   - [MLServer.NwOptions]
//   - [MLServer.Packet]
//   - [MLServer.Q]
//   - [MLServer.SetLoadCommand]
//   - [MLServer.SetLoadFunction]
//   - [MLServer.SetPredictCommand]
//   - [MLServer.SetPredictFunction]
//   - [MLServer.SetTextCommand]
//   - [MLServer.SetTextFunction]
//   - [MLServer.SetUnLoadCommand]
//   - [MLServer.SetUnLoadFunction]
//   - [MLServer.Start]
//   - [MLServer.Stop]
//   - [MLServer.InitWithOptions]
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLServer
type MLServer struct {
	objectivec.Object
}

// MLServerFromID constructs a [MLServer] from an objc.ID.
func MLServerFromID(id objc.ID) MLServer {
	return MLServer{objectivec.Object{ID: id}}
}

// Ensure MLServer implements IMLServer.
var _ IMLServer = MLServer{}

// An interface definition for the [MLServer] class.
//
// # Methods
//
//   - [IMLServer.DoReceiveContextIsCompleteError]
//   - [IMLServer.NwObj]
//   - [IMLServer.NwOptions]
//   - [IMLServer.Packet]
//   - [IMLServer.Q]
//   - [IMLServer.SetLoadCommand]
//   - [IMLServer.SetLoadFunction]
//   - [IMLServer.SetPredictCommand]
//   - [IMLServer.SetPredictFunction]
//   - [IMLServer.SetTextCommand]
//   - [IMLServer.SetTextFunction]
//   - [IMLServer.SetUnLoadCommand]
//   - [IMLServer.SetUnLoadFunction]
//   - [IMLServer.Start]
//   - [IMLServer.Stop]
//   - [IMLServer.InitWithOptions]
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLServer
type IMLServer interface {
	objectivec.IObject

	// Topic: Methods

	DoReceiveContextIsCompleteError(receive objectivec.IObject, context objectivec.IObject, complete bool, error_ objectivec.IObject)
	NwObj() *MLNetworking
	NwOptions() *MLNetworkOptions
	Packet() *MLNetworkPacket
	Q() objectivec.Object
	SetLoadCommand(command VoidHandler)
	SetLoadFunction(function VoidHandler)
	SetPredictCommand(command VoidHandler)
	SetPredictFunction(function VoidHandler)
	SetTextCommand(command VoidHandler)
	SetTextFunction(function VoidHandler)
	SetUnLoadCommand(command VoidHandler)
	SetUnLoadFunction(function VoidHandler)
	Start()
	Stop()
	InitWithOptions(options objectivec.IObject) MLServer
}

// Init initializes the instance.
func (m MLServer) Init() MLServer {
	rv := objc.Send[MLServer](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLServer) Autorelease() MLServer {
	rv := objc.Send[MLServer](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLServer creates a new MLServer instance.
func NewMLServer() MLServer {
	class := getMLServerClass()
	rv := objc.Send[MLServer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLServer/initWithOptions:
func NewMLServerWithOptions(options objectivec.IObject) MLServer {
	instance := getMLServerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOptions:"), options)
	return MLServerFromID(rv)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLServer/doReceive:context:isComplete:error:
func (m MLServer) DoReceiveContextIsCompleteError(receive objectivec.IObject, context objectivec.IObject, complete bool, error_ objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("doReceive:context:isComplete:error:"), receive, context, complete, error_)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLServer/setLoadCommand:
func (m MLServer) SetLoadCommand(command VoidHandler) {
	_block0, _ := NewVoidBlock(command)
	objc.Send[objc.ID](m.ID, objc.Sel("setLoadCommand:"), _block0)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLServer/setLoadFunction:
func (m MLServer) SetLoadFunction(function VoidHandler) {
	_block0, _ := NewVoidBlock(function)
	objc.Send[objc.ID](m.ID, objc.Sel("setLoadFunction:"), _block0)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLServer/setPredictCommand:
func (m MLServer) SetPredictCommand(command VoidHandler) {
	_block0, _ := NewVoidBlock(command)
	objc.Send[objc.ID](m.ID, objc.Sel("setPredictCommand:"), _block0)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLServer/setPredictFunction:
func (m MLServer) SetPredictFunction(function VoidHandler) {
	_block0, _ := NewVoidBlock(function)
	objc.Send[objc.ID](m.ID, objc.Sel("setPredictFunction:"), _block0)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLServer/setTextCommand:
func (m MLServer) SetTextCommand(command VoidHandler) {
	_block0, _ := NewVoidBlock(command)
	objc.Send[objc.ID](m.ID, objc.Sel("setTextCommand:"), _block0)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLServer/setTextFunction:
func (m MLServer) SetTextFunction(function VoidHandler) {
	_block0, _ := NewVoidBlock(function)
	objc.Send[objc.ID](m.ID, objc.Sel("setTextFunction:"), _block0)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLServer/setUnLoadCommand:
func (m MLServer) SetUnLoadCommand(command VoidHandler) {
	_block0, _ := NewVoidBlock(command)
	objc.Send[objc.ID](m.ID, objc.Sel("setUnLoadCommand:"), _block0)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLServer/setUnLoadFunction:
func (m MLServer) SetUnLoadFunction(function VoidHandler) {
	_block0, _ := NewVoidBlock(function)
	objc.Send[objc.ID](m.ID, objc.Sel("setUnLoadFunction:"), _block0)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLServer/start
func (m MLServer) Start() {
	objc.Send[objc.ID](m.ID, objc.Sel("start"))
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLServer/stop
func (m MLServer) Stop() {
	objc.Send[objc.ID](m.ID, objc.Sel("stop"))
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLServer/initWithOptions:
func (m MLServer) InitWithOptions(options objectivec.IObject) MLServer {
	rv := objc.Send[MLServer](m.ID, objc.Sel("initWithOptions:"), options)
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLServer/nwObj
func (m MLServer) NwObj() *MLNetworking {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("nwObj"))
	if rv == 0 {
		return nil
	}
	val := MLNetworkingFromID(objc.ID(rv))
	return &val
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLServer/nwOptions
func (m MLServer) NwOptions() *MLNetworkOptions {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("nwOptions"))
	if rv == 0 {
		return nil
	}
	val := MLNetworkOptionsFromID(objc.ID(rv))
	return &val
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLServer/packet
func (m MLServer) Packet() *MLNetworkPacket {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("packet"))
	if rv == 0 {
		return nil
	}
	val := MLNetworkPacketFromID(objc.ID(rv))
	return &val
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLServer/q
func (m MLServer) Q() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("q"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

// SetLoadCommandSync is a synchronous wrapper around [MLServer.SetLoadCommand].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLServer) SetLoadCommandSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	m.SetLoadCommand(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetLoadFunctionSync is a synchronous wrapper around [MLServer.SetLoadFunction].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLServer) SetLoadFunctionSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	m.SetLoadFunction(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetPredictCommandSync is a synchronous wrapper around [MLServer.SetPredictCommand].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLServer) SetPredictCommandSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	m.SetPredictCommand(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetPredictFunctionSync is a synchronous wrapper around [MLServer.SetPredictFunction].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLServer) SetPredictFunctionSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	m.SetPredictFunction(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetTextCommandSync is a synchronous wrapper around [MLServer.SetTextCommand].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLServer) SetTextCommandSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	m.SetTextCommand(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetTextFunctionSync is a synchronous wrapper around [MLServer.SetTextFunction].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLServer) SetTextFunctionSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	m.SetTextFunction(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetUnLoadCommandSync is a synchronous wrapper around [MLServer.SetUnLoadCommand].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLServer) SetUnLoadCommandSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	m.SetUnLoadCommand(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetUnLoadFunctionSync is a synchronous wrapper around [MLServer.SetUnLoadFunction].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLServer) SetUnLoadFunctionSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	m.SetUnLoadFunction(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
