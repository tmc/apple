// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPersistentCloudKitContainerEvent] class.
var (
	_NSPersistentCloudKitContainerEventClass     NSPersistentCloudKitContainerEventClass
	_NSPersistentCloudKitContainerEventClassOnce sync.Once
)

func getNSPersistentCloudKitContainerEventClass() NSPersistentCloudKitContainerEventClass {
	_NSPersistentCloudKitContainerEventClassOnce.Do(func() {
		_NSPersistentCloudKitContainerEventClass = NSPersistentCloudKitContainerEventClass{class: objc.GetClass("NSPersistentCloudKitContainerEvent")}
	})
	return _NSPersistentCloudKitContainerEventClass
}

// GetNSPersistentCloudKitContainerEventClass returns the class object for NSPersistentCloudKitContainerEvent.
func GetNSPersistentCloudKitContainerEventClass() NSPersistentCloudKitContainerEventClass {
	return getNSPersistentCloudKitContainerEventClass()
}

type NSPersistentCloudKitContainerEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPersistentCloudKitContainerEventClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPersistentCloudKitContainerEventClass) Alloc() NSPersistentCloudKitContainerEvent {
	rv := objc.Send[NSPersistentCloudKitContainerEvent](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents activity in a persistent CloudKit container.
//
// # Inspecting Event Properties
//
//   - [NSPersistentCloudKitContainerEvent.Type]: The type of event, either setup, import, or export.
//   - [NSPersistentCloudKitContainerEvent.Identifier]: A unique identifier for the event in a container.
//   - [NSPersistentCloudKitContainerEvent.StoreIdentifier]: The associated store identifier in the container for the event.
//   - [NSPersistentCloudKitContainerEvent.Succeeded]: A Boolean value that indicates whether the operation the event represents is successful.
//   - [NSPersistentCloudKitContainerEvent.StartDate]: The start date of the operation that the event represents.
//   - [NSPersistentCloudKitContainerEvent.EndDate]: The end date of the operation that the event represents.
//   - [NSPersistentCloudKitContainerEvent.Error]: An error that indicates why an operation fails.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/Event
type NSPersistentCloudKitContainerEvent struct {
	objectivec.Object
}

// NSPersistentCloudKitContainerEventFromID constructs a [NSPersistentCloudKitContainerEvent] from an objc.ID.
//
// An object that represents activity in a persistent CloudKit container.
func NSPersistentCloudKitContainerEventFromID(id objc.ID) NSPersistentCloudKitContainerEvent {
	return NSPersistentCloudKitContainerEvent{objectivec.Object{ID: id}}
}

// NOTE: NSPersistentCloudKitContainerEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPersistentCloudKitContainerEvent] class.
//
// # Inspecting Event Properties
//
//   - [INSPersistentCloudKitContainerEvent.Type]: The type of event, either setup, import, or export.
//   - [INSPersistentCloudKitContainerEvent.Identifier]: A unique identifier for the event in a container.
//   - [INSPersistentCloudKitContainerEvent.StoreIdentifier]: The associated store identifier in the container for the event.
//   - [INSPersistentCloudKitContainerEvent.Succeeded]: A Boolean value that indicates whether the operation the event represents is successful.
//   - [INSPersistentCloudKitContainerEvent.StartDate]: The start date of the operation that the event represents.
//   - [INSPersistentCloudKitContainerEvent.EndDate]: The end date of the operation that the event represents.
//   - [INSPersistentCloudKitContainerEvent.Error]: An error that indicates why an operation fails.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/Event
type INSPersistentCloudKitContainerEvent interface {
	objectivec.IObject

	// Topic: Inspecting Event Properties

	// The type of event, either setup, import, or export.
	Type() NSPersistentCloudKitContainerEventType
	// A unique identifier for the event in a container.
	Identifier() foundation.NSUUID
	// The associated store identifier in the container for the event.
	StoreIdentifier() string
	// A Boolean value that indicates whether the operation the event represents is successful.
	Succeeded() bool
	// The start date of the operation that the event represents.
	StartDate() foundation.NSDate
	// The end date of the operation that the event represents.
	EndDate() foundation.NSDate
	// An error that indicates why an operation fails.
	Error() foundation.NSError
}

// Init initializes the instance.
func (p NSPersistentCloudKitContainerEvent) Init() NSPersistentCloudKitContainerEvent {
	rv := objc.Send[NSPersistentCloudKitContainerEvent](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPersistentCloudKitContainerEvent) Autorelease() NSPersistentCloudKitContainerEvent {
	rv := objc.Send[NSPersistentCloudKitContainerEvent](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPersistentCloudKitContainerEvent creates a new NSPersistentCloudKitContainerEvent instance.
func NewNSPersistentCloudKitContainerEvent() NSPersistentCloudKitContainerEvent {
	class := getNSPersistentCloudKitContainerEventClass()
	rv := objc.Send[NSPersistentCloudKitContainerEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The type of event, either setup, import, or export.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/Event/type
func (p NSPersistentCloudKitContainerEvent) Type() NSPersistentCloudKitContainerEventType {
	rv := objc.Send[NSPersistentCloudKitContainerEventType](p.ID, objc.Sel("type"))
	return NSPersistentCloudKitContainerEventType(rv)
}

// A unique identifier for the event in a container.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/Event/identifier
func (p NSPersistentCloudKitContainerEvent) Identifier() foundation.NSUUID {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("identifier"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}

// The associated store identifier in the container for the event.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/Event/storeIdentifier
func (p NSPersistentCloudKitContainerEvent) StoreIdentifier() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("storeIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// A Boolean value that indicates whether the operation the event represents
// is successful.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/Event/succeeded
func (p NSPersistentCloudKitContainerEvent) Succeeded() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("succeeded"))
	return rv
}

// The start date of the operation that the event represents.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/Event/startDate
func (p NSPersistentCloudKitContainerEvent) StartDate() foundation.NSDate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("startDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// The end date of the operation that the event represents.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/Event/endDate
func (p NSPersistentCloudKitContainerEvent) EndDate() foundation.NSDate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("endDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// An error that indicates why an operation fails.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/Event/error
func (p NSPersistentCloudKitContainerEvent) Error() foundation.NSError {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
