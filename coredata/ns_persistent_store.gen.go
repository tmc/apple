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

// The class instance for the [NSPersistentStore] class.
var (
	_NSPersistentStoreClass     NSPersistentStoreClass
	_NSPersistentStoreClassOnce sync.Once
)

func getNSPersistentStoreClass() NSPersistentStoreClass {
	_NSPersistentStoreClassOnce.Do(func() {
		_NSPersistentStoreClass = NSPersistentStoreClass{class: objc.GetClass("NSPersistentStore")}
	})
	return _NSPersistentStoreClass
}

// GetNSPersistentStoreClass returns the class object for NSPersistentStore.
func GetNSPersistentStoreClass() NSPersistentStoreClass {
	return getNSPersistentStoreClass()
}

type NSPersistentStoreClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPersistentStoreClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPersistentStoreClass) Alloc() NSPersistentStore {
	rv := objc.Send[NSPersistentStore](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The abstract base class for all Core Data persistent stores.
//
// # Overview
//
// Core Data provides four store types—SQLite, Binary, XML, and In-Memory
// (the XML store is not available on iOS); these are described in Persistent
// Store Features. Core Data also provides subclasses of [NSPersistentStore]
// that you can use to define your own store types: [NSAtomicStore] and
// [NSIncrementalStore]. The Binary and XML stores are examples of atomic
// stores that inherit functionality from [NSAtomicStore].
//
// # Subclassing Notes
//
// You should not subclass [NSPersistentStore] directly. Core Data only
// supports subclassing of [NSAtomicStore] and [NSIncrementalStore].
//
// The designated initializer is
// [NSPersistentStore.InitWithPersistentStoreCoordinatorConfigurationNameURLOptions].
// When you implement the initializer, you must ensure you load metadata
// during initialization and set it using [NSPersistentStore.Metadata].
//
// You must override these methods:
//
// - [NSPersistentStore.Type] - [NSPersistentStore.Metadata] -
// [NSPersistentStoreClass.MetadataForPersistentStoreWithURLError] -
// [NSPersistentStoreClass.SetMetadataForPersistentStoreWithURLError]
//
// # Creating a Persistent Store
//
//   - [NSPersistentStore.InitWithPersistentStoreCoordinatorConfigurationNameURLOptions]: Returns a store initialized with the given arguments.
//
// # Getting Store Configuration
//
//   - [NSPersistentStore.ConfigurationName]: The name of the managed object model configuration that creates the persistent store.
//   - [NSPersistentStore.Options]: The options that Core Data uses to create the store.
//   - [NSPersistentStore.PersistentStoreCoordinator]: The persistent store coordinator that loads the persistent store.
//   - [NSPersistentStore.Type]: The type string of the persistent store.
//
// # Managing Store Attributes
//
//   - [NSPersistentStore.Identifier]: The unique identifier for the persistent store.
//   - [NSPersistentStore.SetIdentifier]
//   - [NSPersistentStore.IsReadOnly]: A Boolean value that indicates whether the persistent store is read-only.
//   - [NSPersistentStore.SetReadOnly]
//   - [NSPersistentStore.URL]: The URL for the persistent store.
//   - [NSPersistentStore.SetURL]
//
// # Managing Store Metadata
//
//   - [NSPersistentStore.LoadMetadata]: Instructs the persistent store to load its metadata.
//   - [NSPersistentStore.Metadata]: The metadata for the persistent store.
//   - [NSPersistentStore.SetMetadata]
//
// # Responding to the Store Life Cycle
//
//   - [NSPersistentStore.DidAddToPersistentStoreCoordinator]: Invoked after the persistent store has been added to the persistent store coordinator.
//   - [NSPersistentStore.WillRemoveFromPersistentStoreCoordinator]: Invoked before the persistent store is removed from the persistent store coordinator.
//
// # Integrating with Spotlight
//
//   - [NSPersistentStore.CoreSpotlightExporter]: The spotlight exporter associated with this persistent store.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStore
type NSPersistentStore struct {
	objectivec.Object
}

// NSPersistentStoreFromID constructs a [NSPersistentStore] from an objc.ID.
//
// The abstract base class for all Core Data persistent stores.
func NSPersistentStoreFromID(id objc.ID) NSPersistentStore {
	return NSPersistentStore{objectivec.Object{ID: id}}
}

// NOTE: NSPersistentStore adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPersistentStore] class.
//
// # Creating a Persistent Store
//
//   - [INSPersistentStore.InitWithPersistentStoreCoordinatorConfigurationNameURLOptions]: Returns a store initialized with the given arguments.
//
// # Getting Store Configuration
//
//   - [INSPersistentStore.ConfigurationName]: The name of the managed object model configuration that creates the persistent store.
//   - [INSPersistentStore.Options]: The options that Core Data uses to create the store.
//   - [INSPersistentStore.PersistentStoreCoordinator]: The persistent store coordinator that loads the persistent store.
//   - [INSPersistentStore.Type]: The type string of the persistent store.
//
// # Managing Store Attributes
//
//   - [INSPersistentStore.Identifier]: The unique identifier for the persistent store.
//   - [INSPersistentStore.SetIdentifier]
//   - [INSPersistentStore.IsReadOnly]: A Boolean value that indicates whether the persistent store is read-only.
//   - [INSPersistentStore.SetReadOnly]
//   - [INSPersistentStore.URL]: The URL for the persistent store.
//   - [INSPersistentStore.SetURL]
//
// # Managing Store Metadata
//
//   - [INSPersistentStore.LoadMetadata]: Instructs the persistent store to load its metadata.
//   - [INSPersistentStore.Metadata]: The metadata for the persistent store.
//   - [INSPersistentStore.SetMetadata]
//
// # Responding to the Store Life Cycle
//
//   - [INSPersistentStore.DidAddToPersistentStoreCoordinator]: Invoked after the persistent store has been added to the persistent store coordinator.
//   - [INSPersistentStore.WillRemoveFromPersistentStoreCoordinator]: Invoked before the persistent store is removed from the persistent store coordinator.
//
// # Integrating with Spotlight
//
//   - [INSPersistentStore.CoreSpotlightExporter]: The spotlight exporter associated with this persistent store.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStore
type INSPersistentStore interface {
	objectivec.IObject

	// Topic: Creating a Persistent Store

	// Returns a store initialized with the given arguments.
	InitWithPersistentStoreCoordinatorConfigurationNameURLOptions(root INSPersistentStoreCoordinator, name string, url foundation.NSURL, options foundation.INSDictionary) NSPersistentStore

	// Topic: Getting Store Configuration

	// The name of the managed object model configuration that creates the persistent store.
	ConfigurationName() string
	// The options that Core Data uses to create the store.
	Options() foundation.INSDictionary
	// The persistent store coordinator that loads the persistent store.
	PersistentStoreCoordinator() INSPersistentStoreCoordinator
	// The type string of the persistent store.
	Type() string

	// Topic: Managing Store Attributes

	// The unique identifier for the persistent store.
	Identifier() string
	SetIdentifier(value string)
	// A Boolean value that indicates whether the persistent store is read-only.
	IsReadOnly() bool
	SetReadOnly(value bool)
	// The URL for the persistent store.
	URL() foundation.NSURL
	SetURL(value foundation.NSURL)

	// Topic: Managing Store Metadata

	// Instructs the persistent store to load its metadata.
	LoadMetadata() (bool, error)
	// The metadata for the persistent store.
	Metadata() foundation.INSDictionary
	SetMetadata(value foundation.INSDictionary)

	// Topic: Responding to the Store Life Cycle

	// Invoked after the persistent store has been added to the persistent store coordinator.
	DidAddToPersistentStoreCoordinator(coordinator INSPersistentStoreCoordinator)
	// Invoked before the persistent store is removed from the persistent store coordinator.
	WillRemoveFromPersistentStoreCoordinator(coordinator INSPersistentStoreCoordinator)

	// Topic: Integrating with Spotlight

	// The spotlight exporter associated with this persistent store.
	CoreSpotlightExporter() INSCoreDataCoreSpotlightDelegate
}

// Init initializes the instance.
func (p NSPersistentStore) Init() NSPersistentStore {
	rv := objc.Send[NSPersistentStore](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPersistentStore) Autorelease() NSPersistentStore {
	rv := objc.Send[NSPersistentStore](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPersistentStore creates a new NSPersistentStore instance.
func NewNSPersistentStore() NSPersistentStore {
	class := getNSPersistentStoreClass()
	rv := objc.Send[NSPersistentStore](objc.ID(class.class), objc.Sel("new"))
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
func NewPersistentStoreWithPersistentStoreCoordinatorConfigurationNameURLOptions(root INSPersistentStoreCoordinator, name string, url foundation.NSURL, options foundation.INSDictionary) NSPersistentStore {
	instance := getNSPersistentStoreClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPersistentStoreCoordinator:configurationName:URL:options:"), root, objc.String(name), url, options)
	return NSPersistentStoreFromID(rv)
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
func (p NSPersistentStore) InitWithPersistentStoreCoordinatorConfigurationNameURLOptions(root INSPersistentStoreCoordinator, name string, url foundation.NSURL, options foundation.INSDictionary) NSPersistentStore {
	rv := objc.Send[NSPersistentStore](p.ID, objc.Sel("initWithPersistentStoreCoordinator:configurationName:URL:options:"), root, objc.String(name), url, options)
	return rv
}

// Instructs the persistent store to load its metadata.
//
// # Discussion
//
// There is no way to return an error if the store is invalid.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStore/loadMetadata()
func (p NSPersistentStore) LoadMetadata() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](p.ID, objc.Sel("loadMetadata:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("loadMetadata: returned NO with nil NSError")
	}
	return rv, nil

}

// Invoked after the persistent store has been added to the persistent store
// coordinator.
//
// coordinator: The persistent store coordinator to which the receiver was added.
//
// # Discussion
//
// The default implementation does nothing. You can override this method in a
// subclass in order to perform any kind of setup necessary before the load
// method is invoked.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStore/didAdd(to:)
func (p NSPersistentStore) DidAddToPersistentStoreCoordinator(coordinator INSPersistentStoreCoordinator) {
	objc.Send[objc.ID](p.ID, objc.Sel("didAddToPersistentStoreCoordinator:"), coordinator)
}

// Invoked before the persistent store is removed from the persistent store
// coordinator.
//
// coordinator: The persistent store coordinator from which the receiver was removed.
//
// # Discussion
//
// The default implementation does nothing. You can override this method in a
// subclass in order to perform any clean-up before the store is removed from
// the coordinator (and deallocated).
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStore/willRemove(from:)
func (p NSPersistentStore) WillRemoveFromPersistentStoreCoordinator(coordinator INSPersistentStoreCoordinator) {
	objc.Send[objc.ID](p.ID, objc.Sel("willRemoveFromPersistentStoreCoordinator:"), coordinator)
}

// Returns the metadata from the persistent store at the given URL.
//
// url: The location of the store.
//
// # Return Value
//
// The metadata from the persistent store at `url`. Returns `nil` if there is
// an error.
//
// # Discussion
//
// Subclasses must override this method.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStore/metadataForPersistentStore(with:)
func (_NSPersistentStoreClass NSPersistentStoreClass) MetadataForPersistentStoreWithURLError(url foundation.NSURL) (foundation.INSDictionary, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentStoreClass.class), objc.Sel("metadataForPersistentStoreWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDictionaryFromID(rv), nil

}

// Sets the metadata for the store at a given URL.
//
// metadata: The metadata for the store at `url`.
//
// url: The location of the store.
//
// # Discussion
//
// Subclasses must override this method to set metadata appropriately.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStore/setMetadata(_:forPersistentStoreAt:)
func (_NSPersistentStoreClass NSPersistentStoreClass) SetMetadataForPersistentStoreWithURLError(metadata foundation.INSDictionary, url foundation.NSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_NSPersistentStoreClass.class), objc.Sel("setMetadata:forPersistentStoreWithURL:error:"), metadata, url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setMetadata:forPersistentStoreWithURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns the migration manager class for this store class.
//
// # Return Value
//
// The [NSMigrationManager] class for this store class
//
// # Discussion
//
// In a subclass of [NSPersistentStore], you can override this to provide a
// custom migration manager subclass (for example, to take advantage of
// store-specific functionality to improve migration performance).
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStore/migrationManagerClass()
func (_NSPersistentStoreClass NSPersistentStoreClass) MigrationManagerClass() objectivec.Class {
	rv := objc.Send[objectivec.Class](objc.ID(_NSPersistentStoreClass.class), objc.Sel("migrationManagerClass"))
	return objectivec.Class(rv)
}

// The name of the managed object model configuration that creates the
// persistent store.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStore/configurationName
func (p NSPersistentStore) ConfigurationName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("configurationName"))
	return foundation.NSStringFromID(rv).String()
}

// The options that Core Data uses to create the store.
//
// # Discussion
//
// See [NSPersistentStoreCoordinator] for a list of key names for options in
// this dictionary.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStore/options
func (p NSPersistentStore) Options() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("options"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The persistent store coordinator that loads the persistent store.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStore/persistentStoreCoordinator
func (p NSPersistentStore) PersistentStoreCoordinator() INSPersistentStoreCoordinator {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("persistentStoreCoordinator"))
	return NSPersistentStoreCoordinatorFromID(objc.ID(rv))
}

// The type string of the persistent store.
//
// # Discussion
//
// This string is used when specifying the type of store to add to a
// persistent store coordinator.
//
// # Special Considerations
//
// Subclasses must override this method to provide a unique type.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStore/type
func (p NSPersistentStore) Type() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("type"))
	return foundation.NSStringFromID(rv).String()
}

// The unique identifier for the persistent store.
//
// # Discussion
//
// The identifier is used as part of the managed object IDs for each object in
// the store.
//
// # Special Considerations
//
// [NSPersistentStore] provides a default implementation to provide a globally
// unique identifier for the store instance.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStore/identifier
func (p NSPersistentStore) Identifier() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}
func (p NSPersistentStore) SetIdentifier(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setIdentifier:"), objc.String(value))
}

// A Boolean value that indicates whether the persistent store is read-only.
//
// # Discussion
//
// true if the receiver is read-only, otherwise false.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStore/isReadOnly
func (p NSPersistentStore) IsReadOnly() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isReadOnly"))
	return rv
}
func (p NSPersistentStore) SetReadOnly(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setReadOnly:"), value)
}

// The URL for the persistent store.
//
// # Discussion
//
// To alter the location of a store, send the persistent store coordinator a
// [NSPersistentStoreCoordinator.SetURLForPersistentStore] message.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStore/url
func (p NSPersistentStore) URL() foundation.NSURL {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (p NSPersistentStore) SetURL(value foundation.NSURL) {
	objc.Send[struct{}](p.ID, objc.Sel("setURL:"), value)
}

// The metadata for the persistent store.
//
// # Discussion
//
// The dictionary must include the store type ([NSStoreTypeKey]) and UUID
// ([NSStoreUUIDKey]).
//
// # Special Considerations
//
// Subclasses must override this property to provide storage and persistence
// for the store metadata.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStore/metadata
//
// [NSStoreTypeKey]: https://developer.apple.com/documentation/CoreData/NSStoreTypeKey
// [NSStoreUUIDKey]: https://developer.apple.com/documentation/CoreData/NSStoreUUIDKey
func (p NSPersistentStore) Metadata() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("metadata"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (p NSPersistentStore) SetMetadata(value foundation.INSDictionary) {
	objc.Send[struct{}](p.ID, objc.Sel("setMetadata:"), value)
}

// The spotlight exporter associated with this persistent store.
//
// # Discussion
//
// Spotlight support isn’t available in a compatible iPad or iPhone app
// running in visionOS.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStore/coreSpotlightExporter
func (p NSPersistentStore) CoreSpotlightExporter() INSCoreDataCoreSpotlightDelegate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("coreSpotlightExporter"))
	return NSCoreDataCoreSpotlightDelegateFromID(objc.ID(rv))
}
