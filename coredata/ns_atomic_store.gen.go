// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAtomicStore] class.
var (
	_NSAtomicStoreClass     NSAtomicStoreClass
	_NSAtomicStoreClassOnce sync.Once
)

func getNSAtomicStoreClass() NSAtomicStoreClass {
	_NSAtomicStoreClassOnce.Do(func() {
		_NSAtomicStoreClass = NSAtomicStoreClass{class: objc.GetClass("NSAtomicStore")}
	})
	return _NSAtomicStoreClass
}

// GetNSAtomicStoreClass returns the class object for NSAtomicStore.
func GetNSAtomicStoreClass() NSAtomicStoreClass {
	return getNSAtomicStoreClass()
}

type NSAtomicStoreClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSAtomicStoreClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAtomicStoreClass) Alloc() NSAtomicStore {
	rv := objc.Send[NSAtomicStore](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract superclass that you subclass to create a Core Data atomic
// store.
//
// # Overview
//
// Use an atomic store to handle data sets that can be expressed in memory.
// The atomic store API favors simplicity over performance.
//
// This class provides default implementations of some utility methods. Create
// a custom atomic store subclass when you have a custom file format that you
// want to integrate with a Core Data app. When you create a subclass,
// override the following [NSAtomicStore] methods:
//
// - [NSAtomicStore.Load] - [NSAtomicStore.NewCacheNodeForManagedObject] -
// [NSAtomicStore.NewReferenceObjectForManagedObject] - [NSAtomicStore.Save] -
// [NSAtomicStore.UpdateCacheNodeFromManagedObject]
//
// Also override the following properties and methods of [NSPersistentStore],
// from which the atomic store class inherits:
//
// - [NSPersistentStore.Type] - [NSPersistentStore.Identifier] -
// [NSPersistentStore.Metadata] -
// [NSPersistentStoreClass.MetadataForPersistentStoreWithURLError] -
// [NSPersistentStoreClass.SetMetadataForPersistentStoreWithURLError]
//
// [NSAtomicStore] provides a default dictionary of metadata. This dictionary
// contains the store type and identifier ([NSStoreTypeKey] and
// [NSStoreUUIDKey]) as well as store versioning information. Subclasses must
// ensure that the metadata is saved along with the store data.
//
// # Loading a Store
//
//   - [NSAtomicStore.Load]: Loads the cache nodes for the receiver.
//   - [NSAtomicStore.ObjectIDForEntityReferenceObject]: Returns a managed object ID from the reference data for a specified entity.
//   - [NSAtomicStore.AddCacheNodes]: Registers a set of cache nodes with the receiver.
//
// # Updating Cache Nodes
//
//   - [NSAtomicStore.NewCacheNodeForManagedObject]: Returns a new cache node for a given managed object.
//   - [NSAtomicStore.NewReferenceObjectForManagedObject]: Returns a new reference object for a given managed object.
//   - [NSAtomicStore.UpdateCacheNodeFromManagedObject]: Updates the given cache node using the values in a given managed object.
//   - [NSAtomicStore.WillRemoveCacheNodes]: Method invoked before the store removes the given collection of cache nodes.
//
// # Saving a Store
//
//   - [NSAtomicStore.Save]: Saves the cache nodes.
//
// # Utility Methods
//
//   - [NSAtomicStore.CacheNodes]: Returns the set of cache nodes registered with the receiver.
//   - [NSAtomicStore.CacheNodeForObjectID]: Returns the cache node for a given managed object ID.
//   - [NSAtomicStore.ReferenceObjectForObjectID]: Returns the reference object for a given managed object ID.
//
// See: https://developer.apple.com/documentation/CoreData/NSAtomicStore
//
// [NSStoreTypeKey]: https://developer.apple.com/documentation/CoreData/NSStoreTypeKey
// [NSStoreUUIDKey]: https://developer.apple.com/documentation/CoreData/NSStoreUUIDKey
type NSAtomicStore struct {
	NSPersistentStore
}

// NSAtomicStoreFromID constructs a [NSAtomicStore] from an objc.ID.
//
// An abstract superclass that you subclass to create a Core Data atomic
// store.
func NSAtomicStoreFromID(id objc.ID) NSAtomicStore {
	return NSAtomicStore{NSPersistentStore: NSPersistentStoreFromID(id)}
}

// NOTE: NSAtomicStore adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSAtomicStore] class.
//
// # Loading a Store
//
//   - [INSAtomicStore.Load]: Loads the cache nodes for the receiver.
//   - [INSAtomicStore.ObjectIDForEntityReferenceObject]: Returns a managed object ID from the reference data for a specified entity.
//   - [INSAtomicStore.AddCacheNodes]: Registers a set of cache nodes with the receiver.
//
// # Updating Cache Nodes
//
//   - [INSAtomicStore.NewCacheNodeForManagedObject]: Returns a new cache node for a given managed object.
//   - [INSAtomicStore.NewReferenceObjectForManagedObject]: Returns a new reference object for a given managed object.
//   - [INSAtomicStore.UpdateCacheNodeFromManagedObject]: Updates the given cache node using the values in a given managed object.
//   - [INSAtomicStore.WillRemoveCacheNodes]: Method invoked before the store removes the given collection of cache nodes.
//
// # Saving a Store
//
//   - [INSAtomicStore.Save]: Saves the cache nodes.
//
// # Utility Methods
//
//   - [INSAtomicStore.CacheNodes]: Returns the set of cache nodes registered with the receiver.
//   - [INSAtomicStore.CacheNodeForObjectID]: Returns the cache node for a given managed object ID.
//   - [INSAtomicStore.ReferenceObjectForObjectID]: Returns the reference object for a given managed object ID.
//
// See: https://developer.apple.com/documentation/CoreData/NSAtomicStore
type INSAtomicStore interface {
	INSPersistentStore

	// Topic: Loading a Store

	// Loads the cache nodes for the receiver.
	Load() (bool, error)
	// Returns a managed object ID from the reference data for a specified entity.
	ObjectIDForEntityReferenceObject(entity INSEntityDescription, data objectivec.IObject) INSManagedObjectID
	// Registers a set of cache nodes with the receiver.
	AddCacheNodes(cacheNodes foundation.INSSet)

	// Topic: Updating Cache Nodes

	// Returns a new cache node for a given managed object.
	NewCacheNodeForManagedObject(managedObject INSManagedObject) INSAtomicStoreCacheNode
	// Returns a new reference object for a given managed object.
	NewReferenceObjectForManagedObject(managedObject INSManagedObject) objectivec.IObject
	// Updates the given cache node using the values in a given managed object.
	UpdateCacheNodeFromManagedObject(node INSAtomicStoreCacheNode, managedObject INSManagedObject)
	// Method invoked before the store removes the given collection of cache nodes.
	WillRemoveCacheNodes(cacheNodes foundation.INSSet)

	// Topic: Saving a Store

	// Saves the cache nodes.
	Save() (bool, error)

	// Topic: Utility Methods

	// Returns the set of cache nodes registered with the receiver.
	CacheNodes() foundation.INSSet
	// Returns the cache node for a given managed object ID.
	CacheNodeForObjectID(objectID INSManagedObjectID) INSAtomicStoreCacheNode
	// Returns the reference object for a given managed object ID.
	ReferenceObjectForObjectID(objectID INSManagedObjectID) objectivec.IObject
}

// Init initializes the instance.
func (a NSAtomicStore) Init() NSAtomicStore {
	rv := objc.Send[NSAtomicStore](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAtomicStore) Autorelease() NSAtomicStore {
	rv := objc.Send[NSAtomicStore](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAtomicStore creates a new NSAtomicStore instance.
func NewNSAtomicStore() NSAtomicStore {
	class := getNSAtomicStoreClass()
	rv := objc.Send[NSAtomicStore](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an atomic store at the specified location.
//
// coordinator: The persistent store coordinator.
//
// configurationName: The name of the store’s configuration in the managed object model.
//
// url: The URL of the store to load. This value can’t be `nil`.
//
// options: A dictionary that contains the store’s options. For possible values, see
// [Store options].
//
// # Discussion
//
// Typically, you don’t invoke this method yourself; instead, the persistent
// store coordinator invokes the method when it creates a new store or adds an
// existing one.
//
// In your implementation, check whether a file exists at `url`. If it
// doesn’t exist, create a zero-length file at `url` so that the file exists
// before the coordinator calls the store’s [NSAtomicStore.Load] method. A
// zero-length file indicates to the system that it should create a new store
// at that location. If the coordinator removes the store without first
// calling [NSAtomicStore.Save], delete the zero-length file.
//
// It’s your responsibility to load the store’s metadata during
// initialization and set it using the
// [NSPersistentStoreClass.SetMetadataForPersistentStoreWithURLError] method.
//
// See: https://developer.apple.com/documentation/CoreData/NSAtomicStore/init(persistentStoreCoordinator:configurationName:at:options:)
//
// [Store options]: https://developer.apple.com/documentation/CoreData/store-options
func NewAtomicStoreWithPersistentStoreCoordinatorConfigurationNameURLOptions(coordinator INSPersistentStoreCoordinator, configurationName string, url foundation.NSURL, options foundation.INSDictionary) NSAtomicStore {
	instance := getNSAtomicStoreClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPersistentStoreCoordinator:configurationName:URL:options:"), coordinator, objc.String(configurationName), url, options)
	return NSAtomicStoreFromID(rv)
}

// Loads the cache nodes for the receiver.
//
// # Discussion
//
// You override this method to load the data from the URL specified in
// [NSAtomicStore.InitWithPersistentStoreCoordinatorConfigurationNameURLOptions]
// and create cache nodes for the represented objects. You must respect the
// configuration specified for the store, as well as the options.
//
// Any subclass of [NSAtomicStore] must be able to handle being initialized
// with a URL pointing to a zero-length file. This serves as an indicator that
// a new store is to be constructed at the specified location and allows you
// to securely create reservation files in known locations which can then be
// passed to Core Data to construct stores. You may choose to create
// zero-length reservation files during
// [NSAtomicStore.InitWithPersistentStoreCoordinatorConfigurationNameURLOptions]
// or [NSAtomicStore.Load]. If you do so, you must remove the reservation file
// if the store is removed from the coordinator before it is saved.
//
// You must override this method in a subclass of [NSAtomicStore].
//
// See: https://developer.apple.com/documentation/CoreData/NSAtomicStore/load()
func (a NSAtomicStore) Load() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("load:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("load: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns a managed object ID from the reference data for a specified entity.
//
// entity: An entity description object.
//
// data: Reference data for which the managed object ID is required.
//
// # Return Value
//
// # The managed object ID from the reference data for a specified entity
//
// # Discussion
//
// You use this method to create managed object IDs which are then used to
// create cache nodes for information being loaded into the store.
//
// # Special Considerations
//
// You should not override this method.
//
// See: https://developer.apple.com/documentation/CoreData/NSAtomicStore/objectID(for:withReferenceObject:)
func (a NSAtomicStore) ObjectIDForEntityReferenceObject(entity INSEntityDescription, data objectivec.IObject) INSManagedObjectID {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("objectIDForEntity:referenceObject:"), entity, data)
	return NSManagedObjectIDFromID(rv)
}

// Registers a set of cache nodes with the receiver.
//
// cacheNodes: A set of cache nodes.
//
// # Discussion
//
// You should invoke this method in a subclass during the call to
// [NSAtomicStore.Load] to register the loaded information with the store.
//
// See: https://developer.apple.com/documentation/CoreData/NSAtomicStore/addCacheNodes(_:)
func (a NSAtomicStore) AddCacheNodes(cacheNodes foundation.INSSet) {
	objc.Send[objc.ID](a.ID, objc.Sel("addCacheNodes:"), cacheNodes)
}

// Returns a new cache node for a given managed object.
//
// managedObject: A managed object.
//
// # Return Value
//
// A new cache node for `managedObject`.
//
// # Discussion
//
// This method is invoked by the framework during a save operation, once for
// each newly-inserted managed object. It should pull information from the
// managed object and return a cache node containing the information (the node
// will be registered by the framework).
//
// # Special Considerations
//
// You must override this method.
//
// See: https://developer.apple.com/documentation/CoreData/NSAtomicStore/newCacheNode(for:)
func (a NSAtomicStore) NewCacheNodeForManagedObject(managedObject INSManagedObject) INSAtomicStoreCacheNode {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("newCacheNodeForManagedObject:"), managedObject)
	return NSAtomicStoreCacheNodeFromID(rv)
}

// Returns a new reference object for a given managed object.
//
// managedObject: A managed object. At the time this method is called, it has a temporary ID.
//
// # Return Value
//
// A new reference object for `managedObject`.
//
// # Discussion
//
// This method is invoked by the framework after a save operation on a managed
// object context, once for each newly-inserted managed object. The value
// returned is used to create a permanent ID for the object and must be unique
// for an instance within its entity’s inheritance hierarchy (in this
// store).
//
// # Special Considerations
//
// You must override this method.
//
// This method must return a stable (unchanging) value for a given object,
// otherwise Save As and migration will not work correctly. This means that
// you can use arbitrary numbers, UUIDs, or other random values only if they
// are persisted with the raw data. If you cannot save the originally-assigned
// reference object with the data, then the method must derive the reference
// object from the managed object’s values. For more details, see [Atomic
// Store Programming Topics].
//
// See: https://developer.apple.com/documentation/CoreData/NSAtomicStore/newReferenceObject(for:)
//
// [Atomic Store Programming Topics]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/AtomicStore_Concepts/Introduction/Introduction.html#//apple_ref/doc/uid/TP40004521
func (a NSAtomicStore) NewReferenceObjectForManagedObject(managedObject INSManagedObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("newReferenceObjectForManagedObject:"), managedObject)
	return objectivec.Object{ID: rv}
}

// Updates the given cache node using the values in a given managed object.
//
// node: The cache node to update.
//
// managedObject: The managed object with which to update `node`.
//
// # Discussion
//
// This method is invoked by the framework after a save operation on a managed
// object context, once for each updated [NSManagedObject] instance.
//
// You override this method in a subclass to take the information from
// `managedObject` and update `node`.
//
// # Special Considerations
//
// You must override this method.
//
// See: https://developer.apple.com/documentation/CoreData/NSAtomicStore/updateCacheNode(_:from:)
func (a NSAtomicStore) UpdateCacheNodeFromManagedObject(node INSAtomicStoreCacheNode, managedObject INSManagedObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("updateCacheNode:fromManagedObject:"), node, managedObject)
}

// Method invoked before the store removes the given collection of cache
// nodes.
//
// cacheNodes: The set of cache nodes to remove.
//
// # Discussion
//
// This method is invoked by the store before the call to [NSAtomicStore.Save]
// with the collection of cache nodes marked as deleted by a managed object
// context. You can override this method to track the nodes which will not be
// made persistent in the [NSAtomicStore.Save] method.
//
// You should not invoke this method directly in a subclass.
//
// See: https://developer.apple.com/documentation/CoreData/NSAtomicStore/willRemoveCacheNodes(_:)
func (a NSAtomicStore) WillRemoveCacheNodes(cacheNodes foundation.INSSet) {
	objc.Send[objc.ID](a.ID, objc.Sel("willRemoveCacheNodes:"), cacheNodes)
}

// Saves the cache nodes.
//
// # Discussion
//
// You override this method to make persistent the necessary information from
// the cache nodes to the URL specified for the receiver.
//
// # Special Considerations
//
// You must override this method.
//
// See: https://developer.apple.com/documentation/CoreData/NSAtomicStore/save()
func (a NSAtomicStore) Save() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("save:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("save: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns the set of cache nodes registered with the receiver.
//
// # Return Value
//
// The set of cache nodes registered with the receiver.
//
// # Discussion
//
// You should modify this collection using [NSAtomicStore.AddCacheNodes]: and
// [NSAtomicStore.WillRemoveCacheNodes].
//
// See: https://developer.apple.com/documentation/CoreData/NSAtomicStore/cacheNodes()
func (a NSAtomicStore) CacheNodes() foundation.INSSet {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("cacheNodes"))
	return foundation.NSSetFromID(rv)
}

// Returns the cache node for a given managed object ID.
//
// objectID: A managed object ID.
//
// # Return Value
//
// The cache node for `objectID`.
//
// # Discussion
//
// This method is normally used by cache nodes to locate related cache nodes
// (by relationships).
//
// See: https://developer.apple.com/documentation/CoreData/NSAtomicStore/cacheNode(for:)
func (a NSAtomicStore) CacheNodeForObjectID(objectID INSManagedObjectID) INSAtomicStoreCacheNode {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("cacheNodeForObjectID:"), objectID)
	return NSAtomicStoreCacheNodeFromID(rv)
}

// Returns the reference object for a given managed object ID.
//
// objectID: A managed object ID.
//
// # Return Value
//
// The reference object for `objectID`.
//
// # Discussion
//
// Subclasses should invoke this method to extract the reference data from the
// object ID for each cache node if the data is to be made persistent.
//
// See: https://developer.apple.com/documentation/CoreData/NSAtomicStore/referenceObject(for:)
func (a NSAtomicStore) ReferenceObjectForObjectID(objectID INSManagedObjectID) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("referenceObjectForObjectID:"), objectID)
	return objectivec.Object{ID: rv}
}
