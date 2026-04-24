// Code generated from Apple documentation for remotecoreml. DO NOT EDIT.

package remotecoreml

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNetworking] class.
var (
	_MLNetworkingClass     MLNetworkingClass
	_MLNetworkingClassOnce sync.Once
)

func getMLNetworkingClass() MLNetworkingClass {
	_MLNetworkingClassOnce.Do(func() {
		_MLNetworkingClass = MLNetworkingClass{class: objc.GetClass("_MLNetworking")}
	})
	return _MLNetworkingClass
}

// GetMLNetworkingClass returns the class object for _MLNetworking.
func GetMLNetworkingClass() MLNetworkingClass {
	return getMLNetworkingClass()
}

type MLNetworkingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNetworkingClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNetworkingClass) Alloc() MLNetworking {
	rv := objc.Send[MLNetworking](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLNetworking.Connection]
//   - [MLNetworking.SetConnection]
//   - [MLNetworking.IsClient]
//   - [MLNetworking.Listener]
//   - [MLNetworking.LogType]
//   - [MLNetworking.NwOptions]
//   - [MLNetworking.Parameters]
//   - [MLNetworking.Protocol_stack]
//   - [MLNetworking.Q]
//   - [MLNetworking.ReceiveLoop]
//   - [MLNetworking.RestartConnection]
//   - [MLNetworking.SendData]
//   - [MLNetworking.SetListenerReceiveDataCallBack]
//   - [MLNetworking.SetReceiveDataCallBack]
//   - [MLNetworking.StartConnection]
//   - [MLNetworking.Stop]
//   - [MLNetworking.InitConnection]
//   - [MLNetworking.InitListener]
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworking
type MLNetworking struct {
	objectivec.Object
}

// MLNetworkingFromID constructs a [MLNetworking] from an objc.ID.
func MLNetworkingFromID(id objc.ID) MLNetworking {
	return MLNetworking{objectivec.Object{ID: id}}
}

// Ensure MLNetworking implements IMLNetworking.
var _ IMLNetworking = MLNetworking{}

// An interface definition for the [MLNetworking] class.
//
// # Methods
//
//   - [IMLNetworking.Connection]
//   - [IMLNetworking.SetConnection]
//   - [IMLNetworking.IsClient]
//   - [IMLNetworking.Listener]
//   - [IMLNetworking.LogType]
//   - [IMLNetworking.NwOptions]
//   - [IMLNetworking.Parameters]
//   - [IMLNetworking.Protocol_stack]
//   - [IMLNetworking.Q]
//   - [IMLNetworking.ReceiveLoop]
//   - [IMLNetworking.RestartConnection]
//   - [IMLNetworking.SendData]
//   - [IMLNetworking.SetListenerReceiveDataCallBack]
//   - [IMLNetworking.SetReceiveDataCallBack]
//   - [IMLNetworking.StartConnection]
//   - [IMLNetworking.Stop]
//   - [IMLNetworking.InitConnection]
//   - [IMLNetworking.InitListener]
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworking
type IMLNetworking interface {
	objectivec.IObject

	// Topic: Methods

	Connection() objectivec.Object
	SetConnection(value objectivec.Object)
	IsClient() bool
	Listener() objectivec.Object
	LogType() objectivec.Object
	NwOptions() unsafe.Pointer
	Parameters() objectivec.Object
	Protocol_stack() objectivec.Object
	Q() objectivec.Object
	ReceiveLoop(loop VoidHandler)
	RestartConnection()
	SendData(data objectivec.IObject)
	SetListenerReceiveDataCallBack(back VoidHandler)
	SetReceiveDataCallBack(back VoidHandler)
	StartConnection()
	Stop()
	InitConnection(connection objectivec.IObject) MLNetworking
	InitListener(listener objectivec.IObject) MLNetworking
}

// Init initializes the instance.
func (m MLNetworking) Init() MLNetworking {
	rv := objc.Send[MLNetworking](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLNetworking) Autorelease() MLNetworking {
	rv := objc.Send[MLNetworking](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNetworking creates a new MLNetworking instance.
func NewMLNetworking() MLNetworking {
	class := getMLNetworkingClass()
	rv := objc.Send[MLNetworking](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworking/initConnection:
func NewMLNetworkingConnection(connection objectivec.IObject) MLNetworking {
	instance := getMLNetworkingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initConnection:"), connection)
	return MLNetworkingFromID(rv)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworking/initListener:
func NewMLNetworkingListener(listener objectivec.IObject) MLNetworking {
	instance := getMLNetworkingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initListener:"), listener)
	return MLNetworkingFromID(rv)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworking/receiveLoop:
func (m MLNetworking) ReceiveLoop(loop VoidHandler) {
	_block0, _ := NewVoidBlock(loop)
	objc.Send[objc.ID](m.ID, objc.Sel("receiveLoop:"), _block0)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworking/restartConnection
func (m MLNetworking) RestartConnection() {
	objc.Send[objc.ID](m.ID, objc.Sel("restartConnection"))
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworking/sendData:
func (m MLNetworking) SendData(data objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("sendData:"), data)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworking/setListenerReceiveDataCallBack:
func (m MLNetworking) SetListenerReceiveDataCallBack(back VoidHandler) {
	_block0, _ := NewVoidBlock(back)
	objc.Send[objc.ID](m.ID, objc.Sel("setListenerReceiveDataCallBack:"), _block0)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworking/setReceiveDataCallBack:
func (m MLNetworking) SetReceiveDataCallBack(back VoidHandler) {
	_block0, _ := NewVoidBlock(back)
	objc.Send[objc.ID](m.ID, objc.Sel("setReceiveDataCallBack:"), _block0)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworking/startConnection
func (m MLNetworking) StartConnection() {
	objc.Send[objc.ID](m.ID, objc.Sel("startConnection"))
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworking/stop
func (m MLNetworking) Stop() {
	objc.Send[objc.ID](m.ID, objc.Sel("stop"))
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworking/initConnection:
func (m MLNetworking) InitConnection(connection objectivec.IObject) MLNetworking {
	rv := objc.Send[MLNetworking](m.ID, objc.Sel("initConnection:"), connection)
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworking/initListener:
func (m MLNetworking) InitListener(listener objectivec.IObject) MLNetworking {
	rv := objc.Send[MLNetworking](m.ID, objc.Sel("initListener:"), listener)
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworking/connection
func (m MLNetworking) Connection() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("connection"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (m MLNetworking) SetConnection(value objectivec.Object) {
	objc.Send[struct{}](m.ID, objc.Sel("setConnection:"), value)
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworking/isClient
func (m MLNetworking) IsClient() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isClient"))
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworking/listener
func (m MLNetworking) Listener() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("listener"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworking/logType
func (m MLNetworking) LogType() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("logType"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworking/nwOptions
func (m MLNetworking) NwOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("nwOptions"))
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworking/parameters
func (m MLNetworking) Parameters() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parameters"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworking/protocol_stack
func (m MLNetworking) Protocol_stack() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("protocol_stack"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworking/q
func (m MLNetworking) Q() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("q"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

// ReceiveLoopSync is a synchronous wrapper around [MLNetworking.ReceiveLoop].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLNetworking) ReceiveLoopSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	m.ReceiveLoop(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetListenerReceiveDataCallBackSync is a synchronous wrapper around [MLNetworking.SetListenerReceiveDataCallBack].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLNetworking) SetListenerReceiveDataCallBackSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	m.SetListenerReceiveDataCallBack(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetReceiveDataCallBackSync is a synchronous wrapper around [MLNetworking.SetReceiveDataCallBack].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLNetworking) SetReceiveDataCallBackSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	m.SetReceiveDataCallBack(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
