// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

// Package cloudkit provides Go bindings for the CloudKit framework.
//
// Store structured app and user data in iCloud containers that all users of your app can share.
//
// The CloudKit framework provides interfaces for moving data between your app and your iCloud containers. You use CloudKit to store your app’s existing data in the cloud so that the user can access it on multiple devices. You can also store data in a public area where all users can access it.
//
// # Essentials
//
//   - Deciding whether CloudKit is right for your app: Explore the various options you have for using iCloud to store and sync your app’s data.
//   - Enabling CloudKit in Your App: Configure your app to store data in iCloud using CloudKit.
//
// # Schemas
//
//   - Designing and Creating a CloudKit Database: Create a schema to store your app’s objects as records in iCloud using CloudKit.
//   - Managing iCloud Containers with CloudKit Database App: Inspect and modify the schema and data for your app’s iCloud container.
//   - CKRecordZone: A database partition that contains related records.
//   - CKRecord: A collection of key-value pairs that store your app’s data. ([CKRecordKeyValueSetting], [CKRecordValue])
//   - CKRecord.Reference: A relationship between two records in a record zone.
//   - CKAsset: An external file that belongs to a record.
//   - Integrating a Text-Based Schema into Your Workflow: Define and update your schema with the CloudKit Schema Language.
//
// # Records
//
//   - Local Records: Manipulate records on-device and save changes to the server. ([CKModifyRecordZonesOperation], [CKModifyRecordsOperation], [CKFetchRecordZonesOperation], [CKFetchRecordsOperation], [CKQuery])
//   - Remote Records: Use subscriptions and change tokens to efficiently manage modifications to remote records. ([CKDatabaseSubscription], [CKDatabaseNotification], [CKFetchDatabaseChangesOperation], [CKRecordZoneSubscription], [CKRecordZoneNotification])
//   - CKSyncEngine: An object that manages the synchronization of local and remote record data. ([CKSyncEngineDelegate])
//   - Shared Records: Share one or more records with other iCloud users. ([CKShare], [CKAllowedSharingOptions], [CKSystemSharingUIObserver], [CKFetchShareMetadataOperation], [CKAcceptSharesOperation])
//
// # User discovery
//
//   - CKUserIdentity: The identity of a user.
//   - CKUserIdentity.LookupInfo: The criteria to use when searching for discoverable iCloud users.
//
// # Core objects
//
//   - CKContainer: A conduit to your app’s databases. ([CKAccountStatus])
//   - CKDatabase: An object that represents a collection of record zones and subscriptions.
//   - CKOperationGroup: An explicit association between two or more operations.
//
// # Privacy
//
//   - Encrypting User Data: Deploy industry-standard security technologies using CloudKit encryption.
//   - Providing User Access to CloudKit Data: Provide users access to the data your app stores on their behalf.
//   - Changing Access Controls on User Data: Restrict access to or remove restrictions from a user’s data at their request.
//   - CKFetchWebAuthTokenOperation: An operation that creates an authentication token for use with CloudKit web services.
//   - Responding to Requests to Delete Data: Provide options for users to delete their CloudKit data from your app.
//   - Identifying an App’s Containers: Use Xcode’s Project navigator to find the identifiers of active CloudKit containers.
//
// # Errors
//
//   - CKErrorDomain: The error domain for CloudKit errors.
//   - CKError: A type that describes a CloudKit error.
//   - CKError.Code: The error codes that CloudKit returns.
//   - CKErrorRetryAfterKey: The key to retrieve the number of seconds to wait before you retry a request.
//   - CKErrorUserDidResetEncryptedDataKey: The key that determines whether CloudKit deletes a record zone because of a user action.
//   - CKPartialErrorsByItemIDKey: The key to retrieve partial errors.
//   - Record Changed Error Keys: Constants that represent conflicting records in a save operation.
//
// # Classes
//
//   - CKShareRequestAccessOperation
//
// # Variables
//
//   - CKRecordParentKey
//   - CKRecordShareKey
//   - CKRecordTypeShare
//   - CKRecordTypeUserRecord
//   - CKRecordZoneDefaultName
//   - CKShareThumbnailImageDataKey
//   - CKShareTitleKey
//   - CKShareTypeKey
//
// # Key Types
//
//   - [CKShare] - A specialized record type that manages a collection of shared records.
//   - [CKContainer] - A conduit to your app’s databases.
//   - [CKNotification] - The abstract base class for CloudKit notifications.
//   - [CKNotificationInfo] - An object that describes the configuration of a subscription’s push notifications.
//   - [CKRecord] - A collection of key-value pairs that store your app’s data.
//   - [CKModifyRecordsOperation] - An operation that modifies one or more records.
//   - [CKShareMetadata] - An object that describes a shared record’s metadata.
//   - [CKSyncEngineEvent] - An event that occurs during a sync operation.
//   - [CKDatabase] - An object that represents a collection of record zones and subscriptions.
//   - [CKFetchDatabaseChangesOperation] - An operation that fetches database changes.
//
// [CloudKit Documentation]: https://developer.apple.com/documentation/CloudKit
package cloudkit

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the CloudKit library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/CloudKit.framework/CloudKit",
	"/usr/lib/libCloudKit.dylib",
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
	fmt.Fprintf(os.Stderr, "warning: CloudKit: failed to load framework from any known path\n")
}
