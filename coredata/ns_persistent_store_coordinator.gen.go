// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPersistentStoreCoordinator] class.
var (
	_NSPersistentStoreCoordinatorClass     NSPersistentStoreCoordinatorClass
	_NSPersistentStoreCoordinatorClassOnce sync.Once
)

func getNSPersistentStoreCoordinatorClass() NSPersistentStoreCoordinatorClass {
	_NSPersistentStoreCoordinatorClassOnce.Do(func() {
		_NSPersistentStoreCoordinatorClass = NSPersistentStoreCoordinatorClass{class: objc.GetClass("NSPersistentStoreCoordinator")}
	})
	return _NSPersistentStoreCoordinatorClass
}

// GetNSPersistentStoreCoordinatorClass returns the class object for NSPersistentStoreCoordinator.
func GetNSPersistentStoreCoordinatorClass() NSPersistentStoreCoordinatorClass {
	return getNSPersistentStoreCoordinatorClass()
}

type NSPersistentStoreCoordinatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPersistentStoreCoordinatorClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPersistentStoreCoordinatorClass) Alloc() NSPersistentStoreCoordinator {
	rv := objc.Send[NSPersistentStoreCoordinator](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that enables an app’s contexts and the underlying persistent
// stores to work together.
//
// # Overview
//
// A managed object context uses a coordinator to facilitate the persistence
// of its entities in the coordinator’s registered stores. A context can’t
// function without a coordinator because it relies on the coordinator’s
// access to the managed object model. The coordinator presents its registered
// stores as an aggregate, allowing a context to operate on the union of those
// stores instead of on each individually. A coordinator performs its work on
// a private queue and executes that work serially. You can use multiple
// coordinators if the work requires separate queues.
//
// Use a coordinator to add or remove persistent stores, change the type or
// location on-disk of those stores, query the metadata of a specific store,
// defer a store’s migrations, determine whether two objects originate from
// the same store, and so on.
//
// # Creating a persistent store coordinator
//
//   - [NSPersistentStoreCoordinator.InitWithManagedObjectModel]: Creates a persistent store coordinator with the specified managed object model.
//
// # Managing configuration
//
//   - [NSPersistentStoreCoordinator.Name]: The coordinator’s name.
//   - [NSPersistentStoreCoordinator.SetName]
//   - [NSPersistentStoreCoordinator.ManagedObjectModel]: The coordinator’s managed object model.
//   - [NSPersistentStoreCoordinator.PersistentStores]: The coordinator’s persistent stores.
//
// # Adding or removing a store
//
//   - [NSPersistentStoreCoordinator.AddPersistentStoreWithTypeConfigurationURLOptionsError]: Adds a specific type of persistent store at the provided location.
//   - [NSPersistentStoreCoordinator.AddPersistentStoreWithDescriptionCompletionHandler]: Adds a persistent store using the provided description.
//   - [NSPersistentStoreCoordinator.RemovePersistentStoreError]: Removes the specified persistent store from the coordinator.
//
// # Modifying a store
//
//   - [NSPersistentStoreCoordinator.DestroyPersistentStoreAtURLWithTypeOptionsError]: Deletes a specific type of persistent store at the provided location.
//   - [NSPersistentStoreCoordinator.MigratePersistentStoreToURLOptionsWithTypeError]: Changes the location and, if necessary, the store type of the specified persistent store.
//   - [NSPersistentStoreCoordinator.ReplacePersistentStoreAtURLDestinationOptionsWithPersistentStoreFromURLSourceOptionsStoreTypeError]: Replaces one persistent store with another.
//
// # Managing a store’s location
//
//   - [NSPersistentStoreCoordinator.SetURLForPersistentStore]: Changes the location of the specified persistent store.
//   - [NSPersistentStoreCoordinator.PersistentStoreForURL]: Returns the persistent store for the specified file URL.
//   - [NSPersistentStoreCoordinator.URLForPersistentStore]: Returns the location of the provided persistent store.
//
// # Managing a store’s metadata
//
//   - [NSPersistentStoreCoordinator.MetadataForPersistentStore]: Returns the metadata of the specified persistent store.
//   - [NSPersistentStoreCoordinator.SetMetadataForPersistentStore]: Updates the metadata for the specified persistent store.
//
// # Deferring a store’s migrations
//
//   - [NSPersistentStoreCoordinator.FinishDeferredLightweightMigrationTask]: Executes a single pending task of a deferred lightweight migration.
//   - [NSPersistentStoreCoordinator.FinishDeferredLightweightMigration]: Executes all remaining tasks of a deferred lightweight migration.
//
// # Performing tasks
//
//   - [NSPersistentStoreCoordinator.PerformBlock]: Executes the provided closure asynchronously on the coordinator’s queue.
//   - [NSPersistentStoreCoordinator.PerformBlockAndWait]: Executes the provided closure on the coordinator’s queue and waits for it to finish.
//   - [NSPersistentStoreCoordinator.ExecuteRequestWithContextError]: Executes the specified request on each of the coordinator’s persistent stores.
//
// # Maintaining a record of changes
//
//   - [NSPersistentStoreCoordinator.CurrentPersistentHistoryTokenFromStores]: Returns a single persistent history token representing all of the specified stores.
//
// # Getting individual object identifiers
//
//   - [NSPersistentStoreCoordinator.ManagedObjectIDForURIRepresentation]: Returns the object identifier for the specified URI representation.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator
type NSPersistentStoreCoordinator struct {
	objectivec.Object
}

// NSPersistentStoreCoordinatorFromID constructs a [NSPersistentStoreCoordinator] from an objc.ID.
//
// An object that enables an app’s contexts and the underlying persistent
// stores to work together.
func NSPersistentStoreCoordinatorFromID(id objc.ID) NSPersistentStoreCoordinator {
	return NSPersistentStoreCoordinator{objectivec.Object{ID: id}}
}

// NOTE: NSPersistentStoreCoordinator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPersistentStoreCoordinator] class.
//
// # Creating a persistent store coordinator
//
//   - [INSPersistentStoreCoordinator.InitWithManagedObjectModel]: Creates a persistent store coordinator with the specified managed object model.
//
// # Managing configuration
//
//   - [INSPersistentStoreCoordinator.Name]: The coordinator’s name.
//   - [INSPersistentStoreCoordinator.SetName]
//   - [INSPersistentStoreCoordinator.ManagedObjectModel]: The coordinator’s managed object model.
//   - [INSPersistentStoreCoordinator.PersistentStores]: The coordinator’s persistent stores.
//
// # Adding or removing a store
//
//   - [INSPersistentStoreCoordinator.AddPersistentStoreWithTypeConfigurationURLOptionsError]: Adds a specific type of persistent store at the provided location.
//   - [INSPersistentStoreCoordinator.AddPersistentStoreWithDescriptionCompletionHandler]: Adds a persistent store using the provided description.
//   - [INSPersistentStoreCoordinator.RemovePersistentStoreError]: Removes the specified persistent store from the coordinator.
//
// # Modifying a store
//
//   - [INSPersistentStoreCoordinator.DestroyPersistentStoreAtURLWithTypeOptionsError]: Deletes a specific type of persistent store at the provided location.
//   - [INSPersistentStoreCoordinator.MigratePersistentStoreToURLOptionsWithTypeError]: Changes the location and, if necessary, the store type of the specified persistent store.
//   - [INSPersistentStoreCoordinator.ReplacePersistentStoreAtURLDestinationOptionsWithPersistentStoreFromURLSourceOptionsStoreTypeError]: Replaces one persistent store with another.
//
// # Managing a store’s location
//
//   - [INSPersistentStoreCoordinator.SetURLForPersistentStore]: Changes the location of the specified persistent store.
//   - [INSPersistentStoreCoordinator.PersistentStoreForURL]: Returns the persistent store for the specified file URL.
//   - [INSPersistentStoreCoordinator.URLForPersistentStore]: Returns the location of the provided persistent store.
//
// # Managing a store’s metadata
//
//   - [INSPersistentStoreCoordinator.MetadataForPersistentStore]: Returns the metadata of the specified persistent store.
//   - [INSPersistentStoreCoordinator.SetMetadataForPersistentStore]: Updates the metadata for the specified persistent store.
//
// # Deferring a store’s migrations
//
//   - [INSPersistentStoreCoordinator.FinishDeferredLightweightMigrationTask]: Executes a single pending task of a deferred lightweight migration.
//   - [INSPersistentStoreCoordinator.FinishDeferredLightweightMigration]: Executes all remaining tasks of a deferred lightweight migration.
//
// # Performing tasks
//
//   - [INSPersistentStoreCoordinator.PerformBlock]: Executes the provided closure asynchronously on the coordinator’s queue.
//   - [INSPersistentStoreCoordinator.PerformBlockAndWait]: Executes the provided closure on the coordinator’s queue and waits for it to finish.
//   - [INSPersistentStoreCoordinator.ExecuteRequestWithContextError]: Executes the specified request on each of the coordinator’s persistent stores.
//
// # Maintaining a record of changes
//
//   - [INSPersistentStoreCoordinator.CurrentPersistentHistoryTokenFromStores]: Returns a single persistent history token representing all of the specified stores.
//
// # Getting individual object identifiers
//
//   - [INSPersistentStoreCoordinator.ManagedObjectIDForURIRepresentation]: Returns the object identifier for the specified URI representation.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator
type INSPersistentStoreCoordinator interface {
	objectivec.IObject

	// Topic: Creating a persistent store coordinator

	// Creates a persistent store coordinator with the specified managed object model.
	InitWithManagedObjectModel(model INSManagedObjectModel) NSPersistentStoreCoordinator

	// Topic: Managing configuration

	// The coordinator’s name.
	Name() string
	SetName(value string)
	// The coordinator’s managed object model.
	ManagedObjectModel() INSManagedObjectModel
	// The coordinator’s persistent stores.
	PersistentStores() []NSPersistentStore

	// Topic: Adding or removing a store

	// Adds a specific type of persistent store at the provided location.
	AddPersistentStoreWithTypeConfigurationURLOptionsError(storeType string, configuration string, storeURL foundation.NSURL, options foundation.INSDictionary) (INSPersistentStore, error)
	// Adds a persistent store using the provided description.
	AddPersistentStoreWithDescriptionCompletionHandler(storeDescription INSPersistentStoreDescription, block PersistentStoreDescriptionErrorHandler)
	// Removes the specified persistent store from the coordinator.
	RemovePersistentStoreError(store INSPersistentStore) (bool, error)

	// Topic: Modifying a store

	// Deletes a specific type of persistent store at the provided location.
	DestroyPersistentStoreAtURLWithTypeOptionsError(url foundation.NSURL, storeType string, options foundation.INSDictionary) (bool, error)
	// Changes the location and, if necessary, the store type of the specified persistent store.
	MigratePersistentStoreToURLOptionsWithTypeError(store INSPersistentStore, URL foundation.NSURL, options foundation.INSDictionary, storeType string) (INSPersistentStore, error)
	// Replaces one persistent store with another.
	ReplacePersistentStoreAtURLDestinationOptionsWithPersistentStoreFromURLSourceOptionsStoreTypeError(destinationURL foundation.NSURL, destinationOptions foundation.INSDictionary, sourceURL foundation.NSURL, sourceOptions foundation.INSDictionary, storeType string) (bool, error)

	// Topic: Managing a store’s location

	// Changes the location of the specified persistent store.
	SetURLForPersistentStore(url foundation.NSURL, store INSPersistentStore) bool
	// Returns the persistent store for the specified file URL.
	PersistentStoreForURL(URL foundation.NSURL) INSPersistentStore
	// Returns the location of the provided persistent store.
	URLForPersistentStore(store INSPersistentStore) foundation.NSURL

	// Topic: Managing a store’s metadata

	// Returns the metadata of the specified persistent store.
	MetadataForPersistentStore(store INSPersistentStore) foundation.INSDictionary
	// Updates the metadata for the specified persistent store.
	SetMetadataForPersistentStore(metadata foundation.INSDictionary, store INSPersistentStore)

	// Topic: Deferring a store’s migrations

	// Executes a single pending task of a deferred lightweight migration.
	FinishDeferredLightweightMigrationTask() (bool, error)
	// Executes all remaining tasks of a deferred lightweight migration.
	FinishDeferredLightweightMigration() (bool, error)

	// Topic: Performing tasks

	// Executes the provided closure asynchronously on the coordinator’s queue.
	PerformBlock(block VoidHandler)
	// Executes the provided closure on the coordinator’s queue and waits for it to finish.
	PerformBlockAndWait(block VoidHandler)
	// Executes the specified request on each of the coordinator’s persistent stores.
	ExecuteRequestWithContextError(request INSPersistentStoreRequest, context INSManagedObjectContext) (objectivec.IObject, error)

	// Topic: Maintaining a record of changes

	// Returns a single persistent history token representing all of the specified stores.
	CurrentPersistentHistoryTokenFromStores(stores foundation.INSArray) INSPersistentHistoryToken

	// Topic: Getting individual object identifiers

	// Returns the object identifier for the specified URI representation.
	ManagedObjectIDForURIRepresentation(url foundation.NSURL) INSManagedObjectID
}

// Init initializes the instance.
func (p NSPersistentStoreCoordinator) Init() NSPersistentStoreCoordinator {
	rv := objc.Send[NSPersistentStoreCoordinator](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPersistentStoreCoordinator) Autorelease() NSPersistentStoreCoordinator {
	rv := objc.Send[NSPersistentStoreCoordinator](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPersistentStoreCoordinator creates a new NSPersistentStoreCoordinator instance.
func NewNSPersistentStoreCoordinator() NSPersistentStoreCoordinator {
	class := getNSPersistentStoreCoordinatorClass()
	rv := objc.Send[NSPersistentStoreCoordinator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a persistent store coordinator with the specified managed object
// model.
//
// model: A managed object model.
//
// # Return Value
//
// The receiver, initialized with `model`.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/init(managedObjectModel:)
func NewPersistentStoreCoordinatorWithManagedObjectModel(model INSManagedObjectModel) NSPersistentStoreCoordinator {
	instance := getNSPersistentStoreCoordinatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithManagedObjectModel:"), model)
	return NSPersistentStoreCoordinatorFromID(rv)
}

// Creates a persistent store coordinator with the specified managed object
// model.
//
// model: A managed object model.
//
// # Return Value
//
// The receiver, initialized with `model`.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/init(managedObjectModel:)
func (p NSPersistentStoreCoordinator) InitWithManagedObjectModel(model INSManagedObjectModel) NSPersistentStoreCoordinator {
	rv := objc.Send[NSPersistentStoreCoordinator](p.ID, objc.Sel("initWithManagedObjectModel:"), model)
	return rv
}

// Adds a specific type of persistent store at the provided location.
//
// storeType: A string constant (such as [NSSQLiteStoreType]) that specifies the store
// type—see [Persistent Store Types] for possible values.
//
// configuration: The name of a configuration in the receiver’s managed object model that
// will be used by the new store. The configuration can be `nil`, in which
// case no other configurations are allowed.
//
// storeURL: The file location of the persistent store.
//
// options: A dictionary containing key-value pairs that specify whether the store
// should be read-only, and whether (for an XML store) the XML file should be
// validated against the DTD before it is read. For key definitions, see
// [Store options] and [Migration options]. This value may be `nil`.
//
// # Return Value
//
// The newly created store or, if an error occurs, `nil`.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/addPersistentStore(ofType:configurationName:at:options:)
//
// [Persistent Store Types]: https://developer.apple.com/documentation/CoreData/persistent-store-types
// [Migration options]: https://developer.apple.com/documentation/CoreData/migration-options
// [Store options]: https://developer.apple.com/documentation/CoreData/store-options
func (p NSPersistentStoreCoordinator) AddPersistentStoreWithTypeConfigurationURLOptionsError(storeType string, configuration string, storeURL foundation.NSURL, options foundation.INSDictionary) (INSPersistentStore, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](p.ID, objc.Sel("addPersistentStoreWithType:configuration:URL:options:error:"), objc.String(storeType), objc.String(configuration), storeURL, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSPersistentStore{}, foundation.NSErrorFrom(errorPtr)
	}
	return NSPersistentStoreFromID(rv), nil

}

// Adds a persistent store using the provided description.
//
// storeDescription: A description object used to create and load a persistent store.
//
// block: The completion handler block that’s invoked after the store is added.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/addPersistentStore(with:completionHandler:)
func (p NSPersistentStoreCoordinator) AddPersistentStoreWithDescriptionCompletionHandler(storeDescription INSPersistentStoreDescription, block PersistentStoreDescriptionErrorHandler) {
	_block1, _ := NewPersistentStoreDescriptionErrorBlock(block)
	objc.Send[objc.ID](p.ID, objc.Sel("addPersistentStoreWithDescription:completionHandler:"), storeDescription, _block1)
}

// Removes the specified persistent store from the coordinator.
//
// store: A persistent store.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/remove(_:)
func (p NSPersistentStoreCoordinator) RemovePersistentStoreError(store INSPersistentStore) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](p.ID, objc.Sel("removePersistentStore:error:"), store, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("removePersistentStore:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Deletes a specific type of persistent store at the provided location.
//
// url: The store’s location.
//
// storeType: The store type. For possible values, see [NSPersistentStore.StoreType].
//
// options: A dictionary containing key-value pairs that specify store behavior and
// characteristics. For more information, see [Store options].
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/destroyPersistentStore(at:ofType:options:)
//
// [NSPersistentStore.StoreType]: https://developer.apple.com/documentation/CoreData/NSPersistentStore/StoreType
// [Store options]: https://developer.apple.com/documentation/CoreData/store-options
func (p NSPersistentStoreCoordinator) DestroyPersistentStoreAtURLWithTypeOptionsError(url foundation.NSURL, storeType string, options foundation.INSDictionary) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](p.ID, objc.Sel("destroyPersistentStoreAtURL:withType:options:error:"), url, objc.String(storeType), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("destroyPersistentStoreAtURL:withType:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Changes the location and, if necessary, the store type of the specified
// persistent store.
//
// store: A persistent store.
//
// URL: An URL object that specifies the location for the new store.
//
// options: A dictionary containing key-value pairs that specify whether the store
// should be read-only, and whether (for an XML store) the XML file should be
// validated against the DTD before it is read. For key definitions, see
// [Store options].
//
// storeType: A string constant (such as [NSSQLiteStoreType]) that specifies the type of
// the new store—see [Persistent Store Types].
//
// # Return Value
//
// If the migration is successful, the new store, otherwise `nil`.
//
// # Discussion
//
// This method is typically used for “Save As” operations. Performance may
// vary depending on the type of old and new store. For more details of the
// action of this method, see Persistent Store Features in [Core Data
// Programming Guide].
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/migratePersistentStore(_:to:options:withType:)
//
// [Store options]: https://developer.apple.com/documentation/CoreData/store-options
// [Persistent Store Types]: https://developer.apple.com/documentation/CoreData/persistent-store-types
// [Core Data Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CoreData/index.html#//apple_ref/doc/uid/TP40001075
func (p NSPersistentStoreCoordinator) MigratePersistentStoreToURLOptionsWithTypeError(store INSPersistentStore, URL foundation.NSURL, options foundation.INSDictionary, storeType string) (INSPersistentStore, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](p.ID, objc.Sel("migratePersistentStore:toURL:options:withType:error:"), store, URL, options, objc.String(storeType), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSPersistentStore{}, foundation.NSErrorFrom(errorPtr)
	}
	return NSPersistentStoreFromID(rv), nil

}

// Replaces one persistent store with another.
//
// destinationURL: The location of the store to replace.
//
// destinationOptions: A dictionary containing key-value pairs that specify the behavior and
// characteristics of the store to replace. For more information, see [Store
// options].
//
// sourceURL: The location of the store to use as the replacement.
//
// sourceOptions: A dictionary containing key-value pairs that specify the behavior and
// characteristics of the replacement store. For more information, see [Store
// options].
//
// storeType: The store type of the replacement store. For possible values, see
// [NSPersistentStore.StoreType].
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/replacePersistentStore(at:destinationOptions:withPersistentStoreFrom:sourceOptions:ofType:)
//
// [Store options]: https://developer.apple.com/documentation/CoreData/store-options
// [NSPersistentStore.StoreType]: https://developer.apple.com/documentation/CoreData/NSPersistentStore/StoreType
//
// [Store options]: https://developer.apple.com/documentation/CoreData/store-options
func (p NSPersistentStoreCoordinator) ReplacePersistentStoreAtURLDestinationOptionsWithPersistentStoreFromURLSourceOptionsStoreTypeError(destinationURL foundation.NSURL, destinationOptions foundation.INSDictionary, sourceURL foundation.NSURL, sourceOptions foundation.INSDictionary, storeType string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](p.ID, objc.Sel("replacePersistentStoreAtURL:destinationOptions:withPersistentStoreFromURL:sourceOptions:storeType:error:"), destinationURL, destinationOptions, sourceURL, sourceOptions, objc.String(storeType), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("replacePersistentStoreAtURL:destinationOptions:withPersistentStoreFromURL:sourceOptions:storeType:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Changes the location of the specified persistent store.
//
// url: The new location for `store`.
//
// store: A persistent store associated with the receiver.
//
// # Return Value
//
// true if the store was relocated, otherwise false.
//
// # Discussion
//
// For atomic stores, this method alters the location to which the next save
// operation will write the file; for non-atomic stores, invoking this method
// will relinquish the existing connection and create a new one at the
// specified URL. (For non-atomic stores, a store must already exist at the
// destination URL; a new store will not be created.)
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/setURL(_:for:)
func (p NSPersistentStoreCoordinator) SetURLForPersistentStore(url foundation.NSURL, store INSPersistentStore) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("setURL:forPersistentStore:"), url, store)
	return rv
}

// Returns the persistent store for the specified file URL.
//
// URL: An URL object that specifies the location of a persistent store.
//
// # Return Value
//
// The persistent store at the location specified by [URL].
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/persistentStore(for:)
func (p NSPersistentStoreCoordinator) PersistentStoreForURL(URL foundation.NSURL) INSPersistentStore {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("persistentStoreForURL:"), URL)
	return NSPersistentStoreFromID(rv)
}

// Returns the location of the provided persistent store.
//
// store: A persistent store.
//
// # Return Value
//
// The URL for `store`.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/url(for:)
func (p NSPersistentStoreCoordinator) URLForPersistentStore(store INSPersistentStore) foundation.NSURL {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("URLForPersistentStore:"), store)
	return foundation.NSURLFromID(rv)
}

// Returns the metadata of the specified persistent store.
//
// store: A persistent store.
//
// # Return Value
//
// A dictionary that contains the metadata currently stored or to-be-stored in
// `store`.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/metadata(for:)
func (p NSPersistentStoreCoordinator) MetadataForPersistentStore(store INSPersistentStore) foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("metadataForPersistentStore:"), store)
	return foundation.NSDictionaryFromID(rv)
}

// Updates the metadata for the specified persistent store.
//
// metadata: A dictionary containing metadata for the store.
//
// store: A persistent store.
//
// # Discussion
//
// The store type and UUID ([NSStoreTypeKey] and [NSStoreUUIDKey]) are always
// added automatically, however [NSStoreUUIDKey] is only added if it is not
// set manually as part of the dictionary argument.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/setMetadata(_:for:)
func (p NSPersistentStoreCoordinator) SetMetadataForPersistentStore(metadata foundation.INSDictionary, store INSPersistentStore) {
	objc.Send[objc.ID](p.ID, objc.Sel("setMetadata:forPersistentStore:"), metadata, store)
}

// Executes a single pending task of a deferred lightweight migration.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/finishDeferredLightweightMigrationTask()
func (p NSPersistentStoreCoordinator) FinishDeferredLightweightMigrationTask() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](p.ID, objc.Sel("finishDeferredLightweightMigrationTask:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("finishDeferredLightweightMigrationTask: returned NO with nil NSError")
	}
	return rv, nil

}

// Executes all remaining tasks of a deferred lightweight migration.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/finishDeferredLightweightMigration()
func (p NSPersistentStoreCoordinator) FinishDeferredLightweightMigration() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](p.ID, objc.Sel("finishDeferredLightweightMigration:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("finishDeferredLightweightMigration: returned NO with nil NSError")
	}
	return rv, nil

}

// Executes the provided closure asynchronously on the coordinator’s queue.
//
// block: The closure to execute.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/perform(_:)-7jqb
func (p NSPersistentStoreCoordinator) PerformBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](p.ID, objc.Sel("performBlock:"), _block0)
}

// Executes the provided closure on the coordinator’s queue and waits for it
// to finish.
//
// block: The closure to execute.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/performAndWait(_:)-d3kq
func (p NSPersistentStoreCoordinator) PerformBlockAndWait(block VoidHandler) {
	_block0, _cleanup0 := NewVoidBlock(block)
	defer _cleanup0()
	objc.Send[objc.ID](p.ID, objc.Sel("performBlockAndWait:"), _block0)
}

// Executes the specified request on each of the coordinator’s persistent
// stores.
//
// request: A fetch or save request.
//
// context: The context against which `request` should be executed.
//
// # Return Value
//
// An array containing managed objects, managed object IDs, or dictionaries as
// appropriate for a fetch request; an empty array if `request` is a save
// request, or `nil` if an error occurred.
//
// # Discussion
//
// User defined requests return arrays of arrays, where a nested array is the
// result returned from a single store.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/execute(_:with:)
func (p NSPersistentStoreCoordinator) ExecuteRequestWithContextError(request INSPersistentStoreRequest, context INSManagedObjectContext) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](p.ID, objc.Sel("executeRequest:withContext:error:"), request, context, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// Returns a single persistent history token representing all of the specified
// stores.
//
// stores: The persistent stores of interest.
//
// # Return Value
//
// A persistent history token, or `nil` if the coordinator can’t create one.
//
// # Discussion
//
// If you specify `nil` or provide an empty array, the coordinator attempts to
// create a token for all of its registered stores.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/currentPersistentHistoryToken(fromStores:)
func (p NSPersistentStoreCoordinator) CurrentPersistentHistoryTokenFromStores(stores foundation.INSArray) INSPersistentHistoryToken {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("currentPersistentHistoryTokenFromStores:"), stores)
	return NSPersistentHistoryTokenFromID(rv)
}

// Returns the object identifier for the specified URI representation.
//
// url: An URL object containing a URI that specify a managed object.
//
// # Return Value
//
// An object ID for the object specified by `url`.
//
// # Discussion
//
// The URI representation contains a UUID of the store the ID is coming from,
// and the coordinator can match it against the stores added to it.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/managedObjectID(forURIRepresentation:)
func (p NSPersistentStoreCoordinator) ManagedObjectIDForURIRepresentation(url foundation.NSURL) INSManagedObjectID {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("managedObjectIDForURIRepresentation:"), url)
	return NSManagedObjectIDFromID(rv)
}

// Registers a persistent store subclass using the specified store type
// identifier.
//
// storeClass: The [NSPersistentStore] subclass to use for the store of type `storeType`.
//
// storeType: A unique string that identifies a store type.
//
// # Discussion
//
// You must invoke this method before a custom subclass of [NSPersistentStore]
// can be loaded into a persistent store coordinator.
//
// You can pass `nil` for `storeClass` to unregister the store type.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/registerStoreClass(_:forStoreType:)
func (_NSPersistentStoreCoordinatorClass NSPersistentStoreCoordinatorClass) RegisterStoreClassForStoreType(storeClass objectivec.Class, storeType string) {
	objc.Send[objc.ID](objc.ID(_NSPersistentStoreCoordinatorClass.class), objc.Sel("registerStoreClass:forStoreType:"), storeClass, objc.String(storeType))
}

// Updates the metadata of a specific type of persistent store at the provided
// location.
//
// metadata: A dictionary that contains the metadata to store.
//
// storeType: The type of store. If `nil`, Core Data automatically attempts to determine
// the store class to use.
//
// url: The file URL of the store.
//
// options: A dictionary that contains options for the store.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/setMetadata(_:forPersistentStoreOfType:at:options:)
func (_NSPersistentStoreCoordinatorClass NSPersistentStoreCoordinatorClass) SetMetadataForPersistentStoreOfTypeURLOptionsError(metadata foundation.INSDictionary, storeType string, url foundation.NSURL, options foundation.INSDictionary) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_NSPersistentStoreCoordinatorClass.class), objc.Sel("setMetadata:forPersistentStoreOfType:URL:options:error:"), metadata, objc.String(storeType), url, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setMetadata:forPersistentStoreOfType:URL:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns the metadata of a specific type of persistent store at the provided
// location.
//
// storeType: The type of the store. If `nil`, Core Data automatically attempts to
// determine the store class to use.
//
// url: The file URL of the store.
//
// options: A dictionary that contains options for the store.
//
// # Return Value
//
// A dictionary that contains, at a minimum, values for the [NSStoreTypeKey]
// and [NSStoreUUIDKey] keys.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/metadataForPersistentStore(ofType:at:options:)
//
// [NSStoreTypeKey]: https://developer.apple.com/documentation/CoreData/NSStoreTypeKey
// [NSStoreUUIDKey]: https://developer.apple.com/documentation/CoreData/NSStoreUUIDKey
func (_NSPersistentStoreCoordinatorClass NSPersistentStoreCoordinatorClass) MetadataForPersistentStoreOfTypeURLOptionsError(storeType string, url foundation.NSURL, options foundation.INSDictionary) (foundation.INSDictionary, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentStoreCoordinatorClass.class), objc.Sel("metadataForPersistentStoreOfType:URL:options:error:"), objc.String(storeType), url, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDictionaryFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/cachedModelForPersistentStore(at:options:)
func (_NSPersistentStoreCoordinatorClass NSPersistentStoreCoordinatorClass) CachedModelForPersistentStoreAtURLOptionsError(url foundation.NSURL, options foundation.INSDictionary) (NSManagedObjectModel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentStoreCoordinatorClass.class), objc.Sel("cachedModelForPersistentStoreAtURL:options:error:"), url, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSManagedObjectModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return NSManagedObjectModelFromID(rv), nil

}

// The coordinator’s name.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/name
func (p NSPersistentStoreCoordinator) Name() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (p NSPersistentStoreCoordinator) SetName(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setName:"), objc.String(value))
}

// The coordinator’s managed object model.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/managedObjectModel
func (p NSPersistentStoreCoordinator) ManagedObjectModel() INSManagedObjectModel {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("managedObjectModel"))
	return NSManagedObjectModelFromID(objc.ID(rv))
}

// The coordinator’s persistent stores.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/persistentStores
func (p NSPersistentStoreCoordinator) PersistentStores() []NSPersistentStore {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("persistentStores"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSPersistentStore {
		return NSPersistentStoreFromID(id)
	})
}

// The coordinator’s registered store types.
//
// # Return Value
//
// A dictionary of the registered store types—the keys are the store type
// strings, and the values are the [NSPersistentStore] subclasses.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/registeredStoreTypes
func (_NSPersistentStoreCoordinatorClass NSPersistentStoreCoordinatorClass) RegisteredStoreTypes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentStoreCoordinatorClass.class), objc.Sel("registeredStoreTypes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// AddPersistentStoreWithDescription is a synchronous wrapper around [NSPersistentStoreCoordinator.AddPersistentStoreWithDescriptionCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (p NSPersistentStoreCoordinator) AddPersistentStoreWithDescription(ctx context.Context, storeDescription INSPersistentStoreDescription) (*NSPersistentStoreDescription, error) {
	type result struct {
		val *NSPersistentStoreDescription
		err error
	}
	done := make(chan result, 1)
	p.AddPersistentStoreWithDescriptionCompletionHandler(storeDescription, func(val *NSPersistentStoreDescription, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// PerformBlockSync is a synchronous wrapper around [NSPersistentStoreCoordinator.PerformBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (p NSPersistentStoreCoordinator) PerformBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	p.PerformBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PerformBlockAndWaitSync is a synchronous wrapper around [NSPersistentStoreCoordinator.PerformBlockAndWait].
// It blocks until the completion handler fires or the context is cancelled.
func (p NSPersistentStoreCoordinator) PerformBlockAndWaitSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	p.PerformBlockAndWait(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
