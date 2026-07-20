// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSXPCService] class.
var (
	_SLSXPCServiceClass     SLSXPCServiceClass
	_SLSXPCServiceClassOnce sync.Once
)

func getSLSXPCServiceClass() SLSXPCServiceClass {
	_SLSXPCServiceClassOnce.Do(func() {
		_SLSXPCServiceClass = SLSXPCServiceClass{class: objc.GetClass("SLSXPCService")}
	})
	return _SLSXPCServiceClass
}

// GetSLSXPCServiceClass returns the class object for SLSXPCService.
func GetSLSXPCServiceClass() SLSXPCServiceClass {
	return getSLSXPCServiceClass()
}

type SLSXPCServiceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSXPCServiceClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSXPCServiceClass) Alloc() SLSXPCService {
	rv := objc.Send[SLSXPCService](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSXPCService.Autoreconnect]
//   - [SLSXPCService.SetAutoreconnect]
//   - [SLSXPCService.CfStringToCStringPtr]
//   - [SLSXPCService.Connected]
//   - [SLSXPCService.SetConnected]
//   - [SLSXPCService.Connection]
//   - [SLSXPCService.SetConnection]
//   - [SLSXPCService.CreateCancellableMachRecvSourceWithQueueErrorCancelAction]
//   - [SLSXPCService.CreateNoSenderRecvPairWithQueueErrorHandlerEventHandler]
//   - [SLSXPCService.CreateXPCDictionary]
//   - [SLSXPCService.Enabled]
//   - [SLSXPCService.SetEnabled]
//   - [SLSXPCService.GetConnectionQueue]
//   - [SLSXPCService.HandleXPCEvent]
//   - [SLSXPCService.MakeNSErrorForCGError]
//   - [SLSXPCService.MakeNSErrorForCocoaError]
//   - [SLSXPCService.MakeNSErrorForMachError]
//   - [SLSXPCService.MakeNSErrorForOSStatus]
//   - [SLSXPCService.MakeNSErrorForPOSIXError]
//   - [SLSXPCService.NotifyQueue]
//   - [SLSXPCService.SetNotifyQueue]
//   - [SLSXPCService.ReinitConnection]
//   - [SLSXPCService.SendNSError]
//   - [SLSXPCService.SendXPCConnectionClosed]
//   - [SLSXPCService.SendXPCDictionary]
//   - [SLSXPCService.SendXPCDictionarySync]
//   - [SLSXPCService.SetClientErrorBlock]
//   - [SLSXPCService.SetClientNotificationBlock]
//   - [SLSXPCService.SetErrorBlock]
//   - [SLSXPCService.SetNotificationBlock]
//   - [SLSXPCService.SetTarget]
//   - [SLSXPCService.SetSetTarget]
//   - [SLSXPCService.InitConnectionWithNameNotificationQueueWithAutoreconnectErrorHandlerNotificationBlock]
//   - [SLSXPCService.InitWithConnectionErrorHandlerNotificationBlock]
//   - [SLSXPCService.DebugDescription]
//   - [SLSXPCService.Description]
//   - [SLSXPCService.Hash]
//   - [SLSXPCService.Superclass]
type SLSXPCService struct {
	objectivec.Object
}

// SLSXPCServiceFromID constructs a [SLSXPCService] from an objc.ID.
func SLSXPCServiceFromID(id objc.ID) SLSXPCService {
	return SLSXPCService{objectivec.Object{ID: id}}
}

// Ensure SLSXPCService implements ISLSXPCService.
var _ ISLSXPCService = SLSXPCService{}

// An interface definition for the [SLSXPCService] class.
//
// # Methods
//
//   - [ISLSXPCService.Autoreconnect]
//   - [ISLSXPCService.SetAutoreconnect]
//   - [ISLSXPCService.CfStringToCStringPtr]
//   - [ISLSXPCService.Connected]
//   - [ISLSXPCService.SetConnected]
//   - [ISLSXPCService.Connection]
//   - [ISLSXPCService.SetConnection]
//   - [ISLSXPCService.CreateCancellableMachRecvSourceWithQueueErrorCancelAction]
//   - [ISLSXPCService.CreateNoSenderRecvPairWithQueueErrorHandlerEventHandler]
//   - [ISLSXPCService.CreateXPCDictionary]
//   - [ISLSXPCService.Enabled]
//   - [ISLSXPCService.SetEnabled]
//   - [ISLSXPCService.GetConnectionQueue]
//   - [ISLSXPCService.HandleXPCEvent]
//   - [ISLSXPCService.MakeNSErrorForCGError]
//   - [ISLSXPCService.MakeNSErrorForCocoaError]
//   - [ISLSXPCService.MakeNSErrorForMachError]
//   - [ISLSXPCService.MakeNSErrorForOSStatus]
//   - [ISLSXPCService.MakeNSErrorForPOSIXError]
//   - [ISLSXPCService.NotifyQueue]
//   - [ISLSXPCService.SetNotifyQueue]
//   - [ISLSXPCService.ReinitConnection]
//   - [ISLSXPCService.SendNSError]
//   - [ISLSXPCService.SendXPCConnectionClosed]
//   - [ISLSXPCService.SendXPCDictionary]
//   - [ISLSXPCService.SendXPCDictionarySync]
//   - [ISLSXPCService.SetClientErrorBlock]
//   - [ISLSXPCService.SetClientNotificationBlock]
//   - [ISLSXPCService.SetErrorBlock]
//   - [ISLSXPCService.SetNotificationBlock]
//   - [ISLSXPCService.SetTarget]
//   - [ISLSXPCService.SetSetTarget]
//   - [ISLSXPCService.InitConnectionWithNameNotificationQueueWithAutoreconnectErrorHandlerNotificationBlock]
//   - [ISLSXPCService.InitWithConnectionErrorHandlerNotificationBlock]
//   - [ISLSXPCService.DebugDescription]
//   - [ISLSXPCService.Description]
//   - [ISLSXPCService.Hash]
//   - [ISLSXPCService.Superclass]
type ISLSXPCService interface {
	objectivec.IObject

	// Topic: Methods

	Autoreconnect() bool
	SetAutoreconnect(value bool)
	CfStringToCStringPtr(ptr objectivec.IObject) string
	Connected() bool
	SetConnected(value bool)
	Connection() objectivec.Object
	SetConnection(value objectivec.Object)
	CreateCancellableMachRecvSourceWithQueueErrorCancelAction(queue objectivec.IObject, error_ []objectivec.IObject, action VoidHandler) objectivec.IObject
	CreateNoSenderRecvPairWithQueueErrorHandlerEventHandler(queue objectivec.IObject, handler VoidHandler, handler2 VoidHandler)
	CreateXPCDictionary(xPCDictionary uint64) objectivec.IObject
	Enabled() bool
	SetEnabled(value bool)
	GetConnectionQueue() objectivec.IObject
	HandleXPCEvent(xPCEvent objectivec.IObject)
	MakeNSErrorForCGError(cGError int64) objectivec.IObject
	MakeNSErrorForCocoaError(error_ int64) objectivec.IObject
	MakeNSErrorForMachError(error_ int64) objectivec.IObject
	MakeNSErrorForOSStatus(oSStatus int64) objectivec.IObject
	MakeNSErrorForPOSIXError(pOSIXError int64) objectivec.IObject
	NotifyQueue() objectivec.Object
	SetNotifyQueue(value objectivec.Object)
	ReinitConnection() bool
	SendNSError(nSError objectivec.IObject)
	SendXPCConnectionClosed() int
	SendXPCDictionary(xPCDictionary objectivec.IObject) int
	SendXPCDictionarySync(sync objectivec.IObject) objectivec.IObject
	SetClientErrorBlock(block VoidHandler)
	SetClientNotificationBlock(block VoidHandler)
	SetErrorBlock(block VoidHandler)
	SetNotificationBlock(block VoidHandler)
	SetTarget() bool
	SetSetTarget(value bool)
	InitConnectionWithNameNotificationQueueWithAutoreconnectErrorHandlerNotificationBlock(name objectivec.IObject, queue objectivec.IObject, autoreconnect bool, handler VoidHandler, block VoidHandler) SLSXPCService
	InitWithConnectionErrorHandlerNotificationBlock(connection objectivec.IObject, handler VoidHandler, block VoidHandler) SLSXPCService
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (s SLSXPCService) Init() SLSXPCService {
	rv := objc.Send[SLSXPCService](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSXPCService) Autorelease() SLSXPCService {
	rv := objc.Send[SLSXPCService](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSXPCService creates a new SLSXPCService instance.
func NewSLSXPCService() SLSXPCService {
	class := getSLSXPCServiceClass()
	rv := objc.Send[SLSXPCService](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (s SLSXPCService) CfStringToCStringPtr(ptr objectivec.IObject) string {
	rv := objc.Send[*byte](s.ID, objc.Sel("cfStringToCStringPtr:"), ptr)
	return objc.GoString(rv)
}
func (s SLSXPCService) CreateCancellableMachRecvSourceWithQueueErrorCancelAction(queue objectivec.IObject, error_ []objectivec.IObject, action VoidHandler) objectivec.IObject {
	_block2, _ := NewVoidBlock(action)
	rv := objc.Send[objc.ID](s.ID, objc.Sel("createCancellableMachRecvSourceWithQueue:error:cancelAction:"), queue, error_, _block2)
	return objectivec.Object{ID: rv}
}
func (s SLSXPCService) CreateNoSenderRecvPairWithQueueErrorHandlerEventHandler(queue objectivec.IObject, handler VoidHandler, handler2 VoidHandler) {
	_block1, _ := NewVoidBlock(handler)
	_block2, _ := NewVoidBlock(handler2)
	objc.Send[objc.ID](s.ID, objc.Sel("createNoSenderRecvPairWithQueue:errorHandler:eventHandler:"), queue, _block1, _block2)
}
func (s SLSXPCService) CreateXPCDictionary(xPCDictionary uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("createXPCDictionary:"), xPCDictionary)
	return objectivec.Object{ID: rv}
}
func (s SLSXPCService) GetConnectionQueue() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("getConnectionQueue"))
	return objectivec.Object{ID: rv}
}
func (s SLSXPCService) HandleXPCEvent(xPCEvent objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("handleXPCEvent:"), xPCEvent)
}
func (s SLSXPCService) MakeNSErrorForCGError(cGError int64) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeNSErrorForCGError:"), cGError)
	return objectivec.Object{ID: rv}
}
func (s SLSXPCService) MakeNSErrorForCocoaError(error_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeNSErrorForCocoaError:"), error_)
	return objectivec.Object{ID: rv}
}
func (s SLSXPCService) MakeNSErrorForMachError(error_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeNSErrorForMachError:"), error_)
	return objectivec.Object{ID: rv}
}
func (s SLSXPCService) MakeNSErrorForOSStatus(oSStatus int64) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeNSErrorForOSStatus:"), oSStatus)
	return objectivec.Object{ID: rv}
}
func (s SLSXPCService) MakeNSErrorForPOSIXError(pOSIXError int64) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeNSErrorForPOSIXError:"), pOSIXError)
	return objectivec.Object{ID: rv}
}
func (s SLSXPCService) ReinitConnection() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("reinitConnection"))
	return rv
}
func (s SLSXPCService) SendNSError(nSError objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("sendNSError:"), nSError)
}
func (s SLSXPCService) SendXPCConnectionClosed() int {
	rv := objc.Send[int](s.ID, objc.Sel("sendXPCConnectionClosed"))
	return rv
}
func (s SLSXPCService) SendXPCDictionary(xPCDictionary objectivec.IObject) int {
	rv := objc.Send[int](s.ID, objc.Sel("sendXPCDictionary:"), xPCDictionary)
	return rv
}
func (s SLSXPCService) SendXPCDictionarySync(sync objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("sendXPCDictionarySync:"), sync)
	return objectivec.Object{ID: rv}
}
func (s SLSXPCService) SetClientErrorBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](s.ID, objc.Sel("setClientErrorBlock:"), _block0)
}
func (s SLSXPCService) SetClientNotificationBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](s.ID, objc.Sel("setClientNotificationBlock:"), _block0)
}
func (s SLSXPCService) SetErrorBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](s.ID, objc.Sel("setErrorBlock:"), _block0)
}
func (s SLSXPCService) SetNotificationBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](s.ID, objc.Sel("setNotificationBlock:"), _block0)
}
func (s SLSXPCService) InitConnectionWithNameNotificationQueueWithAutoreconnectErrorHandlerNotificationBlock(name objectivec.IObject, queue objectivec.IObject, autoreconnect bool, handler VoidHandler, block VoidHandler) SLSXPCService {
	_block3, _ := NewVoidBlock(handler)
	_block4, _ := NewVoidBlock(block)
	rv := objc.Send[SLSXPCService](s.ID, objc.Sel("initConnectionWithName:notificationQueue:withAutoreconnect:errorHandler:notificationBlock:"), name, queue, autoreconnect, _block3, _block4)
	return rv
}
func (s SLSXPCService) InitWithConnectionErrorHandlerNotificationBlock(connection objectivec.IObject, handler VoidHandler, block VoidHandler) SLSXPCService {
	_block1, _ := NewVoidBlock(handler)
	_block2, _ := NewVoidBlock(block)
	rv := objc.Send[SLSXPCService](s.ID, objc.Sel("initWithConnection:errorHandler:notificationBlock:"), connection, _block1, _block2)
	return rv
}

func (s SLSXPCService) Autoreconnect() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("autoreconnect"))
	return rv
}
func (s SLSXPCService) SetAutoreconnect(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setAutoreconnect:"), value)
}
func (s SLSXPCService) Connected() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("connected"))
	return rv
}
func (s SLSXPCService) SetConnected(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setConnected:"), value)
}
func (s SLSXPCService) Connection() objectivec.Object {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("connection"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (s SLSXPCService) SetConnection(value objectivec.Object) {
	objc.Send[struct{}](s.ID, objc.Sel("setConnection:"), value)
}
func (s SLSXPCService) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (s SLSXPCService) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (s SLSXPCService) Enabled() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("enabled"))
	return rv
}
func (s SLSXPCService) SetEnabled(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setEnabled:"), value)
}
func (s SLSXPCService) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}
func (s SLSXPCService) NotifyQueue() objectivec.Object {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("notifyQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (s SLSXPCService) SetNotifyQueue(value objectivec.Object) {
	objc.Send[struct{}](s.ID, objc.Sel("setNotifyQueue:"), value)
}
func (s SLSXPCService) SetTarget() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("setTarget"))
	return rv
}
func (s SLSXPCService) SetSetTarget(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setSetTarget:"), value)
}
func (s SLSXPCService) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](s.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

// CreateNoSenderRecvPairWithQueueErrorHandlerEventHandlerSync is a synchronous wrapper around [SLSXPCService.CreateNoSenderRecvPairWithQueueErrorHandlerEventHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLSXPCService) CreateNoSenderRecvPairWithQueueErrorHandlerEventHandlerSync(ctx context.Context, queue objectivec.IObject, handler func()) error {
	done := make(chan struct{}, 1)
	s.CreateNoSenderRecvPairWithQueueErrorHandlerEventHandler(queue, handler, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetClientErrorBlockSync is a synchronous wrapper around [SLSXPCService.SetClientErrorBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLSXPCService) SetClientErrorBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	s.SetClientErrorBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetClientNotificationBlockSync is a synchronous wrapper around [SLSXPCService.SetClientNotificationBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLSXPCService) SetClientNotificationBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	s.SetClientNotificationBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetErrorBlockSync is a synchronous wrapper around [SLSXPCService.SetErrorBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLSXPCService) SetErrorBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	s.SetErrorBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetNotificationBlockSync is a synchronous wrapper around [SLSXPCService.SetNotificationBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLSXPCService) SetNotificationBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	s.SetNotificationBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// InitConnectionWithNameNotificationQueueWithAutoreconnectErrorHandlerNotificationBlockSync is a synchronous wrapper around [SLSXPCService.InitConnectionWithNameNotificationQueueWithAutoreconnectErrorHandlerNotificationBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLSXPCService) InitConnectionWithNameNotificationQueueWithAutoreconnectErrorHandlerNotificationBlockSync(ctx context.Context, name objectivec.IObject, queue objectivec.IObject, autoreconnect bool, handler func()) error {
	done := make(chan struct{}, 1)
	s.InitConnectionWithNameNotificationQueueWithAutoreconnectErrorHandlerNotificationBlock(name, queue, autoreconnect, handler, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// InitWithConnectionErrorHandlerNotificationBlockSync is a synchronous wrapper around [SLSXPCService.InitWithConnectionErrorHandlerNotificationBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLSXPCService) InitWithConnectionErrorHandlerNotificationBlockSync(ctx context.Context, connection objectivec.IObject, handler func()) error {
	done := make(chan struct{}, 1)
	s.InitWithConnectionErrorHandlerNotificationBlock(connection, handler, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
