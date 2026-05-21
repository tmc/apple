// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSRemoteViewEventClient] class.
var (
	_SLSRemoteViewEventClientClass     SLSRemoteViewEventClientClass
	_SLSRemoteViewEventClientClassOnce sync.Once
)

func getSLSRemoteViewEventClientClass() SLSRemoteViewEventClientClass {
	_SLSRemoteViewEventClientClassOnce.Do(func() {
		_SLSRemoteViewEventClientClass = SLSRemoteViewEventClientClass{class: objc.GetClass("SLSRemoteViewEventClient")}
	})
	return _SLSRemoteViewEventClientClass
}

// GetSLSRemoteViewEventClientClass returns the class object for SLSRemoteViewEventClient.
func GetSLSRemoteViewEventClientClass() SLSRemoteViewEventClientClass {
	return getSLSRemoteViewEventClientClass()
}

type SLSRemoteViewEventClientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSRemoteViewEventClientClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSRemoteViewEventClientClass) Alloc() SLSRemoteViewEventClient {
	rv := objc.Send[SLSRemoteViewEventClient](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSRemoteViewEventClient.ActivateWithHandlerInvalidationHandler]
//   - [SLSRemoteViewEventClient.DeferringEnvironmentFromEvent]
//   - [SLSRemoteViewEventClient.DeferringTokenFromEvent]
//   - [SLSRemoteViewEventClient.Delegate]
//   - [SLSRemoteViewEventClient.SetDelegate]
//   - [SLSRemoteViewEventClient.Invalidate]
//   - [SLSRemoteViewEventClient.SendEventToHostFullDispatchReply]
//   - [SLSRemoteViewEventClient.ServicePassEventUpstreamToHostFullDispatchReply]
//   - [SLSRemoteViewEventClient.SetDeferringTokenAndEnvironmentInEvent]
//   - [SLSRemoteViewEventClient.InitWithConfig]
//   - [SLSRemoteViewEventClient.DebugDescription]
//   - [SLSRemoteViewEventClient.Description]
//   - [SLSRemoteViewEventClient.Hash]
//   - [SLSRemoteViewEventClient.Superclass]
type SLSRemoteViewEventClient struct {
	objectivec.Object
}

// SLSRemoteViewEventClientFromID constructs a [SLSRemoteViewEventClient] from an objc.ID.
func SLSRemoteViewEventClientFromID(id objc.ID) SLSRemoteViewEventClient {
	return SLSRemoteViewEventClient{objectivec.Object{ID: id}}
}

// Ensure SLSRemoteViewEventClient implements ISLSRemoteViewEventClient.
var _ ISLSRemoteViewEventClient = SLSRemoteViewEventClient{}

// An interface definition for the [SLSRemoteViewEventClient] class.
//
// # Methods
//
//   - [ISLSRemoteViewEventClient.ActivateWithHandlerInvalidationHandler]
//   - [ISLSRemoteViewEventClient.DeferringEnvironmentFromEvent]
//   - [ISLSRemoteViewEventClient.DeferringTokenFromEvent]
//   - [ISLSRemoteViewEventClient.Delegate]
//   - [ISLSRemoteViewEventClient.SetDelegate]
//   - [ISLSRemoteViewEventClient.Invalidate]
//   - [ISLSRemoteViewEventClient.SendEventToHostFullDispatchReply]
//   - [ISLSRemoteViewEventClient.ServicePassEventUpstreamToHostFullDispatchReply]
//   - [ISLSRemoteViewEventClient.SetDeferringTokenAndEnvironmentInEvent]
//   - [ISLSRemoteViewEventClient.InitWithConfig]
//   - [ISLSRemoteViewEventClient.DebugDescription]
//   - [ISLSRemoteViewEventClient.Description]
//   - [ISLSRemoteViewEventClient.Hash]
//   - [ISLSRemoteViewEventClient.Superclass]
type ISLSRemoteViewEventClient interface {
	objectivec.IObject

	// Topic: Methods

	ActivateWithHandlerInvalidationHandler(handler VoidHandler, handler2 VoidHandler)
	DeferringEnvironmentFromEvent(event unsafe.Pointer) objectivec.IObject
	DeferringTokenFromEvent(event unsafe.Pointer) objectivec.IObject
	Delegate() unsafe.Pointer
	SetDelegate(value kernel.Pointer)
	Invalidate()
	SendEventToHostFullDispatchReply(host objectivec.IObject, dispatch objectivec.IObject, reply VoidHandler)
	ServicePassEventUpstreamToHostFullDispatchReply(host unsafe.Pointer, dispatch bool, reply VoidHandler)
	SetDeferringTokenAndEnvironmentInEvent(token objectivec.IObject, environment objectivec.IObject, event unsafe.Pointer)
	InitWithConfig(config objectivec.IObject) SLSRemoteViewEventClient
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (s SLSRemoteViewEventClient) Init() SLSRemoteViewEventClient {
	rv := objc.Send[SLSRemoteViewEventClient](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSRemoteViewEventClient) Autorelease() SLSRemoteViewEventClient {
	rv := objc.Send[SLSRemoteViewEventClient](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSRemoteViewEventClient creates a new SLSRemoteViewEventClient instance.
func NewSLSRemoteViewEventClient() SLSRemoteViewEventClient {
	class := getSLSRemoteViewEventClientClass()
	rv := objc.Send[SLSRemoteViewEventClient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSRemoteViewEventClientWithConfig(config objectivec.IObject) SLSRemoteViewEventClient {
	instance := getSLSRemoteViewEventClientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfig:"), config)
	return SLSRemoteViewEventClientFromID(rv)
}

func (s SLSRemoteViewEventClient) ActivateWithHandlerInvalidationHandler(handler VoidHandler, handler2 VoidHandler) {
	_block0, _ := NewVoidBlock(handler)
	_block1, _ := NewVoidBlock(handler2)
	objc.Send[objc.ID](s.ID, objc.Sel("activateWithHandler:invalidationHandler:"), _block0, _block1)
}
func (s SLSRemoteViewEventClient) DeferringEnvironmentFromEvent(event unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("deferringEnvironmentFromEvent:"), event)
	return objectivec.Object{ID: rv}
}
func (s SLSRemoteViewEventClient) DeferringTokenFromEvent(event unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("deferringTokenFromEvent:"), event)
	return objectivec.Object{ID: rv}
}
func (s SLSRemoteViewEventClient) Invalidate() {
	objc.Send[objc.ID](s.ID, objc.Sel("invalidate"))
}
func (s SLSRemoteViewEventClient) SendEventToHostFullDispatchReply(host objectivec.IObject, dispatch objectivec.IObject, reply VoidHandler) {
	_block2, _ := NewVoidBlock(reply)
	objc.Send[objc.ID](s.ID, objc.Sel("sendEventToHost:fullDispatch:reply:"), host, dispatch, _block2)
}
func (s SLSRemoteViewEventClient) ServicePassEventUpstreamToHostFullDispatchReply(host unsafe.Pointer, dispatch bool, reply VoidHandler) {
	_block2, _ := NewVoidBlock(reply)
	objc.Send[objc.ID](s.ID, objc.Sel("servicePassEventUpstreamToHost:fullDispatch:reply:"), host, dispatch, _block2)
}
func (s SLSRemoteViewEventClient) SetDeferringTokenAndEnvironmentInEvent(token objectivec.IObject, environment objectivec.IObject, event unsafe.Pointer) {
	objc.Send[objc.ID](s.ID, objc.Sel("setDeferringToken:andEnvironment:inEvent:"), token, environment, event)
}
func (s SLSRemoteViewEventClient) InitWithConfig(config objectivec.IObject) SLSRemoteViewEventClient {
	rv := objc.Send[SLSRemoteViewEventClient](s.ID, objc.Sel("initWithConfig:"), config)
	return rv
}

func (_SLSRemoteViewEventClientClass SLSRemoteViewEventClientClass) PossibleEventDescriptorForVirtualKeyCode(code uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSRemoteViewEventClientClass.class), objc.Sel("possibleEventDescriptorForVirtualKeyCode:"), code)
	return objectivec.Object{ID: rv}
}
func (_SLSRemoteViewEventClientClass SLSRemoteViewEventClientClass) SharedInstance() SLSRemoteViewEventClient {
	rv := objc.Send[objc.ID](objc.ID(_SLSRemoteViewEventClientClass.class), objc.Sel("sharedInstance"))
	return SLSRemoteViewEventClientFromID(rv)
}

func (s SLSRemoteViewEventClient) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (s SLSRemoteViewEventClient) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s.ID, objc.Sel("delegate"))
	return rv
}
func (s SLSRemoteViewEventClient) SetDelegate(value kernel.Pointer) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}
func (s SLSRemoteViewEventClient) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (s SLSRemoteViewEventClient) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}
func (s SLSRemoteViewEventClient) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](s.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

// ActivateWithHandlerInvalidationHandlerSync is a synchronous wrapper around [SLSRemoteViewEventClient.ActivateWithHandlerInvalidationHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLSRemoteViewEventClient) ActivateWithHandlerInvalidationHandlerSync(ctx context.Context, handler VoidHandler) error {
	done := make(chan struct{}, 1)
	s.ActivateWithHandlerInvalidationHandler(handler, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SendEventToHostFullDispatchReplySync is a synchronous wrapper around [SLSRemoteViewEventClient.SendEventToHostFullDispatchReply].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLSRemoteViewEventClient) SendEventToHostFullDispatchReplySync(ctx context.Context, host objectivec.IObject, dispatch objectivec.IObject) error {
	done := make(chan struct{}, 1)
	s.SendEventToHostFullDispatchReply(host, dispatch, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ServicePassEventUpstreamToHostFullDispatchReplySync is a synchronous wrapper around [SLSRemoteViewEventClient.ServicePassEventUpstreamToHostFullDispatchReply].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLSRemoteViewEventClient) ServicePassEventUpstreamToHostFullDispatchReplySync(ctx context.Context, host unsafe.Pointer, dispatch bool) error {
	done := make(chan struct{}, 1)
	s.ServicePassEventUpstreamToHostFullDispatchReply(host, dispatch, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
