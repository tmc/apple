// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXRemoteViewEventManager] class.
var (
	_CPXRemoteViewEventManagerClass     CPXRemoteViewEventManagerClass
	_CPXRemoteViewEventManagerClassOnce sync.Once
)

func getCPXRemoteViewEventManagerClass() CPXRemoteViewEventManagerClass {
	_CPXRemoteViewEventManagerClassOnce.Do(func() {
		_CPXRemoteViewEventManagerClass = CPXRemoteViewEventManagerClass{class: objc.GetClass("CPXRemoteViewEventManager")}
	})
	return _CPXRemoteViewEventManagerClass
}

// GetCPXRemoteViewEventManagerClass returns the class object for CPXRemoteViewEventManager.
func GetCPXRemoteViewEventManagerClass() CPXRemoteViewEventManagerClass {
	return getCPXRemoteViewEventManagerClass()
}

type CPXRemoteViewEventManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXRemoteViewEventManagerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXRemoteViewEventManagerClass) Alloc() CPXRemoteViewEventManager {
	rv := objc.Send[CPXRemoteViewEventManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXRemoteViewEventManager.ClientCount]
//   - [CPXRemoteViewEventManager.DidReceiveConnectionConfig]
//   - [CPXRemoteViewEventManager.InvalidateConnections]
//   - [CPXRemoteViewEventManager.PassEventUpstreamToHostFullDispatchReply]
//   - [CPXRemoteViewEventManager.PidForCurrentConnection]
//   - [CPXRemoteViewEventManager.SendEventToHostPidFullDispatchReply]
//   - [CPXRemoteViewEventManager.InitWithDeliveryManager]
type CPXRemoteViewEventManager struct {
	objectivec.Object
}

// CPXRemoteViewEventManagerFromID constructs a [CPXRemoteViewEventManager] from an objc.ID.
func CPXRemoteViewEventManagerFromID(id objc.ID) CPXRemoteViewEventManager {
	return CPXRemoteViewEventManager{objectivec.Object{ID: id}}
}

// Ensure CPXRemoteViewEventManager implements ICPXRemoteViewEventManager.
var _ ICPXRemoteViewEventManager = CPXRemoteViewEventManager{}

// An interface definition for the [CPXRemoteViewEventManager] class.
//
// # Methods
//
//   - [ICPXRemoteViewEventManager.ClientCount]
//   - [ICPXRemoteViewEventManager.DidReceiveConnectionConfig]
//   - [ICPXRemoteViewEventManager.InvalidateConnections]
//   - [ICPXRemoteViewEventManager.PassEventUpstreamToHostFullDispatchReply]
//   - [ICPXRemoteViewEventManager.PidForCurrentConnection]
//   - [ICPXRemoteViewEventManager.SendEventToHostPidFullDispatchReply]
//   - [ICPXRemoteViewEventManager.InitWithDeliveryManager]
type ICPXRemoteViewEventManager interface {
	objectivec.IObject

	// Topic: Methods

	ClientCount() uint64
	DidReceiveConnectionConfig(connection objectivec.IObject, config objectivec.IObject)
	InvalidateConnections()
	PassEventUpstreamToHostFullDispatchReply(host objectivec.IObject, dispatch objectivec.IObject, reply VoidHandler)
	PidForCurrentConnection() int
	SendEventToHostPidFullDispatchReply(event unsafe.Pointer, pid int, dispatch objectivec.IObject, reply VoidHandler)
	InitWithDeliveryManager(manager objectivec.IObject) CPXRemoteViewEventManager
}

// Init initializes the instance.
func (c CPXRemoteViewEventManager) Init() CPXRemoteViewEventManager {
	rv := objc.Send[CPXRemoteViewEventManager](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXRemoteViewEventManager) Autorelease() CPXRemoteViewEventManager {
	rv := objc.Send[CPXRemoteViewEventManager](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXRemoteViewEventManager creates a new CPXRemoteViewEventManager instance.
func NewCPXRemoteViewEventManager() CPXRemoteViewEventManager {
	class := getCPXRemoteViewEventManagerClass()
	rv := objc.Send[CPXRemoteViewEventManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewCPXRemoteViewEventManagerWithDeliveryManager(manager objectivec.IObject) CPXRemoteViewEventManager {
	instance := getCPXRemoteViewEventManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDeliveryManager:"), manager)
	return CPXRemoteViewEventManagerFromID(rv)
}

func (c CPXRemoteViewEventManager) ClientCount() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("clientCount"))
	return rv
}
func (c CPXRemoteViewEventManager) DidReceiveConnectionConfig(connection objectivec.IObject, config objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("didReceiveConnection:config:"), connection, config)
}
func (c CPXRemoteViewEventManager) InvalidateConnections() {
	objc.Send[objc.ID](c.ID, objc.Sel("invalidateConnections"))
}
func (c CPXRemoteViewEventManager) PassEventUpstreamToHostFullDispatchReply(host objectivec.IObject, dispatch objectivec.IObject, reply VoidHandler) {
	_block2, _ := NewVoidBlock(reply)
	objc.Send[objc.ID](c.ID, objc.Sel("passEventUpstreamToHost:fullDispatch:reply:"), host, dispatch, _block2)
}
func (c CPXRemoteViewEventManager) PidForCurrentConnection() int {
	rv := objc.Send[int](c.ID, objc.Sel("pidForCurrentConnection"))
	return rv
}
func (c CPXRemoteViewEventManager) SendEventToHostPidFullDispatchReply(event unsafe.Pointer, pid int, dispatch objectivec.IObject, reply VoidHandler) {
	_block3, _ := NewVoidBlock(reply)
	objc.Send[objc.ID](c.ID, objc.Sel("sendEvent:toHostPid:fullDispatch:reply:"), event, pid, dispatch, _block3)
}
func (c CPXRemoteViewEventManager) InitWithDeliveryManager(manager objectivec.IObject) CPXRemoteViewEventManager {
	rv := objc.Send[CPXRemoteViewEventManager](c.ID, objc.Sel("initWithDeliveryManager:"), manager)
	return rv
}

// PassEventUpstreamToHostFullDispatchReplySync is a synchronous wrapper around [CPXRemoteViewEventManager.PassEventUpstreamToHostFullDispatchReply].
// It blocks until the completion handler fires or the context is cancelled.
func (c CPXRemoteViewEventManager) PassEventUpstreamToHostFullDispatchReplySync(ctx context.Context, host objectivec.IObject, dispatch objectivec.IObject) error {
	done := make(chan struct{}, 1)
	c.PassEventUpstreamToHostFullDispatchReply(host, dispatch, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SendEventToHostPidFullDispatchReplySync is a synchronous wrapper around [CPXRemoteViewEventManager.SendEventToHostPidFullDispatchReply].
// It blocks until the completion handler fires or the context is cancelled.
func (c CPXRemoteViewEventManager) SendEventToHostPidFullDispatchReplySync(ctx context.Context, event unsafe.Pointer, pid int, dispatch objectivec.IObject) error {
	done := make(chan struct{}, 1)
	c.SendEventToHostPidFullDispatchReply(event, pid, dispatch, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
