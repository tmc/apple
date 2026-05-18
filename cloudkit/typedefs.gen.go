// Code generated from Apple documentation. DO NOT EDIT.

package cloudkit

import (
	"github.com/tmc/apple/foundation"
)

// CKApplicationPermissionBlock is a closure that processes the outcome of a permissions request.
//
// Deprecated: Deprecated since macOS 14.0. No longer supported. Please see Sharing CloudKit Data with Other iCloud Users.
//
// See: https://developer.apple.com/documentation/CloudKit/CKContainer/ApplicationPermissionBlock
type CKApplicationPermissionBlock = func(CKApplicationPermissionStatus, foundation.NSError)

// CKOperationID is a type that represents the ID of an operation.
//
// See: https://developer.apple.com/documentation/CloudKit/CKOperationID
type CKOperationID = string

// CKRecordFieldKey is a data type that CloudKit requires for record field names.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordFieldKey
type CKRecordFieldKey = string

// CKRecordType is a data type that CloudKit requires for record types.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordType
type CKRecordType = string

// See: https://developer.apple.com/documentation/CloudKit/CKSharePreparationCompletionHandler
type CKSharePreparationCompletionHandler = func(CKShare, foundation.NSError)

// See: https://developer.apple.com/documentation/CloudKit/CKSharePreparationHandler
type CKSharePreparationHandler = func(func(*CKShare, *foundation.NSError))

// CKSubscriptionID is a type that represents a subscription’s identifier.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSubscriptionID
type CKSubscriptionID = string
