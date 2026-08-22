// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSPersistentCloudKitContainerEventRequest] class.
var (
	_NSPersistentCloudKitContainerEventRequestClass     NSPersistentCloudKitContainerEventRequestClass
	_NSPersistentCloudKitContainerEventRequestClassOnce sync.Once
)

func getNSPersistentCloudKitContainerEventRequestClass() NSPersistentCloudKitContainerEventRequestClass {
	_NSPersistentCloudKitContainerEventRequestClassOnce.Do(func() {
		_NSPersistentCloudKitContainerEventRequestClass = NSPersistentCloudKitContainerEventRequestClass{class: objc.GetClass("NSPersistentCloudKitContainerEventRequest")}
	})
	return _NSPersistentCloudKitContainerEventRequestClass
}

// GetNSPersistentCloudKitContainerEventRequestClass returns the class object for NSPersistentCloudKitContainerEventRequest.
func GetNSPersistentCloudKitContainerEventRequestClass() NSPersistentCloudKitContainerEventRequestClass {
	return getNSPersistentCloudKitContainerEventRequestClass()
}

type NSPersistentCloudKitContainerEventRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPersistentCloudKitContainerEventRequestClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPersistentCloudKitContainerEventRequestClass) Alloc() NSPersistentCloudKitContainerEventRequest {
	rv := objc.Send[NSPersistentCloudKitContainerEventRequest](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A request to fetch setup, import, or export events in a persistent CloudKit
// container.
//
// # Fetching Events
//
//   - [NSPersistentCloudKitContainerEventRequest.ResultType]: The type of result that the request returns.
//   - [NSPersistentCloudKitContainerEventRequest.SetResultType]
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventRequest
type NSPersistentCloudKitContainerEventRequest struct {
	NSPersistentStoreRequest
}

// NSPersistentCloudKitContainerEventRequestFromID constructs a [NSPersistentCloudKitContainerEventRequest] from an objc.ID.
//
// A request to fetch setup, import, or export events in a persistent CloudKit
// container.
func NSPersistentCloudKitContainerEventRequestFromID(id objc.ID) NSPersistentCloudKitContainerEventRequest {
	return NSPersistentCloudKitContainerEventRequest{NSPersistentStoreRequest: NSPersistentStoreRequestFromID(id)}
}

// NOTE: NSPersistentCloudKitContainerEventRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPersistentCloudKitContainerEventRequest] class.
//
// # Fetching Events
//
//   - [INSPersistentCloudKitContainerEventRequest.ResultType]: The type of result that the request returns.
//   - [INSPersistentCloudKitContainerEventRequest.SetResultType]
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventRequest
type INSPersistentCloudKitContainerEventRequest interface {
	INSPersistentStoreRequest

	// Topic: Fetching Events

	// The type of result that the request returns.
	ResultType() NSPersistentCloudKitContainerEventResultType
	SetResultType(value NSPersistentCloudKitContainerEventResultType)
}

// Init initializes the instance.
func (p NSPersistentCloudKitContainerEventRequest) Init() NSPersistentCloudKitContainerEventRequest {
	rv := objc.Send[NSPersistentCloudKitContainerEventRequest](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPersistentCloudKitContainerEventRequest) Autorelease() NSPersistentCloudKitContainerEventRequest {
	rv := objc.Send[NSPersistentCloudKitContainerEventRequest](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPersistentCloudKitContainerEventRequest creates a new NSPersistentCloudKitContainerEventRequest instance.
func NewNSPersistentCloudKitContainerEventRequest() NSPersistentCloudKitContainerEventRequest {
	class := getNSPersistentCloudKitContainerEventRequestClass()
	rv := objc.Send[NSPersistentCloudKitContainerEventRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a fetch request for events after a specified date from a persistent
// CloudKit container.
//
// date: The earliest date to return events for.
//
// # Return Value
//
// A request object that fetches persistent CloudKit container events by
// executing in a managed object context.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventRequest/fetchEvents(after:)-5izg7
func (_NSPersistentCloudKitContainerEventRequestClass NSPersistentCloudKitContainerEventRequestClass) FetchEventsAfterDate(date foundation.NSDate) NSPersistentCloudKitContainerEventRequest {
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentCloudKitContainerEventRequestClass.class), objc.Sel("fetchEventsAfterDate:"), date)
	return NSPersistentCloudKitContainerEventRequestFromID(rv)
}

// Creates a fetch request for events that occur after a specified event from
// a persistent CloudKit container.
//
// event: An event that precedes other events.
//
// # Return Value
//
// A request object that fetches persistent CloudKit container events by
// executing in a managed object context.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventRequest/fetchEvents(after:)-3yfp
func (_NSPersistentCloudKitContainerEventRequestClass NSPersistentCloudKitContainerEventRequestClass) FetchEventsAfterEvent(event INSPersistentCloudKitContainerEvent) NSPersistentCloudKitContainerEventRequest {
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentCloudKitContainerEventRequestClass.class), objc.Sel("fetchEventsAfterEvent:"), event)
	return NSPersistentCloudKitContainerEventRequestFromID(rv)
}

// Creates a fetch request for events that match a specified fetch request
// from a persistent CloudKit container.
//
// fetchRequest: A fetch request to identify matching events.
//
// # Return Value
//
// A request object that fetches persistent CloudKit container events by
// executing in a managed object context.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventRequest/fetchEvents(matchingFetch:)
func (_NSPersistentCloudKitContainerEventRequestClass NSPersistentCloudKitContainerEventRequestClass) FetchEventsMatchingFetchRequest(fetchRequest INSFetchRequest) NSPersistentCloudKitContainerEventRequest {
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentCloudKitContainerEventRequestClass.class), objc.Sel("fetchEventsMatchingFetchRequest:"), fetchRequest)
	return NSPersistentCloudKitContainerEventRequestFromID(rv)
}

// Creates a fetch request for all events in a persistent CloudKit container.
//
// # Return Value
//
// A request object that fetches persistent CloudKit container events by
// executing in a managed object context.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventRequest/fetchForEvents()
func (_NSPersistentCloudKitContainerEventRequestClass NSPersistentCloudKitContainerEventRequestClass) FetchRequestForEvents() NSFetchRequest {
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentCloudKitContainerEventRequestClass.class), objc.Sel("fetchRequestForEvents"))
	return NSFetchRequestFromID(rv)
}

// The type of result that the request returns.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventRequest/resultType
func (p NSPersistentCloudKitContainerEventRequest) ResultType() NSPersistentCloudKitContainerEventResultType {
	rv := objc.Send[NSPersistentCloudKitContainerEventResultType](p.ID, objc.Sel("resultType"))
	return NSPersistentCloudKitContainerEventResultType(rv)
}
func (p NSPersistentCloudKitContainerEventRequest) SetResultType(value NSPersistentCloudKitContainerEventResultType) {
	objc.Send[struct{}](p.ID, objc.Sel("setResultType:"), value)
}
