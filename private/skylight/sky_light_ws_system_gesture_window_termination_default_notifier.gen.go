// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightWSSystemGestureWindowTerminationDefaultNotifier] class.
var (
	_SkyLightWSSystemGestureWindowTerminationDefaultNotifierClass     SkyLightWSSystemGestureWindowTerminationDefaultNotifierClass
	_SkyLightWSSystemGestureWindowTerminationDefaultNotifierClassOnce sync.Once
)

func getSkyLightWSSystemGestureWindowTerminationDefaultNotifierClass() SkyLightWSSystemGestureWindowTerminationDefaultNotifierClass {
	_SkyLightWSSystemGestureWindowTerminationDefaultNotifierClassOnce.Do(func() {
		_SkyLightWSSystemGestureWindowTerminationDefaultNotifierClass = SkyLightWSSystemGestureWindowTerminationDefaultNotifierClass{class: objc.GetClass("SkyLight.WSSystemGestureWindowTerminationDefaultNotifier")}
	})
	return _SkyLightWSSystemGestureWindowTerminationDefaultNotifierClass
}

// GetSkyLightWSSystemGestureWindowTerminationDefaultNotifierClass returns the class object for SkyLight.WSSystemGestureWindowTerminationDefaultNotifier.
func GetSkyLightWSSystemGestureWindowTerminationDefaultNotifierClass() SkyLightWSSystemGestureWindowTerminationDefaultNotifierClass {
	return getSkyLightWSSystemGestureWindowTerminationDefaultNotifierClass()
}

type SkyLightWSSystemGestureWindowTerminationDefaultNotifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightWSSystemGestureWindowTerminationDefaultNotifierClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightWSSystemGestureWindowTerminationDefaultNotifierClass) Alloc() SkyLightWSSystemGestureWindowTerminationDefaultNotifier {
	rv := objc.SendIfResponds[SkyLightWSSystemGestureWindowTerminationDefaultNotifier](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SkyLightWSSystemGestureWindowTerminationDefaultNotifier.SubscribeForWindowTermination]
type SkyLightWSSystemGestureWindowTerminationDefaultNotifier struct {
	objectivec.Object
}

// SkyLightWSSystemGestureWindowTerminationDefaultNotifierFromID constructs a [SkyLightWSSystemGestureWindowTerminationDefaultNotifier] from an objc.ID.
func SkyLightWSSystemGestureWindowTerminationDefaultNotifierFromID(id objc.ID) SkyLightWSSystemGestureWindowTerminationDefaultNotifier {
	return SkyLightWSSystemGestureWindowTerminationDefaultNotifier{objectivec.Object{ID: id}}
}

// Ensure SkyLightWSSystemGestureWindowTerminationDefaultNotifier implements ISkyLightWSSystemGestureWindowTerminationDefaultNotifier.
var _ ISkyLightWSSystemGestureWindowTerminationDefaultNotifier = SkyLightWSSystemGestureWindowTerminationDefaultNotifier{}

// An interface definition for the [SkyLightWSSystemGestureWindowTerminationDefaultNotifier] class.
//
// # Methods
//
//   - [ISkyLightWSSystemGestureWindowTerminationDefaultNotifier.SubscribeForWindowTermination]
type ISkyLightWSSystemGestureWindowTerminationDefaultNotifier interface {
	objectivec.IObject

	// Topic: Methods

	SubscribeForWindowTermination(termination VoidHandler) objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightWSSystemGestureWindowTerminationDefaultNotifier) Init() SkyLightWSSystemGestureWindowTerminationDefaultNotifier {
	rv := objc.SendIfResponds[SkyLightWSSystemGestureWindowTerminationDefaultNotifier](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightWSSystemGestureWindowTerminationDefaultNotifier) Autorelease() SkyLightWSSystemGestureWindowTerminationDefaultNotifier {
	rv := objc.SendIfResponds[SkyLightWSSystemGestureWindowTerminationDefaultNotifier](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightWSSystemGestureWindowTerminationDefaultNotifier creates a new SkyLightWSSystemGestureWindowTerminationDefaultNotifier instance.
func NewSkyLightWSSystemGestureWindowTerminationDefaultNotifier() SkyLightWSSystemGestureWindowTerminationDefaultNotifier {
	class := getSkyLightWSSystemGestureWindowTerminationDefaultNotifierClass()
	rv := objc.SendIfResponds[SkyLightWSSystemGestureWindowTerminationDefaultNotifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (s SkyLightWSSystemGestureWindowTerminationDefaultNotifier) SubscribeForWindowTermination(termination VoidHandler) objectivec.IObject {
	_block0, _ := NewVoidBlock(termination)
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("subscribeForWindowTermination:"), _block0)
	return objectivec.Object{ID: rv}
}

// SubscribeForWindowTerminationSync is a synchronous wrapper around [SkyLightWSSystemGestureWindowTerminationDefaultNotifier.SubscribeForWindowTermination].
// It blocks until the completion handler fires or the context is cancelled.
func (s SkyLightWSSystemGestureWindowTerminationDefaultNotifier) SubscribeForWindowTerminationSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	s.SubscribeForWindowTermination(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
