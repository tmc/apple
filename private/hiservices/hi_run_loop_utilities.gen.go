// Code generated from Apple documentation for hiservices. DO NOT EDIT.

package hiservices

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [HIRunLoopUtilities] class.
var (
	_HIRunLoopUtilitiesClass     HIRunLoopUtilitiesClass
	_HIRunLoopUtilitiesClassOnce sync.Once
)

func getHIRunLoopUtilitiesClass() HIRunLoopUtilitiesClass {
	_HIRunLoopUtilitiesClassOnce.Do(func() {
		_HIRunLoopUtilitiesClass = HIRunLoopUtilitiesClass{class: objc.GetClass("HIRunLoopUtilities")}
	})
	return _HIRunLoopUtilitiesClass
}

// GetHIRunLoopUtilitiesClass returns the class object for HIRunLoopUtilities.
func GetHIRunLoopUtilitiesClass() HIRunLoopUtilitiesClass {
	return getHIRunLoopUtilitiesClass()
}

type HIRunLoopUtilitiesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (hc HIRunLoopUtilitiesClass) Class() objc.Class {
	return hc.class
}

// Alloc allocates memory for a new instance of the class.
func (hc HIRunLoopUtilitiesClass) Alloc() HIRunLoopUtilities {
	rv := objc.Send[HIRunLoopUtilities](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [HIRunLoopUtilities._blockQueueDepth]
//
// See: https://developer.apple.com/documentation/HIServices/HIRunLoopUtilities
type HIRunLoopUtilities struct {
	objectivec.Object
}

// HIRunLoopUtilitiesFromID constructs a [HIRunLoopUtilities] from an objc.ID.
func HIRunLoopUtilitiesFromID(id objc.ID) HIRunLoopUtilities {
	return HIRunLoopUtilities{objectivec.Object{ID: id}}
}

// Ensure HIRunLoopUtilities implements IHIRunLoopUtilities.
var _ IHIRunLoopUtilities = HIRunLoopUtilities{}

// An interface definition for the [HIRunLoopUtilities] class.
//
// # Methods
//
//   - [IHIRunLoopUtilities._blockQueueDepth]
//
// See: https://developer.apple.com/documentation/HIServices/HIRunLoopUtilities
type IHIRunLoopUtilities interface {
	objectivec.IObject

	// Topic: Methods

	_blockQueueDepth() uint32
}

// Init initializes the instance.
func (h HIRunLoopUtilities) Init() HIRunLoopUtilities {
	rv := objc.Send[HIRunLoopUtilities](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h HIRunLoopUtilities) Autorelease() HIRunLoopUtilities {
	rv := objc.Send[HIRunLoopUtilities](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewHIRunLoopUtilities creates a new HIRunLoopUtilities instance.
func NewHIRunLoopUtilities() HIRunLoopUtilities {
	class := getHIRunLoopUtilitiesClass()
	rv := objc.Send[HIRunLoopUtilities](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/HIServices/HIRunLoopUtilities/_blockQueueDepth
func (h HIRunLoopUtilities) _blockQueueDepth() uint32 {
	rv := objc.Send[uint32](h.ID, objc.Sel("_blockQueueDepth"))
	return rv
}

// BlockQueueDepth is an exported wrapper for the private method _blockQueueDepth.
func (h HIRunLoopUtilities) BlockQueueDepth() uint32 {
	return h._blockQueueDepth()
}

// See: https://developer.apple.com/documentation/HIServices/HIRunLoopUtilities/_currentRunLoopMode
func (_HIRunLoopUtilitiesClass HIRunLoopUtilitiesClass) _currentRunLoopMode() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_HIRunLoopUtilitiesClass.class), objc.Sel("_currentRunLoopMode"))
	return objectivec.Object{ID: rv}
}

// CurrentRunLoopMode is an exported wrapper for the private method _currentRunLoopMode.
func (_HIRunLoopUtilitiesClass HIRunLoopUtilitiesClass) CurrentRunLoopMode() objectivec.IObject {
	return _HIRunLoopUtilitiesClass._currentRunLoopMode()
}

// See: https://developer.apple.com/documentation/HIServices/HIRunLoopUtilities/addRunLoopModesForDeferredActions:
func (_HIRunLoopUtilitiesClass HIRunLoopUtilitiesClass) AddRunLoopModesForDeferredActions(actions objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_HIRunLoopUtilitiesClass.class), objc.Sel("addRunLoopModesForDeferredActions:"), actions)
}

// See: https://developer.apple.com/documentation/HIServices/HIRunLoopUtilities/deferActions:
func (_HIRunLoopUtilitiesClass HIRunLoopUtilitiesClass) DeferActions(actions VoidHandler) {
	_block0, _ := NewVoidBlock(actions)
	objc.Send[objc.ID](objc.ID(_HIRunLoopUtilitiesClass.class), objc.Sel("deferActions:"), _block0)
}

// DeferActionsSync is a synchronous wrapper around [HIRunLoopUtilities.DeferActions].
// It blocks until the completion handler fires or the context is cancelled.
func (hc HIRunLoopUtilitiesClass) DeferActionsSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	hc.DeferActions(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
