// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSBatchDeleteRequest] class.
var (
	_NSBatchDeleteRequestClass     NSBatchDeleteRequestClass
	_NSBatchDeleteRequestClassOnce sync.Once
)

func getNSBatchDeleteRequestClass() NSBatchDeleteRequestClass {
	_NSBatchDeleteRequestClassOnce.Do(func() {
		_NSBatchDeleteRequestClass = NSBatchDeleteRequestClass{class: objc.GetClass("NSBatchDeleteRequest")}
	})
	return _NSBatchDeleteRequestClass
}

// GetNSBatchDeleteRequestClass returns the class object for NSBatchDeleteRequest.
func GetNSBatchDeleteRequestClass() NSBatchDeleteRequestClass {
	return getNSBatchDeleteRequestClass()
}

type NSBatchDeleteRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSBatchDeleteRequestClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSBatchDeleteRequestClass) Alloc() NSBatchDeleteRequest {
	rv := objc.Send[NSBatchDeleteRequest](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A request that deletes objects in the SQLite persistent store without
// loading them into memory.
//
// # Overview
//
// [NSBatchDeleteRequest] — available only when using a SQLite persistent
// store — deletes managed objects at the SQL level of the persistent store.
// This request is quicker and more efficient than using a context to fetch a
// large number of objects into memory, delete them, and then save those
// deletions back to the store. You create a request using an instance of
// [NSFetchRequest] that identifies the objects to delete. Alternatively, you
// can provide an array of identifiers from specific objects of the same
// entity type; mixing entity types results in an error when you execute the
// request.
//
// [NSManagedObjectContext] doesn’t automatically merge a request’s
// deletions because they happen at the SQL level. Subsequently, you must
// remove any deleted objects from memory after the request finishes. To
// determine the objects a request deletes, configure it to return the
// [NSManagedObjectID] of each deleted object and use those identifiers to
// update your contexts, as the following example shows:
//
// Alternatively, you can use persistent history tracking to make your
// contexts aware of changes that happen at the persistent store level. For
// more information, see [Consuming relevant store changes].
//
// # Creating a Request
//
//   - [NSBatchDeleteRequest.InitWithFetchRequest]: Creates a request that deletes the results of the specified fetch request.
//   - [NSBatchDeleteRequest.InitWithObjectIDs]: Creates a request that deletes the managed objects with the specified identifiers.
//
// # Accessing the Fetch Request
//
//   - [NSBatchDeleteRequest.FetchRequest]: The fetch request that identifies the managed objects to delete.
//
// # Configuring the Result Type
//
//   - [NSBatchDeleteRequest.ResultType]: The type of result the request provides when it executes.
//   - [NSBatchDeleteRequest.SetResultType]
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchDeleteRequest
//
// [Consuming relevant store changes]: https://developer.apple.com/documentation/CoreData/consuming-relevant-store-changes
type NSBatchDeleteRequest struct {
	NSPersistentStoreRequest
}

// NSBatchDeleteRequestFromID constructs a [NSBatchDeleteRequest] from an objc.ID.
//
// A request that deletes objects in the SQLite persistent store without
// loading them into memory.
func NSBatchDeleteRequestFromID(id objc.ID) NSBatchDeleteRequest {
	return NSBatchDeleteRequest{NSPersistentStoreRequest: NSPersistentStoreRequestFromID(id)}
}

// NOTE: NSBatchDeleteRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSBatchDeleteRequest] class.
//
// # Creating a Request
//
//   - [INSBatchDeleteRequest.InitWithFetchRequest]: Creates a request that deletes the results of the specified fetch request.
//   - [INSBatchDeleteRequest.InitWithObjectIDs]: Creates a request that deletes the managed objects with the specified identifiers.
//
// # Accessing the Fetch Request
//
//   - [INSBatchDeleteRequest.FetchRequest]: The fetch request that identifies the managed objects to delete.
//
// # Configuring the Result Type
//
//   - [INSBatchDeleteRequest.ResultType]: The type of result the request provides when it executes.
//   - [INSBatchDeleteRequest.SetResultType]
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchDeleteRequest
type INSBatchDeleteRequest interface {
	INSPersistentStoreRequest

	// Topic: Creating a Request

	// Creates a request that deletes the results of the specified fetch request.
	InitWithFetchRequest(fetch INSFetchRequest) NSBatchDeleteRequest
	// Creates a request that deletes the managed objects with the specified identifiers.
	InitWithObjectIDs(objects []NSManagedObjectID) NSBatchDeleteRequest

	// Topic: Accessing the Fetch Request

	// The fetch request that identifies the managed objects to delete.
	FetchRequest() INSFetchRequest

	// Topic: Configuring the Result Type

	// The type of result the request provides when it executes.
	ResultType() NSBatchDeleteRequestResultType
	SetResultType(value NSBatchDeleteRequestResultType)
}

// Init initializes the instance.
func (b NSBatchDeleteRequest) Init() NSBatchDeleteRequest {
	rv := objc.Send[NSBatchDeleteRequest](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b NSBatchDeleteRequest) Autorelease() NSBatchDeleteRequest {
	rv := objc.Send[NSBatchDeleteRequest](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSBatchDeleteRequest creates a new NSBatchDeleteRequest instance.
func NewNSBatchDeleteRequest() NSBatchDeleteRequest {
	class := getNSBatchDeleteRequestClass()
	rv := objc.Send[NSBatchDeleteRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a request that deletes the results of the specified fetch request.
//
// fetch: The fetch request that identifies the managed objects to delete.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchDeleteRequest/init(fetchRequest:)
func NewBatchDeleteRequestWithFetchRequest(fetch INSFetchRequest) NSBatchDeleteRequest {
	instance := getNSBatchDeleteRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFetchRequest:"), fetch)
	return NSBatchDeleteRequestFromID(rv)
}

// Creates a request that deletes the managed objects with the specified
// identifiers.
//
// objects: The array that contains the identifiers of the managed objects to delete.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchDeleteRequest/init(objectIDs:)
func NewBatchDeleteRequestWithObjectIDs(objects []NSManagedObjectID) NSBatchDeleteRequest {
	instance := getNSBatchDeleteRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObjectIDs:"), objectivec.IObjectSliceToNSArray(objects))
	return NSBatchDeleteRequestFromID(rv)
}

// Creates a request that deletes the results of the specified fetch request.
//
// fetch: The fetch request that identifies the managed objects to delete.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchDeleteRequest/init(fetchRequest:)
func (b NSBatchDeleteRequest) InitWithFetchRequest(fetch INSFetchRequest) NSBatchDeleteRequest {
	rv := objc.Send[NSBatchDeleteRequest](b.ID, objc.Sel("initWithFetchRequest:"), fetch)
	return rv
}

// Creates a request that deletes the managed objects with the specified
// identifiers.
//
// objects: The array that contains the identifiers of the managed objects to delete.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchDeleteRequest/init(objectIDs:)
func (b NSBatchDeleteRequest) InitWithObjectIDs(objects []NSManagedObjectID) NSBatchDeleteRequest {
	rv := objc.Send[NSBatchDeleteRequest](b.ID, objc.Sel("initWithObjectIDs:"), objectivec.IObjectSliceToNSArray(objects))
	return rv
}

// The fetch request that identifies the managed objects to delete.
//
// # Discussion
//
// This property contains the fetch request that identifies the managed
// objects to delete. If you initialize [NSBatchDeleteRequest] with an array
// of [NSManagedObjectID], Core Data automatically generates a fetch request
// with a predicate that matches the identifiers in that array.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchDeleteRequest/fetchRequest
func (b NSBatchDeleteRequest) FetchRequest() INSFetchRequest {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("fetchRequest"))
	return NSFetchRequestFromID(objc.ID(rv))
}

// The type of result the request provides when it executes.
//
// # Discussion
//
// Set this property before you execute the request if you require a result
// type other than the default of
// [NSBatchDeleteRequestResultType.resultTypeStatusOnly].
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchDeleteRequest/resultType
//
// [NSBatchDeleteRequestResultType.resultTypeStatusOnly]: https://developer.apple.com/documentation/CoreData/NSBatchDeleteRequestResultType/resultTypeStatusOnly
func (b NSBatchDeleteRequest) ResultType() NSBatchDeleteRequestResultType {
	rv := objc.Send[NSBatchDeleteRequestResultType](b.ID, objc.Sel("resultType"))
	return NSBatchDeleteRequestResultType(rv)
}
func (b NSBatchDeleteRequest) SetResultType(value NSBatchDeleteRequestResultType) {
	objc.Send[struct{}](b.ID, objc.Sel("setResultType:"), value)
}
