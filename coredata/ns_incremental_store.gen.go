// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSIncrementalStore] class.
var (
	_NSIncrementalStoreClass     NSIncrementalStoreClass
	_NSIncrementalStoreClassOnce sync.Once
)

func getNSIncrementalStoreClass() NSIncrementalStoreClass {
	_NSIncrementalStoreClassOnce.Do(func() {
		_NSIncrementalStoreClass = NSIncrementalStoreClass{class: objc.GetClass("NSIncrementalStore")}
	})
	return _NSIncrementalStoreClass
}

// GetNSIncrementalStoreClass returns the class object for NSIncrementalStore.
func GetNSIncrementalStoreClass() NSIncrementalStoreClass {
	return getNSIncrementalStoreClass()
}

type NSIncrementalStoreClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSIncrementalStoreClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSIncrementalStoreClass) Alloc() NSIncrementalStore {
	rv := objc.Send[NSIncrementalStore](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract superclass defining the API through which Core Data
// communicates with a store.
//
// # Overview
//
// You use this interface to create persistent stores that load and save data
// incrementally, allowing for the management of large and/or shared datasets.
//
// # Subclassing Notes
//
// # Methods to Override
//
// In a subclass of [NSIncrementalStore], you must override the following
// methods to provide behavior appropriate for your store:
//
// - [NSIncrementalStore.LoadMetadata] -
// [NSIncrementalStore.ExecuteRequestWithContextError] -
// [NSIncrementalStore.NewValuesForObjectWithIDWithContextError] -
// [NSIncrementalStore.NewValueForRelationshipForObjectWithIDWithContextError]
// - [NSIncrementalStore.ObtainPermanentIDsForObjectsError]
//
// You can also optionally override the following methods:
//
// - [NSIncrementalStoreClass.IdentifierForNewStoreAtURL] -
// [NSIncrementalStore.ManagedObjectContextDidRegisterObjectsWithIDs] -
// [NSIncrementalStore.ManagedObjectContextDidUnregisterObjectsWithIDs]
//
// There is no need to override the methods that you must otherwise override
// for a subclass of [NSPersistentStore].
//
// # Methods that Should Not Be Overridden
//
// In a subclass of [NSIncrementalStore], you should not override the
// following methods:
//
// - [NSIncrementalStore.NewObjectIDForEntityReferenceObject] -
// [NSIncrementalStore.ReferenceObjectForObjectID]
//
// # Manipulating Managed Objects
//
//   - [NSIncrementalStore.ExecuteRequestWithContextError]: Returns a value as appropriate for the given request, or nil if the request cannot be completed.
//   - [NSIncrementalStore.NewValuesForObjectWithIDWithContextError]: Returns an incremental store node encapsulating the persistent external values of the object with a given object ID.
//   - [NSIncrementalStore.NewValueForRelationshipForObjectWithIDWithContextError]: Returns the relationship for the given relationship of the object with a given object ID.
//   - [NSIncrementalStore.ObtainPermanentIDsForObjectsError]: Returns an array containing the object IDs for a given array of newly-inserted objects.
//   - [NSIncrementalStore.NewObjectIDForEntityReferenceObject]: Returns a new object ID that uses given data as the key.
//   - [NSIncrementalStore.ReferenceObjectForObjectID]: Returns the reference data used to construct a given object ID.
//
// # Responding to Context Changes
//
//   - [NSIncrementalStore.ManagedObjectContextDidRegisterObjectsWithIDs]: Indicates that objects identified by a given array of object IDs are in use in a managed object context.
//   - [NSIncrementalStore.ManagedObjectContextDidUnregisterObjectsWithIDs]: Indicates that objects identified by a given array of object IDs are no longer being used by a managed object context.
//
// See: https://developer.apple.com/documentation/CoreData/NSIncrementalStore
type NSIncrementalStore struct {
	NSPersistentStore
}

// NSIncrementalStoreFromID constructs a [NSIncrementalStore] from an objc.ID.
//
// An abstract superclass defining the API through which Core Data
// communicates with a store.
func NSIncrementalStoreFromID(id objc.ID) NSIncrementalStore {
	return NSIncrementalStore{NSPersistentStore: NSPersistentStoreFromID(id)}
}

// NOTE: NSIncrementalStore adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSIncrementalStore] class.
//
// # Manipulating Managed Objects
//
//   - [INSIncrementalStore.ExecuteRequestWithContextError]: Returns a value as appropriate for the given request, or nil if the request cannot be completed.
//   - [INSIncrementalStore.NewValuesForObjectWithIDWithContextError]: Returns an incremental store node encapsulating the persistent external values of the object with a given object ID.
//   - [INSIncrementalStore.NewValueForRelationshipForObjectWithIDWithContextError]: Returns the relationship for the given relationship of the object with a given object ID.
//   - [INSIncrementalStore.ObtainPermanentIDsForObjectsError]: Returns an array containing the object IDs for a given array of newly-inserted objects.
//   - [INSIncrementalStore.NewObjectIDForEntityReferenceObject]: Returns a new object ID that uses given data as the key.
//   - [INSIncrementalStore.ReferenceObjectForObjectID]: Returns the reference data used to construct a given object ID.
//
// # Responding to Context Changes
//
//   - [INSIncrementalStore.ManagedObjectContextDidRegisterObjectsWithIDs]: Indicates that objects identified by a given array of object IDs are in use in a managed object context.
//   - [INSIncrementalStore.ManagedObjectContextDidUnregisterObjectsWithIDs]: Indicates that objects identified by a given array of object IDs are no longer being used by a managed object context.
//
// See: https://developer.apple.com/documentation/CoreData/NSIncrementalStore
type INSIncrementalStore interface {
	INSPersistentStore

	// Topic: Manipulating Managed Objects

	// Returns a value as appropriate for the given request, or nil if the request cannot be completed.
	ExecuteRequestWithContextError(request INSPersistentStoreRequest, context INSManagedObjectContext) (objectivec.IObject, error)
	// Returns an incremental store node encapsulating the persistent external values of the object with a given object ID.
	NewValuesForObjectWithIDWithContextError(objectID INSManagedObjectID, context INSManagedObjectContext) (INSIncrementalStoreNode, error)
	// Returns the relationship for the given relationship of the object with a given object ID.
	NewValueForRelationshipForObjectWithIDWithContextError(relationship INSRelationshipDescription, objectID INSManagedObjectID, context INSManagedObjectContext) (objectivec.IObject, error)
	// Returns an array containing the object IDs for a given array of newly-inserted objects.
	ObtainPermanentIDsForObjectsError(array []NSManagedObject) ([]NSManagedObjectID, error)
	// Returns a new object ID that uses given data as the key.
	NewObjectIDForEntityReferenceObject(entity INSEntityDescription, data objectivec.IObject) INSManagedObjectID
	// Returns the reference data used to construct a given object ID.
	ReferenceObjectForObjectID(objectID INSManagedObjectID) objectivec.IObject

	// Topic: Responding to Context Changes

	// Indicates that objects identified by a given array of object IDs are in use in a managed object context.
	ManagedObjectContextDidRegisterObjectsWithIDs(objectIDs []NSManagedObjectID)
	// Indicates that objects identified by a given array of object IDs are no longer being used by a managed object context.
	ManagedObjectContextDidUnregisterObjectsWithIDs(objectIDs []NSManagedObjectID)
}

// Init initializes the instance.
func (i NSIncrementalStore) Init() NSIncrementalStore {
	rv := objc.Send[NSIncrementalStore](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i NSIncrementalStore) Autorelease() NSIncrementalStore {
	rv := objc.Send[NSIncrementalStore](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSIncrementalStore creates a new NSIncrementalStore instance.
func NewNSIncrementalStore() NSIncrementalStore {
	class := getNSIncrementalStoreClass()
	rv := objc.Send[NSIncrementalStore](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a store initialized with the given arguments.
//
// root: A persistent store coordinator.
//
// name: The name of the managed object model configuration to use. Pass `nil` if
// you do not want to specify a configuration.
//
// url: The URL of the store to load.
//
// options: A dictionary containing configuration options. See
// [NSPersistentStoreCoordinator] for a list of key names for options in this
// dictionary.
//
// # Return Value
//
// A new store object, associated with `coordinator`, that represents a
// persistent store at url using the options in `options` and—if it is not
// `nil`—the managed object model configuration `configurationName`.
//
// # Discussion
//
// You must ensure that you load metadata during initialization and set it
// using [NSPersistentStore.Metadata].
//
// # Special Considerations
//
// This is the designated initializer for persistent stores.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStore/init(persistentStoreCoordinator:configurationName:at:options:)
func NewIncrementalStoreWithPersistentStoreCoordinatorConfigurationNameURLOptions(root INSPersistentStoreCoordinator, name string, url foundation.NSURL, options foundation.INSDictionary) NSIncrementalStore {
	instance := getNSIncrementalStoreClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPersistentStoreCoordinator:configurationName:URL:options:"), root, objc.String(name), url, options)
	return NSIncrementalStoreFromID(rv)
}

// Returns a value as appropriate for the given request, or nil if the request
// cannot be completed.
//
// request: A fetch request.
//
// context: The managed object context used to execute `request`.
//
// # Return Value
//
// A value as appropriate for `request`, or `nil` if the request cannot be
// completed
//
// # Discussion
//
// The value to return depends on the result type (see
// [NSFetchRequest.ResultType]) of `request`:
//
// - If it is [NSManagedObjectResultType], [NSManagedObjectIDResultType], or
// [NSDictionaryResultType], the method should return an array containing all
// objects in the store matching the request. - If it is [NSCountResultType],
// the method should return an array containing an [NSNumber] whose value is
// the count of all objects in the store matching the request. - If the
// request is a save request, the method should return an empty array.
//
// If the save request contains nil values for the
// inserted/updated/deleted/locked collections; you should treat it as a
// request to save the store metadata.
//
// You should implement this method conservatively, and expect that unknown
// request types may at some point be passed to the method. The correct
// behavior in these cases is to return `nil` and an error.
//
// See: https://developer.apple.com/documentation/CoreData/NSIncrementalStore/execute(_:with:)
func (i NSIncrementalStore) ExecuteRequestWithContextError(request INSPersistentStoreRequest, context INSManagedObjectContext) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](i.ID, objc.Sel("executeRequest:withContext:error:"), request, context, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// Returns an incremental store node encapsulating the persistent external
// values of the object with a given object ID.
//
// objectID: The ID of the object for which values are requested.
//
// context: The managed object context into which values will be returned.
//
// # Return Value
//
// An incremental store node encapsulating the persistent external values of
// the object with object ID `objectID`, or `nil` if the corresponding object
// cannot be found.
//
// # Discussion
//
// The returned node should include all attributes values and may include
// to-one relationship values as instances of [NSManagedObjectID].
//
// If an object with object ID `objectID` cannot be found, the method should
// return `nil` and—if `error` is not [NULL]—create and return an
// appropriate error object in `error`.
//
// See: https://developer.apple.com/documentation/CoreData/NSIncrementalStore/newValuesForObject(with:with:)
func (i NSIncrementalStore) NewValuesForObjectWithIDWithContextError(objectID INSManagedObjectID, context INSManagedObjectContext) (INSIncrementalStoreNode, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](i.ID, objc.Sel("newValuesForObjectWithID:withContext:error:"), objectID, context, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSIncrementalStoreNode{}, foundation.NSErrorFrom(errorPtr)
	}
	return NSIncrementalStoreNodeFromID(rv), nil

}

// Returns the relationship for the given relationship of the object with a
// given object ID.
//
// relationship: The relationship for which values are requested.
//
// objectID: The ID of the object for which values are requested.
//
// context: The managed object context into which values will be returned.
//
// # Return Value
//
// The value of the relationship specified `relationship` of the object with
// object ID `objectID`, or `nil` if an error occurs.
//
// # Discussion
//
// If the relationship is a to-one, the method should return an
// [NSManagedObjectID] instance that identifies the destination, or an
// instance of [NSNull] if the relationship value is `nil`.
//
// If the relationship is a to-many, the method should return a collection
// object containing [NSManagedObjectID] instances to identify the related
// objects. Using an [NSArray] instance is preferred because it will be the
// most efficient. A store may also return an instance of [NSSet] or
// [NSOrderedSet]; an instance of [NSDictionary] is not acceptable.
//
// If an object with object ID `objectID` cannot be found, the method should
// return `nil` and—if `error` is not [NULL]—create and return an
// appropriate error object in `error`.
//
// See: https://developer.apple.com/documentation/CoreData/NSIncrementalStore/newValue(forRelationship:forObjectWith:with:)
//
// [NSNull]: https://developer.apple.com/documentation/Foundation/NSNull
func (i NSIncrementalStore) NewValueForRelationshipForObjectWithIDWithContextError(relationship INSRelationshipDescription, objectID INSManagedObjectID, context INSManagedObjectContext) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](i.ID, objc.Sel("newValueForRelationship:forObjectWithID:withContext:error:"), relationship, objectID, context, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// Returns an array containing the object IDs for a given array of
// newly-inserted objects.
//
// array: An array of newly-inserted objects.
//
// # Return Value
//
// An array containing the object IDs for the objects in `array`.
//
// # Discussion
//
// The returned array must return the object IDs in the same order as the
// objects appear in `array`.
//
// # Discussion
//
// This method is called before
// [NSIncrementalStore.ExecuteRequestWithContextError] with a save request, to
// assign permanent IDs to newly-inserted objects.
//
// See: https://developer.apple.com/documentation/CoreData/NSIncrementalStore/obtainPermanentIDs(for:)
func (i NSIncrementalStore) ObtainPermanentIDsForObjectsError(array []NSManagedObject) ([]NSManagedObjectID, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](i.ID, objc.Sel("obtainPermanentIDsForObjects:error:"), objectivec.IObjectSliceToNSArray(array), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objc.ConvertSlice(rv, func(id objc.ID) NSManagedObjectID {
		return NSManagedObjectIDFromID(id)
	}), nil

}

// Returns a new object ID that uses given data as the key.
//
// entity: The entity for the new object ID.
//
// data: An object of type [NSString] or [NSNumber] to use as the key.
//
// # Return Value
//
// A new object ID for an instance of the entity specified by `entity` and
// that uses `data` as the key.
//
// # Discussion
//
// You should not override this method.
//
// See: https://developer.apple.com/documentation/CoreData/NSIncrementalStore/newObjectID(for:referenceObject:)
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
func (i NSIncrementalStore) NewObjectIDForEntityReferenceObject(entity INSEntityDescription, data objectivec.IObject) INSManagedObjectID {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("newObjectIDForEntity:referenceObject:"), entity, data)
	return NSManagedObjectIDFromID(rv)
}

// Returns the reference data used to construct a given object ID.
//
// objectID: An object ID created by the receiver.
//
// # Return Value
//
// The reference data used to construct objectID.
//
// # Discussion
//
// This method raises an [invalidArgumentException] if the object ID was not
// created by the receiving store.
//
// You should not override this method.
//
// See: https://developer.apple.com/documentation/CoreData/NSIncrementalStore/referenceObject(for:)
//
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
func (i NSIncrementalStore) ReferenceObjectForObjectID(objectID INSManagedObjectID) objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("referenceObjectForObjectID:"), objectID)
	return objectivec.Object{ID: rv}
}

// Indicates that objects identified by a given array of object IDs are in use
// in a managed object context.
//
// objectIDs: An array of object IDs.
//
// # Discussion
//
// This method and
// [NSIncrementalStore.ManagedObjectContextDidUnregisterObjectsWithIDs] allow
// managed object contexts to communicate interest in the row data of specific
// objects in a manner akin to reference counting. For more details, see
// [NSIncrementalStore.ManagedObjectContextDidUnregisterObjectsWithIDs].
//
// See: https://developer.apple.com/documentation/CoreData/NSIncrementalStore/managedObjectContextDidRegisterObjects(with:)
func (i NSIncrementalStore) ManagedObjectContextDidRegisterObjectsWithIDs(objectIDs []NSManagedObjectID) {
	objc.Send[objc.ID](i.ID, objc.Sel("managedObjectContextDidRegisterObjectsWithIDs:"), objectivec.IObjectSliceToNSArray(objectIDs))
}

// Indicates that objects identified by a given array of object IDs are no
// longer being used by a managed object context.
//
// objectIDs: An array of object IDs.
//
// # Discussion
//
// This method is the counterpart to
// [NSIncrementalStore.ManagedObjectContextDidRegisterObjectsWithIDs].
//
// Passing an object ID in the object IDs array of
// [NSIncrementalStore.ManagedObjectContextDidRegisterObjectsWithIDs] is akin
// to incrementing the object ID’s reference count by 1; passing an object
// ID in the object IDs array of
// [NSIncrementalStore.ManagedObjectContextDidUnregisterObjectsWithIDs] is
// akin to decrementing the object ID’s reference count by 1. It is only
// when an object ID’s reference count is 0 that no contexts indicate that
// they are using the corresponding managed object. (Object IDs start with a
// reference count of 0.)
//
// For example, if the register methods is invoked on two occasions when the
// object IDs array contains a given object ID, and the unregister method is
// invoked once when the object IDs array contains that object ID, then a
// context is still using the object with the given ID.
//
// See: https://developer.apple.com/documentation/CoreData/NSIncrementalStore/managedObjectContextDidUnregisterObjects(with:)
func (i NSIncrementalStore) ManagedObjectContextDidUnregisterObjectsWithIDs(objectIDs []NSManagedObjectID) {
	objc.Send[objc.ID](i.ID, objc.Sel("managedObjectContextDidUnregisterObjectsWithIDs:"), objectivec.IObjectSliceToNSArray(objectIDs))
}

// Returns the identifier for the store at a given URL.
//
// storeURL: The URL of a persistent store.
//
// # Return Value
//
// The identifier for the store at `storeURL`.
//
// See: https://developer.apple.com/documentation/CoreData/NSIncrementalStore/identifierForNewStore(at:)
func (_NSIncrementalStoreClass NSIncrementalStoreClass) IdentifierForNewStoreAtURL(storeURL foundation.NSURL) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NSIncrementalStoreClass.class), objc.Sel("identifierForNewStoreAtURL:"), storeURL)
	return objectivec.Object{ID: rv}
}
