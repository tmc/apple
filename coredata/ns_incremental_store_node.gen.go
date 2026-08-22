// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSIncrementalStoreNode] class.
var (
	_NSIncrementalStoreNodeClass     NSIncrementalStoreNodeClass
	_NSIncrementalStoreNodeClassOnce sync.Once
)

func getNSIncrementalStoreNodeClass() NSIncrementalStoreNodeClass {
	_NSIncrementalStoreNodeClassOnce.Do(func() {
		_NSIncrementalStoreNodeClass = NSIncrementalStoreNodeClass{class: objc.GetClass("NSIncrementalStoreNode")}
	})
	return _NSIncrementalStoreNodeClass
}

// GetNSIncrementalStoreNodeClass returns the class object for NSIncrementalStoreNode.
func GetNSIncrementalStoreNodeClass() NSIncrementalStoreNodeClass {
	return getNSIncrementalStoreNodeClass()
}

type NSIncrementalStoreNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSIncrementalStoreNodeClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSIncrementalStoreNodeClass) Alloc() NSIncrementalStoreNode {
	rv := objc.Send[NSIncrementalStoreNode](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A concrete class used to represent basic nodes in a Core Data incremental
// store.
//
// # Overview
//
// A node represents a single record in a persistent store.
//
// You can subclass [NSIncrementalStoreNode] to provide custom behavior.
//
// # Initializing a Node
//
//   - [NSIncrementalStoreNode.InitWithObjectIDWithValuesVersion]: Returns an object initialized with the given values.
//
// # Managing Node Data
//
//   - [NSIncrementalStoreNode.ObjectID]: The object ID that identifies the data stored by the receiver.
//   - [NSIncrementalStoreNode.UpdateWithValuesVersion]: Update the values and version to reflect new data being saved to or loaded from the external store.
//   - [NSIncrementalStoreNode.ValueForPropertyDescription]: Returns the value for the given property.
//   - [NSIncrementalStoreNode.Version]: The version of data in the receiver.
//
// See: https://developer.apple.com/documentation/CoreData/NSIncrementalStoreNode
type NSIncrementalStoreNode struct {
	objectivec.Object
}

// NSIncrementalStoreNodeFromID constructs a [NSIncrementalStoreNode] from an objc.ID.
//
// A concrete class used to represent basic nodes in a Core Data incremental
// store.
func NSIncrementalStoreNodeFromID(id objc.ID) NSIncrementalStoreNode {
	return NSIncrementalStoreNode{objectivec.Object{ID: id}}
}

// NOTE: NSIncrementalStoreNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSIncrementalStoreNode] class.
//
// # Initializing a Node
//
//   - [INSIncrementalStoreNode.InitWithObjectIDWithValuesVersion]: Returns an object initialized with the given values.
//
// # Managing Node Data
//
//   - [INSIncrementalStoreNode.ObjectID]: The object ID that identifies the data stored by the receiver.
//   - [INSIncrementalStoreNode.UpdateWithValuesVersion]: Update the values and version to reflect new data being saved to or loaded from the external store.
//   - [INSIncrementalStoreNode.ValueForPropertyDescription]: Returns the value for the given property.
//   - [INSIncrementalStoreNode.Version]: The version of data in the receiver.
//
// See: https://developer.apple.com/documentation/CoreData/NSIncrementalStoreNode
type INSIncrementalStoreNode interface {
	objectivec.IObject

	// Topic: Initializing a Node

	// Returns an object initialized with the given values.
	InitWithObjectIDWithValuesVersion(objectID INSManagedObjectID, values foundation.INSDictionary, version uint64) NSIncrementalStoreNode

	// Topic: Managing Node Data

	// The object ID that identifies the data stored by the receiver.
	ObjectID() INSManagedObjectID
	// Update the values and version to reflect new data being saved to or loaded from the external store.
	UpdateWithValuesVersion(values foundation.INSDictionary, version uint64)
	// Returns the value for the given property.
	ValueForPropertyDescription(prop INSPropertyDescription) objectivec.IObject
	// The version of data in the receiver.
	Version() uint64
}

// Init initializes the instance.
func (i NSIncrementalStoreNode) Init() NSIncrementalStoreNode {
	rv := objc.Send[NSIncrementalStoreNode](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i NSIncrementalStoreNode) Autorelease() NSIncrementalStoreNode {
	rv := objc.Send[NSIncrementalStoreNode](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSIncrementalStoreNode creates a new NSIncrementalStoreNode instance.
func NewNSIncrementalStoreNode() NSIncrementalStoreNode {
	class := getNSIncrementalStoreNodeClass()
	rv := objc.Send[NSIncrementalStoreNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an object initialized with the given values.
//
// objectID: A managed object ID.
//
// values: A dictionary containing the values persisted in an external store with keys
// corresponding to the names of the property description in the
// [NSEntityDescription] object described by `objectID`:
//
// - For attributes: an immutable value (an instance of a value class such as
// [NSNumber], [NSString], [NSData]). Missing attribute keys will assume a nil
// value. - For to-one relationships: the managed object ID of the related
// object or an instance of [NSNull] for nil relationship values. A missing
// key will be resolved lazily through calling “ on the [NSPersistentStore]
// object. Lazy resolution for to-one relationships is discouraged. - For
// to-many relationships: an instance of [NSArray] or [NSSet] containing the
// managed object IDs of the related objects. Empty to-many relationships must
// be represented by an empty non-nil collection. A missing key will be
// resolved lazily through calling “ on the [NSPersistentStore] object. Lazy
// resolution for to-many relationships is encouraged.
//
// Unknown or unmodeled keys are stripped out.
//
// version: The revision number of this state. This value is used for conflict
// detection and merging.
//
// # Return Value
//
// An object initialized with the given values.
//
// See: https://developer.apple.com/documentation/CoreData/NSIncrementalStoreNode/init(objectID:withValues:version:)
func NewIncrementalStoreNodeWithObjectIDWithValuesVersion(objectID INSManagedObjectID, values foundation.INSDictionary, version uint64) NSIncrementalStoreNode {
	instance := getNSIncrementalStoreNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObjectID:withValues:version:"), objectID, values, version)
	return NSIncrementalStoreNodeFromID(rv)
}

// Returns an object initialized with the given values.
//
// objectID: A managed object ID.
//
// values: A dictionary containing the values persisted in an external store with keys
// corresponding to the names of the property description in the
// [NSEntityDescription] object described by `objectID`:
//
// - For attributes: an immutable value (an instance of a value class such as
// [NSNumber], [NSString], [NSData]). Missing attribute keys will assume a nil
// value. - For to-one relationships: the managed object ID of the related
// object or an instance of [NSNull] for nil relationship values. A missing
// key will be resolved lazily through calling “ on the [NSPersistentStore]
// object. Lazy resolution for to-one relationships is discouraged. - For
// to-many relationships: an instance of [NSArray] or [NSSet] containing the
// managed object IDs of the related objects. Empty to-many relationships must
// be represented by an empty non-nil collection. A missing key will be
// resolved lazily through calling “ on the [NSPersistentStore] object. Lazy
// resolution for to-many relationships is encouraged.
//
// Unknown or unmodeled keys are stripped out.
//
// version: The revision number of this state. This value is used for conflict
// detection and merging.
//
// # Return Value
//
// An object initialized with the given values.
//
// See: https://developer.apple.com/documentation/CoreData/NSIncrementalStoreNode/init(objectID:withValues:version:)
func (i NSIncrementalStoreNode) InitWithObjectIDWithValuesVersion(objectID INSManagedObjectID, values foundation.INSDictionary, version uint64) NSIncrementalStoreNode {
	rv := objc.Send[NSIncrementalStoreNode](i.ID, objc.Sel("initWithObjectID:withValues:version:"), objectID, values, version)
	return rv
}

// Update the values and version to reflect new data being saved to or loaded
// from the external store.
//
// values: A dictionary containing updated values, in the same format as that
// described in [NSIncrementalStoreNode.InitWithObjectIDWithValuesVersion].
//
// version: The version number for the transaction.
//
// # Discussion
//
// Update the values and version to reflect new data being saved to or loaded
// from the external store. // The values dictionary is in the same format as
// the initializer
//
// See: https://developer.apple.com/documentation/CoreData/NSIncrementalStoreNode/update(withValues:version:)
func (i NSIncrementalStoreNode) UpdateWithValuesVersion(values foundation.INSDictionary, version uint64) {
	objc.Send[objc.ID](i.ID, objc.Sel("updateWithValues:version:"), values, version)
}

// Returns the value for the given property.
//
// prop: A property description for one of the properties in the receiver.
//
// # Return Value
//
// The value for the property specified by `prop`. May return an instance of
// [NSNull] for to-one relationships.
//
// # Discussion
//
// If a relationship is `nil`, you should create a new value by invoking “ on
// the [NSPersistentStore] object.
//
// See: https://developer.apple.com/documentation/CoreData/NSIncrementalStoreNode/value(for:)
func (i NSIncrementalStoreNode) ValueForPropertyDescription(prop INSPropertyDescription) objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("valueForPropertyDescription:"), prop)
	return objectivec.Object{ID: rv}
}

// The object ID that identifies the data stored by the receiver.
//
// See: https://developer.apple.com/documentation/CoreData/NSIncrementalStoreNode/objectID
func (i NSIncrementalStoreNode) ObjectID() INSManagedObjectID {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("objectID"))
	return NSManagedObjectIDFromID(objc.ID(rv))
}

// The version of data in the receiver.
//
// # Discussion
//
// The version number is used by the persistent store coordinator to detect
// and handle merge conflicts. The version number should be stored with the
// record. The version number should (implicitly) start at zero (where zero
// indicates an unsaved object in memory) and be incremented by exactly one
// every time you save. In addition, you increment the version number when you
// or the Core Data framework have marked the associated managed object for
// optimistic locking.
//
// See: https://developer.apple.com/documentation/CoreData/NSIncrementalStoreNode/version
func (i NSIncrementalStoreNode) Version() uint64 {
	rv := objc.Send[uint64](i.ID, objc.Sel("version"))
	return rv
}
