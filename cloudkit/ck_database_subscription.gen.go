// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CKDatabaseSubscription] class.
var (
	_CKDatabaseSubscriptionClass     CKDatabaseSubscriptionClass
	_CKDatabaseSubscriptionClassOnce sync.Once
)

func getCKDatabaseSubscriptionClass() CKDatabaseSubscriptionClass {
	_CKDatabaseSubscriptionClassOnce.Do(func() {
		_CKDatabaseSubscriptionClass = CKDatabaseSubscriptionClass{class: objc.GetClass("CKDatabaseSubscription")}
	})
	return _CKDatabaseSubscriptionClass
}

// GetCKDatabaseSubscriptionClass returns the class object for CKDatabaseSubscription.
func GetCKDatabaseSubscriptionClass() CKDatabaseSubscriptionClass {
	return getCKDatabaseSubscriptionClass()
}

type CKDatabaseSubscriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKDatabaseSubscriptionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKDatabaseSubscriptionClass) Alloc() CKDatabaseSubscription {
	rv := objc.Send[CKDatabaseSubscription](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A subscription that generates push notifications when CloudKit modifies
// records in a database.
//
// # Overview
//
// Subscriptions track the creation, modification, and deletion of records in
// a database, and are fundamental in keeping data on the user’s device up
// to date. A subscription applies only to the user that creates it. When a
// subscription registers a change, such as CloudKit saving a new record, it
// sends push notifications to the user’s devices to inform your app about
// the change. You can then fetch the changes and cache them on-device. When
// appropriate, the server excludes the device where the change originates.
//
// A database subscription executes whenever a change occurs in a custom
// record zone that resides in the database where you save the subscription.
// This is important for the shared database because you don’t know what
// record zones exist in advance. The only exception to this is the default
// record zone in the user’s private database, which doesn’t participate
// in database subscriptions.
//
// You can further specialize a database subscription by setting its
// [CKDatabaseSubscription.RecordType] property to a specific record type. This limits the scope of
// the subscription to only track changes to records of that type and reduces
// the number of notifications it generates.
//
// Create any subscriptions on your app’s first launch. After you initialize
// a subscription, save it to the server using
// [CKModifySubscriptionsOperation]. After the operation completes, record
// that state on-device (in [UserDefaults], for example). You can then check
// that state on subsequent launches to prevent unnecessary trips to the
// server.
//
// To configure the notification that the subscription generates, set the
// subscription’s [CKDatabaseSubscription.NotificationInfo] property. Because the system coalesces
// notifications, don’t rely on them for specific changes. CloudKit can omit
// data to keep the payload size under the APNs size limit. Consider
// notifications an indication of remote changes, and use
// [CKFetchDatabaseChangesOperation] to fetch the record zones that contain
// those changes. After you have the record zones, use
// [CKFetchRecordZoneChangesOperation] to fetch the changed records in each
// zone. Server change tokens allow you to limit the fetch results to just the
// changes since your previous fetch.
//
// The example below shows how to create a database subscription in the
// user’s private database, configure the notifications it generates — in
// this case, silent push notifications — and then save that subscription to
// the server:
//
// # Accessing the Subscription Metadata
//
//   - [CKDatabaseSubscription.RecordType]: The type of record that the subscription queries.
//   - [CKDatabaseSubscription.SetRecordType]
//
// See: https://developer.apple.com/documentation/CloudKit/CKDatabaseSubscription
//
// [UserDefaults]: https://developer.apple.com/documentation/Foundation/UserDefaults
type CKDatabaseSubscription struct {
	CKSubscription
}

// CKDatabaseSubscriptionFromID constructs a [CKDatabaseSubscription] from an objc.ID.
//
// A subscription that generates push notifications when CloudKit modifies
// records in a database.
func CKDatabaseSubscriptionFromID(id objc.ID) CKDatabaseSubscription {
	return CKDatabaseSubscription{CKSubscription: CKSubscriptionFromID(id)}
}

// NOTE: CKDatabaseSubscription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKDatabaseSubscription] class.
//
// # Accessing the Subscription Metadata
//
//   - [ICKDatabaseSubscription.RecordType]: The type of record that the subscription queries.
//   - [ICKDatabaseSubscription.SetRecordType]
//
// See: https://developer.apple.com/documentation/CloudKit/CKDatabaseSubscription
type ICKDatabaseSubscription interface {
	ICKSubscription

	// Topic: Accessing the Subscription Metadata

	// The type of record that the subscription queries.
	RecordType() unsafe.Pointer
	SetRecordType(value unsafe.Pointer)
}

// Init initializes the instance.
func (c CKDatabaseSubscription) Init() CKDatabaseSubscription {
	rv := objc.Send[CKDatabaseSubscription](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKDatabaseSubscription) Autorelease() CKDatabaseSubscription {
	rv := objc.Send[CKDatabaseSubscription](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKDatabaseSubscription creates a new CKDatabaseSubscription instance.
func NewCKDatabaseSubscription() CKDatabaseSubscription {
	class := getCKDatabaseSubscriptionClass()
	rv := objc.Send[CKDatabaseSubscription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a database subscription from a serialized instance.
//
// aDecoder: The object that decodes the serialized database subscription.
//
// See: https://developer.apple.com/documentation/CloudKit/CKDatabaseSubscription/init(coder:)
func NewCKDatabaseSubscriptionWithCoder(aDecoder foundation.INSCoder) CKDatabaseSubscription {
	instance := getCKDatabaseSubscriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return CKDatabaseSubscriptionFromID(rv)
}

// The type of record that the subscription queries.
//
// See: https://developer.apple.com/documentation/cloudkit/ckdatabasesubscription/recordtype-46v7a
func (c CKDatabaseSubscription) RecordType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("recordType"))
	return rv
}
func (c CKDatabaseSubscription) SetRecordType(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setRecordType:"), value)
}
