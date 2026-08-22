// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPersistentStoreDescription] class.
var (
	_NSPersistentStoreDescriptionClass     NSPersistentStoreDescriptionClass
	_NSPersistentStoreDescriptionClassOnce sync.Once
)

func getNSPersistentStoreDescriptionClass() NSPersistentStoreDescriptionClass {
	_NSPersistentStoreDescriptionClassOnce.Do(func() {
		_NSPersistentStoreDescriptionClass = NSPersistentStoreDescriptionClass{class: objc.GetClass("NSPersistentStoreDescription")}
	})
	return _NSPersistentStoreDescriptionClass
}

// GetNSPersistentStoreDescriptionClass returns the class object for NSPersistentStoreDescription.
func GetNSPersistentStoreDescriptionClass() NSPersistentStoreDescriptionClass {
	return getNSPersistentStoreDescriptionClass()
}

type NSPersistentStoreDescriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPersistentStoreDescriptionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPersistentStoreDescriptionClass) Alloc() NSPersistentStoreDescription {
	rv := objc.Send[NSPersistentStoreDescription](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A description object used to create and load a persistent store.
//
// # Creating a Persistent Store Description
//
//   - [NSPersistentStoreDescription.InitWithURL]: Initializes the receiver with a URL for the store.
//
// # Configuring a Persistent Store Description
//
//   - [NSPersistentStoreDescription.URL]: The URL that the store will use for its location.
//   - [NSPersistentStoreDescription.SetURL]
//   - [NSPersistentStoreDescription.Configuration]: The name of the configuration used by this store.
//   - [NSPersistentStoreDescription.SetConfiguration]
//   - [NSPersistentStoreDescription.Timeout]: The connection timeout for the associated store.
//   - [NSPersistentStoreDescription.SetTimeout]
//   - [NSPersistentStoreDescription.Type]: The type of store this description represents.
//   - [NSPersistentStoreDescription.SetType]
//   - [NSPersistentStoreDescription.IsReadOnly]: A flag that indicates whether this store will be read-only.
//   - [NSPersistentStoreDescription.SetReadOnly]
//   - [NSPersistentStoreDescription.ShouldAddStoreAsynchronously]: A flag that determines whether the store is added asynchronously.
//   - [NSPersistentStoreDescription.SetShouldAddStoreAsynchronously]
//   - [NSPersistentStoreDescription.ShouldInferMappingModelAutomatically]: A flag indicating whether a mapping model should be created automatically.
//   - [NSPersistentStoreDescription.SetShouldInferMappingModelAutomatically]
//   - [NSPersistentStoreDescription.ShouldMigrateStoreAutomatically]: A flag indicating whether the associated persistent store should be migrated automatically.
//   - [NSPersistentStoreDescription.SetShouldMigrateStoreAutomatically]
//   - [NSPersistentStoreDescription.SetOptionForKey]: Sets an option on the store.
//   - [NSPersistentStoreDescription.SetValueForPragmaNamed]: Allows you to set pragmas for the SQLite store.
//
// # Accessing the Configuration Options
//
//   - [NSPersistentStoreDescription.Options]: A dictionary representation of the options set on the associated persistent store.
//   - [NSPersistentStoreDescription.SqlitePragmas]: The SQLite pragmas set for the associated persistent store. (read-only)
//
// # Syncing to CloudKit
//
//   - [NSPersistentStoreDescription.CloudKitContainerOptions]: Options that customize how this store description aligns with a CloudKit database.
//   - [NSPersistentStoreDescription.SetCloudKitContainerOptions]
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription
type NSPersistentStoreDescription struct {
	objectivec.Object
}

// NSPersistentStoreDescriptionFromID constructs a [NSPersistentStoreDescription] from an objc.ID.
//
// A description object used to create and load a persistent store.
func NSPersistentStoreDescriptionFromID(id objc.ID) NSPersistentStoreDescription {
	return NSPersistentStoreDescription{objectivec.Object{ID: id}}
}

// NOTE: NSPersistentStoreDescription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPersistentStoreDescription] class.
//
// # Creating a Persistent Store Description
//
//   - [INSPersistentStoreDescription.InitWithURL]: Initializes the receiver with a URL for the store.
//
// # Configuring a Persistent Store Description
//
//   - [INSPersistentStoreDescription.URL]: The URL that the store will use for its location.
//   - [INSPersistentStoreDescription.SetURL]
//   - [INSPersistentStoreDescription.Configuration]: The name of the configuration used by this store.
//   - [INSPersistentStoreDescription.SetConfiguration]
//   - [INSPersistentStoreDescription.Timeout]: The connection timeout for the associated store.
//   - [INSPersistentStoreDescription.SetTimeout]
//   - [INSPersistentStoreDescription.Type]: The type of store this description represents.
//   - [INSPersistentStoreDescription.SetType]
//   - [INSPersistentStoreDescription.IsReadOnly]: A flag that indicates whether this store will be read-only.
//   - [INSPersistentStoreDescription.SetReadOnly]
//   - [INSPersistentStoreDescription.ShouldAddStoreAsynchronously]: A flag that determines whether the store is added asynchronously.
//   - [INSPersistentStoreDescription.SetShouldAddStoreAsynchronously]
//   - [INSPersistentStoreDescription.ShouldInferMappingModelAutomatically]: A flag indicating whether a mapping model should be created automatically.
//   - [INSPersistentStoreDescription.SetShouldInferMappingModelAutomatically]
//   - [INSPersistentStoreDescription.ShouldMigrateStoreAutomatically]: A flag indicating whether the associated persistent store should be migrated automatically.
//   - [INSPersistentStoreDescription.SetShouldMigrateStoreAutomatically]
//   - [INSPersistentStoreDescription.SetOptionForKey]: Sets an option on the store.
//   - [INSPersistentStoreDescription.SetValueForPragmaNamed]: Allows you to set pragmas for the SQLite store.
//
// # Accessing the Configuration Options
//
//   - [INSPersistentStoreDescription.Options]: A dictionary representation of the options set on the associated persistent store.
//   - [INSPersistentStoreDescription.SqlitePragmas]: The SQLite pragmas set for the associated persistent store. (read-only)
//
// # Syncing to CloudKit
//
//   - [INSPersistentStoreDescription.CloudKitContainerOptions]: Options that customize how this store description aligns with a CloudKit database.
//   - [INSPersistentStoreDescription.SetCloudKitContainerOptions]
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription
type INSPersistentStoreDescription interface {
	objectivec.IObject

	// Topic: Creating a Persistent Store Description

	// Initializes the receiver with a URL for the store.
	InitWithURL(url foundation.NSURL) NSPersistentStoreDescription

	// Topic: Configuring a Persistent Store Description

	// The URL that the store will use for its location.
	URL() foundation.NSURL
	SetURL(value foundation.NSURL)
	// The name of the configuration used by this store.
	Configuration() string
	SetConfiguration(value string)
	// The connection timeout for the associated store.
	Timeout() foundation.NSTimeInterval
	SetTimeout(value foundation.NSTimeInterval)
	// The type of store this description represents.
	Type() string
	SetType(value string)
	// A flag that indicates whether this store will be read-only.
	IsReadOnly() bool
	SetReadOnly(value bool)
	// A flag that determines whether the store is added asynchronously.
	ShouldAddStoreAsynchronously() bool
	SetShouldAddStoreAsynchronously(value bool)
	// A flag indicating whether a mapping model should be created automatically.
	ShouldInferMappingModelAutomatically() bool
	SetShouldInferMappingModelAutomatically(value bool)
	// A flag indicating whether the associated persistent store should be migrated automatically.
	ShouldMigrateStoreAutomatically() bool
	SetShouldMigrateStoreAutomatically(value bool)
	// Sets an option on the store.
	SetOptionForKey(option objectivec.NSObject, key string)
	// Allows you to set pragmas for the SQLite store.
	SetValueForPragmaNamed(value objectivec.NSObject, name string)

	// Topic: Accessing the Configuration Options

	// A dictionary representation of the options set on the associated persistent store.
	Options() foundation.INSDictionary
	// The SQLite pragmas set for the associated persistent store. (read-only)
	SqlitePragmas() foundation.INSDictionary

	// Topic: Syncing to CloudKit

	// Options that customize how this store description aligns with a CloudKit database.
	CloudKitContainerOptions() INSPersistentCloudKitContainerOptions
	SetCloudKitContainerOptions(value INSPersistentCloudKitContainerOptions)
}

// Init initializes the instance.
func (p NSPersistentStoreDescription) Init() NSPersistentStoreDescription {
	rv := objc.Send[NSPersistentStoreDescription](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPersistentStoreDescription) Autorelease() NSPersistentStoreDescription {
	rv := objc.Send[NSPersistentStoreDescription](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPersistentStoreDescription creates a new NSPersistentStoreDescription instance.
func NewNSPersistentStoreDescription() NSPersistentStoreDescription {
	class := getNSPersistentStoreDescriptionClass()
	rv := objc.Send[NSPersistentStoreDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes the receiver with a URL for the store.
//
// url: Location for the store.
//
// # Return Value
//
// Initialized [NSPersistentStoreDescription] configured with the given URL.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/init(url:)
func NewPersistentStoreDescriptionWithURL(url foundation.NSURL) NSPersistentStoreDescription {
	instance := getNSPersistentStoreDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), url)
	return NSPersistentStoreDescriptionFromID(rv)
}

// Initializes the receiver with a URL for the store.
//
// url: Location for the store.
//
// # Return Value
//
// Initialized [NSPersistentStoreDescription] configured with the given URL.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/init(url:)
func (p NSPersistentStoreDescription) InitWithURL(url foundation.NSURL) NSPersistentStoreDescription {
	rv := objc.Send[NSPersistentStoreDescription](p.ID, objc.Sel("initWithURL:"), url)
	return rv
}

// Sets an option on the store.
//
// option: The value to be set for an option on the store.
//
// key: The key of the value to be set for an option on the store.
//
// # Discussion
//
// If a value was previously set for the given option, that value is replaced
// with the given value. Note that the keys are case-sensitive. For a list of
// the available options, see [NSPersistentStoreCoordinator].
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/setOption(_:forKey:)
func (p NSPersistentStoreDescription) SetOptionForKey(option objectivec.NSObject, key string) {
	objc.Send[objc.ID](p.ID, objc.Sel("setOption:forKey:"), option, objc.String(key))
}

// Allows you to set pragmas for the SQLite store.
//
// value: The value of the pragma to be set.
//
// name: The name of the pragma to be set.
//
// # Discussion
//
// Pragma options are for SQLite stores only. All pragma values must be
// specified as [NSString]objects. The `fullfsync` and `synchronous` pragmas
// control the tradeoff between write performance (write to disk speed and
// cache utilization) and durability (data loss/corruption sensitivity to
// power interruption). For more information on pragma settings, see
// [http://sqlite.org/pragma.html].
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/setValue(_:forPragmaNamed:)
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [http://sqlite.org/pragma.html]: http://sqlite.org/pragma.html
func (p NSPersistentStoreDescription) SetValueForPragmaNamed(value objectivec.NSObject, name string) {
	objc.Send[objc.ID](p.ID, objc.Sel("setValue:forPragmaNamed:"), value, objc.String(name))
}

// The URL that the store will use for its location.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/url
func (p NSPersistentStoreDescription) URL() foundation.NSURL {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (p NSPersistentStoreDescription) SetURL(value foundation.NSURL) {
	objc.Send[struct{}](p.ID, objc.Sel("setURL:"), value)
}

// The name of the configuration used by this store.
//
// # Discussion
//
// This displays the name of a configuration in the receiver’s managed
// object model that will be used by the new store. The configuration can be
// `nil`, in which case no other configurations are allowed.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/configuration
func (p NSPersistentStoreDescription) Configuration() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("configuration"))
	return foundation.NSStringFromID(rv).String()
}
func (p NSPersistentStoreDescription) SetConfiguration(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setConfiguration:"), objc.String(value))
}

// The connection timeout for the associated store.
//
// # Discussion
//
// This is a convenience method for setting the
// [NSPersistentStoreTimeoutOption] on the associated store.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/timeout
//
// [NSPersistentStoreTimeoutOption]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreTimeoutOption
func (p NSPersistentStoreDescription) Timeout() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](p.ID, objc.Sel("timeout"))
	return foundation.NSTimeInterval(rv)
}
func (p NSPersistentStoreDescription) SetTimeout(value foundation.NSTimeInterval) {
	objc.Send[struct{}](p.ID, objc.Sel("setTimeout:"), value)
}

// The type of store this description represents.
//
// # Discussion
//
// A string constant (such as [NSSQLiteStoreType]) that specifies the type of
// the new store—see [NSPersistentStoreCoordinator].
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/type
func (p NSPersistentStoreDescription) Type() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("type"))
	return foundation.NSStringFromID(rv).String()
}
func (p NSPersistentStoreDescription) SetType(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setType:"), objc.String(value))
}

// A flag that indicates whether this store will be read-only.
//
// # Discussion
//
// This is a convenience method for setting the
// [NSReadOnlyPersistentStoreOption] on the associated store.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/isReadOnly
//
// [NSReadOnlyPersistentStoreOption]: https://developer.apple.com/documentation/CoreData/NSReadOnlyPersistentStoreOption
func (p NSPersistentStoreDescription) IsReadOnly() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isReadOnly"))
	return rv
}
func (p NSPersistentStoreDescription) SetReadOnly(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setReadOnly:"), value)
}

// A flag that determines whether the store is added asynchronously.
//
// # Discussion
//
// By default, the store is added to the [NSPersistentStoreCoordinator]
// synchronously on the calling thread. If this flag is set to true, the store
// is added asynchronously on a background queue. The default for this flag is
// false.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/shouldAddStoreAsynchronously
func (p NSPersistentStoreDescription) ShouldAddStoreAsynchronously() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("shouldAddStoreAsynchronously"))
	return rv
}
func (p NSPersistentStoreDescription) SetShouldAddStoreAsynchronously(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setShouldAddStoreAsynchronously:"), value)
}

// A flag indicating whether a mapping model should be created automatically.
//
// # Discussion
//
// If this flag is set to true and the value of the
// [NSPersistentStoreDescription.ShouldMigrateStoreAutomatically] is true, the
// coordinator attempts to infer a mapping model if none can be found. The
// default for this flag is true.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/shouldInferMappingModelAutomatically
func (p NSPersistentStoreDescription) ShouldInferMappingModelAutomatically() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("shouldInferMappingModelAutomatically"))
	return rv
}
func (p NSPersistentStoreDescription) SetShouldInferMappingModelAutomatically(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setShouldInferMappingModelAutomatically:"), value)
}

// A flag indicating whether the associated persistent store should be
// migrated automatically.
//
// # Discussion
//
// If this is set to false and the store is out of sync, attempting to load
// the store produces an error. If this is set to true and the store is out of
// sync, attempting to load the store causes Core Data to attempt a migration.
// This flag is set to true by default.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/shouldMigrateStoreAutomatically
func (p NSPersistentStoreDescription) ShouldMigrateStoreAutomatically() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("shouldMigrateStoreAutomatically"))
	return rv
}
func (p NSPersistentStoreDescription) SetShouldMigrateStoreAutomatically(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setShouldMigrateStoreAutomatically:"), value)
}

// A dictionary representation of the options set on the associated persistent
// store.
//
// # Discussion
//
// A dictionary containing key-value pairs that specify numerous settings for
// the persistent store. For key definitions, see
// [NSPersistentStoreCoordinator].
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/options
func (p NSPersistentStoreDescription) Options() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("options"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The SQLite pragmas set for the associated persistent store. (read-only)
//
// # Discussion
//
// This property contains all of the pragmas set on the associated persistent
// store. This property is only relevant when the
// [NSPersistentStoreDescription.Type] is set to [NSSQLiteStoreType].
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/sqlitePragmas
//
// [NSSQLiteStoreType]: https://developer.apple.com/documentation/CoreData/NSSQLiteStoreType
func (p NSPersistentStoreDescription) SqlitePragmas() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("sqlitePragmas"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// Options that customize how this store description aligns with a CloudKit
// database.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/cloudKitContainerOptions
func (p NSPersistentStoreDescription) CloudKitContainerOptions() INSPersistentCloudKitContainerOptions {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("cloudKitContainerOptions"))
	return NSPersistentCloudKitContainerOptionsFromID(objc.ID(rv))
}
func (p NSPersistentStoreDescription) SetCloudKitContainerOptions(value INSPersistentCloudKitContainerOptions) {
	objc.Send[struct{}](p.ID, objc.Sel("setCloudKitContainerOptions:"), value)
}
