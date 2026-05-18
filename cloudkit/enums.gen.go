// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/CloudKit/CKAccountStatus
type CKAccountStatus int

const (
	// CKAccountStatusAvailable: The user’s iCloud account is available.
	CKAccountStatusAvailable CKAccountStatus = 1
	// CKAccountStatusCouldNotDetermine: CloudKit can’t determine the status of the user’s iCloud account.
	CKAccountStatusCouldNotDetermine CKAccountStatus = 0
	// CKAccountStatusNoAccount: The device doesn’t have an iCloud account.
	CKAccountStatusNoAccount CKAccountStatus = 3
	// CKAccountStatusRestricted: The system denies access to the user’s iCloud account.
	CKAccountStatusRestricted CKAccountStatus = 2
	// CKAccountStatusTemporarilyUnavailable: The user’s iCloud account is temporarily unavailable.
	CKAccountStatusTemporarilyUnavailable CKAccountStatus = 4
)

func (e CKAccountStatus) String() string {
	switch e {
	case CKAccountStatusAvailable:
		return "CKAccountStatusAvailable"
	case CKAccountStatusCouldNotDetermine:
		return "CKAccountStatusCouldNotDetermine"
	case CKAccountStatusNoAccount:
		return "CKAccountStatusNoAccount"
	case CKAccountStatusRestricted:
		return "CKAccountStatusRestricted"
	case CKAccountStatusTemporarilyUnavailable:
		return "CKAccountStatusTemporarilyUnavailable"
	default:
		return fmt.Sprintf("CKAccountStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKContainer/ApplicationPermissionStatus
type CKApplicationPermissionStatus int

const (
	// CKApplicationPermissionStatusCouldNotComplete: An error that occurs while processing the permission request.
	CKApplicationPermissionStatusCouldNotComplete CKApplicationPermissionStatus = 1
	// CKApplicationPermissionStatusDenied: The user denies the permission.
	CKApplicationPermissionStatusDenied CKApplicationPermissionStatus = 2
	// CKApplicationPermissionStatusGranted: The user grants the permission.
	CKApplicationPermissionStatusGranted CKApplicationPermissionStatus = 3
	// CKApplicationPermissionStatusInitialState: The app is yet to request the permission.
	CKApplicationPermissionStatusInitialState CKApplicationPermissionStatus = 0
)

func (e CKApplicationPermissionStatus) String() string {
	switch e {
	case CKApplicationPermissionStatusCouldNotComplete:
		return "CKApplicationPermissionStatusCouldNotComplete"
	case CKApplicationPermissionStatusDenied:
		return "CKApplicationPermissionStatusDenied"
	case CKApplicationPermissionStatusGranted:
		return "CKApplicationPermissionStatusGranted"
	case CKApplicationPermissionStatusInitialState:
		return "CKApplicationPermissionStatusInitialState"
	default:
		return fmt.Sprintf("CKApplicationPermissionStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKContainer/ApplicationPermissions
type CKApplicationPermissions uint

const (
	// Deprecated.
	CKApplicationPermissionUserDiscoverability CKApplicationPermissions = 1
)

func (e CKApplicationPermissions) String() string {
	switch e {
	case CKApplicationPermissionUserDiscoverability:
		return "CKApplicationPermissionUserDiscoverability"
	default:
		return fmt.Sprintf("CKApplicationPermissions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKDatabase/Scope
type CKDatabaseScope int

const (
	// CKDatabaseScopePrivate: The private database.
	CKDatabaseScopePrivate CKDatabaseScope = 2
	// CKDatabaseScopePublic: The public database.
	CKDatabaseScopePublic CKDatabaseScope = 1
	// CKDatabaseScopeShared: The shared database.
	CKDatabaseScopeShared CKDatabaseScope = 3
)

func (e CKDatabaseScope) String() string {
	switch e {
	case CKDatabaseScopePrivate:
		return "CKDatabaseScopePrivate"
	case CKDatabaseScopePublic:
		return "CKDatabaseScopePublic"
	case CKDatabaseScopeShared:
		return "CKDatabaseScopeShared"
	default:
		return fmt.Sprintf("CKDatabaseScope(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKError/Code
type CKErrorCode int

const (
	// CKErrorAccountTemporarilyUnavailable: An error that occurs when the user’s iCloud account is temporarily unavailable.
	CKErrorAccountTemporarilyUnavailable CKErrorCode = 36
	// CKErrorAlreadyShared: An error that occurs when CloudKit attempts to share a record with an existing share.
	CKErrorAlreadyShared CKErrorCode = 30
	// CKErrorAssetFileModified: An error that occurs when the system modifies an asset while saving it.
	CKErrorAssetFileModified CKErrorCode = 17
	// CKErrorAssetFileNotFound: An error that occurs when the system can’t find the specified asset.
	CKErrorAssetFileNotFound CKErrorCode = 16
	// CKErrorAssetNotAvailable: An error that occurs when the system can’t access the specified asset.
	CKErrorAssetNotAvailable CKErrorCode = 35
	// CKErrorBadContainer: An error that occurs when you use an unknown or unauthorized container.
	CKErrorBadContainer CKErrorCode = 5
	// CKErrorBadDatabase: An error that occurs when the operation can’t complete for the specified database.
	CKErrorBadDatabase CKErrorCode = 24
	// CKErrorBatchRequestFailed: An error that occurs when the system rejects the entire batch of changes.
	CKErrorBatchRequestFailed CKErrorCode = 22
	// CKErrorChangeTokenExpired: An error that occurs when the change token expires.
	CKErrorChangeTokenExpired CKErrorCode = 21
	// CKErrorConstraintViolation: An error that occurs when the server rejects the request because of a unique constraint violation.
	CKErrorConstraintViolation CKErrorCode = 19
	// CKErrorIncompatibleVersion: An error that occurs when the current app version is older than the oldest allowed version.
	CKErrorIncompatibleVersion CKErrorCode = 18
	// CKErrorInternalError: A nonrecoverable error that CloudKit encounters.
	CKErrorInternalError CKErrorCode = 1
	// CKErrorInvalidArguments: An error that occurs when the request contains invalid information.
	CKErrorInvalidArguments CKErrorCode = 12
	// CKErrorLimitExceeded: An error that occurs when a request’s size exceeds the limit.
	CKErrorLimitExceeded CKErrorCode = 27
	// CKErrorManagedAccountRestricted: An error that occurs when CloudKit rejects a request due to a managed-account restriction.
	CKErrorManagedAccountRestricted CKErrorCode = 32
	// CKErrorMissingEntitlement: An error that occurs when the app is missing a required entitlement.
	CKErrorMissingEntitlement CKErrorCode = 8
	// CKErrorNetworkFailure: An error that occurs when a network is available, but CloudKit is inaccessible.
	CKErrorNetworkFailure CKErrorCode = 4
	// CKErrorNetworkUnavailable: An error that occurs when the network is unavailable.
	CKErrorNetworkUnavailable CKErrorCode = 3
	// CKErrorNotAuthenticated: An error that occurs when the user is unauthenticated.
	CKErrorNotAuthenticated CKErrorCode = 9
	// CKErrorOperationCancelled: An error that occurs when an operation cancels.
	CKErrorOperationCancelled CKErrorCode = 20
	// CKErrorPartialFailure: An error that occurs when an operation completes with partial failures.
	CKErrorPartialFailure CKErrorCode = 2
	// CKErrorParticipantAlreadyInvited: The user is already an invited participant on this share.
	CKErrorParticipantAlreadyInvited CKErrorCode = 37
	// CKErrorParticipantMayNeedVerification: An error that occurs when the user isn’t a participant of the share.
	CKErrorParticipantMayNeedVerification CKErrorCode = 33
	// CKErrorPermissionFailure: An error that occurs when the user doesn’t have permission to save or fetch data.
	CKErrorPermissionFailure CKErrorCode = 10
	// CKErrorQuotaExceeded: An error that occurs when saving a record exceeds the user’s storage quota.
	CKErrorQuotaExceeded CKErrorCode = 25
	// CKErrorReferenceViolation: An error that occurs when CloudKit can’t find the target of a reference.
	CKErrorReferenceViolation CKErrorCode = 31
	// CKErrorRequestRateLimited: An error that occurs when CloudKit rate-limits requests.
	CKErrorRequestRateLimited CKErrorCode = 7
	// CKErrorResultsTruncated: An error that occurs when CloudKit truncates a query’s results.
	CKErrorResultsTruncated CKErrorCode = 13
	// CKErrorServerRecordChanged: An error that occurs when CloudKit rejects a record because the server’s version is different.
	CKErrorServerRecordChanged CKErrorCode = 14
	// CKErrorServerRejectedRequest: An error that occurs when CloudKit rejects the request.
	CKErrorServerRejectedRequest CKErrorCode = 15
	// CKErrorServerResponseLost: An error that occurs when CloudKit is unable to maintain the network connection and provide a response.
	CKErrorServerResponseLost CKErrorCode = 34
	// CKErrorServiceUnavailable: An error that occurs when CloudKit is unavailable.
	CKErrorServiceUnavailable CKErrorCode = 6
	// CKErrorTooManyParticipants: An error that occurs when a share has too many participants.
	CKErrorTooManyParticipants CKErrorCode = 29
	// CKErrorUnknownItem: An error that occurs when the specified record doesn’t exist.
	CKErrorUnknownItem CKErrorCode = 11
	// CKErrorUserDeletedZone: An error that occurs when the user deletes a record zone using the Settings app.
	CKErrorUserDeletedZone CKErrorCode = 28
	// CKErrorZoneBusy: An error that occurs when the server is too busy to handle the record zone operation.
	CKErrorZoneBusy CKErrorCode = 23
	// CKErrorZoneNotFound: An error that occurs when the specified record zone doesn’t exist.
	CKErrorZoneNotFound CKErrorCode = 26
)

func (e CKErrorCode) String() string {
	switch e {
	case CKErrorAccountTemporarilyUnavailable:
		return "CKErrorAccountTemporarilyUnavailable"
	case CKErrorAlreadyShared:
		return "CKErrorAlreadyShared"
	case CKErrorAssetFileModified:
		return "CKErrorAssetFileModified"
	case CKErrorAssetFileNotFound:
		return "CKErrorAssetFileNotFound"
	case CKErrorAssetNotAvailable:
		return "CKErrorAssetNotAvailable"
	case CKErrorBadContainer:
		return "CKErrorBadContainer"
	case CKErrorBadDatabase:
		return "CKErrorBadDatabase"
	case CKErrorBatchRequestFailed:
		return "CKErrorBatchRequestFailed"
	case CKErrorChangeTokenExpired:
		return "CKErrorChangeTokenExpired"
	case CKErrorConstraintViolation:
		return "CKErrorConstraintViolation"
	case CKErrorIncompatibleVersion:
		return "CKErrorIncompatibleVersion"
	case CKErrorInternalError:
		return "CKErrorInternalError"
	case CKErrorInvalidArguments:
		return "CKErrorInvalidArguments"
	case CKErrorLimitExceeded:
		return "CKErrorLimitExceeded"
	case CKErrorManagedAccountRestricted:
		return "CKErrorManagedAccountRestricted"
	case CKErrorMissingEntitlement:
		return "CKErrorMissingEntitlement"
	case CKErrorNetworkFailure:
		return "CKErrorNetworkFailure"
	case CKErrorNetworkUnavailable:
		return "CKErrorNetworkUnavailable"
	case CKErrorNotAuthenticated:
		return "CKErrorNotAuthenticated"
	case CKErrorOperationCancelled:
		return "CKErrorOperationCancelled"
	case CKErrorPartialFailure:
		return "CKErrorPartialFailure"
	case CKErrorParticipantAlreadyInvited:
		return "CKErrorParticipantAlreadyInvited"
	case CKErrorParticipantMayNeedVerification:
		return "CKErrorParticipantMayNeedVerification"
	case CKErrorPermissionFailure:
		return "CKErrorPermissionFailure"
	case CKErrorQuotaExceeded:
		return "CKErrorQuotaExceeded"
	case CKErrorReferenceViolation:
		return "CKErrorReferenceViolation"
	case CKErrorRequestRateLimited:
		return "CKErrorRequestRateLimited"
	case CKErrorResultsTruncated:
		return "CKErrorResultsTruncated"
	case CKErrorServerRecordChanged:
		return "CKErrorServerRecordChanged"
	case CKErrorServerRejectedRequest:
		return "CKErrorServerRejectedRequest"
	case CKErrorServerResponseLost:
		return "CKErrorServerResponseLost"
	case CKErrorServiceUnavailable:
		return "CKErrorServiceUnavailable"
	case CKErrorTooManyParticipants:
		return "CKErrorTooManyParticipants"
	case CKErrorUnknownItem:
		return "CKErrorUnknownItem"
	case CKErrorUserDeletedZone:
		return "CKErrorUserDeletedZone"
	case CKErrorZoneBusy:
		return "CKErrorZoneBusy"
	case CKErrorZoneNotFound:
		return "CKErrorZoneNotFound"
	default:
		return fmt.Sprintf("CKErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKNotification/NotificationType-swift.enum
type CKNotificationType int

const (
	// CKNotificationTypeDatabase: A notification that CloudKit generates when the contents of a database change.
	CKNotificationTypeDatabase CKNotificationType = 4
	// CKNotificationTypeQuery: A notification that CloudKit generates from a query subscription’s predicate.
	CKNotificationTypeQuery CKNotificationType = 1
	// CKNotificationTypeReadNotification: A notification that your app marks as read.
	CKNotificationTypeReadNotification CKNotificationType = 3
	// CKNotificationTypeRecordZone: A notification that CloudKit generates when the contents of a record zone change.
	CKNotificationTypeRecordZone CKNotificationType = 2
)

func (e CKNotificationType) String() string {
	switch e {
	case CKNotificationTypeDatabase:
		return "CKNotificationTypeDatabase"
	case CKNotificationTypeQuery:
		return "CKNotificationTypeQuery"
	case CKNotificationTypeReadNotification:
		return "CKNotificationTypeReadNotification"
	case CKNotificationTypeRecordZone:
		return "CKNotificationTypeRecordZone"
	default:
		return fmt.Sprintf("CKNotificationType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/TransferSize
type CKOperationGroupTransferSize int

const (
	// CKOperationGroupTransferSizeGigabytes: A transfer size that represents 1 or more gigabytes.
	CKOperationGroupTransferSizeGigabytes CKOperationGroupTransferSize = 5
	// CKOperationGroupTransferSizeHundredsOfGigabytes: A transfer size that represents hundreds of gigabytes.
	CKOperationGroupTransferSizeHundredsOfGigabytes CKOperationGroupTransferSize = 7
	// CKOperationGroupTransferSizeHundredsOfMegabytes: A transfer size that represents hundreds of megabytes.
	CKOperationGroupTransferSizeHundredsOfMegabytes CKOperationGroupTransferSize = 4
	// CKOperationGroupTransferSizeKilobytes: A transfer size that represents 1 or more kilobytes.
	CKOperationGroupTransferSizeKilobytes CKOperationGroupTransferSize = 1
	// CKOperationGroupTransferSizeMegabytes: A transfer size that represents 1 or more megabytes.
	CKOperationGroupTransferSizeMegabytes CKOperationGroupTransferSize = 2
	// CKOperationGroupTransferSizeTensOfGigabytes: A transfer size that represents tens of gigabytes.
	CKOperationGroupTransferSizeTensOfGigabytes CKOperationGroupTransferSize = 6
	// CKOperationGroupTransferSizeTensOfMegabytes: A transfer size that represents tens of megabytes.
	CKOperationGroupTransferSizeTensOfMegabytes CKOperationGroupTransferSize = 3
	// CKOperationGroupTransferSizeUnknown: An unknown transfer size.
	CKOperationGroupTransferSizeUnknown CKOperationGroupTransferSize = 0
)

func (e CKOperationGroupTransferSize) String() string {
	switch e {
	case CKOperationGroupTransferSizeGigabytes:
		return "CKOperationGroupTransferSizeGigabytes"
	case CKOperationGroupTransferSizeHundredsOfGigabytes:
		return "CKOperationGroupTransferSizeHundredsOfGigabytes"
	case CKOperationGroupTransferSizeHundredsOfMegabytes:
		return "CKOperationGroupTransferSizeHundredsOfMegabytes"
	case CKOperationGroupTransferSizeKilobytes:
		return "CKOperationGroupTransferSizeKilobytes"
	case CKOperationGroupTransferSizeMegabytes:
		return "CKOperationGroupTransferSizeMegabytes"
	case CKOperationGroupTransferSizeTensOfGigabytes:
		return "CKOperationGroupTransferSizeTensOfGigabytes"
	case CKOperationGroupTransferSizeTensOfMegabytes:
		return "CKOperationGroupTransferSizeTensOfMegabytes"
	case CKOperationGroupTransferSizeUnknown:
		return "CKOperationGroupTransferSizeUnknown"
	default:
		return fmt.Sprintf("CKOperationGroupTransferSize(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKQueryNotification/Reason
type CKQueryNotificationReason int

const (
	// CKQueryNotificationReasonRecordCreated: A notification that indicates the creation of a record matching the subscription’s predicate.
	CKQueryNotificationReasonRecordCreated CKQueryNotificationReason = 1
	// CKQueryNotificationReasonRecordDeleted: A notification that indicates the deletion of a record matching the subscription’s predicate.
	CKQueryNotificationReasonRecordDeleted CKQueryNotificationReason = 3
	// CKQueryNotificationReasonRecordUpdated: A notification that indicates the update of a record matching the subscription’s predicate.
	CKQueryNotificationReasonRecordUpdated CKQueryNotificationReason = 2
)

func (e CKQueryNotificationReason) String() string {
	switch e {
	case CKQueryNotificationReasonRecordCreated:
		return "CKQueryNotificationReasonRecordCreated"
	case CKQueryNotificationReasonRecordDeleted:
		return "CKQueryNotificationReasonRecordDeleted"
	case CKQueryNotificationReasonRecordUpdated:
		return "CKQueryNotificationReasonRecordUpdated"
	default:
		return fmt.Sprintf("CKQueryNotificationReason(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKQuerySubscription/Options
type CKQuerySubscriptionOptions uint

const (
	// CKQuerySubscriptionOptionsFiresOnRecordCreation: An option that instructs CloudKit to send a push notification when it creates a record that matches a subscription’s criteria.
	CKQuerySubscriptionOptionsFiresOnRecordCreation CKQuerySubscriptionOptions = 1
	// CKQuerySubscriptionOptionsFiresOnRecordDeletion: An option that instructs CloudKit to send a push notification when it deletes a record that matches a subscription’s criteria.
	CKQuerySubscriptionOptionsFiresOnRecordDeletion CKQuerySubscriptionOptions = 4
	// CKQuerySubscriptionOptionsFiresOnRecordUpdate: An option that instructs CloudKit to send a push notification when it modifies a record that matches a subscription’s criteria.
	CKQuerySubscriptionOptionsFiresOnRecordUpdate CKQuerySubscriptionOptions = 2
	// CKQuerySubscriptionOptionsFiresOnce: An option that instructs CloudKit to send a push notification only once.
	CKQuerySubscriptionOptionsFiresOnce CKQuerySubscriptionOptions = 8
)

func (e CKQuerySubscriptionOptions) String() string {
	switch e {
	case CKQuerySubscriptionOptionsFiresOnRecordCreation:
		return "CKQuerySubscriptionOptionsFiresOnRecordCreation"
	case CKQuerySubscriptionOptionsFiresOnRecordDeletion:
		return "CKQuerySubscriptionOptionsFiresOnRecordDeletion"
	case CKQuerySubscriptionOptionsFiresOnRecordUpdate:
		return "CKQuerySubscriptionOptionsFiresOnRecordUpdate"
	case CKQuerySubscriptionOptionsFiresOnce:
		return "CKQuerySubscriptionOptionsFiresOnce"
	default:
		return fmt.Sprintf("CKQuerySubscriptionOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/RecordSavePolicy
type CKRecordSavePolicy int

const (
	// CKRecordSaveAllKeys: A policy that instructs CloudKit to save all keys of a record, even those without changes.
	CKRecordSaveAllKeys CKRecordSavePolicy = 2
	// CKRecordSaveChangedKeys: A policy that instructs CloudKit to save only the fields of a record that contain changes.
	CKRecordSaveChangedKeys CKRecordSavePolicy = 1
	// CKRecordSaveIfServerRecordUnchanged: A policy that instructs CloudKit to only proceed if the record’s change tag matches that of the server’s copy.
	CKRecordSaveIfServerRecordUnchanged CKRecordSavePolicy = 0
)

func (e CKRecordSavePolicy) String() string {
	switch e {
	case CKRecordSaveAllKeys:
		return "CKRecordSaveAllKeys"
	case CKRecordSaveChangedKeys:
		return "CKRecordSaveChangedKeys"
	case CKRecordSaveIfServerRecordUnchanged:
		return "CKRecordSaveIfServerRecordUnchanged"
	default:
		return fmt.Sprintf("CKRecordSavePolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKRecordZone/Capabilities-swift.struct
type CKRecordZoneCapabilities uint

const (
	// CKRecordZoneCapabilityAtomic: A capability that allows atomic changes of multiple records.
	CKRecordZoneCapabilityAtomic CKRecordZoneCapabilities = 2
	// CKRecordZoneCapabilityFetchChanges: A capability for fetching only the changed records from a zone.
	CKRecordZoneCapabilityFetchChanges CKRecordZoneCapabilities = 1
	// CKRecordZoneCapabilitySharing: A capability for sharing a specific hierarchy of records.
	CKRecordZoneCapabilitySharing CKRecordZoneCapabilities = 4
	// CKRecordZoneCapabilityZoneWideSharing: A capability for sharing the entire contents of a record zone.
	CKRecordZoneCapabilityZoneWideSharing CKRecordZoneCapabilities = 8
)

func (e CKRecordZoneCapabilities) String() string {
	switch e {
	case CKRecordZoneCapabilityAtomic:
		return "CKRecordZoneCapabilityAtomic"
	case CKRecordZoneCapabilityFetchChanges:
		return "CKRecordZoneCapabilityFetchChanges"
	case CKRecordZoneCapabilitySharing:
		return "CKRecordZoneCapabilitySharing"
	case CKRecordZoneCapabilityZoneWideSharing:
		return "CKRecordZoneCapabilityZoneWideSharing"
	default:
		return fmt.Sprintf("CKRecordZoneCapabilities(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKRecordZone/EncryptionScope-swift.enum
type CKRecordZoneEncryptionScope int

const (
	// CKRecordZoneEncryptionScopePerRecord: Zone uses per-record encryption keys for any encrypted values on a record or share.
	CKRecordZoneEncryptionScopePerRecord CKRecordZoneEncryptionScope = 0
	// CKRecordZoneEncryptionScopePerZone: Zone uses per-zone encryption keys for encrypted values across all records and the zone-wide share, if present.
	CKRecordZoneEncryptionScopePerZone CKRecordZoneEncryptionScope = 1
)

func (e CKRecordZoneEncryptionScope) String() string {
	switch e {
	case CKRecordZoneEncryptionScopePerRecord:
		return "CKRecordZoneEncryptionScopePerRecord"
	case CKRecordZoneEncryptionScopePerZone:
		return "CKRecordZoneEncryptionScopePerZone"
	default:
		return fmt.Sprintf("CKRecordZoneEncryptionScope(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKRecord/ReferenceAction
type CKReferenceAction uint

const (
	// CKReferenceActionDeleteSelf: A reference action that cascades deletions.
	CKReferenceActionDeleteSelf CKReferenceAction = 1
	// CKReferenceActionNone: A reference action that has no cascading behavior.
	CKReferenceActionNone CKReferenceAction = 0
)

func (e CKReferenceAction) String() string {
	switch e {
	case CKReferenceActionDeleteSelf:
		return "CKReferenceActionDeleteSelf"
	case CKReferenceActionNone:
		return "CKReferenceActionNone"
	default:
		return fmt.Sprintf("CKReferenceAction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantAcceptanceStatus
type CKShareParticipantAcceptanceStatus int

const (
	// CKShareParticipantAcceptanceStatusAccepted: The participant accepted the share request.
	CKShareParticipantAcceptanceStatusAccepted CKShareParticipantAcceptanceStatus = 2
	// CKShareParticipantAcceptanceStatusPending: The participant’s acceptance of the share request is pending.
	CKShareParticipantAcceptanceStatusPending CKShareParticipantAcceptanceStatus = 1
	// CKShareParticipantAcceptanceStatusRemoved: The system removed the participant from the share.
	CKShareParticipantAcceptanceStatusRemoved CKShareParticipantAcceptanceStatus = 3
	// CKShareParticipantAcceptanceStatusUnknown: The participant’s status is unknown.
	CKShareParticipantAcceptanceStatusUnknown CKShareParticipantAcceptanceStatus = 0
)

func (e CKShareParticipantAcceptanceStatus) String() string {
	switch e {
	case CKShareParticipantAcceptanceStatusAccepted:
		return "CKShareParticipantAcceptanceStatusAccepted"
	case CKShareParticipantAcceptanceStatusPending:
		return "CKShareParticipantAcceptanceStatusPending"
	case CKShareParticipantAcceptanceStatusRemoved:
		return "CKShareParticipantAcceptanceStatusRemoved"
	case CKShareParticipantAcceptanceStatusUnknown:
		return "CKShareParticipantAcceptanceStatusUnknown"
	default:
		return fmt.Sprintf("CKShareParticipantAcceptanceStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantPermission
type CKShareParticipantPermission int

const (
	// CKShareParticipantPermissionNone: The participant doesn’t have any permissions for the share.
	CKShareParticipantPermissionNone CKShareParticipantPermission = 1
	// CKShareParticipantPermissionReadOnly: The participant has read-only permissions for the share.
	CKShareParticipantPermissionReadOnly CKShareParticipantPermission = 2
	// CKShareParticipantPermissionReadWrite: The participant has read-and-write permissions for the share.
	CKShareParticipantPermissionReadWrite CKShareParticipantPermission = 3
	// CKShareParticipantPermissionUnknown: The participant’s permissions are unknown.
	CKShareParticipantPermissionUnknown CKShareParticipantPermission = 0
)

func (e CKShareParticipantPermission) String() string {
	switch e {
	case CKShareParticipantPermissionNone:
		return "CKShareParticipantPermissionNone"
	case CKShareParticipantPermissionReadOnly:
		return "CKShareParticipantPermissionReadOnly"
	case CKShareParticipantPermissionReadWrite:
		return "CKShareParticipantPermissionReadWrite"
	case CKShareParticipantPermissionUnknown:
		return "CKShareParticipantPermissionUnknown"
	default:
		return fmt.Sprintf("CKShareParticipantPermission(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantRole
type CKShareParticipantRole int

const (
	// CKShareParticipantRoleAdministrator: The participant has the administrator role.
	CKShareParticipantRoleAdministrator CKShareParticipantRole = 2
	// CKShareParticipantRoleOwner: The participant is the share’s owner.
	CKShareParticipantRoleOwner CKShareParticipantRole = 1
	// CKShareParticipantRolePrivateUser: The participant has the private role.
	CKShareParticipantRolePrivateUser CKShareParticipantRole = 3
	// CKShareParticipantRolePublicUser: The participant has the public role.
	CKShareParticipantRolePublicUser CKShareParticipantRole = 4
	// CKShareParticipantRoleUnknown: The participant’s role is unknown.
	CKShareParticipantRoleUnknown CKShareParticipantRole = 0
)

func (e CKShareParticipantRole) String() string {
	switch e {
	case CKShareParticipantRoleAdministrator:
		return "CKShareParticipantRoleAdministrator"
	case CKShareParticipantRoleOwner:
		return "CKShareParticipantRoleOwner"
	case CKShareParticipantRolePrivateUser:
		return "CKShareParticipantRolePrivateUser"
	case CKShareParticipantRolePublicUser:
		return "CKShareParticipantRolePublicUser"
	case CKShareParticipantRoleUnknown:
		return "CKShareParticipantRoleUnknown"
	default:
		return fmt.Sprintf("CKShareParticipantRole(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKSharingParticipantAccessOption
type CKSharingParticipantAccessOption uint

const (
	// CKSharingParticipantAccessOptionAny: The permission option the system uses to control whether a user can share publicly or privately.
	CKSharingParticipantAccessOptionAny CKSharingParticipantAccessOption = 1
	// CKSharingParticipantAccessOptionAnyoneWithLink: The permission option the system uses to control whether a user can share publicly.
	CKSharingParticipantAccessOptionAnyoneWithLink CKSharingParticipantAccessOption = 1
	// CKSharingParticipantAccessOptionSpecifiedRecipientsOnly: The permission option the system uses to control whether a user can share privately.
	CKSharingParticipantAccessOptionSpecifiedRecipientsOnly CKSharingParticipantAccessOption = 2
)

func (e CKSharingParticipantAccessOption) String() string {
	switch e {
	case CKSharingParticipantAccessOptionAny:
		return "CKSharingParticipantAccessOptionAny"
	case CKSharingParticipantAccessOptionSpecifiedRecipientsOnly:
		return "CKSharingParticipantAccessOptionSpecifiedRecipientsOnly"
	default:
		return fmt.Sprintf("CKSharingParticipantAccessOption(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKSharingParticipantPermissionOption
type CKSharingParticipantPermissionOption uint

const (
	// CKSharingParticipantPermissionOptionAny: The permission option the system uses to control whether a user can grant read-only or write access.
	CKSharingParticipantPermissionOptionAny CKSharingParticipantPermissionOption = 1
	// CKSharingParticipantPermissionOptionReadOnly: The permission option the system uses to control whether a user can grant read-only access.
	CKSharingParticipantPermissionOptionReadOnly CKSharingParticipantPermissionOption = 1
	// CKSharingParticipantPermissionOptionReadWrite: The permission option the system uses to control whether a user can grant write access.
	CKSharingParticipantPermissionOptionReadWrite CKSharingParticipantPermissionOption = 2
)

func (e CKSharingParticipantPermissionOption) String() string {
	switch e {
	case CKSharingParticipantPermissionOptionAny:
		return "CKSharingParticipantPermissionOptionAny"
	case CKSharingParticipantPermissionOptionReadWrite:
		return "CKSharingParticipantPermissionOptionReadWrite"
	default:
		return fmt.Sprintf("CKSharingParticipantPermissionOption(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKSubscription/SubscriptionType-swift.enum
type CKSubscriptionType int

const (
	// CKSubscriptionTypeDatabase: A constant that indicates the subscription is database-based.
	CKSubscriptionTypeDatabase CKSubscriptionType = 3
	// CKSubscriptionTypeQuery: A constant that indicates the subscription is query-based.
	CKSubscriptionTypeQuery CKSubscriptionType = 1
	// CKSubscriptionTypeRecordZone: A constant that indicates the subscription is zone-based.
	CKSubscriptionTypeRecordZone CKSubscriptionType = 2
)

func (e CKSubscriptionType) String() string {
	switch e {
	case CKSubscriptionTypeDatabase:
		return "CKSubscriptionTypeDatabase"
	case CKSubscriptionTypeQuery:
		return "CKSubscriptionTypeQuery"
	case CKSubscriptionTypeRecordZone:
		return "CKSubscriptionTypeRecordZone"
	default:
		return fmt.Sprintf("CKSubscriptionType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineAccountChangeType
type CKSyncEngineAccountChangeType int

const (
	// CKSyncEngineAccountChangeTypeSignIn: A change indicating a sign-in to an iCloud account.
	CKSyncEngineAccountChangeTypeSignIn CKSyncEngineAccountChangeType = 0
	// CKSyncEngineAccountChangeTypeSignOut: A change indicating a sign-out of an iCloud account.
	CKSyncEngineAccountChangeTypeSignOut CKSyncEngineAccountChangeType = 1
	// CKSyncEngineAccountChangeTypeSwitchAccounts: A change indicating a switch between two iCloud accounts.
	CKSyncEngineAccountChangeTypeSwitchAccounts CKSyncEngineAccountChangeType = 2
)

func (e CKSyncEngineAccountChangeType) String() string {
	switch e {
	case CKSyncEngineAccountChangeTypeSignIn:
		return "CKSyncEngineAccountChangeTypeSignIn"
	case CKSyncEngineAccountChangeTypeSignOut:
		return "CKSyncEngineAccountChangeTypeSignOut"
	case CKSyncEngineAccountChangeTypeSwitchAccounts:
		return "CKSyncEngineAccountChangeTypeSwitchAccounts"
	default:
		return fmt.Sprintf("CKSyncEngineAccountChangeType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEventType
type CKSyncEngineEventType int

const (
	// CKSyncEngineEventTypeAccountChange: The user signed in or out of their account.
	CKSyncEngineEventTypeAccountChange CKSyncEngineEventType = 1
	// CKSyncEngineEventTypeDidFetchChanges: The sync engine finished fetching changes from the server.
	CKSyncEngineEventTypeDidFetchChanges CKSyncEngineEventType = 9
	// CKSyncEngineEventTypeDidFetchRecordZoneChanges: The sync engine has completed fetching record zone changes from the server for a specific zone.
	CKSyncEngineEventTypeDidFetchRecordZoneChanges CKSyncEngineEventType = 8
	// CKSyncEngineEventTypeDidSendChanges: The sync engine finished sending changes to the server.
	CKSyncEngineEventTypeDidSendChanges CKSyncEngineEventType = 11
	// CKSyncEngineEventTypeFetchedDatabaseChanges: The sync engine has fetched new database changes from the server.
	CKSyncEngineEventTypeFetchedDatabaseChanges CKSyncEngineEventType = 2
	// CKSyncEngineEventTypeFetchedRecordZoneChanges: The sync engine fetched new record zone changes from the server.
	CKSyncEngineEventTypeFetchedRecordZoneChanges CKSyncEngineEventType = 3
	// CKSyncEngineEventTypeSentDatabaseChanges: The sync engine sent a batch of database changes to the server.
	CKSyncEngineEventTypeSentDatabaseChanges CKSyncEngineEventType = 4
	// CKSyncEngineEventTypeSentRecordZoneChanges: The sync engine sent a batch of record zone changes to the server.
	CKSyncEngineEventTypeSentRecordZoneChanges CKSyncEngineEventType = 5
	// CKSyncEngineEventTypeStateUpdate: The sync engine updated its state.
	CKSyncEngineEventTypeStateUpdate CKSyncEngineEventType = 0
	// CKSyncEngineEventTypeWillFetchChanges: The sync engine is about to fetch changes from the server.
	CKSyncEngineEventTypeWillFetchChanges CKSyncEngineEventType = 6
	// CKSyncEngineEventTypeWillFetchRecordZoneChanges: The sync engine is about to fetch record zone changes from the server for a specific zone.
	CKSyncEngineEventTypeWillFetchRecordZoneChanges CKSyncEngineEventType = 7
	// CKSyncEngineEventTypeWillSendChanges: The sync engine is about to send changes to the server.
	CKSyncEngineEventTypeWillSendChanges CKSyncEngineEventType = 10
)

func (e CKSyncEngineEventType) String() string {
	switch e {
	case CKSyncEngineEventTypeAccountChange:
		return "CKSyncEngineEventTypeAccountChange"
	case CKSyncEngineEventTypeDidFetchChanges:
		return "CKSyncEngineEventTypeDidFetchChanges"
	case CKSyncEngineEventTypeDidFetchRecordZoneChanges:
		return "CKSyncEngineEventTypeDidFetchRecordZoneChanges"
	case CKSyncEngineEventTypeDidSendChanges:
		return "CKSyncEngineEventTypeDidSendChanges"
	case CKSyncEngineEventTypeFetchedDatabaseChanges:
		return "CKSyncEngineEventTypeFetchedDatabaseChanges"
	case CKSyncEngineEventTypeFetchedRecordZoneChanges:
		return "CKSyncEngineEventTypeFetchedRecordZoneChanges"
	case CKSyncEngineEventTypeSentDatabaseChanges:
		return "CKSyncEngineEventTypeSentDatabaseChanges"
	case CKSyncEngineEventTypeSentRecordZoneChanges:
		return "CKSyncEngineEventTypeSentRecordZoneChanges"
	case CKSyncEngineEventTypeStateUpdate:
		return "CKSyncEngineEventTypeStateUpdate"
	case CKSyncEngineEventTypeWillFetchChanges:
		return "CKSyncEngineEventTypeWillFetchChanges"
	case CKSyncEngineEventTypeWillFetchRecordZoneChanges:
		return "CKSyncEngineEventTypeWillFetchRecordZoneChanges"
	case CKSyncEngineEventTypeWillSendChanges:
		return "CKSyncEngineEventTypeWillSendChanges"
	default:
		return fmt.Sprintf("CKSyncEngineEventType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingDatabaseChangeType
type CKSyncEnginePendingDatabaseChangeType int

const (
	CKSyncEnginePendingDatabaseChangeTypeDeleteZone CKSyncEnginePendingDatabaseChangeType = 1
	CKSyncEnginePendingDatabaseChangeTypeSaveZone   CKSyncEnginePendingDatabaseChangeType = 0
)

func (e CKSyncEnginePendingDatabaseChangeType) String() string {
	switch e {
	case CKSyncEnginePendingDatabaseChangeTypeDeleteZone:
		return "CKSyncEnginePendingDatabaseChangeTypeDeleteZone"
	case CKSyncEnginePendingDatabaseChangeTypeSaveZone:
		return "CKSyncEnginePendingDatabaseChangeTypeSaveZone"
	default:
		return fmt.Sprintf("CKSyncEnginePendingDatabaseChangeType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingRecordZoneChangeType
type CKSyncEnginePendingRecordZoneChangeType int

const (
	CKSyncEnginePendingRecordZoneChangeTypeDeleteRecord CKSyncEnginePendingRecordZoneChangeType = 1
	CKSyncEnginePendingRecordZoneChangeTypeSaveRecord   CKSyncEnginePendingRecordZoneChangeType = 0
)

func (e CKSyncEnginePendingRecordZoneChangeType) String() string {
	switch e {
	case CKSyncEnginePendingRecordZoneChangeTypeDeleteRecord:
		return "CKSyncEnginePendingRecordZoneChangeTypeDeleteRecord"
	case CKSyncEnginePendingRecordZoneChangeTypeSaveRecord:
		return "CKSyncEnginePendingRecordZoneChangeTypeSaveRecord"
	default:
		return fmt.Sprintf("CKSyncEnginePendingRecordZoneChangeType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSyncReason
type CKSyncEngineSyncReason int

const (
	// CKSyncEngineSyncReasonManual: A manual sync operation.
	CKSyncEngineSyncReasonManual CKSyncEngineSyncReason = 1
	// CKSyncEngineSyncReasonScheduled: The sync engine automatically scheduled this sync.
	CKSyncEngineSyncReasonScheduled CKSyncEngineSyncReason = 0
)

func (e CKSyncEngineSyncReason) String() string {
	switch e {
	case CKSyncEngineSyncReasonManual:
		return "CKSyncEngineSyncReasonManual"
	case CKSyncEngineSyncReasonScheduled:
		return "CKSyncEngineSyncReasonScheduled"
	default:
		return fmt.Sprintf("CKSyncEngineSyncReason(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineZoneDeletionReason
type CKSyncEngineZoneDeletionReason int

const (
	// CKSyncEngineZoneDeletionReasonDeleted: Your app deleted the record zone.
	CKSyncEngineZoneDeletionReasonDeleted CKSyncEngineZoneDeletionReason = 0
	// CKSyncEngineZoneDeletionReasonEncryptedDataReset: The owner of the iCloud account reset their encrypted data.
	CKSyncEngineZoneDeletionReasonEncryptedDataReset CKSyncEngineZoneDeletionReason = 2
	// CKSyncEngineZoneDeletionReasonPurged: The owner of the iCloud account purged your app’s data using the Settings app.
	CKSyncEngineZoneDeletionReasonPurged CKSyncEngineZoneDeletionReason = 1
)

func (e CKSyncEngineZoneDeletionReason) String() string {
	switch e {
	case CKSyncEngineZoneDeletionReasonDeleted:
		return "CKSyncEngineZoneDeletionReasonDeleted"
	case CKSyncEngineZoneDeletionReasonEncryptedDataReset:
		return "CKSyncEngineZoneDeletionReasonEncryptedDataReset"
	case CKSyncEngineZoneDeletionReasonPurged:
		return "CKSyncEngineZoneDeletionReasonPurged"
	default:
		return fmt.Sprintf("CKSyncEngineZoneDeletionReason(%d)", e)
	}
}
