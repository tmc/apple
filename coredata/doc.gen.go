// Code generated from Apple documentation for CoreData. DO NOT EDIT.

// Package coredata provides Go bindings for the CoreData framework.
//
// Persist or cache data on a single device, or sync data to multiple devices
// with CloudKit.
//
// Use Core Data to save your application’s permanent data for offline use,
// to cache temporary data, and to add undo functionality to your app on a
// single device. To sync data across multiple devices in a single iCloud
// account, Core Data automatically mirrors your schema to a CloudKit
// container.
//
// # Essentials
//
//   - [Creating a Core Data model]: Define your app’s object structure with a data model file.
//   - [Setting up a Core Data stack]: Set up the classes that manage and persist your app’s objects.
//   - [Core Data stack]: Manage and persist your app’s model layer. ([NSPersistentContainer], [NSManagedObjectModel], [NSEntityDescription], [NSPropertyDescription], [NSAttributeDescription])
//   - [Handling Different Data Types in Core Data]: Create, store, and present records for a variety of data types.
//   - [Linking Data Between Two Core Data Stores]: Organize data in two different stores and implement a link between them.
//
// # Data modeling
//
//   - [Modeling data]: Configure the data model file to contain your app’s object graph.
//   - [Core Data model]: Describe your app’s object structure. ([NSManagedObject], [NSEntityDescription], [NSPropertyDescription], [NSAttributeDescription], [NSAttributeType])
//
// # Fetch requests
//
//   - [NSFetchRequest]: A description of search criteria used to retrieve data from a persistent store. ([NSFetchRequestResultType], [NSFetchRequestExpression], [NSExpressionDescription], [NSFetchedPropertyDescription], [NSFetchRequestResultType])
//   - [NSAsynchronousFetchRequest]: A fetch request that retrieves results asynchronously and supports progress notification.
//   - [NSAsynchronousFetchResult]: A fetch result object that encompasses the response from an executed asynchronous fetch request.
//   - [NSFetchedResultsController]: A controller that you use to manage the results of a Core Data fetch request and to display data to the user. ([NSFetchedResultsControllerDelegate], [NSFetchedResultsSectionInfo], [NSFetchRequestResultType], [NSFetchedResultsChangeType])
//
// # SwiftData migration and coexistence
//
//   - [Adopting SwiftData for a Core Data app]: Persist data in your app intuitively with the Swift native persistence framework.
//
// # CloudKit mirroring
//
//   - [Mirroring a Core Data store with CloudKit]: Back user interfaces with a local replica of a CloudKit private database.
//   - [Synchronizing a local store to the cloud]: Share data between a user’s devices and other iCloud users.
//   - [NSPersistentCloudKitContainer]: A container that encapsulates the Core Data stack in your app, and mirrors select persistent stores to a CloudKit private database. ([NSPersistentCloudKitContainerSchemaInitializationOptions], [NSPersistentCloudKitContainerEventRequest], [NSPersistentCloudKitContainerEventResult])
//   - [NSPersistentCloudKitContainerOptions]: An object that customizes how a store description aligns with a CloudKit database.
//   - [Sharing Core Data objects between iCloud users]: Use Core Data and CloudKit to synchronize data between devices of an iCloud user and share data between different iCloud users.
//
// # Change processing
//
//   - [Accessing data when the store changes]: Guarantee that a context won’t see store changes until you tell it to look.
//   - [Consuming relevant store changes]: Filter store transactions for changes relevant to the current view.
//   - [Persistent history]: Use persistent history tracking to determine what changes have occurred in the store since the enabling of persistent history tracking. ([NSPersistentHistoryToken], [NSPersistentHistoryChangeRequest], [NSPersistentHistoryResult], [NSPersistentHistoryTransaction], [NSPersistentHistoryChange])
//
// # Background tasks
//
//   - [Using Core Data in the background]: Use Core Data in both a single-threaded and multithreaded app.
//   - [Loading and displaying a large data feed]: Consume data in the background, and lower memory use by batching imports and preventing duplicate records.
//   - [Conflict resolution]: Detect and resolve conflicts that occur when data is changed on multiple threads. ([NSConstraintConflict], [NSMergeConflict], [NSMergePolicy], [NSQueryGenerationToken])
//   - [Batch processing]: Use batch processes to manage large data changes. ([NSBatchInsertRequest], [NSBatchInsertResult], [NSBatchUpdateRequest], [NSBatchUpdateResult], [NSBatchDeleteRequest])
//
// # Data model migration
//
//   - [Migrating your data model automatically]: Enable lightweight migrations to keep your data model and the underlying data in a consistent state.
//   - [Staged migrations]: Migrate complex data models containing changes that are incompatible with lightweight migrations. ([NSStagedMigrationManager], [NSLightweightMigrationStage], [NSCustomMigrationStage], [NSMigrationStage])
//   - [Manual migrations]: Migrate elaborate data models with changes that go beyond the capabilities of both lightweight and staged migrations. ([NSMigrationManager], [NSMappingModel], [NSEntityMapping], [NSEntityMigrationPolicy], [NSEntityMappingType])
//
// # Related types
//
//   - [Core Data Constants]: Keys to use with persistent stores and notifications from Core Data.//
//
// # Key Types
//
//   - [NSManagedObjectContext] - An object space to manipulate and track changes to managed objects.
//   - [NSManagedObject] - The base class that all Core Data model objects inherit from.
//   - [NSPersistentStoreCoordinator] - An object that enables an app’s contexts and the underlying persistent stores to work together.
//   - [NSEntityDescription] - A description of a Core Data entity.
//   - [NSFetchRequest] - A description of search criteria used to retrieve data from a persistent store.
//   - [NSManagedObjectModel] - A programmatic representation of the `XCUIElementTypeXcdatamodeld` file describing your objects.
//   - [NSMigrationManager] - A migration manager instance that performs a migration of data from one persistent store to another using a given mapping model.
//   - [NSPersistentStore] - The abstract base class for all Core Data persistent stores.
//   - [NSFetchedResultsController] - A controller that you use to manage the results of a Core Data fetch request and to display data to the user.
//   - [NSPersistentStoreDescription] - A description object used to create and load a persistent store.
//
// [Accessing data when the store changes]: https://developer.apple.com/documentation/coredata/accessing-data-when-the-store-changes
// [Adopting SwiftData for a Core Data app]: https://developer.apple.com/documentation/coredata/adopting-swiftdata-for-a-core-data-app
// [Batch processing]: https://developer.apple.com/documentation/coredata/batch-processing
// [Conflict resolution]: https://developer.apple.com/documentation/coredata/conflict-resolution
// [Consuming relevant store changes]: https://developer.apple.com/documentation/coredata/consuming-relevant-store-changes
// [Core Data Constants]: https://developer.apple.com/documentation/coredata/core-data-constants
// [Core Data model]: https://developer.apple.com/documentation/coredata/core-data-model
// [Core Data stack]: https://developer.apple.com/documentation/coredata/core-data-stack
// [Creating a Core Data model]: https://developer.apple.com/documentation/coredata/creating-a-core-data-model
// [Handling Different Data Types in Core Data]: https://developer.apple.com/documentation/coredata/handling-different-data-types-in-core-data
// [Linking Data Between Two Core Data Stores]: https://developer.apple.com/documentation/coredata/linking-data-between-two-core-data-stores
// [Loading and displaying a large data feed]: https://developer.apple.com/documentation/SwiftUI/loading-and-displaying-a-large-data-feed
// [Manual migrations]: https://developer.apple.com/documentation/coredata/manual-migrations
// [Migrating your data model automatically]: https://developer.apple.com/documentation/coredata/migrating-your-data-model-automatically
// [Mirroring a Core Data store with CloudKit]: https://developer.apple.com/documentation/coredata/mirroring-a-core-data-store-with-cloudkit
// [Modeling data]: https://developer.apple.com/documentation/coredata/modeling-data
// [Persistent history]: https://developer.apple.com/documentation/coredata/persistent-history
// [Setting up a Core Data stack]: https://developer.apple.com/documentation/coredata/setting-up-a-core-data-stack
// [Sharing Core Data objects between iCloud users]: https://developer.apple.com/documentation/coredata/sharing-core-data-objects-between-icloud-users
// [Staged migrations]: https://developer.apple.com/documentation/coredata/staged-migrations
// [Synchronizing a local store to the cloud]: https://developer.apple.com/documentation/coredata/synchronizing-a-local-store-to-the-cloud
// [Using Core Data in the background]: https://developer.apple.com/documentation/coredata/using-core-data-in-the-background
package coredata

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the CoreData library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/CoreData.framework/CoreData",
	"/usr/lib/libCoreData.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: CoreData: failed to load framework from any known path\n")
	}
}
