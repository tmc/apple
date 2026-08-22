// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSManagedObjectID] class.
var (
	_NSManagedObjectIDClass     NSManagedObjectIDClass
	_NSManagedObjectIDClassOnce sync.Once
)

func getNSManagedObjectIDClass() NSManagedObjectIDClass {
	_NSManagedObjectIDClassOnce.Do(func() {
		_NSManagedObjectIDClass = NSManagedObjectIDClass{class: objc.GetClass("NSManagedObjectID")}
	})
	return _NSManagedObjectIDClass
}

// GetNSManagedObjectIDClass returns the class object for NSManagedObjectID.
func GetNSManagedObjectIDClass() NSManagedObjectIDClass {
	return getNSManagedObjectIDClass()
}

type NSManagedObjectIDClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSManagedObjectIDClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSManagedObjectIDClass) Alloc() NSManagedObjectID {
	rv := objc.Send[NSManagedObjectID](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A compact, universal identifier for a managed object.
//
// # Overview
//
// This identifier forms the basis for uniquing in the Core Data Framework. A
// managed object ID uniquely identifies the same managed object both between
// managed object contexts in a single application, and in multiple
// applications (as in distributed systems). Identifiers contain the
// information needed to exactly describe an object in a persistent store
// (like the primary key in the database), although the detailed information
// is not exposed. The framework completely encapsulates the “external”
// information and presents a clean object oriented interface.
//
// Object IDs can be transformed into a URI representation which can be
// archived and recreated later to refer back to a given object (using
// [NSPersistentStoreCoordinator.ManagedObjectIDForURIRepresentation]
// ([NSPersistentStoreCoordinator]) and [NSManagedObjectContext.ObjectWithID]
// ([NSManagedObjectContext]). For example, the last selected group in an
// application could be stored in the user defaults through the group
// object’s ID. You can also use object ID URI representations to store
// “weak” relationships across persistent stores (where no hard join is
// possible).
//
// # Getting Managed Object ID Information
//
//   - [NSManagedObjectID.Entity]: The entity description associated with the object ID.
//   - [NSManagedObjectID.IsTemporaryID]: A Boolean value that indicates whether the object ID is temporary.
//   - [NSManagedObjectID.PersistentStore]: The persistent store that fetched the object for the object ID.
//   - [NSManagedObjectID.URIRepresentation]: Returns a URI that provides an archiveable reference to the object for the object ID.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectID
type NSManagedObjectID struct {
	objectivec.Object
}

// NSManagedObjectIDFromID constructs a [NSManagedObjectID] from an objc.ID.
//
// A compact, universal identifier for a managed object.
func NSManagedObjectIDFromID(id objc.ID) NSManagedObjectID {
	return NSManagedObjectID{objectivec.Object{ID: id}}
}

// NOTE: NSManagedObjectID adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSManagedObjectID] class.
//
// # Getting Managed Object ID Information
//
//   - [INSManagedObjectID.Entity]: The entity description associated with the object ID.
//   - [INSManagedObjectID.IsTemporaryID]: A Boolean value that indicates whether the object ID is temporary.
//   - [INSManagedObjectID.PersistentStore]: The persistent store that fetched the object for the object ID.
//   - [INSManagedObjectID.URIRepresentation]: Returns a URI that provides an archiveable reference to the object for the object ID.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectID
type INSManagedObjectID interface {
	objectivec.IObject
	NSFetchRequestResult

	// Topic: Getting Managed Object ID Information

	// The entity description associated with the object ID.
	Entity() INSEntityDescription
	// A Boolean value that indicates whether the object ID is temporary.
	IsTemporaryID() bool
	// The persistent store that fetched the object for the object ID.
	PersistentStore() INSPersistentStore
	// Returns a URI that provides an archiveable reference to the object for the object ID.
	URIRepresentation() foundation.NSURL
}

// Init initializes the instance.
func (m NSManagedObjectID) Init() NSManagedObjectID {
	rv := objc.Send[NSManagedObjectID](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSManagedObjectID) Autorelease() NSManagedObjectID {
	rv := objc.Send[NSManagedObjectID](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSManagedObjectID creates a new NSManagedObjectID instance.
func NewNSManagedObjectID() NSManagedObjectID {
	class := getNSManagedObjectIDClass()
	rv := objc.Send[NSManagedObjectID](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a URI that provides an archiveable reference to the object for the
// object ID.
//
// # Return Value
//
// An [NSURL] object containing a URI that provides an archiveable reference
// to the object which the receiver represents.
//
// # Discussion
//
// If the corresponding managed object has not yet been saved, the object ID
// (and hence URI) is a temporary value that will change when the
// corresponding managed object is saved.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectID/uriRepresentation()
func (m NSManagedObjectID) URIRepresentation() foundation.NSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("URIRepresentation"))
	return foundation.NSURLFromID(rv)
}

// The entity description associated with the object ID.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectID/entity
func (m NSManagedObjectID) Entity() INSEntityDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("entity"))
	return NSEntityDescriptionFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the object ID is temporary.
//
// # Discussion
//
// true if the receiver is temporary, otherwise false. Most object IDs return
// false. New objects inserted into a managed object context are assigned a
// temporary ID which is replaced with a permanent one once the object gets
// saved to a persistent store.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectID/isTemporaryID
func (m NSManagedObjectID) IsTemporaryID() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isTemporaryID"))
	return rv
}

// The persistent store that fetched the object for the object ID.
//
// # Discussion
//
// `nil` if the ID is for a newly-inserted object that has not yet been saved
// to a persistent store.
//
// See: https://developer.apple.com/documentation/CoreData/NSManagedObjectID/persistentStore
func (m NSManagedObjectID) PersistentStore() INSPersistentStore {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("persistentStore"))
	return NSPersistentStoreFromID(objc.ID(rv))
}

// Protocol methods for NSFetchRequestResult
