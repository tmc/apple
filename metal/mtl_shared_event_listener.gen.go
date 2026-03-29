// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLSharedEventListener] class.
var (
	_MTLSharedEventListenerClass     MTLSharedEventListenerClass
	_MTLSharedEventListenerClassOnce sync.Once
)

func getMTLSharedEventListenerClass() MTLSharedEventListenerClass {
	_MTLSharedEventListenerClassOnce.Do(func() {
		_MTLSharedEventListenerClass = MTLSharedEventListenerClass{class: objc.GetClass("MTLSharedEventListener")}
	})
	return _MTLSharedEventListenerClass
}

// GetMTLSharedEventListenerClass returns the class object for MTLSharedEventListener.
func GetMTLSharedEventListenerClass() MTLSharedEventListenerClass {
	return getMTLSharedEventListenerClass()
}

type MTLSharedEventListenerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLSharedEventListenerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLSharedEventListenerClass) Alloc() MTLSharedEventListener {
	rv := objc.Send[MTLSharedEventListener](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A listener for shareable event notifications.
//
// # Initializing a shareable event listener
//
//   - [MTLSharedEventListener.InitWithDispatchQueue]: Creates a new shareable event listener with a specific dispatch queue.
//
// # Getting the dispatch queue
//
//   - [MTLSharedEventListener.DispatchQueue]: The dispatch queue used to dispatch any notifications.
//
// See: https://developer.apple.com/documentation/Metal/MTLSharedEventListener
type MTLSharedEventListener struct {
	objectivec.Object
}

// MTLSharedEventListenerFromID constructs a [MTLSharedEventListener] from an objc.ID.
//
// A listener for shareable event notifications.
func MTLSharedEventListenerFromID(id objc.ID) MTLSharedEventListener {
	return MTLSharedEventListener{objectivec.Object{ID: id}}
}
// NOTE: MTLSharedEventListener adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLSharedEventListener] class.
//
// # Initializing a shareable event listener
//
//   - [IMTLSharedEventListener.InitWithDispatchQueue]: Creates a new shareable event listener with a specific dispatch queue.
//
// # Getting the dispatch queue
//
//   - [IMTLSharedEventListener.DispatchQueue]: The dispatch queue used to dispatch any notifications.
//
// See: https://developer.apple.com/documentation/Metal/MTLSharedEventListener
type IMTLSharedEventListener interface {
	objectivec.IObject

	// Topic: Initializing a shareable event listener

	// Creates a new shareable event listener with a specific dispatch queue.
	InitWithDispatchQueue(dispatchQueue dispatch.Queue) MTLSharedEventListener

	// Topic: Getting the dispatch queue

	// The dispatch queue used to dispatch any notifications.
	DispatchQueue() dispatch.Queue
}

// Init initializes the instance.
func (s MTLSharedEventListener) Init() MTLSharedEventListener {
	rv := objc.Send[MTLSharedEventListener](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MTLSharedEventListener) Autorelease() MTLSharedEventListener {
	rv := objc.Send[MTLSharedEventListener](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLSharedEventListener creates a new MTLSharedEventListener instance.
func NewMTLSharedEventListener() MTLSharedEventListener {
	class := getMTLSharedEventListenerClass()
	rv := objc.Send[MTLSharedEventListener](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new shareable event listener with a specific dispatch queue.
//
// See: https://developer.apple.com/documentation/Metal/MTLSharedEventListener/init(dispatchQueue:)
func NewSharedEventListenerWithDispatchQueue(dispatchQueue dispatch.Queue) MTLSharedEventListener {
	instance := getMTLSharedEventListenerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDispatchQueue:"), uintptr(dispatchQueue.Handle()))
	return MTLSharedEventListenerFromID(rv)
}

// Creates a new shareable event listener with a specific dispatch queue.
//
// See: https://developer.apple.com/documentation/Metal/MTLSharedEventListener/init(dispatchQueue:)
func (s MTLSharedEventListener) InitWithDispatchQueue(dispatchQueue dispatch.Queue) MTLSharedEventListener {
	rv := objc.Send[MTLSharedEventListener](s.ID, objc.Sel("initWithDispatchQueue:"), uintptr(dispatchQueue.Handle()))
	return rv
}

// See: https://developer.apple.com/documentation/Metal/MTLSharedEventListener/shared()
func (_MTLSharedEventListenerClass MTLSharedEventListenerClass) SharedListener() MTLSharedEventListener {
	rv := objc.Send[objc.ID](objc.ID(_MTLSharedEventListenerClass.class), objc.Sel("sharedListener"))
	return MTLSharedEventListenerFromID(rv)
}

// The dispatch queue used to dispatch any notifications.
//
// See: https://developer.apple.com/documentation/Metal/MTLSharedEventListener/dispatchQueue
func (s MTLSharedEventListener) DispatchQueue() dispatch.Queue {
	rv := objc.Send[uintptr](s.ID, objc.Sel("dispatchQueue"))
	return dispatch.QueueFromHandle(rv)
}

