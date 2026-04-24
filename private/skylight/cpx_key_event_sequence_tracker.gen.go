// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXKeyEventSequenceTracker] class.
var (
	_CPXKeyEventSequenceTrackerClass     CPXKeyEventSequenceTrackerClass
	_CPXKeyEventSequenceTrackerClassOnce sync.Once
)

func getCPXKeyEventSequenceTrackerClass() CPXKeyEventSequenceTrackerClass {
	_CPXKeyEventSequenceTrackerClassOnce.Do(func() {
		_CPXKeyEventSequenceTrackerClass = CPXKeyEventSequenceTrackerClass{class: objc.GetClass("CPXKeyEventSequenceTracker")}
	})
	return _CPXKeyEventSequenceTrackerClass
}

// GetCPXKeyEventSequenceTrackerClass returns the class object for CPXKeyEventSequenceTracker.
func GetCPXKeyEventSequenceTrackerClass() CPXKeyEventSequenceTrackerClass {
	return getCPXKeyEventSequenceTrackerClass()
}

type CPXKeyEventSequenceTrackerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXKeyEventSequenceTrackerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXKeyEventSequenceTrackerClass) Alloc() CPXKeyEventSequenceTracker {
	rv := objc.Send[CPXKeyEventSequenceTracker](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXKeyEventSequenceTracker.Count]
//   - [CPXKeyEventSequenceTracker.DestinationForEventExtras]
//   - [CPXKeyEventSequenceTracker.NoteKeyEventProcessedDestination]
//   - [CPXKeyEventSequenceTracker.InitWithProvider]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXKeyEventSequenceTracker
type CPXKeyEventSequenceTracker struct {
	objectivec.Object
}

// CPXKeyEventSequenceTrackerFromID constructs a [CPXKeyEventSequenceTracker] from an objc.ID.
func CPXKeyEventSequenceTrackerFromID(id objc.ID) CPXKeyEventSequenceTracker {
	return CPXKeyEventSequenceTracker{objectivec.Object{ID: id}}
}

// Ensure CPXKeyEventSequenceTracker implements ICPXKeyEventSequenceTracker.
var _ ICPXKeyEventSequenceTracker = CPXKeyEventSequenceTracker{}

// An interface definition for the [CPXKeyEventSequenceTracker] class.
//
// # Methods
//
//   - [ICPXKeyEventSequenceTracker.Count]
//   - [ICPXKeyEventSequenceTracker.DestinationForEventExtras]
//   - [ICPXKeyEventSequenceTracker.NoteKeyEventProcessedDestination]
//   - [ICPXKeyEventSequenceTracker.InitWithProvider]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXKeyEventSequenceTracker
type ICPXKeyEventSequenceTracker interface {
	objectivec.IObject

	// Topic: Methods

	Count() uint64
	DestinationForEventExtras(event *SLSEventRecordRef, extras []objectivec.IObject) objectivec.IObject
	NoteKeyEventProcessedDestination(processed *SLSEventRecordRef, destination objectivec.IObject) bool
	InitWithProvider(provider objectivec.IObject) CPXKeyEventSequenceTracker
}

// Init initializes the instance.
func (c CPXKeyEventSequenceTracker) Init() CPXKeyEventSequenceTracker {
	rv := objc.Send[CPXKeyEventSequenceTracker](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXKeyEventSequenceTracker) Autorelease() CPXKeyEventSequenceTracker {
	rv := objc.Send[CPXKeyEventSequenceTracker](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXKeyEventSequenceTracker creates a new CPXKeyEventSequenceTracker instance.
func NewCPXKeyEventSequenceTracker() CPXKeyEventSequenceTracker {
	class := getCPXKeyEventSequenceTrackerClass()
	rv := objc.Send[CPXKeyEventSequenceTracker](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyEventSequenceTracker/initWithProvider:
func NewCPXKeyEventSequenceTrackerWithProvider(provider objectivec.IObject) CPXKeyEventSequenceTracker {
	instance := getCPXKeyEventSequenceTrackerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProvider:"), provider)
	return CPXKeyEventSequenceTrackerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyEventSequenceTracker/count
func (c CPXKeyEventSequenceTracker) Count() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("count"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyEventSequenceTracker/destinationForEvent:extras:
func (c CPXKeyEventSequenceTracker) DestinationForEventExtras(event *SLSEventRecordRef, extras []objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("destinationForEvent:extras:"), event, objectivec.IObjectSliceToNSArray(extras))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyEventSequenceTracker/noteKeyEventProcessed:destination:
func (c CPXKeyEventSequenceTracker) NoteKeyEventProcessedDestination(processed *SLSEventRecordRef, destination objectivec.IObject) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("noteKeyEventProcessed:destination:"), processed, destination)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyEventSequenceTracker/initWithProvider:
func (c CPXKeyEventSequenceTracker) InitWithProvider(provider objectivec.IObject) CPXKeyEventSequenceTracker {
	rv := objc.Send[CPXKeyEventSequenceTracker](c.ID, objc.Sel("initWithProvider:"), provider)
	return rv
}
