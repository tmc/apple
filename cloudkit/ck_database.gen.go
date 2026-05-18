// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKDatabase] class.
var (
	_CKDatabaseClass     CKDatabaseClass
	_CKDatabaseClassOnce sync.Once
)

func getCKDatabaseClass() CKDatabaseClass {
	_CKDatabaseClassOnce.Do(func() {
		_CKDatabaseClass = CKDatabaseClass{class: objc.GetClass("CKDatabase")}
	})
	return _CKDatabaseClass
}

// GetCKDatabaseClass returns the class object for CKDatabase.
func GetCKDatabaseClass() CKDatabaseClass {
	return getCKDatabaseClass()
}

type CKDatabaseClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKDatabaseClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKDatabaseClass) Alloc() CKDatabase {
	rv := objc.Send[CKDatabase](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a collection of record zones and subscriptions.
//
// # Overview
//
// A database takes requests and operations and applies them to the objects it
// contains, whether that’s record zones, records, or subscriptions. Each of
// your app’s users has access to the three separate databases:
//
// - A public database that’s accessible to all users of your app. - A
// private database that’s accessible only to the user of the current
// device. - A shared database that’s accessible only to the user of the
// current device, which contains records that other iCloud users share with
// them.
//
// The public database is always available, even when the device doesn’t
// have an active iCloud account. In this scenario, your app can fetch
// specific records and perform searches, but it can’t create or modify
// records. CloudKit requires an iCloud account for writing to the public
// database so it can identify the authors of any changes. All access to the
// private and shared databases requires an iCloud account.
//
// You don’t create instances of [CKDatabase], nor do you subclass it.
// Instead, you access the required database using one of your app’s
// containers. For more information, see [CKContainer].
//
// By default, CloudKit executes the methods in this class with a low-priority
// quality of service (QoS). To use a higher-priority QoS, perform the
// following:
//
// - Create an instance of [CKOperationConfiguration] and set its
// [CKDatabase.QualityOfService] property to the preferred value. - Call the databaseʼs
// [configuredWith(configuration:group:body:)] method and provide the
// configuration and a trailing closure. - In the closure, use the provided
// database to execute the relevant methods at the preferred QoS.
//
// # Fetching Records
//
//   - [CKDatabase.FetchRecordWithIDCompletionHandler]: Fetches a specific record.
//
// # Modifying Records
//
//   - [CKDatabase.SaveRecordCompletionHandler]: Saves a specific record.
//   - [CKDatabase.DeleteRecordWithIDCompletionHandler]: Deletes a specific record.
//
// # Fetching Record Zones
//
//   - [CKDatabase.FetchRecordZoneWithIDCompletionHandler]: Fetches a specific record zone.
//
// # Modifying Record Zones
//
//   - [CKDatabase.SaveRecordZoneCompletionHandler]: Saves a specific record zone.
//   - [CKDatabase.DeleteRecordZoneWithIDCompletionHandler]: Deletes a specific record zone.
//
// # Modifying Subscriptions
//
//   - [CKDatabase.SaveSubscriptionCompletionHandler]: Saves a specific subscription.
//
// # Running Operations
//
//   - [CKDatabase.AddOperation]: Executes the specified operation in the current database.
//
// # Getting the Database Type
//
//   - [CKDatabase.DatabaseScope]: The type of database.
//
// See: https://developer.apple.com/documentation/CloudKit/CKDatabase
//
// [configuredWith(configuration:group:body:)]: https://developer.apple.com/documentation/CloudKit/CKDatabase/configuredWith(configuration:group:body:)-637p1
type CKDatabase struct {
	objectivec.Object
}

// CKDatabaseFromID constructs a [CKDatabase] from an objc.ID.
//
// An object that represents a collection of record zones and subscriptions.
func CKDatabaseFromID(id objc.ID) CKDatabase {
	return CKDatabase{objectivec.Object{ID: id}}
}

// NOTE: CKDatabase adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKDatabase] class.
//
// # Fetching Records
//
//   - [ICKDatabase.FetchRecordWithIDCompletionHandler]: Fetches a specific record.
//
// # Modifying Records
//
//   - [ICKDatabase.SaveRecordCompletionHandler]: Saves a specific record.
//   - [ICKDatabase.DeleteRecordWithIDCompletionHandler]: Deletes a specific record.
//
// # Fetching Record Zones
//
//   - [ICKDatabase.FetchRecordZoneWithIDCompletionHandler]: Fetches a specific record zone.
//
// # Modifying Record Zones
//
//   - [ICKDatabase.SaveRecordZoneCompletionHandler]: Saves a specific record zone.
//   - [ICKDatabase.DeleteRecordZoneWithIDCompletionHandler]: Deletes a specific record zone.
//
// # Modifying Subscriptions
//
//   - [ICKDatabase.SaveSubscriptionCompletionHandler]: Saves a specific subscription.
//
// # Running Operations
//
//   - [ICKDatabase.AddOperation]: Executes the specified operation in the current database.
//
// # Getting the Database Type
//
//   - [ICKDatabase.DatabaseScope]: The type of database.
//
// See: https://developer.apple.com/documentation/CloudKit/CKDatabase
type ICKDatabase interface {
	objectivec.IObject

	// Topic: Fetching Records

	// Fetches a specific record.
	FetchRecordWithIDCompletionHandler(recordID ICKRecordID, completionHandler CKRecordErrorHandler)

	// Topic: Modifying Records

	// Saves a specific record.
	SaveRecordCompletionHandler(record ICKRecord, completionHandler CKRecordErrorHandler)
	// Deletes a specific record.
	DeleteRecordWithIDCompletionHandler(recordID ICKRecordID, completionHandler CKRecordIDErrorHandler)

	// Topic: Fetching Record Zones

	// Fetches a specific record zone.
	FetchRecordZoneWithIDCompletionHandler(zoneID ICKRecordZoneID, completionHandler CKRecordZoneErrorHandler)

	// Topic: Modifying Record Zones

	// Saves a specific record zone.
	SaveRecordZoneCompletionHandler(zone ICKRecordZone, completionHandler CKRecordZoneErrorHandler)
	// Deletes a specific record zone.
	DeleteRecordZoneWithIDCompletionHandler(zoneID ICKRecordZoneID, completionHandler CKRecordZoneIDErrorHandler)

	// Topic: Modifying Subscriptions

	// Saves a specific subscription.
	SaveSubscriptionCompletionHandler(subscription ICKSubscription, completionHandler CKSubscriptionErrorHandler)

	// Topic: Running Operations

	// Executes the specified operation in the current database.
	AddOperation(operation ICKDatabaseOperation)

	// Topic: Getting the Database Type

	// The type of database.
	DatabaseScope() CKDatabaseScope

	// The priority that the system uses when it allocates resources to the operations that use this configuration.
	QualityOfService() foundation.NSQualityOfService
	SetQualityOfService(value foundation.NSQualityOfService)
}

// Init initializes the instance.
func (c CKDatabase) Init() CKDatabase {
	rv := objc.Send[CKDatabase](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKDatabase) Autorelease() CKDatabase {
	rv := objc.Send[CKDatabase](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKDatabase creates a new CKDatabase instance.
func NewCKDatabase() CKDatabase {
	class := getCKDatabaseClass()
	rv := objc.Send[CKDatabase](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Fetches a specific record.
//
// recordID: The identifier of the record to fetch.
//
// completionHandler: The closure to execute with the fetch results.
//
// # Discussion
//
// The completion handler takes the following parameters:
//
// - The requested record, or `nil` if CloudKit can’t provide that record. -
// An error if a problem occurs, or `nil` if the fetch completes successfully.
//
// For information on a more convenient way to fetch specific records, see
// [records(for:desiredKeys:)].
//
// See: https://developer.apple.com/documentation/CloudKit/CKDatabase/fetch(withRecordID:completionHandler:)
//
// [records(for:desiredKeys:)]: https://developer.apple.com/documentation/CloudKit/CKDatabase/records(for:desiredKeys:)
func (c CKDatabase) FetchRecordWithIDCompletionHandler(recordID ICKRecordID, completionHandler CKRecordErrorHandler) {
	_block1, _ := NewCKRecordErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("fetchRecordWithID:completionHandler:"), recordID, _block1)
}

// Saves a specific record.
//
// record: The record to save.
//
// completionHandler: The closure to execute after CloudKit saves the record.
//
// # Discussion
//
// The completion handler takes the following parameters:
//
// - The saved record (as it appears on the server), or `nil` if there’s an
// error. - An error if a problem occurs, or `nil` if CloudKit successfully
// saves the record.
//
// The save succeeds only when the specified record is new, or is a more
// recent version than the one on the server.
//
// For information on a more convenient way to save records, see
// [modifyRecords(saving:deleting:savePolicy:atomically:)].
//
// See: https://developer.apple.com/documentation/CloudKit/CKDatabase/save(_:completionHandler:)-3tatz
//
// [modifyRecords(saving:deleting:savePolicy:atomically:)]: https://developer.apple.com/documentation/CloudKit/CKDatabase/modifyRecords(saving:deleting:savePolicy:atomically:)
func (c CKDatabase) SaveRecordCompletionHandler(record ICKRecord, completionHandler CKRecordErrorHandler) {
	_block1, _ := NewCKRecordErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("saveRecord:completionHandler:"), record, _block1)
}

// Deletes a specific record.
//
// recordID: The identifier of the record to delete.
//
// completionHandler: The closure to execute after CloudKit deletes the record.
//
// # Discussion
//
// The completion handler takes the following parameters:
//
// - The identifier of the deleted record, or `nil` if there’s an error. -
// An error if a problem occurs, or `nil` if CloudKit successfully deletes the
// record.
//
// Deleting a record may cause additional deletions if other records in the
// database reference the deleted record. CloudKit doesn’t provide the
// identifiers of any additional records it deletes.
//
// For information on a more convenient way to delete records, see
// [modifyRecords(saving:deleting:savePolicy:atomically:)].
//
// See: https://developer.apple.com/documentation/CloudKit/CKDatabase/delete(withRecordID:completionHandler:)
//
// [modifyRecords(saving:deleting:savePolicy:atomically:)]: https://developer.apple.com/documentation/CloudKit/CKDatabase/modifyRecords(saving:deleting:savePolicy:atomically:)
func (c CKDatabase) DeleteRecordWithIDCompletionHandler(recordID ICKRecordID, completionHandler CKRecordIDErrorHandler) {
	_block1, _ := NewCKRecordIDErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("deleteRecordWithID:completionHandler:"), recordID, _block1)
}

// Fetches a specific record zone.
//
// zoneID: The identifier of the record zone to fetch.
//
// completionHandler: The closure to execute with the fetch results.
//
// # Discussion
//
// The completion handler takes the following parameters:
//
// - The fetched record zone, or `nil` if there’s an error. - An error if a
// problem occurs, or `nil` if CloudKit successfully fetches the specified
// record zone.
//
// For information on a more convenient way to fetch specific record zones,
// see [recordZones(for:)] in Swift or [CKFetchRecordZonesOperation] in
// Objective-C.
//
// See: https://developer.apple.com/documentation/CloudKit/CKDatabase/fetch(withRecordZoneID:completionHandler:)
//
// [recordZones(for:)]: https://developer.apple.com/documentation/CloudKit/CKDatabase/recordZones(for:)
func (c CKDatabase) FetchRecordZoneWithIDCompletionHandler(zoneID ICKRecordZoneID, completionHandler CKRecordZoneErrorHandler) {
	_block1, _ := NewCKRecordZoneErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("fetchRecordZoneWithID:completionHandler:"), zoneID, _block1)
}

// Saves a specific record zone.
//
// zone: The record zone to save.
//
// completionHandler: The closure to execute after CloudKit saves the record.
//
// # Discussion
//
// The completion handler takes the following parameters:
//
// - The saved record zone (as it appears on the server), or `nil` if
// there’s an error. - An error if a problem occurs, or `nil` if CloudKit
// successfully saves the record zone.
//
// For information on a more convenient way to save record zones, see
// [modifyRecordZones(saving:deleting:)].
//
// See: https://developer.apple.com/documentation/CloudKit/CKDatabase/save(_:completionHandler:)-32ffr
//
// [modifyRecordZones(saving:deleting:)]: https://developer.apple.com/documentation/CloudKit/CKDatabase/modifyRecordZones(saving:deleting:)
func (c CKDatabase) SaveRecordZoneCompletionHandler(zone ICKRecordZone, completionHandler CKRecordZoneErrorHandler) {
	_block1, _ := NewCKRecordZoneErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("saveRecordZone:completionHandler:"), zone, _block1)
}

// Deletes a specific record zone.
//
// zoneID: The identifier of the record zone to delete.
//
// completionHandler: The closure to execute after CloudKit deletes the record zone.
//
// # Discussion
//
// The completion handler takes the following parameters:
//
// - The identifier of the deleted record zone, or `nil` if there’s an
// error. - An error if a problem occurs, or `nil` if CloudKit successfully
// deletes the record zone.
//
// For information on a more convenient way to delete record zones, see
// [modifyRecordZones(saving:deleting:)].
//
// See: https://developer.apple.com/documentation/CloudKit/CKDatabase/delete(withRecordZoneID:completionHandler:)
//
// [modifyRecordZones(saving:deleting:)]: https://developer.apple.com/documentation/CloudKit/CKDatabase/modifyRecordZones(saving:deleting:)
func (c CKDatabase) DeleteRecordZoneWithIDCompletionHandler(zoneID ICKRecordZoneID, completionHandler CKRecordZoneIDErrorHandler) {
	_block1, _ := NewCKRecordZoneIDErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("deleteRecordZoneWithID:completionHandler:"), zoneID, _block1)
}

// Saves a specific subscription.
//
// subscription: The subscription to save.
//
// completionHandler: The closure to execute after CloudKit saves the subscription.
//
// # Discussion
//
// The completion handler takes the following parameters:
//
// - The saved subscription (as it appears on the server), or `nil` if
// there’s an error. - An error if a problem occurs, or `nil` if CloudKit
// successfully saves the subscription.
//
// For information on a more convenient way to save subscriptions, see
// [modifySubscriptions(saving:deleting:)].
//
// See: https://developer.apple.com/documentation/CloudKit/CKDatabase/save(_:completionHandler:)-9pona
//
// [modifySubscriptions(saving:deleting:)]: https://developer.apple.com/documentation/CloudKit/CKDatabase/modifySubscriptions(saving:deleting:)
func (c CKDatabase) SaveSubscriptionCompletionHandler(subscription ICKSubscription, completionHandler CKSubscriptionErrorHandler) {
	_block1, _ := NewCKSubscriptionErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("saveSubscription:completionHandler:"), subscription, _block1)
}

// Executes the specified operation in the current database.
//
// operation: The operation to execute.
//
// # Discussion
//
// Configure the operation fully before you call this method. Prior to the
// operation executing, CloudKit sets its [Database] property to the current
// database. The operation executes at the priority and quality of service
// (QoS) that you specify using the [queuePriority] and [QualityOfService]
// properties.
//
// See: https://developer.apple.com/documentation/CloudKit/CKDatabase/add(_:)
//
// [queuePriority]: https://developer.apple.com/documentation/Foundation/Operation/queuePriority-swift.property
func (c CKDatabase) AddOperation(operation ICKDatabaseOperation) {
	objc.Send[objc.ID](c.ID, objc.Sel("addOperation:"), operation)
}

// The type of database.
//
// # Discussion
//
// For possible values, see [CKDatabase.Scope].
//
// See: https://developer.apple.com/documentation/CloudKit/CKDatabase/databaseScope
//
// [CKDatabase.Scope]: https://developer.apple.com/documentation/CloudKit/CKDatabase/Scope
func (c CKDatabase) DatabaseScope() CKDatabaseScope {
	rv := objc.Send[CKDatabaseScope](c.ID, objc.Sel("databaseScope"))
	return CKDatabaseScope(rv)
}

// The priority that the system uses when it allocates resources to the
// operations that use this configuration.
//
// See: https://developer.apple.com/documentation/cloudkit/ckoperation/configuration-swift.class/qualityofservice
func (c CKDatabase) QualityOfService() foundation.NSQualityOfService {
	rv := objc.Send[foundation.NSQualityOfService](c.ID, objc.Sel("qualityOfService"))
	return foundation.NSQualityOfService(rv)
}
func (c CKDatabase) SetQualityOfService(value foundation.NSQualityOfService) {
	objc.Send[struct{}](c.ID, objc.Sel("setQualityOfService:"), value)
}

// FetchRecordWithID is a synchronous wrapper around [CKDatabase.FetchRecordWithIDCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CKDatabase) FetchRecordWithID(ctx context.Context, recordID ICKRecordID) (*CKRecord, error) {
	type result struct {
		val *CKRecord
		err error
	}
	done := make(chan result, 1)
	c.FetchRecordWithIDCompletionHandler(recordID, func(val *CKRecord, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// SaveRecord is a synchronous wrapper around [CKDatabase.SaveRecordCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CKDatabase) SaveRecord(ctx context.Context, record ICKRecord) (*CKRecord, error) {
	type result struct {
		val *CKRecord
		err error
	}
	done := make(chan result, 1)
	c.SaveRecordCompletionHandler(record, func(val *CKRecord, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// DeleteRecordWithID is a synchronous wrapper around [CKDatabase.DeleteRecordWithIDCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CKDatabase) DeleteRecordWithID(ctx context.Context, recordID ICKRecordID) (*CKRecordID, error) {
	type result struct {
		val *CKRecordID
		err error
	}
	done := make(chan result, 1)
	c.DeleteRecordWithIDCompletionHandler(recordID, func(val *CKRecordID, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// FetchRecordZoneWithID is a synchronous wrapper around [CKDatabase.FetchRecordZoneWithIDCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CKDatabase) FetchRecordZoneWithID(ctx context.Context, zoneID ICKRecordZoneID) (*CKRecordZone, error) {
	type result struct {
		val *CKRecordZone
		err error
	}
	done := make(chan result, 1)
	c.FetchRecordZoneWithIDCompletionHandler(zoneID, func(val *CKRecordZone, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// SaveRecordZone is a synchronous wrapper around [CKDatabase.SaveRecordZoneCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CKDatabase) SaveRecordZone(ctx context.Context, zone ICKRecordZone) (*CKRecordZone, error) {
	type result struct {
		val *CKRecordZone
		err error
	}
	done := make(chan result, 1)
	c.SaveRecordZoneCompletionHandler(zone, func(val *CKRecordZone, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// DeleteRecordZoneWithID is a synchronous wrapper around [CKDatabase.DeleteRecordZoneWithIDCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CKDatabase) DeleteRecordZoneWithID(ctx context.Context, zoneID ICKRecordZoneID) (*CKRecordZoneID, error) {
	type result struct {
		val *CKRecordZoneID
		err error
	}
	done := make(chan result, 1)
	c.DeleteRecordZoneWithIDCompletionHandler(zoneID, func(val *CKRecordZoneID, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// SaveSubscription is a synchronous wrapper around [CKDatabase.SaveSubscriptionCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CKDatabase) SaveSubscription(ctx context.Context, subscription ICKSubscription) (*CKSubscription, error) {
	type result struct {
		val *CKSubscription
		err error
	}
	done := make(chan result, 1)
	c.SaveSubscriptionCompletionHandler(subscription, func(val *CKSubscription, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
