// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [NSFetchedPropertyDescription] class.
var (
	_NSFetchedPropertyDescriptionClass     NSFetchedPropertyDescriptionClass
	_NSFetchedPropertyDescriptionClassOnce sync.Once
)

func getNSFetchedPropertyDescriptionClass() NSFetchedPropertyDescriptionClass {
	_NSFetchedPropertyDescriptionClassOnce.Do(func() {
		_NSFetchedPropertyDescriptionClass = NSFetchedPropertyDescriptionClass{class: objc.GetClass("NSFetchedPropertyDescription")}
	})
	return _NSFetchedPropertyDescriptionClass
}

// GetNSFetchedPropertyDescriptionClass returns the class object for NSFetchedPropertyDescription.
func GetNSFetchedPropertyDescriptionClass() NSFetchedPropertyDescriptionClass {
	return getNSFetchedPropertyDescriptionClass()
}

type NSFetchedPropertyDescriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSFetchedPropertyDescriptionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFetchedPropertyDescriptionClass) Alloc() NSFetchedPropertyDescription {
	rv := objc.Send[NSFetchedPropertyDescription](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A description object used to define which properties are fetched from Core
// Data.
//
// # Overview
//
// An example might be a iTunes playlist, if expressed as a property of a
// containing object. Songs don’t belong to a particular playlist,
// especially in the case that they’re on a remote server. The playlist may
// remain even after the songs have been deleted, or the remote server has
// become inaccessible. Note, however, that unlike a playlist a fetched
// property is static—it does not dynamically update itself as objects in
// the destination entity change.
//
// The effect of a fetched property is similar to executing a fetch request
// yourself and placing the results in a transient attribute, although with
// the framework managing the details. In particular, a fetched property is
// not fetched until it is requested, and the results are then cached until
// the object is turned into a fault. You use
// [NSManagedObjectContext.RefreshObjectMergeChanges]
// ([NSManagedObjectContext]) to manually refresh the properties—this causes
// the fetch request associated with this property to be executed again when
// the object fault is next fired.
//
// Unlike other relationships, which are all sets, fetched properties are
// represented by an ordered [NSArray] object just as if you executed the
// fetch request yourself. The fetch request associated with the property can
// have a sort ordering. The value for a fetched property of a managed object
// does not support “.
//
// # Fetch Request Variables
//
// Fetch requests set on an fetched property have 2 special variable bindings
// you can use: `$FETCH_SOURCE` and `$FETCHED_PROPERTY`. The source refers to
// the specific managed object that has this property; the property refers to
// the [NSFetchedPropertyDescription] object itself (which may have a user
// info associated with it that you want to use).
//
// # Editing Fetched Property Descriptions
//
// Fetched Property descriptions are editable until they are used by an object
// graph manager. This allows you to create or modify them dynamically.
// However, once a description is used (when the managed object model to which
// it belongs is associated with a persistent store coordinator), it must not
// (indeed cannot) be changed. This is enforced at runtime: any attempt to
// mutate a model or any of its subjects after the model is associated with a
// persistent store coordinator causes an exception to be thrown. If you need
// to modify a model that is in use, create a copy, modify the copy, and then
// discard the objects with the old model.
//
// # Getting and setting the fetch request
//
//   - [NSFetchedPropertyDescription.FetchRequest]: The fetch request of the receiver.
//   - [NSFetchedPropertyDescription.SetFetchRequest]
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedPropertyDescription
type NSFetchedPropertyDescription struct {
	NSPropertyDescription
}

// NSFetchedPropertyDescriptionFromID constructs a [NSFetchedPropertyDescription] from an objc.ID.
//
// A description object used to define which properties are fetched from Core
// Data.
func NSFetchedPropertyDescriptionFromID(id objc.ID) NSFetchedPropertyDescription {
	return NSFetchedPropertyDescription{NSPropertyDescription: NSPropertyDescriptionFromID(id)}
}

// NOTE: NSFetchedPropertyDescription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFetchedPropertyDescription] class.
//
// # Getting and setting the fetch request
//
//   - [INSFetchedPropertyDescription.FetchRequest]: The fetch request of the receiver.
//   - [INSFetchedPropertyDescription.SetFetchRequest]
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedPropertyDescription
type INSFetchedPropertyDescription interface {
	INSPropertyDescription

	// Topic: Getting and setting the fetch request

	// The fetch request of the receiver.
	FetchRequest() INSFetchRequest
	SetFetchRequest(value INSFetchRequest)
}

// Init initializes the instance.
func (f NSFetchedPropertyDescription) Init() NSFetchedPropertyDescription {
	rv := objc.Send[NSFetchedPropertyDescription](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFetchedPropertyDescription) Autorelease() NSFetchedPropertyDescription {
	rv := objc.Send[NSFetchedPropertyDescription](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFetchedPropertyDescription creates a new NSFetchedPropertyDescription instance.
func NewNSFetchedPropertyDescription() NSFetchedPropertyDescription {
	class := getNSFetchedPropertyDescriptionClass()
	rv := objc.Send[NSFetchedPropertyDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The fetch request of the receiver.
//
// # Discussion
//
// Setting the fetch request raises an exception if the receiver’s model has
// been used by an object graph manager.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedPropertyDescription/fetchRequest
func (f NSFetchedPropertyDescription) FetchRequest() INSFetchRequest {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fetchRequest"))
	return NSFetchRequestFromID(objc.ID(rv))
}
func (f NSFetchedPropertyDescription) SetFetchRequest(value INSFetchRequest) {
	objc.Send[struct{}](f.ID, objc.Sel("setFetchRequest:"), value)
}
