// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKSyncEngineConfiguration] class.
var (
	_CKSyncEngineConfigurationClass     CKSyncEngineConfigurationClass
	_CKSyncEngineConfigurationClassOnce sync.Once
)

func getCKSyncEngineConfigurationClass() CKSyncEngineConfigurationClass {
	_CKSyncEngineConfigurationClassOnce.Do(func() {
		_CKSyncEngineConfigurationClass = CKSyncEngineConfigurationClass{class: objc.GetClass("CKSyncEngineConfiguration")}
	})
	return _CKSyncEngineConfigurationClass
}

// GetCKSyncEngineConfigurationClass returns the class object for CKSyncEngineConfiguration.
func GetCKSyncEngineConfigurationClass() CKSyncEngineConfigurationClass {
	return getCKSyncEngineConfigurationClass()
}

type CKSyncEngineConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineConfigurationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineConfigurationClass) Alloc() CKSyncEngineConfiguration {
	rv := objc.Send[CKSyncEngineConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A type that configures the attributes and behavior of a sync engine.
//
// # Creating configurations
//
//   - [CKSyncEngineConfiguration.InitWithDatabaseStateSerializationDelegate]: Creates a configuration for the specified database and serialized state.
//
// # Handling record changes
//
//   - [CKSyncEngineConfiguration.Delegate]: The object that provides the records to sync and handles any related events.
//   - [CKSyncEngineConfiguration.SetDelegate]
//
// # Managing attributes
//
//   - [CKSyncEngineConfiguration.AutomaticallySync]: A Boolean value that determines whether the engine syncs automatically.
//   - [CKSyncEngineConfiguration.SetAutomaticallySync]
//   - [CKSyncEngineConfiguration.Database]: The associated database.
//   - [CKSyncEngineConfiguration.SetDatabase]
//   - [CKSyncEngineConfiguration.SubscriptionID]: The subscription identifier for the associated database.
//   - [CKSyncEngineConfiguration.SetSubscriptionID]
//   - [CKSyncEngineConfiguration.StateSerialization]: The sync engine’s serialized state.
//   - [CKSyncEngineConfiguration.SetStateSerialization]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration
type CKSyncEngineConfiguration struct {
	objectivec.Object
}

// CKSyncEngineConfigurationFromID constructs a [CKSyncEngineConfiguration] from an objc.ID.
//
// A type that configures the attributes and behavior of a sync engine.
func CKSyncEngineConfigurationFromID(id objc.ID) CKSyncEngineConfiguration {
	return CKSyncEngineConfiguration{objectivec.Object{ID: id}}
}

// Ensure CKSyncEngineConfiguration implements ICKSyncEngineConfiguration.
var _ ICKSyncEngineConfiguration = CKSyncEngineConfiguration{}

// An interface definition for the [CKSyncEngineConfiguration] class.
//
// # Creating configurations
//
//   - [ICKSyncEngineConfiguration.InitWithDatabaseStateSerializationDelegate]: Creates a configuration for the specified database and serialized state.
//
// # Handling record changes
//
//   - [ICKSyncEngineConfiguration.Delegate]: The object that provides the records to sync and handles any related events.
//   - [ICKSyncEngineConfiguration.SetDelegate]
//
// # Managing attributes
//
//   - [ICKSyncEngineConfiguration.AutomaticallySync]: A Boolean value that determines whether the engine syncs automatically.
//   - [ICKSyncEngineConfiguration.SetAutomaticallySync]
//   - [ICKSyncEngineConfiguration.Database]: The associated database.
//   - [ICKSyncEngineConfiguration.SetDatabase]
//   - [ICKSyncEngineConfiguration.SubscriptionID]: The subscription identifier for the associated database.
//   - [ICKSyncEngineConfiguration.SetSubscriptionID]
//   - [ICKSyncEngineConfiguration.StateSerialization]: The sync engine’s serialized state.
//   - [ICKSyncEngineConfiguration.SetStateSerialization]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration
type ICKSyncEngineConfiguration interface {
	objectivec.IObject

	// Topic: Creating configurations

	// Creates a configuration for the specified database and serialized state.
	InitWithDatabaseStateSerializationDelegate(database ICKDatabase, stateSerialization ICKSyncEngineStateSerialization, delegate CKSyncEngineDelegate) CKSyncEngineConfiguration

	// Topic: Handling record changes

	// The object that provides the records to sync and handles any related events.
	Delegate() CKSyncEngineDelegate
	SetDelegate(value CKSyncEngineDelegate)

	// Topic: Managing attributes

	// A Boolean value that determines whether the engine syncs automatically.
	AutomaticallySync() bool
	SetAutomaticallySync(value bool)
	// The associated database.
	Database() ICKDatabase
	SetDatabase(value ICKDatabase)
	// The subscription identifier for the associated database.
	SubscriptionID() CKSubscriptionID
	SetSubscriptionID(value CKSubscriptionID)
	// The sync engine’s serialized state.
	StateSerialization() ICKSyncEngineStateSerialization
	SetStateSerialization(value ICKSyncEngineStateSerialization)
}

// Init initializes the instance.
func (c CKSyncEngineConfiguration) Init() CKSyncEngineConfiguration {
	rv := objc.Send[CKSyncEngineConfiguration](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineConfiguration) Autorelease() CKSyncEngineConfiguration {
	rv := objc.Send[CKSyncEngineConfiguration](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineConfiguration creates a new CKSyncEngineConfiguration instance.
func NewCKSyncEngineConfiguration() CKSyncEngineConfiguration {
	class := getCKSyncEngineConfigurationClass()
	rv := objc.Send[CKSyncEngineConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a configuration for the specified database and serialized state.
//
// database: The database to sync — either a person’s private database or their
// shared database.
//
// stateSerialization: If this is the first initialization of the associated sync engine, specify
// `nil`; otherwise, specify the state from the most recent
// [CKSyncEngineStateUpdateEvent] that your delegate handled.
//
// delegate: The object that provides the records to sync and handles any related
// events.
//
// # Return Value
//
// An initialized configuration instance, or `nil` if CloudKit can’t create
// one.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration/initWithDatabase:stateSerialization:delegate:
func NewCKSyncEngineConfigurationWithDatabaseStateSerializationDelegate(database ICKDatabase, stateSerialization ICKSyncEngineStateSerialization, delegate CKSyncEngineDelegate) CKSyncEngineConfiguration {
	instance := getCKSyncEngineConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDatabase:stateSerialization:delegate:"), database, stateSerialization, delegate)
	return CKSyncEngineConfigurationFromID(rv)
}

// Creates a configuration for the specified database and serialized state.
//
// database: The database to sync — either a person’s private database or their
// shared database.
//
// stateSerialization: If this is the first initialization of the associated sync engine, specify
// `nil`; otherwise, specify the state from the most recent
// [CKSyncEngineStateUpdateEvent] that your delegate handled.
//
// delegate: The object that provides the records to sync and handles any related
// events.
//
// # Return Value
//
// An initialized configuration instance, or `nil` if CloudKit can’t create
// one.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration/initWithDatabase:stateSerialization:delegate:
func (c CKSyncEngineConfiguration) InitWithDatabaseStateSerializationDelegate(database ICKDatabase, stateSerialization ICKSyncEngineStateSerialization, delegate CKSyncEngineDelegate) CKSyncEngineConfiguration {
	rv := objc.Send[CKSyncEngineConfiguration](c.ID, objc.Sel("initWithDatabase:stateSerialization:delegate:"), database, stateSerialization, delegate)
	return rv
}

// The object that provides the records to sync and handles any related
// events.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration/delegate
func (c CKSyncEngineConfiguration) Delegate() CKSyncEngineDelegate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("delegate"))
	return CKSyncEngineDelegateObjectFromID(rv)
}
func (c CKSyncEngineConfiguration) SetDelegate(value CKSyncEngineDelegate) {
	objc.Send[struct{}](c.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that determines whether the engine syncs automatically.
//
// # Discussion
//
// By default, the sync engine uses the system scheduler to automatically
// schedule both send and fetch operations. If an operation fails due to a
// recoverable error, such as a network failure or when the server is
// enforcing request limits, the engine reschedules those operations as
// necessary. Unless you have a specific need, prefer to use the default
// behavior in your app.
//
// If you set this property’s value to false, use
// [FetchChangesWithCompletionHandler] and [SendChangesWithCompletionHandler]
// to invoke immediate sync operations, allowing for more control over when
// your app syncs its records. For example, you may want to sync at a specific
// time of day or deterministically simulate certain conditions in your unit
// tests.
//
// The default value is true.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration/automaticallySync
func (c CKSyncEngineConfiguration) AutomaticallySync() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("automaticallySync"))
	return rv
}
func (c CKSyncEngineConfiguration) SetAutomaticallySync(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAutomaticallySync:"), value)
}

// The associated database.
//
// # Discussion
//
// Multiple sync engines can run in the same process, each targeting a
// different database. For example, you may use one sync engine for a
// person’s private database and another for their shared database.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration/database
func (c CKSyncEngineConfiguration) Database() ICKDatabase {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("database"))
	return CKDatabaseFromID(objc.ID(rv))
}
func (c CKSyncEngineConfiguration) SetDatabase(value ICKDatabase) {
	objc.Send[struct{}](c.ID, objc.Sel("setDatabase:"), value)
}

// The subscription identifier for the associated database.
//
// # Discussion
//
// By default, a sync engine attempts to discover an existing subscription for
// the synced database. If one isn’t found, the engine creates an internal
// [CKDatabaseSubscription] and uses that to receive notifications about
// remote record changes.
//
// If you require the sync engine to use a specific database subscription,
// assign that subscription’s identifier to this property. Doing so enables
// your app to be backwards compatible if you’re migrating to [CKSyncEngine]
// from a custom CloudKit sync implementation.
//
// The default value is `nil`.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration/subscriptionID
//
// [CKSyncEngine]: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-5sie5
func (c CKSyncEngineConfiguration) SubscriptionID() CKSubscriptionID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("subscriptionID"))
	return CKSubscriptionID(foundation.NSStringFromID(rv).String())
}
func (c CKSyncEngineConfiguration) SetSubscriptionID(value CKSubscriptionID) {
	objc.Send[struct{}](c.ID, objc.Sel("setSubscriptionID:"), objc.String(string(value)))
}

// The sync engine’s serialized state.
//
// # Discussion
//
// This property returns the value you specify for the initializer’s
// `stateSerialization` parameter. If you choose to set this property after
// initialization, assign the state from the most recent
// [CKSyncEngineStateUpdateEvent] handled by your delegate. However, If this
// is the first initialization of the associated sync engine, specify `nil`
// instead.
//
// The default value is `nil`.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration/stateSerialization
func (c CKSyncEngineConfiguration) StateSerialization() ICKSyncEngineStateSerialization {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("stateSerialization"))
	return CKSyncEngineStateSerializationFromID(objc.ID(rv))
}
func (c CKSyncEngineConfiguration) SetStateSerialization(value ICKSyncEngineStateSerialization) {
	objc.Send[struct{}](c.ID, objc.Sel("setStateSerialization:"), value)
}
