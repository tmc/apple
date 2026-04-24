// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXEventDispatcher] class.
var (
	_CPXEventDispatcherClass     CPXEventDispatcherClass
	_CPXEventDispatcherClassOnce sync.Once
)

func getCPXEventDispatcherClass() CPXEventDispatcherClass {
	_CPXEventDispatcherClassOnce.Do(func() {
		_CPXEventDispatcherClass = CPXEventDispatcherClass{class: objc.GetClass("CPXEventDispatcher")}
	})
	return _CPXEventDispatcherClass
}

// GetCPXEventDispatcherClass returns the class object for CPXEventDispatcher.
func GetCPXEventDispatcherClass() CPXEventDispatcherClass {
	return getCPXEventDispatcherClass()
}

type CPXEventDispatcherClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXEventDispatcherClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXEventDispatcherClass) Alloc() CPXEventDispatcher {
	rv := objc.Send[CPXEventDispatcher](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXEventDispatcher.PostBackgroundEvent]
//   - [CPXEventDispatcher.PostEventToConnectionID]
//   - [CPXEventDispatcher.PostEventToDestination]
//   - [CPXEventDispatcher.DebugDescription]
//   - [CPXEventDispatcher.Description]
//   - [CPXEventDispatcher.Hash]
//   - [CPXEventDispatcher.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXEventDispatcher
type CPXEventDispatcher struct {
	objectivec.Object
}

// CPXEventDispatcherFromID constructs a [CPXEventDispatcher] from an objc.ID.
func CPXEventDispatcherFromID(id objc.ID) CPXEventDispatcher {
	return CPXEventDispatcher{objectivec.Object{ID: id}}
}

// Ensure CPXEventDispatcher implements ICPXEventDispatcher.
var _ ICPXEventDispatcher = CPXEventDispatcher{}

// An interface definition for the [CPXEventDispatcher] class.
//
// # Methods
//
//   - [ICPXEventDispatcher.PostBackgroundEvent]
//   - [ICPXEventDispatcher.PostEventToConnectionID]
//   - [ICPXEventDispatcher.PostEventToDestination]
//   - [ICPXEventDispatcher.DebugDescription]
//   - [ICPXEventDispatcher.Description]
//   - [ICPXEventDispatcher.Hash]
//   - [ICPXEventDispatcher.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXEventDispatcher
type ICPXEventDispatcher interface {
	objectivec.IObject

	// Topic: Methods

	PostBackgroundEvent(event *SLSEventRecordRef)
	PostEventToConnectionID(event *SLSEventRecordRef, id uint32)
	PostEventToDestination(event *SLSEventRecordRef, destination objectivec.IObject)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c CPXEventDispatcher) Init() CPXEventDispatcher {
	rv := objc.Send[CPXEventDispatcher](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXEventDispatcher) Autorelease() CPXEventDispatcher {
	rv := objc.Send[CPXEventDispatcher](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXEventDispatcher creates a new CPXEventDispatcher instance.
func NewCPXEventDispatcher() CPXEventDispatcher {
	class := getCPXEventDispatcherClass()
	rv := objc.Send[CPXEventDispatcher](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDispatcher/postBackgroundEvent:
func (c CPXEventDispatcher) PostBackgroundEvent(event *SLSEventRecordRef) {
	objc.Send[objc.ID](c.ID, objc.Sel("postBackgroundEvent:"), event)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDispatcher/postEvent:toConnectionID:
func (c CPXEventDispatcher) PostEventToConnectionID(event *SLSEventRecordRef, id uint32) {
	objc.Send[objc.ID](c.ID, objc.Sel("postEvent:toConnectionID:"), event, id)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDispatcher/postEvent:toDestination:
func (c CPXEventDispatcher) PostEventToDestination(event *SLSEventRecordRef, destination objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("postEvent:toDestination:"), event, destination)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDispatcher/debugDescription
func (c CPXEventDispatcher) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDispatcher/description
func (c CPXEventDispatcher) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDispatcher/hash
func (c CPXEventDispatcher) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXEventDispatcher/superclass
func (c CPXEventDispatcher) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}
