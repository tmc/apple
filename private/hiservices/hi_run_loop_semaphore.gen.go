// Code generated from Apple documentation for hiservices. DO NOT EDIT.

package hiservices

import (
	"context"
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [HIRunLoopSemaphore] class.
var (
	_HIRunLoopSemaphoreClass     HIRunLoopSemaphoreClass
	_HIRunLoopSemaphoreClassOnce sync.Once
)

func getHIRunLoopSemaphoreClass() HIRunLoopSemaphoreClass {
	_HIRunLoopSemaphoreClassOnce.Do(func() {
		_HIRunLoopSemaphoreClass = HIRunLoopSemaphoreClass{class: objc.GetClass("HIRunLoopSemaphore")}
	})
	return _HIRunLoopSemaphoreClass
}

// GetHIRunLoopSemaphoreClass returns the class object for HIRunLoopSemaphore.
func GetHIRunLoopSemaphoreClass() HIRunLoopSemaphoreClass {
	return getHIRunLoopSemaphoreClass()
}

type HIRunLoopSemaphoreClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (hc HIRunLoopSemaphoreClass) Class() objc.Class {
	return hc.class
}

// Alloc allocates memory for a new instance of the class.
func (hc HIRunLoopSemaphoreClass) Alloc() HIRunLoopSemaphore {
	rv := objc.Send[HIRunLoopSemaphore](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [HIRunLoopSemaphore.InvokeLoopInModeForDurationWithBlock]
//   - [HIRunLoopSemaphore.Legend]
//   - [HIRunLoopSemaphore.SetLegend]
//   - [HIRunLoopSemaphore.Mode]
//   - [HIRunLoopSemaphore.SetLegacyWake]
//   - [HIRunLoopSemaphore.Signal]
//   - [HIRunLoopSemaphore.Wait]
//   - [HIRunLoopSemaphore.WaitWithWait]
//   - [HIRunLoopSemaphore.InitWithMode]
//
// See: https://developer.apple.com/documentation/HIServices/HIRunLoopSemaphore
type HIRunLoopSemaphore struct {
	objectivec.Object
}

// HIRunLoopSemaphoreFromID constructs a [HIRunLoopSemaphore] from an objc.ID.
func HIRunLoopSemaphoreFromID(id objc.ID) HIRunLoopSemaphore {
	return HIRunLoopSemaphore{objectivec.Object{ID: id}}
}

// Ensure HIRunLoopSemaphore implements IHIRunLoopSemaphore.
var _ IHIRunLoopSemaphore = HIRunLoopSemaphore{}

// An interface definition for the [HIRunLoopSemaphore] class.
//
// # Methods
//
//   - [IHIRunLoopSemaphore.InvokeLoopInModeForDurationWithBlock]
//   - [IHIRunLoopSemaphore.Legend]
//   - [IHIRunLoopSemaphore.SetLegend]
//   - [IHIRunLoopSemaphore.Mode]
//   - [IHIRunLoopSemaphore.SetLegacyWake]
//   - [IHIRunLoopSemaphore.Signal]
//   - [IHIRunLoopSemaphore.Wait]
//   - [IHIRunLoopSemaphore.WaitWithWait]
//   - [IHIRunLoopSemaphore.InitWithMode]
//
// See: https://developer.apple.com/documentation/HIServices/HIRunLoopSemaphore
type IHIRunLoopSemaphore interface {
	objectivec.IObject

	// Topic: Methods

	InvokeLoopInModeForDurationWithBlock(duration float64, block VoidHandler)
	Legend() string
	SetLegend(value string)
	Mode() corefoundation.CFStringRef
	SetLegacyWake()
	Signal()
	Wait()
	WaitWithWait(wait float64) bool
	InitWithMode(mode corefoundation.CFStringRef) HIRunLoopSemaphore
}

// Init initializes the instance.
func (h HIRunLoopSemaphore) Init() HIRunLoopSemaphore {
	rv := objc.Send[HIRunLoopSemaphore](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h HIRunLoopSemaphore) Autorelease() HIRunLoopSemaphore {
	rv := objc.Send[HIRunLoopSemaphore](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewHIRunLoopSemaphore creates a new HIRunLoopSemaphore instance.
func NewHIRunLoopSemaphore() HIRunLoopSemaphore {
	class := getHIRunLoopSemaphoreClass()
	rv := objc.Send[HIRunLoopSemaphore](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/HIServices/HIRunLoopSemaphore/initWithMode:
func NewHIRunLoopSemaphoreWithMode(mode corefoundation.CFStringRef) HIRunLoopSemaphore {
	instance := getHIRunLoopSemaphoreClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMode:"), mode)
	return HIRunLoopSemaphoreFromID(rv)
}

// See: https://developer.apple.com/documentation/HIServices/HIRunLoopSemaphore/invokeLoopInModeForDuration:withBlock:
func (h HIRunLoopSemaphore) InvokeLoopInModeForDurationWithBlock(duration float64, block VoidHandler) {
	_block1, _ := NewVoidBlock(block)
	objc.Send[objc.ID](h.ID, objc.Sel("invokeLoopInModeForDuration:withBlock:"), duration, _block1)
}

// See: https://developer.apple.com/documentation/HIServices/HIRunLoopSemaphore/setLegacyWake
func (h HIRunLoopSemaphore) SetLegacyWake() {
	objc.Send[objc.ID](h.ID, objc.Sel("setLegacyWake"))
}

// See: https://developer.apple.com/documentation/HIServices/HIRunLoopSemaphore/signal
func (h HIRunLoopSemaphore) Signal() {
	objc.Send[objc.ID](h.ID, objc.Sel("signal"))
}

// See: https://developer.apple.com/documentation/HIServices/HIRunLoopSemaphore/wait
func (h HIRunLoopSemaphore) Wait() {
	objc.Send[objc.ID](h.ID, objc.Sel("wait"))
}

// See: https://developer.apple.com/documentation/HIServices/HIRunLoopSemaphore/wait:
func (h HIRunLoopSemaphore) WaitWithWait(wait float64) bool {
	rv := objc.Send[bool](h.ID, objc.Sel("wait:"), wait)
	return rv
}

// See: https://developer.apple.com/documentation/HIServices/HIRunLoopSemaphore/initWithMode:
func (h HIRunLoopSemaphore) InitWithMode(mode corefoundation.CFStringRef) HIRunLoopSemaphore {
	rv := objc.Send[HIRunLoopSemaphore](h.ID, objc.Sel("initWithMode:"), mode)
	return rv
}

// See: https://developer.apple.com/documentation/HIServices/HIRunLoopSemaphore/_invocations
func (_HIRunLoopSemaphoreClass HIRunLoopSemaphoreClass) _invocations() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_HIRunLoopSemaphoreClass.class), objc.Sel("_invocations"))
	return objectivec.Object{ID: rv}
}

// Invocations is an exported wrapper for the private method _invocations.
func (_HIRunLoopSemaphoreClass HIRunLoopSemaphoreClass) Invocations() objectivec.IObject {
	return _HIRunLoopSemaphoreClass._invocations()
}

// See: https://developer.apple.com/documentation/HIServices/HIRunLoopSemaphore/_observe:whilePerforming:
func (_HIRunLoopSemaphoreClass HIRunLoopSemaphoreClass) _observeWhilePerforming(_observe corefoundation.CFStringRef, performing VoidHandler) {
	_block1, _ := NewVoidBlock(performing)
	objc.Send[objc.ID](objc.ID(_HIRunLoopSemaphoreClass.class), objc.Sel("_observe:whilePerforming:"), _observe, _block1)
}

// ObserveWhilePerforming is an exported wrapper for the private method _observeWhilePerforming.
func (_HIRunLoopSemaphoreClass HIRunLoopSemaphoreClass) ObserveWhilePerforming(_observe corefoundation.CFStringRef, performing VoidHandler) {
	_HIRunLoopSemaphoreClass._observeWhilePerforming(_observe, performing)
}

// See: https://developer.apple.com/documentation/HIServices/HIRunLoopSemaphore/legend
func (h HIRunLoopSemaphore) Legend() string {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("legend"))
	return foundation.NSStringFromID(rv).String()
}
func (h HIRunLoopSemaphore) SetLegend(value string) {
	objc.Send[struct{}](h.ID, objc.Sel("setLegend:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/HIServices/HIRunLoopSemaphore/mode
func (h HIRunLoopSemaphore) Mode() corefoundation.CFStringRef {
	rv := objc.Send[corefoundation.CFStringRef](h.ID, objc.Sel("mode"))
	return corefoundation.CFStringRef(rv)
}

// InvokeLoopInModeForDurationWithBlockSync is a synchronous wrapper around [HIRunLoopSemaphore.InvokeLoopInModeForDurationWithBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (h HIRunLoopSemaphore) InvokeLoopInModeForDurationWithBlockSync(ctx context.Context, duration float64) error {
	done := make(chan struct{}, 1)
	h.InvokeLoopInModeForDurationWithBlock(duration, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// _observeWhilePerformingSync is a synchronous wrapper around [HIRunLoopSemaphore._observeWhilePerforming].
// It blocks until the completion handler fires or the context is cancelled.
func (hc HIRunLoopSemaphoreClass) _observeWhilePerformingSync(ctx context.Context, _observe corefoundation.CFStringRef) error {
	done := make(chan struct{}, 1)
	hc._observeWhilePerforming(_observe, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
